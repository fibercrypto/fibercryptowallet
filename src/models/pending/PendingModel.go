package pending

import (
	"strconv"
	"github.com/therecipe/qt/core"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/simelo/FiberCryptoWallet/src/util"
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

	_ func()                    `slot:"load"`
}

type PendingTransaction struct {
	core.QObject
	
	_ int              `property:"sky"`
	_ int              `property:"coinHours"`
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

	m.ConnectLoad(m.load)
	m.load()
}

func (m *PendingTransactionList) load() {
	transactions, err := getAll()
	if err != nil {
		return 
	}
	
	m.SetTransactions(transactions)
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

func getAll() ([]*PendingTransaction, error) {
	c := util.NewClient()
	transactions, err := c.PendingTransactionsVerbose()
	if err != nil {
		return nil, err
	}
	ptModels := make([]*PendingTransaction, 0)
	for _, pt := range transactions {
		ptModel := UnconfirmedTransactionToPendingTransaction(&pt)
		ptModels = append(ptModels, ptModel)
	}
	return ptModels, nil
}

func getMine() ([]*PendingTransaction, error) {
	c := util.NewClient()
	wallets, err := c.Wallets()
	if err != nil {
		return nil, err
	}
	ptModels := make([]*PendingTransaction, 0)
	for _, w := range wallets {
		response, err := c.WalletUnconfirmedTransactionsVerbose(w.Meta.Filename)
		if err != nil {
			return nil, err
		}
		for _, pt := range response.Transactions {
			ptModel := UnconfirmedTransactionToPendingTransaction(&pt)
			ptModels = append(ptModels, ptModel)
		}
	}
	return ptModels, nil
}

func UnconfirmedTransactionToPendingTransaction(ut *readable.UnconfirmedTransactionVerbose) *PendingTransaction {
	pt := NewPendingTransaction(nil)
	//pt.SetTimeStamp(ut.Received)
	//year, month, day, h, m, s := util.ParseDate(ut.Received)
	pt.SetTimeStamp(core.NewQDateTime3(core.NewQDate3(2000, 1, 1), core.NewQTime3(10, 0, 0, 0), core.Qt__LocalTime))
	//bs.SetTimeStamp(core.NewQDateTime3(core.NewQDate3(year, month, day), core.NewQTime3(h, m, s, 0), core.Qt__LocalTime))
	pt.SetTransactionID(ut.Transaction.Hash)
	coins, hours := 0, 0
	for _, output := range ut.Transaction.Out {
		outCoin, _ := strconv.Atoi(output.Coins)
		coins = coins + outCoin
		hours = hours + int(output.Hours)
	}
	pt.SetSky(coins)
	pt.SetCoinHours(hours)
	return pt
}