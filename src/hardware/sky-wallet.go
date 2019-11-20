package hardware

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	"github.com/fibercrypto/skywallet-protob/go"
	fce "github.com/fibercrypto/FiberCryptoWallet/src/errors"
)

const (
	urnPrefix = "Hardware:SkyWallet:"
)
type SkyWallet struct {
	dev skyWallet.Devicer
	callback func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error)
}

func (sw SkyWallet) ReadyForTxn(core.Wallet, core.Transaction) (bool, error) {
	panic("implement me")
}

// skyWallet.NewDevice(skyWallet.DeviceTypeUSB),
func NewSkyWallet(dev skyWallet.Devicer, callback func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error)) *SkyWallet {
	return &SkyWallet{
		dev: dev,
		callback: callback,
	}
}

// SignTransaction using hardware wallet
func (sw SkyWallet) SignTransaction(tr core.Transaction, pr core.PasswordReader, s []string) (core.Transaction, error) {
	device := sw.dev
	if device == nil {
		// TODO i18n
		return nil, errors.New("error creating hardware wallet deice handler")
	}
	//defer device.Close()
	signed, err := tr.IsFullySigned()
	if err != nil {
		return tr, err
	}
	if signed {
		return nil, errors.New("Transaction is fully signed")
	}
	var transactionInputs []*messages.SkycoinTransactionInput
	for i, input := range tr.GetInputs() {
		var transactionInput messages.SkycoinTransactionInput
		transactionInput.HashIn = proto.String(input.GetId())
		transactionInput.Index = proto.Uint32(uint32(i))
		transactionInputs = append(transactionInputs, &transactionInput)
	}
	var transactionOutputs []*messages.SkycoinTransactionOutput
	for i, output := range tr.GetOutputs() {
		var transactionOutput messages.SkycoinTransactionOutput
		transactionOutput.Address = proto.String(output.GetId())
		coin, err := output.GetCoins(params.SkycoinTicker)
		if err != nil {
			return nil, err
		}
		transactionOutput.Coin = proto.Uint64(uint64(coin))
		coinHours, err := output.GetCoins(params.CoinHoursTicker)
		if err != nil {
			return nil, err
		}
		transactionOutput.Hour = proto.Uint64(uint64(coinHours))
		transactionOutput.AddressIndex = proto.Uint32(uint32(i))
		transactionOutput.Address = proto.String(output.GetAddress().String())
		println("output.GetAddress()", output.GetAddress().String())

		// TODO find index to check if it is determined as a owned address
		//	transactionOutput.AddressIndex = proto.Uint32(uint32(addressIndex[i]))
		transactionOutputs = append(transactionOutputs, &transactionOutput)
	}
	// TODO use from library
	walletType := "deterministic"
	msg, err := device.TransactionSign(transactionInputs, transactionOutputs, walletType)
	if err != nil {
		logrus.WithError(err).Error("error signing transaction")
		return nil, errors.New("unable to sign transaction")
	}
	msg, err = sw.callback(sw.dev, msg, len(tr.GetOutputs()))
	if err != nil {
		if err == fce.ErrHwSignTransactionFailed {
			logrus.WithError(err).Errorln("failed to sign transaction")
		} else if err == fce.ErrHwSignTransactionCanceled {
			logrus.WithError(err).Warnln("action canceled from device")
		}
		return tr, err
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_ResponseTransactionSign) {
		signatures, err := skyWallet.DecodeResponseTransactionSign(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return tr, err
		}
		for _, sign := range signatures {
			spew.Dump("str", sign)
		}
	} else if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
		msgStr, err := skyWallet.DecodeFailMsg(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return nil, err
		}
		// TODO i18n
		spew.Dump("msgStr", msgStr)
		return nil, err
	} else {
		logrus.WithField("msg", msg).Errorln("unexpected error signing transaction with hw")
		return nil, errors.New("unexpected error signing transaction with hw")
	}
	return tr, nil
}

// GetSignerUID this signer uid using the hardware wallet label
func (sw SkyWallet) GetSignerUID() core.UID {
	device := sw.dev
	if device == nil {
		// TODO i18n
		logrus.Error("error, nil hardware wallet device handler")
		return "undefined"
	}
	// FIXME: This should not be closed, it's a lower level detail.
	// defer device.Close()
	msg, err := device.GetFeatures()
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Error("error getting device features")
		return "undefined"
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_Features):
		features := &messages.Features{}
		err = proto.Unmarshal(msg.Data, features)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return "undefined"
		}
		return core.UID(*features.DeviceId)
	case uint16(messages.MessageType_MessageType_Failure), uint16(messages.MessageType_MessageType_Success):
		msgData, err := skyWallet.DecodeSuccessOrFailMsg(msg)
		if err != nil {
			logrus.WithError(err).Error(msgData)
		}
		// TODO i18n
		return "undefined"
	default:
		// TODO i18n
		logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return "undefined"
	}
}

// GetSignerDescription facilitates a human readable caption identifying this signing strategy
// in urn(https://en.wikipedia.org/wiki/Uniform_Resource_Name) format.
func (sw SkyWallet) GetSignerDescription() string {
	device := sw.dev
	if device == nil {
		// TODO i18n
		logrus.Error("error creating hardware wallet deice handler")
		return "undefined"
	}
	// FIXME this should be solved in the hw-go cli too
	// defer device.Close()
	msg, err := device.GetFeatures()
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Error("error getting device features")
		return "undefined"
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_Features):
		features := &messages.Features{}
		err = proto.Unmarshal(msg.Data, features)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return "undefined"
		}
		return urnPrefix+*features.Label
	case uint16(messages.MessageType_MessageType_Failure), uint16(messages.MessageType_MessageType_Success):
		msgData, err := skyWallet.DecodeSuccessOrFailMsg(msg)
		if err != nil {
			logrus.WithError(err).Error(msgData)
		}
		return "undefined"
	default:
		// TODO i18n
		logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return "undefined"
	}
}
