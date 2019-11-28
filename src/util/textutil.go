package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

// EmptyPassword read no password
func EmptyPassword(string, core.KeyValueStore) (string, error) {
	return "", nil
}
