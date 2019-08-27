package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
)

type SkycoinConnectionFactory struct {
	url string
}

func (cf *SkycoinConnectionFactory) Create() core.PooledObject {
	return api.NewClient(cf.url)
}

func NewSkycoinConnectionFactory(url string) *SkycoinConnectionFactory {
	return &SkycoinConnectionFactory{
		url: url,
	}
}
