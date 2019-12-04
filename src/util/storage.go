package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

// KeyValueMap
type KeyValueMap struct {
	values map[string]interface{}
}

func (tOpt *KeyValueMap) GetValue(key string) interface{} {
	return tOpt.values[key]
}

func (tOpt *KeyValueMap) SetValue(key string, value interface{}) {
	tOpt.values[key] = value
}

func NewKeyValueMap() *KeyValueMap {
	tOptions := KeyValueMap{
		values: make(map[string]interface{}),
	}
	return &tOptions
}

// Type assertions
var (
	_ core.KeyValueStorage = &KeyValueMap{}
)
