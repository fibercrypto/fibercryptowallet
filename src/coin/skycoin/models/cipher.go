package skycoin

import (
	"github.com/FiberCrypto/FiberCryptoWallet/src/core"
)

type SkycoinAddress struct {
}

func (*SkycoinAddress) IsBip32() bool {
	return false
}

func (*SkycoinAddress) String() string {
	return "249iEtdhniFppBTpkzq7syoiBaydLi6USnU"
}

type SkycoinAddressIterator struct {
	index int
	value SkycoinAddress
}

func (iter *SkycoinAddressIterator) Value() core.Address {
	if index == 0 && iter.value == nil {
		iter.value = &SkycoinAddress{}
	}
	return iter.value
}

func (iter *SkycoinAddressIterator) Next() bool {
	if iter.index < 3 {
		iter.value = &SkycoinAddress{}
		iter.index++
		return true
	}
	return false
}

func (iter *SkycoinAddressIterator) HasNext() bool {
	return iter.index < 3
}
