package history

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/transactions"
	"github.com/therecipe/qt/core"
)

func init() {
	TransactionList_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionList")
}

type TransactionList struct {
	core.QAbstractListModel

	_ map[int]*core.QByteArray `property:"roles"`

	_ func() `constructor:"init"`

	_ func(transaction *transactions.TransactionDetails) `signal:"addTransaction,auto"`
	_ func(index int)                                    `signal:"removeTransaction,auto"`
	_ func(wallets []string)                             `signal:"loadHistory,auto"`

	_ []*transactions.TransactionDetails `property:"transactions"`
}

func (hm *TransactionList) init() {
	hm.SetRoles(map[int]*core.QByteArray{
		transactions.Date:            core.NewQByteArray2("date", -1),
		transactions.Status:          core.NewQByteArray2("status", -1),
		transactions.Type:            core.NewQByteArray2("type", -1),
		transactions.Amount:          core.NewQByteArray2("amount", -1),
		transactions.HoursReceived:   core.NewQByteArray2("hoursReceived", -1),
		transactions.HoursBurned:     core.NewQByteArray2("hoursBurned", -1),
		transactions.TransactionID:   core.NewQByteArray2("transactionID", -1),
		transactions.SentAddress:     core.NewQByteArray2("sentAddress", -1),
		transactions.ReceivedAddress: core.NewQByteArray2("receivedAddress", -1),
		transactions.Inputs:          core.NewQByteArray2("inputs", -1),
		transactions.Outputs:         core.NewQByteArray2("outputs", -1),
	})

	hm.ConnectRowCount(hm.rowCount)
	hm.ConnectData(hm.data)
	hm.ConnectRoleNames(hm.roleNames)

	hm.loadHistory([]string{"example"})
}

func (hm *TransactionList) rowCount(*core.QModelIndex) int {
	return len(hm.Transactions())
}

func (hm *TransactionList) roleNames() map[int]*core.QByteArray {
	return hm.Roles()
}

func (hm *TransactionList) addTransaction(transaction *transactions.TransactionDetails) {
	hm.BeginInsertRows(core.NewQModelIndex(), len(hm.Transactions()), len(hm.Transactions()))
	hm.SetTransactions(append(hm.Transactions(), transaction))
	hm.EndInsertRows()
}

func (hm *TransactionList) removeTransaction(index int) {
	hm.BeginRemoveRows(core.NewQModelIndex(), index, index)
	hm.SetTransactions(append(hm.Transactions()[:index], hm.Transactions()[index+1:]...))
	hm.EndRemoveRows()
}

func (hm *TransactionList) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() >= len(hm.Transactions()) {
		return core.NewQVariant()
	}

	transaction := hm.Transactions()[index.Row()]

	switch role {
	case transactions.Date:
		{
			return core.NewQVariant1(transaction.Date())
		}
	case transactions.Status:
		{
			return core.NewQVariant1(transaction.Status())
		}
	case transactions.Type:
		{
			return core.NewQVariant1(transaction.Type())
		}
	case transactions.Amount:
		{
			return core.NewQVariant1(transaction.Amount())
		}
	case transactions.HoursReceived:
		{
			return core.NewQVariant1(transaction.HoursReceived())
		}
	case transactions.HoursBurned:
		{
			return core.NewQVariant1(transaction.HoursBurned())
		}
	case transactions.TransactionID:
		{
			return core.NewQVariant1(transaction.TransactionID())
		}
	case transactions.SentAddress:
		{
			return core.NewQVariant1(transaction.SentAddress())
		}
	case transactions.ReceivedAddress:
		{
			return core.NewQVariant1(transaction.ReceivedAddress())
		}
	case transactions.Inputs:
		{
			return core.NewQVariant1(transaction.Inputs())
		}
	case transactions.Outputs:
		{
			return core.NewQVariant1(transaction.Outputs())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (hm *TransactionList) loadHistory(wallets []string) {
	td := transactions.NewTransactionDetails(nil)
	td.SetDate(core.NewQDateTime3(core.NewQDate3(2000, 1, 1), core.NewQTime3(10, 0, 0, 0), core.Qt__LocalTime))
	td.SetStatus(transactions.TransactionStatusPending)
	td.SetType(transactions.TransactionTypeSend)
	td.SetAmount(100)
	td.SetHoursReceived(100)
	td.SetHoursBurned(100)
	td.SetTransactionID("transactionID")
	td.SetSentAddress("addrexample1")
	td.SetReceivedAddress("addrexample2")
	inputs := address.NewAddressList(nil)
	inputs.RemoveAddress(0)
	td.SetInputs(inputs)
	td.SetOutputs(address.NewAddressList(nil))
	hm.addTransaction(td)
}
