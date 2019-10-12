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
	GetSection(string) (MultiPoolSection, error)
	ListSections() ([]string, error)
	CreateSection(string, PooledObjectFactory) error
}

type MultiPoolSection interface {
	Get() (PooledObject, error)
	Put(interface{}) error
}

type NotAvailableObjectsError struct {
	poolSection string
}

func (err NotAvailableObjectsError) Error() string {
	return fmt.Sprintf("There is not exist %s poolSection", err.poolSection)
}

type MultiConnectionsPool struct {
	capacity int
	sections map[string]*PoolSection
}

type PoolSection struct {
	capacity  int
	available []PooledObject
	inUse     []PooledObject
	mutex     *sync.Mutex
	factory   PooledObjectFactory
}

func (ps *PoolSection) Get() (PooledObject, error) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	if len(ps.available) == 0 {
		if len(ps.inUse) == ps.capacity {
			return nil, errors.New("There is not available objects")
		}
		obj, err := ps.factory.Create()
		if err != nil {
			return nil, err
		}
		ps.inUse = append(ps.inUse, obj)
		return obj, nil
	} else {
		var obj PooledObject
		obj, ps.available = ps.available[0], ps.available[1:]
		ps.inUse = append(ps.inUse, obj)
		return obj, nil
	}
}

func (mp *MultiConnectionsPool) GetSection(poolSection string) (*PoolSection, error) {
	section, ok := mp.sections[poolSection]
	if !ok {
		return nil, errors.New("Invalid Section")
	}
	return section, nil
}

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

func (mp *MultiConnectionsPool) CreateSection(name string, factory PooledObjectFactory) error {

	if _, ok := mp.factories[name]; ok {
		return errors.New("Invalid section")
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
