package hardware

import (
	"github.com/chebyrash/promise"
	"github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
	"github.com/fibercrypto/skywallet-go/src/integration/proxy"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
	"sync"
)

type SkyWalletInteraction struct {
	dev skywallet.Devicer
	initializeWasWarn bool
	secureWasWarn bool
}

var once sync.Once
var skyWltInteraction hardware_wallet.DeviceInteraction

// CreateSkyWltInteractionInstanceOnce initialize the device instance.
// Trying to change the device interaction behavior by using
// different arguments does not works because this is a singleton
// like function method.
func CreateSkyWltInteractionInstanceOnce(deviceType skywallet.DeviceType, inputReader func(skywallet.InputRequestKind, string, string)(string, error)) hardware_wallet.DeviceInteraction {
	once.Do(func() {
		skyWltInteraction = &SkyWalletInteraction{
			dev: proxy.NewSequencer(skywallet.NewDevice(deviceType), true, inputReader),
			initializeWasWarn: false,
			secureWasWarn: false,
		}
	})
	return skyWltInteraction
}

// SkyWltInteractionInstance return the shared device interaction instance
func SkyWltInteractionInstance() hardware_wallet.DeviceInteraction {
	if skyWltInteraction == nil {
		logSkyWallet.Errorln("device instance is null, a previous call to CreateSkyWltInteractionInstanceOnce it's required")
	}
	return skyWltInteraction
}

func(wlt *SkyWalletInteraction) AddressGen(addressN, startIndex uint32, confirmAddress bool, walletType string) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.AddressGen(addressN, startIndex, confirmAddress, walletType)
		if err != nil {
			reject(err)
			return
		}
		addresses, err := skywallet.DecodeResponseSkycoinAddress(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(addresses)
	})
}

func(wlt *SkyWalletInteraction) ApplySettings(usePassphrase *bool, label string, language string) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.ApplySettings(usePassphrase, label, language)
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) Backup() *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.Backup()
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) CheckMessageSignature(message, signature, address string) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.CheckMessageSignature(message, signature, address)
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) ChangePin(removePin *bool) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.ChangePin(removePin)
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) IsAvailable() *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve(wlt.dev.Available())
	})
}

func(wlt *SkyWalletInteraction) UploadFirmware(payload []byte, hash [32]byte) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.UploadFirmware(payload, hash)
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) Features() *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.GetFeatures()
		if err != nil {
			reject(err)
			return
		}
		features := messages.Features{}
		if err = proto.Unmarshal(msg.Data, &features); err != nil {
			reject(err)
			return
		}
		resolve(features)
	})
}

func(wlt *SkyWalletInteraction) GenerateMnemonic(wordCount uint32, usePassphrase bool) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.GenerateMnemonic(wordCount, usePassphrase)
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) Recovery(wordCount uint32, usePassphrase *bool, dryRun bool) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.Recovery(wordCount, usePassphrase, dryRun)
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) TransactionSign(inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput, walletType string) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.TransactionSign(inputs, outputs, walletType)
		if err != nil {
			reject(err)
			return
		}
		signatures, err := skywallet.DecodeResponseTransactionSign(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(signatures)
	})
}

func(wlt *SkyWalletInteraction) SignMessage(addressN, addressIndex int, message string, walletType string) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.SignMessage(addressN, addressIndex, message, walletType)
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeResponseSkycoinSignMessage(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func(wlt *SkyWalletInteraction) Wipe() *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.Wipe()
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeSuccessMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

// TODO refactor all success into a reusable function
func (wlt *SkyWalletInteraction) Cancel() *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		msg, err := wlt.dev.Cancel()
		if err != nil {
			reject(err)
			return
		}
		msgStr, err := skywallet.DecodeFailMsg(msg)
		if err != nil {
			reject(err)
			return
		}
		resolve(msgStr)
	})
}

func (wlt *SkyWalletInteraction) SetInitializeWasWarn() {
	wlt.initializeWasWarn = true
}

func (wlt *SkyWalletInteraction) InitializeWasWarn() bool {
	return wlt.initializeWasWarn
}

func (wlt *SkyWalletInteraction) SetSecureWasWarn() {
	wlt.secureWasWarn = true
}

func (wlt *SkyWalletInteraction) SecureWasWarn() bool {
	return wlt.secureWasWarn
}