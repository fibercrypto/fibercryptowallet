package transactions

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/therecipe/qt/core"
)

func init() {
	TransactionDetails_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionDetail")
}

const (
	Date = int(core.Qt__UserRole) + 1<<iota
	Status
	Type
	Amount
	HoursReceived
	HoursBurned
	TransactionID
	SentAddress
	ReceivedAddress
	Inputs
	Outputs
)

const (
	TransactionStatusConfirmed = iota
	TransactionStatusPending
	TransactionStatusPreview
)

const (
	TransactionTypeSend = iota
	TransactionTypeReceive
)

type TransactionDetails struct {
	core.QObject

	_ *core.QDateTime      `property:"date"`
	_ int                  `property:"status"`
	_ int                  `property:"type"`
	_ int                  `property:"amount"`
	_ int                  `property:"hoursReceived"`
	_ int                  `property:"hoursBurned"`
	_ string               `property:"transactionID"`
	_ string               `property:"sentAddress"`
	_ string               `property:"receivedAddress"`
	_ *address.AddressList `property:"inputs"`
	_ *address.AddressList `property:"outputs"`
}
