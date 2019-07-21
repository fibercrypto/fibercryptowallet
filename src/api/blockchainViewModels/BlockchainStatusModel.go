package blockchainViewModels

import (
	"github.com/skycoin/skycoin/src/api"
	"github.com/therecipe/qt/core"
)

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

func init() {
	BlockchainStatusModel_QmlRegisterType2("BlockchainModels", 1, 0, "BlockchainStatusModel")
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
		return
	}
	return
}

//NewClient returns a new client
func NewClient() *api.Client {
	addr := "http://127.0.0.1:38391" //
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

	lastBlock := blocks.Blocks[0]
	numberOfBlocks := 1 // TODO: number of blocks?
	lastBlockHash := lastBlock.Head.Hash
	// timeStampLastBlock := lastBlock.Head.Time //TODO: how to extract the time from it

	year, month, day := 2000, 6, 25
	h, m, _ := 12, 12, 0

	// coinSup, err := c.CoinSupply()
	if err != nil {
		return err
	}

	currentSkySupply := 123      //(*coinSup).CurrentSupply // TODO: Parse int
	totalSkySupply := 211        //(*coinSup).TotalSupply
	currentCoinHoursSupply := 12 //(*coinSup).CurrentCoinHourSupply
	totalCoinHoursSupply := 122  //(*coinSup).TotalCoinHourSupply

	// block details
	println("updated-status")
	bs.SetNumberOfBlocks(numberOfBlocks)
	bs.SetTimestampLastBlock(core.NewQDateTime3(core.NewQDate3(year, month, day), core.NewQTime3(h, m, 0, 0), core.Qt__LocalTime)) //TODO: datetime
	bs.SetHashLastBlock(lastBlockHash)

	// sky details
	bs.SetCurrentSkySupply(currentSkySupply)
	bs.SetTotalSkySupply(totalSkySupply)
	bs.SetCurrentCoinHoursSupply(currentCoinHoursSupply)
	bs.SetTotalCoinHoursSupply(totalCoinHoursSupply)

	return nil
}
