package core

type KeyValueStorage interface {
	getValue(key string) interface{}
}
