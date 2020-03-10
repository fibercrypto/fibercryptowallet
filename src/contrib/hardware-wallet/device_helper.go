package hardware_wallet

import (
	"github.com/chebyrash/promise"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

type DeviceHelper interface {
	// FirstAddress return the first address in the sequence if there is a configured
	// device connected, error if not device found or some thing fail.
	FirstAddress(walletType string) *promise.Promise

	// DeviceMatch check if there is an attached device that matches with the giving wallet
	DeviceMatch(wlt core.Wallet) *promise.Promise

	// ShouldBeSecured check if the device needs to be secured
	ShouldBeSecured() *promise.Promise

	// ShouldBeInitialized check if the device needs to be initialized
	ShouldBeInitialized() *promise.Promise

	// ShouldUploadFirmware check if the device require a new firmware upload
	ShouldUploadFirmware() *promise.Promise

	// IsBootloaderMode check if the device is running in bootloader mode
	IsBootloaderMode() *promise.Promise
}
