package core

import (
	"errors"
	"fmt"
	"sync"
)

var once sync.Once
var multiConnectionsPool *MultiConnectionsPool

type PEX interface {
	GetTxnPool() (TransactionIterator, error)
	GetConnections() (PexNodeSet, error)
	BroadcastTxn(txn Transaction) error
}

type PexNodeIterator interface {
	Value() PexNode
	Next() bool
	HasNext() bool
}

type PexNodeSet interface {
	ListPeers() PexNodeIterator
}

type PexNode interface {
	GetIp() string
	GetPort() uint16
	GetBlockHeight() uint64
	IsTrusted() bool
	GetLastSeenIn() int64
	GetLastSeenOut() int64
}

type PooledObject interface {
}

type PooledObjectFactory interface {
	Create() (PooledObject, error)
}

type MultiPool interface {
	Get(poolSection string) (PooledObject, error)
	Return(poolSection string, obj PooledObject) error
	CreateSection(name string, factory PooledObjectFactory)
	ListSections() []string
}

type NotAvailableObjectsError struct {
	poolSection string
}

func (err NotAvailableObjectsError) Error() string {
	return fmt.Sprintf("There is not exist %s poolSection", err.poolSection)
}

type MultiConnectionsPool struct {
	capacity  int
	available map[string][]PooledObject
	inUse     map[string][]PooledObject
	mutexs    map[string]*sync.Mutex
	factories map[string]PooledObjectFactory
}

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

func (mp *MultiConnectionsPool) Return(poolSection string, obj PooledObject) error {
	mutex, ok := mp.mutexs[poolSection]
	if !ok {
		return fmt.Errorf("There is not exist %s poolSection", poolSection)
	}
	mutex.Lock()
	defer mutex.Unlock()
	index := findIndex(mp.inUse[poolSection], obj)
	if index == -1 {
		return fmt.Errorf("That object is no from this pool")
	}
	mp.available[poolSection] = append(mp.available[poolSection], obj)
	mp.inUse[poolSection] = append(mp.inUse[poolSection][:index], mp.inUse[poolSection][index+1:]...)
	return nil
}

func (mp *MultiConnectionsPool) CreateSection(name string, factory PooledObjectFactory) {

	if _, ok := mp.factories[name]; ok {
		return
	}

	mp.factories[name] = factory
	mp.available[name] = make([]PooledObject, 0)
	mp.inUse[name] = make([]PooledObject, 0)
	mp.mutexs[name] = new(sync.Mutex)

}

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
