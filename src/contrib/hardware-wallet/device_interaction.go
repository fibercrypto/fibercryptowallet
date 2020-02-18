package hardware_wallet

import (
	"github.com/chebyrash/promise"
	messages "github.com/fibercrypto/skywallet-protob/go"
)

type DeviceInteraction interface {
	AddressGen(addressN, startIndex uint32, confirmAddress bool, walletType string) *promise.Promise
	ApplySettings(usePassphrase *bool, label string, language string) *promise.Promise
	Backup() *promise.Promise
	CheckMessageSignature(message, signature, address string) *promise.Promise
	ChangePin(removePin *bool) *promise.Promise
	IsAvailable() *promise.Promise
	UploadFirmware(payload []byte, hash [32]byte) *promise.Promise
	Features() *promise.Promise
	GenerateMnemonic(wordCount uint32, usePassphrase bool) *promise.Promise
	Recovery(wordCount uint32, usePassphrase *bool, dryRun bool) *promise.Promise
	// FIXME: change inputs and outputs to the generic types:
	// core.TransactionInputIterator and core.TransactionOutputIterator respectively
	Cancel() *promise.Promise
	TransactionSign(inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput, walletType string) *promise.Promise
	SignMessage(addressN, addressIndex int, message string, walletType string) *promise.Promise
	Wipe() *promise.Promise
	SetInitializeWasWarn()
	InitializeWasWarn() bool
	SetSecureWasWarn()
	SecureWasWarn() bool
}