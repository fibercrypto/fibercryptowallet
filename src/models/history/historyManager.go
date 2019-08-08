package history

import (
	"strconv"
	"time"

	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/transactions"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	qtcore "github.com/therecipe/qt/core"
)

type HistoryManager struct {
	qtcore.QObject

	_ func() `constructor:"init"`

	_ func(addresses map[string]string, filterAddresses []string) []*transactions.TransactionDetails `slot:"loadHistoryWithFilters"`
	_ func(addresses map[string]string)                                                              `slot:"loadHistory"`
}

func (hm *HistoryManager) init() {
	hm.ConnectLoadHistoryWithFilters(hm.loadHistoryWithFilters)
}

func (hm *HistoryManager) loadHistoryWithFilters(addresses map[string]string, filterAddresses []string) []*transactions.TransactionDetails {

	var sent, internally bool
	var traspassedHours, skyAmount uint64

	find := make(map[string]struct{}, len(addresses))
	for _, addr := range addresses {
		find[addr] = struct{}{}
	}

	c := util.NewClient()
	txns, err := c.TransactionsVerbose(filterAddresses)
	if err != nil {
		return nil
	}
	txnsDetails := make([]*transactions.TransactionDetails, 0)
	for _, txn := range txns {
		traspassedHours = 0
		skyAmount = 0
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

		for _, out := range txn.Transaction.Out {
			traspassedHours += out.Hours
			sky, _ := strconv.ParseUint(out.Coins, 10, 64)
			skyAmount += sky
			qOu := address.NewAddressDetails(nil)
			qOu.SetAddress(out.Address)
			qOu.SetAddressSky(out.Coins)
			qOu.SetAddressCoinHours(strconv.FormatUint(out.Hours, 10))
			outputs.AddAddress(qOu)
			if !sent {
				_, ok := find[out.Address]
				if ok {
					txnAddresses.AddAddress(qOu)
				}
			} else if addresses[txn.Transaction.In[0].Address] != out.Address {
				internally = false
			}
		}
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
		txnDetails.SetHoursTraspassed(strconv.FormatUint(traspassedHours, 10))
		txnDetails.SetAmount(strconv.FormatUint(skyAmount, 10))
		txnDetails.SetTransactionID(txn.Transaction.Hash)
		txnsDetails = append(txnsDetails, txnDetails)

	}

	return txnsDetails
}
