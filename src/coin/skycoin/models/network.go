package skycoin

import (
	"errors"
	"fmt"

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

func NewSkycoinApiClient(section string) (*api.Client, error) {
	pool := core.GetMultiPool()
	obj, err := pool.Get(section)

	if err != nil {
		for _, ok := err.(core.NotAvailableObjectsError); ok; _, ok = err.(core.NotAvailableObjectsError) {
			obj, err = pool.Get(section)
			if err == nil {
				break
			}
		}
		return nil, err
	}

	skyApi, ok := obj.(*api.Client)
	if !ok {
		return nil, errors.New(fmt.Sprintf("There is not propers client in %s pool", section))
	}
	return skyApi, nil
}

type SkycoinPEX struct { //Implements PEX interface
	NodeAddress string
} 

func (spex *SkycoinPEX) GetConnections()  {
	//TODO
}

func (spex *SkycoinPEX) BroadcastTxn(txn core.Transaction)  {
	//TODO
}

func (spex *SkycoinPEX) GetTxnPool() (core.TransactionIterator, error) {
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	txns, err2 := c.PendingTransactionsVerbose()
	if err2 != nil {
		return nil, err2
	}
	skycoinTxns := make([]core.Transaction, 0)
	for _, txn := range txns {
		skycoinTxns = append(skycoinTxns, &SkycoinPendingTransaction{Transaction: txn})
	}
	return NewSkycoinTransactionIterator(skycoinTxns), nil
}
