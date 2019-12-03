package skycoin

import (
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

type SkycoinAddressIterator struct { //Implements AddressIterator interfaces
	current   int
	addresses []core.Address
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
	return (it.current + 1) < len(it.addresses)
}

func NewSkycoinAddressIterator(addresses []core.Address) *SkycoinAddressIterator {
	return &SkycoinAddressIterator{addresses: addresses, current: -1}
}

type SkycoinAddress struct { //Implements Address and CryptoAccount interfaces
	address     string
	poolSection string
}

func (addr *SkycoinAddress) IsBip32() bool {
	return false
}

func (addr *SkycoinAddress) String() string {
	return addr.address
}

func (addr *SkycoinAddress) GetCryptoAccount() core.CryptoAccount {
	return addr
}

func (addr SkycoinAddress) ToSkycoinCipherAddress() (*cipher.Address, error) {
	pubkey, err := cipher.PubKeyFromHex(addr.String())
	if err != nil {
		return nil, err
	}
	sAddr := cipher.AddressFromPubKey(pubkey)

	return &sAddr, nil
}
