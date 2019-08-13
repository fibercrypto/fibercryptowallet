package models

import (
	"github.com/therecipe/qt/qml"
	//"github.com/fibercrypto/FiberCryptoWallet/src/models/history"

	qtcore "github.com/therecipe/qt/core"
)

type ModelManager struct {
	qtcore.QObject
	wltManager     WalletManager
	addressesModel map[string]*AddressesModel
	_              func()                       `constructor:"init"`
	_              func(*WalletManager)         `slot:"setWalletManager"`
	_              func(string) *AddressesModel `slot:"getAddressModel"`
}

func (mm *ModelManager) init() {
	mm.ConnectSetWalletManager(mm.setWalletManager)
	mm.ConnectGetAddressModel(mm.getAddressModel)
	qml.QQmlEngine_SetObjectOwnership(mm, qml.QQmlEngine__CppOwnership)
	mm.addressesModel = make(map[string]*AddressesModel, 0)
}

func (mm *ModelManager) setWalletManager(wm *WalletManager) {
	mm.wltManager = *wm

}

func (mm *ModelManager) getAddressModel(wltName string) *AddressesModel {
	addrModel, ok := mm.addressesModel[wltName]
	if !ok {

		addrModel = NewAddressesModel(nil)

		qml.QQmlEngine_SetObjectOwnership(addrModel, qml.QQmlEngine__CppOwnership)
		mm.wltManager.getNewSeed(128)
		addrModel.loadModel(mm.wltManager.getAddresses(wltName))
		addrModel.removeAddress(0)

		mm.addressesModel[wltName] = addrModel
	}

	return addrModel
}
