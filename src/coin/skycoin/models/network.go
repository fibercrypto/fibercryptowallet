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
		//skycoinTxn := UnconfirmedTransactionToTransaction(&txn)
		skycoinTxns = append(skycoinTxns, &SkycoinPendingTransaction{Transaction: txn})
	}
	return NewSkycoinTransactionIterator(skycoinTxns), nil
}

// func UnconfirmedTransactionToTransaction(ut *readable.UnconfirmedTransactionVerbose) *SkycoinTransaction {
// 	txn := new(SkycoinTransaction)
// 	txn.Timestamp = core.TransactionTimestamp(ut.Received.Unix())
// 	txn.Id = ut.Transaction.Hash
// 	txn.Status = core.TXN_STATUS_PENDING
// 	txn.Fee = ut.Transaction.Fee
	
// 	/*TODO: set Intputs using array ut.In
// 		type TransactionInput struct {
// 			Hash            string `json:"uxid"`
// 			Address         string `json:"owner"`
// 			Coins           string `json:"coins"`
// 			Hours           uint64 `json:"hours"`
// 			CalculatedHours uint64 `json:"calculated_hours"`
// 		}
// 	*/

// 	txnOutputs := make([]core.TransactionOutput, 0)
// 	for _, output := range ut.Transaction.Out {
// 		txnOutput := new(SkycoinTransactionOutput)
// 		txnOutput.Id = output.Hash
// 		txnOutput.Address = SkycoinAddress{address: output.Address}
// 		coin, _ := strconv.ParseFloat(output.Coins, 64)
// 		txnOutput.Sky = uint64(coin)
// 		txnOutput.CoinHours = output.Hours
// 		txnOutputs = append(txnOutputs, txnOutput)
// 	}
// 	txn.Outputs = txnOutputs
// 	return txn
// }
