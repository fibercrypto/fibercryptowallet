package history

import (
	"strconv"
	"time"

	skycoin "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/transactions"
	qtcore "github.com/therecipe/qt/core"
)

type HistoryManager struct {
	qtcore.QObject
	filters []string
	_       func() `constructor:"init"`

	_         func() []*transactions.TransactionDetails `slot:"loadHistoryWithFilters"`
	_         func() []*transactions.TransactionDetails `slot:"loadHistory"`
	_         func(string)                              `slot:"addFilter"`
	_         func(string)                              `slot:"removeFilter"`
	walletEnv core.WalletEnv
}

func (hm *HistoryManager) init() {
	hm.ConnectLoadHistoryWithFilters(hm.loadHistoryWithFilters)
	hm.ConnectLoadHistory(hm.loadHistory)
	hm.ConnectAddFilter(hm.addFilter)
	hm.ConnectRemoveFilter(hm.removeFilter)
	hm.walletEnv = &skycoin.WalletDirectory{WalletDir: "/home/kid/.skycoin/wallets"}
}

func (hm *HistoryManager) loadHistoryWithFilters() []*transactions.TransactionDetails {
	filterAddresses := hm.filters

	addresses := hm.getAddressesWithWallets()

	var sent, internally bool
	var traspassedHoursIn, traspassedHoursOut, skyAmountIn, skyAmountOut uint64

	find := make(map[string]struct{}, len(filterAddresses))
	for _, addr := range filterAddresses {
		find[addr] = struct{}{}
	}
	txnFind := make(map[string]struct{})
	txns := make([]core.Transaction, 0)

	wltIter := hm.walletEnv.GetWalletSet().ListWallets()
	for wltIter.Next() {
		addrIter, _ := wltIter.Value().GetLoadedAddresses()
		for addrIter.Next() {
			_, ok := find[addrIter.Value().String()]
			if ok {
				tnxnsIter := addrIter.Value().GetCryptoAccount().ListTransactions()
				for tnxnsIter.Next() {
					_, ok2 := txnFind[tnxnsIter.Value().GetId()]
					if !ok2 {
						txns = append(txns, tnxnsIter.Value())
						txnFind[tnxnsIter.Value().GetId()] = struct{}{}
					}
				}
			}

		}
	}

	txnsDetails := make([]*transactions.TransactionDetails, 0)
	for _, txn := range txns {
		traspassedHoursIn = 0
		traspassedHoursOut = 0
		skyAmountIn = 0
		skyAmountOut = 0
		internally = true
		sent = false
		txnDetails := transactions.NewTransactionDetails(nil)
		txnAddresses := address.NewAddressList(nil)
		inAddresses := make(map[string]struct{}, 0)
		inputs := address.NewAddressList(nil)
		outputs := address.NewAddressList(nil)
		txnIns := txn.GetInputs()

		for _, in := range txnIns {
			qIn := address.NewAddressDetails(nil)
			qIn.SetAddress(in.GetSpentOutput().GetAddress().String())
			skyUint64 := in.GetCoins("SKY")
			skyFloat := float64(skyUint64) / 1e6
			qIn.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
			chUint64 := in.GetCoins("SKYCH")
			qIn.SetAddressCoinHours(strconv.FormatUint(chUint64, 10))
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

			sky := out.GetCoins("SKY")
			qOu := address.NewAddressDetails(nil)
			qOu.SetAddress(out.GetAddress().String())
			skyFloat := float64(sky) / 1e6
			qOu.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
			qOu.SetAddressCoinHours(strconv.FormatUint(out.GetCoins("SKYCH"), 10))
			outputs.AddAddress(qOu)
			if sent {

				if addresses[txn.GetInputs()[0].GetSpentOutput().GetAddress().String()] == addresses[out.GetAddress().String()] {
					skyAmountOut -= sky

				} else {
					internally = false
					traspassedHoursOut += out.GetCoins("SKYCH")
				}
			} else {
				_, ok := find[out.GetAddress().String()]
				if ok {
					traspassedHoursIn += out.GetCoins("SKYCH")
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
		txnDetails.SetDate(qtcore.NewQDateTime3(qtcore.NewQDate3(t.Year(), int(t.Month()), t.Day()), qtcore.NewQTime3(t.Hour(), t.Minute(), 0, 0), qtcore.Qt__LocalTime))
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
		txnDetails.SetHoursBurned(strconv.FormatUint(txn.ComputeFee("SKYCH"), 10))

		switch txnDetails.Type() {
		case transactions.TransactionTypeReceive:
			{
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursIn, 10))
				val := float64(skyAmountIn)
				val = val / 1000000
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
						hours, _ := strconv.ParseUint(addr.AddressCoinHours(), 10, 64)
						traspassedHoursMoved += hours
						skyf, _ := strconv.ParseFloat(addr.AddressSky(), 64)
						sky := uint64(skyf * 1e6)
						skyAmountMoved += sky
					}

				}
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursMoved, 10))
				val := float64(skyAmountMoved)
				val = val / 1000000
				txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))

			}
		case transactions.TransactionTypeSend:
			{
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursOut, 10))
				val := float64(skyAmountOut)
				val = val / 1000000
				txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))

			}
		}
		txnDetails.SetAddresses(txnAddresses)
		txnDetails.SetTransactionID(txn.GetId())

		txnsDetails = append(txnsDetails, txnDetails)

	}

	return txnsDetails
}

func (hm *HistoryManager) loadHistory() []*transactions.TransactionDetails {

	addresses := hm.getAddressesWithWallets()

	filterAddresses := make([]string, 0)
	for addr, _ := range addresses {
		filterAddresses = append(filterAddresses, addr)
	}

	var sent, internally bool
	var traspassedHoursIn, traspassedHoursOut, skyAmountIn, skyAmountOut uint64

	find := make(map[string]struct{}, len(addresses))
	for addr, _ := range addresses {
		find[addr] = struct{}{}
	}
	txnFind := make(map[string]struct{})
	txns := make([]core.Transaction, 0)
	wltIter := hm.walletEnv.GetWalletSet().ListWallets()
	for wltIter.Next() {
		txnsIter := wltIter.Value().GetCryptoAccount().ListTransactions()
		for txnsIter.Next() {

			_, ok := txnFind[txnsIter.Value().GetId()]
			if !ok {

				txnFind[txnsIter.Value().GetId()] = struct{}{}
				txns = append(txns, txnsIter.Value())
			}

		}
	}

	txnsDetails := make([]*transactions.TransactionDetails, 0)
	for _, txn := range txns {
		traspassedHoursIn = 0
		traspassedHoursOut = 0
		skyAmountIn = 0
		skyAmountOut = 0
		internally = true
		sent = false
		txnDetails := transactions.NewTransactionDetails(nil)
		txnAddresses := address.NewAddressList(nil)
		inAddresses := make(map[string]struct{}, 0)
		inputs := address.NewAddressList(nil)
		outputs := address.NewAddressList(nil)
		txnIns := txn.GetInputs()

		for _, in := range txnIns {
			qIn := address.NewAddressDetails(nil)
			qIn.SetAddress(in.GetSpentOutput().GetAddress().String())
			skyUint64 := in.GetCoins("SKY")
			skyFloat := float64(skyUint64) / 1e6
			qIn.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
			chUint64 := in.GetCoins("SKYCH")
			qIn.SetAddressCoinHours(strconv.FormatUint(chUint64, 10))
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

			sky := out.GetCoins("SKY")
			qOu := address.NewAddressDetails(nil)
			qOu.SetAddress(out.GetAddress().String())
			skyFloat := float64(sky) / 1e6
			qOu.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
			qOu.SetAddressCoinHours(strconv.FormatUint(out.GetCoins("SKYCH"), 10))
			outputs.AddAddress(qOu)
			if sent {

				if addresses[txn.GetInputs()[0].GetSpentOutput().GetAddress().String()] == addresses[out.GetAddress().String()] {
					skyAmountOut -= sky

				} else {
					internally = false
					traspassedHoursOut += out.GetCoins("SKYCH")
				}
			} else {
				_, ok := find[out.GetAddress().String()]
				if ok {
					traspassedHoursIn += out.GetCoins("SKYCH")
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
		txnDetails.SetDate(qtcore.NewQDateTime3(qtcore.NewQDate3(t.Year(), int(t.Month()), t.Day()), qtcore.NewQTime3(t.Hour(), t.Minute(), 0, 0), qtcore.Qt__LocalTime))
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
		txnDetails.SetHoursBurned(strconv.FormatUint(txn.ComputeFee("SKYCH"), 10))

		switch txnDetails.Type() {
		case transactions.TransactionTypeReceive:
			{
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursIn, 10))
				val := float64(skyAmountIn)
				val = val / 1000000
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
						hours, _ := strconv.ParseUint(addr.AddressCoinHours(), 10, 64)
						traspassedHoursMoved += hours
						skyf, _ := strconv.ParseFloat(addr.AddressSky(), 64)
						sky := uint64(skyf * 1e6)
						skyAmountMoved += sky
					}

				}
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursMoved, 10))
				val := float64(skyAmountMoved)
				val = val / 1000000
				txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))

			}
		case transactions.TransactionTypeSend:
			{
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursOut, 10))
				val := float64(skyAmountOut)
				val = val / 1000000
				txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))

			}
		}
		txnDetails.SetAddresses(txnAddresses)
		txnDetails.SetTransactionID(txn.GetId())

		txnsDetails = append(txnsDetails, txnDetails)

	}

	return txnsDetails

}

func (hm *HistoryManager) addFilter(addr string) {
	alreadyIs := false
	for _, filt := range hm.filters {
		if filt == addr {
			alreadyIs = true
			break
		}
	}
	if !alreadyIs {
		hm.filters = append(hm.filters, addr)
	}

}

func (hm *HistoryManager) removeFilter(addr string) {

	for i := 0; i < len(hm.filters); i++ {
		if hm.filters[i] == addr {
			hm.filters = append(hm.filters[0:i], hm.filters[i+1:]...)
			break
		}
	}

}
func (hm *HistoryManager) getAddressesWithWallets() map[string]string {
	response := make(map[string]string, 0)
	it := hm.walletEnv.GetWalletSet().ListWallets()
	for it.Next() {
		wlt := it.Value()
		addrs, _ := wlt.GetLoadedAddresses()

		for addrs.Next() {
			response[addrs.Value().String()] = wlt.GetId()
		}

	}

	return response
}
