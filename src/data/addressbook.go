package data

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/SkycoinProject/skycoin/src/cipher/pbkdf2"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

const (
	// NoSecurity  No security
	NoSecurity = iota
	// ObfuscationSecurity data obfuscation security
	ObfuscationSecurity
	// PasswordSecurity password security
	PasswordSecurity
)

const (
	hash         = "hash"
	entropy      = "entropy"
	securityType = "secType"
	bCryptCost   = 12
)

var (
	// Errors
	errInvalidContact      = errors.New("you try to inserted a invalid contact")
	errInvalidSecType      = errors.New("invalid security type")
	errAddrsBookHasNotInit = errors.New("address book not has init")
)

// addrsBook implement AddressBook interface for boltdb database.
type addrsBook struct {
	storage core.Storage
	key     []byte
}

var logDb = logging.MustGetLogger("AddressBook Data")

// NewAddressBook create a new instance of AddessBook.
func NewAddressBook(storage core.Storage) core.AddressBook {
	return &addrsBook{
		storage: storage,
		key:     nil,
	}
}

// ChangeSecurity change the security type of the Address Book, this method work to change the password too.
func (addrsBook *addrsBook) ChangeSecurity(NewSecType int, oldPassword, newPassword string) error {
	logDb.Info("changing Address Book security")

	if err := addrsBook.Authenticate(oldPassword); err != nil {
		return err
	}

	contactsList, err := addrsBook.ListContact()
	if err != nil && err != errBucketEmpty {
		return err
	}

	switch NewSecType {
	case NoSecurity, ObfuscationSecurity:
		if err := addrsBook.insertConfig(NewSecType, nil, nil); err != nil {
			logDb.Error(err)
			return err
		}
		break
	case PasswordSecurity:
		var hash, entropy []byte
		if entropy, err = genEntropy(); err != nil {
			logDb.Error(err)
			return err
		}

		hash, err = bcrypt.GenerateFromPassword([]byte(newPassword), bCryptCost)
		if err != nil {
			logDb.Error(err)
			return err
		}

		if err := addrsBook.insertConfig(NewSecType, hash, entropy); err != nil {
			logDb.Error(err)
			return err
		}

		addrsBook.key = pbkdf2.Key([]byte(newPassword), entropy, 4096, 32, sha512.New)
		break
	default:
		logDb.Error(errInvalidSecType)
		return errInvalidSecType
	}
	for e := range contactsList {
		if err := addrsBook.UpdateContact(contactsList[e].GetID(), contactsList[e]); err != nil {
			return err
		}
	}

	return nil
}

// Init initialize an address book. Pass secType(security type) and password if is PasswordSecurity.
func (addrsBook *addrsBook) Init(secType int, password string) error {
	logDb.Info("initialize AddressBook")
	if !addrsBook.IsOpen() {
		return errDatabaseNotOpen
	}

	if addrsBook.HasInit() {
		return fmt.Errorf("address book has init")
	}

	var hash, entropy []byte
	var err error

	switch secType {
	case NoSecurity, ObfuscationSecurity:
		if err := addrsBook.insertConfig(secType, hash, entropy); err != nil {
			logDb.Error(err)
			return err
		}
		break
	case PasswordSecurity:
		if entropy, err = genEntropy(); err != nil {
			logDb.Error(err)
			return err
		}

		hash, err = bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			logDb.Error(err)
			return err
		}

		if err := addrsBook.insertConfig(secType, hash, entropy); err != nil {
			logDb.Error(err)
			return err
		}

		addrsBook.key = pbkdf2.Key([]byte(password), entropy, 4096, 32, sha512.New)
		break
	default:
		logDb.Error(errInvalidSecType)
		return errInvalidSecType
	}

	return nil
}

// Authenticate authentic a user in the Address Book. ( Only SecType : PasswordSecurity )
func (addrsBook *addrsBook) Authenticate(password string) error {
	logDb.Info("authenticate AddressBook")
	if !addrsBook.IsOpen() {
		logDb.Error(errDatabaseNotOpen)
		return errDatabaseNotOpen
	}

	if !addrsBook.HasInit() {
		logDb.Error(errAddrsBookHasNotInit)
		return errAddrsBookHasNotInit
	}

	if addrsBook.GetSecType() == PasswordSecurity {
		if err := bcrypt.CompareHashAndPassword(addrsBook.getHashFromConfig(), []byte(password)); err != nil {
			logDb.Error(err)
			return err
		}
		addrsBook.key = pbkdf2.Key([]byte(password), addrsBook.getEntropyFromConfig(), 4096, 32, sha512.New)
	}
	return nil
}

// InsertContact insert a contact into the address book.
// If any of its address exist return error.
func (addrsBook *addrsBook) InsertContact(contact core.Contact) (uint64, error) {
	logDb.Info("Inserting a contact into the Address Book.")
	if !addrsBook.validateContact(contact) {
		return 0, errInvalidContact
	}

	encryptedData, err := contact.Encrypt(addrsBook.GetSecType(), addrsBook.key)
	if err != nil {
		return 0, err
	}

	// Commit transaction before exit.
	return addrsBook.GetStorage().InsertValue(encryptedData)
}

// GetContact get a contact by ID.
func (addrsBook *addrsBook) GetContact(id uint64) (core.Contact, error) {
	logDb.Info("Getting a contact from the Address Book")
	encryptData, err := addrsBook.GetStorage().GetValue(id)
	if err != nil {
		logDb.Error(err)
		return nil, err
	}
	if _, ok := encryptData.([]byte); !ok {
		logDb.Error(errValueNoMatch(encryptData.([]byte), []byte{}))
		return nil, errValueNoMatch(encryptData.([]byte), []byte{})
	}
	var contact Contact

	if err := contact.Decrypt(addrsBook.GetSecType(), encryptData.([]byte), addrsBook.key); err != nil {
		return nil, err
	}
	contact.SetID(id)
	return &contact, nil
}

// ListContact list all contact in the address book.
func (addrsBook *addrsBook) ListContact() ([]core.Contact, error) {
	logDb.Info("Getting all contact in the Address Book.")
	var contactsList []core.Contact
	encryptContactList, err := addrsBook.GetStorage().ListValues()
	if err != nil {
		logDb.Error(err)
		return nil, err
	}
	for id, encryptContact := range encryptContactList {
		if _, ok := encryptContact.([]byte); !ok {
			return nil, errValueNoMatch(encryptContact, []byte{})
		}
		var contact Contact
		if err := contact.Decrypt(addrsBook.GetSecType(), encryptContact.([]byte), addrsBook.key); err != nil {
			return nil, err
		}
		contact.SetID(id)
		contactsList = append(contactsList, &contact)
	}
	return contactsList, nil
}

// DeleteContact delete a contact from the address book by its ID.
func (addrsBook *addrsBook) DeleteContact(id uint64) error {
	logDb.Info("Removing a contact from AddressBook")
	return addrsBook.GetStorage().DeleteValue(id)
}

// UpdateContact update a contact in the address book by its ID.
func (addrsBook *addrsBook) UpdateContact(id uint64, newContact core.Contact) error {
	logDb.Infof("Updating contact with id:%d", id)
	if err := addrsBook.DeleteContact(id); err != nil {
		return err
	}

	if !addrsBook.validateContact(newContact) {
		return errInvalidContact
	}
	encryptedData, err := newContact.Encrypt(addrsBook.GetSecType(), addrsBook.key)
	if err != nil {
		return err
	}
	return addrsBook.GetStorage().UpdateValue(id, encryptedData)
}

// GetPath return database path
func (addrsBook *addrsBook) GetPath() string {
	return addrsBook.GetStorage().Path()
}

// Close shuts down the database.
func (addrsBook *addrsBook) Close() error {
	logDb.Info("Closing the AddressBook")
	if err := addrsBook.GetStorage().Close(); err != nil {
		return err
	}
	return nil
}

// HasInit verify if database has been initialize.
func (addrsBook *addrsBook) HasInit() bool {
	if addrsBook.storage.GetConfig() != nil {
		return true
	}
	return false
}

// IsOpen verify if database is open.
func (addrsBook *addrsBook) IsOpen() bool {
	if addrsBook.storage.Path() != "" {
		return true
	}
	return false
}

// GetStorage return the Storage parameter. This is in charge of manager the database interaction.
func (addrsBook *addrsBook) GetStorage() core.Storage {
	logDb.Info("Getting  Storage.")
	return addrsBook.storage
}

// GetSecType return the Security Type of the Address Book.
func (addrsBook *addrsBook) GetSecType() int {
	logDb.Info("Getting security type.")
	secType, err := strconv.Atoi(addrsBook.GetStorage().GetConfig()[securityType])
	if err != nil {
		logDb.Error(err)
		return -1
	}
	return secType
}

// getEntropyFromConfig retrieve the entropy parameter from the database.
func (addrsBook *addrsBook) getEntropyFromConfig() []byte {
	logDb.Info("Getting entropy.")
	return []byte(addrsBook.GetStorage().GetConfig()[entropy])

}

// getHashFromConfig get hash from config bucket in the database..
func (addrsBook *addrsBook) getHashFromConfig() []byte {
	logDb.Info("Getting hash.")
	return []byte(addrsBook.GetStorage().GetConfig()[hash])
}

func (addrsBook *addrsBook) validateContact(contact core.Contact) bool {
	if !contact.IsValid() {
		logDb.Error(errInvalidContact)
		return false
	}

	contactsList, err := addrsBook.ListContact()
	if err != nil && err != errBucketEmpty {
		logDb.Error(err)
		return false
	}

	if err := isRepeatName(contact, contactsList); err != nil {
		return false
	}

	for e := range contact.GetAddresses() {
		if err := isRepeatAddress(contact.GetAddresses()[e], contactsList); err != nil {
			return false
		}
	}
	return true
}

// insertConfig insert the configuration parameter into Config block in the database.
func (addrsBook *addrsBook) insertConfig(secType int, hashVal, entropyVal []byte) error {
	if err := addrsBook.GetStorage().InsertConfig(
		map[string]string{
			securityType: strconv.Itoa(secType),
			hash:         string(hashVal),
			entropy:      string(entropyVal)}); err != nil {
		logDb.Error(err)
		return err
	}
	return nil
}
