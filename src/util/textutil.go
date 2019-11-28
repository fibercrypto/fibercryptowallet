package util

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

// EmptyPassword read no password
func EmptyPassword(string, core.KeyValueStore) (string, error) {
	return "", nil
}
