package address

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"github.com/therecipe/qt/core"
)

var logAddressModel = logging.MustGetLogger("Address Model")

func init() {
	AddressDetails_QmlRegisterType2("HistoryModels", 1, 0, "QAddress")
	AddressList_QmlRegisterType2("HistoryModels", 1, 0, "QAddressList")
}

const (
	Address = int(core.Qt__UserRole) + 1<<iota
	AddressSky
	AddressCoinHours
)

type AddressDetails struct {
	core.QObject

	_ string `property:"address"`
	_ string `property:"addressSky"`
	_ string `property:"addressCoinHours"`
}

type AddressList struct {
	core.QAbstractListModel

	_ map[int]*core.QByteArray `property:"roles"`

	_ func() `constructor:"init"`

	_ func(transaction *AddressDetails) `signal:"addAddress,auto"`
	_ func(index int)                   `signal:"removeAddress,auto"`

	_ []*AddressDetails `property:"addresses"`
}

func (al *AddressList) init() {
	al.SetRoles(map[int]*core.QByteArray{
		Address:          core.NewQByteArray2("address", -1),
		AddressSky:       core.NewQByteArray2("addressSky", -1),
		AddressCoinHours: core.NewQByteArray2("addressCoinHours", -1),
	})

	al.ConnectRowCount(al.rowCount)
	al.ConnectData(al.data)
	al.ConnectRoleNames(al.roleNames)

}

func (al *AddressList) rowCount(*core.QModelIndex) int {
	return len(al.Addresses())
}

func (al *AddressList) roleNames() map[int]*core.QByteArray {
	return al.Roles()
}

func (al *AddressList) addAddress(address *AddressDetails) {
	logAddressModel.Info("Adding Address")
	al.BeginInsertRows(core.NewQModelIndex(), len(al.Addresses()), len(al.Addresses()))
	al.SetAddresses(append(al.Addresses(), address))
	al.EndInsertRows()
}

func (al *AddressList) removeAddress(index int) {
	logAddressModel.Info("Removing Address")
	al.BeginRemoveRows(core.NewQModelIndex(), index, index)
	al.SetAddresses(append(al.Addresses()[:index], al.Addresses()[index+1:]...))
	al.EndRemoveRows()
}

func (al *AddressList) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() >= len(al.Addresses()) {
		return core.NewQVariant()
	}

	address := al.Addresses()[index.Row()]

	switch role {
	case Address:
		{
			return core.NewQVariant1(address.Address())
		}
	case AddressCoinHours:
		{
			return core.NewQVariant1(address.AddressCoinHours())
		}
	case AddressSky:
		{
			return core.NewQVariant1(address.AddressSky())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}
