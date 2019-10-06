package skycoin

import(
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

func TestWalletListPendingTransactions(t *testing.T){
	global_mock.On("WalletUnconfirmedTransactionsVerbose", "42").Return(
		&api.UnconfirmedTxnsVerboseResponse {
			Transactions: []readable.UnconfirmedTransactionVerbose {
				readable.UnconfirmedTransactionVerbose{
					Transaction: readable.BlockTransactionVerbose{
						Out:    		[]readable.TransactionOutput{
							readable.TransactionOutput{
								Coins:    		"1",
								Hours:    		2000,
							},
							readable.TransactionOutput{
								Coins:    		"1",
								Hours:    		2000,
							},
						},
					},
				},
				readable.UnconfirmedTransactionVerbose{
					Transaction: readable.BlockTransactionVerbose{
						Out:    		[]readable.TransactionOutput{
							readable.TransactionOutput{
								Coins:    		"1",
								Hours:    		2000,
							},
							readable.TransactionOutput{
								Coins:    		"1",
								Hours:    		2000,
							},
						},
					},
				},
			},
		},
	nil)
	
	wlt := &RemoteWallet{}
	wlt.Id = "42"
	txns, err2 := wlt.ListPendingTransactions()
	assert.Nil(t, err2)

	for txns.Next() {
		iter := NewSkycoinTransactionOutputIterator(txns.Value().GetOutputs())
		for iter.Next() {
			output := iter.Value()
			val, err3 := output.GetCoins(Sky)
			assert.Nil(t, err3)
			assert.Equal(t, val, uint64(1000000))
			val, err3 = output.GetCoins(CoinHour)
			assert.Nil(t, err3)
			assert.Equal(t, val, uint64(2000))
		}
	}
}
