package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
)

type SkycoinPEX struct { //Implements PEX interface
	NodeAddress string
} 

func (spex *SkycoinPEX) newClient() *api.Client {
	c := api.NewClient(spex.NodeAddress)
	return c
}

func (spex *SkycoinPEX) GetConnections()  {
	//TODO
}

func (spex *SkycoinPEX) BroadcastTxn(txn core.Transaction)  {
	//TODO
}

func (spex *SkycoinPEX) GetTxnPool() (core.TransactionIterator, error) {
	c := spex.newClient()
	txns, err := c.PendingTransactionsVerbose()
	if err != nil {
		return nil, err
	}
	skycoinTxns := make([]core.Transaction, 0)
	for _, txn := range txns {
		skycoinTxns = append(skycoinTxns, &SkycoinPendingTransaction{Transaction: txn})
	}
	return NewSkycoinTransactionIterator(skycoinTxns), nil
}
