package core

import (
	"fmt"
	"sync"

	"github.com/fibercrypto/FiberCryptoWallet/src/errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
)

var logConnectionPool = logging.MustGetLogger("Connection Pool")

var once sync.Once
var multiConnectionsPool *MultiConnectionsPool

// PEX exposes cryptocurrency API for peer-to-peer communication
type PEX interface {
	// GetTxnPool return transactions pending for confirmation by network peers
	GetTxnPool() (TransactionIterator, error)
	// GetConnection enumerate connectionns to peer nodes
	GetConnections() (PexNodeSet, error)
	// BroadcastTxn injects a transaction for confirmation by network peers
	BroadcastTxn(txn Transaction) error
}

// PexNodeIterator scans nodes in a set
type PexNodeIterator interface {
	// Value of PEX node data instance at iterator pointer position
	Value() PexNode
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
}

// PexNodeSet represent a set of nodes
type PexNodeSet interface {
	// ListPeers offers an iterator over this set of nodes
	ListPeers() PexNodeIterator
}

// PexNode represents a peer in he cryptocurrency network
type PexNode interface {
	// GetIp returns node IP network address
	GetIp() string
	// GetPort retrieves IP port used to connect to peer node
	GetPort() uint16
	// GetBlockHeight provides sequence number of the block a the tip of peer's chain
	GetBlockHeight() uint64
	// IsTrusted determines if peer node is a network seed node
	IsTrusted() bool
	// GetLastSeenIn
	// TODO: Document method overview
	GetLastSeenIn() int64
	// GetLastSeenOut
	// TODO: Document method overview
	GetLastSeenOut() int64
}

// PooledObject represents any object that can be added to a connnection pool
// PooledObjectFactory instantiates pooled objects
type PooledObjectFactory interface {
	Create() (interface{}, error)
}

// MultiPool implements a pool supporting multiple object factories
type MultiPool interface {
	GetSection(string) (MultiPoolSection, error)
	ListSections() ([]string, error)
	CreateSection(string, PooledObjectFactory) error
}

type MultiPoolSection interface {
	Get() interface{}
	Put(interface{})
}

// NotAvailableObjectsError is returned when name is not bound to any pool factory
type NotAvailableObjectsError struct {
	poolSection string
}

// Error describes error condition
func (err NotAvailableObjectsError) Error() string {
	return fmt.Sprintf("There is not exist %s poolSection", err.poolSection)
}

// MultiConnectionsPool implements a generic pool supporting multiple object factories
type MultiConnectionsPool struct {
	capacity int
	sections map[string]*PoolSection
}

func (mp *MultiConnectionsPool) GetSection(poolSection string) (MultiPoolSection, error) {
	logConnectionPool.Info("Getting " + poolSection + "pool section")
	section, ok := mp.sections[poolSection]
	if !ok {
		return nil, errors.ErrInvalidPoolSection
	}
	return section, nil
}

func (mp *MultiConnectionsPool) CreateSection(name string, factory PooledObjectFactory) error {
	logConnectionPool.Info("Creating pool section")
	mp.sections[name] = &PoolSection{
		mutex:     new(sync.Mutex),
		capacity:  mp.capacity,
		factory:   factory,
		inUse:     make([]interface{}, 0),
		available: make([]interface{}, 0),
	}
	return nil
}

func (mp *MultiConnectionsPool) ListSections() ([]string, error) {
	logConnectionPool.Info("Listing pool sections")
	sections := make([]string, 0)
	for key, _ := range mp.sections {
		sections = append(sections, key)
	}
	return sections, nil
}

type PoolSection struct {
	capacity  int
	available []interface{}
	inUse     []interface{}
	mutex     *sync.Mutex
	factory   PooledObjectFactory
}

func (ps *PoolSection) Get() interface{} {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	if len(ps.available) == 0 {
		if len(ps.inUse) == ps.capacity {
			return errors.ErrObjectPoolUndeflow
		}
		obj, err := ps.factory.Create()
		if err != nil {
			return err
		}
		ps.inUse = append(ps.inUse, obj)
		return obj
	} else {
		var obj interface{}
		obj, ps.available = ps.available[0], ps.available[1:]
		ps.inUse = append(ps.inUse, obj)
		return obj
	}
}

func (ps *PoolSection) Put(obj interface{}) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	index := findIndex(ps.inUse, obj)
	if index == -1 {
		return
	}
	ps.available = append(ps.available, obj)
	ps.inUse = append(ps.inUse[:index], ps.inUse[index+1:]...)

}

func newMultiConnectionPool(capacity int) *MultiConnectionsPool {
	return &MultiConnectionsPool{
		capacity: capacity,
		sections: make(map[string]*PoolSection),
	}
}

// GetMultiPool instantiates singleton connection pool object
func GetMultiPool() MultiPool {

	once.Do(func() {
		multiConnectionsPool = newMultiConnectionPool(10)
	})

	return multiConnectionsPool
}

func findIndex(collection []interface{}, obj interface{}) int {
	for i := 0; i < len(collection); i++ {
		if collection[i] == obj {
			return i
		}
	}
	return -1
}
