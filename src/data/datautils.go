package data

import (
	"crypto/sha512"
	"golang.org/x/crypto/pbkdf2"
)

func derivePassphrase(entropy, password []byte) []byte {
	return pbkdf2.Key(entropy, []byte("entropy:"+string(password)), 4096, 32, sha512.New)
}
