package skycoin

import (
	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/coin"
	"github.com/SkycoinProject/skycoin/src/readable"
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

	for _, addr := range addresses {
		mock.On(
			"OutputsForAddresses",
			[]string{
				addr,
			},
		).Return(response, nil)
	}
}

func mockSkyApiTransactionsVerbose(mock *SkycoinApiMock, addresses []string) {
	response := readable.TransactionWithStatusVerbose{
		Status: readable.TransactionStatus{
			Confirmed: false,
		},
	}
	response.Transaction.Hash = "hash1"

	for _, addr := range addresses {
		mock.On("TransactionsVerbose", []string{addr}).Return(
			[]readable.TransactionWithStatusVerbose{
				response,
			},
			nil,
		)
	}
}
