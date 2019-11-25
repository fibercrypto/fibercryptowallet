package core

// KeyValueStorage provides read / write access to values given a key
type KeyValueStorage interface {
	// GetValue lookup value for key
	GetValue(key string) interface{}
	// SetValue bind value o known key
	SetValue(key string, value interface{})
}
