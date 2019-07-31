package pending

import (
	"strconv"
	"github.com/therecipe/qt/core"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

const (
	Sky = int(core.Qt__UserRole) + 1
	CoinHours = int(core.Qt__UserRole) + 2
	TimeStamp = int(core.Qt__UserRole) + 3
	TransactionID = int(core.Qt__UserRole) + 4
)

type PendingTransactionList struct {
	core.QAbstractListModel

	_ func()                    `constructor:"init"`

	_ map[int]*core.QByteArray  `property:"roles"`
	_ []*PendingTransaction     `property:"transactions"`

	_ func()                    `slot:"getMine"`
	_ func()                    `slot:"getAll"`
}

type PendingTransaction struct {
	core.QObject
	
	_ float64          `property:"sky"`
	_ uint64           `property:"coinHours"`
	_ *core.QDateTime  `property:"timeStamp"`
	_ string           `property:"transactionID"`
}

func (m *PendingTransactionList) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Sky:             core.NewQByteArray2("sky", -1),
		CoinHours:       core.NewQByteArray2("coinHours", -1),
		TimeStamp:       core.NewQByteArray2("timestamp", -1),
		TransactionID:   core.NewQByteArray2("transactionID", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)

	m.ConnectGetMine(m.getMine)
	m.ConnectGetAll(m.getAll)
	m.getAll()
}


func (m *PendingTransactionList) rowCount(*core.QModelIndex) int {
	return len(m.Transactions())
}

func (m *PendingTransactionList) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *PendingTransactionList) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Transactions()){
		return core.NewQVariant()
	}

	pt := m.Transactions()[index.Row()]

	switch role{
	case Sky:
		{
			return core.NewQVariant1(pt.Sky())
		}
	case CoinHours:
		{
			return core.NewQVariant1(pt.CoinHours())
		}
	case TimeStamp:
		{
			return core.NewQVariant1(pt.TimeStamp())
		}

	case TransactionID:
		{
			return core.NewQVariant1(pt.TransactionID())
		} 
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *PendingTransactionList) getAll() {
	println("getAll")
	c := util.NewClient()
	transactions, err := c.PendingTransactionsVerbose()
	if err != nil {
		return
	}
	ptModels := make([]*PendingTransaction, 0)
	for _, pt := range transactions {
		ptModel := UnconfirmedTransactionToPendingTransaction(&pt)
		ptModels = append(ptModels, ptModel)
	}
	m.SetTransactions(ptModels)
}

func (m *PendingTransactionList) getMine() {
	println("getMine")
	c := util.NewClient()
	wallets, err := c.Wallets()
	if err != nil {
		return
	}
	ptModels := make([]*PendingTransaction, 0)
	for _, w := range wallets {
		response, err := c.WalletUnconfirmedTransactionsVerbose(w.Meta.Filename)
		if err != nil {
			return
		}
		for _, pt := range response.Transactions {
			ptModel := UnconfirmedTransactionToPendingTransaction(&pt)
			ptModels = append(ptModels, ptModel)
		}
	}
	m.SetTransactions(ptModels)
}

func UnconfirmedTransactionToPendingTransaction(ut *readable.UnconfirmedTransactionVerbose) *PendingTransaction {
	pt := NewPendingTransaction(nil)
	year, month, day, h, m, s := util.ParseDate(ut.Received.Unix())
	pt.SetTimeStamp(core.NewQDateTime3(core.NewQDate3(year, month, day), core.NewQTime3(h, m, s, 0), core.Qt__LocalTime))
	pt.SetTransactionID(ut.Transaction.Hash)
	coins, hours := 0.0, uint64(0)
	for _, output := range ut.Transaction.Out {
		outCoin, _ := strconv.ParseFloat(output.Coins, 64)
		coins = coins + outCoin
		hours = hours + output.Hours
	}
	pt.SetSky(coins)
	pt.SetCoinHours(hours)
	return pt
}