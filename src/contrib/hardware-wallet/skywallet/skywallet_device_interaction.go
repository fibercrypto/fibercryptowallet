package hardware

import (
	"github.com/chebyrash/promise"
	"github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
)

type SkyWalletInteraction struct {
	dev skywallet.Devicer
}

func NewSkyWalletInteraction() hardware_wallet.DeviceInteraction {
	return &SkyWalletInteraction{dev: SkyWltDeviceInstance()}
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

func(wlt *SkyWalletInteraction) FirmwareUpload(payload []byte, hash [32]byte) *promise.Promise {
	return promise.New(func(resolve func(interface{}), reject func(error)) {
		if err := wlt.dev.FirmwareUpload(payload, hash); err != nil {
			reject(err)
			return
		}
		resolve(nil)
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