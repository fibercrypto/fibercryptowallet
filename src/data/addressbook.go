package data

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/sirupsen/logrus"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	"github.com/skycoin/skycoin/src/util/file"
	"golang.org/x/crypto/bcrypt"
	"io"
	"time"
)

const (
	Type1 = iota //  No security
	Type2        //  data obfuscation security
	Type3        //  password security
)

var (
	// Errors
	errBucketEmpty    = errors.New(" Error: bucket are empty")
	errValEmpty       = errors.New(" Error: value are empty")
	errParseContact   = errors.New(" The inserted cannot be parse")
	errInvalidSecType = fmt.Errorf("invalid security type")
	// Db buckets.
	dbAddrsBookBkt = []byte("AddressBook")
	dbConfigBkt    = []byte("Config")
)

// DB implement AddressBook interface for boltdb database.
type DB struct {
	db  *bolt.DB // db is a point of bolt.DB object. This handle all interaction with the db.
	key []byte
}

// Create a new instance of AddessBook and open the database of the given route.
// If database is open return bolt.errDatabaseOpen.
func NewAddressBook(path string) (core.AddressBook, error) {
	var ab DB
	var err error
	// Open database.
	db, err := bolt.Open(path,
		0600,
		&bolt.Options{
			Timeout: 1 * time.Second,
		})
	if err != nil {
		return nil, errorWrapper(fmt.Errorf(" error opening data base: %s", err))
	}
	ab.db = db
	return &ab, nil
}

// Init initialize an address book. Pass secType(security type) and password if is Type3.
func (ab *DB) Init(secType int, password string) error {
	if !ab.IsOpen() {
		return errorWrapper(bolt.ErrDatabaseNotOpen)
	}

	if ab.HasInit() {
		return errorWrapper(fmt.Errorf("address book has init"))
	}

	var hash, entropy []byte
	var err error

	switch secType {
	case Type1, Type2:
		if err := ab.insertConfig(secType, hash, entropy); err != nil {
			return errorWrapper(err)
		}
		break
	case Type3:
		ab.key = []byte(password)
		if entropy, err = ab.genEntropy(); err != nil {
			return errorWrapper(err)
		}

		hash, err = bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			return errorWrapper(err)
		}

		if err := ab.insertConfig(secType, hash, entropy); err != nil {
			return errorWrapper(err)
		}
		break
	default:
		return errorWrapper(errInvalidSecType)
	}

	return nil
}

// Authenticate authentic a user in the Address Book. ( Only SecType : Type3 )
func (ab *DB) Authenticate(password string) error {
	if !ab.IsOpen() {
		return errorWrapper(bolt.ErrDatabaseNotOpen)
	}

	if !ab.HasInit() {
		return errorWrapper(fmt.Errorf("address book not has init"))
	}

	secType, err := ab.getSecTypeFromConfig()
	if err != nil {
		return errorWrapper(err)
	}

	if secType != Type3 {
		return nil
	}

	ab.key = []byte(password)
	if err := ab.verifyHash(); err != nil {
		return errorWrapper(err)
	}

	return nil
}

// InsertContact insert a contact into the address book.
// If any of its address exist return error.
func (ab *DB) InsertContact(c core.Contact) (uint64, error) {
	// Start a writeable transaction.
	tx, err := ab.db.Begin(true)
	if err != nil {
		return 0, errorWrapper(err)
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			logrus.Fatal(err)
		}
	}()

	contacts, err := ab.ListContact()
	if err != nil {
		return 0, err
	}
	for _, v := range c.GetAddresses() {
		if err := ab.addressExists(v, contacts); err != nil {
			return 0, errorWrapper(err)
		}
	}
	if err := ab.nameExists(c, contacts); err != nil {
		return 0, errorWrapper(err)
	}

	bkt := tx.Bucket(dbAddrsBookBkt)
	if bkt == nil {
		return 0, errBucketEmpty
	}

	// The sequence is an auto-incrementing integer that is transactionally safe.
	seq, err := bkt.NextSequence()
	c.SetID(seq)

	if cc, ok := c.(*Contact); ok {
		encryptedData, err := ab.encryptContact(cc)
		if err != nil {
			return 0, errorWrapper(err)
		}

		// Save contact to the bucket.
		if err := bkt.Put(itob(c.GetID()), encryptedData); err != nil {
			return 0, errorWrapper(err)
		}
	} else {
		return 0, errParseContact
	}

	// Commit transaction before exit.
	return c.GetID(), tx.Commit()
}

// GetContact get a contact by Id.
func (ab *DB) GetContact(id uint64) (core.Contact, error) {
	// Start a readable transaction.
	tx, err := ab.db.Begin(false)
	if err != nil {
		return nil, errorWrapper(err)
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()

	bkt := tx.Bucket(dbAddrsBookBkt)
	if bkt == nil {
		return nil, errorWrapper(errBucketEmpty)
	}

	// Read encoded contact bytes.
	encryptData := bkt.Get(itob(id))
	if encryptData == nil {
		return nil, errorWrapper(errValEmpty)
	}
	c, err := ab.decryptContact(encryptData)
	if err != nil {
		return nil, errorWrapper(err)
	}
	return c, nil
}

// ListContact list all contact in the address book.
func (ab *DB) ListContact() ([]core.Contact, error) {
	// Start a redeable transaction.
	tx, err := ab.db.Begin(false)
	if err != nil {
		return nil, errorWrapper(err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()

	bkt := tx.Bucket(dbAddrsBookBkt)
	if bkt == nil {
		return nil, errorWrapper(errBucketEmpty)
	}
	var contacts []core.Contact
	if err := bkt.ForEach(func(k, v []byte) error {
		c, err := ab.decryptContact(v)
		if err != nil {
			return errorWrapper(err)
		}

		contacts = append(contacts, c)
		return nil
	}); err != nil {
		return nil, errorWrapper(err)
	}
	return contacts, nil
}

// DeleteContact delete a contact from the address book by its Id.
func (ab *DB) DeleteContact(id uint64) error {
	return ab.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(dbAddrsBookBkt)
		if bkt == nil {
			return errorWrapper(errBucketEmpty)
		}
		if val := bkt.Get(itob(id)); val == nil {
			return errorWrapper(errValEmpty)
		}

		return bkt.Delete(itob(id))
	})

}

// UpdateContact update a contact in the address book by its Id.
func (ab *DB) UpdateContact(id uint64, newContact core.Contact) error {
	var contacts []core.Contact
	var err error
	if contacts, err = ab.ListContact(); err != nil {
		return errorWrapper(err)
	}
	for e := range contacts {
		if contacts[e].GetID() == id {
			contacts[e] = nil
			break
		}
	}

	for _, ncAddrs := range newContact.GetAddresses() {
		if err := ab.addressExists(ncAddrs, contacts); err != nil {
			return errorWrapper(err)
		}
	}
	if err := ab.nameExists(newContact, contacts); err != nil {
		return errorWrapper(err)
	}

	if err := ab.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(dbAddrsBookBkt)
		if bkt == nil {
			return errorWrapper(errBucketEmpty)
		}
		if cc, ok := newContact.(*Contact); ok {
			cc.SetID(id)
			encryptedData, err := ab.encryptContact(cc)
			if err != nil {
				return errorWrapper(err)
			}
			if err := bkt.Put(itob(id), encryptedData); err != nil {
				return errorWrapper(err)
			}
		} else {
			return errorWrapper(errParseContact)
		}

		return nil
	}); err != nil {
		return errorWrapper(err)
	}
	return nil
}

func (ab *DB) GetPath() string {
	return ab.db.Path()
}

func (ab *DB) GetSecType() int {
	if ab.db != nil {
		secType, _ := ab.getSecTypeFromConfig()
		return secType
	}
	return -1
}

// Close shuts down the database.
func (ab *DB) Close() error {
	if err := ab.db.Close(); err != nil {
		return errorWrapper(err)
	}
	return nil
}

func (ab *DB) HasInit() bool {
	tx, _ := ab.db.Begin(false)
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()
	if tx.Bucket(dbConfigBkt) != nil {
		return true
	}
	return false
}

func (ab *DB) IsOpen() bool {
	if ab.db.Path() != "" {
		return true
	}
	return false
}

func (ab *DB) Exist(path string) bool {
	ok, _ := file.Exists(path)
	return ok
}

func (ab *DB) getSecTypeFromConfig() (int, error) {
	tx, err := ab.db.Begin(false)
	if err != nil {
		return -1, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()
	buck := tx.Bucket(dbConfigBkt)
	if buck == nil {
		return -1, errBucketEmpty
	}
	return int(btoi(buck.Get([]byte("secType")))), nil
}

// genEntropy generate annil Entropy by a mnemonic. If mnemonic is nil,
// it generate a random.
func (ab *DB) genEntropy() ([]byte, error) {
	mn, err := bip39.NewDefaultMnemonic()
	if err != nil {
		return nil, err
	}
	e, err := bip39.EntropyFromMnemonic(mn)
	if err != nil {
		return nil, err
	}
	return e, err
}

func (ab *DB) getEntropyFromConfig() ([]byte, error) {
	tx, err := ab.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()
	buck := tx.Bucket(dbConfigBkt)
	if buck == nil {
		return nil, errBucketEmpty
	}
	return buck.Get([]byte("entropy")), nil
}

// getHashFromConfig get hash from config bucket.
func (ab *DB) getHashFromConfig() ([]byte, error) {
	tx, err := ab.db.Begin(false)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()
	buck := tx.Bucket(dbConfigBkt)
	if buck == nil {
		return nil, errBucketEmpty
	}

	hash := buck.Get([]byte("hash"))
	if hash == nil {
		return nil, errValEmpty
	}

	return hash, nil
}

// encryptContact encrypt a contact using a password with AES-GCM.
func (ab *DB) encryptContact(c *Contact) ([]byte, error) {
	secType, err := ab.getSecTypeFromConfig()
	if err != nil {
		return nil, err
	}
	switch secType {
	case Type1:
		return c.MarshalBinary()
	case Type2:
		data, err := c.MarshalBinary()
		if err != nil {
			return nil, err
		}
		return toBase64(data), nil
	case Type3:
		return ab.encryptAESGCM(c)
	}

	return nil, fmt.Errorf("invalid security type")
}

// Decrypt a cipher message using a password with AES-GCM and return a Contact.
func (ab *DB) decryptContact(cipherMsg []byte) (core.Contact, error) {
	secType, err := ab.getSecTypeFromConfig()
	if err != nil {
		return nil, err
	}
	switch secType {
	case Type1:
		c := Contact{}
		if err := c.UnmarshalBinary(cipherMsg); err != nil {
			return nil, err
		} else {
			return &c, nil
		}
	case Type2:
		c := Contact{}
		data, err := fromBase64(cipherMsg)
		if err != nil {
			return nil, err
		}
		if err := c.UnmarshalBinary(data); err != nil {
			return nil, err
		}
		return &c, nil
	case Type3:
		return ab.decryptAESGCM(cipherMsg)
	}

	return nil, errInvalidSecType
}

//
func (ab *DB) verifyHash() error {
	hash, err := ab.getHashFromConfig()
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hash, ab.key)
}

// addressExists search an address in the list of contacts into the AddressBook.
// If find the address return error, else return nil.
func (ab *DB) addressExists(address core.StringAddress, contacts []core.Contact) error {
	for _, v := range contacts {
		c, ok := v.(*Contact)
		if ok {
			for _, addrs := range c.Address {
				if bytes.Compare(addrs.GetValue(), address.GetValue()) == 0 &&
					bytes.Compare(addrs.GetCoinType(), address.GetCoinType()) == 0 {
					return fmt.Errorf("Address with value: %s  and Cointype: %s alredy exist",
						address.GetValue(), address.GetCoinType())
				}
			}
		}
	}

	return nil
}

// nameExists search an name in the list of contacts into the AddressBook.
// If find the address return error, else return nil.
func (ab *DB) nameExists(contact core.Contact, contacts []core.Contact) error {
	for _, c := range contacts {
		if dataContact, ok := c.(*Contact); ok {
			if bytes.Compare(contact.(*Contact).Name, dataContact.Name) == 0 {
				return fmt.Errorf(" Contact with name: %s alredy exist", contact.(*Contact).Name)
			}
		}
	}

	return nil
}

func (ab *DB) insertConfig(secType int, hash, entropy []byte) error {
	tx, err := ab.db.Begin(true)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			logrus.Fatal(err)
		}
	}()

	_, err = tx.CreateBucketIfNotExists(dbAddrsBookBkt)
	if err != nil {
		return err
	}
	_, err = tx.CreateBucketIfNotExists(dbConfigBkt)
	if err != nil {
		return err
	}

	bConf := tx.Bucket(dbConfigBkt)
	if err := bConf.Put([]byte("hash"), hash); err != nil {
		return err
	}
	if err := bConf.Put([]byte("entropy"), entropy); err != nil {
		return err
	}
	if err := bConf.Put([]byte("secType"), itob(uint64(secType))); err != nil {
		return err
	}

	return tx.Commit()
}

func (ab *DB) encryptAESGCM(c *Contact) ([]byte, error) {
	e, err := ab.getEntropyFromConfig()
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(derivePassphrase(e, ab.key))
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	bc, err := c.MarshalBinary()
	if err != nil {
		return nil, err
	}

	cipherText := aesGCM.Seal(nonce, nonce, bc, nil)
	return cipherText, nil
}

func (ab *DB) decryptAESGCM(cipherMsg []byte) (core.Contact, error) {
	e, err := ab.getEntropyFromConfig()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(derivePassphrase(e, ab.key))
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	var c Contact
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := cipherMsg[:nonceSize], cipherMsg[nonceSize:]

	data, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	if err := c.UnmarshalBinary(data); err != nil {
		return nil, err
	}
	return &c, nil
}
