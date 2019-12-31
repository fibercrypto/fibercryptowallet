package skycoin

import (
	"encoding/hex"
	"testing"

	"github.com/SkycoinProject/skycoin/src/visor"

	"github.com/stretchr/testify/require"

	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
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

func TestSkycoinPexNode(t *testing.T) {
	addr := "addr"
	port := uint16(8000)
	sent := int64(20)
	recived := int64(42)
	height := uint64(200)
	trusted := true

	conn := readable.Connection{
		Addr:          addr,
		ListenPort:    port,
		LastSent:      sent,
		LastReceived:  recived,
		Height:        height,
		IsTrustedPeer: trusted,
	}
	pex := connectionsToNetwork(conn)

	require.Equal(t, addr, pex.GetIp())
	require.Equal(t, port, pex.GetPort())
	require.Equal(t, sent, pex.GetLastSeenIn())
	require.Equal(t, recived, pex.GetLastSeenOut())
	require.Equal(t, height, pex.GetBlockHeight())
	require.Equal(t, trusted, pex.IsTrusted())
}

func TestSkycoinPexNodeIterator(t *testing.T) {
	pexs := make([]core.PexNode, 0)
	addrs := []string{"addr1", "addr2", "addr3"}
	for _, addr := range addrs {
		pex := &mocks.PexNode{}
		pex.On("GetIp").Return(addr)
		pexs = append(pexs, pex)
	}

	it := NewSkycoinPexNodeIterator(pexs)
	for _, addr := range addrs {
		require.Equal(t, true, it.Next())
		require.Equal(t, addr, it.Value().GetIp())
	}
	require.Equal(t, false, it.Next())
}
