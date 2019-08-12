package models

import (
	"strconv"

	skycoin "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"

	qtcore "github.com/therecipe/qt/core"
)

// BlockchainStatusModel Contains info about the blockchain to be show.
type BlockchainStatusModel struct {
	qtcore.QObject
	_ func() `constructor:"init"`

	infoRequester *core.BlockchainStatus

	_ string            `property:"numberOfBlocks"`
	_ *qtcore.QDateTime `property:"timestampLastBlock"`
	_ string            `property:"hashLastBlock"`
	_ string            `property:"currentSkySupply"`
	_ string            `property:"totalSkySupply"`
	_ string            `property:"currentCoinHoursSupply"`
	_ string            `property:"totalCoinHoursSupply"`

	_ func() `signal:"update,auto"`
}

func (bs *BlockchainStatusModel) init() {
	// block details
	bs.SetNumberOfBlocksDefault("0")
	bs.SetTimestampLastBlockDefault(qtcore.NewQDateTime3(qtcore.NewQDate3(2000, 1, 1), qtcore.NewQTime3(0, 0, 0, 0), qtcore.Qt__LocalTime))
	bs.SetHashLastBlockDefault("")
	// sky details
	bs.SetCurrentSkySupplyDefault("0")
	bs.SetTotalSkySupplyDefault("0")
	bs.SetCurrentCoinHoursSupplyDefault("0")
	bs.SetTotalCoinHoursSupplyDefault("0")
}

func (bs *BlockchainStatusModel) update() {
	// update info
	if err := bs.updateInfo(); err != nil {
		println(err.Error())
		return
	}
	return
}

// updateInfo request the needed information
func (bs *BlockchainStatusModel) updateInfo() error {
	c := util.NewClient()

	blocks, err := c.LastBlocks(1)
	if err != nil {
		return err
	}

	lastBlock := blocks.Blocks[len(blocks.Blocks)-1]
	numberOfBlocks := strconv.FormatUint(lastBlock.Head.BkSeq, 10)
	lastBlockHash := lastBlock.Head.Hash
	year, month, day, h, m, s := util.ParseDate(int64(lastBlock.Head.Time)) // Fixme: the conversion its save, right??

	coinSup, err := c.CoinSupply()
	if err != nil {
		return err
	}

	currentSkySupply := (*coinSup).CurrentSupply
	totalSkySupply := (*coinSup).TotalSupply
	currentCoinHoursSupply := (*coinSup).CurrentCoinHourSupply
	totalCoinHoursSupply := (*coinSup).TotalCoinHourSupply

	// block details
	bs.SetNumberOfBlocks(numberOfBlocks)
	bs.SetTimestampLastBlock(core.NewQDateTime3(core.NewQDate3(year, month, day), core.NewQTime3(h, m, s, 0), core.Qt__LocalTime))
	bs.SetHashLastBlock(lastBlockHash)

	// sky details
	bs.SetCurrentSkySupply(currentSkySupply)
	bs.SetTotalSkySupply(totalSkySupply)
	bs.SetCurrentCoinHoursSupply(currentCoinHoursSupply)
	bs.SetTotalCoinHoursSupply(totalCoinHoursSupply)

	return nil
}
