package models

func init() {
	BlockchainStatusModel_QmlRegisterType2("BlockchainModels", 1, 0, "BlockchainStatusModel")
	WalletModel_QmlRegisterType2("WalletsManager", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("WalletsManager", 1, 0, "QWallet")
	AddressesModel_QmlRegisterType2("WalletsManager", 1, 0, "AddressModel")
	QAddress_QmlRegisterType2("WalletsManager", 1, 0, "QAddress")
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")
	ConfigManager_QmlRegisterType2("Config", 1, 0, "ConfigManager")
	WalletSource_QmlRegisterType2("Config", 1, 0, "WalletSource")
	ModelManager_QmlRegisterType2("WalletsManager", 1, 0, "ModelManager")
	ModelWallets_QmlRegisterType2("OutputsModels", 1, 0, "QWallets")
	ModelAddresses_QmlRegisterType2("OutputsModels", 1, 0, "QAddresses")
	ModelOutputs_QmlRegisterType2("OutputsModels", 1, 0, "QOutputs")
}
