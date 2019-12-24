package data

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/SkycoinProject/skycoin/src/cipher/bip39"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"io"
)

// isRepeatAddress search an address in the list of contacts into the AddressBook.
// If find the address return error, else return nil.
func isRepeatAddress(address core.StringAddress, contacts []core.Contact) error {
	for _, contact := range contacts {
		for _, addrs := range contact.GetAddresses() {
			if bytes.Compare(addrs.GetValue(), address.GetValue()) == 0 &&
				bytes.Compare(addrs.GetCoinType(), address.GetCoinType()) == 0 {
				return fmt.Errorf("Address with value: %s  and Cointype: %s alredy exist",
					address.GetValue(), address.GetCoinType())
			}
		}
	}
	return nil
}

// isRepeatName search an name in the list of contacts into the AddressBook.
// If find the address return error, else return nil.
func isRepeatName(newContact core.Contact, contactsList []core.Contact) error {
	for _, contact := range contactsList {
		if dataContact, ok := contact.(*Contact); ok {
			if bytes.Compare(newContact.(*Contact).Name, dataContact.Name) == 0 {
				return fmt.Errorf(" Contact with name: %s alredy exist", newContact.(*Contact).Name)
			}
		}
	}

	return nil
}

// genEntropy generate an entropy by a random mnemonic.
func genEntropy() ([]byte, error) {
	mn, err := bip39.NewDefaultMnemonic()
	if err != nil {
		return nil, err
	}
	e, err := bip39.EntropyFromMnemonic(mn)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// decryptAESGCM encrypt a data with AESGCM algorithm. http://golang.org/crypto/
func encryptAESGCM(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := aesGCM.Seal(nonce, nonce, data, nil)
	return cipherText, nil

}

// decryptAESGCM decrypt a data with AESGCM algorithm. http://golang.org/crypto/
func decryptAESGCM(cipherData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := cipherData[:nonceSize], cipherData[nonceSize:]

	return aesGCM.Open(nil, nonce, cipherText, nil)
}

func errValueNoMatch(value, valType interface{}) error {
	return fmt.Errorf("value %v type %T no match with type %T", value, value, valType)
}
