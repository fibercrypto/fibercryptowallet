package models

import (
	hardware "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type QDeviceInteraction struct {
	core.QObject
	_ func()                    `constructor:"init"`
	_ func()                    `slot:"wipeDevice"`
	_ func()                    `slot:"changePin"`
	_ func()                    `slot:"deviceFeatures"`
	_ func(hasPin bool)         `signal:"hasPinDetermined"`
	_ func(name string)         `signal:"nameDetermined"`
	_ func(needsBackup bool)    `signal:"needsBackupDetermined"`
	_ func()                    `signal:"operationDone"`
}

func (devI *QDeviceInteraction) init() {
	logWalletsModel.Info("Initialize Qt device interaction")
	qml.QQmlEngine_SetObjectOwnership(devI, qml.QQmlEngine__CppOwnership)
	devI.ConnectWipeDevice(devI.wipeDevice)
	devI.ConnectChangePin(devI.changePin)
	devI.ConnectDeviceFeatures(devI.deviceFeatures)
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

func (devI *QDeviceInteraction) deviceFeatures() {
	dev := hardware.NewSkyWalletInteraction()
	dev.Features().Then(func(data interface{}) interface{} {
		features := data.(messages.Features)
		devI.HasPinDetermined(*features.PinProtection)
		devI.NameDetermined(*features.Label)
		devI.NeedsBackupDetermined(*features.NeedsBackup)
		return data
	}).Catch(func(err error) error {
		return err
	})
}