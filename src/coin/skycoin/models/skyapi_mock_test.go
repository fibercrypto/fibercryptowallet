package skycoin

import (
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
)

func mockSkyApiUxOut(mock *SkycoinApiMock, ux coin.UxOut) {
	rUxOut := makeSpentOutput(ux, 0, cipher.SHA256{})
	mock.On("UxOut", ux.Hash().Hex()).Return(&rUxOut, nil)
}

func mockSkyApiTransactions(mock *SkycoinApiMock, addressesN []string) {
	mock.On("Transactions", []string{}).Return(nil, nil)
	mock.On("Transactions", []string{addressesN[0]}).Return(
		[]readable.TransactionWithStatus{},
		nil)
	mock.On("Transactions", []string{addressesN[1]}).Return(
		[]readable.TransactionWithStatus{
			readable.TransactionWithStatus{
				Status: readable.TransactionStatus{
					Confirmed: true,
				},
			},
		},
		nil)
	mock.On("Transactions", []string{addressesN[2]}).Return(
		[]readable.TransactionWithStatus{
			readable.TransactionWithStatus{
				Status: readable.TransactionStatus{
					Confirmed: true,
				},
			},
			readable.TransactionWithStatus{
				Status: readable.TransactionStatus{
					Confirmed: false,
				},
			},
		},
		nil)
}

func mockSkyApiCreateWallet(mock *SkycoinApiMock, wltOpt *api.CreateWalletOptions, label string, encrypted bool) {
	mock.On("CreateWallet", *wltOpt).Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     label,
				Encrypted: encrypted,
			},
		},
		nil)
}

func mockSkyApiWalletCreateTransaction(mock *SkycoinApiMock, wreq *api.WalletCreateTransactionRequest, crtTxn *api.CreateTransactionResponse) {
	mock.On("WalletCreateTransaction", *wreq).Return(
		crtTxn,
		nil)
}

func mockSkyApiCreateTransaction(mock *SkycoinApiMock, req *api.CreateTransactionRequest, crtTxn *api.CreateTransactionResponse) {
	mock.On("CreateTransaction", *req).Return(
		crtTxn,
		nil)
}

func mockSkyApiOutputsForAddresses(mock *SkycoinApiMock, addresses []string) {
	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         "2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs: readable.UnspentOutputs{usOut, usOut},
	}

	mock.On(
		"OutputsForAddresses",
		[]string{
			"6gnBM5gMSSb7XRUEap7q3WxFnuvbN9usTq",
		},
	).Return(response, nil)

	mock.On(
		"OutputsForAddresses",
		[]string{
			"2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
		},
	).Return(response, nil)

	mock.On(
		"OutputsForAddresses",
		[]string{
			"7wqRjpVwg5uSsz72oAZcDrBevHQRHQudyj",
		},
	).Return(response, nil)

	mock.On(
		"OutputsForAddresses",
		[]string{
			"2G9wDPX14WsbZuZU1f7MveYc9vpLxj2qNsz",
		},
	).Return(response, nil)
}

func mockSkyApiTransactionsVerbose(mock *SkycoinApiMock, addresses []string) {
	response := readable.TransactionWithStatusVerbose{
		Status: readable.TransactionStatus{
			Confirmed: false,
		},
	}
	response.Transaction.Hash = "hash1"

	mock.On("TransactionsVerbose", []string{"2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP"}).Return(
		[]readable.TransactionWithStatusVerbose{
			response,
		},
		nil,
	)

	mock.On("TransactionsVerbose", []string{"7wqRjpVwg5uSsz72oAZcDrBevHQRHQudyj"}).Return(
		[]readable.TransactionWithStatusVerbose{
			response,
		},
		nil,
	)

	mock.On("TransactionsVerbose", []string{"2G9wDPX14WsbZuZU1f7MveYc9vpLxj2qNsz"}).Return(
		[]readable.TransactionWithStatusVerbose{
			response,
		},
		nil,
	)

	mock.On("TransactionsVerbose", []string{"6gnBM5gMSSb7XRUEap7q3WxFnuvbN9usTq"}).Return(
		[]readable.TransactionWithStatusVerbose{
			response,
		},
		nil,
	)
}
