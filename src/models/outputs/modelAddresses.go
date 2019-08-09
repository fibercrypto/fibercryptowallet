package outputs

import (
	"github.com/therecipe/qt/core"
)

const (
	Address = int(core.Qt__UserRole) + 1
)

type ModelAddresses struct {
	core.QAbstractListModel

	_ func()                    `constructor:"init"`

	_ map[int]*core.QByteArray  `property:"roles"`
	_ []*ModelOutputs        	`property:"outputs"`
	_ string					`property:"name"`

	_ func([]*ModelOutputs) 	`slot:"addOutputs"`
}

func (m *ModelAddresses) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Address: 			 core.NewQByteArray2("address", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddOutputs(m.addOutputs)
}

func (m *ModelAddresses) rowCount(*core.QModelIndex) int {
	println("Row count for ModelAddresses:", len(m.Outputs()))
	//println("Address of ModelAddresses: ", m)
	return len(m.Outputs())
}

func (m *ModelAddresses) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *ModelAddresses) data(index *core.QModelIndex, role int) *core.QVariant {
	println("data requested for ModelAddresses")
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

func (m *ModelAddresses) insertRows(row int, count int) bool {
	m.BeginInsertRows(core.NewQModelIndex(), row, row + count)
	m.EndInsertRows()
	return true
}

func (m *ModelAddresses) addOutputs(mo []*ModelOutputs) {
	m.SetOutputs(mo)
	m.insertRows(len(m.Outputs()), len(mo))
}