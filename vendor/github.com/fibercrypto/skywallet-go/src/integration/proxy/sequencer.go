package proxy

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"sync"
)

// Sequencer implementation force all messages to be sequential and make the
// command atomic
type Sequencer struct {
	dev skywallet.Devicer
	m *sync.Mutex
}

func NewSequencer(dev skywallet.Devicer) skywallet.Devicer {
	return &Sequencer{dev:dev, m: &sync.Mutex{}}
}

func (sq *Sequencer) AddressGen(addressN, startIndex uint32, confirmAddress bool, walletType string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	var pinEnc string
	msg, err := sq.dev.AddressGen(addressN, startIndex, confirmAddress, walletType)
	if err != nil {
		return wire.Message{}, err
	}
	for msg.Kind != uint16(messages.MessageType_MessageType_ResponseSkycoinAddress) && msg.Kind != uint16(messages.MessageType_MessageType_Failure) {
		if msg.Kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) {
			// FIXME use a reader from sq
			// fmt.Printf("PinMatrixRequest response: ")
			// fmt.Scanln(&pinEnc)
			pinAckResponse, err := sq.dev.PinMatrixAck(pinEnc)
			if err != nil {
				return wire.Message{}, err
			}
			// TODO this log
			logrus.Errorln("PinMatrixAck response: %s", pinAckResponse)
			continue
		}

		if msg.Kind == uint16(messages.MessageType_MessageType_PassphraseRequest) {
			var passphrase string
			// FIXME use a reader from sq
			//fmt.Printf("Input passphrase: ")
			//fmt.Scanln(&passphrase)
			passphraseAckResponse, err := sq.dev.PassphraseAck(passphrase)
			if err != nil {
				return wire.Message{}, nil
			}
			// TODO this log
			logrus.Errorln("PinMatrixAck response: %s", passphraseAckResponse)
			continue
		}

		if msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
			msg, err = sq.dev.ButtonAck()
			if err != nil {
				return wire.Message{}, err
			}
			continue
		}
	}
	if msg.Kind == uint16(messages.MessageType_MessageType_ResponseSkycoinAddress) {
		return msg, nil
	}
	failMsg, err := skywallet.DecodeFailMsg(msg)
	if err != nil {
		return wire.Message{}, err
	}
	logrus.WithError(err).Errorln(failMsg)
	return wire.Message{}, err
}

func (sq *Sequencer) ApplySettings(usePassphrase *bool, label string, language string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.ApplySettings(usePassphrase, label, language)
}

func (sq *Sequencer) Backup() (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Backup()
}

func (sq *Sequencer) Cancel() (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Cancel()
}

func (sq *Sequencer) CheckMessageSignature(message, signature, address string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.CheckMessageSignature(message, signature, address)
}

func (sq *Sequencer) ChangePin(removePin *bool) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.ChangePin(removePin)
}

func (sq *Sequencer) Connected() bool {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Connected()
}

func (sq *Sequencer) Available() bool {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Available()
}

func (sq *Sequencer) FirmwareUpload(payload []byte, hash [32]byte) error {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.FirmwareUpload(payload, hash)
}

func (sq *Sequencer) GetFeatures() (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	msg, err := sq.dev.GetFeatures()
	if err != nil {
		return wire.Message{}, err
	}
	switch msg.Kind {
	case uint16(messages.MessageType_MessageType_Features):
		return msg, nil
	case uint16(messages.MessageType_MessageType_Failure), uint16(messages.MessageType_MessageType_Success):
		msgData, err := skywallet.DecodeFailMsg(msg)
		if err != nil {
			return wire.Message{}, err
		}
		logrus.Errorln(msgData)
		return wire.Message{}, err
	default:
		logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
		return wire.Message{}, errors.New("unexpected msg")
	}
}

func (sq *Sequencer) GenerateMnemonic(wordCount uint32, usePassphrase bool) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.GenerateMnemonic(wordCount, usePassphrase)
}

func (sq *Sequencer) Recovery(wordCount uint32, usePassphrase *bool, dryRun bool) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Recovery(wordCount, usePassphrase, dryRun)
}

func (sq *Sequencer) SetMnemonic(mnemonic string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.SetMnemonic(mnemonic)
}

func (sq *Sequencer) TransactionSign(inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput, walletType string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
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
	if err != nil {
		return wire.Message{}, err
	}
	for {
		logrus.Errorln("msg.Kind", msg.Kind)
		switch msg.Kind {
		case uint16(messages.MessageType_MessageType_ResponseTransactionSign):
			return msg, nil
		case uint16(messages.MessageType_MessageType_Success):
			return wire.Message{}, errors.New("should end with ResponseTransactionSign request")
		case uint16(messages.MessageType_MessageType_ButtonRequest):
			msg, err = sq.dev.ButtonAck()
			if err != nil {
				return wire.Message{}, err
			}
		case uint16(messages.MessageType_MessageType_PassphraseRequest):
			var passphrase string
			// FIXME
			//fmt.Printf("Input passphrase: ")
			//fmt.Scanln(&passphrase)
			msg, err = sq.dev.PassphraseAck(passphrase)
			if err != nil {
				return wire.Message{}, err
			}
		case uint16(messages.MessageType_MessageType_PinMatrixRequest):
			var pinEnc string
			// FIXME
			//fmt.Printf("PinMatrixRequest response: ")
			//fmt.Scanln(&pinEnc)
			msg, err = sq.dev.PinMatrixAck(pinEnc)
			if err != nil {
				return wire.Message{}, err
			}
		case uint16(messages.MessageType_MessageType_Failure):
			failMsg, err := skywallet.DecodeFailMsg(msg)
			if err != nil {
				return wire.Message{}, err
			}
			logrus.WithError(err).Errorln(failMsg)
			return wire.Message{}, errors.New("failed")
		default:
			logrus.Errorf("received unexpected message type: %s", messages.MessageType(msg.Kind))
			return wire.Message{}, errors.New("unexpected message")
		}
	}
}

func (sq *Sequencer) SignMessage(addressN, addressIndex int, message string, walletType string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.SignMessage(addressN, addressIndex, message, walletType)
}

func (sq *Sequencer) Wipe() (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Wipe()
}

func (sq *Sequencer) PinMatrixAck(p string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.PinMatrixAck(p)
}

func (sq *Sequencer) WordAck(word string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.WordAck(word)
}

func (sq *Sequencer) PassphraseAck(passphrase string) (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.PassphraseAck(passphrase)
}

func (sq *Sequencer) ButtonAck() (wire.Message, error) {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.ButtonAck()
}

func (sq *Sequencer) SetAutoPressButton(simulateButtonPress bool, simulateButtonType skywallet.ButtonType) error {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.SetAutoPressButton(simulateButtonPress, simulateButtonType)
}

func (sq *Sequencer) Close() {
	sq.m.Lock()
	defer sq.m.Unlock()
	sq.dev.Close()
}

func (sq *Sequencer) Connect() error {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Connect()
}

func (sq *Sequencer) Disconnect() error {
	sq.m.Lock()
	defer sq.m.Unlock()
	return sq.dev.Disconnect()
}