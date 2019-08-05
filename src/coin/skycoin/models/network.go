package skycoin

import (
	"strconv"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type SkycoinPEX struct { //Implements PEX interface

} 

func (spex *SkycoinPEX) getConnections()  {
	//TODO
}

func (spex *SkycoinPEX) broadcastTxn(txn core.Transaction)  {
	//TODO
}

func (spex *SkycoinPEX) getTxnPool() (core.TransactionIterator, error) {
	c := util.NewClient()
	txns, err := c.PendingTransactionsVerbose()
	if err != nil {
		return nil, err
	}
	skycoinTxns := make([]core.Transaction, 0)
	for _, txn := range txns {
		skycoinTxn := UnconfirmedTransactionToTransaction(&txn)
		skycoinTxns = append(skycoinTxns, skycoinTxn)
	}
	return NewSkycoinTransactionIterator(skycoinTxns), nil
}

func UnconfirmedTransactionToTransaction(ut *readable.UnconfirmedTransactionVerbose) *SkycoinTransaction {
	txn := new(SkycoinTransaction)
	txn.Timestamp = core.TransactionTimestamp(ut.Received.Unix())
	txn.Id = ut.Transaction.Hash
	txn.Status = core.TXN_STATUS_PENDING
	
	/*TODO: set Intputs using array ut.In
		type TransactionInput struct {
			Hash            string `json:"uxid"`
			Address         string `json:"owner"`
			Coins           string `json:"coins"`
			Hours           uint64 `json:"hours"`
			CalculatedHours uint64 `json:"calculated_hours"`
		}
	*/

	txnOutputs := make([]core.TransactionOutput, 0)
	for _, output := range ut.Transaction.Out {
		txnOutput := new(SkycoinTransactionOutput)
		txnOutput.Id = output.Hash
		txnOutput.Address = SkycoinAddress{address: output.Address}
		coin, _ := strconv.ParseFloat(output.Coins, 64)
		//TODO: ask if is correct to use uint64 instead of float64
		txnOutput.Sky = uint64(coin)
		txnOutput.CoinHours = output.Hours
		txnOutputs = append(txnOutputs, txnOutput)
	}
	return txn
}
