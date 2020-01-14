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
	_ func()                    `slot:"backupDevice"`
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
	devI.ConnectBackupDevice(devI.backupDevice)
}

func (devI *QDeviceInteraction) wipeDevice() {
	dev := hardware.NewSkyWalletInteraction()
	dev.Wipe().Then(func(data interface{}) interface{} {
		logWalletsModel.Infoln(data.(string))
		devI.OperationDone()
		return data
	}).Catch(func(err error) error {
		logWalletsModel.WithError(err).Errorln("unable to wipe device")
		devI.OperationDone()
		return err
	})
}

func (devI *QDeviceInteraction) backupDevice() {
	dev := hardware.NewSkyWalletInteraction()
	dev.Backup().Then(func(data interface{}) interface{} {
		devI.OperationDone()
		return data
	}).Catch(func(err error) error {
		devI.OperationDone()
		return err
	})
}

func (devI *QDeviceInteraction) changePin() {
	dev := hardware.NewSkyWalletInteraction()
	rm := false
	dev.ChangePin(&rm).Then(func(data interface{}) interface{} {
		logWalletsModel.Infoln(data.(string))
		devI.OperationDone()
		return data
	}).Catch(func(err error) error {
		logWalletsModel.WithError(err).Errorln("unable to change pin")
		devI.OperationDone()
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
		devI.OperationDone()
		return data
	}).Catch(func(err error) error {
		devI.OperationDone()
		return err
	})
}