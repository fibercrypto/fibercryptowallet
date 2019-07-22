package history

import (
	"fmt"

	"github.com/therecipe/qt/core"
)

func init() {
	HistoryModel_QmlRegisterType2("HistoryModels", 1, 0, "QHistory")
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

type Address struct {
	address          string
	addressSky       int64
	addressCoinHours int64
}

type TransactionDetails struct {
	core.QObject

	// _ func() `constructor:"init"`

	_ string `property:"date"`
	_ int    `property:"status"`
	_ int    `property:"type"`
	_ int    `property:"amount"`
	_ int    `property:"hoursReceived"`
	_ int    `property:"hoursBurned"`
	_ string `property:"transactionID"`
	_ string `property:"sentAddress"`
	_ string `property:"receivedAddress"`
}

// func (td *TransactionDetails) init() {
// 	td.SetProperty("date", core.NewQVariant12("2000-01-01"))
// 	td.SetProperty("status", core.NewQVariant5(transactionStatusConfirmed))
// 	td.SetProperty("type", core.NewQVariant5(transactionTypeReceive))
// 	td.SetProperty("amount", core.NewQVariant5(100))
// 	td.SetProperty("hoursReceived", core.NewQVariant5(100))
// 	td.SetProperty("hoursBurned", core.NewQVariant5(100))
// 	td.SetProperty("transactionID", core.NewQVariant12("transactionID"))
// 	td.SetProperty("sentAddress", core.NewQVariant12("addrexample1"))
// 	td.SetProperty("receivedAddress", core.NewQVariant12("addrexample2"))
// }

type HistoryModel struct {
	core.QAbstractListModel

	_ map[int]*core.QByteArray `property:"roles"`

	_ func() `constructor:"init"`

	_ func(transaction *TransactionDetails) `signal:"addTransaction"`

	_ []*TransactionDetails `property:"transactions"`
}

func (hm *HistoryModel) init() {
	hm.SetRoles(map[int]*core.QByteArray{
		Date:            core.NewQByteArray2("date", -1),
		Status:          core.NewQByteArray2("status", -1),
		Type:            core.NewQByteArray2("type", -1),
		Amount:          core.NewQByteArray2("amount", -1),
		HoursReceived:   core.NewQByteArray2("hoursReceived", -1),
		HoursBurned:     core.NewQByteArray2("hoursBurned", -1),
		TransactionID:   core.NewQByteArray2("transactionID", -1),
		SentAddress:     core.NewQByteArray2("sentAddress", -1),
		ReceivedAddress: core.NewQByteArray2("receivedAddress", -1),
	})

	hm.ConnectRowCount(hm.rowCount)
	hm.ConnectData(hm.data)
	hm.ConnectRoleNames(hm.roleNames)

	// hm.ConnectAddTransaction(hm.addTransaction)

	fmt.Println("Initialicing History Model")

	td := new(TransactionDetails)
	td.SetDate("2000-01-01")
	td.SetStatus(transactionStatusConfirmed)
	td.SetType(transactionTypeReceive)
	td.SetAmount(100)
	td.SetHoursReceived(100)
	td.SetHoursBurned(100)
	td.SetTransactionID("transactionID")
	td.SetSentAddress("addrexample1")
	td.SetReceivedAddress("addrexample2")
	hm.addTransaction(td)
	println("transaction->", td.Amount(), "<-")
	// fmt.Println(hm.Transactions())
	// fmt.Println(td)
	fmt.Println("Done")
}

func (hm *HistoryModel) rowCount(*core.QModelIndex) int {
	return len(hm.Transactions())
}

func (hm *HistoryModel) roleNames() map[int]*core.QByteArray {
	return hm.Roles()
}

func (hm *HistoryModel) addTransaction(transaction *TransactionDetails) {
	hm.BeginInsertRows(core.NewQModelIndex(), len(hm.Transactions()), len(hm.Transactions()))
	hm.SetTransactions(append(hm.Transactions(), transaction))
	hm.EndInsertRows()
}

func (hm *HistoryModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() >= len(hm.Transactions()) {
		return core.NewQVariant()
	}

	transaction := hm.Transactions()[index.Row()]

	switch role {
	case Date:
		{
			return core.NewQVariant1(transaction.Date())
		}
	case Status:
		{
			return core.NewQVariant1(transaction.Status())
		}
	case Type:
		{
			return core.NewQVariant1(transaction.Type())
		}
	case Amount:
		{
			return core.NewQVariant1(transaction.Amount())
		}
	case HoursReceived:
		{
			return core.NewQVariant1(transaction.HoursReceived())
		}
	case HoursBurned:
		{
			return core.NewQVariant1(transaction.HoursBurned())
		}
	case TransactionID:
		{
			return core.NewQVariant1(transaction.TransactionID())
		}
	case SentAddress:
		{
			return core.NewQVariant1(transaction.SentAddress())
		}
	case ReceivedAddress:
		{
			return core.NewQVariant1(transaction.ReceivedAddress())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}
