package core

// Iterator defines a generic iterator.
type Iterator interface {
	CurrentData(interface{}) error
	HasNext() bool
	Next() bool
}
