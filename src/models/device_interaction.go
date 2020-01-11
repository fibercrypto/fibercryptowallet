package models

import (
	hardware "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet/skywallet"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type QDeviceInteraction struct {
	core.QObject
	_ func()      `constructor:"init"`
	_ func()      `slot:"wipeDevice"`
	_ func()      `slot:"changePin"`
}

func (devI *QDeviceInteraction) init() {
	logWalletsModel.Info("Initialize Qt device interaction")
	qml.QQmlEngine_SetObjectOwnership(devI, qml.QQmlEngine__CppOwnership)
	devI.ConnectWipeDevice(devI.wipeDevice)
	devI.ConnectChangePin(devI.changePin)
}

func (devI *QDeviceInteraction) wipeDevice() {
	dev := hardware.NewSkyWalletInteraction()
	dev.Wipe().Then(func(data interface{}) interface{} {
		logWalletsModel.Infoln(data.(string))
		return data
	}).Catch(func(err error) error {
		logWalletsModel.WithError(err).Errorln("unable to wipe device")
		return err
	})
}

func (devI *QDeviceInteraction) changePin() {
	dev := hardware.NewSkyWalletInteraction()
	dev.Wipe().Then(func(data interface{}) interface{} {
		logWalletsModel.Infoln(data.(string))
		return data
	}).Catch(func(err error) error {
		logWalletsModel.WithError(err).Errorln("unable to change pin")
		return err
	})
}

