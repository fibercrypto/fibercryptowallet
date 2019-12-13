package data

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/skycoin/skycoin/src/visor/dbutil"
	"time"
)

type boltStorage struct {
	*bolt.DB
}

const (
	// Db buckets.
	dbAddrsBookBkt = "AddressBook"
	dbConfigBkt    = "Config"
)

var (
	// Errors
	errDatabaseNotOpen = errors.New("database not open")
	errBucketEmpty     = errors.New("database: bucket are empty")
	errValEmpty        = errors.New(" database: result are empty")
)

// GetBoltStorage generate a new instance of boltStorage by path.
func GetBoltStorage(path string) (*boltStorage, error) {
	db, err := bolt.Open(path, 0600,
		&bolt.Options{
			Timeout: 1 * time.Second,
		})

	if err != nil {
		logDb.Error(err)
		return nil, err
	}

	return &boltStorage{db}, nil
}

// GetConfig Returns the config bucket content.
func (b *boltStorage) GetConfig() map[string]string {
	tx, err := b.Begin(false)
	if err != nil {
		logDb.Error(err)
		return nil
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logDb.Fatal(err)
		}
	}()

	if conf := tx.Bucket([]byte(dbConfigBkt)); conf != nil {
		confMap := make(map[string]string, 0)
		if err := conf.ForEach(func(k, v []byte) error {
			confMap[string(k)] = string(v)
			return nil
		}); err != nil {
			logDb.Error(err)
			return nil
		}
		return confMap
	}
	return nil
}

// InsertConfig set into config bucket the config parameters (securityType, hash, entropy)
func (b *boltStorage) InsertConfig(options map[string]string) error {
	// Start a writeable transaction.
	tx, err := b.Begin(true)
	if err != nil {
		logDb.Error(err)
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			// Start a writeable transaction.
			logDb.Fatal(err)
		}
	}()

	bkt, err := tx.CreateBucketIfNotExists([]byte(dbConfigBkt))
	if err != nil {
		logDb.Error(err)
		return err
	}
	for k, v := range options {
		if err := bkt.Put([]byte(k), []byte(v)); err != nil {
			logDb.Error(err)
			return err
		}
	}
	return tx.Commit()
}

// InsertValue insert a value in AddressBook bucket.
func (b *boltStorage) InsertValue(value interface{}) (uint64, error) {
	// Start a writeable transaction.
	tx, err := b.Begin(true)
	if err != nil {
		logDb.Error(err)
		return 0, err
	}
	element, ok := value.([]byte)
	if !ok {
		err := errValueNoMatch(element, []byte{})
		logDb.Error(err)
		return 0, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			logDb.Fatal(err)
		}
	}()

	bkt, err := tx.CreateBucketIfNotExists([]byte(dbAddrsBookBkt))
	if err != nil {
		logDb.Fatal(err)
		return 0, err
	}
	// The sequence is an auto-incrementing integer that is transactionally safe.
	id, err := bkt.NextSequence()

	if err != nil {
		logDb.Fatal(err)
		return 0, err
	}

	if err := bkt.Put(dbutil.Itob(id), element); err != nil {
		logDb.Fatal(err)
		return 0, err
	}
	return id, tx.Commit()
}

// GetValue get a value from the AddressBook bucket.
func (b *boltStorage) GetValue(key uint64) (interface{}, error) {
	tx, err := b.Begin(false)
	if err != nil {
		logDb.Error(err)
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logDb.Fatal(err)
		}
	}()
	if addrsBookBkt := tx.Bucket([]byte(dbAddrsBookBkt)); addrsBookBkt != nil {
		result := addrsBookBkt.Get(dbutil.Itob(key))
		if result == nil {
			return nil, errValEmpty
		}
		return result, nil
	}

	return nil, errBucketEmpty
}

// ListValues returns all values from AddressBook bucket.
func (b *boltStorage) ListValues() (map[uint64]interface{}, error) {
	tx, err := b.Begin(false)
	if err != nil {
		logDb.Error(err)
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			logDb.Fatal(err)
		}
	}()

	if addrsBookBkt := tx.Bucket([]byte(dbAddrsBookBkt)); addrsBookBkt != nil {
		resultsMap := make(map[uint64]interface{}, 0)
		if err := addrsBookBkt.ForEach(func(k, v []byte) error {
			resultsMap[dbutil.Btoi(k)] = v
			return nil
		}); err != nil {
			logDb.Error(err)
			return nil, err
		}
		return resultsMap, nil
	}
	return nil, errBucketEmpty
}

// DeleteValue remove a value from the AddressBook bucket by its id.
func (b *boltStorage) DeleteValue(key uint64) error {
	return b.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(dbAddrsBookBkt))
		if bkt == nil {
			return errBucketEmpty
		}

		if val := bkt.Get(dbutil.Itob(key)); val == nil {
			return errValEmpty
		}

		return bkt.Delete(dbutil.Itob(key))
	})

}

// UpdateValue update a element into the AddressBook bucket by its id.
func (b *boltStorage) UpdateValue(key uint64, newVal interface{}) error {
	return b.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(dbAddrsBookBkt))
		if bkt == nil {
			logDb.Error(errBucketEmpty)
			return errBucketEmpty
		}
		element, ok := newVal.([]byte)
		if !ok {
			err := errValueNoMatch(element, []byte{})
			logDb.Error(err)
			return err
		}
		return bkt.Put(dbutil.Itob(key), newVal.([]byte))
	})
}

func errValueNoMatch(value, valType interface{}) error {
	return fmt.Errorf("value %value type %T no match with type %T", value, value, valType)
}
