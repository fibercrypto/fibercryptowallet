package util

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

// KeyValueMap
type KeyValueMap struct {
	values map[string]interface{}
}

// GetValue lookup value for key
func (tOpt *KeyValueMap) GetValue(key string) interface{} {
	return tOpt.values[key]
}

// SetValue bind value o known key
func (tOpt *KeyValueMap) SetValue(key string, value interface{}) {
	tOpt.values[key] = value
}

// NewKeyValueMap instantiate key value map storage
func NewKeyValueMap() *KeyValueMap {
	tOptions := KeyValueMap{
		values: make(map[string]interface{}),
	}
	return &tOptions
}

// NewKeyValueMap instantiate key value map storage
func NewMapWithSingleKey(k string, v interface{}) *KeyValueMap {
	store := NewKeyValueMap()
	store.values[k] = v
	return store
}

// KeyValuesWithDefaults retrieve and set values of another KeyValueStore with defaults for keys not found
type KeyValuesWithDefaults struct {
	values, defaults core.KeyValueStore
}

// GetValue lookup value for key
func (tOpt *KeyValuesWithDefaults) GetValue(key string) interface{} {
	v := tOpt.values.GetValue(key)
	if v != nil {
		return v
	}
	return tOpt.defaults.GetValue(key)
}

// SetValue bind value o known key
func (tOpt *KeyValuesWithDefaults) SetValue(key string, value interface{}) {
	tOpt.values.SetValue(key, value)
}

// NewKeyValuesWithDefaults instantiate key value map storage
func NewKeyValuesWithDefaults(values, defaults core.KeyValueStore) core.KeyValueStore {
	if defaults == nil {
		return values
	}
	if values == nil {
		values = NewKeyValueMap()
	}
	tOptions := KeyValuesWithDefaults{
		values:   values,
		defaults: defaults,
	}
	return &tOptions
}

// Type assertions
var (
	_ core.KeyValueStore = &KeyValueMap{}
)
