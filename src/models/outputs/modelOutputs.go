package outputs

import (
	"github.com/therecipe/qt/core"
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
	_ []*QOutput       			`property:"outputs"`
	_ string					`property:"address"`
}

type QOutput struct {
	core.QObject

	_ string 	`property:"outputID"`
	_ uint64 	`property:"addressSky"`
	_ uint64 	`property:"addressCoinHours"`
}

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
	return len(m.Outputs())
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

	qo := m.Outputs()[index.Row()]

	switch role{
	case OutputID:
		{
			return core.NewQVariant1(qo.OutputID())
		}
	case AddressSky:
		{
			return core.NewQVariant1(qo.AddressSky())
		}
	case AddressCoinHours:
		{
			return core.NewQVariant1(qo.AddressCoinHours())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}