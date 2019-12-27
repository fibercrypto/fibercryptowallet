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

	_ func() `constructor:"init"`

	_ int                        `property:"currentPage"`
	_ map[int]*qtcore.QByteArray `property:"roles"`
	_ []*QBlock                  `property:"blocks"`

	_ func(pageNum int) `signal:"loadPage"`
	_ func()            `signal:"update,auto"`

	_ func([]*QBlock) `slot:"addBlocks"`
	_ func() 		  `slot:"loadModel"`
}

// QBlock Contains info about the block to be show.
type QBlock struct {
	qtcore.QObject

	_ string `property:"time"`
	_ string `property:"blockNumber"`
	_ string `property:"Transactions"`
	_ string `property:"Blockhash"`
}

func (m *BlocksModel) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Time:         qtcore.NewQByteArray2("time", -1),
		BlockNumber:  qtcore.NewQByteArray2("blockNumber", -1),
		Transactions: qtcore.NewQByteArray2("transactions", -1),
		Blockhash:    qtcore.NewQByteArray2("blockhash", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddBlocks(m.addBlocks)
	m.ConnectLoadModel(m.loadModel)

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

func (m *BlocksModel) rowCount(*qtcore.QModelIndex) int {
	return len(m.Blocks())
}

func (m *BlocksModel) roleNames() map[int]*qtcore.QByteArray {
	return m.Roles()
}

func (m *BlocksModel) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	if !index.IsValid() || index.Row() >= len(m.Blocks()) {
		return qtcore.NewQVariant()
	}

	qb := m.Blocks()[index.Row()]

	switch role {
	case Time:
		{
			return qtcore.NewQVariant1(qb.Time())
		}
	case BlockNumber:
		{
			return qtcore.NewQVariant1(qb.BlockNumber())
		}
	case Transactions:
		{
			return qtcore.NewQVariant1(qb.Transactions())
		}
	case Blockhash:
		{
			return qtcore.NewQVariant1(qb.Blockhash())
		}
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *BlocksModel) insertRows(row int, count int) bool {
	m.BeginInsertRows(qtcore.NewQModelIndex(), row, row+count)
	m.EndInsertRows()
	return true
}

func (m *BlocksModel) addBlocks(qb []*QBlock) {
	m.SetBlocks(qb)
	m.insertRows(len(m.Blocks()), len(qb))
}

func (m *BlocksModel) loadModel() {
	//m.SetOutputs(mo)
}
