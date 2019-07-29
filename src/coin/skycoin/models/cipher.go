package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

type SkycoinAddressIterator struct { //Implements AddressIterator interfaces
	current   int
	addresses []Address
}

func (it *SkycoinAddressIterator) Value() core.Address {
	return it.addresses[it.current]
}

func (it *SkycoinAddressIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *SkycoinAddressIterator) HasNext() bool {
	if (it.current + 1) >= len(it.addresses) {
		return false
	}
	return true
}

func NewSkycoinAddressIterator(addresses []Address) *SkycoinAddressIterator {
	return &SkycoinAddressIterator{addresses: addresses, current: -1}
}

type Address struct { //Implements Address and CryptoAccount interfaces
	address string
}

func (addr Address) IsBip32() bool {
	return true
}

func (addr Address) String() string {
	return addr.address
}

func (addr Address) GetCryptoAccount() core.CryptoAccount {
	return addr
}
