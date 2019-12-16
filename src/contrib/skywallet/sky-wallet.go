package hardware

import (
	"encoding/hex"
	"errors"
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/coin"
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/skytypes"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	fce "github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	"github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"sync"
)

var logSkyWallet = logging.MustGetLogger("Skycoin hardware wallet")

type SkyWallet struct {
	wlt core.Wallet
	dev skyWallet.Devicer
	handleButtonAckSequence func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error)
	sequencer *sync.Mutex
}

const (
	urnPrefix = "signer:skywallet:"
)

// HwFirstAddr return the first address in the deterministic sequence if there is a configured
// device connected, error if not device found or some thing fail.
func HwFirstAddr(dev skyWallet.Devicer, derivationType string) (string, error) {
	msg, err := dev.AddressGen(1, 0, false, derivationType)
	if err != nil {
		logSkyWallet.WithError(err).Debugln("error getting address from device")
		return "", fce.ErrHwUnexpected
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_ResponseSkycoinAddress):
		addrs, err := skywallet.DecodeResponseSkycoinAddress(msg)
		if err != nil {
			logSkyWallet.WithError(err).Error("error decoding device response")
			return "", fce.ErrHwUnexpected
		}
		if len(addrs) != 1 {
			logSkyWallet.WithField("addr_len", len(addrs)).Error("unexpected address count in response")
			return "", fce.ErrHwUnexpected
		}
		return addrs[0], nil
	case uint16(messages.MessageType_MessageType_Failure):
		msgData, err := skyWallet.DecodeFailMsg(msg)
		if err != nil {
			logSkyWallet.WithError(err).Error("error decoding device response")
			return "", fce.ErrHwUnexpected
		}
		logSkyWallet.Error(msgData)
		return "", fce.ErrHwUnexpected
	default:
		logSkyWallet.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return "", fce.ErrHwUnexpected
	}
}

func hwMatchWallet(hw SkyWallet, wlt core.Wallet) bool {
	checkForDerivation := func(dt string) bool {
		firstAddr, err := HwFirstAddr(hw.dev, dt)
		if err != nil {
			logSkyWallet.WithError(err).Errorln("unable to get first address from hw")
			return false
		}
		addrs := wlt.GenAddresses(core.AccountAddress, 0, 1, nil)
		if addrs.Next() {
			addr := addrs.Value()
			return addr.String() == firstAddr
		}
		return false
	}
	if checkForDerivation(skyWallet.WalletTypeDeterministic) {
		return true
	}
	return checkForDerivation(skyWallet.WalletTypeBip44)
}

func (sw SkyWallet) ReadyForTxn(wlt core.Wallet, txn core.Transaction) (bool, error) {
	sw.sequencer.Lock()
	defer sw.sequencer.Unlock()
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

// NewSkyWallet create a new sky wallet instance
func NewSkyWallet(wlt core.Wallet, dev skyWallet.Devicer, buttonAckHandler func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error), seq *sync.Mutex) *SkyWallet {
	return &SkyWallet{
		wlt: wlt,
		dev: dev,
		handleButtonAckSequence: buttonAckHandler,
		sequencer: seq,
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

func findValInSlice(val int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

func getInputs(txn coin.Transaction, indexes []int) (inputs []*messages.SkycoinTransactionInput, err error) {
	for idx, input := range txn.In {
		var transactionInput messages.SkycoinTransactionInput
		transactionInput.HashIn = proto.String(input.String())
		if findValInSlice(idx, indexes) {
			transactionInput.Index = proto.Uint32(uint32(idx))
		}
		inputs = append(inputs, &transactionInput)
	}
	return inputs, nil
}

func getAddrIndex(wlt core.Wallet, addr string) (uint32, error) {
	addrs, err := wlt.GetLoadedAddresses()
	if err != nil {
		return 0, err
	}
	for i := uint32(0); addrs.Next(); i++ {
		if addrs.Value().String() == addr {
			return i, nil
		}
	}
	return 0, errors.New("unable to find a matching address from wallet")
}

func getOutputs(wlt core.Wallet, tr coin.Transaction) (inputs []*messages.SkycoinTransactionOutput, err error) {
	for _, output := range tr.Out {
		var transactionOutput messages.SkycoinTransactionOutput
		transactionOutput.Address = proto.String(output.Address.String())
		transactionOutput.Coin = proto.Uint64(output.Coins)
		transactionOutput.Hour = proto.Uint64(output.Hours)
		index, err := getAddrIndex(wlt, output.Address.String())
		if err == nil {
			transactionOutput.AddressIndex = proto.Uint32(index)
		}
		inputs = append(inputs, &transactionOutput)
	}
	return
}

func readableTxn2Transaction(txn skytypes.ReadableTxn) (*coin.Transaction, error) {
	cTxn, err := txn.ToCreatedTransaction()
	if err != nil {
		return nil, err
	}
	skyTxn, err := cTxn.ToTransaction()
	if err != nil {
		return nil, err
	}
	return skyTxn, nil
}

func rawTxn2Transaction (txn skytypes.SkycoinTxn) (*coin.Transaction, error) {
	buf, err := txn.EncodeSkycoinTransaction()
	if err != nil {

	}
	skyTxn, err := coin.DeserializeTransaction(buf)
	if err != nil {

	}
	return &skyTxn, nil
}

func toTransaction(txn core.Transaction) (*coin.Transaction, error) {
	if rTxn, isReadableTxn := txn.(skytypes.ReadableTxn); isReadableTxn {
		t, err := readableTxn2Transaction(rTxn)
		if err != nil {
			logSkyWallet.WithError(err).Errorln("error transforming core.Transaction to coin.Transaction")
			return nil, fce.ErrInvalidTypeAssertion
		}
		return t, nil
	} else {
		// Raw transaction
		unTxn, ok := txn.(*skycoin.SkycoinUninjectedTransaction)
		if !ok {
			logSkyWallet.Errorln("error transforming core.Transaction to coin.Transaction")
			return nil, fce.ErrInvalidTypeAssertion
		}
		t, err := rawTxn2Transaction(unTxn)
		if err != nil {
			logSkyWallet.WithError(err).Errorln("error transforming core.Transaction to coin.Transaction")
			return nil, fce.ErrInvalidTypeAssertion
		}
		return t, nil
	}
}

func coin2Core(txn *coin.Transaction, fee uint64) (core.Transaction, error) {
	unTxn, err := skycoin.NewUninjectedTransaction(txn, fee)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("unable to create uninjected transaction")
		return nil, fce.ErrTxnSignFailure
	}
	return unTxn, nil
}

func (sw *SkyWallet) signTxn(txn *coin.Transaction, idxs []int, dt string) (*coin.Transaction, error) {
	transactionInputs, err := getInputs(*txn, idxs)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("unable to get inputs")
		return nil, fce.ErrTxnSignFailure
	}
	transactionOutputs, err := getOutputs(sw.wlt, *txn)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("unable to get outputs")
		return nil, fce.ErrTxnSignFailure
	}
	msg, err := sw.dev.TransactionSign(transactionInputs, transactionOutputs, dt)
	if err != nil {
		logSkyWallet.WithError(err).Error("error signing transaction")
		return nil, fce.ErrTxnSignFailure
	}
	if msg.Kind != uint16(messages.MessageType_MessageType_ButtonRequest) {
		if msg.Kind == uint16(messages.MessageType_MessageType_ResponseTransactionSign) {
			return nil, fce.ErrTxnSignFailure
		}
		if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
			msgStr, err := skyWallet.DecodeFailMsg(msg)
			if err != nil {
				logSkyWallet.WithError(err).Errorln("error decoding failed response")
				return nil, fce.ErrTxnSignFailure
			}
			logSkyWallet.Errorln(msgStr)
			return nil, fce.ErrTxnSignFailure
		}
		logSkyWallet.WithField("msgResponse", msg).Errorln("error signing transaction with hardware wallet")
		return nil, fce.ErrTxnSignFailure
	}
	msg, err = sw.handleButtonAckSequence(sw.dev, msg, len(transactionOutputs))
	if err != nil {
		if err == fce.ErrTxnSignFailure {
			logSkyWallet.WithError(err).Errorln("failed to sign transaction")
		} else if err == fce.ErrHwSignTransactionCanceled {
			logSkyWallet.WithError(err).Warnln("action canceled from device")
		} else {
			logSkyWallet.WithError(err).Errorln("unable to sign transaction with device")
			return nil, fce.ErrTxnSignFailure
		}
		return nil, err
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_ResponseTransactionSign) {
		signatures, err := skyWallet.DecodeResponseTransactionSign(msg)
		if err != nil {
			logSkyWallet.WithError(err).Error("error decoding device response")
			return nil, fce.ErrTxnSignFailure
		}
		if txn.Sigs == nil {
			logSkyWallet.Warnln("nil slice in transaction signatures detected, creating a new one")
			txn.Sigs = make([]cipher.Sig, len(transactionInputs))
		}
		if len(signatures) != len(transactionInputs) {
			logSkyWallet.WithFields(
				logrus.Fields{
					"signatures_len": len(signatures),
					"transactionInputs_len": len(transactionInputs)}).Errorln("signatures response len should match inputs one")
			return nil, fce.ErrTxnSignFailure
		}
		for idx, sign := range signatures {
			if !findValInSlice(idx, idxs) {
				// NOTE only sign required inputs
				continue
			}
			buf, err := hex.DecodeString(sign)
			if err != nil {
				logSkyWallet.WithError(err).Error("unable to decode signature")
				return nil, fce.ErrTxnSignFailure
			}
			sgn, err := cipher.NewSig(buf)
			if err != nil {
				logSkyWallet.WithError(err).Errorln("unable to get Skycoin address from buffer")
				return nil, errors.New("unable to get Skycoin address from buffer")
			}
			txn.Sigs[idx] = sgn
		}
	} else if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
		msgStr, err := skyWallet.DecodeFailMsg(msg)
		if err != nil {
			logSkyWallet.WithError(err).Error("error decoding device response")
			return nil, fce.ErrTxnSignFailure
		}
		logSkyWallet.Errorln(msgStr)
		return nil, fce.ErrTxnSignFailure
	} else {
		logSkyWallet.WithField("msg", msg).Errorln("unexpected error signing transaction with hw")
		return nil, fce.ErrTxnSignFailure
	}
	return txn, nil
}

func (sw SkyWallet) signTransaction(txn core.Transaction, idxs []int) (core.Transaction, error) {
	fee, err := txn.ComputeFee(params.CoinHoursTicker)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("unable to get fee")
		return nil, fce.ErrTxnSignFailure
	}
	t, err := toTransaction(txn)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("unable to get txn as coin.Transaction")
		return nil, fce.ErrTxnSignFailure
	}
	dt := derivationType(txn)
	signed, err := sw.signTxn(t, idxs, dt)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("unable to sign transaction")
		return nil, fce.ErrTxnSignFailure
	}
	return coin2Core(signed, fee)
}

func derivationType(txn core.Transaction) string {
	wt := skyWallet.WalletTypeDeterministic
	inputs := txn.GetInputs()
	if len(inputs) > 0 {
		addr := inputs[0].GetSpentOutput().GetAddress()
		if addr.IsBip32() {
			wt = skyWallet.WalletTypeBip44
		}
	}
	return wt
}

func verifyInputsGrouping(txn core.Transaction) error {
	areBip32 := derivationType(txn) == skyWallet.WalletTypeBip44
	for _, in := range txn.GetInputs() {
		addr := in.GetSpentOutput().GetAddress()
		if addr.IsBip32() != areBip32 {
			return errors.New("all inputs should be grouped by derivation type")
		}
	}
	return nil
}

// SignTransaction using hardware wallet
func (sw SkyWallet) SignTransaction(txn core.Transaction, pr core.PasswordReader, indexes []string) (core.Transaction, error) {
	sw.sequencer.Lock()
	defer sw.sequencer.Unlock()
	if sw.dev == nil {
		logSkyWallet.Errorln("error creating hardware wallet device handler")
		return nil, fce.ErrTxnSignFailure
	}
	//defer device.Close()
	if err := verifyInputsGrouping(txn); err != nil {
		logSkyWallet.WithError(err).Errorln("unable to sign transaction using skycoin hardware wallet")
		return nil, fce.ErrTxnSignFailure
	}
	isFullySigned, err := txn.IsFullySigned()
	if err != nil {
		return txn, err
	}
	if isFullySigned {
		// TODO named var, this should be used in tests assertions too
		return nil, errors.New("Transaction is fully signed")
	}
	idxs, err := skycoin.GetHashIndices(txn.GetInputs(), indexes)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("unable to get indexes slice as int slice")
		return nil, fce.ErrTxnSignFailure
	}
	if len(idxs) == 0 {
		logSkyWallet.Debugln("not inputs to sign specified, assuming all")
		idxs = getAllIndexesFromTxn(txn)
	}
	signedTxn, err := sw.signTransaction(txn, idxs)
	if err != nil {
		logSkyWallet.WithError(err).Errorln("error signing transaction with device")
		return nil, fce.ErrTxnSignFailure
	}
	return signedTxn, nil
}

func (sw SkyWallet) getDeviceFeatures() (messages.Features, error) {
	if sw.dev == nil {
		logSkyWallet.Error("error, nil hardware wallet device handler")
		return messages.Features{}, fce.ErrHwUnexpected
	}
	// FIXME: This should not be closed, it's a lower level detail (check out in cli tool too).
	// defer sw.dev.Close()
	msg, err := sw.dev.GetFeatures()
	if err != nil {
		logSkyWallet.WithError(err).Error("error getting device features")
		return messages.Features{}, fce.ErrHwUnexpected
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_Features):
		features := messages.Features{}
		err = proto.Unmarshal(msg.Data, &features)
		if err != nil {
			logSkyWallet.WithError(err).Error("error decoding device response")
			return messages.Features{}, fce.ErrHwUnexpected
		}
		return features, nil
	case uint16(messages.MessageType_MessageType_Failure), uint16(messages.MessageType_MessageType_Success):
		msgData, err := skyWallet.DecodeSuccessOrFailMsg(msg)
		if err != nil {
			logSkyWallet.WithError(err).Errorln("error decoding device response")
		} else {
			logSkyWallet.Errorln(msgData)
		}
		return messages.Features{}, fce.ErrHwUnexpected
	default:
		logSkyWallet.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return messages.Features{}, fce.ErrHwUnexpected
	}
}

// GetSignerUID this signer uid using the hardware wallet label
func (sw SkyWallet) GetSignerUID() (core.UID, error) {
	sw.sequencer.Lock()
	defer sw.sequencer.Unlock()
	features, err := sw.getDeviceFeatures()
	if err != nil {
		logSkyWallet.WithError(err).Error("unable to get device features")
		return "", fce.ErrHwUnexpected
	}
	return core.UID(*features.DeviceId), nil
}

// GetSignerDescription facilitates a human readable caption identifying this signing strategy
// in urn(https://en.wikipedia.org/wiki/Uniform_Resource_Name) format.
func (sw SkyWallet) GetSignerDescription() (string, error) {
	sw.sequencer.Lock()
	defer sw.sequencer.Unlock()
	features, err := sw.getDeviceFeatures()
	if err != nil {
		logSkyWallet.WithError(err).Error("unable to get device features")
		return "", fce.ErrHwUnexpected
	}
	return urnPrefix+*features.Label, nil
}

// Type assertions
var (
	_ core.TxnSigner = &SkyWallet{}
)