package skycoin

import (
	"errors"
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
	remote := &RemoteWallet{}
	remote.Id = "42"
	local := &LocalWallet{}
	local.Id = "G1"

	method := "WalletUnconfirmedTransactionsVerbose"
	wallets := NewSkycoinWalletIterator([]core.Wallet{remote, local})
	for wallets.Next() {
		wlt := wallets.Value()
		account := wlt.GetCryptoAccount()
		global_mock.On(method, wlt.GetId()).Return(response, errors.New("failure")).Once()
		txns, err := account.ListPendingTransactions()
		require.Error(t, err)
		global_mock.On(method, wlt.GetId()).Return(response, nil).Once()
		txns, err = account.ListPendingTransactions()
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
	CleanGlobalMock()
	response := new(api.BalanceResponse)
	addr, err := NewSkycoinAddress("2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt")
	require.NoError(t, err)
	response.Confirmed = readable.Balance{Coins: uint64(42000000), Hours: uint64(200)}
	skyAddrs := addr.GetCryptoAccount()
	global_mock.On("Balance", []string{addr.String()}).Return(response, errors.New("failure")).Once()
	val, err := skyAddrs.GetBalance(Sky)
	require.Error(t, err)
	global_mock.On("Balance", []string{addr.String()}).Return(response, nil)
	val, err = skyAddrs.GetBalance(Sky)
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
	CleanGlobalMock()
	addr := "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"
	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         addr,
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs: readable.UnspentOutputs{usOut, usOut},
	}
	addrs, err := NewSkycoinAddress(addr)
	assert.NoError(t, err)
	skyAddrs := addrs.GetCryptoAccount()

	global_mock.On("OutputsForAddresses", []string{addr}).Return(response, errors.New("failure")).Once()
	it := skyAddrs.ScanUnspentOutputs()
	require.Nil(t, it)

	global_mock.On("OutputsForAddresses", []string{addr}).Return(response, nil)
	it = skyAddrs.ScanUnspentOutputs()
	for it.Next() {
		output := it.Value()
		require.Equal(t, output.GetId(), "hash1")
		require.Equal(t, output.GetAddress().String(), addr)
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
	dir := "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"
	response := readable.TransactionWithStatusVerbose{
		Status: readable.TransactionStatus{
			Confirmed: false,
		},
	}
	response.Transaction.Hash = "hash1"
	global_mock.On("TransactionsVerbose", []string{dir}).Return(
		[]readable.TransactionWithStatusVerbose{},
		errors.New("failure"),
	).Once()

	addr, err := NewSkycoinAddress(dir)
	assert.NoError(t, err)
	skyAddr := addr.GetCryptoAccount()
	it := skyAddr.ListTransactions()
	require.Nil(t, it)
	global_mock.On("TransactionsVerbose", []string{dir}).Return(
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
	it = skyAddr.ListTransactions()
	it.Next()
	thx := it.Value()
	require.Equal(t, thx.GetStatus(), core.TXN_STATUS_CONFIRMED)
	it.Next()
	thx = it.Value()
	require.Equal(t, thx.GetStatus(), core.TXN_STATUS_PENDING)
}

func TestLocalWalletGetBalance(t *testing.T) {
	CleanGlobalMock()

	emptyOutput := readable.UnspentOutput{}
	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         "2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs:     readable.UnspentOutputs{usOut, usOut, emptyOutput},
		OutgoingOutputs: readable.UnspentOutputs{usOut, usOut},
		IncomingOutputs: readable.UnspentOutputs{usOut, usOut},
	}
	addrs := []string{
		"2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
		"7wqRjpVwg5uSsz72oAZcDrBevHQRHQudyj",
		"2G9wDPX14WsbZuZU1f7MveYc9vpLxj2qNsz",
		"6gnBM5gMSSb7XRUEap7q3WxFnuvbN9usTq",
	}
	method := "OutputsForAddresses"
	global_mock.On(method, addrs).Return(response, errors.New("failure")).Once()
	global_mock.On(method, addrs).Return(response, nil)

	// wallet not found
	wlt := &LocalWallet{WalletDir: "./testdata", Id: "no_wallet.wlt"}
	val, err := wlt.GetBalance(Sky)
	require.Error(t, err)

	// api interaction error
	wlt = &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}
	val, err = wlt.GetBalance(Sky)
	require.Error(t, err)

	// invalid HeadOutputs
	wlt = &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}
	val, err = wlt.GetBalance(Sky)
	require.Error(t, err)
	response.HeadOutputs = response.HeadOutputs[:len(response.HeadOutputs)-1]

	// all well
	val, err = wlt.GetBalance(Sky)
	require.NoError(t, err)
	require.Equal(t, uint64(84000000), val)
	val, err = wlt.GetBalance(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(84), val)

	//invalid ticker
	_, err = wlt.GetBalance("INVALID_TICKER") // nolint gosec
	require.Error(t, err)
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

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	global_mock.On("WalletBalance", "wallet").Return(response, errors.New("failure")).Once()
	val, err := wlt.GetBalance(Sky)
	require.Error(t, err)
	global_mock.On("WalletBalance", "wallet").Return(response, nil)
	val, err = wlt.GetBalance(Sky)
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

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	global_mock.On("Wallet", "wallet").Return(
		new(api.WalletResponse),
		errors.New("failure"),
	).Once()
	iter := wlt.ScanUnspentOutputs()
	require.Nil(t, iter)

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

	iter = wlt.ScanUnspentOutputs()
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
	dir := "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	global_mock.On("Wallet", "wallet").Return(new(api.WalletResponse), errors.New("failure")).Once()
	iter := wlt.ListTransactions()
	require.Nil(t, iter)

	global_mock.On("TransactionsVerbose", []string{dir}).Return(
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
				{Address: dir},
			},
		},
		nil)

	iter = wlt.ListTransactions()
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

	wlt := &LocalWallet{WalletDir: "./testdata", Id: "no_wallet.wlt"}
	iter := wlt.ScanUnspentOutputs()
	require.Nil(t, iter)

	wlt = &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}
	iter = wlt.ScanUnspentOutputs()
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

	wlt := &LocalWallet{WalletDir: "./testdata", Id: "no_wallet.wlt"}
	iter := wlt.ListTransactions()
	require.Nil(t, iter)

	wlt = &LocalWallet{WalletDir: "./testdata", Id: "test.wlt"}
	iter = wlt.ListTransactions()
	items := 0
	for iter.Next() {
		tx := iter.Value()
		items++
		require.Equal(t, "hash1", tx.GetId())
	}
	require.Equal(t, 4, items)
}

func Test_getBalanceOfAddresses(t *testing.T) {
	addrs := []string{"addr1"}
	emptyOutput := readable.UnspentOutput{}
	outBadCoin := readable.UnspentOutput{
		Coins:   "42a",
		Address: "addr1",
	}
	outValid := readable.UnspentOutput{Address: "addr1", Coins: "0"}
	tests := []struct {
		name      string
		summary   *readable.UnspentOutputsSummary
		wantError bool
	}{
		{
			name: "invalid_HeadOutput",
			summary: &readable.UnspentOutputsSummary{
				HeadOutputs: readable.UnspentOutputs{emptyOutput},
			},
			wantError: true,
		},
		{
			name: "invalid_coin_in_HeadOutput",
			summary: &readable.UnspentOutputsSummary{
				HeadOutputs: readable.UnspentOutputs{outBadCoin},
			},
			wantError: true,
		},
		{
			name: "invalid_IncomingOutput",
			summary: &readable.UnspentOutputsSummary{
				IncomingOutputs: readable.UnspentOutputs{emptyOutput},
			},
			wantError: true,
		},
		{
			name: "invalid_coin_in_IncomingOutput",
			summary: &readable.UnspentOutputsSummary{
				IncomingOutputs: readable.UnspentOutputs{outBadCoin},
			},
			wantError: true,
		},
		{
			name: "valid",
			summary: &readable.UnspentOutputsSummary{
				HeadOutputs:     readable.UnspentOutputs{outValid},
				IncomingOutputs: readable.UnspentOutputs{outValid},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bl, err := getBalanceOfAddresses(tt.summary, addrs)
			if tt.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, bl)
			}
		})
	}
}
