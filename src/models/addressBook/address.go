package addressBook

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	qtcore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
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
	// _ map[int]*qtcore.QByteArray `slot:"addAddress"`
}

func (m *AddrsBkAddressModel) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Value:    qtcore.NewQByteArray2("value", -1),
		CoinType: qtcore.NewQByteArray2("coin", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(m, qml.QQmlEngine__CppOwnership)
	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
	// m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

}

func (m *AddrsBkAddressModel) rowCount(*qtcore.QModelIndex) int {
	return len(m.Address())
}

func (m *AddrsBkAddressModel) roleNames() map[int]*qtcore.QByteArray {
	return m.Roles()
}

func (m *AddrsBkAddressModel) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	logAddressBook.Info("Loading address data")
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}

	if index.Row() >= len(m.Address()) {
		return qtcore.NewQVariant()
	}

	address := m.Address()[index.Row()]

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

// func (m *AddrsBkAddressModel) insertRows(row int, count int) bool {
// 	m.BeginInsertRows(qtcore.NewQModelIndex(), row, row+count)
// 	m.EndInsertRows()
// 	return true
// }
// func (m *AddrsBkAddressModel) addAddress(mo []*AddrsBookModel) {
//
// 	// m.insertRows(len(m.Outputs()), len(mo))
// }

func FromAddressToQAddress(addresses []core.ReadableAddress) []*QAddress {
	var qAddresses = make([]*QAddress, 0)
	for _, addrs := range addresses {
		qa := NewQAddress(nil)
		qa.SetCoinType(string(addrs.GetCoinType()))
		qa.SetValue(string(addrs.GetValue()))
		qAddresses = append(qAddresses, qa)
	}
	// logAddressBook.Infof("%#v",qAddresses)
	return qAddresses
}
