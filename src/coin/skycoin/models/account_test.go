package skycoin

import (
	"testing"

	"github.com/stretchr/testify/require"

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
		require.NoError(t, err)
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
}

func TestSkycoinAddressGetBalance(t *testing.T) {
	response := new(api.BalanceResponse)
	response.Confirmed = readable.Balance{Coins: uint64(42000000), Hours: uint64(200)}
	global_mock.On("Balance", []string{"addr1"}).Return(response, nil)

	addr := &SkycoinAddress{address: "addr1"}
	val, err := addr.GetBalance(Sky)
	require.NoError(t, err)
	require.Equal(t, val, uint64(42000000))
	val, err = addr.GetBalance(CoinHour)
	require.NoError(t, err)
	require.Equal(t, val, uint64(200))
	val, err = addr.GetBalance("INVALID_TICKER")
	require.Error(t, err)
	require.Equal(t, val, uint64(0))
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
		HeadOutputs: readable.UnspentOutputs{usOut, usOut},
	}

	global_mock.On("OutputsForAddresses", []string{"addr1"}).Return(response, nil)

	addr := &SkycoinAddress{address: "addr1"}
	it := addr.ScanUnspentOutputs()

	for it.Next() {
		output := it.Value()
		require.Equal(t, output.GetId(), "hash1")
		require.Equal(t, output.GetAddress().String(), "addr1")
		val, err := output.GetCoins(Sky)
		require.NoError(t, err)
		require.Equal(t, val, uint64(42000000))
		val, err = output.GetCoins(CoinHour)
		require.NoError(t, err)
		require.Equal(t, val, uint64(42))
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
	require.Equal(t, thx.GetStatus(), core.TXN_STATUS_CONFIRMED)
	it.Next()
	thx = it.Value()
	require.Equal(t, thx.GetStatus(), core.TXN_STATUS_PENDING)
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
	wlt := &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}
	val, err := wlt.GetBalance(Sky)
	require.NoError(t, err)
	require.Equal(t, uint64(84000000), val)
	val, err = wlt.GetBalance(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(84), val)
	val, err = wlt.GetBalance("INVALID_TICKER")
	require.Error(t, err)
	require.Equal(t, uint64(0), val)
}

func TestRemoteWalletGetBalance(t *testing.T) {
	CleanGlobalMock()
	response := new(api.BalanceResponse)
	response.Confirmed = readable.Balance{Coins: uint64(42000000), Hours: uint64(200)}
	global_mock.On("WalletBalance", "wallet").Return(response, nil)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	val, err := wlt.GetBalance(Sky)
	require.NoError(t, err)
	require.Equal(t, val, uint64(42000000))
	val, err = wlt.GetBalance(CoinHour)
	require.NoError(t, err)
	require.Equal(t, val, uint64(200))
	val, err = wlt.GetBalance("INVALID_TICKER")
	require.Error(t, err)
	require.Equal(t, val, uint64(0))
}

func TestRemoteWalletScanUnspentOutputs(t *testing.T) {
	CleanGlobalMock()

	global_mock.On("Wallet", "wallet").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "wallet",
				Encrypted: true,
			},
			Entries: []readable.WalletEntry{
				readable.WalletEntry{Address: "addr"},
			},
		},
		nil)

	global_mock.On("Wallet", "wallet_no_outputs").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "wallet_no_outputs",
				Encrypted: true,
			},
		},
		nil)

	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         "addr",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs: readable.UnspentOutputs{usOut},
	}

	global_mock.On("OutputsForAddresses", []string{"addr"}).Return(response, nil)

	global_mock.On("OutputsForAddresses", []string{"no_outputs"}).Return(&readable.UnspentOutputsSummary{}, nil)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	iter := wlt.ScanUnspentOutputs()
	items := 0
	for iter.Next() {
		to := iter.Value()
		items++
		require.Equal(t, "addr", to.GetAddress().String())
	}
	require.Equal(t, 1, items)

	//No outputs
	wlt = &RemoteWallet{
		Id:          "wallet_no_outputs",
		poolSection: PoolSection,
	}
	iter = wlt.ScanUnspentOutputs()
	items = 0
	for iter.Next() {
		to := iter.Value()
		items++
		require.Nil(t, to)
	}
	require.Equal(t, 0, items)
}

func TestRemoteWalletListTransactions(t *testing.T) {
	CleanGlobalMock()

	response := readable.TransactionWithStatusVerbose{
		Status: readable.TransactionStatus{
			Confirmed: false,
		},
	}
	response.Transaction.Hash = "hash1"

	global_mock.On("TransactionsVerbose", []string{"addr"}).Return(
		[]readable.TransactionWithStatusVerbose{
			response,
		},
		nil,
	)

	global_mock.On("Wallet", "wallet").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "wallet",
				Encrypted: true,
			},
			Entries: []readable.WalletEntry{
				readable.WalletEntry{Address: "addr"},
			},
		},
		nil)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	iter := wlt.ListTransactions()
	items := 0
	for iter.Next() {
		tx := iter.Value()
		items++
		require.Equal(t, "hash1", tx.GetId())
	}
	require.Equal(t, 1, items)
}
