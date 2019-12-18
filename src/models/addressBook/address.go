package addressBook

import (
	"bytes"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	qtcore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sort"
)

const (
	Value = int(qtcore.Qt__UserRole + (iota + 1))
	CoinType
)

func init() { AddrsBkAddressModel_QmlRegisterType2("AddrsBookManager", 1, 0, "AddrsBkAddressModel") }

type QAddress struct {
	qtcore.QObject
	_ string `property:"value"`
	_ string `property:"coinType"`
}

type AddrsBkAddressModel struct {
	qtcore.QAbstractListModel
	_ func()                     `constructor:"init"`
	_ []*QAddress                `property:"address"`
	_ map[int]*qtcore.QByteArray `property:"roles"`
}

func (addrsBookAddressModel *AddrsBkAddressModel) init() {
	addrsBookAddressModel.SetRoles(map[int]*qtcore.QByteArray{
		Value:    qtcore.NewQByteArray2("value", -1),
		CoinType: qtcore.NewQByteArray2("coinType", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(addrsBookAddressModel, qml.QQmlEngine__CppOwnership)
	addrsBookAddressModel.ConnectRowCount(addrsBookAddressModel.rowCount)
	addrsBookAddressModel.ConnectData(addrsBookAddressModel.data)
	// addrsBookAddressModel.ConnectColumnCount(addrsBookAddressModel.columnCount)
	addrsBookAddressModel.ConnectRoleNames(addrsBookAddressModel.roleNames)

}

func (addrsBookAddressModel *AddrsBkAddressModel) rowCount(*qtcore.QModelIndex) int {
	return len(addrsBookAddressModel.Address())
}

func (addrsBookAddressModel *AddrsBkAddressModel) roleNames() map[int]*qtcore.QByteArray {
	return addrsBookAddressModel.Roles()
}

func (addrsBookAddressModel *AddrsBkAddressModel) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {

	logAddressBook.Info("Loading address data")
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}

	if index.Row() >= len(addrsBookAddressModel.Address()) {
		return qtcore.NewQVariant()
	}

	address := addrsBookAddressModel.Address()[index.Row()]

	switch role {
	case Value:
		{
			return qtcore.NewQVariant1(address.Value())
		}
	case CoinType:
		{
			return qtcore.NewQVariant1(address.CoinType())
		}
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func fromAddressToQAddress(addresses []core.StringAddress) []*QAddress {
	var qAddressesList = make([]*QAddress, 0)
	sort.Slice(addresses, func(i, j int) bool {
		return bytes.Compare(addresses[i].GetCoinType(), addresses[j].GetCoinType()) == -1
	})
	for _, addrs := range addresses {
		qAddress := NewQAddress(nil)
		qAddress.SetCoinType(string(addrs.GetCoinType()))
		qAddress.SetValue(string(addrs.GetValue()))
		qAddressesList = append(qAddressesList, qAddress)
	}
	return qAddressesList
}
