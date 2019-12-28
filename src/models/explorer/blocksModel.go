package explorer

import (
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/params"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtcore "github.com/therecipe/qt/core"
)

// Properties:

// For blocks:
// - Time
// - Block number
// - Transactions (number of Transactions)
// - BlockHash

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
	BlockHash
)

// BlocksModel List of Blocks to be show.
type BlocksModel struct {
	qtcore.QAbstractListModel
	blockChain core.BlockchainStatus
	_          func() `constructor:"init"`

	_ int                        `property:"currentPage"`
	_ int                        `property:"countPage"`
	_ int                        `property:"blockNumber"`
	_ map[int]*qtcore.QByteArray `property:"roles"`
	_ []*QBlock                  `property:"blocks"`

	_ func(pageNum int) `signal:"loadPage"`
	_ func()            `signal:"update,auto"`

	_ func([]*QBlock) `slot:"addBlocks"`
	_ func()          `slot:"loadModel"`
}

// QBlock Contains info about the block to be show.
type QBlock struct {
	qtcore.QObject

	_ qtcore.QDateTime `property:"time"`
	_ uint64           `property:"blockNumber"`
	_ int              `property:"transactions"`
	_ string           `property:"blockHash"`
}

func (m *BlocksModel) init() {
	logExplorer.Info("Init Explorer Model")
	m.SetRoles(map[int]*qtcore.QByteArray{
		Time:         qtcore.NewQByteArray2("time", -1),
		BlockNumber:  qtcore.NewQByteArray2("blockNumber", -1),
		Transactions: qtcore.NewQByteArray2("transactions", -1),
		BlockHash:    qtcore.NewQByteArray2("blockHash", -1),
	})
	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddBlocks(m.addBlocks)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectUpdate(m.update)
	m.ConnectLoadPage(m.loadPage)
	m.blockChain = skycoin.NewSkycoinBlockchain(params.DataRefreshTimeout)
	m.update()
	m.loadPage(1)
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
	numberOfBlocks, err := blocksM.blockChain.GetNumberOfBlocks()
	if err != nil {
		logExplorer.WithError(err).Warn("Couldn't get the number of blocks")
		return err
	}
	blocksM.SetCurrentPage(1)
	blocksM.SetBlockNumber(int(numberOfBlocks))

	if numberOfBlocks%10 != 0 {
		blocksM.SetCountPage(int(numberOfBlocks/10) + 1)
	} else {
		blocksM.SetCountPage(int(numberOfBlocks / 10))
	}
	return nil
}

func (blocksM *BlocksModel) loadPage(pageNum int) {
	if pageNum > 0 && pageNum <= blocksM.CountPage() {
		blocksM.SetCurrentPage(pageNum)
		blocksM.loadModel()
	}
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
	case BlockHash:
		{
			return qtcore.NewQVariant1(qb.BlockHash())
		}
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *BlocksModel) addBlocks(qb []*QBlock) {
	logExplorer.Info("Add Block")
	m.BeginResetModel()
	m.SetBlocks(qb)
	m.EndResetModel()

}

func (m *BlocksModel) loadModel() {
	logExplorer.Info("Loading Explorer")

	blocks, err := m.blockChain.GetRangeBlocks(uint64(m.BlockNumber()-(m.CurrentPage()*10))+1,
		uint64((m.BlockNumber()-(m.CurrentPage()*10))+10))
	logExplorer.Info(len(blocks))
	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the range of blocks")
	}
	logExplorer.Info(len(m.Blocks()))
	var qBlocks []*QBlock
	for e := range blocks {
		h, _ := blocks[e].GetHeight()
		logExplorer.Info(h)
		qBlocks = append(qBlocks, blockToQBlock(blocks[e]))
	}
	m.addBlocks(qBlocks)
}

func blockToQBlock(block core.Block) *QBlock {
	var qBlock = NewQBlock(nil)
	timestamp, err := block.GetTime()

	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the time of block")
	}

	year, month, day, h, m, s := util.ParseDate(int64(timestamp))
	qBlock.SetTime(qtcore.NewQDateTime3(qtcore.NewQDate3(year, month, day),
		qtcore.NewQTime3(h, m, s, 0), qtcore.Qt__LocalTime))

	hash, err := block.GetHash()

	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the hash of block")
	}

	qBlock.SetBlockHash(string(hash))

	height, err := block.GetHeight()

	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the height of block")
	}

	qBlock.SetBlockNumber(height)

	transactions, err := block.GetTransactions()

	if err != nil {
		logExplorer.Error(err)
	}

	qBlock.SetTransactions(len(transactions))

	return qBlock
}
