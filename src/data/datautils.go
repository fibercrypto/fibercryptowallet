package data

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

func derivePassphrase(entropy, password []byte) []byte {
	return pbkdf2.Key(entropy, password, 4096, 32, sha512.New)
}

// Itob converts uint64 to bytes
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// Btoi converts bytes to uint64
func btoi(v []byte) uint64 {
	return binary.BigEndian.Uint64(v)
}

func toBase64(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}

func fromBase64(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func errorWrapper(err error) error {
	return fmt.Errorf("address book error: %s", err)
}
