package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type SkycoinPEX struct { //Implements PEX interface

} 

func (spex *SkycoinPEX) getTxnPool() (core.TransactionIterator, error) {
	c := util.NewClient()
	transactions, err := c.PendingTransactionsVerbose()
	if err != nil {
		return nil, err
	}
	ptModels := make([]*core.Transaction, 0)
	for _, pt := range transactions {
		ptModel := UnconfirmedTransactionToTransaction(&pt)
		ptModels = append(ptModels, ptModel)
	}
	return nil, nil
}

func UnconfirmedTransactionToTransaction(ut *readable.UnconfirmedTransactionVerbose) *core.Transaction {
	//TODO: When the Transaction implementation is finished, create a new one and fill it out through ut
	return nil
}