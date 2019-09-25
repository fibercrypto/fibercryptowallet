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

func NewSkycoinPEX(poolSection string) *SkycoinPEX {
	return &SkycoinPEX{poolSection}
}

type SkycoinPEX struct { //Implements PEX interface
	poolSection string
}

func (spex *SkycoinPEX) GetConnections() (core.PexNodeSet, error) {
	//TODO
	return nil, nil
}

func (spex *SkycoinPEX) BroadcastTxn(txn core.Transaction) error {
	//TODO
	return nil
}

func (spex *SkycoinPEX) GetTxnPool() (core.TransactionIterator, error) {
	c, err := NewSkycoinApiClient(spex.poolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(spex.poolSection, c)
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
