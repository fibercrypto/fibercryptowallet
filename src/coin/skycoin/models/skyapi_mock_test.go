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
