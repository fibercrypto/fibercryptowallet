package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/hardware"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	fce "github.com/fibercrypto/FiberCryptoWallet/src/errors"
	"github.com/sirupsen/logrus"
	wlcore "github.com/fibercrypto/FiberCryptoWallet/src/main"
	messages "github.com/fibercrypto/skywallet-protob/go"
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
var hwConnectedOn = -1

type WalletModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QWallet               `property:"wallets"`

	_ func(*QWallet)                                                                    `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, sky string, coinHours string)  `slot:"editWallet"`
	_ func(row int)                                                                     `slot:"removeWallet"`
	_ func([]*QWallet)                                                                  `slot:"loadModel"`
	_ func([]*QWallet)                                                                  `slot:"updateModel"`
	_ func()                                                                            `slot:"sniffHw"`
	_ int                                                                               `property:"count"`
}

type QWallet struct {
	core.QObject
	_ string  `property:"name"`
	_ int     `property:"encryptionEnabled"`
	_ string  `property:"sky"`
	_ string  `property:"coinHours"`
	_ string  `property:"fileName"`
	_ bool    `property:"expand"`
	_ bool    `property:"hasHardwareWallet"`
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
func attachHwAsSigner(dev skyWallet.Devicer) error {
	pb := func(dev skyWallet.Devicer, prvMsg wire.Message, nextMsg messages.MessageType) (wire.Message, error) {
		msg, err := dev.ButtonAck()
		if err != nil {
			// TODO i18n
			logrus.WithError(err).Errorln("unexpected error")
			return msg, err
		}
		if msg.Kind != uint16(nextMsg) {
			if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
				str, err := skyWallet.DecodeFailMsg(msg)
				if err != nil {
					// TODO i18n
					logrus.WithField("msg", msg).Errorln("error decoding msg")
					return msg, fce.ErrTxnSignFailure
				}
				// FIXME: this can become broken with device internationalization
				if str == "Action cancelled by user" {
					return msg, fce.ErrHwSignTransactionCanceled
				}
			}
			// TODO i18n
			logrus.WithFields(
				logrus.Fields{
					"expected": nextMsg,
					"actual": msg.Kind}).Errorln("unexpected msg type")
			return msg, fce.ErrTxnSignFailure
		}
		// FIXME maybe only a MessageType_MessageType_TransactionSign
		if msg.Kind == uint16(messages.MessageType_MessageType_Success) {
			successMsg, err := skyWallet.DecodeSuccessMsg(msg)
			if err != nil {
				// TODO i18n
				logrus.WithError(err).Errorln("error decoding msg")
				return wire.Message{}, err
			}
			// TODO i18n
			logrus.Debugln("signing transaction with hw", successMsg)
			return msg, err// FIXME bug in err value
		}
		return msg, nil
	}
	cb := func(dev skyWallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error) {
		var msg wire.Message
		for outsLen > 0 {
			var err error
			msg, err = pb(dev, prvMsg, messages.MessageType_MessageType_ButtonRequest)
			if err != nil {
				return wire.Message{}, err
			}
			if outsLen == 1 {
				msg, err = pb(dev, prvMsg, messages.MessageType_MessageType_ResponseTransactionSign)
			} else {
				msg, err = pb(dev, prvMsg, messages.MessageType_MessageType_ButtonRequest)
			}
			if err != nil {
				return wire.Message{}, err
			}
			outsLen--
		}
		return msg, nil
	}
	hw := hardware.NewSkyWallet(dev, cb)
	am := wlcore.LoadAltcoinManager()
	if err := am.AttachSignService(hw); err != nil {
		logrus.Errorln("error registering hardware wallet as signer")
		return err
	}
	return nil
}

// sniffHw notify the model about available hardware wallet device if any
func (walletModel *WalletModel) sniffHw() {
	dev := skyWallet.NewDevice(skyWallet.DeviceTypeUSB)
	addr, err := hardware.HwFirstAddr(dev)
	if err == nil {
		wlt, err := walletManager.WalletEnv.GetWallet(addr)
		if err != nil {
			// TODO i18n
			logrus.Warnln("can not find a wallet matching the hardware one")
			// FIXME handle this scenario with a wallet registration.
			return
		}
		err = attachHwAsSigner(dev)
		if err != nil {
			logrus.WithError(err).Errorln("unable to attach signer")
			return
		}
		hadHwConnected = true
		walletModel.updateWallet(wlt.GetId())
	} else {
		if hadHwConnected {
			hadHwConnected = false
			beginIndex := walletModel.Index(0, 0, core.NewQModelIndex())
			endIndex := walletModel.Index(walletModel.rowCount(core.NewQModelIndex()) - 1, 0, core.NewQModelIndex())
			walletModel.DataChanged(beginIndex, endIndex, []int{HasHardwareWallet})
			logrus.WithError(err).Info("connection to hardware wallet was lose")
		}
	}
}

func (walletModel *WalletModel) updateWallet(fn string) {
	index := &core.QModelIndex{}
	for row := 0; row < walletModel.rowCount(core.NewQModelIndex()); row++ {
		index = walletModel.Index(row, 0, core.NewQModelIndex())
		fileName := walletModel.data(index, FileName)
		if  fileName.ToString() == fn {
			hwConnectedOn = row
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
			return core.NewQVariant1(w.CoinHours())
		}
	case FileName:
		{
			return core.NewQVariant1(w.FileName())
		}
	case HasHardwareWallet:
		{
			// FIXME: consider a double checking here instead of hadHwConnected
			// be careful this can have a big performance impact
			return core.NewQVariant1(hwConnectedOn == index.Row() && hadHwConnected)
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
	for i, wlt := range wallets {
		walletModel.editWallet(i, wlt.Name(), wlt.EncryptionEnabled() == 1, wlt.Sky(), wlt.CoinHours())
	}
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
