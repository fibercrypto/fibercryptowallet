package data

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/data/internal"
	"github.com/gogo/protobuf/proto"
)

// Contact is a contact of the DB
type Contact struct {
	id      uint64
	Address []Address
	Name    []byte
}

// Address is the relation of an address and his coin type.
type Address struct {
	Value []byte
	Coin  []byte
}

// MarshalBinary encodes a user to binary format.
func (c *Contact) MarshalBinary() ([]byte, error) {
	var intAddrs []internal.Address
	for _, v := range c.Address {
		intAddrs = append(intAddrs, internal.Address{
			Address:  v.GetValue(),
			CoinType: v.GetCoinType(),
		})
	}
	return proto.Marshal(&internal.Contact{
		ID:      c.id,
		Name:    c.Name,
		Address: intAddrs,
	})
}

// UnmarshalBinary decodes a user from binary data.
func (c *Contact) UnmarshalBinary(data []byte) error {
	var pb internal.Contact
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}
	c.id = pb.ID
	c.Name = pb.Name
	for _, v := range pb.Address {
		c.Address = append(c.Address, Address{
			Value: v.Address,
			Coin:  v.CoinType,
		})
	}
	return nil
}

// GetID get id of current contact.
func (c *Contact) GetID() uint64 {
	return c.id
}

// SetID set an id to current contact.
func (c *Contact) SetID(id uint64) {
	c.id = id
}

// GetAddresses get address list of current contact.
func (c *Contact) GetAddresses() []core.ReadableAddress {
	var addresses []core.ReadableAddress
	for e := range c.Address {
		addresses = append(addresses, &c.Address[e])
	}
	return addresses
}

// SetAddresses set an address list to the current contact.
func (c *Contact) SetAddresses(addrs []core.ReadableAddress) {
	for e := range addrs {
		if v, ok := addrs[e].(*Address); ok {
			c.Address = append(c.Address, *v)
		} else {
			panic("Error in SetAddress: addrs cannot parse to type data.Address")

		}
	}
}

// .....Address

// GetValue get address string.
func (ad *Address) GetValue() []byte {
	return ad.Value
}

// SetValue set an address string.
func (ad *Address) SetValue(val []byte) {
	ad.Value = val
}

// GetCoinType get coin type of an address.
func (ad *Address) GetCoinType() []byte {
	return ad.Coin
}

// SetCoinType set the coin type to current address.
func (ad *Address) SetCoinType(coinType []byte) {
	ad.Coin = coinType
}
