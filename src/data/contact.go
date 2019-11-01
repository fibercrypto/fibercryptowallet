package data

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/data/internal"
	"github.com/gogo/protobuf/proto"
)

// Contact is a contact of the AddressBookWithBolt
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

//
func (c *Contact) GetID() uint64 {
	return c.id
}

//
func (c *Contact) SetID(id uint64) {
	c.id = id
}

//
func (c *Contact) GetAddress(pos int64) core.ReadableAddress {
	return &c.Address[pos]
}

//
func (c *Contact) SetAddress(addrs core.ReadableAddress) {
	if v, ok := addrs.(*Address); ok {
		c.Address = append(c.Address, *v)
	} else {
		panic("Error in SetAddress: addrs cannot parse to type data.Address")
	}
}

// Address
func (ad *Address) GetValue() []byte {
	return ad.Value
}

func (ad *Address) SetValue(val []byte) {
	ad.Value = val
}

func (ad *Address) GetCoinType() []byte {
	return ad.Coin
}

func (ad *Address) SetType(coinType []byte) {
	ad.Coin = coinType
}
