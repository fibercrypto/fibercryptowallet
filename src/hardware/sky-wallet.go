package hardware

import (
	"errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	skyWallet "github.com/skycoin/hardware-wallet-go/src/skywallet"
	"github.com/skycoin/hardware-wallet-protob/go"
)

const (
	urnPrefix = "Hardware:SkyWallet:"
)
type SkyWallet struct {
	dev skyWallet.Devicer
}

// skyWallet.NewDevice(skyWallet.DeviceTypeUSB),
func NewSkyWallet(dev skyWallet.Devicer) *SkyWallet {
	return &SkyWallet{
		dev: dev,
	}
}

// SignTransaction using hardware wallet
func (sw SkyWallet) SignTransaction(tr core.Transaction, pr core.PasswordReader, s []string) (core.Transaction, error) {
	device := skyWallet.NewDevice(skyWallet.DeviceTypeUSB)
	if device == nil {
		return nil, errors.New("error creating hardware wallet deice handler") // TODO i18n
	}
	defer device.Close()
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

		// TODO find index to check if it is determined as a owned address
		//	transactionOutput.AddressIndex = proto.Uint32(uint32(addressIndex[i]))
		transactionOutputs = append(transactionOutputs, &transactionOutput)
	}
	walletType := "deterministic" // todo use from library
	msg, err := device.TransactionSign(transactionInputs, transactionOutputs, walletType)
	if err != nil {
		msgStr, err1 := skyWallet.DecodeFailMsg(msg)
		if err1 == nil {
			logrus.WithFields(logrus.Fields{"transaction": tr, "error": err}).Error(msgStr)
			return nil, err
		}
		logrus.WithError(err1).Error("error decoding device response")
		return nil, err
	}
	return tr, nil
}

// GetSignerUID this signer uid using the hardware wallet id
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
		logrus.WithError(err).Error("error getting device features")
		return "undefined"
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_Features):
		features := &messages.Features{}
		err = proto.Unmarshal(msg.Data, features)
		if err != nil {
			logrus.WithError(err).Error("error decoding device response")
			return "undefined" // i18n
		}
		return core.UID(*features.DeviceId)
	case uint16(messages.MessageType_MessageType_Failure), uint16(messages.MessageType_MessageType_Success):
		msgData, err := skyWallet.DecodeSuccessOrFailMsg(msg)
		if err != nil {
			logrus.WithError(err).Error(msgData)
		}
		return "undefined" // i18n
	default:
		logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return "undefined" // i18n
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
