package pending

import (
	"strconv"
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
	Mine          = int(qtcore.Qt__UserRole) + 5
)

type PendingTransactionList struct {
	qtcore.QAbstractListModel
	PEX          core.PEX
	WalletEnv    core.WalletEnv

	_ func()                      `constructor:"init"`

	_ map[int]*qtcore.QByteArray  `property:"roles"`
	_ []*PendingTransaction       `property:"transactions"`

	_ func()                      `slot:"loadModel"`
}

type PendingTransaction struct {
	qtcore.QObject
	
	_ string             `property:"sky"`
	_ string             `property:"coinHours"`
	_ *qtcore.QDateTime  `property:"timeStamp"`
	_ string             `property:"transactionID"`
	_ int                `property:"mine"`
}

func (m *PendingTransactionList) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Sky:             qtcore.NewQByteArray2("sky", -1),
		CoinHours:       qtcore.NewQByteArray2("coinHours", -1),
		TimeStamp:       qtcore.NewQByteArray2("timestamp", -1),
		TransactionID:   qtcore.NewQByteArray2("transactionID", -1),
		Mine:            qtcore.NewQByteArray2("mine", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)

	m.ConnectLoadModel(m.loadModel)

	//Set the correct NodeAddress
	addr := "http://127.0.0.1:46480" 
	m.PEX = &skycoin.SkycoinPEX{NodeAddress: addr}
	m.WalletEnv = &skycoin.WalletNode{NodeAddress: addr}

	m.loadModel()
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
	case Mine:
		{
			return qtcore.NewQVariant1(pt.Mine())
		} 
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *PendingTransactionList) loadModel() {
	txns, err := m.PEX.GetTxnPool()
	if err != nil {
		//display an error in qml app when All is selected
		println(err)
		return
	}
	ptModels := make([]*PendingTransaction, 0)
	for txns.Next() {
		ptModel := TransactionToPendingTransaction(txns.Value())
		ptModel.SetMine(0)
		ptModels = append(ptModels, ptModel)
	}
	wallets := m.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		return
	}
	for wallets.Next() {
		crypto := wallets.Value().GetCryptoAccount()
		txns, err2 = crypto.ListPendingTransactions()
		if err2 != nil {
			//display an error in qml app when Mine is selected
			return
		}
		for txns.Next() {
			txn := txns.Value()
			if txn.GetStatus() == core.TXN_STATUS_PENDING {
				ptModel := TransactionToPendingTransaction(txn)
				ptModel.SetMine(1)
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
	iter := skycoin.NewSkycoinTransactionOutputIterator(stxn.GetOutputs())
	sky, coinHours := uint64(0), uint64(0)
	skyErr, coinHoursErr := false, false
	for iter.Next() {
		output := iter.Value()
		val, err := output.GetCoins(skycoin.Sky)
		skyErr = skyErr || (err != nil)
		if !skyErr {
			sky = sky + val
		}
		val, err = output.GetCoins(skycoin.CoinHour)
		coinHoursErr = coinHoursErr || (err != nil)
		if !coinHoursErr {
			coinHours = coinHours + val
		}
	}
	skyAccuracy, err := util.AltcoinQuotient(skycoin.Sky)
	skyErr = skyErr || (err != nil)
	float_sky := ""
	if skyErr {
		float_sky = "NA"
	} else {
		val := float64(sky) / float64(skyAccuracy)
		float_sky = strconv.FormatFloat(val, 'f', -1, 64)
	}
	pt.SetSky(float_sky)
	skychAccuracy, err2 := util.AltcoinQuotient(skycoin.CoinHour)
	coinHoursErr = coinHoursErr || (err2 != nil)
	uint_ch := ""
	if coinHoursErr {
		uint_ch = "NA"
	} else {
		uint_ch = strconv.FormatUint(coinHours / skychAccuracy, 10)
	}
	pt.SetCoinHours(uint_ch)
	return pt
}