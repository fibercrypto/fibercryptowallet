package core

import (
	"errors"
	"fmt"
	"sync"
)

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
type PooledObject interface {
}

// PooledObjectFactory instantiates pooled objects
type PooledObjectFactory interface {
	Create() (PooledObject, error)
}

// MultiPool implements a pool supporting multiple object factories
type MultiPool interface {
	// Get an object from the pool
	Get(poolSection string) (PooledObject, error)
	// Return put back an object into the pool
	Return(poolSection string, obj PooledObject) error
	// CreateSection allocates extra space for a new factory in the pool
	CreateSection(name string, factory PooledObjectFactory)
	// ListSections returns a list of section names previously added to the pool
	// Each name represents a different factory
	ListSections() []string
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
	capacity  int
	available map[string][]PooledObject
	inUse     map[string][]PooledObject
	mutexs    map[string]*sync.Mutex
	factories map[string]PooledObjectFactory
}

// Get an object from the pool
func (mp *MultiConnectionsPool) Get(poolSection string) (PooledObject, error) {
	mutex, ok := mp.mutexs[poolSection]

	if !ok {
		return nil, NotAvailableObjectsError{poolSection}
	}
	mutex.Lock()
	defer mutex.Unlock()

	if len(mp.available[poolSection]) == 0 {
		if len(mp.inUse[poolSection]) == mp.capacity {
			return nil, errors.New("There is not available objects")
		}
		obj, err := mp.factories[poolSection].Create()
		if err != nil {
			return nil, err
		}
		mp.inUse[poolSection] = append(mp.inUse[poolSection], obj)
		return obj, nil
	} else {
		var obj PooledObject
		obj, mp.available[poolSection] = mp.available[poolSection][0], mp.available[poolSection][1:]
		mp.inUse[poolSection] = append(mp.inUse[poolSection], obj)
		return obj, nil
	}
}

// Return put back an object into the pool
func (mp *MultiConnectionsPool) Return(poolSection string, obj PooledObject) error {
	mutex, ok := mp.mutexs[poolSection]
	if !ok {
		return errors.New(fmt.Sprintf("There is not exist %s poolSection", poolSection))
	}
	mutex.Lock()
	defer mutex.Unlock()
	index := findIndex(mp.inUse[poolSection], obj)
	if index == -1 {
		return errors.New(fmt.Sprintf("That object is no from this pool"))
	}
	mp.available[poolSection] = append(mp.available[poolSection], obj)
	mp.inUse[poolSection] = append(mp.inUse[poolSection][:index], mp.inUse[poolSection][index+1:]...)
	return nil
}

// CreateSection allocates extra space for a new factory in the pool
func (mp *MultiConnectionsPool) CreateSection(name string, factory PooledObjectFactory) {

	if _, ok := mp.factories[name]; ok {
		return
	}

	mp.factories[name] = factory
	mp.available[name] = make([]PooledObject, 0)
	mp.inUse[name] = make([]PooledObject, 0)
	mp.mutexs[name] = new(sync.Mutex)

}

// ListSections returns a list of section names previously registred with the pool
func (mp *MultiConnectionsPool) ListSections() []string {
	sections := make([]string, 0)
	for key, _ := range mp.factories {
		sections = append(sections, key)
	}
	return sections
}

func newMultiConnectionPool(capacity int) *MultiConnectionsPool {
	return &MultiConnectionsPool{
		capacity:  capacity,
		available: make(map[string][]PooledObject),
		inUse:     make(map[string][]PooledObject),
		factories: make(map[string]PooledObjectFactory),
		mutexs:    make(map[string]*sync.Mutex),
	}
}

// GetMultiPool instantiates singleton connection pool obj=ect
func GetMultiPool() MultiPool {

	once.Do(func() {
		multiConnectionsPool = newMultiConnectionPool(10)
	})

	return multiConnectionsPool
}

func findIndex(collection []PooledObject, obj PooledObject) int {
	for i := 0; i < len(collection); i++ {
		if collection[i] == obj {
			return i
		}
	}
	return -1
}
