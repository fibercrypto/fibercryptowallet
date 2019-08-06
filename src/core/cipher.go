package core

type Address interface {
	IsBip32() bool
	String() string
	GetCryptoAccount() CryptoAccount
}

type AddressIterator interface {
	Value() Address
	Next() bool
	HasNext() bool
}
