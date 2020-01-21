package history

import (
	"fmt"
	"strconv"
	"time"

	"github.com/therecipe/qt/qml"

	"github.com/davecgh/go-spew/spew"

	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"

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
	newTxn          []*transactions.TransactionDetails
	txnFinded       map[string]struct{}
	filters         []string
	txnForAddresses map[string][]*transactions.TransactionDetails
	mutexForUpdate  sync.Mutex
	walletsIterator core.WalletIterator
	end             chan bool
	_               func() `constructor:"init"`

	//_ func() []*transactions.TransactionDetails `slot:"loadHistoryWithFilters"`
	//_ func() []*transactions.TransactionDetails `slot:"loadHistory"`
	_ func()                                    `signal:"newTransactions"`
	_ func() []*transactions.TransactionDetails `slot:"getTransactions"`
	_ func() []*transactions.TransactionDetails `slot:"getTransactionsWithFilters"`
	_ func() []*transactions.TransactionDetails `slot:"getNewTransactions"`
	_ func() []*transactions.TransactionDetails `slot:"getNewTransactionsWithFilters"`
	_ func(string)                              `slot:"addFilter"`
	_ func(string)                              `slot:"removeFilter"`
	_ func()                                    `slot:"update"`
}

func (hm *HistoryManager) init() {
	//hm.ConnectLoadHistoryWithFilters(hm.loadHistoryWithFilters)
	//hm.ConnectLoadHistory(hm.loadHistory)
	hm.ConnectGetTransactions(hm.getTransactions)
	hm.ConnectGetTransactionsWithFilters(hm.getTransansactionsWithFilters)
	hm.ConnectGetNewTransactions(hm.getNewTransactions)
	hm.ConnectGetNewTransactionsWithFilters(hm.getNewTransactionsWithFilters)
	hm.ConnectAddFilter(hm.addFilter)
	hm.ConnectRemoveFilter(hm.removeFilter)
	hm.ConnectUpdate(hm.updateTxns)
	hm.walletEnv = models.GetWalletEnv()
	hm.txnForAddresses = make(map[string][]*transactions.TransactionDetails, 0)
	//updateTime := config.GetDataUpdateTime()
	//uptimeTicker := time.NewTicker(time.Duration(updateTime) * time.Microsecond * 2)
	historyManager = hm
	hm.txnFinded = make(map[string]struct{}, 0)
	hm.updateTxns()
	//go func() {
	//	for {
	//		select {
	//		case <-uptimeTicker.C:
	//			logHistoryManager.Debug("Updating history")
	//			go hm.updateTxns()
	//		}
	//		historyManager = hm
	//	}
	//}()
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

//func (hm *HistoryManager) getTransactionsOfAddresses(filterAddresses []string) []*transactions.TransactionDetails {
//
//	txnsDetails := make([]*transactions.TransactionDetails, 0)
//	addrs := make([]string, 0)
//	for _, addr := range filterAddresses {
//		val, ok := hm.txnForAddresses[addr]
//		if !ok {
//			addrs = append(addrs, addr)
//			continue
//		}
//		logHistoryManager.Debug("Getting txns from ", addr)
//		txnsDetails = append(txnsDetails, val...)
//	}
//
//	go hm.updateTxnOfAddresses(addrs)
//
//	return txnsDetails
//}

func (hm *HistoryManager) updateTxns( /*filterAddresses []string*/ ) {
	logHistoryManager.Info("Getting transactions of Addresses")
	hm.mutexForUpdate.Lock()
	defer hm.mutexForUpdate.Unlock()
	addresses := hm.getAddressesWithWallets()
	//var sent, internally bool
	//var traspassedHoursIn, traspassedHoursOut, skyAmountIn, skyAmountOut uint64

	// find := make(map[string]struct{}, len(filterAddresses))
	// for _, addr := range filterAddresses {
	// 	find[addr] = struct{}{}
	// }
	txnFind := make(map[string]struct{})
	txns := make([]core.Transaction, 0)

	wltIterator := hm.walletEnv.GetWalletSet().ListWallets()
	if wltIterator == nil {
		logHistoryManager.WithError(nil).Warn("Couldn't get transactions of Addresses")
		return //make([]*transactions.TransactionDetails, 0)
	}
	for wltIterator.Next() {
		logHistoryManager.Debug("Getting addresses history for wallet ", wltIterator.Value().GetId())
		addressIterator, err := wltIterator.Value().GetLoadedAddresses()
		if err != nil {
			logHistoryManager.Warn("Couldn't get address iterator")
			continue
		}
		for addressIterator.Next() {
			// _, ok := find[addressIterator.Value().String()]
			// if ok {
			txnsIterator := addressIterator.Value().GetCryptoAccount().ListTransactions()
			if txnsIterator == nil {
				logHistoryManager.Warn("Couldn't get transaction iterator")
				continue
			}
			for txnsIterator.Next() {
				_, ok2 := txnFind[txnsIterator.Value().GetId()]
				if !ok2 {
					txns = append(txns, txnsIterator.Value())
					txnFind[txnsIterator.Value().GetId()] = struct{}{}
				}
			}
			//}

		}
	}

	txnsDetails := make([]*transactions.TransactionDetails, 0)
	for _, txn := range txns {
		if _, ok := hm.txnFinded[txn.GetId()]; ok {
			continue
		} else {
			hm.txnFinded[txn.GetId()] = struct{}{}
		}
		txnDetail, err := TransactionDetailsFromCoreTxn(txn, addresses)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't conver core transaction")
		}
		txnsDetails = append(txnsDetails, txnDetail)
	}
	//sort.Sort(ByDate(txnsDetails))

	hm.newTxn = append(hm.newTxn, txnsDetails...)
	// If there is at least one new txn send a signal
	if len(txnsDetails) >= 1 {
		hm.NewTransactions()
	}
}

func (hm *HistoryManager) getTransactions() []*transactions.TransactionDetails {
	hm.mutexForUpdate.Lock()
	defer hm.mutexForUpdate.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	fmt.Println("AAAAAAAAAAAAAAAA")
	spew.Dump(hm.txnForAddresses)
	for _, txns := range hm.txnForAddresses {
		fmt.Println(len(txns))
		for _, txn := range txns {
			if _, exist := hm.txnForAddresses[txn.TransactionID()]; !exist {
				fmt.Println("Adding")
				txnsForReturn = append(txnsForReturn, txn)
				added[txn.TransactionID()] = struct{}{}
			}
		}
	}
	return txnsForReturn
}

func (hm *HistoryManager) getTransansactionsWithFilters() []*transactions.TransactionDetails {
	hm.mutexForUpdate.Lock()
	defer hm.mutexForUpdate.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	added := make(map[string]struct{}, 0)
	for _, addr := range hm.filters {
		for _, txn := range hm.txnForAddresses[addr] {
			if _, exist := added[txn.TransactionID()]; !exist {
				txnsForReturn = append(txnsForReturn, txn)
				added[txn.TransactionID()] = struct{}{}
			}
		}
	}
	return txnsForReturn
}

func (hm *HistoryManager) getNewTransactions() []*transactions.TransactionDetails {
	hm.mutexForUpdate.Lock()
	defer hm.mutexForUpdate.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	for _, txn := range hm.newTxn {
		txnsForReturn = append(txnsForReturn, txn)
		addrs := txn.Addresses()
		for _, addr := range addrs.Addresses() {
			hm.txnForAddresses[addr.Address()] = append(hm.txnForAddresses[addr.Address()], txn)
		}
	}
	hm.newTxn = make([]*transactions.TransactionDetails, 0)
	return txnsForReturn
}

func (hm *HistoryManager) getNewTransactionsWithFilters() []*transactions.TransactionDetails {
	hm.mutexForUpdate.Lock()
	defer hm.mutexForUpdate.Unlock()
	txnsForReturn := make([]*transactions.TransactionDetails, 0)
	for _, txn := range hm.newTxn {
		forReturn := false
		addrs := txn.Addresses()
		for _, addr := range addrs.Addresses() {
			hm.txnForAddresses[addr.Address()] = append(hm.txnForAddresses[addr.Address()], txn)
			for _, fAddr := range hm.filters {
				if fAddr == addr.Address() {
					forReturn = true
					break
				}
			}
		}
		if forReturn {
			txnsForReturn = append(txnsForReturn, txn)
		}
	}
	return txnsForReturn
}

//func (hm *HistoryManager) loadHistoryWithFilters() []*transactions.TransactionDetails {
//	logHistoryManager.Info("Loading history with some filters")
//	filterAddresses := hm.filters
//	return hm.getTransactionsOfAddresses(filterAddresses)
//
//}

//func (hm *HistoryManager) loadHistory() []*transactions.TransactionDetails {
//	logHistoryManager.Info("Loading history")
//	hm.getAddressesWithWallets()
//	addresses := hm.addresses
//
//	filterAddresses := make([]string, 0)
//	for addr, _ := range addresses {
//		filterAddresses = append(filterAddresses, addr)
//	}
//	logHistoryManager.WithField("Addresses to filter", filterAddresses).Debug("Addresses to filter")
//	return hm.getTransactionsOfAddresses(filterAddresses)
//
//}

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
	inAddresses := make(map[string]struct{}, 0)
	inputs := address.NewAddressList(nil)
	outputs := address.NewAddressList(nil)
	qml.QQmlEngine_SetObjectOwnership(inputs, qml.QQmlEngine__CppOwnership)
	qml.QQmlEngine_SetObjectOwnership(outputs, qml.QQmlEngine__CppOwnership)
	txnIns := txn.GetInputs()
	//spew.Dump(txnIns)
	for _, in := range txnIns {
		qIn := address.NewAddressDetails(nil)
		qml.QQmlEngine_SetObjectOwnership(qIn, qml.QQmlEngine__CppOwnership)
		qIn.SetAddress(in.GetSpentOutput().GetAddress().String())
		skyUint64, err := in.GetCoins(params.SkycoinTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Skycoins balance")
			fmt.Println(1)
			return nil, err
		}
		accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Skycoins quotient")
			fmt.Println(2)
			return nil, err
		}
		skyFloat := float64(skyUint64) / float64(accuracy)
		qIn.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
		chUint64, err := in.GetCoins(params.CoinHoursTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours balance")
			fmt.Println(3)
			return nil, err
		}
		accuracy, err = util.AltcoinQuotient(params.CoinHoursTicker)
		if err != nil {
			logHistoryManager.WithError(err).Warn("Couldn't get Coin Hours quotient")
			fmt.Println(4)
			return nil, err
		}
		qIn.SetAddressCoinHours(util.FormatCoins(chUint64, accuracy))
		inputs.AddAddress(qIn)
		fmt.Println(5)
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
	//fmt.Println(txn.GetId())
	//fmt.Println(txnDetails.TransactionID())
	//spew.Dump(txnDetails)
	return txnDetails, nil
	//txnsDetails = append(txnsDetails, txnDetails)

	//for _, addrInTxn := range txnAddresses.Addresses() {
	//	hm.txnForAddresses[addrInTxn.Address()] = append(hm.txnForAddresses[addrInTxn.Address()], txnDetails)
	//}

}
