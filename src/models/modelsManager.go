package models

import (
	"time"

	"github.com/therecipe/qt/qml"

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
	go func() {
		uptimeTicker := time.NewTicker(time.Second * 2)

		for {
			<-uptimeTicker.C
			for wlt, _ := range mm.addressesModel {
				addrModel := NewAddressesModel(nil)
				qml.QQmlEngine_SetObjectOwnership(addrModel, qml.QQmlEngine__CppOwnership)
				addrModel.loadModel(mm.wltManager.getAddresses(wlt))
				addrModel.removeAddress(0)
				mm.addressesModel[wlt] = addrModel
			}

		}
	}()
}

func (mm *ModelManager) setWalletManager(wm *WalletManager) {
	mm.wltManager = *wm

}

func (mm *ModelManager) getAddressModel(wltName string) *AddressesModel {
	addrModel, ok := mm.addressesModel[wltName]
	if !ok {
		addrModel = NewAddressesModel(nil)
		qml.QQmlEngine_SetObjectOwnership(addrModel, qml.QQmlEngine__CppOwnership)
		addrModel.loadModel(mm.wltManager.getAddresses(wltName))
		addrModel.removeAddress(0)
		mm.addressesModel[wltName] = addrModel
	}

	return addrModel
}
