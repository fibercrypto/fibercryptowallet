package transactions

import (
	coin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models/address"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"strconv"
	"time"
)

var logTransactionDetails = logging.MustGetLogger("TransactionDetails")

func init() {
	TransactionDetails_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionDetail")
}

const (
	Date = int(qtCore.Qt__UserRole) + 1<<iota
	Status
	Type
	Amount
	HoursTraspassed
	HoursBurned
	TransactionID
	Addresses
	Inputs
	Outputs
)

const (
	TransactionStatusConfirmed = iota
	TransactionStatusPending
	TransactionStatusPreview
)

const (
	TransactionTypeSend = iota
	TransactionTypeReceive
	TransactionTypeInternal
	TransactionTypeGeneric
)

type TransactionDetails struct {
	qtCore.QObject
	_ *qtCore.QDateTime    `property:"date"`
	_ int                  `property:"status"`
	_ int                  `property:"type"`
	_ string               `property:"amount"`
	_ string               `property:"hoursTraspassed"`
	_ string               `property:"hoursBurned"`
	_ string               `property:"transactionID"`
	_ *address.AddressList `property:"addresses"`
	_ *address.AddressList `property:"inputs"`
	_ *address.AddressList `property:"outputs"`
}

func NewTransactionDetailFromCoreTransaction(transaction core.Transaction, txType int) (*TransactionDetails, error) {
	txnDetails := NewTransactionDetails(nil)
	t := time.Unix(int64(transaction.GetTimestamp()), 0)

	txnDetails.SetDate(qtCore.NewQDateTime3(qtCore.NewQDate3(t.Year(), int(t.Month()), t.Day()),
		qtCore.NewQTime3(t.Hour(), t.Minute(), 0, 0), qtCore.Qt__LocalTime))
	switch transaction.GetStatus() {
	case core.TXN_STATUS_CONFIRMED:
		txnDetails.SetStatus(TransactionStatusConfirmed)
	case core.TXN_STATUS_PENDING:
		txnDetails.SetStatus(TransactionStatusPending)
	default:
		txnDetails.SetStatus(TransactionStatusPreview)
	}

	txnDetails.SetTransactionID(transaction.GetId())

	if txType >= 0 {
		txnDetails.SetType(txType)
	} else {
		txnDetails.SetType(TransactionTypeGeneric)
	}

	inputs := address.NewAddressList(nil)
	outputs := address.NewAddressList(nil)

	txnIns := transaction.GetInputs()

	for e := range txnIns {

		qIn := address.NewAddressDetails(nil)
		qIn.SetAddress(txnIns[e].GetSpentOutput().GetAddress().String())

		//  TODO Do this generic for all coins
		skyUint64, err := txnIns[e].GetCoins(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins balance")
			continue
		}
		accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins quotient")
			continue
		}
		skyFloat := float64(skyUint64) / float64(accuracy)
		qIn.SetAddressSky(strconv.FormatFloat(skyFloat, 'f', -1, 64))
		chUint64, err := txnIns[e].GetCoins(params.CoinHoursTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours balance")
			continue
		}
		accuracy, err = util.AltcoinQuotient(params.CoinHoursTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours quotient")
			continue
		}
		qIn.SetAddressCoinHours(util.FormatCoins(chUint64, accuracy))
		inputs.AddAddress(qIn)
	}

	txnDetails.SetInputs(inputs)

	for _, out := range transaction.GetOutputs() {
		sky, err := out.GetCoins(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins balance")
			continue
		}
		qOu := address.NewAddressDetails(nil)
		qOu.SetAddress(out.GetAddress().String())
		accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Skycoins quotient")
			continue
		}
		qOu.SetAddressSky(util.FormatCoins(sky, accuracy))
		val, err := out.GetCoins(params.CoinHoursTicker)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours balance")
			continue
		}
		accuracy, err = util.AltcoinQuotient(coin.CoinHour)
		if err != nil {
			logTransactionDetails.WithError(err).Warn("Couldn't get Coin Hours quotient")
			continue
		}
		qOu.SetAddressCoinHours(util.FormatCoins(val, accuracy))
		outputs.AddAddress(qOu)
	}

	txnDetails.SetOutputs(outputs)

	fee, err := transaction.ComputeFee(params.CoinHoursTicker)
	if err != nil {
		logTransactionDetails.WithError(err).Warn("Couldn't compute fee of the operation")
		return nil, err
	}
	accuracy, err := util.AltcoinQuotient(coin.CoinHoursTicker)
	if err != nil {
		logTransactionDetails.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
	}
	txnDetails.SetHoursBurned(util.FormatCoins(fee, accuracy))

	return txnDetails, nil
}
