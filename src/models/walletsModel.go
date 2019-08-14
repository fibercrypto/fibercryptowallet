package models

import (
	"github.com/therecipe/qt/core"
)

const (
	Name              = int(core.Qt__UserRole) + 1
	EncryptionEnabled = int(core.Qt__UserRole) + 2
	Sky               = int(core.Qt__UserRole) + 3
	CoinHours         = int(core.Qt__UserRole) + 4
	FileName          = int(core.Qt__UserRole) + 5
)

type WalletModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QWallet               `property:"wallets"`

	_ func(*QWallet)                                                                    `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, sky float64, coinHours uint64) `slot:"editWallet"`
	_ func(row int)                                                                     `slot:"removeWallet"`
	_ func([]*QWallet)                                                                  `slot:"loadModel"`
	_ int                                                                               `property:"count"`
}

type QWallet struct {
	core.QObject

	_ string  `property:"name"`
	_ int     `property:"encryptionEnabled"`
	_ float64 `property:"sky"`
	_ uint64  `property:"coinHours"`
	_ string  `property:"fileName"`
}

func (m *WalletModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name:              core.NewQByteArray2("name", -1),
		EncryptionEnabled: core.NewQByteArray2("encryptionEnabled", -1),
		Sky:               core.NewQByteArray2("sky", -1),
		CoinHours:         core.NewQByteArray2("coinHours", -1),
		FileName:          core.NewQByteArray2("fileName", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddWallet(m.addWallet)
	m.ConnectEditWallet(m.editWallet)
	m.ConnectRemoveWallet(m.removeWallet)
	m.ConnectLoadModel(m.loadModel)

}

func (m *WalletModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Wallets()) {
		return core.NewQVariant()
	}

	var w = m.Wallets()[index.Row()]

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

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *WalletModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Wallets())
}

func (m *WalletModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *WalletModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *WalletModel) addWallet(w *QWallet) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Wallets()), len(m.Wallets()))
	m.SetWallets(append(m.Wallets(), w))
	m.EndInsertRows()
	m.updateCount()

}

func (m *WalletModel) editWallet(row int, name string, encrypted bool, sky float64, coinHours uint64) {
	w := m.Wallets()[row]
	w.SetName(name)
	w.SetEncryptionEnabled(0)
	if encrypted {
		w.SetEncryptionEnabled(1)
	}
	w.SetSky(sky)
	w.SetCoinHours(coinHours)
	w.SetFileName(w.FileName())

	pIndex := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{Name, EncryptionEnabled, Sky, CoinHours, FileName})
	m.updateCount()
}

func (m *WalletModel) removeWallet(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetWallets(append(m.Wallets()[:row], m.Wallets()[row+1:]...))
	m.EndRemoveRows()
	m.updateCount()

}

func (m *WalletModel) loadModel(wallets []*QWallet) {

	m.BeginResetModel()
	m.SetWallets(wallets)
	m.EndResetModel()
	m.updateCount()

}

func (m *WalletModel) updateCount() {
	m.SetCount(len(m.Wallets()))
}
