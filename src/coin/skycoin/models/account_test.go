package skycoin

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

func TestWalletListPendingTransactions(t *testing.T) {
	response := &api.UnconfirmedTxnsVerboseResponse{
		Transactions: []readable.UnconfirmedTransactionVerbose{
			{
				Transaction: readable.BlockTransactionVerbose{
					Out: []readable.TransactionOutput{
						{
							Coins: "1",
							Hours: 2000,
						},
						{
							Coins: "1",
							Hours: 2000,
						},
					},
				},
			},
			{
				Transaction: readable.BlockTransactionVerbose{
					Out: []readable.TransactionOutput{
						{
							Coins: "1",
							Hours: 2000,
						},
						{
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

func TestSkycoinAddressListAssets(t *testing.T) {
	addr, err := NewSkycoinAddress("2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt")
	assert.NoError(t, err)
	assets := addr.ListAssets()
	require.Equal(t, 2, len(assets))
	require.True(t, assets[0] == Sky || assets[1] == Sky)
	require.True(t, assets[0] == CoinHour || assets[1] == CoinHour)
}

func TestSkycoinAddressGetBalance(t *testing.T) {
	response := new(api.BalanceResponse)
	addr, err := NewSkycoinAddress("2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt")
	assert.NoError(t, err)
	response.Confirmed = readable.Balance{Coins: uint64(42000000), Hours: uint64(200)}
	global_mock.On("Balance", []string{addr.String()}).Return(response, nil)
	skyAddrs := addr.GetCryptoAccount()
	val, err := skyAddrs.GetBalance(Sky)
	require.NoError(t, err)
	require.Equal(t, val, uint64(42000000))
	val, err = skyAddrs.GetBalance(CoinHour)
	require.NoError(t, err)
	require.Equal(t, val, uint64(200))
	val, err = skyAddrs.GetBalance("INVALID_TICKER")
	require.Error(t, err)
	require.Equal(t, val, uint64(0))
}

func TestSkycoinAddressScanUnspentOutputs(t *testing.T) {
	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs: readable.UnspentOutputs{usOut, usOut},
	}

	global_mock.On("OutputsForAddresses", []string{"2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"}).Return(response, nil)

	addrs, err := NewSkycoinAddress("2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8")
	assert.NoError(t, err)
	skyAddrs := addrs.GetCryptoAccount()
	it := skyAddrs.ScanUnspentOutputs()

	for it.Next() {
		output := it.Value()
		require.Equal(t, output.GetId(), "hash1")
		require.Equal(t, output.GetAddress().String(), "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8")
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
	global_mock.On("TransactionsVerbose", []string{"2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"}).Return(
		[]readable.TransactionWithStatusVerbose{
			{
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
	addr, err := NewSkycoinAddress("2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8")
	assert.NoError(t, err)
	skyAddr := addr.GetCryptoAccount()
	it := skyAddr.ListTransactions()
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

func TestWalletsListAssets(t *testing.T) {
	wlts := []core.CryptoAccount{&RemoteWallet{}, &LocalWallet{}}
	assets := []string{Sky, CoinHour}
	sort.Strings(assets)

	for _, wlt := range wlts {
		wllt_assets := wlt.ListAssets()
		sort.Strings(wllt_assets)
		require.Equal(t, assets, wllt_assets)
	}
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
				{Address: "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt"},
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
		Address:         "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs: readable.UnspentOutputs{usOut},
	}

	// addrs
	global_mock.On("OutputsForAddresses", []string{"2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt"}).Return(response, nil)
	// no_output
	global_mock.On("OutputsForAddresses", []string{"2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"}).Return(&readable.UnspentOutputsSummary{}, nil)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	iter := wlt.ScanUnspentOutputs()
	items := 0
	for iter.Next() {
		to := iter.Value()
		items++
		require.Equal(t, "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt", to.GetAddress().String())
	}
	require.Equal(t, 1, items)

	// No outputs
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

	global_mock.On("TransactionsVerbose", []string{"2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"}).Return(
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
				{Address: "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"},
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

func TestLocalWalletScanUnspentOutputs(t *testing.T) {
	CleanGlobalMock()

	global_mock.On("Wallet", "test.wlt").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "test.wlt",
				Encrypted: true,
			},
			Entries: []readable.WalletEntry{
				{Address: "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"},
			},
		},
		nil)

	addresses := []string{
		"2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
		"7wqRjpVwg5uSsz72oAZcDrBevHQRHQudyj",
		"2G9wDPX14WsbZuZU1f7MveYc9vpLxj2qNsz",
		"6gnBM5gMSSb7XRUEap7q3WxFnuvbN9usTq",
	}

	mockSkyApiOutputsForAddresses(global_mock, addresses)

	wlt := &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}

	iter := wlt.ScanUnspentOutputs()
	items := 0
	for iter.Next() {
		to := iter.Value()
		items++
		require.Equal(t, "2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP", to.GetAddress().String())
	}
	require.Equal(t, 8, items)
}

func TestLocalWalletListTransactions(t *testing.T) {
	CleanGlobalMock()

	global_mock.On("Wallet", "test.wlt").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "test.wlt",
				Encrypted: true,
			},
			Entries: []readable.WalletEntry{
				{Address: "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"},
			},
		},
		nil)

	addresses := []string{
		"2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
		"7wqRjpVwg5uSsz72oAZcDrBevHQRHQudyj",
		"2G9wDPX14WsbZuZU1f7MveYc9vpLxj2qNsz",
		"6gnBM5gMSSb7XRUEap7q3WxFnuvbN9usTq",
	}

	mockSkyApiTransactionsVerbose(global_mock, addresses)

	wlt := &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}

	iter := wlt.ListTransactions()
	items := 0
	for iter.Next() {
		tx := iter.Value()
		items++
		require.Equal(t, "hash1", tx.GetId())
	}
	require.Equal(t, 4, items)
}
