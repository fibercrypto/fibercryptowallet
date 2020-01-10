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
}
