package data

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/sirupsen/logrus"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	"github.com/skycoin/skycoin/src/visor/dbutil"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --gogo_out=. internal/contact.proto
//go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --gogo_out=. internal/config.proto

var (
	addrsBook = addressBook{}
	// Errors
	errBucketEmpty = errors.New(" Error: bucket are empty.")
	errValEmpty    = errors.New(" Error: value are empty.")
	// Db buckets.
	dbAddrsBookBkt = []byte("addressBook")
	dbConfigBkt    = []byte("Config")
)

type addressBook struct {
	// db is a point of bolt.DB object. This handle all interaction with the db.
	db     *bolt.DB
	dbPath string // ".fiber/data.db"
	// hasPassword bool
	hash    []byte
	entropy []byte
}

//
func (ab *addressBook) verifyHash(password []byte) error {
	if ab.hash == nil {
		if err := ab.GetHashFromConfig(); err != nil {
			return err
		}
	}
	return bcrypt.CompareHashAndPassword(ab.hash, password)
}

//
func GetAddressBook() *addressBook {
	return &addrsBook
}

func (ab *addressBook) New(password []byte, path, mnemonic string) error {
	ab.dbPath = path
	if err := ab.Open(); err != nil {
		return err
	}
	if err := ab.GenerateEntropy(mnemonic); err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword(password, 14)

	if err != nil {
		return err
	}
	ab.hash = hash

	tx, err := ab.db.Begin(true)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			logrus.Fatal(err)
		}
	}()

	bConf := tx.Bucket(dbConfigBkt)
	if err := bConf.Put([]byte("hash"), ab.hash); err != nil {
		return err
	}
	if err := bConf.Put([]byte("entropy"), ab.entropy); err != nil {
		return err
	}

	return tx.Commit()
}

// startUp function open the database.
func (ab *addressBook) Open() error {
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
func (ab *addressBook) InsertContact(c core.Contact, password []byte) error {
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
	c.SetID(seq)

	if err := ab.verifyHash(password); err != nil {
		return fmt.Errorf(" Error verify password: %s", err)
	}

	encryptedData, err := c.EncryptContact(password, ab.entropy)
	if err != nil {
		return err
	}
	// Save contact to the bucket.
	if err := bkt.Put(dbutil.Itob(c.GetID()), encryptedData); err != nil {
		return err
	}
	// Commit transaction and exit.
	return tx.Commit()
}

// Get a contact by ID.
func (ab *addressBook) GetContact(id uint64, password []byte) (core.Contact, error) {
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
	if err := c.DecryptContact(encryptData, password, ab.entropy); err != nil {
		return nil, err
	}
	return &c, nil
}

//
func (ab *addressBook) ListContact(password []byte) ([]core.Contact, error) {
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
		var c Contact
		if err := c.DecryptContact(v, password, ab.entropy); err != nil {
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
func (ab *addressBook) DeleteContact(id uint64) error {
	return ab.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(dbAddrsBookBkt)
		if bkt == nil {
			return errBucketEmpty
		}
		return bkt.Delete(dbutil.Itob(id))
	})

}

//
func (ab *addressBook) Update(id uint64, newcontact Contact) error {
	return nil
}

// Close shuts down the database.
func (ab *addressBook) Close() error {
	return ab.db.Close()
}

//
// func (ab *addressBook) GetConfig() error {
// 	tx, err := ab.db.Begin(false)
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		if err := tx.Rollback(); err != nil {
// 			logrus.Fatal(err)
// 		}
// 	}()
//
// 	bkt := tx.Bucket(dbConfigBkt)
// 	if bkt == nil {
// 		return errBucketEmpty
// 	}
//
// 	// Read encoded contact bytes.
// 	encryptData := bkt.Get(dbConfigDataKey)
// 	if encryptData == nil {
// 		return errValEmpty
// 	}
// 	return ab.UnmarshalBinary(encryptData)
// }

//
func (ab *addressBook) GenerateEntropy(mnemonic string) error {
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

func (ab *addressBook) SetPath(path string) {
	ab.dbPath = path
}

func (ab *addressBook) GetPath() string {
	return ab.dbPath
}

func (ab *addressBook) GetHashFromConfig() error {
	tx, err := ab.db.Begin(false)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
	}()
	buck := tx.Bucket(dbConfigBkt)
	if buck == nil {
		return errBucketEmpty
	}

	hash := buck.Get([]byte("hash"))
	if hash == nil {
		return errValEmpty
	}
	ab.hash = hash
	return nil
}
