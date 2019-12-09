package models

import (
	"encoding/json"
	"strings"

	local "github.com/fibercrypto/fibercryptowallet/src/main"

	"github.com/therecipe/qt/qml"

	qtcore "github.com/therecipe/qt/core"
)

type ConfigManager struct {
	qtcore.QObject
	configManager *local.ConfigManager
	_             func()                      `constructor:"init"`
	_             func(string, string)        `slot:"setValue"`
	_             func(string) string         `slot:"getValue"`
	_             func() []string             `slot:"getSections"`
	_             func(string) *ConfigSection `slot:"getSection"`
	_             func(string) string         `slot:"getDefaultValue"`
}

func (cm *ConfigManager) init() {
	qml.QQmlEngine_SetObjectOwnership(cm, qml.QQmlEngine__CppOwnership)
	cm.configManager = local.GetConfigManager()
	cm.ConnectSetValue(cm.setValue)
	cm.ConnectGetSections(cm.getSections)
	cm.ConnectGetSection(cm.getSection)
	cm.ConnectGetValue(cm.getValue)
	cm.ConnectGetDefaultValue(cm.getDefaultValue)

}

func (cm *ConfigManager) getSections() []string {
	return cm.configManager.GetSections()
}

func (cm *ConfigManager) getSection(sectionName string) *ConfigSection {
	sm := NewConfigSection(nil)
	sm.sm = cm.configManager.GetSectionManager(sectionName)
	return sm
}

func (cm *ConfigManager) setValue(path, value string) {

	splitedPath := strings.Split(path, "/")
	section := splitedPath[0]
	optPath := splitedPath[1 : len(splitedPath)-1]
	name := splitedPath[len(splitedPath)-1]
	optName := optPath[len(optPath)-1]
	optPath = optPath[:len(optPath)-1]
	cm.GetSection(section).saveOptionValue(optName, optPath, name, value)

}

func (cm *ConfigManager) getValue(path string) string {
	splitedPath := strings.Split(path, "/")
	section := splitedPath[0]
	optPath := splitedPath[1 : len(splitedPath)-1]
	name := splitedPath[len(splitedPath)-1]
	optName := optPath[len(optPath)-1]
	optPath = optPath[:len(optPath)-1]
	return cm.GetSection(section).getValue(optName, optPath, name)
}

func (cm *ConfigManager) getDefaultValue(path string) string {
	splitedPath := strings.Split(path, "/")
	section := splitedPath[0]
	optPath := splitedPath[1 : len(splitedPath)-1]
	name := splitedPath[len(splitedPath)-1]
	optName := optPath[len(optPath)-1]
	optPath = optPath[:len(optPath)-1]
	return cm.GetSection(section).getDefaultValue(optName, optPath, name)
}

type ConfigSection struct {
	qtcore.QObject
	sm *local.SectionManager
	_  func() [][]string                 `slot:"getPaths"`
	_  func([]string) []*KeyValueStorage `slot:"getOptions"`
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

func (cs *ConfigSection) getOption(name string, path []string) *KeyValueStorage {
	opt, err := cs.sm.GetValue(name, path)
	if err != nil {
		//log error
		return nil
	}

	kv := NewKeyValueStorage(nil)
	data := make(map[string]string, 0)
	json.Unmarshal([]byte(opt), &data)
	kv.setValues(data)
	return kv
}

func (cs *ConfigSection) saveOptionValue(opt string, path []string, name string, value string) {
	optV := cs.getOption(opt, path)
	if optV == nil {
		//log error
		return
	}

	optV.setValue(name, value)
	data, err := json.Marshal(optV.keyValues)
	if err != nil {
		//log error
		return
	}
	cs.sm.Save(opt, path, string(data))

}

func (cs *ConfigSection) getValue(opt string, path []string, name string) string {
	optV := cs.getOption(opt, path)
	if optV == nil {
		//log error
		return ""
	}
	return optV.getValue(name)

}

func (cs *ConfigSection) getDefaultValue(opt string, path []string, name string) string {
	val, err := cs.sm.GetDefaultValue(opt, path, name)
	if err != nil {
		return ""
	}
	return val
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
