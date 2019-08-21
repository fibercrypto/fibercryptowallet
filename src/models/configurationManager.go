package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"

	"github.com/therecipe/qt/qml"

	qtcore "github.com/therecipe/qt/core"
)

type WalletSource struct {
	qtcore.QObject
	_ int    `property:"sourceType"`
	_ string `property:"source"`
}

type ConfigManager struct {
	qtcore.QObject
	configManager *core.ConfigManager
	_             func()                 `constructor:"init"`
	_             func() []*WalletSource `slot:"getSources"`
}

func (cm *ConfigManager) init() {
	qml.QQmlEngine_SetObjectOwnership(cm, qml.QQmlEngine__CppOwnership)
	cm.configManager = core.GetConfigManager()
	cm.ConnectGetSources(cm.getSources)

}

func (cm *ConfigManager) getSources() []*WalletSource {
	wltsSource := make([]*WalletSource, 0)
	for _, wltS := range cm.configManager.GetSources() {
		newWltSource := NewWalletSource(nil)
		qml.QQmlEngine_SetObjectOwnership(newWltSource, qml.QQmlEngine__CppOwnership)
		newWltSource.SetSourceType(wltS.GetType())
		newWltSource.SetSource(wltS.GetSource())
		wltsSource = append(wltsSource, newWltSource)
	}
	return wltsSource
}
