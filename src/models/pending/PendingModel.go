package pending

import (
	qtcore "github.com/therecipe/qt/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models" //callable as skycoin
)

const (
	Sky           = int(qtcore.Qt__UserRole) + 1
	CoinHours     = int(qtcore.Qt__UserRole) + 2
	TimeStamp     = int(qtcore.Qt__UserRole) + 3
	TransactionID = int(qtcore.Qt__UserRole) + 4
)

type PendingTransactionList struct {
	qtcore.QAbstractListModel
	PEX          core.PEX
	WalletEnv    core.WalletEnv

	_ func()                    `constructor:"init"`

	_ map[int]*qtcore.QByteArray  `property:"roles"`
	_ []*PendingTransaction     `property:"transactions"`

	_ func()                    `slot:"getMine"`
	_ func()                    `slot:"getAll"`
}

type PendingTransaction struct {
	qtcore.QObject
	
	_ uint64             `property:"sky"`
	_ uint64             `property:"coinHours"`
	_ *qtcore.QDateTime  `property:"timeStamp"`
	_ string             `property:"transactionID"`
}

func (m *PendingTransactionList) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Sky:             qtcore.NewQByteArray2("sky", -1),
		CoinHours:       qtcore.NewQByteArray2("coinHours", -1),
		TimeStamp:       qtcore.NewQByteArray2("timestamp", -1),
		TransactionID:   qtcore.NewQByteArray2("transactionID", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)

	m.ConnectGetMine(m.getMine)
	m.ConnectGetAll(m.getAll)

	//Set the correct NodeAddress
	addr := "http://127.0.0.1:37039" 
	m.PEX = &skycoin.SkycoinPEX{NodeAddress: addr}
	m.WalletEnv = &skycoin.WalletNode{NodeAddress: addr}

	m.getMine()
}

func (m *PendingTransactionList) rowCount(*qtcore.QModelIndex) int {
	return len(m.Transactions())
}

func (m *PendingTransactionList) roleNames() map[int]*qtcore.QByteArray {
	return m.Roles()
}

func (m *PendingTransactionList) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}

	if index.Row() >= len(m.Transactions()){
		return qtcore.NewQVariant()
	}

	pt := m.Transactions()[index.Row()]

	switch role{
	case Sky:
		{
			return qtcore.NewQVariant1(pt.Sky())
		}
	case CoinHours:
		{
			return qtcore.NewQVariant1(pt.CoinHours())
		}
	case TimeStamp:
		{
			return qtcore.NewQVariant1(pt.TimeStamp())
		}

	case TransactionID:
		{
			return qtcore.NewQVariant1(pt.TransactionID())
		} 
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *PendingTransactionList) getAll() {
	println("getAll")
	txns, err := m.PEX.GetTxnPool()
	if err != nil {
		println(err)
		return
	}
	ptModels := make([]*PendingTransaction, 0)
	for txns.Next() {
		ptModel := TransactionToPendingTransaction(txns.Value())
		ptModels = append(ptModels, ptModel)
	}
	m.SetTransactions(ptModels)
}

func (m *PendingTransactionList) getMine() {
	println("getMine")
	wallets := m.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		return
	}
	ptModels := make([]*PendingTransaction, 0)
	for wallets.Next() {
		crypto := wallets.Value().GetCryptoAccount()
		txns := crypto.ListTransactions()
		if txns == nil {
			return
		}
		for txns.Next() {
			txn := txns.Value()
			if txn.GetStatus() == core.TXN_STATUS_PENDING {
				ptModel := TransactionToPendingTransaction(txn)
				ptModels = append(ptModels, ptModel)
			}
		}
	}
	m.SetTransactions(ptModels)
}

func TransactionToPendingTransaction(stxn core.Transaction) *PendingTransaction {
	pt := NewPendingTransaction(nil)
	year, month, day, h, m, s := util.ParseDate(int64(stxn.GetTimestamp()))
	pt.SetTimeStamp(qtcore.NewQDateTime3(qtcore.NewQDate3(year, month, day), qtcore.NewQTime3(h, m, s, 0), qtcore.Qt__LocalTime))
	pt.SetTransactionID(stxn.GetId())
	pt.SetSky(stxn.ComputeFee("SKY"))
	pt.SetCoinHours(stxn.ComputeFee("SKYCH"))
	return pt
}