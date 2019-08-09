package history

import (
	"fmt"
	"strconv"
	"time"

	skycoin "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/transactions"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	qtcore "github.com/therecipe/qt/core"
)

type HistoryManager struct {
	qtcore.QObject

	_ func() `constructor:"init"`

	_         func(addresses map[string]string, filterAddresses []string) []*transactions.TransactionDetails `slot:"loadHistoryWithFilters"`
	_         func() []*transactions.TransactionDetails                                                      `slot:"loadHistory"`
	walletEnv core.WalletEnv
}

func (hm *HistoryManager) init() {
	hm.ConnectLoadHistoryWithFilters(hm.loadHistoryWithFilters)
	hm.ConnectLoadHistory(hm.loadHistory)
	hm.walletEnv = &skycoin.WalletDirectory{WalletDir: "/home/kid/.skycoin/wallets"}
}

func (hm *HistoryManager) loadHistoryWithFilters(addresses map[string]string, filterAddresses []string) []*transactions.TransactionDetails {

	var sent, internally bool
	var traspassedHoursIn, traspassedHoursOut, skyAmountIn, skyAmountOut uint64

	find := make(map[string]struct{}, len(addresses))
	for addr, _ := range addresses {
		find[addr] = struct{}{}
	}

	c := util.NewClient()
	txns, err := c.TransactionsVerbose(filterAddresses)
	if err != nil {
		return nil
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
		inputs := address.NewAddressList(nil)
		outputs := address.NewAddressList(nil)
		for _, in := range txn.Transaction.In {
			qIn := address.NewAddressDetails(nil)
			qIn.SetAddress(in.Address)
			qIn.SetAddressSky(in.Coins)
			qIn.SetAddressCoinHours(strconv.FormatUint(in.Hours, 10))
			inputs.AddAddress(qIn)
			_, ok := find[in.Address]
			if ok {
				sent = true
				txnAddresses.AddAddress(qIn)
			}
		}
		txnDetails.SetInputs(inputs)

		for _, out := range txn.Transaction.Out {

			sky, _ := strconv.ParseUint(out.Coins, 10, 64)
			//skyAmount += sky
			qOu := address.NewAddressDetails(nil)
			qOu.SetAddress(out.Address)
			qOu.SetAddressSky(out.Coins)
			qOu.SetAddressCoinHours(strconv.FormatUint(out.Hours, 10))
			outputs.AddAddress(qOu)
			_, ok := find[out.Address]
			if ok {
				traspassedHoursIn += out.Hours
				skyAmountIn += sky
			} else {
				traspassedHoursOut += out.Hours
				skyAmountOut += sky
			}
			if addresses[txn.Transaction.In[0].Address] != out.Address {
				internally = false
			}

			//if !sent {
			//	_, ok := find[out.Address]
			//	if ok {
			//		traspassedHoursIn += out.Hours
			//		txnAddresses.AddAddress(qOu)
			//	}
			//} else {
			//	traspassedHoursOut += out.Hours
			//	if addresses[txn.Transaction.In[0].Address] != out.Address {
			//		internally = false
			//	}
			//}
		}
		txnDetails.SetOutputs(outputs)
		t := time.Unix(int64(txn.Time), 0)
		txnDetails.SetDate(qtcore.NewQDateTime3(qtcore.NewQDate3(t.Year(), int(t.Month()), t.Day()), qtcore.NewQTime3(t.Hour(), t.Minute(), 0, 0), qtcore.Qt__LocalTime))
		txnDetails.SetStatus(transactions.TransactionStatusPending)
		if txn.Status.Confirmed {
			txnDetails.SetStatus(transactions.TransactionStatusConfirmed)
		}
		txnDetails.SetType(transactions.TransactionTypeReceive)
		if sent {
			txnDetails.SetType(transactions.TransactionTypeSend)
			if internally {
				txnDetails.SetType(transactions.TransactionTypeInternal)
			}
		}
		txnDetails.SetHoursBurned(strconv.FormatUint(txn.Transaction.Fee, 10))

		switch txnDetails.Type() {
		case transactions.TransactionTypeReceive:
			{
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursIn, 10))
				txnDetails.SetAmount(strconv.FormatUint(skyAmountIn, 10))
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
						sky, _ := strconv.ParseUint(addr.AddressSky(), 10, 64)
						skyAmountMoved += sky
					}

				}
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursMoved, 10))
				txnDetails.SetAmount(strconv.FormatUint(skyAmountMoved, 10))

			}
		case transactions.TransactionTypeSend:
			{
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursOut, 10))
				txnDetails.SetAmount(strconv.FormatUint(skyAmountOut, 10))
			}
		}

		txnDetails.SetTransactionID(txn.Transaction.Hash)
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

	c := util.NewClient()
	txns, err := c.TransactionsVerbose(filterAddresses)
	if err != nil {
		return nil
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
		inputs := address.NewAddressList(nil)
		outputs := address.NewAddressList(nil)
		for _, in := range txn.Transaction.In {
			qIn := address.NewAddressDetails(nil)
			qIn.SetAddress(in.Address)
			qIn.SetAddressSky(in.Coins)
			qIn.SetAddressCoinHours(strconv.FormatUint(in.Hours, 10))
			inputs.AddAddress(qIn)
			_, ok := find[in.Address]
			if ok {
				sent = true
				txnAddresses.AddAddress(qIn)
			}
		}
		txnDetails.SetInputs(inputs)

		for _, out := range txn.Transaction.Out {

			fmt.Println(out.Coins)

			skyf, _ := strconv.ParseFloat(out.Coins, 64)
			sky := uint64(skyf * 1000000)

			qOu := address.NewAddressDetails(nil)
			qOu.SetAddress(out.Address)
			qOu.SetAddressSky(out.Coins)
			qOu.SetAddressCoinHours(strconv.FormatUint(out.Hours, 10))
			outputs.AddAddress(qOu)
			_, ok := find[out.Address]
			if ok {
				traspassedHoursIn += out.Hours
				skyAmountIn += sky
			} else {
				traspassedHoursOut += out.Hours
				skyAmountOut += sky
			}
			if addresses[txn.Transaction.In[0].Address] != out.Address {
				internally = false
			}

		}
		txnDetails.SetOutputs(outputs)
		t := time.Unix(int64(txn.Time), 0)
		txnDetails.SetDate(qtcore.NewQDateTime3(qtcore.NewQDate3(t.Year(), int(t.Month()), t.Day()), qtcore.NewQTime3(t.Hour(), t.Minute(), 0, 0), qtcore.Qt__LocalTime))
		txnDetails.SetStatus(transactions.TransactionStatusPending)
		if txn.Status.Confirmed {
			txnDetails.SetStatus(transactions.TransactionStatusConfirmed)
		}
		txnDetails.SetType(transactions.TransactionTypeReceive)
		if sent {
			txnDetails.SetType(transactions.TransactionTypeSend)
			if internally {
				txnDetails.SetType(transactions.TransactionTypeInternal)
			}
		}
		txnDetails.SetHoursBurned(strconv.FormatUint(txn.Transaction.Fee, 10))

		switch txnDetails.Type() {
		case transactions.TransactionTypeReceive:
			{
				txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHoursIn, 10))
				val := float64(skyAmountIn)
				val = val / 1000000
				txnDetails.SetAmount(strconv.FormatFloat(val, 'f', -1, 64))
				//fmt.Println(skyAmountIn)
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
						sky, _ := strconv.ParseUint(addr.AddressSky(), 10, 64)
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
				//txnDetails.SetAmount(strconv.FormatUint(skyAmountOut, 10))

			}
		}

		txnDetails.SetTransactionID(txn.Transaction.Hash)
		txnsDetails = append(txnsDetails, txnDetails)

	}

	return txnsDetails

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
