package models

import (
	"github.com/therecipe/qt/core"
)

const (
	OAddress = int(core.Qt__UserRole) + iota + 1
	QOutputs
)

type ModelAddresses struct {
	core.QAbstractListModel
	outputs []*ModelOutputs

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ string                   `property:"name"`
	_ string                   `property:"id"`

	_ func([]*ModelOutputs) `slot:"addOutputs"`
}

func (m *ModelAddresses) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Address:  core.NewQByteArray2("address", -1),
		QOutputs: core.NewQByteArray2("qoutputs", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddOutputs(m.addOutputs)
	m.outputs = make([]*ModelOutputs, 0)
}

func (m *ModelAddresses) rowCount(*core.QModelIndex) int {
	return len(m.outputs)
}

func (m *ModelAddresses) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *ModelAddresses) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.outputs) {
		return core.NewQVariant()
	}

	wa := m.outputs[index.Row()]

	switch role {
	case OAddress:
		{
			return core.NewQVariant1(wa.Address())
		}
	case QOutputs:
		{
			return core.NewQVariant1(wa)
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *ModelAddresses) insertRows(row int, count int) bool {
	m.BeginInsertRows(core.NewQModelIndex(), row, row+count)
	m.EndInsertRows()
	return true
}

func (m *ModelAddresses) addOutputs(mo []*ModelOutputs) {
	for row, mOut := range mo {
		find := false
		for _, mOutSet := range m.outputs {
			if mOut.Address() == mOutSet.Address() {
				mOutSet.addOutputs(mOut.outputs)
				find = true
				break
			}
		}
		if !find  {
			m.BeginInsertRows(core.NewQModelIndex(), row, row)
			m.outputs = append(m.outputs, mOut)
			m.EndInsertRows()
		} else {
			m.DataChanged(m.Index(len(m.outputs)-1, row, core.NewQModelIndex()), m.Index(len(m.outputs)-1, row+1, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
		}
	}
}
