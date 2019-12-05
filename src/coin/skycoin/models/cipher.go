package skycoin

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/skycoin/skycoin/src/cipher"
)

type SkycoinAddressIterator struct { // Implements AddressIterator interfaces
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

func NewSkycoinAddress(addrStr string, isBip32 bool) (core.Address, error) {
	var skyAddr cipher.Address
	var err error
	if skyAddr, err = cipher.DecodeBase58Address(addrStr); err != nil {
		return nil, err
	}

	return &SkycoinAddress{
		isBip32:     isBip32,
		address:     skyAddr,
		poolSection: PoolSection,
	}, nil
}

type SkycoinAddress struct { // Implements Address and CryptoAccount interfaces
	isBip32     bool
	address     cipher.Address
	poolSection string
}

func (addr *SkycoinAddress) IsBip32() bool {
	return addr.isBip32
}

func (addr *SkycoinAddress) String() string {
	return addr.address.String()
}

func (addr *SkycoinAddress) GetCryptoAccount() core.CryptoAccount {
	return addr
}

func (addr *SkycoinAddress) ToSkycoinCipherAddress() (cipher.Address, error) {
	return cipher.AddressFromBytes(addr.Bytes())
}

// Bytes binary representation for address
func (addr *SkycoinAddress) Bytes() []byte {
	return addr.address.Bytes()
}

// Checksum computes address consistency token
func (addr *SkycoinAddress) Checksum() core.Checksum {
	checksum := addr.address.Checksum()
	return checksum[:]
}

func toSkycoinPubKey(pk core.PubKey) (cipher.PubKey, error) {
	if pk1, isSkyPK := pk.(*SkycoinPubKey); isSkyPK {
		return pk1.pubkey, nil
	}
	return cipher.NewPubKey(pk.Bytes())
}

func toSkycoinSecKey(sk core.SecKey) (cipher.SecKey, error) {
	if sk1, isSkySK := sk.(*SkycoinSecKey); isSkySK {
		return sk1.seckey, nil
	}
	return cipher.NewSecKey(sk.Bytes())
}

// Verify checks that the address appears valid for the public key
func (addr *SkycoinAddress) Verify(pk core.PubKey) error {

	skyPK, err := toSkycoinPubKey(pk)
	if err != nil {
		return err
	}
	return addr.address.Verify(skyPK)
}

// Null returns true if the address is null
func (addr *SkycoinAddress) Null() bool {
	return addr.address.Null()
}

// SkycoinSecKey Skycoin private key wrapper
type SkycoinSecKey struct {
	seckey cipher.SecKey
}

// Verify checks that the private key appears valid
func (sk *SkycoinSecKey) Verify() error {
	return sk.seckey.Verify()
}

// Null returns true if the private key is null
func (sk *SkycoinSecKey) Null() bool {
	return sk.seckey.Null()
}

// Bytes binary representation for private key
func (sk *SkycoinSecKey) Bytes() []byte {
	return sk.seckey[:]
}

func skySecKeyFromBytes(b []byte) (*SkycoinSecKey, error) {
	sk, err := cipher.NewSecKey(b)
	if err != nil {
		return nil, err
	}
	return &SkycoinSecKey{
		seckey: sk,
	}, nil
}

func skyPubKeyFromBytes(b []byte) (*SkycoinPubKey, error) {
	pk, err := cipher.NewPubKey(b)
	if err != nil {
		return nil, err
	}
	return &SkycoinPubKey{
		pubkey: pk,
	}, nil
}

// SkycoinPubKey Skycoin public key wrapper
type SkycoinPubKey struct {
	pubkey cipher.PubKey
}

// Verify checks that the public key appears valid
func (pk *SkycoinPubKey) Verify() error {
	return pk.pubkey.Verify()
}

// Null returns true if the public key is null
func (pk *SkycoinPubKey) Null() bool {
	return pk.pubkey.Null()
}

// Bytes binary representation for public key
func (pk *SkycoinPubKey) Bytes() []byte {
	return pk.pubkey[:]
}

// Type assertions
var (
	_ core.Address = &SkycoinAddress{}
	_ core.PubKey  = &SkycoinPubKey{}
	_ core.SecKey  = &SkycoinSecKey{}
)
