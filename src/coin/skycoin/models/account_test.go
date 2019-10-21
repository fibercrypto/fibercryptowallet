package skycoin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

func TestWalletListPendingTransactions(t *testing.T) {
	response := &api.UnconfirmedTxnsVerboseResponse{
		Transactions: []readable.UnconfirmedTransactionVerbose{
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
		},
	}
	global_mock.On("WalletUnconfirmedTransactionsVerbose", "42").Return(response, nil)
	global_mock.On("WalletUnconfirmedTransactionsVerbose", "G1").Return(response, nil)

	remote := &RemoteWallet{}
	remote.Id = "42"
	local := &LocalWallet{}
	local.Id = "G1"

	wallets := NewSkycoinWalletIterator([]core.Wallet{remote, local})

	for wallets.Next() {
		txns, err := wallets.Value().GetCryptoAccount().ListPendingTransactions()
		assert.Nil(t, err)
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
}

func TestSkycoinAddressGetBalance(t *testing.T) {
	response := new(api.BalanceResponse)
	response.Confirmed = readable.Balance{Coins: uint64(42000000), Hours: uint64(200)}
	global_mock.On("Balance", []string{"addr1"}).Return(response, nil)

	addr := &SkycoinAddress{address: "addr1"}
	val, err := addr.GetBalance(Sky)
	assert.Nil(t, err)
	assert.Equal(t, val, uint64(42000000))
	val, err = addr.GetBalance(CoinHour)
	assert.Nil(t, err)
	assert.Equal(t, val, uint64(200))
}

func TestSkycoinAddressScanUnspentOutputs(t *testing.T) {
	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         "addr1",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs:     readable.UnspentOutputs{usOut, usOut},
		OutgoingOutputs: readable.UnspentOutputs{usOut, usOut},
	}
	global_mock.On("OutputsForAddresses", []string{"addr1"}).Return(response, nil)

	addr := &SkycoinAddress{address: "addr1"}
	it := addr.ScanUnspentOutputs()

	for it.Next() {
		output := it.Value()
		assert.Equal(t, output.GetId(), "hash1")
		assert.Equal(t, output.GetAddress().String(), "addr1")
		val, err := output.GetCoins(Sky)
		assert.Nil(t, err)
		assert.Equal(t, val, uint64(42000000))
		val, err = output.GetCoins(CoinHour)
		assert.Nil(t, err)
		assert.Equal(t, val, uint64(42))
	}
}

func TestSkycoinAddressListTransactions(t *testing.T) {
	CleanGlobalMock()
	response := readable.TransactionWithStatusVerbose{
		Status: readable.TransactionStatus{
			Confirmed: false,
		},
	}
	response.Transaction.Hash = "hash1"
	global_mock.On("TransactionsVerbose", []string{"addr1"}).Return(
		[]readable.TransactionWithStatusVerbose{
			readable.TransactionWithStatusVerbose{
				Status: readable.TransactionStatus{
					Confirmed: true,
				},
			},
			response,
		},
		nil,
	)
	global_mock.On("Transaction", "hash1").Return(
		&readable.TransactionWithStatus{Status: response.Status},
		nil,
	)
	addr := &SkycoinAddress{address: "addr1"}
	it := addr.ListTransactions()
	it.Next()
	thx := it.Value()
	assert.Equal(t, thx.GetStatus(), core.TXN_STATUS_CONFIRMED)
	it.Next()
	thx = it.Value()
	assert.Equal(t, thx.GetStatus(), core.TXN_STATUS_PENDING)
}

func TestLocalWalletGetBalance(t *testing.T) {
	CleanGlobalMock()
	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         "2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs:     readable.UnspentOutputs{usOut, usOut},
		OutgoingOutputs: readable.UnspentOutputs{usOut, usOut},
		IncomingOutputs: readable.UnspentOutputs{usOut, usOut},
	}
	global_mock.On(
		"OutputsForAddresses",
		[]string{
			"2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
			"7wqRjpVwg5uSsz72oAZcDrBevHQRHQudyj",
			"2G9wDPX14WsbZuZU1f7MveYc9vpLxj2qNsz",
			"6gnBM5gMSSb7XRUEap7q3WxFnuvbN9usTq",
		},
	).Return(response, nil)
	addr := &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}
	val, err := addr.GetBalance(Sky)
	assert.Nil(t, err)
	assert.Equal(t, val, uint64(84000000))
	val, err = addr.GetBalance(CoinHour)
	assert.Nil(t, err)
	assert.Equal(t, val, uint64(84))
}
