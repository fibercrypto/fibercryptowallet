package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
)

const (
	PoolSection = "skycoin"
)

type SkycoinConnectionFactory struct {
	url string
}

func (cf *SkycoinConnectionFactory) Create() (core.PooledObject, error) {

	return api.NewClient(cf.url), nil
}

func NewSkycoinConnectionFactory(url string) *SkycoinConnectionFactory {

	return &SkycoinConnectionFactory{
		url: url,
	}
}

func WaitForPooledObject(pool core.MultiPool, section string) (core.PooledObject, error) {
	obj, err := pool.Get(section)

	if err != nil {
		for _, ok := err.(core.NotAvailableObjectsError); ok; _, ok = err.(core.NotAvailableObjectsError) {
			obj, err = pool.Get(section)
			if err == nil {
				return obj, nil
			}
		}
		return nil, err
	}

	return obj, nil
}
