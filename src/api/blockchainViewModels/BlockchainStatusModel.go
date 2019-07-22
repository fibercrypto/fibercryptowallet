package blockchainViewModels

import (
	"time"

	"github.com/skycoin/skycoin/src/api"
	"github.com/therecipe/qt/core"
)

func init() {
	BlockchainStatusModel_QmlRegisterType2("BlockchainModels", 1, 0, "BlockchainStatusModel")
}

// BlockchainStatusModel Contains info about the blockchain to be show.
type BlockchainStatusModel struct {
	core.QObject
	_ func() `constructor:"init"`

	_ int             `property:"numberOfBlocks"`
	_ *core.QDateTime `property:"timestampLastBlock"`
	_ string          `property:"hashLastBlock"`
	_ int             `property:"currentSkySupply"`
	_ int             `property:"totalSkySupply"`
	_ int             `property:"currentCoinHoursSupply"`
	_ int             `property:"totalCoinHoursSupply"`
}

func (bs *BlockchainStatusModel) init() {
	println("init-BlockchainStatusModel")

	// block details
	bs.SetNumberOfBlocksDefault(0)
	bs.SetTimestampLastBlockDefault(core.NewQDateTime())
	bs.SetHashLastBlockDefault("")
	// sky details
	bs.SetCurrentSkySupplyDefault(0)
	bs.SetTotalSkySupplyDefault(0)
	bs.SetCurrentCoinHoursSupplyDefault(0)
	bs.SetTotalCoinHoursSupplyDefault(0)

	// update info
	if err := bs.updateInfo(); err != nil {
		println(err.Error())
		return
	}
	return
}

//NewClient returns a new client
func NewClient() *api.Client {
	addr := "http://127.0.0.1:38391" // example only
	return api.NewClient(addr)
}

// updateInfo request the needed information
func (bs *BlockchainStatusModel) updateInfo() error {
	// TODO: api work

	c := NewClient()

	blocks, err := c.LastBlocks(1)
	if err != nil {
		return err
	}

	lastBlock := blocks.Blocks[len(blocks.Blocks)-1]
	numberOfBlocks := 1 // TODO: number of blocks?
	lastBlockHash := lastBlock.Head.Hash
	// timeStampLastBlock := lastBlock.Head.Time //TODO: how to extract the time from it

	year, month, day := 2000, 6, 25
	h, m, _ := 12, 12, 0

	coinSup, err := c.CoinSupply()
	if err != nil {
		return err
	}

	// TODO: what i supose to use int or float to representate money values
	// TODO: resolve conflicts with int64 and qlonglong
	//
	_currentSkySupply, err := strconv.ParseInt((*coinSup).CurrentSupply, 10, 0)
	if err != nil {
		return err
	}
	currentSkySupply := int(_currentSkySupply)
	_totalSkySupply, err := strconv.ParseInt((*coinSup).TotalSupply, 10, 0)
	if err != nil {
		return err
	}
	totalSkySupply := int(_totalSkySupply)
	_currentCoinHoursSupply, err := strconv.ParseInt((*coinSup).CurrentCoinHourSupply, 10, 0)
	if err != nil {
		return err
	}
	currentCoinHoursSupply := int(_currentCoinHoursSupply)

	_totalCoinHoursSupply, err := strconv.ParseInt((*coinSup).TotalCoinHourSupply, 10, 0)
	if err != nil {
		return err
	}
	totalCoinHoursSupply := int(_totalCoinHoursSupply)

	// block details
	bs.SetNumberOfBlocks(numberOfBlocks)
	bs.SetTimestampLastBlock(core.NewQDateTime3(core.NewQDate3(year, month, day), core.NewQTime3(h, m, 0, 0), core.Qt__LocalTime)) //TODO: datetime
	bs.SetHashLastBlock(lastBlockHash)

	// sky details
	bs.SetCurrentSkySupply(currentSkySupply)
	bs.SetTotalSkySupply(totalSkySupply)
	bs.SetCurrentCoinHoursSupply(currentCoinHoursSupply)
	bs.SetTotalCoinHoursSupply(totalCoinHoursSupply)

	println("Status-updated")

	return nil
}

func parseDate(timeStamp int64) (int, int, int, int, int, int) {
	t := time.Unix(timeStamp, 0) //Fixme: use or not UTC() for local time or for server time?
	y, _m, d := t.Date()

	return y, int(_m), d, t.Hour(), t.Minute(), t.Second()
}
