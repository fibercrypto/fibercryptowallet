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
	handleButtonAckSequence func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error)
}

const (
	urnPrefix = "Hardware:SkyWallet:"
)

// HwFirstAddr return the first address in the deterministic sequence if there is a configured
// device connected, error if not device found or some thing fail.
func HwFirstAddr(dev skyWallet.Devicer) (string, error) {
	msg, err := dev.AddressGen(1, 0, false, skyWallet.WalletTypeDeterministic)
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Debugln("error getting address from device")
		return "", fce.ErrHwUnexpected
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_ResponseSkycoinAddress):
		addrs, err := skywallet.DecodeResponseSkycoinAddress(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return "", fce.ErrHwUnexpected
		}
		if len(addrs) != 1 {
			// TODO i18n
			logrus.WithField("addr_len", len(addrs)).Error("unexpected address count in response")
			return "", fce.ErrHwUnexpected
		}
		return addrs[0], nil
	case uint16(messages.MessageType_MessageType_Failure):
		msgData, err := skyWallet.DecodeFailMsg(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return "", fce.ErrHwUnexpected
		}
		logrus.Error(msgData)
		return "", fce.ErrHwUnexpected
	default:
		// TODO i18n
		logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return "", fce.ErrHwUnexpected
	}
}

func hwMatchWallet(hw SkyWallet, wlt core.Wallet) bool {
	firstAddr, err := HwFirstAddr(hw.dev)
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Errorln("unable to get first address from hw")
		return false
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
		// FIXME i18n
		return false, errors.New("a non Skycoin wallet received in ReadyForTxn")
	}
	return hwMatchWallet(sw, wlt), nil
}

// NewSkyWallet create a new sky wallet instance
func NewSkyWallet(dev skyWallet.Devicer, buttonAckHandler func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error)) *SkyWallet {
	return &SkyWallet{
		dev: dev,
		handleButtonAckSequence: buttonAckHandler,
	}
}

func getAllIndexesFromTxn(txn core.Transaction) []int {
	inputs := txn.GetInputs()
	indexes := make([]int, len(inputs), cap(inputs))
	for i := 0; i < len(inputs); i++ {
		indexes[i] = i
	}
	return indexes
}

func getInputs(txn core.Transaction, indexes []int) (inputs []*messages.SkycoinTransactionInput, err error) {
	for _, i := range indexes {
		if i >= len(txn.GetInputs()) {
			return nil, fce.ErrInvalidIndex
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
			return nil, fce.ErrInvalidIndex
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
		// FIXME: should be possible to send the address index
		// find index to check if it is determined as a owned address
		//	transactionOutput.AddressIndex = proto.Uint32(uint32(addressIndex[i]))
		// check out hw implementation
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
		inputs = append(inputs, &transactionOutput)
	}
	return
}

// SignTransaction using hardware wallet
func (sw SkyWallet) SignTransaction(txn core.Transaction, pr core.PasswordReader, indexes []string) (core.Transaction, error) {
	cloned, err := txn.Clone()
	if err != nil {
		logrus.WithError(err).Errorln("error cloning transaction")
		return nil, fce.ErrTxnSignFailure
	}
	tr2Sign, ok := cloned.(core.Transaction)
	if !ok {
		// FIXME i18n
		logrus.Errorln("unable to get cloned object as a core.Transaction")
		return nil, fce.ErrTxnSignFailure
	}
	if sw.dev == nil {
		// TODO i18n
		logrus.Errorln("error creating hardware wallet device handler")
		return nil, fce.ErrTxnSignFailure
	}
	//defer device.Close()
	isFullySigned, err := tr2Sign.IsFullySigned()
	if err != nil {
		return tr2Sign, err
	}
	if isFullySigned {
		// FIXME i18n or named var, this should be used in tests assertions too
		return nil, errors.New("Transaction is fully signed")
	}
	idxs, err := util.StrSlice2IntSlice(indexes)
	if err != nil {
		// FIXME i18n
		logrus.WithError(err).Errorln("unable to get indexes slice as int slice")
		return nil, fce.ErrTxnSignFailure
	}
	if len(idxs) == 0 {
		idxs = getAllIndexesFromTxn(txn)
	}
	transactionInputs, err := getInputs(tr2Sign, idxs)
	if err != nil {
		// FIXME i18n
		logrus.WithError(err).Errorln("unable to get inputs")
		return nil, fce.ErrTxnSignFailure
	}
	transactionOutputs, err := getOutputs(tr2Sign, idxs)
	if err != nil {
		// FIXME i18n
		logrus.WithError(err).Errorln("unable to get outputs")
		return nil, fce.ErrTxnSignFailure
	}
	msg, err := sw.dev.TransactionSign(transactionInputs, transactionOutputs, skyWallet.WalletTypeDeterministic)
	if err != nil {
		logrus.WithError(err).Error("error signing transaction")
		return nil, fce.ErrTxnSignFailure
	}
	if msg.Kind != uint16(messages.MessageType_MessageType_ButtonRequest) {
		if msg.Kind == uint16(messages.MessageType_MessageType_ResponseTransactionSign) {
			return nil, fce.ErrTxnSignFailure
		}
		if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
			msgStr, err := skyWallet.DecodeFailMsg(msg)
			if err != nil {
				// FIXME i18n
				logrus.WithError(err).Errorln("error decoding failed response")
				return nil, fce.ErrTxnSignFailure
			}
			logrus.Errorln(msgStr)
			// FIXME i18n
			return nil, fce.ErrTxnSignFailure
		}
		// FIXME i18n
		logrus.WithField("msgResponse", msg).Errorln("error signing transaction with hardware wallet")
		return nil, fce.ErrTxnSignFailure
	}
	msg, err = sw.handleButtonAckSequence(sw.dev, msg, len(transactionOutputs))
	if err != nil {
		if err == fce.ErrTxnSignFailure {
			logrus.WithError(err).Errorln("failed to sign transaction")
		} else if err == fce.ErrHwSignTransactionCanceled {
			logrus.WithError(err).Warnln("action canceled from device")
		} else {
			logrus.WithError(err).Errorln("unable to sign transaction with device")
			return nil, fce.ErrTxnSignFailure
		}
		return nil, err
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_ResponseTransactionSign) {
		signatures, err := skyWallet.DecodeResponseTransactionSign(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return nil, fce.ErrTxnSignFailure
		}
		for idx, idxSign := range idxs {
			sign := signatures[idx]
			b, err := hex.DecodeString(sign)
			if err != nil {
				logrus.WithError(err).Errorln("error decoding signature")
				return nil, fce.ErrTxnSignFailure
			}
			err = tr2Sign.AddSignature(uint64(idxSign), b)
			if err != nil {
				logrus.WithError(err).Errorln("error setting signature to transaction")
				return nil, fce.ErrTxnSignFailure
			}
		}
	} else if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
		msgStr, err := skyWallet.DecodeFailMsg(msg)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return nil, fce.ErrTxnSignFailure
		}
		logrus.Errorln(msgStr)
		return nil, fce.ErrTxnSignFailure
	} else {
		logrus.WithField("msg", msg).Errorln("unexpected error signing transaction with hw")
		return nil, fce.ErrTxnSignFailure
	}
	return tr2Sign, nil
}

func (sw SkyWallet) getDeviceFeatures() (messages.Features, error) {
	if sw.dev == nil {
		// TODO i18n
		logrus.Error("error, nil hardware wallet device handler")
		return messages.Features{}, fce.ErrHwUnexpected
	}
	// FIXME: This should not be closed, it's a lower level detail (check out in cli tool too).
	// defer sw.dev.Close()
	msg, err := sw.dev.GetFeatures()
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Error("error getting device features")
		return messages.Features{}, fce.ErrHwUnexpected
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_Features):
		features := messages.Features{}
		err = proto.Unmarshal(msg.Data, &features)
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Error("error decoding device response")
			return messages.Features{}, fce.ErrHwUnexpected
		}
		return features, nil
	case uint16(messages.MessageType_MessageType_Failure), uint16(messages.MessageType_MessageType_Success):
		msgData, err := skyWallet.DecodeSuccessOrFailMsg(msg)
		if err != nil {
			logrus.WithError(err).Errorln("error decoding device response")
		} else {
			logrus.Errorln(msgData)
		}
		return messages.Features{}, fce.ErrHwUnexpected
	default:
		// TODO i18n
		logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return messages.Features{}, fce.ErrHwUnexpected
	}
}

// GetSignerUID this signer uid using the hardware wallet label
func (sw SkyWallet) GetSignerUID() (core.UID, error) {
	features, err := sw.getDeviceFeatures()
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Error("unable to get device features")
		return "", fce.ErrHwUnexpected
	}
	return core.UID(*features.DeviceId), nil
}

// GetSignerDescription facilitates a human readable caption identifying this signing strategy
// in urn(https://en.wikipedia.org/wiki/Uniform_Resource_Name) format.
func (sw SkyWallet) GetSignerDescription() (string, error) {
	features, err := sw.getDeviceFeatures()
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Error("unable to get device features")
		return "", fce.ErrHwUnexpected
	}
	return urnPrefix+*features.Label, nil
}
