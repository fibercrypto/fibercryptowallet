package skycoin

import(
	"testing"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/stretchr/testify/assert"

)

func TestPending(t *testing.T){
	cf := local.GetConfigManager()
	client, err1 := NewSkycoinConnectionFactory(cf.GetNode()).Create()
	assert.Nil(t, err1)

	mock := &SkycoinApiMock{Node: client.(*api.Client)}
	mock.On("PendingTransactionsVerbose").Return(
		[]readable.UnconfirmedTransactionVerbose {
			readable.UnconfirmedTransactionVerbose{
				Transaction: readable.BlockTransactionVerbose{
					Length:        	15,
					Type:          	1,
					Hash:        	"hash",
					InnerHash:    	"inner_hash",
					Fee:        	30,
					Sigs:    		[]string{"s1", "s2"},
					In:    			[]readable.TransactionInput{
						readable.TransactionInput{
							Hash:        		"uxid",
							Address:        	"owner",
							Coins:        		"1000000",
							Hours:        		2000,
							CalculatedHours:    15,
						},
						readable.TransactionInput{
							Hash:        		"uxid",
							Address:        	"owner",
							Coins:        		"1000000",
							Hours:        		2000,
							CalculatedHours:    15,
						},
					},
					Out:    		[]readable.TransactionOutput{
						readable.TransactionOutput{
							Hash:    		"uxid",
							Address:    	"dst",
							Coins:    		"1",
							Hours:    		2000,
						},
						readable.TransactionOutput{
							Hash:    		"uxid",
							Address:    	"dst",
							Coins:    		"1",
							Hours:    		2000,
						},
					},
				},
			},
			readable.UnconfirmedTransactionVerbose{
				Transaction: readable.BlockTransactionVerbose{
					Length:        	15,
					Type:          	1,
					Hash:        	"hash",
					InnerHash:    	"inner_hash",
					Fee:        	30,
					Sigs:    		[]string{"s1", "s2"},
					In:    			[]readable.TransactionInput{
						readable.TransactionInput{
							Hash:        		"uxid",
							Address:        	"owner",
							Coins:        		"1000000",
							Hours:        		2000,
							CalculatedHours:    15,
						},
						readable.TransactionInput{
							Hash:        		"uxid",
							Address:        	"owner",
							Coins:        		"1000000",
							Hours:        		2000,
							CalculatedHours:    15,
						},
					},
					Out:    		[]readable.TransactionOutput{
						readable.TransactionOutput{
							Hash:    		"uxid",
							Address:    	"dst",
							Coins:    		"1",
							Hours:    		2000,
						},
						readable.TransactionOutput{
							Hash:    		"uxid",
							Address:    	"dst",
							Coins:    		"1",
							Hours:    		2000,
						},
					},
				},
			},
		}, nil)
	core.GetMultiPool().CreateSection(PoolSection, mock)
	util.RegisterAltcoin(NewSkyFiberPlugin(SkycoinMainNetParams))

	pex := &SkycoinPEX{}
	txns, err2 := pex.GetTxnPool()
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
