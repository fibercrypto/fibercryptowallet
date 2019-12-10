package skycoin

import (
	"encoding/hex"
	"testing"

	"github.com/SkycoinProject/skycoin/src/visor"

	"github.com/stretchr/testify/require"

	"github.com/SkycoinProject/skycoin/src/readable"
)

func TestSkycoinPEXGetTxnPool(t *testing.T) {
	global_mock.On("PendingTransactionsVerbose").Return(
		[]readable.UnconfirmedTransactionVerbose{
			readable.UnconfirmedTransactionVerbose{
				Transaction: readable.BlockTransactionVerbose{
					Out: []readable.TransactionOutput{
						readable.TransactionOutput{
							Coins: "1",
							Hours: 2000,
						},
						readable.TransactionOutput{
							Coins: "1",
							Hours: 2000,
						},
					},
				},
			},
			readable.UnconfirmedTransactionVerbose{
				Transaction: readable.BlockTransactionVerbose{
					Out: []readable.TransactionOutput{
						readable.TransactionOutput{
							Coins: "1",
							Hours: 2000,
						},
						readable.TransactionOutput{
							Coins: "1",
							Hours: 2000,
						},
					},
				},
			},
		}, nil)

	pex := &SkycoinPEX{poolSection: PoolSection}
	txns, err2 := pex.GetTxnPool()
	require.NoError(t, err2)

	for txns.Next() {
		iter := NewSkycoinTransactionOutputIterator(txns.Value().GetOutputs())
		for iter.Next() {
			output := iter.Value()
			val, err3 := output.GetCoins(Sky)
			require.NoError(t, err3)
			require.Equal(t, val, uint64(1000000))
			val, err3 = output.GetCoins(CoinHour)
			require.NoError(t, err3)
			require.Equal(t, val, uint64(2000))
		}
	}
}

func TestSkycoinPEXBroadcastTxn(t *testing.T) {
	CleanGlobalMock()

	txn, err := makeTransaction(t)
	require.NoError(t, err)
	txn.In = nil
	txn.Out = nil

	txnV, err := readable.NewTransactionVerbose(
		visor.Transaction{
			Transaction: txn,
		},
		nil)

	require.NoError(t, err)

	skyTxn := &SkycoinTransaction{
		skyTxn: txnV,
	}

	txnBytes, err := serializeCreatedTransaction(skyTxn)
	require.NoError(t, err)

	global_mock.On("InjectEncodedTransaction", hex.EncodeToString(txnBytes)).Return("", nil)

	pex := &SkycoinPEX{poolSection: PoolSection}
	err = pex.BroadcastTxn(skyTxn)
	require.NoError(t, err)
}
