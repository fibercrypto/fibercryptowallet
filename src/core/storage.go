package core

// KeyValueStorage provides read / write access to values given a key
type KeyValueStorage interface {
	GetValue(key string) interface{}
}
