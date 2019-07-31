package models

import "github.com/therecipe/qt/core"

// ip, port, source, block, lastSeenIn, lastSeenOut
const (
	Ip	              = int(core.Qt__UserRole) + 1
	Port			  = int(core.Qt__UserRole) + 2
	Source            = int(core.Qt__UserRole) + 3
	Block	          = int(core.Qt__UserRole) + 4
	LastSeenIn        = int(core.Qt__UserRole) + 5
	LastSeenOut       = int(core.Qt__UserRole) + 6
)

type NetworkingModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QNetworking           `property:"networks"`

	_ func(*QNetworking)                                                        `slot:"addNetwork"`
	_ func(row int)                                                             `slot:"removeNetwork"`
	_ func([]*QNetworking)                                                      `slot:"loadModel"`
	_ int                                                                       `property:"count"`
}

type QNetworking struct {
	core.QObject
// ip, port, source, block, lastSeenIn, lastSeenOut
	_ string `property:"ip"`
	_ int    `property:"port"`
	_ uint64 `property:"source"`
	_ uint64 `property:"block"`
	_ string `property:"lastSeenIn"`
	_ string `property:"lastSeenOut"`
}

func (m *NetworkingModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Ip:             core.NewQByteArray2("ip", -1),
		Port: 			core.NewQByteArray2("port", -1),
		Source:         core.NewQByteArray2("source", -1),
		Block:          core.NewQByteArray2("block", -1),
		LastSeenIn:     core.NewQByteArray2("lastSeenIn", -1),
		LastSeenOut:    core.NewQByteArray2("lastSeenOut", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddNetwork(m.addNetwork)
	m.ConnectRemoveWallet(m.removeNetwork)
	m.ConnectLoadModel(m.loadModel)

}

func (m *NetworkingModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.networks()) {
		return core.NewQVariant()
	}

	var w = m.networks()[index.Row()]

	switch role {
	case Ip:
		{
			return core.NewQVariant1(w.Ip())
		}

	case Port:
		{
			return core.NewQVariant1(w.Port())
		}

	case Source:
		{
			return core.NewQVariant1(w.Source())
		}

	case Block:
		{
			return core.NewQVariant1(w.Block())
		}
	case LastSeenIn:
		{
			return core.NewQVariant1(w.LastSeenIn())
		}
	case LastSeenOut:
		{
			return core.NewQVariant1(w.LastSeenOut())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *NetworkingModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Networks())
}

func (m *NetworkingModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *NetworkingModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *NetworkingModel) addNetwork(w *QNetworking) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Networks()), len(m.Networks()))
	m.SetWallets(append(m.Networks(), w))
	m.EndInsertRows()
	m.updateCount()

}

func (m *WalletModel) removeNetwork(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetWallets(append(m.Networks()[:row], m.Networks()[row+1:]...))
	m.EndRemoveRows()
	m.updateCount()

}

func (m *NetworkingModel) loadModel(networks []*QNetworking) {

	m.BeginResetModel()
	m.SetWallets(networks)
	m.EndResetModel()
	m.updateCount()

}

func (m *NetworkingModel) updateCount() {
	m.SetCount(len(m.Networks()))
}
