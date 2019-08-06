package outputs

import (
	"github.com/therecipe/qt/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

const (
	Address = int(core.Qt__UserRole) + 1
)

type ModelAddresses struct {
	core.QAbstractListModel

	_ func()                    `constructor:"init"`

	_ map[int]*core.QByteArray  `property:"roles"`
	_ []*ModelOutputs        `property:"outputs"`
}

func (m *ModelAddresses) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Address: 			 core.NewQByteArray2("address", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)

}

func (m *ModelAddresses) rowCount(*core.QModelIndex) int {
	return len(m.Transactions())
}

func (m *ModelAddresses) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *ModelAddresses) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Outputs()){
		return core.NewQVariant()
	}

	wa := m.Outputs()[index.Row()]

	switch role{
	case Address:
		{
			return core.NewQVariant1(wa.Address())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}