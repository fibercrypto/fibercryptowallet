package models

import (
	"github.com/therecipe/qt/core"
)

const (
	OutputID = int(core.Qt__UserRole) + iota + 1
	AddressSky
	AddressCoinHours
	AddressOwner
	WalletOwner
)

type ModelOutputs struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QOutput               `property:"outputs"`
	_ string                   `property:"address"`

	_ func([]*QOutput) `slot:"addOutputs"`
	_ func([]*QOutput) `slot:"insertOutputs"`
	_ func([]*QOutput) `slot:"loadModel"`
	_ func()           `slot:"cleanModel"`
	_ func(string)     `slot:"removeOutputsFromAddress"`
	_ func(string)     `slot:"removeOutputsFromWallet"`
}

type QOutput struct {
	core.QObject

	_ string `property:"outputID"`
	_ string `property:"addressSky"`
	_ string `property:"addressCoinHours"`
	_ string `property:"addressOwner"`
	_ string `property:"walletOwner"`
}

func (m *ModelOutputs) init() {
	m.SetRoles(map[int]*core.QByteArray{
		OutputID:         core.NewQByteArray2("outputID", -1),
		AddressSky:       core.NewQByteArray2("addressSky", -1),
		AddressCoinHours: core.NewQByteArray2("addressCoinHours", -1),
		AddressOwner:     core.NewQByteArray2("addressOwner", -1),
		WalletOwner:      core.NewQByteArray2("walletOwner", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddOutputs(m.addOutputs)
	m.ConnectInsertOutputs(m.insertOutputs)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectCleanModel(m.cleanModel)
	m.ConnectRemoveOutputsFromAddress(m.removeOutputsFromAddress)
	m.ConnectRemoveOutputsFromWallet(m.removeOutputsFromWallet)
}

func (m *ModelOutputs) removeOutputsFromAddress(addr string) {
	old := m.Outputs()
	new := make([]*QOutput, 0)
	for _, out := range old {
		if out.AddressOwner() != addr {
			new = append(new, out)
		}
	}
	m.loadModel(new)
}

func (m *ModelOutputs) removeOutputsFromWallet(wltId string) {
	old := m.Outputs()
	new := make([]*QOutput, 0)
	for _, out := range old {
		if out.WalletOwner() != wltId {
			new = append(new, out)
		}
	}
	m.loadModel(new)
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

	if index.Row() >= len(m.Outputs()) {
		return core.NewQVariant()
	}

	qo := m.Outputs()[index.Row()]

	switch role {
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
	case AddressOwner:
		{
			return core.NewQVariant1(qo.AddressOwner())
		}
	case WalletOwner:
		{
			return core.NewQVariant1(qo.WalletOwner())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *ModelOutputs) insertRows(row int, count int) bool {
	m.BeginInsertRows(core.NewQModelIndex(), row, row+count)
	m.EndInsertRows()
	return true
}

func (m *ModelOutputs) addOutputs(mo []*QOutput) {
	m.SetOutputs(mo)
	m.insertRows(len(m.Outputs()), len(mo))
}

func (m *ModelOutputs) insertOutputs(mo []*QOutput) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Outputs()), len(m.Outputs())+len(mo)-1)
	m.SetOutputs(append(m.Outputs(), mo...))
	m.EndInsertRows()
}

func (m *ModelOutputs) loadModel(mo []*QOutput) {
	m.BeginResetModel()
	m.SetOutputs(mo)
	m.EndResetModel()
}

func (m *ModelOutputs) cleanModel() {
	m.BeginResetModel()
	m.SetOutputs(make([]*QOutput, 0))
	m.EndResetModel()
}
