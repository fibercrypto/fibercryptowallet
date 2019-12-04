package parseutil

import (
	"errors"
	skycoin "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

// CoreAddressFromString returns a core.Address if match with string address.
// If the poolSection parameter not match with any address type returns 'poolSection not match' error.
func CoreAddressFromString(addrs, poolSection string) (core.Address, error) {
	switch poolSection {
	case "skycoin":
		return skycoin.NewSkycoinAddress(addrs)
	default:
		return nil, errors.New("poolSection not match")
	}
}
