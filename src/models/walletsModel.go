package models

import (
	"errors"
	hardware "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet/skywallet"
	fce "github.com/fibercrypto/fibercryptowallet/src/errors"
	"strconv"

	coin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	fccore "github.com/fibercrypto/fibercryptowallet/src/core"
	wlcore "github.com/fibercrypto/fibercryptowallet/src/main"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"time"
)

const (
	Name = int(core.Qt__UserRole) + iota + 1
	EncryptionEnabled
	Sky
	CoinHours
	FileName
	Expand
	HasHardwareWallet
)

var logWalletsModel = logging.MustGetLogger("Wallets Model")
var hadHwConnected = false
var hwConnectedOn []int

type WalletModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QWallet               `property:"wallets"`

	_ func(*QWallet)                                                                   `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, sky string, coinHours string) `slot:"editWallet"`
	_ func(row int)                                                                    `slot:"removeWallet"`
	_ func([]*QWallet)                                                                 `slot:"loadModel"`
	_ func([]*QWallet)                                                                 `slot:"updateModel"`
	_ func(devI *QDeviceInteraction, locker *QBridge)                                  `slot:"sniffHw"`
	_ int                                                                              `property:"count"`
}

type QWallet struct {
	core.QObject
	_ string `property:"name"`
	_ int    `property:"encryptionEnabled"`
	_ string `property:"sky"`
	_ string `property:"coinHours"`
	_ string `property:"fileName"`
	_ bool   `property:"expand"`
	_ bool   `property:"hasHardwareWallet"`
}

func createSkyHardwareWallet(bridgeForPassword *QBridge) {
	requestKind2Prompt := func(kind skyWallet.InputRequestKind, bridge *QBridge) (func(string, string), error) {
		switch kind {
		case skyWallet.RequestKindWord:
			return bridge.GetBip39Word, nil
		case skyWallet.RequestKindPinMatrix:
			return bridge.GetSkyHardwareWalletPin, nil
		case skyWallet.RequestInformUserOkAndCancel:
			return bridge.DeviceRequireConfirmableAction, nil
		case skyWallet.RequestInformUserOnlyCancel:
			return bridge.DeviceRequireCancelableAction, nil
		case skyWallet.RequestInformUserOnlyOk:
			return bridge.DeviceRequireAction, nil
		default:
			errStr := "invalid request kind"
			logWalletsModel.WithField("kind", kind).Errorln(errStr)
			return nil, errors.New(errStr)
		}
	}
	hardware.CreateSkyWltInteractionInstanceOnce(
		skyWallet.DeviceTypeEmulator,
		func(kind skyWallet.InputRequestKind, tittle, message string)(string, error) {
			prompt, err := requestKind2Prompt(kind, bridgeForPassword)
			if err != nil {
				return "", err
			}
			bridgeForPassword.BeginUse()
			defer bridgeForPassword.EndUse()
			switch kind {
			case skyWallet.RequestKindPinMatrix, skyWallet.RequestKindWord:
				bridgeForPassword.lock()
				prompt(tittle, message)
				bridgeForPassword.lock()
				pass, err := bridgeForPassword.getOptionalResult()
				bridgeForPassword.unlock()
				if err != nil {
					logWalletsModel.WithError(err).Warningln("error handling user interaction")
					return "", skyWallet.ErrUserCancelledFromInputReader
				}
				return pass, nil
			case skyWallet.RequestInformUserOkAndCancel, skyWallet.RequestInformUserOnlyCancel, skyWallet.RequestInformUserOnlyOk:
				prompt(tittle, message)
				return "", nil
			default:
				errStr := "invalid request kind"
				logWalletsModel.WithField("kind", kind).Errorln(errStr)
				return "", errors.New(errStr)
			}
		},
	)
}

func (walletModel *WalletModel) init() {
	logWalletsModel.Info("Initialize Wallet model")
	walletModel.SetRoles(map[int]*core.QByteArray{
		Name:              core.NewQByteArray2("name", -1),
		EncryptionEnabled: core.NewQByteArray2("encryptionEnabled", -1),
		Sky:               core.NewQByteArray2("sky", -1),
		CoinHours:         core.NewQByteArray2("coinHours", -1),
		FileName:          core.NewQByteArray2("fileName", -1),
		Expand:            core.NewQByteArray2("expand", -1),
		HasHardwareWallet: core.NewQByteArray2("hasHardwareWallet", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(walletModel, qml.QQmlEngine__CppOwnership)
	walletModel.ConnectData(walletModel.data)
	walletModel.ConnectSetData(walletModel.setData)
	walletModel.ConnectRowCount(walletModel.rowCount)
	walletModel.ConnectColumnCount(walletModel.columnCount)
	walletModel.ConnectRoleNames(walletModel.roleNames)

	walletModel.ConnectAddWallet(walletModel.addWallet)
	walletModel.ConnectEditWallet(walletModel.editWallet)
	walletModel.ConnectRemoveWallet(walletModel.removeWallet)
	walletModel.ConnectLoadModel(walletModel.loadModel)
	walletModel.ConnectUpdateModel(walletModel.updateModel)
	walletModel.ConnectSniffHw(walletModel.sniffHw)
}

// attachHwAsSigner add a hw as signer
func attachHwAsSigner(wlt fccore.Wallet) error {
	hw := hardware.NewSkyWallet(wlt)
	am := wlcore.LoadAltcoinManager()
	if err := am.AttachSignService(hw); err != nil {
		logSignersModel.Errorln("error registering hardware wallet as signer")
		return err
	}
	return nil
}

// sniffHw notify the model about available hardware wallet device if any
func (walletModel *WalletModel) sniffHw(qmlDevI *QDeviceInteraction, locker *QBridge) {
	blockingCheck := func() {
		registerWlt := func(wlt fccore.Wallet) {
			if wlt == nil {
				return
			}
			hadHwConnected = true
			walletModel.updateWallet(wlt.GetId())
			attachHwAsSigner(wlt)
		}
		logError := func(err error) {
			if err != nil {
				if err == fce.ErrWltFromAddrNotFound {
					logWalletsModel.WithError(err).Debugln("wallet not found")
					return
				}
				logWalletsModel.WithError(err).Errorln("unexpected error")
			}
		}
		dev := hardware.NewSkyWalletHelper()
		dev.FirstAddress(skyWallet.WalletTypeDeterministic).Then(func(data interface{}) interface{} {
			wlt, err := walletManager.WalletEnv.LookupWallet(data.(string))
			logError(err)
			registerWlt(wlt)
			return dev.FirstAddress(skyWallet.WalletTypeBip44)
		}).Then(func(data interface{}) interface{} {
			wlt, err := walletManager.WalletEnv.LookupWallet(data.(string))
			logError(err)
			registerWlt(wlt)
			return data
		}).Catch(func(err error) error {
			if hadHwConnected {
				hadHwConnected = false
				hwConnectedOn = []int{}
				beginIndex := walletModel.Index(0, 0, core.NewQModelIndex())
				endIndex := walletModel.Index(walletModel.rowCount(core.NewQModelIndex())-1, 0, core.NewQModelIndex())
				walletModel.DataChanged(beginIndex, endIndex, []int{HasHardwareWallet})
				logSignersModel.WithError(err).Info("connection to hardware wallet was lose")
			}
			if err == skyWallet.ErrNoDeviceConnected {
				logWalletsModel.Warningln("00000000000000000")
				return err
			}
			return err
		}).Await()
		openDialog := func(prompt func()) {
			locker.BeginUse()
			defer locker.EndUse()
			locker.lock()
			prompt()
			locker.lock()
			locker.unlock()
		}
		dev.ShouldBeInitialized().Then(func(data interface{}) interface{} {
			if data.(bool) {
				openDialog(qmlDevI.InitializeDevice)
			}
			return dev.ShouldBeSecured()
		}).Then(func(data interface{}) interface{} {
			if data.(bool) {
				openDialog(qmlDevI.SecureDevice)
			}
			return data
		}).Catch(func(err error) error {
			logWalletsModel.WithError(err).Errorln("uf")
			return err
		}).Await()
	}
	go func() {
		for {
			hwConnectedOn = []int{}
			blockingCheck()
			time.Sleep(time.Millisecond * 500)
		}
	}()
}

func (walletModel *WalletModel) updateWallet(fn string) {
	index := &core.QModelIndex{}
	for row := 0; row < walletModel.rowCount(core.NewQModelIndex()); row++ {
		index = walletModel.Index(row, 0, core.NewQModelIndex())
		fileName := walletModel.data(index, FileName)
		if  fileName.ToString() == fn {
			hwConnectedOn = append(hwConnectedOn, row)
			walletModel.DataChanged(index, index, []int{HasHardwareWallet})
			break
		}
	}
}

func (walletModel *WalletModel) data(index *core.QModelIndex, role int) *core.QVariant {
	logWalletsModel.Info("Loading data for index")
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(walletModel.Wallets()) {
		return core.NewQVariant()
	}

	var w = walletModel.Wallets()[index.Row()]

	switch role {
	case Name:
		{
			return core.NewQVariant1(w.Name())
		}

	case EncryptionEnabled:
		{
			return core.NewQVariant1(w.EncryptionEnabled())
		}

	case Sky:
		{
			return core.NewQVariant1(w.Sky())
		}

	case CoinHours:
		{
			accuracy, err := util.AltcoinQuotient(coin.CoinHoursTicker)
			if err != nil {
				logWalletsModel.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
			}
			val, err := strconv.ParseUint(w.CoinHours(), 10, 64)
			return core.NewQVariant1(util.FormatCoins(val, accuracy))
		}
	case FileName:
		{
			return core.NewQVariant1(w.FileName())
		}
	case HasHardwareWallet:
		{
			valInSlice := func() bool {
				for idx := range hwConnectedOn {
					if hwConnectedOn[idx] == index.Row() && hadHwConnected {
						return true
					}
				}
				return false
			}
			// FIXME: consider a double checking here instead of hadHwConnected
			// be careful this can have a big performance impact
			return core.NewQVariant1(valInSlice())
		}
	case Expand:
		{
			return core.NewQVariant1(w.IsExpand())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (walletModel *WalletModel) setData(index *core.QModelIndex, value *core.QVariant, role int) bool {

	if !index.IsValid() {
		return false
	}

	if index.Row() >= len(walletModel.Wallets()) {
		return false
	}

	var w = walletModel.Wallets()[index.Row()]

	switch role {
	case Name:
		{
			w.SetName(value.ToString())
		}
	case EncryptionEnabled:
		{
			w.SetEncryptionEnabled(value.ToInt(nil))
		}
	case Sky:
		{
			w.SetSky(value.ToString())
		}
	case CoinHours:
		{
			w.SetCoinHours(value.ToString())
		}
	case FileName:
		{
			w.SetFileName(value.ToString())
		}
	case Expand:
		{
			w.SetExpand(value.ToBool())
		}
	default:
		{
			return false
		}
	}

	walletModel.DataChanged(index, index, []int{role})
	return true
}

func (walletModel *WalletModel) rowCount(parent *core.QModelIndex) int {
	return len(walletModel.Wallets())
}

func (walletModel *WalletModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (walletModel *WalletModel) roleNames() map[int]*core.QByteArray {
	return walletModel.Roles()
}

func (walletModel *WalletModel) addWallet(w *QWallet) {
	logWalletsModel.Info("Add Wallet")
	if w.Pointer() == nil {
		return
	}
	walletModel.BeginInsertRows(core.NewQModelIndex(), len(walletModel.Wallets()), len(walletModel.Wallets()))
	qml.QQmlEngine_SetObjectOwnership(w, qml.QQmlEngine__CppOwnership)
	walletModel.SetWallets(append(walletModel.Wallets(), w))
	walletModel.SetCount(walletModel.Count() + 1)
	walletModel.EndInsertRows()
}

func (walletModel *WalletModel) editWallet(row int, name string, encrypted bool, sky string, coinHours string) {
	logWalletsModel.Info("Edit Wallet")
	pIndex := walletModel.Index(row, 0, core.NewQModelIndex())

	walletModel.setData(pIndex, core.NewQVariant1(name), Name)
	if encrypted {
		walletModel.setData(pIndex, core.NewQVariant1(1), EncryptionEnabled)
	} else {
		walletModel.setData(pIndex, core.NewQVariant1(0), EncryptionEnabled)
	}
	walletModel.setData(pIndex, core.NewQVariant1(sky), Sky)
	walletModel.setData(pIndex, core.NewQVariant1(coinHours), CoinHours)
}

func (walletModel *WalletModel) removeWallet(row int) {
	logWalletsModel.Info("Remove wallets for index")
	walletModel.BeginRemoveRows(core.NewQModelIndex(), row, row)
	walletModel.SetWallets(append(walletModel.Wallets()[:row], walletModel.Wallets()[row+1:]...))
	walletModel.SetCount(walletModel.Count() - 1)
	walletModel.EndRemoveRows()
}

func (walletModel *WalletModel) updateModel(wallets []*QWallet) {
	go func() {
		for i, wlt := range wallets {
			walletModel.editWallet(i, wlt.Name(), wlt.EncryptionEnabled() == 1, wlt.Sky(), wlt.CoinHours())
		}
	}()
}

func (walletModel *WalletModel) loadModel(wallets []*QWallet) {
	logWalletsModel.Info("Loading wallets")
	for _, wlt := range wallets {
		qml.QQmlEngine_SetObjectOwnership(wlt, qml.QQmlEngine__CppOwnership)
	}
	walletModel.BeginResetModel()
	walletModel.SetWallets(wallets)

	walletModel.EndResetModel()
	walletModel.SetCount(len(walletModel.Wallets()))
}
