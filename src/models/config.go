package models

import (
	"encoding/json"

	local "github.com/fibercrypto/FiberCryptoWallet/src/main"

	"github.com/therecipe/qt/qml"

	qtcore "github.com/therecipe/qt/core"
)

type ConfigManager struct {
	qtcore.QObject
	configManager *local.ConfigManager
	_             func()                                 `constructor:"init"`
	_             func(string, []string, string, string) `slot:"setValue"`
	_             func(string, []string, string) string  `slot:"getValue"`
	_             func() []string                        `slot:"getSections"`
	_             func(string) *ConfigSection            `slot:"getSection"`
}

func (cm *ConfigManager) init() {
	qml.QQmlEngine_SetObjectOwnership(cm, qml.QQmlEngine__CppOwnership)
	cm.configManager = local.GetConfigManager()
	cm.ConnectSetValue(cm.setValue)
	cm.ConnectGetSections(cm.getSections)
	cm.ConnectGetSection(cm.getSection)

}

func (cm *ConfigManager) getSections() []string {
	return cm.configManager.GetSections()
}

func (cm *ConfigManager) getSection(sectionName string) *ConfigSection {
	sm := NewConfigSection(nil)
	sm.sm = cm.configManager.GetSectionManager(sectionName)
	return sm
}

type ConfigSection struct {
	qtcore.QObject
	sm *local.SectionManager
	_  func() [][]string                 `slot:"getPaths"`
	_  func([]string) []*KeyValueStorage `slot:getOptions"`
}

func (cs *ConfigSection) init() {
	qml.QQmlEngine_SetObjectOwnership(cs, qml.QQmlEngine__CppOwnership)
	cs.ConnectGetPaths(cs.getPaths)
	cs.ConnectGetOptions(cs.getOptions)
}

func (cs *ConfigSection) getPaths() [][]string {
	return cs.sm.GetPaths()
}

func (cs *ConfigSection) getOptions(path []string) []*KeyValueStorage {
	opts, err := cs.sm.GetValues(path)
	if err != nil {
		//log error
		return nil
	}

	resul := make([]*KeyValueStorage, 0)
	for _, opt := range opts {
		kv := NewKeyValueStorage(nil)
		data := make(map[string]string, 0)
		json.Unmarshal([]byte(opt), &data)
		kv.setValues(data)
		resul = append(resul, kv)
	}
	return resul
}

func (cm *ConfigManager) getValue(section string, sectionPath []string, name string) string {

	sm := cm.configManager.GetSectionManager(section)
	result, _ := sm.GetValue(name, sectionPath)
	return result

}

func (cm *ConfigManager) setValue(section string, sectionPath []string, name string, value string) {

	sm := cm.configManager.GetSectionManager(section)
	sm.Save(name, sectionPath, value)

}

type KeyValueStorage struct {
	qtcore.QObject
	_         func() []string     `slot:"getKeys"`
	_         func(string) string `slot:"getValue"`
	keyValues map[string]string
}

func (kv *KeyValueStorage) init() {
	kv.ConnectGetKeys(kv.getKeys)
	kv.ConnectGetValue(kv.getValue)
}

func (kv *KeyValueStorage) setValues(values map[string]string) {
	kv.keyValues = values
}

func (kv *KeyValueStorage) setValue(key, value string) {
	kv.keyValues[key] = value
}

func (kv *KeyValueStorage) getKeys() []string {
	keys := make([]string, 0)
	for key, _ := range kv.keyValues {
		keys = append(keys, key)
	}
	return keys
}

func (kv *KeyValueStorage) getValue(key string) string {
	val, ok := kv.keyValues[key]
	if !ok {
		//log error
		return ""
	}
	return val
}
