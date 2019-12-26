package explorer

import (
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtcore "github.com/therecipe/qt/core"
)

// Properties:

// For blocks:
// - Time
// - Block number
// - Transactions (number of Transactions)
// - Blockhash

// For block's details:
// - Height
// - Timestamp
// - Size
// - Hash
// - Parent Hash
// - Total Amount

// - Next block
// - Previous block

// For Transactions:
// - Inputs:
// 	* Coins
// 	* Initial hours
// 	* Final hours

// - Outputs:
// 	* Coins
// 	* Hours

// - Transaction fee (in hours)

// - Links for the transactions

var logExplorer = logging.MustGetLogger("Explorer")

const (
	Time = int(qtcore.Qt__UserRole) + iota + 1
	BlockNumber
	Transactions
	Blockhash
)

// BlocksModel List of Blocks to be show.
type BlocksModel struct {
	qtcore.QAbstractListModel

	_ func()            		`constructor:"init"`

	_ int               		`property:"currentPage"`
	_ map[int]*core.QByteArray  `property:"roles"`
	_ []*QBlock                 `property:"blocks"`

	_ func(pageNum int) 		`signal:"loadPage"`
	_ func()            		`signal:"update,auto"`

	_ func([]*QBlock) 			`slot:"addBlocks"`
	_ func([]*QBlock) 			`slot:"loadModel"`
}

// QBlock Contains info about the block to be show.
type QBlock struct {
	core.QObject

	_ string `property:"time"`
	_ string `property:"blockNumber"`
	_ string `property:"Transactions"`
	_ string `property:"Blockhash"`
}

func (blocksM *BlocksModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		QBlocks: core.NewQByteArray2("qblocks", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddBlocks(m.addBlocks)
	m.ConnectInsertBlocks(m.insertBlocks)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectCleanModel(m.cleanModel)

	m.loadModel()
}

func (blocksM *BlocksModel) update() {
	// update info
	if err := blocksM.updateInfo(); err != nil {
		logExplorer.WithError(err).Warn("Couldn't update blockchain Info")
		return
	}
	return
}

// updateInfo request the needed information
func (blocksM *BlocksModel) updateInfo() error {
	return nil
}

func (blocksM *BlocksModel) loadPage(pageNum int) {

}

func (m *BlocksModel) rowCount(*core.QModelIndex) int {
	return len(m.Outputs())
}

func (m *BlocksModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *BlocksModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() >= len(m.Blocks()) {
		return core.NewQVariant()
	}

	qb := m.Blocks()[index.Row()]

	switch role {
	case Time:
		{
			return core.NewQVariant1(qb.Time())
		}
	case BlockNumber:
		{
			return core.NewQVariant1(qb.BlockNumber())
		}
	case Transactions:
		{
			return core.NewQVariant1(qb.Transactions())
		}
	case Blockhash:
		{
			return core.NewQVariant1(qb.Blockhash())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *BlocksModel) addBlocks(qb []*QBlock) {
	m.SetBlocks(qb)
	m.insertRows(len(m.Blocks()), len(qb))
}

func (m *BlocksModel) loadModel() {
	//m.SetOutputs(mo)
}

