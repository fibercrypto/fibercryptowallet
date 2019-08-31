package skycoin

import (
	"errors"
	"fmt"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
)

type SkycoinAddressIterator struct { //Implements AddressIterator interfaces
	current   int
	addresses []SkycoinAddress
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

func NewSkycoinAddressIterator(addresses []SkycoinAddress) *SkycoinAddressIterator {
	return &SkycoinAddressIterator{addresses: addresses, current: -1}
}

type SkycoinAddress struct { //Implements Address and CryptoAccount interfaces
	address string
	poolSection string
}

func (addr SkycoinAddress) IsBip32() bool {
	return false
}

func (addr SkycoinAddress) String() string {
	return addr.address
}

func (addr SkycoinAddress) GetCryptoAccount() core.CryptoAccount {
	return addr
}

func (addr SkycoinAddress) newClient() (*api.Client, error) {
	pool := core.GetMultiPool()
	conn, err := WaitForPooledObject(pool, "skycoin")
	if err != nil {
		return nil, err
	}

	c, ok := conn.(*api.Client)
	if !ok {
		return nil, errors.New(fmt.Sprintf("There is not propers client in %s pool", addr.poolSection))
	}
	return c, nil
}