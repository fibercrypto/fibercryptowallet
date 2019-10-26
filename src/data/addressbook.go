package data

import (
	"./internal"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	"github.com/skycoin/skycoin/src/visor/dbutil"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --gogo_out=. internal/contact.proto
//go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --gogo_out=. internal/config.proto

var (
	errBucketEmpty  = errors.New(" Error: bucket are empty.")
	errValEmpty     = errors.New(" Error: value are empty.")
	dbConfigDataKey = []byte("ConfigData")
	dbAddrsBookBkt  = []byte("AddressBook")
	dbConfigBkt     = []byte("Config")
)

type AddressBook struct {
	// db is a point of bolt.DB object. This handle all interaction with the db.
	db          *bolt.DB
	dbPath      string // ".fiber/data.db"
	hasPassword bool
	hash        []byte
	entropy     []byte
}

//
func (ab *AddressBook) generateHash(password []byte) error {
	hash, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		return err
	}
	ab.hash = hash
	return nil
}

//
func (ab *AddressBook) verifyHash(password []byte) error {
	if ab.hash == nil {
		if err := ab.GetConfig(); err != nil {
			return err
		}
	}
	return bcrypt.CompareHashAndPassword(ab.hash, password)
}

// startUp function open the database.
func (ab *AddressBook) Open() error {
	var err error
	// Open database.
	db, err := bolt.Open(ab.dbPath,
		0666,
		&bolt.Options{
			Timeout: 1 * time.Second,
		})
	if err != nil {
		return fmt.Errorf(" Error opening data base: %s", err)
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

	// Initialize buckets to guarantee that they exist.
	_, err = tx.CreateBucketIfNotExists(dbAddrsBookBkt)
	if err != nil {
		return err
	}
	_, err = tx.CreateBucketIfNotExists(dbConfigBkt)
	if err != nil {
		return err
	}

	return tx.Commit()
}

//
func (ab *AddressBook) Insert(c *Contact, password []byte) error {
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

	bkt := tx.Bucket(dbAddrsBookBkt)
	if bkt == nil {
		return errBucketEmpty
	}

	// The sequence is an autoincrementing integer that is thansactionally safe.
	seq, err := bkt.NextSequence()
	c.ID = seq

	if err := ab.verifyHash(password); err != nil {
		return fmt.Errorf(" Error verify password: %s", err)
	}

	encryptedData, err := c.encryptContact(password, ab.entropy)
	if err != nil {
		return err
	}
	// Save contact to the bucket.
	if err := bkt.Put(dbutil.Itob(c.ID), encryptedData); err != nil {
		return err
	}

	// Commit transaction and exit.
	return tx.Commit()
}

// Get a contact by ID.
func (ab *AddressBook) Get(id uint64, password []byte) (*Contact, error) {
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
	encryptData := bkt.Get(dbutil.Itob(id))
	if encryptData == nil {
		return nil, errValEmpty
	}

	var c Contact
	if err := c.decryptContact(encryptData, password, ab.entropy); err != nil {
		return nil, err
	}
	return &c, nil
}

//
func (ab *AddressBook) List(password []byte) ([]*Contact, error) {
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
	var contacts []*Contact
	if err := bkt.ForEach(func(k, v []byte) error {
		var c Contact
		if err := c.decryptContact(v, password, ab.entropy); err != nil {
			return err
		}
		contacts = append(contacts, &c)
		return nil
	}); err != nil {
		return nil, err
	}
	return contacts, nil
}

//
func (ab *AddressBook) Delete(id uint64) error {
	return ab.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(dbAddrsBookBkt)
		if bkt == nil {
			return errBucketEmpty
		}
		return bkt.Delete(dbutil.Itob(id))
	})

}

//
func (ab *AddressBook) Update(id uint64, newcontact Contact) error {
	return nil
}

// Close shuts down the database.
func (ab *AddressBook) Close() error {
	return ab.db.Close()
}

//
func (ab *AddressBook) GetConfig() error {
	tx, err := ab.db.Begin(false)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logrus.Fatal(err)
		}
	}()

	bkt := tx.Bucket(dbConfigBkt)
	if bkt == nil {
		return errBucketEmpty
	}

	// Read encoded contact bytes.
	encryptData := bkt.Get(dbConfigDataKey)
	if encryptData == nil {
		return errValEmpty
	}
	return ab.UnmarshalBinary(encryptData)
}

//
func (ab *AddressBook) MarshalBinary() ([]byte, error) {

	return proto.Marshal(&internal.Config{
		Hash:          ab.hash,
		Entropy:       ab.entropy,
		HasPassphrase: ab.hasPassword,
	})
}

//
func (ab *AddressBook) UnmarshalBinary(data []byte) error {
	var pb internal.Config
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}
	ab.hash = pb.Hash
	ab.entropy = pb.Entropy
	ab.hasPassword = pb.HasPassphrase
	return nil
}

//
func (ab *AddressBook) InsertPass(password []byte) error {
	if password == nil {
		ab.hasPassword = true
	}
	return ab.generateHash(password)
}

//
func (ab *AddressBook) GenerateMnemonic(mnemonic string) error {
	if mnemonic != "" {
		if err := bip39.ValidateMnemonic(mnemonic); err != nil {
			return err
		}
		e, err := bip39.EntropyFromMnemonic(mnemonic)
		if err != nil {
			return err
		}
		ab.entropy = e
	} else {
		mn, err := bip39.NewDefaultMnemonic()
		if err != nil {
			return err
		}
		e, err := bip39.EntropyFromMnemonic(mn)
		if err != nil {
			return err
		}
		ab.entropy = e
	}
	return nil
}
