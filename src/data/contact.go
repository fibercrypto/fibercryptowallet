package data

import (
	"./internal"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

// Contact is a contact of the AddressBook
type Contact struct {
	ID      uint64
	Address []Address
	Name    []byte
}

// Address is the relation of an address and his coin type.
type Address struct {
	Value []byte
	Coin  []byte
}

func (c *Contact) encryptContact(password, mnemonic []byte) ([]byte, error) {
	if mnemonic == nil {
		return nil, fmt.Errorf(" Error: Mnemonic are empty.")
	}

	block, err := aes.NewCipher(derivePassphrase(mnemonic, password))
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
	bc, err := c.MarshalBinary()
	if err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, bc, nil)
	return ciphertext, nil
}

func (c *Contact) decryptContact(ciphermsg, password, mnemonic []byte) error {
	if mnemonic == nil {
		return fmt.Errorf(" Error: Mnemonic are empty.")
	}

	block, err := aes.NewCipher(derivePassphrase(mnemonic, password))
	if err != nil {
		return err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphermsg[:nonceSize], ciphermsg[nonceSize:]

	data, err := aesGCM.Open(nil, nonce, ciphertext, nil)

	if err := c.UnmarshalBinary(data); err != nil {
		return err
	}
	return nil
}

// MarshalBinary encodes a user to binary format.
func (c *Contact) MarshalBinary() ([]byte, error) {
	var intrAddrs []internal.Address
	for _, v := range c.Address {
		intrAddrs = append(intrAddrs, internal.Address{
			Address:  v.Value,
			CoinType: v.Coin,
		})
	}
	return proto.Marshal(&internal.Contact{
		ID:      c.ID,
		Name:    c.Name,
		Address: intrAddrs,
	})
}

// UnmarshalBinary decodes a user from binary data.
func (c *Contact) UnmarshalBinary(data []byte) error {
	var pb internal.Contact
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}
	c.ID = pb.ID
	c.Name = pb.Name
	for _, v := range pb.Address {
		c.Address = append(c.Address, Address{
			Value: v.Address,
			Coin:  v.CoinType,
		})
	}
	return nil
}

//
func derivePassphrase(mnemonic, password []byte) []byte {
	return pbkdf2.Key(mnemonic, []byte("mnemonic"+string(password)), 4096, 32, sha512.New)
}
