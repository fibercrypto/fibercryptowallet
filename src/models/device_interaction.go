package models

import (
	"crypto/sha256"
	hardware "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"io/ioutil"
	"strings"
)

type QDeviceInteraction struct {
	core.QObject
	_ func()                    `constructor:"init"`
	_ func()                    `slot:"wipeDevice"`
	_ func()                    `slot:"backupDevice"`
	_ func()                    `slot:"changePin"`
	_ func()                    `slot:"deviceFeatures"`
	_ func(uint, bool)          `slot:"generateMnemonic"`
	_ func(uint, bool)          `slot:"restoreBackup"`
	_ func()                    `slot:"cancelCommand"`
	_ func(file string)         `slot:"firmwareUpload"`
	_ func(hasPin bool)         `signal:"hasPinDetermined"`
	_ func(name string)         `signal:"nameDetermined"`
	_ func(isInitialized bool)  `signal:"isInitializedDetermined"`
	_ func(needsBackup bool)    `signal:"needsBackupDetermined"`
	_ func(bootloader bool)     `signal:"bootModeDetermined"`
	_ func()                    `signal:"operationDone"`
	_ func()                    `signal:"secureDevice"`
	_ func()                    `signal:"initializeDevice"`
}

func (devI *QDeviceInteraction) init() {
	logWalletsModel.Info("Initialize Qt device interaction")
	qml.QQmlEngine_SetObjectOwnership(devI, qml.QQmlEngine__CppOwnership)
	devI.ConnectWipeDevice(devI.wipeDevice)
	devI.ConnectChangePin(devI.changePin)
	devI.ConnectDeviceFeatures(devI.deviceFeatures)
	devI.ConnectBackupDevice(devI.backupDevice)
	devI.ConnectCancelCommand(devI.cancelCommand)
	devI.ConnectFirmwareUpload(devI.firmwareUpload)
	devI.ConnectGenerateMnemonic(devI.generateMnemonic)
	devI.ConnectRestoreBackup(devI.restoreBackup)
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
		if err == skywallet.ErrUserCancelledFromInputReader {
			logWalletsModel.WithError(err).Infoln("unable to change pin")
		} else if err == skywallet.ErrUserCancelledWithDeviceButton {
			logWalletsModel.WithError(err).Warningln("unable to change pin")
		} else {
			logWalletsModel.WithError(err).Errorln("unable to change pin")
		}
		devI.OperationDone()
		return err
	})
}

func (devI* QDeviceInteraction) cancelCommand() {
	dev := hardware.NewSkyWalletInteraction()
	dev.Cancel().Then(func(data interface{}) interface{} {
		logWalletsModel.Infoln(data.(string))
		return data
	}).Catch(func(err error) error {
		logWalletsModel.WithError(err).Errorln("unable to cancel operation")
		return err
	})
}

func (devI *QDeviceInteraction) deviceFeatures() {
	dev := hardware.NewSkyWalletInteraction()
	dev.Features().Then(func(data interface{}) interface{} {
		features := data.(messages.Features)
		if features.Label != nil {
			devI.NameDetermined(*features.Label)
		}
		devI.HasPinDetermined(features.PinProtection != nil && *features.PinProtection)
		devI.NeedsBackupDetermined(features.NeedsBackup != nil && *features.NeedsBackup)
		devI.IsInitializedDetermined(features.Initialized != nil && *features.Initialized)
		devI.BootModeDetermined(features.BootloaderMode != nil && *features.BootloaderMode)
		return data
	}).Catch(func(err error) error {
		return err
	})
}

func (devI *QDeviceInteraction) generateMnemonic(wordCount uint, usePassphrase bool) {
	dev := hardware.NewSkyWalletInteraction()
	dev.GenerateMnemonic(uint32(wordCount), usePassphrase).Then(func(data interface{}) interface{} {
		logWalletsModel.Infoln(data)
		devI.OperationDone()
		return data
	}).Catch(func(err error) error {
		devI.OperationDone()
		logWalletsModel.WithError(err).Errorln("unable to generate mnemonic")
		return err
	})
}

func (devI *QDeviceInteraction) restoreBackup(wordCount uint, usePassphrase bool) {
	dev := hardware.NewSkyWalletInteraction()
	dev.Recovery(uint32(wordCount), &usePassphrase, false).Then(func(data interface{}) interface{} {
		logWalletsModel.Infoln(data)
		devI.OperationDone()
		return data
	}).Catch(func(err error) error {
		logWalletsModel.WithError(err).Errorln("unable to recover device with you seed")
		devI.OperationDone()
		return err
	})
}

func (devI *QDeviceInteraction) firmwareUpload(filePath string) {
	// FIXME portability
	prefix := "file://"
	if strings.HasPrefix(filePath, prefix) {
		filePath = strings.TrimPrefix(filePath, prefix)
	}
	firmware, err := ioutil.ReadFile(filePath)
	if err != nil {
		// FIXME inform the end user
		logWalletsModel.WithError(err).Errorln("unable to upload firmware to device")
		return
	}
	dev := hardware.NewSkyWalletInteraction()
	dev.UploadFirmware(firmware, sha256.Sum256(firmware[0x100:])).Then(func(data interface{}) interface{} {
		devI.OperationDone()
		return data
	}).Catch(func(err error) error {
		devI.OperationDone()
		return err
	})
}
