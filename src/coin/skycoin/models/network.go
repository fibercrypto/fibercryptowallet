package skycoin

import (
	"errors"
	"fmt"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
)

var logNetwork = logging.MustGetLogger("Skycoin network")

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
	logNetwork.Info("Creating Skycoin api client")
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
	logNetwork.Info("Creating new Skycoin PEX")
	return &SkycoinPEX{poolSection}
}

//Implements PEX interface
type SkycoinPEX struct {
	poolSection string
}

func (spex *SkycoinPEX) GetConnections() (core.PexNodeSet, error) {
	//TODO
	return nil, nil
}

func (spex *SkycoinPEX) BroadcastTxn(txn core.Transaction) error {
	logNetwork.Info("Broadcasting transaction")
	unTxn, ok := txn.(*SkycoinUninjectedTransaction)
	if !ok {
		return errors.New("Invalid Transaction")
	}
	c, err := NewSkycoinApiClient(spex.poolSection)
	if err != nil {
		return err
	}
	defer core.GetMultiPool().Return(spex.poolSection, c)
	_, err = c.InjectTransaction(unTxn.txn)
	if err != nil {
		return err
	}

	return nil
}

func (spex *SkycoinPEX) GetTxnPool() (core.TransactionIterator, error) {
	logNetwork.Info("Getting transaction pool")
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
