package parseutil

import (
	"errors"
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

// AddressFromString returns a core.Address if match with string address.
// If the coinTicket parameter not match with any address type returns 'coinTicket not match' error.
func AddressFromString(addrs, coinTicket string) (core.Address, error) {
	switch coinTicket {
	case "skycoin":
		return skycoin.NewSkycoinAddress(addrs, false)
	default:
		return nil, errors.New("coinTicket not match")
	}
}
