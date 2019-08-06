package outputs

import (
	"github.com/therecipe/qt/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

const (
	OutputID = int(core.Qt__UserRole) + 1
	AddressSky = int(core.Qt__UserRole) + 2
	AddressCoinHours = int(core.Qt__UserRole) + 3
)

type ModelOutputs struct {
	core.QAbstractListModel

	_ func()                    `constructor:"init"`

	_ map[int]*core.QByteArray  `property:"roles"`
	_ []*Output       `property:"outputs"`
}

//TODO: Implement Output

func (m *ModelOutputs) init() {
	m.SetRoles(map[int]*core.QByteArray{
		OutputID: 			 core.NewQByteArray2("outputID", -1),
		AddressSky: 	     core.NewQByteArray2("addressSky", -1),
		AddressCoinHours: 	 core.NewQByteArray2("addressCoinHours", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)

}

func (m *ModelOutputs) rowCount(*core.QModelIndex) int {
	return len(m.Transactions())
}

func (m *ModelOutputs) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *ModelOutputs) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Outputs()){
		return core.NewQVariant()
	}

	ao := m.Outputs()[index.Row()]

	switch role{
	case OutputID:
		{
			return core.NewQVariant1(ao.OutputID())
		}
	case AddressSky:
		{
			return core.NewQVariant1(ao.AddressSky())
		}
	case AddressCoinHours:
		{
			return core.NewQVariant1(ao.AddressCoinHours())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}