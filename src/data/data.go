package data

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

//db is a point of bolt.DB object. This handle all interaction with the db.
var db *bolt.DB

const dbPath = ".fiber/data.db"

//Contact is a contact of the AddressBook
type Contact struct {
	Address    []Address
	WalletName []byte
}

//Address is the relation of an address and his coin type.
type Address struct {
	Value []byte
	Coin  []byte
}

//startUp function open the db for once operation.
func startUp(readOnly bool) error {
	var err error
	db, err = bolt.Open(getDbFileDir(),
		0666,
		&bolt.Options{
			Timeout:  1 * time.Second,
			ReadOnly: readOnly,
		})
	if err != nil {
		return fmt.Errorf("Error opening the DDBB: %s", err)
	}
	return nil
}

//Write once element in a bucket of the db.
func dbWrite(key, val []byte) error {
	if err := startUp(false); err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			logrus.Fatal("Error closing the DDBB.", err)
		}
	}()
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("AddressBook"))
		if err != nil {
			return err
		}

		if err = bucket.Put(key, val); err != nil {
			return err
		}

		return nil
	})
}

//Read once element from the db.
func dbRead(key string) ([]byte, error) {
	var err error
	var val []byte
	if err = startUp(true); err != nil {
		return nil, err
	}

	if err = db.View(func(tx *bolt.Tx) error {
		AddressBook := tx.Bucket([]byte("AddressBook"))
		if AddressBook == nil {
			return fmt.Errorf("AddressBook bucket not exist.")
		}
		if AddressBook.Get([]byte(key)) == nil {
			return fmt.Errorf("%s bucket not exist.", key)
		}
		val = AddressBook.Get([]byte(key))
		return nil
	}); err != nil {
		return nil, err
	}
	return val, nil
}

//getDbFileDir catch the path of the db.
func getDbFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, dbPath)
	return fileDir
}
