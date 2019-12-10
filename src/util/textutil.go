package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

// EmptyPassword read no password
func EmptyPassword(string, core.KeyValueStore) (string, error) {
	return "", nil
}

// ConstantPassword always return same known password
func ConstantPassword(pwdText string) core.PasswordReader {
	return func(string, core.KeyValueStore) (string, error) {
		return pwdText, nil
	}
}
