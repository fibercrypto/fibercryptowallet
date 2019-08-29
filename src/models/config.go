package models

import (
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"

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
	configManager *local.ConfigManager
	_             func()                          `constructor:"init"`
	_             func() []*WalletSource          `slot:"getSources"`
	_             func() string                   `slot:"getNodeString"`
	_             func() string                   `slot:"getSourceString"`
	_             func() int                      `slot:"getTypeSource` // 1 is Local, 0 is Remote
	_             func(node, src string, tp bool) `slot:"edit"`
}

func (cm *ConfigManager) init() {
	qml.QQmlEngine_SetObjectOwnership(cm, qml.QQmlEngine__CppOwnership)
	cm.configManager = local.GetConfigManager()
	cm.ConnectGetSources(cm.getSources)
	cm.ConnectGetNodeString(cm.getNodeString)
	cm.ConnectGetSourceString(cm.getSourceString)
	cm.ConnectGetTypeSource(cm.getTypeSource)
	cm.ConnectEdit(cm.edit)

}

func (cm *ConfigManager) edit(node, src string, tp bool) {
	var tpSrc int
	if tp {
		tpSrc = local.RemoteWallet
	} else {
		tpSrc = local.LocalWallet
	}
	cm.configManager.EditWalletSource(1, src, tpSrc)
	cm.configManager.EditNode(node)
	cm.configManager.Save()

}
func (cm *ConfigManager) getNodeString() string {
	return cm.configManager.GetNode()
}

func (cm *ConfigManager) getSourceString() string {
	src := cm.configManager.GetSources()[0]
	return src.GetSource()
}

func (cm *ConfigManager) getTypeSource() int {
	src := cm.configManager.GetSources()[0]
	if src.GetType() == local.LocalWallet {
		return 1
	} else {
		return 0
	}
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
