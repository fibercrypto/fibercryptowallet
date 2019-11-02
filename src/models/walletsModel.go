package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

const (
	Name              = int(core.Qt__UserRole) + iota +1
	EncryptionEnabled
	Sky
	CoinHours
	FileName
	Expand
)

var logWalletsModel = logging.MustGetLogger("Wallets Model")

type WalletModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QWallet               `property:"wallets"`

	_ func(*QWallet)                                                                    `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, sky float64, coinHours uint64) `slot:"editWallet"`
	_ func(row int)                                                                     `slot:"removeWallet"`
	_ func([]*QWallet)                                                                  `slot:"loadModel"`
	_ func([]*QWallet, []bool)                                                          `slot:"loadModelUsingExpanded"`
	_ int                                                                               `property:"count"`
}

type QWallet struct {
	core.QObject
	_ string  `property:"name"`
	_ int     `property:"encryptionEnabled"`
	_ float64 `property:"sky"`
	_ uint64  `property:"coinHours"`
	_ string  `property:"fileName"`
	_ bool    `property:"expand"`
}

func (walletModel *WalletModel) init() {
	logWalletsModel.Info("Initialize Wallet model")
	walletModel.SetRoles(map[int]*core.QByteArray{
		Name:              core.NewQByteArray2("name", -1),
		EncryptionEnabled: core.NewQByteArray2("encryptionEnabled", -1),
		Sky:               core.NewQByteArray2("sky", -1),
		CoinHours:         core.NewQByteArray2("coinHours", -1),
		FileName:          core.NewQByteArray2("fileName", -1),
		Expand:          core.NewQByteArray2("expand", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(walletModel, qml.QQmlEngine__CppOwnership)
	walletModel.ConnectData(walletModel.data)
	walletModel.ConnectRowCount(walletModel.rowCount)
	walletModel.ConnectColumnCount(walletModel.columnCount)
	walletModel.ConnectRoleNames(walletModel.roleNames)

	walletModel.ConnectAddWallet(walletModel.addWallet)
	walletModel.ConnectEditWallet(walletModel.editWallet)
	walletModel.ConnectRemoveWallet(walletModel.removeWallet)
	walletModel.ConnectLoadModel(walletModel.loadModel)
	walletModel.ConnectLoadModelUsingExpanded(walletModel.loadModelUsingExpanded)

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
	walletModel.EndInsertRows()
	walletModel.SetCount(walletModel.Count() + 1)

}

func (walletModel *WalletModel) editWallet(row int, name string, encrypted bool, sky float64, coinHours uint64) {
	logWalletsModel.Info("Edit Wallet")
	w := walletModel.Wallets()[row]
	w.SetName(name)
	w.SetEncryptionEnabled(0)
	if encrypted {
		w.SetEncryptionEnabled(1)
	}

	w.SetSky(sky)
	w.SetCoinHours(coinHours)
	w.SetFileName(w.FileName())

	pIndex := walletModel.Index(row, 0, core.NewQModelIndex())
	walletModel.DataChanged(pIndex, pIndex, []int{Name, EncryptionEnabled, Sky, CoinHours, FileName, Expand})

}

func (walletModel *WalletModel) removeWallet(row int) {
	logWalletsModel.Info("Remove wallets for index")
	walletModel.BeginRemoveRows(core.NewQModelIndex(), row, row)
	walletModel.SetWallets(append(walletModel.Wallets()[:row], walletModel.Wallets()[row+1:]...))
	walletModel.EndRemoveRows()
	walletModel.SetCount(walletModel.Count() - 1)

}

func (walletModel *WalletModel) loadModelUsingExpanded(wallets []*QWallet, exp []bool) {
	logWalletsModel.Info("Loading wallets")
	for i, wlt := range wallets {
		if i < len(exp){
			wallets[i].SetExpand(exp[i])
		}
		qml.QQmlEngine_SetObjectOwnership(wlt, qml.QQmlEngine__CppOwnership)
	}
	walletModel.BeginResetModel()
	walletModel.SetWallets(wallets)

	walletModel.EndResetModel()
	walletModel.SetCount(len(walletModel.Wallets()))
}


func (walletModel *WalletModel) loadModel(wallets []*QWallet) {
	logWalletsModel.Info("Loading wallets")
	for _, wlt := range wallets {
		//wallets[i].SetSky(58)
		qml.QQmlEngine_SetObjectOwnership(wlt, qml.QQmlEngine__CppOwnership)
	}
	walletModel.BeginResetModel()
	walletModel.SetWallets(wallets)

	walletModel.EndResetModel()
	walletModel.SetCount(len(walletModel.Wallets()))
}
