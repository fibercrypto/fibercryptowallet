package hardware_wallet

import "github.com/chebyrash/promise"

type DeviceHelper interface {
	// FirstAddress return the first address in the sequence if there is a configured
	// device connected, error if not device found or some thing fail.
	FirstAddress(walletType string) *promise.Promise
}
