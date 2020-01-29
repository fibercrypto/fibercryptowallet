package history

import (
	"strconv"
	"time"

	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/therecipe/qt/qml"

	"sync"

	coin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	"github.com/fibercrypto/fibercryptowallet/src/models"
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	"github.com/fibercrypto/fibercryptowallet/src/models/transactions"

	"github.com/fibercrypto/fibercryptowallet/src/util"

	qtCore "github.com/therecipe/qt/core"
)

var (
	logHistoryManager = logging.MustGetLogger("modelsHistoryManager")
	historyManager    *HistoryManager
)

const (
	dateTimeFormatForGo  = "2006-01-02T15:04:05"
	dateTimeFormatForQML = "yyyy-MM-ddThh:mm:ss"
)

/*
	HistoryManager
	Represent the controller of history page and all the actions over this page
*/
type HistoryManager struct {
	qtCore.QObject
	walletEnv       core.WalletEnv
	newTxn          map[string][]core.Transaction
	txnFinded       map[string]struct{}
	filters         []string
	txnForAddresses map[string][]core.Transaction
	mutexForNew     sync.Mutex
	mutexForAll     sync.Mutex
	addresses       map[string]string
	walletsIterator core.WalletIterator
	end             chan bool
	_               func()                                    `constructor:"init"`
	_               func()                                    `signal:"newTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getTransactionsWithFilters"`
	_               func() []*transactions.TransactionDetails `slot:"getNewTransactions"`
	_               func() []*transactions.TransactionDetails `slot:"getNewTransactionsWithFilters"`
	_               func(string)                              `slot:"addFilter"`
	_               func(string)                              `slot:"removeFilter"`
	_               func()                                    `slot:"update"`
}

func (hm *HistoryManager) init() {
	hm.ConnectGetTransactions(hm.getTransactions)
	hm.ConnectGetTransactionsWithFilters(hm.getTransansactionsWithFilters)
	hm.ConnectGetNewTransactions(hm.getNewTransactions)
	hm.ConnectGetNewTransactionsWithFilters(hm.getNewTransactionsWithFilters)
	hm.ConnectAddFilter(hm.addFilter)
	hm.ConnectRemoveFilter(hm.removeFilter)
	hm.ConnectUpdate(hm.updateTxns)
	hm.walletEnv = models.GetWalletEnv()

	hm.txnForAddresses = make(map[string][]core.Transaction, 0)
	hm.newTxn = make(map[string][]core.Transaction, 0)
	updateTime := config.GetDataUpdateTime()
	uptimeTicker := time.NewTicker(time.Duration(updateTime) * time.Microsecond * 2)
	historyManager = hm
	hm.txnFinded = make(map[string]struct{}, 0)
	go func() {
		for {
			select {
			case <-uptimeTicker.C:
				logHistoryManager.Debug("Updating history")
				go hm.updateTxns()
				//go hm.reviewForNew()
			}
			historyManager = hm
		}
	}()
}

func (hm *HistoryManager) reviewForNew() {
	hm.mutexForNew.Lock()
	defer hm.mutexForNew.Unlock()
	for _, txns := range hm.newTxn {
		for _, _ = range txns {
			hm.NewTransactions()
			return
		}
	}
}

type ByDate []*transactions.TransactionDetails

func (a ByDate) Len() int {
	return len(a)
}
func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDate) Less(i, j int) bool {
	d1, _ := time.Parse(dateTimeFormatForGo, a[i].Date().ToString(dateTimeFormatForQML))
	d2, _ := time.Parse(dateTimeFormatForGo, a[j].Date().ToString(dateTimeFormatForQML))
	return d1.After(d2)
}

func (hm *HistoryManager) updateTxns() {
	logHistoryManager.Info("Getting transactions of Addresses")

	hm.addresses = hm.getAddressesWithWallets()
	wltIterator := hm.walletEnv.GetWalletSet().ListWallets()
	if wltIterator == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't get transactions of Addresses")
		return
	}
	for wltIterator.Next() {
		logHistoryManager.Debug("Getting addresses history for wallet ", wltIterator.Value().GetId())
		addressIterator, err := wltIterator.Value().GetLoadedAddresses()
		if err != nil {
			logHistoryManager.Warn("Couldn't get address iterator")
			continue
		}
		var newTxnsFinded bool
		for addressIterator.Next() {
			newTxnsFinded = false
			txnsIterator := addressIterator.Value().GetCryptoAccount().ListTransactions()
			if txnsIterator == nil {
				logHistoryManager.Warn("Couldn't get transaction iterator")
				continue
			}
			for txnsIterator.Next() {
				if _, exist := hm.txnFinded[txnsIterator.Value().GetId()]; !exist {
					newTxnsFinded = true
					hm.txnFinded[txnsIterator.Value().GetId()] = struct{}{}
					for _, in := range txnsIterator.Value().GetInputs() {
						if _, exist := hm.addresses[in.GetSpentOutput().GetAddress().String()]; exist {
							hm.mutexForNew.Lock()
							_, exist2 := hm.newTxn[in.GetSpentOutput().GetAddress().String()]
							if exist2 {
								hm.newTxn[in.GetSpentOutput().GetAddress().String()] = append(hm.newTxn[in.GetSpentOutput().GetAddress().String()], txnsIterator.Value())
							} else {
								hm.newTxn[in.GetSpentOutput().GetAddress().String()] = []core.Transaction{txnsIterator.Value()}
							}
							hm.mutexForNew.Unlock()
						}
					}
					for _, out := range txnsIterator.Value().GetOutputs() {
						if _, exist := hm.addresses[out.GetAddress().String()]; exist {
							hm.mutexForNew.Lock()
							_, exist2 := hm.newTxn[out.GetAddress().String()]
							if exist2 {
								hm.newTxn[out.GetAddress().String()] = append(hm.newTxn[out.GetAddress().String()], txnsIterator.Value())
							} else {
								hm.newTxn[out.GetAddress().String()] = []core.Transaction{txnsIterator.Value()}
							}
							hm.mutexForNew.Unlock()
						}
					}
				}
			}
			if newTxnsFinded {
				hm.NewTransactions()
			}

		}
	}

}

func (hm *HistoryManager) getTransactions() []*transactions.TransactionDetails {
	hm.mutexForAll.Lock()

	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, txns := range hm.txnForAddresses {
		for _, txn := range txns {
			if _, exist := hm.txnForAddresses[txn.GetId()]; !exist {
				txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				}
				txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}
		}
	}
	hm.mutexForAll.Unlock()
	newTxns := hm.getNewTransactions()
	txnsForReturn = append(txnsForReturn, newTxns...)
	return txnsForReturn
}

func (hm *HistoryManager) getTransansactionsWithFilters() []*transactions.TransactionDetails {
	hm.mutexForAll.Lock()
	defer hm.mutexForAll.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, addr := range hm.filters {
		for _, txn := range hm.txnForAddresses[addr] {
			if _, exist := added[txn.GetId()]; !exist {
				txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
				}
				txnsForReturn = append(txnsForReturn, txnDetail)
				added[txn.GetId()] = struct{}{}
			}
		}
	}
	return txnsForReturn
}

func (hm *HistoryManager) getNewTransactions() []*transactions.TransactionDetails {
	hm.mutexForNew.Lock()
	defer hm.mutexForNew.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	for addr, _ := range hm.newTxn {
		for _, txn := range hm.newTxn[addr] {
			txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
			}
			txnsForReturn = append(txnsForReturn, txnDetail)
			hm.mutexForAll.Lock()
			_, exist := hm.txnForAddresses[addr]
			if exist {
				hm.txnForAddresses[addr] = append(hm.txnForAddresses[addr], txn)
			} else {
				hm.txnForAddresses[addr] = []core.Transaction{txn}
			}
			hm.mutexForAll.Unlock()
		}
		hm.newTxn[addr] = make([]core.Transaction, 0)
	}
	return txnsForReturn
}

func (hm *HistoryManager) getNewTransactionsWithFilters() []*transactions.TransactionDetails {
	hm.mutexForNew.Lock()
	defer hm.mutexForNew.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	for _, addr := range hm.filters {
		for _, txn := range hm.newTxn[addr] {
			txnDetail, err := TransactionDetailsFromCoreTxn(txn, hm.addresses)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't convert transaction")
			}
			txnsForReturn = append(txnsForReturn, txnDetail)
			hm.mutexForAll.Lock()
			_, exist := hm.txnForAddresses[addr]
			if exist {
				hm.txnForAddresses[addr] = append(hm.txnForAddresses[addr], txn)
			} else {
				hm.txnForAddresses[addr] = []core.Transaction{txn}
			}
			hm.mutexForAll.Unlock()
		}
		hm.newTxn[addr] = make([]core.Transaction, 0)
	}
	return txnsForReturn
}

func (hm *HistoryManager) addFilter(addr string) {
	logHistoryManager.Info("Add filter")
	alreadyIs := false
	for _, filter := range hm.filters {
		if filter == addr {
			alreadyIs = true
			break
		}
	}
	if !alreadyIs {
		hm.filters = append(hm.filters, addr)
	}

}

func (hm *HistoryManager) removeFilter(addr string) {
	logHistoryManager.Info("Remove filter")

	for i := 0; i < len(hm.filters); i++ {
		if hm.filters[i] == addr {
			hm.filters = append(hm.filters[0:i], hm.filters[i+1:]...)
			break
		}
	}

}
func (hm *HistoryManager) getAddressesWithWallets() map[string]string {
	logHistoryManager.Info("Get Addresses with wallets")
	response := make(map[string]string, 0)
	it := hm.walletEnv.GetWalletSet().ListWallets()
	if it == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't load addresses")
		return nil
	}
	for it.Next() {
		wlt := it.Value()
		addresses, err := wlt.GetLoadedAddresses()
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get loaded addresses")
			continue
		}
		for addresses.Next() {
			response[addresses.Value().String()] = wlt.GetId()
		}

	}
	return response
}

func TransactionDetailsFromCoreTxn(txn core.Transaction, addresses map[string]string) (*transactions.TransactionDetails, error) {
	var traspassedHoursIn, traspassedHoursOut, skyAmountIn, skyAmountOut uint64
	traspassedHoursIn = 0
	traspassedHoursOut = 0
	skyAmountIn = 0
	skyAmountOut = 0
	internally := true
	sent := false
	txnDetails := transactions.NewTransactionDetails(nil)
	qml.QQmlEngine_SetObjectOwnership(txnDetails, qml.QQmlEngine__CppOwnership)
	txnAddresses := address.NewAddressList(nil)
	qml.QQmlEngine_SetObjectOwnership(txnAddresses, qml.QQmlEngine__CppOwnership)
	inAddresses := make(map[string]struct{}, 0)
	inputs := address.NewAddressList(nil)
	outputs := address.NewAddressList(nil)
	qml.QQmlEngine_SetObjectOwnership(inputs, qml.QQmlEngine__CppOwnership)
	qml.QQmlEngine_SetObjectOwnership(outputs, qml.QQmlEngine__CppOwnership)
	txnIns := txn.GetInputs()
	for _, in := range txnIns {
		qIn := address.NewAddressDetails(nil)
		qml.QQmlEngine_SetObjectOwnership(qIn, qml.QQmlEngine__CppOwnership)
		qIn.SetAddress(in.GetSpentOutput().GetAddress().String())
		skyUint64, err := in.GetCoins(params.SkycoinTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Skycoins balance")
			return nil, err
		}
		accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Skycoins quotient")

			return nil, err
		}
		skyFloat := float64(skyUint64) / float64(accuracy)
		qIn.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
		chUint64, err := in.GetCoins(params.CalculatedHoursTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours balance")
			return nil, err
		}
		accuracy, err = util.AltcoinQuotient(params.CalculatedHoursTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours quotient")
			return nil, err
		}
		qIn.SetAddressCoinHours(util.FormatCoins(chUint64, accuracy))
		inputs.AddAddress(qIn)
		_, ok := addresses[in.GetSpentOutput().GetAddress().String()]
		if ok {
			skyAmountOut += skyUint64
			sent = true
			_, ok := inAddresses[qIn.Address()]
			if !ok {
				txnAddresses.AddAddress(qIn)
				inAddresses[qIn.Address()] = struct{}{}
			}

		}
	}

	txnDetails.SetInputs(inputs)
	for _, out := range txn.GetOutputs() {
		sky, err := out.GetCoins(params.SkycoinTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Skycoins balance")
			return nil, err
		}
		qOu := address.NewAddressDetails(nil)
		qml.QQmlEngine_SetObjectOwnership(qOu, qml.QQmlEngine__CppOwnership)
		qOu.SetAddress(out.GetAddress().String())
		accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Skycoins quotient")
			return nil, err
		}
		qOu.SetAddressSky(util.FormatCoins(sky, accuracy))
		val, err := out.GetCoins(params.CoinHoursTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours balance")
			return nil, err
		}
		accuracy, err = util.AltcoinQuotient(coin.CoinHour)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours quotient")
			return nil, err
		}
		qOu.SetAddressCoinHours(util.FormatCoins(val, accuracy))
		outputs.AddAddress(qOu)
		if sent {

			if addresses[txn.GetInputs()[0].GetSpentOutput().GetAddress().String()] == addresses[out.GetAddress().String()] {
				skyAmountOut -= sky

			} else {
				internally = false
				val, err = out.GetCoins(params.CoinHoursTicker)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours send it")
					return nil, err
				}
				traspassedHoursOut += val
			}
		} else {
			_, ok := addresses[out.GetAddress().String()]
			if ok {
				val, err = out.GetCoins(params.CoinHoursTicker)
				if err != nil {
					logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours balance")
					return nil, err
				}
				traspassedHoursIn += val
				skyAmountIn += sky

				_, ok := inAddresses[qOu.Address()]
				if !ok {
					txnAddresses.AddAddress(qOu)
					inAddresses[qOu.Address()] = struct{}{}
				}

			}

		}

	}
	txnDetails.SetOutputs(outputs)
	t := time.Unix(int64(txn.GetTimestamp()), 0)
	txnDetails.SetDate(qtCore.NewQDateTime3(qtCore.NewQDate3(t.Year(), int(t.Month()), t.Day()), qtCore.NewQTime3(t.Hour(), t.Minute(), 0, 0), qtCore.Qt__LocalTime))
	txnDetails.SetStatus(transactions.TransactionStatusPending)

	if txn.GetStatus() == core.TXN_STATUS_CONFIRMED {
		txnDetails.SetStatus(transactions.TransactionStatusConfirmed)
	}
	txnDetails.SetType(transactions.TransactionTypeReceive)
	if sent {
		txnDetails.SetType(transactions.TransactionTypeSend)
		if internally {
			txnDetails.SetType(transactions.TransactionTypeInternal)
		}
	}
	fee, err := txn.ComputeFee(params.CoinHoursTicker)
	if err != nil {
		logHistoryManager.WithError(err).Warn("Couldn't compute fee of the operation")
		return nil, err
	}
	accuracy, err := util.AltcoinQuotient(coin.CoinHoursTicker)
	if err != nil {
		logHistoryManager.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
		return nil, err
	}
	txnDetails.SetHoursBurned(util.FormatCoins(fee, accuracy))
	switch txnDetails.Type() {
	case transactions.TransactionTypeReceive:
		{
			accuracy, err := util.AltcoinQuotient(coin.CoinHoursTicker)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
				return nil, err
			}
			txnDetails.SetHoursTraspassed(util.FormatCoins(traspassedHoursIn, accuracy))
			val := float64(skyAmountIn)
			accuracy, err = util.AltcoinQuotient(params.SkycoinTicker)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't get Skycoins quotient")
				return nil, err
			}
			val = val / float64(accuracy)
			txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))

		}
	case transactions.TransactionTypeInternal:
		{
			var traspassedHoursMoved, skyAmountMoved uint64
			traspassedHoursMoved = 0
			skyAmountMoved = 0
			ins := inputs.Addresses()
			inFind := make(map[string]struct{}, len(ins))
			for _, addr := range ins {
				inFind[addr.Address()] = struct{}{}
			}
			outs := outputs.Addresses()
			for _, addr := range outs {
				_, ok := inFind[addr.Address()]
				if !ok {
					hours, err := strconv.ParseUint(addr.AddressCoinHours(), 10, 64)
					if err != nil {
						logHistoryManager.WithError(err).Warn("Couldn't parse Coin Hours from address")
						return nil, err
					}
					traspassedHoursMoved += hours
					skyFloat, err := strconv.ParseFloat(addr.AddressSky(), 64)
					if err != nil {
						logHistoryManager.WithError(err).Warn("Couldn't parse Skycoins from addresses")
						return nil, err
					}
					accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
					if err != nil {
						logHistoryManager.WithError(err).Warn("Couldn't get Skycoins quotient")
						return nil, err
					}
					sky := uint64(skyFloat * float64(accuracy))
					skyAmountMoved += sky
				}

			}
			accuracy, err := util.AltcoinQuotient(coin.CoinHoursTicker)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
				return nil, err
			}
			txnDetails.SetHoursTraspassed(util.FormatCoins(traspassedHoursMoved, accuracy))
			val := float64(skyAmountMoved)
			//FIXME: Error here is skipped
			accuracy, err = util.AltcoinQuotient(params.SkycoinTicker)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't get Skycoins quotient")
				return nil, err
			}
			val = val / float64(accuracy)
			txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))

		}
	case transactions.TransactionTypeSend:
		{
			accuracy, err := util.AltcoinQuotient(coin.CoinHoursTicker)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
				return nil, err
			}
			txnDetails.SetHoursTraspassed(util.FormatCoins(traspassedHoursOut, accuracy))
			val := float64(skyAmountOut)
			accuracy, err = util.AltcoinQuotient(params.SkycoinTicker)
			if err != nil {
				logHistoryManager.WithError(err).Warn("Couldn't get Skycoins quotient")
				return nil, err
			}
			val = val / float64(accuracy)
			txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))

		}
	}
	txnDetails.SetAddresses(txnAddresses)
	txnDetails.SetTransactionID(txn.GetId())
	return txnDetails, nil

}
