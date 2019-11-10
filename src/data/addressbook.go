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
	"golang.org/x/crypto/bcrypt"
	"io"
	"time"
)

//go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --gogo_out=. internal/contact.proto
//go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --gogo_out=. internal/config.proto

var (
	// Errors
	errBucketEmpty  = errors.New(" Error: bucket are empty")
	errValEmpty     = errors.New(" Error: value are empty")
	errParseContact = errors.New(" The inserted cannot be parse")
	// Db buckets.
	dbAddrsBookBkt = []byte("AddressBook")
	dbConfigBkt    = []byte("Config")
)

// DB implement AddressBook interface for boltdb database.
type DB struct {
	// db is a point of bolt.DB object. This handle all interaction with the db.
	db      *bolt.DB
	dbPath  string
	key     []byte
	entropy []byte
}

//
func (ab *DB) verifyHash() error {
	hash, err := ab.GetHashFromConfig()
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hash, ab.key)

}

// Init initialize an address book. Using this before do you use NewAddressBook.
// If the address book has been init, use LoadFromFile.
func Init(password []byte, path string) (*DB, error) {
	var ab DB
	ab.dbPath = path
	if err := ab.open(); err != nil {
		return nil, err
	}
	if err := ab.genEntropy(); err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword(password, 14)

	if err != nil {
		return nil, err
	}

	tx, err := ab.db.Begin(true)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			logrus.Fatal(err)
		}
	}()

	// Initialize buckets to guarantee that they exist.
	_, err = tx.CreateBucketIfNotExists(dbAddrsBookBkt)
	if err != nil {
		return nil, err
	}
	_, err = tx.CreateBucketIfNotExists(dbConfigBkt)
	if err != nil {
		return nil, err
	}
	bConf := tx.Bucket(dbConfigBkt)
	if err := bConf.Put([]byte("hash"), hash); err != nil {
		return nil, err
	}
	if err := bConf.Put([]byte("entropy"), ab.entropy); err != nil {
		return nil, err
	}
	ab.key = password
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &ab, nil
}

// LoadFromFile load a existing db file.
func LoadFromFile(path string, password []byte) (*DB, error) {
	var ab DB
	ab.dbPath = path
	ab.key = password
	if err := ab.open(); err != nil {
		return nil, err
	}
	tx, err := ab.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()
	bConf := tx.Bucket(dbConfigBkt)
	if bConf == nil {
		return nil, errBucketEmpty
	}

	ab.entropy = bConf.Get([]byte("entropy"))

	if err := ab.verifyHash(); err != nil {
		if err := ab.db.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}
	return &ab, nil
}

// open function open the database.
func (ab *DB) open() error {
	var err error
	// Open database.
	db, err := bolt.Open(ab.dbPath,
		0600,
		&bolt.Options{
			Timeout: 1 * time.Second,
		})
	if err != nil {
		return fmt.Errorf(" error opening data base: %s", err)
	}
	ab.db = db
	// Start a transaction.
	tx, err := ab.db.Begin(true)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			logrus.Fatal(err)
		}
	}()
	return tx.Commit()
}

// InsertContact insert a contact into the address book.
// If any of its address exist return error.
func (ab *DB) InsertContact(c core.Contact) error {
	// Start a writeable transaction.
	tx, err := ab.db.Begin(true)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			logrus.Fatal(err)
		}
	}()

	contacts, err := ab.ListContact()
	for _, v := range c.GetAddresses() {
		if err := ab.AddressExists(v, contacts); err != nil {
			return err
		}
	}
	if err := ab.NameExists(c, contacts); err != nil {
		return err
	}

	bkt := tx.Bucket(dbAddrsBookBkt)
	if bkt == nil {
		return errBucketEmpty
	}

	// The sequence is an auto-incrementing integer that is transactionally safe.
	seq, err := bkt.NextSequence()
	c.SetID(seq)

	if cc, ok := c.(*Contact); ok {
		encryptedData, err := ab.encryptAESGCM(cc)
		if err != nil {
			return err
		}

		// Save contact to the bucket.
		if err := bkt.Put(itob(c.GetID()), encryptedData); err != nil {
			return err
		}
	} else {
		return errParseContact
	}

	// Commit transaction before exit.
	return tx.Commit()
}

// GetContact get a contact by id.
func (ab *DB) GetContact(id uint64) (core.Contact, error) {
	// Start a redeable transaction.
	tx, err := ab.db.Begin(false)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()

	bkt := tx.Bucket(dbAddrsBookBkt)
	if bkt == nil {
		return nil, errBucketEmpty
	}

	// Read encoded contact bytes.
	encryptData := bkt.Get(itob(id))
	if encryptData == nil {
		return nil, errValEmpty
	}
	c, err := ab.decryptAESGCM(encryptData)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// ListContact list all contact in the address book.
func (ab *DB) ListContact() ([]core.Contact, error) {
	// Start a redeable transaction.
	tx, err := ab.db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()

	bkt := tx.Bucket(dbAddrsBookBkt)
	if bkt == nil {
		return nil, errBucketEmpty
	}
	var contacts []core.Contact
	if err := bkt.ForEach(func(k, v []byte) error {
		c, err := ab.decryptAESGCM(v)
		if err != nil {
			return err
		}

		contacts = append(contacts, c)
		return nil
	}); err != nil {
		return nil, err
	}
	return contacts, nil
}

// DeleteContact delete a contact from the address book by its id.
func (ab *DB) DeleteContact(id uint64) error {
	return ab.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(dbAddrsBookBkt)
		if bkt == nil {
			return errBucketEmpty
		}
		if val := bkt.Get(itob(id)); val == nil {
			return errValEmpty
		}
		return bkt.Delete(itob(id))
	})

}

// UpdateContact update a contact in the address book by its id.
func (ab *DB) UpdateContact(id uint64, newContact core.Contact) error {
	var contacts []core.Contact
	var err error
	if contacts, err = ab.ListContact(); err != nil {
		return err
	}
	for e := range contacts {
		if contacts[e].GetID() == id {
			contacts[e] = nil
			break
		}
	}

	for _, ncAddrs := range newContact.GetAddresses() {
		if err := ab.AddressExists(ncAddrs, contacts); err != nil {
			return err
		}
	}
	if err := ab.NameExists(newContact, contacts); err != nil {
		return err
	}

	if err := ab.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(dbAddrsBookBkt)
		if bkt == nil {
			return errBucketEmpty
		}
		if cc, ok := newContact.(*Contact); ok {
			encryptedData, err := ab.encryptAESGCM(cc)
			if err != nil {
				return err
			}
			if err := bkt.Put(itob(id), encryptedData); err != nil {
				return err
			}
		} else {
			return errParseContact
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

// Close shuts down the database.
func (ab *DB) Close() error {
	return ab.db.Close()
}

// genEntropy generate an Entropy by a mnemonic. If mnemonic is nil,
// it generate a random.
func (ab *DB) genEntropy() error {
	mn, err := bip39.NewDefaultMnemonic()
	if err != nil {
		return err
	}
	e, err := bip39.EntropyFromMnemonic(mn)
	if err != nil {
		return err
	}
	ab.entropy = e
	return nil
}

// SetPath set database path.
func (ab *DB) SetPath(path string) {
	ab.dbPath = path
}

// GetPath get database path.
func (ab *DB) GetPath() string {
	return ab.dbPath
}

// GetHashFromConfig get hash from config bucket.
func (ab *DB) GetHashFromConfig() ([]byte, error) {
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

// encryptAESGCM encrypt a contact using a password with AES-GCM.
func (ab *DB) encryptAESGCM(c *Contact) ([]byte, error) {
	if ab.entropy == nil {
		return nil, fmt.Errorf(" Error: Mnemonic are empty")
	}
	block, err := aes.NewCipher(derivePassphrase(ab.entropy, ab.key))
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

// Decrypt a cipher message using a password with AES-GCM and return a Contact.
func (ab *DB) decryptAESGCM(cipherMsg []byte) (core.Contact, error) {
	if ab.entropy == nil {
		return nil, fmt.Errorf(" Error: Mnemonic are empty")
	}

	block, err := aes.NewCipher(derivePassphrase(ab.entropy, ab.key))
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

// AddressExists search an address in the list of contacts into the AddressBook.
// If find the address return error, else return nil.
func (ab *DB) AddressExists(address core.ReadableAddress, contacts []core.Contact) error {
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

// NameExists search an name in the list of contacts into the AddressBook.
// If find the address return error, else return nil.
func (ab *DB) NameExists(contact core.Contact, contacts []core.Contact) error {
	for _, c := range contacts {
		if dataContact, ok := c.(*Contact); ok {
			if bytes.Compare(contact.(*Contact).Name, dataContact.Name) == 0 {
				return fmt.Errorf(" Contact with name: %s alredy exist", contact.(*Contact).Name)
			}
		}
	}

	return nil
}
