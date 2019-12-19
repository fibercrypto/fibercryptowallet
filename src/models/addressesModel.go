package models

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

var logAddressesModel = logging.MustGetLogger("Addresses Model")

const (
	Address    = int(core.Qt__UserRole) + 1
	ASky       = int(core.Qt__UserRole) + 2
	ACoinHours = int(core.Qt__UserRole) + 3
	AMarked    = int(core.Qt__UserRole) + 4
	AWallet    = int(core.Qt__UserRole) + 5
	AWalletId  = int(core.Qt__UserRole) + 6
)

type AddressesModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QAddress              `property:"addresses"`

	_ func(*QAddress)                        `slot:"addAddress"`
	_ func([]*QAddress)                      `slot:"addAddresses"`
	_ func(int)                              `slot:"removeAddress"`
	_ func(string)                           `slot:"updateModel"`
	_ func(int, string, uint64, uint64, int) `slot:"editAddress"`
	_ func([]*QAddress)                      `slot:"loadModel"`
	_ int                                    `property:"count"`
	_ func(string)                           `slot:"removeAddressesFromWallet"`
}

type QAddress struct {
	core.QObject

	_ string `property:"address"`
	_ string `property:"addressSky"`
	_ string `property:"addressCoinHours"`
	_ int    `property:"marked"`
	_ string `property:"wallet"`
	_ string `property:"walletId"`
}

func (m *AddressesModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Address:    core.NewQByteArray2("address", -1),
		ASky:       core.NewQByteArray2("addressSky", -1),
		ACoinHours: core.NewQByteArray2("addressCoinHours", -1),
		AMarked:    core.NewQByteArray2("marked", -1),
		AWallet:    core.NewQByteArray2("wallet", -1),
		AWalletId:  core.NewQByteArray2("walletId", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(m, qml.QQmlEngine__CppOwnership)
	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddAddress(m.addAddress)
	m.ConnectUpdateModel(m.updateModel)
	m.ConnectEditAddress(m.editAddress)
	m.ConnectRemoveAddress(m.removeAddress)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectRemoveAddressesFromWallet(m.removeAddressesFromWallet)
	m.ConnectAddAddresses(m.addAddresses)
	m.SetCount(0)

}

func (m *AddressesModel) removeAddressesFromWallet(wltId string) {
	old := m.Addresses()
	new := make([]*QAddress, 0)
	for _, addr := range old {
		if addr.WalletId() != wltId {
			new = append(new, addr)
		}
	}
	m.loadModel(new)
	m.removeAddress(0)
}
func (m *AddressesModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Addresses()) {
		return core.NewQVariant()
	}

	var a = m.Addresses()[index.Row()]

	switch role {

	case Address:
		{
			return core.NewQVariant1(a.Address())
		}
	case ASky:
		{
			return core.NewQVariant1(a.AddressSky())
		}
	case ACoinHours:
		{
			return core.NewQVariant1(a.AddressCoinHours())
		}
	case AMarked:
		{
			return core.NewQVariant1(a.Marked())
		}
	case AWallet:
		{
			return core.NewQVariant1(a.Wallet())
		}
	case AWalletId:
		{
			return core.NewQVariant1(a.WalletId())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *AddressesModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Addresses())
}

func (m *AddressesModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *AddressesModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *AddressesModel) addAddress(address *QAddress) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Addresses()), len(m.Addresses()))
	qml.QQmlEngine_SetObjectOwnership(address, qml.QQmlEngine__CppOwnership)
	m.SetAddresses(append(m.Addresses(), address))
	m.EndInsertRows()
	m.SetCount(m.Count() + 1)
}

func (m *AddressesModel) addAddresses(addresses []*QAddress) {
	for _, addr := range addresses {
		m.addAddress(addr)
	}
}

func (m *AddressesModel) removeAddress(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetAddresses(append(m.Addresses()[:row], m.Addresses()[row+1:]...))
	m.EndRemoveRows()
	m.SetCount(m.Count() - 1)
}

func (m *AddressesModel) editAddress(row int, address string, sky, coinHours uint64, marked int) {
	a := m.Addresses()[row]
	a.SetAddress(address)
	accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
	if err != nil {
		logAddressesModel.WithError(err).Warn("Couldn't get" + params.SkycoinTicker + " quotient")
		return
	}
	a.SetAddressSky(util.FormatCoins(sky, accuracy))
	accuracy, err = util.AltcoinQuotient(params.CoinHoursTicker)
	if err != nil {
		logAddressesModel.WithError(err).Warn("Couldn't get" + params.CoinHoursTicker + " quotient")
		return
	}
	a.SetAddressCoinHours(util.FormatCoins(coinHours, accuracy))
	changeMarked := true
	if marked == a.Marked() {
		changeMarked = false
	}
	a.SetMarked(marked)
	pIndex := m.Index(row, 0, core.NewQModelIndex())
	if changeMarked {
		m.DataChanged(pIndex, pIndex, []int{Address, ASky, ACoinHours, AMarked})
	} else {
		m.DataChanged(pIndex, pIndex, []int{Address, ASky, ACoinHours})
	}

}

func (m *AddressesModel) updateModel(fileName string) {
	go func() {
		walletManager.updateAddresses(fileName)
		m.LoadModel(walletManager.getAddresses(fileName))
	}()
}

func (m *AddressesModel) loadModel(Qaddresses []*QAddress) {
	for _, addr := range Qaddresses {
		qml.QQmlEngine_SetObjectOwnership(addr, qml.QQmlEngine__CppOwnership)
	}
	addresses := make([]*QAddress, 0)
	address := NewQAddress(nil)
	address.SetAddress("--------------------------")
	address.SetAddressSky("0")
	address.SetAddressCoinHours("0")
	address.SetMarked(0)
	qml.QQmlEngine_SetObjectOwnership(address, qml.QQmlEngine__CppOwnership)
	addresses = append(addresses, address)
	addresses = append(addresses, Qaddresses...)

	m.BeginResetModel()
	m.SetAddresses(addresses)
	m.SetCount(len(addresses))

	m.EndResetModel()

}
