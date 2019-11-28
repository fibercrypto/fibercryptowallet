package core

// KeyValueStore provides read / write access to values given a key
type KeyValueStore interface {
	// GetValue lookup value for key
	GetValue(key string) interface{}
	// SetValue bind value o known key
	SetValue(key string, value interface{})
}
