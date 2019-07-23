package networking

import (
	"github.com/therecipe/qt/core"
	//"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

const(
	Peer	=	int(core.Qt__UserRole) + 1
	Source	=	int(core.Qt__UserRole) + 2
	BlockHeight	=	int(core.Qt__UserRole) + 3
	LastSeen	=	int(core.Qt__UserRole) + 4
)


type networkingModel struct {
	core.QAbstractListModel

	_ func()	`constructor:"init"`
	_ map[int]*core.QByteArray	`property:"roles"`
	_ []*QConnection	`property:"connections"`

	
	_ func()	`slot:"loadModel"`
	
}

func (m *networkingModel) init(){
	m.SetRoles(map[int]*core.QByteArray{
		Peer:	core.NewQByteArray2("peer", -1),
		Source:	core.NewQByteArray2("source", -1),
		BlockHeight: core.NewQByteArray2("blockHeight", -1),
		LastSeen:	core.NewQByteArray2("lastSeen", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoles(m.roles)

	m.ConnectLoadModel(m.loadModel)

	m.loadModel()
}

func (m *networkingModel) data(index *core.QModelIndex, role int) *core.QVariant{
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Connections()){
		return core.NewQVariant()
	}

	var c = m.Connections()[index.Row()]

	switch role {
	case Peer:
		{
			return core.NewQVariant1(c.Peer())
		}
	case Source:
		{
			return core.NewQVariant1(c.Source())
		}
	case BlockHeight:
		{
			return core.NewQVariant1(c.BlockHeight())
		}
	case LastSeen:
		{
			return core.NewQVariant1(c.LastSeen())
		}
	default:
		{
			return core.NewQVariant()
		}
		
	}
}


func (m *networkingModel) rowCount(parent *core.QModelIndex) int{
	return len(m.Connections())
}

func (m *networkingModel) columnCount(parent *core.QModelIndex) int{
	return 1
}

func (m *networkingModel) roles() map[int]*core.QByteArray{
	return m.Roles()
}

func (m* networkingModel) loadModel(){
	qconnections := getConnections()
	if qconnections == nil {
		return 
	}

	m.SetConnections(qconnections)
}


type QConnection struct{
	core.QObject
	_ string	`property:"peer"`
	_ string	`property:"source"`
	_ uint64	`property:"blockHeight"`
	_ int64		`property:"lastSeen"`
	
}


func getConnections() []*QConnection{
	c := newClient()

	defaultConnections, err := c.NetworkDefaultPeers()
	if err != nil {
		return nil
	}

	connections, err := c.NetworkConnections(nil)
	if err != nil {
		return nil
	}

	qconnections := make([]*QConnection, 0)
	for _, conn := range connections.Connections{
		qconnection := connectionResponseToQConnection(conn)
		if contains(defaultConnections, qconnection.Peer()){
			qconnection.SetPeer("Default Peer")
		}
		qconnections = append(qconnections, qconnection)
	}

	return qconnections


}

func connectionResponseToQConnection(connection readable.Connection) *QConnection{
	qconnection  := NewQConnection(nil)
	qconnection.SetPeer(connection.Addr)
	qconnection.SetSource("Peer Exchange")
	qconnection.SetBlockHeight(connection.Height)
	qconnection.SetLastSeen(connection.LastReceived)
	return qconnection
}

func contains(list []string, item string) bool{
	for _, a := range list{
		if a == item{
			return true
		}
	}
	return false
}
