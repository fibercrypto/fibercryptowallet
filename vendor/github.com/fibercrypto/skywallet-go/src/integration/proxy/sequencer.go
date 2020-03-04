package proxy //nolint goimports

import (
	"errors"
	"fmt"
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
	scan func()string
}

// NewSequencer create a new sequencer instance
func NewSequencer(dev skywallet.Devicer, cliSpeechless bool, scanner func()string) skywallet.Devicer {
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

func (sq *Sequencer) handleInputInteraction(msg wire.Message) (wire.Message, error) {
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
		switch *pmr.Type {
		case messages.PinMatrixRequestType_PinMatrixRequestType_Current:
			sq.log.Infoln("enter current pin:")
		case messages.PinMatrixRequestType_PinMatrixRequestType_NewFirst:
			sq.log.Infoln("enter new pin:")
		case messages.PinMatrixRequestType_PinMatrixRequestType_NewSecond:
			sq.log.Infoln("confirm new pin:")
		default:
			errStr := "unexpected PinMatrixRequestType"
			sq.log.WithField("type", *pmr.Type).Errorln(errStr)
			return wire.Message{}, errors.New(errStr)
		}
		pinEnc := sq.scan()
		if msg, err = sq.dev.PinMatrixAck(pinEnc); err != nil {
			sq.log.WithError(err).Errorln("pin matrixAck ack: sending message failed")
			return wire.Message{}, err
		}
		return sq.handleInputInteraction(msg)
	} else if msg.Kind == uint16(messages.MessageType_MessageType_PassphraseRequest) {
		sq.log.Println("PassphraseRequest request:")
		passphrase := sq.scan()
		msg, err = sq.dev.PassphraseAck(passphrase)
		msgStr, err := handleResponse(msg, err)
		if err != nil {
			sq.log.WithError(err).Errorln("passphrase ack: sending message failed")
			return msg, err
		}
		sq.logCli.Infof("PassphraseAck response:", msgStr)
	} else if msg.Kind == uint16(messages.MessageType_MessageType_WordRequest) {
		fmt.Printf("Word: ")
		word := sq.scan()
		if msg, err = sq.dev.WordAck(word); err != nil {
			sq.log.WithError(err).Errorln("word ack: sending message failed")
			return msg, err
		}
	} else if msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
		if msg, err = sq.dev.ButtonAck(); err != nil {
			sq.log.WithError(err).Errorln("handling message failed")
			return msg, err
		}
	}
	return msg, nil
}

func (sq *Sequencer) handleFirstCommandResponse(successMsgKind messages.MessageType, commandName string, err error, msg wire.Message) (wire.Message, error) {
	if err != nil {
		sq.log.WithError(err).Errorln(commandName + ": sending message failed")
		return wire.Message{}, err
	}
	for msg.Kind != uint16(successMsgKind) && msg.Kind != uint16(messages.MessageType_MessageType_Failure) {
		if msg.Kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) || msg.Kind == uint16(messages.MessageType_MessageType_PassphraseRequest) || msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
			if msg, err = sq.handleInputInteraction(msg); err != nil {
				sq.log.WithError(err).Errorln("error handling interaction")
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
		sq.log.Errorln(failMsg)
		return wire.Message{}, errors.New(failMsg)
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_ResponseSkycoinAddress, "address gen", err, msg)
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_Success, "apply settings", err, msg)
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_Success, "backup", err, msg)
	if err != nil {
		return wire.Message{}, err
	}
	return sq.handleFinalResponse(msg, messages.MessageType_MessageType_Success)
}

// Cancel forward the call to Device
func (sq *Sequencer) Cancel() (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_Success, "change pin", err, msg)
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

// FirmwareUpload forward the call to Device
func (sq *Sequencer) FirmwareUpload(payload []byte, hash [32]byte) error {
	sq.Lock()
	defer sq.Unlock()
	return sq.dev.FirmwareUpload(payload, hash)
}

// GetFeatures forward the call to Device and handle all the consecutive command as an
// atomic sequence
func (sq *Sequencer) GetFeatures() (wire.Message, error) {
	sq.Lock()
	defer sq.Unlock()
	msg, err := sq.dev.GetFeatures()
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_Features, "features", err, msg)
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_Success, "generate mnemonic", err, msg)
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
	for msg.Kind != uint16(messages.MessageType_MessageType_Success) && msg.Kind != uint16(messages.MessageType_MessageType_Failure) {
		if msg.Kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) || msg.Kind == uint16(messages.MessageType_MessageType_PassphraseRequest) || msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) || msg.Kind == uint16(messages.MessageType_MessageType_WordRequest) {
			if msg, err = sq.handleInputInteraction(msg); err != nil {
				sq.log.WithError(err).Errorln("error handling interaction")
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_Success, "set mnemonic", err, msg)
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_ResponseTransactionSign, "sign transaction", err, msg)
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_ResponseSkycoinSignMessage, "sign message", err, msg)
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
	msg, err = sq.handleFirstCommandResponse(messages.MessageType_MessageType_Success, "wipe", err, msg)
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
