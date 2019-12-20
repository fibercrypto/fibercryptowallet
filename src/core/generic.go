package core

type Iterator interface {
	CurrentData(interface{}) error
	HasNext() bool
	Next() bool
}
