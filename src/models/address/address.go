package address

import (
	"github.com/therecipe/qt/core"
)

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

	_ string  `property:"address"`
	_ float32 `property:"addressSky"`
	_ int     `property:"addressCoinHours"`
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

	al.addExamples()
}

func (al *AddressList) rowCount(*core.QModelIndex) int {
	return len(al.Addresses())
}

func (al *AddressList) roleNames() map[int]*core.QByteArray {
	return al.Roles()
}

func (al *AddressList) addAddress(address *AddressDetails) {
	al.BeginInsertRows(core.NewQModelIndex(), len(al.Addresses()), len(al.Addresses()))
	al.SetAddresses(append(al.Addresses(), address))
	al.EndInsertRows()
}

func (al *AddressList) removeAddress(index int) {
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

func (al *AddressList) addExamples() {
	adr := NewAddressDetails(nil)
	adr.SetAddress("734irweaweygtawieta783cwyc")
	adr.SetAddressSky(38)
	adr.SetAddressCoinHours(5048)
	al.addAddress(adr)
	adr1 := NewAddressDetails(nil)
	adr1.SetAddress("ekq03i3qerwhjqoqh9823yurig")
	adr1.SetAddressSky(61)
	adr1.SetAddressCoinHours(9456)
	al.addAddress(adr1)
	adr2 := NewAddressDetails(nil)
	adr2.SetAddress("1kjher73yiner7wn32nkuwe94v")
	adr2.SetAddressSky(1)
	adr2.SetAddressCoinHours(24)
	al.addAddress(adr2)
	adr3 := NewAddressDetails(nil)
	adr3.SetAddress("oopfwwklfd34iuhjwe83w3h28r")
	adr3.SetAddressSky(111)
	adr3.SetAddressCoinHours(13548)
	al.addAddress(adr3)
}
