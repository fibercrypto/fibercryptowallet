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
	GetObject() interface{}
	ReturnObject() error
}
type Object struct {
	value   interface{}
	section *PoolSection
}

func (obj *Object) GetObject() interface{} {
	return obj.value
}

func (obj *Object) ReturnObject() error {
	return obj.section.Put(obj)
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

func (mp *MultiConnectionsPool) GetSection(poolSection string) (*PoolSection, error) {
	section, ok := mp.sections[poolSection]
	if !ok {
		return nil, errors.New("Invalid Section")
	}
	return section, nil
}

func (mp *MultiConnectionsPool) CreateSection(name string, factory PooledObjectFactory) error {

	mp.sections[name] = &PoolSection{
		mutex:     new(sync.Mutex),
		capacity:  mp.capacity,
		factory:   factory,
		inUse:     make([]PooledObject, 0),
		available: make([]PooledObject, 0),
	}
	return nil
}

func (mp *MultiConnectionsPool) ListSections() []string {
	sections := make([]string, 0)
	for key, _ := range mp.factories {
		sections = append(sections, key)
	}
	return sections
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
		return &Object{
			section: ps,
			value:   obj,
		}, nil
	} else {
		var obj PooledObject
		obj, ps.available = ps.available[0], ps.available[1:]
		ps.inUse = append(ps.inUse, obj)
		return &Object{
			section: ps,
			value:   obj,
		}, nil
	}
}

func (ps *PoolSection) Put(obj PooledObject) error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	index := findIndex(ps.inUse, obj)
	if index == -1 {
		return errors.New(fmt.Sprintf("That object is no from this pool"))
	}
	ps.available = append(ps.available, obj)
	ps.inUse = append(ps.inUse[:index], ps.inUse[index+1:]...)
	return nil
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
