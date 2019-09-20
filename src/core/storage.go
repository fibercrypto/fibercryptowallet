package core

type KeyValueStorage interface {
	GetValue(key string) interface{}
}
