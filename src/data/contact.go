package data

import (
	"encoding/base64"
	"encoding/json"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"strings"
)

// Contact is a contact of the addrsBook
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

// Encrypt encrypt a contact by its security Type. Only use key parameter with secType==PasswordSecurity.
func (contact *Contact) Encrypt(secType int, key []byte) ([]byte, error) {
	data, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}
	switch secType {
	case NoSecurity:
		return data, nil
	case ObfuscationSecurity:
		return []byte(base64.StdEncoding.EncodeToString(data)), nil
	case PasswordSecurity:
		return encryptAESGCM(data, key)
	default:
		return nil, errInvalidSecType
	}
}

// Decrypt decrypt a contact by its security type. Only use key parameter with secType==PasswordSecurity.
func (contact *Contact) Decrypt(secType int, data, key []byte) error {
	var err error
	switch secType {
	case NoSecurity:
		break
	case ObfuscationSecurity:
		data, err = base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			return err
		}
	case PasswordSecurity:
		data, err = decryptAESGCM(data, key)
	default:
		return errInvalidSecType
	}
	return json.Unmarshal(data, contact)
}

// GetID get ID of current contact.
func (contact *Contact) GetID() uint64 {
	return contact.ID
}

// SetID set an ID to current contact.
func (contact *Contact) SetID(id uint64) {
	contact.ID = id
}

// GetAddresses get address list of current contact.
func (contact *Contact) GetAddresses() []core.StringAddress {
	var addresses []core.StringAddress
	for e := range contact.Address {
		addresses = append(addresses, &contact.Address[e])
	}
	return addresses
}

// SetAddresses set an address list to the current contact.
func (contact *Contact) SetAddresses(addrs []core.StringAddress) {
	for e := range addrs {
		if v, ok := addrs[e].(*Address); ok {
			contact.Address = append(contact.Address, *v)
		} else {
			panic("Error in SetAddress: addrs cannot parse to type data.Address")

		}
	}
}

// GetName return contact name
func (contact *Contact) GetName() string {
	return string(contact.Name)
}

// SetName set a name to the contact
func (contact *Contact) SetName(newName string) {
	contact.Name = []byte(newName)
}

// IsValid verify if a contact is valid and all its address.
func (contact *Contact) IsValid() bool {
	if strings.ReplaceAll(contact.GetName(), " ", "") == "" {
		return false
	}

	for e := range contact.GetAddresses() {
		if !contact.GetAddresses()[e].IsValid() {
			return false
		}
	}
	return true
}

// .....Address

// GetValue get address string.
func (address *Address) GetValue() []byte {
	return address.Value
}

// SetValue set an address string.
func (address *Address) SetValue(val []byte) {
	address.Value = val
}

// GetCoinType get coin type of an address.
func (address *Address) GetCoinType() []byte {
	return address.Coin
}

// SetCoinType set the coin type to current address.
func (address *Address) SetCoinType(coinType []byte) {
	address.Coin = coinType
}

// IsValid verify if an address is valid for its coinType.
func (address *Address) IsValid() bool {
	if _, err := util.AddressFromString(string(address.GetValue()),
		string(address.GetCoinType())); err != nil {
		return false
	}
	return true
}
