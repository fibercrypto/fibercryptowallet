package explorer

import (
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/models"
	"github.com/fibercrypto/fibercryptowallet/src/params"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/skycoin/skycoin/src/util/droplet"
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
	TransactionLen
	Blockhash
	PrevBlockhash
	Size
	TotalAmount
	TransactionList
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
	_ *QBlock                    `property:"blockDetail"`
	_ func(pageNum int)          `signal:"loadPage"`
	_ func()                     `signal:"update,auto"`
	_ func(hash string)          `signal:"loadBlockByHash"`

	_ func([]*QBlock) `slot:"addBlocks"`
	_ func()          `slot:"loadModel"`
}

// QBlock Contains info about the block to be show.
type QBlock struct {
	qtcore.QObject

	_ []*models.QTransaction `property:"transactionList"`
	_ qtcore.QDateTime       `property:"time"`
	_ string                 `property:"blockhash"`
	_ string                 `property:"prevBlockhash"`
	_ string                 `property:"totalAmount"`
	_ uint64                 `property:"blockNumber"`
	_ uint64                 `property:"size"`
	_ int                    `property:"transactionLen"`
}

func (m *BlocksModel) init() {
	logExplorer.Info("Init Explorer Model")
	m.SetRoles(map[int]*qtcore.QByteArray{
		Time:            qtcore.NewQByteArray2("time", -1),
		BlockNumber:     qtcore.NewQByteArray2("blockNumber", -1),
		TransactionLen:  qtcore.NewQByteArray2("transactionLen", -1),
		Blockhash:       qtcore.NewQByteArray2("blockhash", -1),
		PrevBlockhash:   qtcore.NewQByteArray2("prevBlockhash", -1),
		Size:            qtcore.NewQByteArray2("size", -1),
		TotalAmount:     qtcore.NewQByteArray2("totalAmount", -1),
		TransactionList: qtcore.NewQByteArray2("transactionList", -1),
	})
	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectUpdate(m.update)
	m.ConnectLoadPage(m.loadPage)
	m.ConnectLoadBlockByHash(m.loadBlockByHash)
	m.blockChain = skycoin.NewSkycoinBlockchain(params.DataRefreshTimeout)
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
	logExplorer.Info("Load page ", pageNum)
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
	case TransactionLen:
		{
			return qtcore.NewQVariant1(qb.TransactionLen())
		}
	case Blockhash:
		{
			return qtcore.NewQVariant1(qb.Blockhash())
		}
	case PrevBlockhash:
		{
			return qtcore.NewQVariant1(qb.PrevBlockhash())
		}
	case Size:
		{
			return qtcore.NewQVariant1(qb.Size())
		}
	case TotalAmount:
		{
			return qtcore.NewQVariant1(qb.TotalAmount())
		}
	case TransactionList:
		{
			return qtcore.NewQVariant1(qb.TransactionList())
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

func (m *BlocksModel) loadBlockByHash(hash string) {
	logExplorer.Info("Loading block details")
	logExplorer.Info(hash)
	block, err := m.blockChain.GetBlockByHash(hash)
	if err != nil {
		logExplorer.WithError(err).Errorf("Couldn't get the detail of block with hash %s", hash)
	}

	m.SetBlockDetail(blockToQBlock(block))
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

	qBlock.SetBlockhash(string(hash))

	height, err := block.GetHeight()
	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the height of block")
	}

	qBlock.SetBlockNumber(height)

	prevHash, err := block.GetPrevHash()
	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the prevBlockHash of block")
	}

	qBlock.SetPrevBlockhash(string(prevHash))

	blockSize, err := block.GetSize()
	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the size of block")
	}

	qBlock.SetSize(blockSize)

	totalAmount, err := block.GetTotalAmount()
	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the total amount of block")
	}

	totalAmountStr, err := droplet.ToString(totalAmount)
	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the total amount of block")
	}
	qBlock.SetTotalAmount(totalAmountStr)

	transactionList, err := block.GetTransactions()
	if err != nil {
		logExplorer.WithError(err).Error("Couldn't get the transactions list of block")
	}
	var qTransactionsList []*models.QTransaction
	for e := range transactionList {

		qTransaction, err := models.NewQTransactionFromTransaction(transactionList[e])
		if err != nil {
			logExplorer.WithError(err).Error("Couldn't get the transactions list of block")
		}
		qTransactionsList = append(qTransactionsList, qTransaction)
	}
	logExplorer.Info("Transaction Count: ", len(qTransactionsList))
	qBlock.SetTransactionList(qTransactionsList)
	qBlock.SetTransactionLen(len(qTransactionsList))

	return qBlock
}
