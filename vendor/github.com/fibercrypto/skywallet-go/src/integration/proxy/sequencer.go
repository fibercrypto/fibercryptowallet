package proxy //nolint goimports

import (
	"errors"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"sync"

	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/skycoin/skycoin/src/util/logging"
)

// Sequencer implementation force all messages to be sequential and make the
// command atomic
type Sequencer struct {
	sync.Mutex
	log *logging.MasterLogger
	logCli *logging.MasterLogger
	dev skywallet.Devicer
	scan func(requestKind skywallet.InputRequestKind, title, message string) (string, error)
}

// ActionCancelableFrom is used for masking options, for example:
// 00000000 ActionConfirmNone
// 00000001 ActionConfirmOkFromDevButton
// 00000010 ActionConfirmCancelFromDevButton
type ActionConfirmFrom uint8
var ActionConfirmNone ActionConfirmFrom = 0x0
var ActionConfirmOkFromDevButton ActionConfirmFrom = 0x1
var ActionConfirmCancelFromDevButton ActionConfirmFrom = 0x2
var ActionConfirmOkAndCancelFromDevButton ActionConfirmFrom = 0x4
var ActionConfirmOkFromWireProtocol ActionConfirmFrom = 0x8
var ActionConfirmCancelFromWireProtocol ActionConfirmFrom = 0x16
var ActionConfirmOkAndCancelFromWireProtocol ActionConfirmFrom = 0x32
var ActionWordRequest ActionConfirmFrom = 0x64

// mixActionConfirmFrom create a merged value from all the masks
func mixActionConfirmFrom(masks ...ActionConfirmFrom) ActionConfirmFrom {
	result := ActionConfirmNone
	for _, mask := range masks {
		result |= mask
	}
	if matchAcf(result, ActionConfirmOkFromDevButton) && matchAcf(result, ActionConfirmCancelFromDevButton) {
		result |= ActionConfirmOkAndCancelFromDevButton
	}
	if matchAcf(result, ActionConfirmOkFromWireProtocol) && matchAcf(result, ActionConfirmCancelFromWireProtocol) {
		result |= ActionConfirmOkAndCancelFromWireProtocol
	}
	if matchAcf(result, ActionWordRequest) {
		result |= ActionConfirmOkAndCancelFromWireProtocol
	}
	return result
}

// matchAcf verify if current have matching enabled by masking
func matchAcf(current, matching ActionConfirmFrom) bool {
	return current & matching == matching
}

// NewSequencer create a new sequencer instance
func NewSequencer(dev skywallet.Devicer, cliSpeechless bool, scanner func(requestKind skywallet.InputRequestKind, title, message string) (string, error)) skywallet.Devicer {
	sq := &Sequencer{
		log: logging.NewMasterLogger(),
		logCli: logging.NewMasterLogger(),
		dev: dev,
		scan: scanner,
	}
	if cliSpeechless {
		sq.logCli.Out = ioutil.Discard
	}
	return sq
}

func (sq *Sequencer) handleInputReaderResponse(err error) error {
	if err == skywallet.ErrUserCancelledWithDeviceButton {
		sq.log.WithError(err).Warningln("action canceled")
		return err
	}
	if err == skywallet.ErrUserCancelledFromInputReader {
		if err := sq.handleInputReaderCanceled(err); err != nil {
			sq.log.WithError(err).Infoln("invalid state")
			return err
		}
		return err
	}
	return err
}

func (sq *Sequencer) handleInputReaderCanceled(err error) error {
	msg, err := sq.dev.Cancel()
	if err != nil {
		sq.log.WithError(err).Errorln("unable to cancel command")
		return err
	}
	msgStr, err := skywallet.DecodeFailMsg(msg)
	if err != nil {
		sq.log.WithError(err).Errorln("unable to decode response")
		return err
	}
	sq.log.Infoln("device response: " + msgStr)
	return nil
}

func (sq *Sequencer) handleInputInteraction(cf ActionConfirmFrom, msg wire.Message) (wire.Message, error) {
	var err error
	handleResponse := func(scopedMsg wire.Message, err error) (string, error) {
		if err != nil {
			sq.log.WithError(err).Errorln("sending message failed" )
			return "", errors.New("sending message failed")
		} else if scopedMsg.Kind == uint16(messages.MessageType_MessageType_Success) {
			msgStr, err := skywallet.DecodeSuccessMsg(scopedMsg)
			if err != nil {
				sq.log.WithError(err).Errorln("unable to decode response")
				return "", errors.New(msgStr)
			}
			return msgStr, nil
		} else if scopedMsg.Kind == uint16(messages.MessageType_MessageType_Failure) {
			msgStr, err := skywallet.DecodeFailMsg(scopedMsg)
			if err != nil {
				sq.log.WithError(err).Errorln("unable to decode response")
				return "", errors.New("unable to decode response")
			}
			return msgStr, nil
		} else {
			return "", errors.New("invalid state")
		}
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) {
		pmr := &messages.PinMatrixRequest{}
		if err = proto.Unmarshal(msg.Data, pmr); err != nil {
			sq.log.WithError(err).Errorln("unable to decode PinMatrixRequest")
			return wire.Message{}, err
		}
		if pmr.Type == nil {
			sq.log.Errorln("pmr.Type should not be null")
			return wire.Message{}, errors.New("unexpected null object")
		}
		str4PinType, err := func(t messages.PinMatrixRequestType) (string, error) {
			switch t {
			case messages.PinMatrixRequestType_PinMatrixRequestType_Current:
				return "enter current pin:", nil
			case messages.PinMatrixRequestType_PinMatrixRequestType_NewFirst:
				return "enter new pin:", nil
			case messages.PinMatrixRequestType_PinMatrixRequestType_NewSecond:
				return "confirm new pin:", nil
			default:
				return "", errors.New("unexpected PinMatrixRequestType")
			}
		}(*pmr.Type)
		if err != nil {
			sq.log.WithField("type", *pmr.Type).Errorln(err.Error())
			return wire.Message{}, err
		}
		pinEnc, err := sq.scan(skywallet.RequestKindPinMatrix, str4PinType, "")
		if err = sq.handleInputReaderResponse(err); err != nil {
			return wire.Message{}, err
		}
		if msg, err = sq.dev.PinMatrixAck(pinEnc); err != nil {
			sq.log.WithError(err).Errorln("pin matrixAck ack: sending message failed")
			return wire.Message{}, err
		}
		return sq.handleInputInteraction(cf, msg)
	} else if msg.Kind == uint16(messages.MessageType_MessageType_PassphraseRequest) {
		passphrase, err := sq.scan(skywallet.RequestKindPassphrase, "PassphraseRequest request:", "passprase TODO chnageit")
		if err = sq.handleInputReaderResponse(err); err != nil {
			return wire.Message{}, err
		}
		msg, err = sq.dev.PassphraseAck(passphrase)
		msgStr, err := handleResponse(msg, err)
		if err != nil {
			sq.log.WithError(err).Errorln("passphrase ack: sending message failed")
			return msg, err
		}
		sq.logCli.Infof("PassphraseAck response:", msgStr)
	} else if msg.Kind == uint16(messages.MessageType_MessageType_WordRequest) {
		cancelableFrom2InputKind := func(cf ActionConfirmFrom) skywallet.InputRequestKind {
			return skywallet.RequestKindWord
		}
		word, err := sq.scan(
			cancelableFrom2InputKind(cf),
			"Word required from device",
			"Look at the device screen and follow the instructions for the required word.")
		if err = sq.handleInputReaderResponse(err); err != nil {
			return wire.Message{}, err
		}
		if msg, err = sq.dev.WordAck(word); err != nil {
			sq.log.WithError(err).Errorln("word ack: sending message failed")
			return msg, err
		}
	} else if msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
		cancelableFrom2InputKind := func(cf ActionConfirmFrom) skywallet.InputRequestKind {
			if matchAcf(cf, ActionConfirmOkAndCancelFromWireProtocol) {
				return skywallet.RequestInformUserOkAndCancel
			}
			if matchAcf(cf, ActionConfirmCancelFromWireProtocol) {
				return skywallet.RequestInformUserOnlyCancel
			}
			return skywallet.RequestInformUserOnlyOk
		}
		_, err := sq.scan(
			cancelableFrom2InputKind(cf),
			"Verify the information in the device",
			"Be careful on checking all the details in the device screen, if all this" +
				" is right, then take the required action to continue...")
		if err = sq.handleInputReaderResponse(err); err != nil {
			return wire.Message{}, err
		}
		if msg, err = sq.dev.ButtonAck(); err != nil {
			sq.log.WithError(err).Errorln("handling message failed")
			return msg, err
		}
	}
	return msg, nil
}

func (sq *Sequencer) handleFirstCommandResponse(cf ActionConfirmFrom, successMsgKind messages.MessageType, commandName string, err error, msg wire.Message) (wire.Message, error) {
	if err != nil {
		if err == skywallet.ErrNoDeviceConnected {
			//sq.log.WithError(err).Infoln(commandName + ": sending message failed")
		} else {
			sq.log.WithError(err).Errorln(commandName + ": sending message failed")
		}
		return wire.Message{}, err
	}
	for msg.Kind != uint16(successMsgKind) && msg.Kind != uint16(messages.MessageType_MessageType_Failure) {
		if msg.Kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) || msg.Kind == uint16(messages.MessageType_MessageType_PassphraseRequest) || msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
			if msg, err = sq.handleInputInteraction(cf, msg); err != nil {
				if err != skywallet.ErrUserCancelledFromInputReader && err != skywallet.ErrUserCancelledWithDeviceButton {
					sq.log.WithError(err).Errorln("error handling interaction")
				}
				return wire.Message{}, err
			}
		}
	}
	return msg, err
}

func (sq *Sequencer) handleFinalResponse(msg wire.Message,  expectedMsgKind messages.MessageType) (wire.Message, error) {
	if msg.Kind == uint16(expectedMsgKind) {
		return msg, nil
	} else if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
		failMsg, err := skywallet.DecodeFailMsg(msg)
		if err != nil {
			sq.log.WithError(err).Errorln("unable to decode response")
			return wire.Message{}, err
		}
		sq.log.WithField("msg", failMsg).Errorln("error handling final response")
		return wire.Message{}, errors.New(failMsg)
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_Cancel) {
		sq.log.Warningln("action canceled")
		return wire.Message{}, skywallet.ErrUserCancelledWithDeviceButton
	}
	sq.log.Errorln("unexpected message")
	return wire.Message{}, errors.New("unexpected message")
}

// AddressGen forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) AddressGen(addressN, startIndex uint32, confirmAddress bool, walletType string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.AddressGen(uint32(addressN), uint32(startIndex), confirmAddress, walletType)
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_ResponseSkycoinAddress, "address gen", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_ResponseSkycoinAddress)
}

// ApplySettings forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) ApplySettings(usePassphrase *bool, label string, language string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.ApplySettings(usePassphrase, label, language)
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "apply settings", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// Backup forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) Backup() (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.Backup()
	confirm := mixActionConfirmFrom(ActionConfirmOkFromDevButton, ActionConfirmCancelFromDevButton, ActionConfirmCancelFromWireProtocol)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "backup", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// Cancel forward the call to Device
func (sq *Sequencer) Cancel() (wire.Message, error) {
	return sq.dev.Cancel()
}

// CheckMessageSignature forward the call to Device
func (sq *Sequencer) CheckMessageSignature(message, signature, address string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.CheckMessageSignature(message, signature, address)
}

// ChangePin forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) ChangePin(removePin *bool) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.ChangePin(removePin)
	confirm := mixActionConfirmFrom(ActionConfirmOkAndCancelFromWireProtocol, ActionConfirmOkAndCancelFromDevButton)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "change pin", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// Connected forward the call to Device
func (sq *Sequencer) Connected() bool {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.Connected()
}

// Available forward the call to Device
func (sq *Sequencer) Available() bool {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.Available()
}

// EraseFirmware erases length bytes in the firmware section in the device
func (sq *Sequencer) EraseFirmware(length uint32) (wire.Message, error) {
	msg, err := sq.dev.EraseFirmware(length)
	confirm := mixActionConfirmFrom(ActionConfirmOkAndCancelFromDevButton, ActionConfirmOkFromWireProtocol)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "erase firmware", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// UploadFirmware upload a new firmware to the device
// erase the old one
func (sq *Sequencer) UploadFirmware(payload []byte, hash [32]byte) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	sq.logCli.Infoln("Length of firmware %d", len(payload))
	if msg, err  := sq.EraseFirmware(uint32(len(payload))); err != nil {
		return msg, err
	}
	msg, err := sq.dev.UploadFirmware(payload, hash)
	confirm := mixActionConfirmFrom(ActionConfirmOkAndCancelFromDevButton, ActionConfirmOkFromWireProtocol)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "upload new firmware", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// GetFeatures forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) GetFeatures() (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.GetFeatures()
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Features, "features", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Features)
}

// GenerateMnemonic forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) GenerateMnemonic(wordCount uint32, usePassphrase bool) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.GenerateMnemonic(wordCount, usePassphrase)
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "generate mnemonic", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// Recovery forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) Recovery(wordCount uint32, usePassphrase *bool, dryRun bool) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.Recovery(wordCount, usePassphrase, dryRun)
	if err != nil {
		sq.log.WithError(err).Errorln("recovery: sending message failed")
		return wire.Message{}, err
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
		msg, err = sq.dev.ButtonAck()
		if err != nil {
			return wire.Message{}, err
		}
	}
	confirm := mixActionConfirmFrom(ActionWordRequest, ActionConfirmOkAndCancelFromWireProtocol)
	for msg.Kind != uint16(messages.MessageType_MessageType_Success) && msg.Kind != uint16(messages.MessageType_MessageType_Failure) {
		if msg.Kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) || msg.Kind == uint16(messages.MessageType_MessageType_PassphraseRequest) || msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) || msg.Kind == uint16(messages.MessageType_MessageType_WordRequest) {
			if msg, err = sq.handleInputInteraction(confirm, msg); err != nil {
				if err != skywallet.ErrUserCancelledFromInputReader && err != skywallet.ErrUserCancelledWithDeviceButton {
					sq.log.WithError(err).Errorln("error handling interaction")
				}
				return wire.Message{}, err
			}
		}
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// SetMnemonic forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) SetMnemonic(mnemonic string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.SetMnemonic(mnemonic)
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "set mnemonic", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// TransactionSign forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) TransactionSign(inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput, walletType string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	// TODO interesting
	//if len(inputs) != len(inputIndex) {
	//	fmt.Println("Every given input hash should have the an inputIndex")
	//	return
	//}
	//if len(outputs) != len(coins) || len(outputs) != len(hours) {
	//	fmt.Println("Every given output should have a coin and hour value")
	//	return
	//}
	msg, err := sq.dev.TransactionSign(inputs, outputs, walletType)
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_ResponseTransactionSign, "sign transaction", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_ResponseTransactionSign)
}

// SignMessage forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) SignMessage(addressN, addressIndex int, message string, walletType string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.SignMessage(1, addressIndex, message, walletType)
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_ResponseSkycoinSignMessage, "sign message", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_ResponseSkycoinSignMessage)
}

// Wipe forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) Wipe() (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.Wipe()
	confirm := mixActionConfirmFrom(ActionConfirmNone)
	msg, err = sq.handleFirstCommandResponse(confirm, messages.MessageType_MessageType_Success, "wipe", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// PinMatrixAck forward the call to Device
func (sq *Sequencer) PinMatrixAck(p string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.PinMatrixAck(p)
}

// WordAck forward the call to Device
func (sq *Sequencer) WordAck(word string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.WordAck(word)
}

// PassphraseAck forward the call to Device
func (sq *Sequencer) PassphraseAck(passphrase string) (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.PassphraseAck(passphrase)
}

// ButtonAck forward the call to Device
func (sq *Sequencer) ButtonAck() (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.ButtonAck()
}

// SetAutoPressButton forward the call to Device
func (sq *Sequencer) SetAutoPressButton(simulateButtonPress bool, simulateButtonType skywallet.ButtonType) error {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.SetAutoPressButton(simulateButtonPress, simulateButtonType)
}

// Close forward the call to Device
func (sq *Sequencer) Close() {
	sq.Lock()
	defer sq.Unlock()
	sq.dev.Close()
}

// Connect forward the call to Device
func (sq *Sequencer) Connect() error {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.Connect()
}

// Disconnect forward the call to Device
func (sq *Sequencer) Disconnect() error {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.Disconnect()
}
