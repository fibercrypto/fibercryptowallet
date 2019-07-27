package pending

import (
	"github.com/therecipe/qt/core"
	"github.com/skycoin/skycoin/src/api"
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
		TimeStamp:       core.NewQByteArray2("timeStamp", -1),
		TransactionID:   core.NewQByteArray2("transactionID", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
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

	if index.Row() >= len(m.Wallets()){
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
	c := util.newClient()
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
	c := util.newClient()
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
		for _, pt := range response.transactions {
			ptModel := UnconfirmedTransactionToPendingTransaction(&pt)
			ptModels = append(ptModels, ptModel)
		}
	}
	return ptModels, nil
}

func UnconfirmedTransactionToPendingTransaction(ut *readable.UnconfirmedTransactionVerbose) *PendingTransaction {
	pt := NewPendingTransaction(nil)
	pt.SetTimeStamp(ut.Received)
	pt.SetTransactionID(ut.Transaction.Hash)
	coins, hours := 0, 0
	for _, output := range ut.Transaction.Out {
		coins = coins + output.Coins
		hours = hours + output.Hours
	}
	pt.SetSky(coins)
	pt.SetCoinHours(hours)
	return pt
}