package hardware

import (
	"encoding/hex"
	"errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/skytypes"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	fce "github.com/fibercrypto/FiberCryptoWallet/src/errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	"github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
)

type SkyWallet struct {
	dev skyWallet.Devicer
	callback func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error)
}

const (
	urnPrefix = "Hardware:SkyWallet:"
)

// HwFirstAddr return the first address in the deterministic sequence if there is a configured
// device connected, error if not device found or some thing fail.
func HwFirstAddr() (string, error) {
	dev := skyWallet.NewDevice(skyWallet.DeviceTypeUSB)
	msg, err := dev.AddressGen(1, 0, false, skyWallet.WalletTypeDeterministic)
	if err != nil {
		// TODO i18n
		return "", errors.New("error getting address from device. " + err.Error())
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_ResponseSkycoinAddress):
		addr := &messages.ResponseSkycoinAddress{}
		err = proto.Unmarshal(msg.Data, addr)
		if err != nil {
			// TODO i18n
			strErr := "error decoding device response"
			logrus.WithError(err).Error(strErr)
			return "", errors.New(strErr)
		}
		if len(addr.Addresses) != 1 {
			// TODO i18n
			strErr := "unexpected address count in response"
			logrus.WithField("addr_len", len(addr.Addresses)).Error(strErr)
			return "", errors.New(strErr)
		}
		return addr.Addresses[0], nil
	case uint16(messages.MessageType_MessageType_Failure):
		msgData, err := skyWallet.DecodeFailMsg(msg)
		if err != nil {
			// TODO i18n
			strErr := "error decoding device response"
			logrus.WithError(err).Error(strErr)
			return "", errors.New(strErr)
		}
		// TODO i18n
		strErr := "device response error"
		logrus.Error(msgData, strErr)
		return "", errors.New(strErr)
	default:
		// TODO i18n
		logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return "", errors.New("unexpected device response")
	}
}

func hwMatchWallet(hw SkyWallet, wlt core.Wallet) bool {
	firstAddr, err := HwFirstAddr()
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Errorln("unable to get first address from hw")
	}
	addrs := wlt.GenAddresses(core.AccountAddress, 0, 1, nil)
	if addrs.Next() {
		addr := addrs.Value()
		return addr.String() == firstAddr
	}
	return false
}

func (sw SkyWallet) ReadyForTxn(wlt core.Wallet, txn core.Transaction) (bool, error) {
	if txn != nil {
		_, isSkycoinTxn := txn.(skytypes.SkycoinTxn)
		if !isSkycoinTxn {
			return false, nil
		}
	}
	_, isSkycoinWlt := wlt.(skytypes.SkycoinWallet)
	if !isSkycoinWlt {
		return false, errors.New("a non Skycoin wallet received in ReadyForTxn")
	}
	return hwMatchWallet(sw, wlt), nil
}

// skyWallet.NewDevice(skyWallet.DeviceTypeUSB),
func NewSkyWallet(dev skyWallet.Devicer, callback func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error)) *SkyWallet {
	return &SkyWallet{
		dev: dev,
		callback: callback,
	}
}

func getInputs(txn core.Transaction, indexes []int) (inputs []*messages.SkycoinTransactionInput, err error) {
	for _, i := range indexes {
		if i >= len(txn.GetInputs()) {
			// FIXME report error
			continue
		}
		input := txn.GetInputs()[i]
		var transactionInput messages.SkycoinTransactionInput
		transactionInput.HashIn = proto.String(input.GetId())
		transactionInput.Index = proto.Uint32(uint32(i))
		inputs = append(inputs, &transactionInput)
	}
	return inputs, nil
}

func getOutputs(tr core.Transaction, indexes []int) (inputs []*messages.SkycoinTransactionOutput, err error) {
	for _, i := range indexes {
		if i >= len(tr.GetOutputs()) {
			// FIXME report error
			continue
		}
		output := tr.GetOutputs()[i]
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
		//addrIndex, err := strconv.Atoi(s[i])
		//if err != nil {
		//      logrus.WithField("str_addr_index", s[i]).Error("unable to get integer from string")
		//      return nil, errors.New("error getting address index as integer")
		//}
		//if addrIndex < 0 {
		//      logrus.WithField("addr_index", addrIndex).Error("addrIndex should be greater than 0")
		//      return nil, errors.New("addrIndex should be greater than 0")
		// }
		//transactionOutput.AddressIndex = proto.Uint32(uint32(addrIndex))
		transactionOutput.Address = proto.String(output.GetAddress().String())
		println("output.GetAddress()", output.GetAddress().String())

		// TODO find index to check if it is determined as a owned address
		//	transactionOutput.AddressIndex = proto.Uint32(uint32(addressIndex[i]))
		inputs = append(inputs, &transactionOutput)
	}
	return
}

// SignTransaction using hardware wallet
func (sw SkyWallet) SignTransaction(tr core.Transaction, pr core.PasswordReader, indexes []string) (core.Transaction, error) {
	cloned, err := tr.Clone()
	if err != nil {
		logrus.WithError(err).Errorln("error cloning transaction")
		return nil, fce.ErrHwSignTransactionFailed
	}
	tr2Sign, ok := cloned.(core.Transaction)
	if !ok {
		// FIXME i18n
		logrus.Errorln("unable to get cloned object as a core.Transaction")
		return nil, fce.ErrHwSignTransactionFailed
	}
	device := sw.dev
	if device == nil {
		// TODO i18n
		logrus.Errorln("error creating hardware wallet device handler")
		return nil, fce.ErrHwSignTransactionFailed
	}
	//defer device.Close()
	signed, err := tr2Sign.IsFullySigned()
	if err != nil {
		return tr2Sign, err
	}
	if signed {
		// FIXME i18n or named var, this may be used in tests as is
		return nil, errors.New("Transaction is fully signed")
	}
	idxs, err := util.StrSlice2IntSlice(indexes)
	if err != nil {
		// FIXME i18n
		logrus.WithError(err).Errorln("unable to get indexes slice as int values")
		return nil, err
	}
	transactionInputs, err := getInputs(tr2Sign, idxs)
	if err != nil {
		// FIXME i18n
		logrus.WithError(err).Errorln("unable to get inputs")
		return nil, errors.New("can not sign transaction")
	}
	transactionOutputs, err := getOutputs(tr2Sign, idxs)
	if err != nil {
		// FIXME i18n
		logrus.WithError(err).Errorln("unable to get inputs")
		return nil, errors.New("can not sign transaction")
	}
	msg, err := device.TransactionSign(transactionInputs, transactionOutputs, skyWallet.WalletTypeDeterministic)
	if err != nil {
		logrus.WithError(err).Error("error signing transaction")
		return nil, errors.New("unable to sign transaction")
	}
	if msg.Kind != uint16(messages.MessageType_MessageType_ButtonRequest) {
		if msg.Kind == uint16(messages.MessageType_MessageType_ResponseTransactionSign) {
			s, e := skyWallet.DecodeResponseTransactionSign(msg)
			print("aaa", s, e)
		}
		if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
			msgStr, err := skyWallet.DecodeFailMsg(msg)
			if err != nil {
				// FIXME i18n
				logrus.WithError(err).Errorln("error decoding failed response")
				return nil, errors.New("error signing transaction with hardware wallet")
			}
			logrus.Errorln(msgStr)
			// FIXME i18n
			return nil, errors.New("error signing transaction with hardware wallet")
		}
		// FIXME i18n
		logrus.WithField("msgResponse", msg).Errorln("error signing transaction with hardware wallet")
		return nil, errors.New("error signing transaction with hardware wallet")
	}
	msg, err = sw.callback(sw.dev, msg, len(transactionOutputs))
	if err != nil {
		if err == fce.ErrHwSignTransactionFailed {
			logrus.WithError(err).Errorln("failed to sign transaction")
		} else if err == fce.ErrHwSignTransactionCanceled {
			logrus.WithError(err).Warnln("action canceled from device")
		}
		return tr2Sign, err
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_ResponseTransactionSign) {
		signatures, err := skyWallet.DecodeResponseTransactionSign(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return tr2Sign, err
		}
		for idx, idxSign := range idxs {
			sign := signatures[idx]
			b, err := hex.DecodeString(sign)
			if err != nil {
				logrus.WithError(err).Errorln("error decoding signature")
				return nil, errors.New("error decoding signature")
			}
			err = tr2Sign.AddSignature(uint64(idxSign), b)
			if err != nil {
				logrus.WithError(err).Errorln("error setting signature to transaction")
				return nil, errors.New("error setting signature to transaction")
			}
		}
	} else if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
		msgStr, err := skyWallet.DecodeFailMsg(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return nil, err
		}
		logrus.Errorln(msgStr)
		return nil, fce.ErrTxnSignFailure
	} else {
		logrus.WithField("msg", msg).Errorln("unexpected error signing transaction with hw")
		return nil, errors.New("unexpected error signing transaction with hw")
	}
	return tr2Sign, nil
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
