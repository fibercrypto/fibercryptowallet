package core

type Address interface {
	IsBip32() bool
	String() string
}

type AddressIterator interface {
	Value() Address
	Next() bool
	HasNext() bool
}
