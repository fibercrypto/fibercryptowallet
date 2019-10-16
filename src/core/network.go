package core

import (
	"errors"
	"fmt"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"sync"
)

var logConnectionPool = logging.MustGetLogger("Connection Pool")

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

type PooledObjectFactory interface {
	Create() (interface{}, error)
}

type MultiPool interface {
	GetSection(string) (MultiPoolSection, error)
	ListSections() ([]string, error)
	CreateSection(string, PooledObjectFactory) error
}

type MultiPoolSection interface {
	Get() interface{}
	Put(interface{})
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

func (mp *MultiConnectionsPool) GetSection(poolSection string) (MultiPoolSection, error) {
	logConnectionPool.Info("Geeting " + poolSection + "pool section")
	section, ok := mp.sections[poolSection]
	if !ok {
		return nil, errors.New("Invalid Section")
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
			return errors.New("There is not available objects")
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
