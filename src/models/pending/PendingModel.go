package pending

import (
	"strconv"
	qtcore "github.com/therecipe/qt/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models" //callable as skycoin
)

type PendingTransactionList struct {
	qtcore.QObject
	PEX          core.PEX
	WalletEnv    core.WalletEnv

	_ func()                      `constructor:"init"`

	_ []*PendingTransaction       `property:"transactions"`

	_ func()                      `slot:"getAll"`
	_ func()                      `slot:"getMine"`
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
	m.ConnectGetAll(m.getAll)
	m.ConnectGetMine(m.getMine)

	//Set the correct NodeAddress
	addr := "http://127.0.0.1:39247" 
	m.PEX = &skycoin.SkycoinPEX{NodeAddress: addr}
	m.WalletEnv = &skycoin.WalletNode{NodeAddress: addr}

	m.getAll()
}

func (m *PendingTransactionList) getAll() {
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
	m.SetTransactions(ptModels)
}

func (m *PendingTransactionList) getMine() {
	wallets := m.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		return
	}
	ptModels := make([]*PendingTransaction, 0)
	for wallets.Next() {
		crypto := wallets.Value().GetCryptoAccount()
		txns, err := crypto.ListPendingTransactions()
		if err != nil {
			//display an error in qml app when Mine is selected
			println(err)
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