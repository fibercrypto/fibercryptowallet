package models

import (
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"

	"github.com/therecipe/qt/qml"

	qtcore "github.com/therecipe/qt/core"
)

var logConfigManager = logging.MustGetLogger("modelsConfigManager")

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

func (configManager *ConfigManager) init() {
	logConfigManager.Info("Initialize config manager")
	qml.QQmlEngine_SetObjectOwnership(configManager, qml.QQmlEngine__CppOwnership)
	configManager.configManager = local.GetConfigManager()
	configManager.ConnectGetSources(configManager.getSources)
	configManager.ConnectGetNodeString(configManager.getNodeString)
	configManager.ConnectGetSourceString(configManager.getSourceString)
	configManager.ConnectGetTypeSource(configManager.getTypeSource)
	configManager.ConnectEdit(configManager.edit)

}

func (configManager *ConfigManager) edit(node, src string, tp bool) {
	logConfigManager.Info("Editing Config Manager")
	var tpSrc int
	if tp {
		tpSrc = local.RemoteWallet
	} else {
		tpSrc = local.LocalWallet
	}
	configManager.configManager.EditWalletSource(1, src, tpSrc)
	configManager.configManager.EditNode(node)
	configManager.configManager.Save()

}
func (configManager *ConfigManager) getNodeString() string {
	return configManager.configManager.GetNode()
}

func (configManager *ConfigManager) getSourceString() string {
	src := configManager.configManager.GetSources()[0]
	return src.GetSource()
}

func (configManager *ConfigManager) getTypeSource() int {
	src := configManager.configManager.GetSources()[0]
	if src.GetType() == local.LocalWallet {
		return 1
	} else {
		return 0
	}
}

func (configManager *ConfigManager) getSources() []*WalletSource {
	logConfigManager.Info("Getting wallets source")
	walletSources := make([]*WalletSource, 0)
	for _, wltS := range configManager.configManager.GetSources() {
		newWltSource := NewWalletSource(nil)
		qml.QQmlEngine_SetObjectOwnership(newWltSource, qml.QQmlEngine__CppOwnership)
		newWltSource.SetSourceType(wltS.GetType())
		newWltSource.SetSource(wltS.GetSource())
		walletSources = append(walletSources, newWltSource)
	}
	return walletSources
}
