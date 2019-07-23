package history

import (
	"github.com/therecipe/qt/core"
)

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
)

const (
	transactionStatusConfirmed = iota
	transactionStatusPending
	transactionStatusPreview
)

const (
	transactionTypeSend = iota
	transactionTypeReceive
)

type TransactionDetails struct {
	core.QObject

	_ *core.QDateTime `property:"date"`
	_ int             `property:"status"`
	_ int             `property:"type"`
	_ int             `property:"amount"`
	_ int             `property:"hoursReceived"`
	_ int             `property:"hoursBurned"`
	_ string          `property:"transactionID"`
	_ string          `property:"sentAddress"`
	_ string          `property:"receivedAddress"`
	_ *AddressList    `property:"inputs"`
	_ *AddressList    `property:"outputs"`
}
