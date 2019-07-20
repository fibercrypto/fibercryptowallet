package blockchainViewModels

import "github.com/therecipe/qt/core"

// BlockchainStatusModel Contains info about the blockchain to be show.
type BlockchainStatusModel struct {
	core.QObject
	_ func() `contructor:"init"`

	_ int64  `property:"numberOfBlocks"`
	_ string `property:"timestampLastBlock"`
	_ string `property:"hashLastBlock"`
	_ int64  `property:"currentSkySupply"`
	_ int64  `property:"totalSkySupply"`
	_ int64  `property:"currentCoinHoursSupply"`
	_ int64  `property:"totalCoinHoursSupply"`
}

func (bs *BlockchainStatusModel) init() {
	println("init-BlockchainStatusModel")

	// block details
	bs.SetNumberOfBlocksDefault(1)
	bs.SetTimestampLastBlockDefault("2000-01-01 00:00")
	bs.SetHashLastBlockDefault("")
	// sky details
	bs.SetCurrentSkySupplyDefault(0)
	bs.SetTotalSkySupplyDefault(0)
	bs.SetCurrentCoinHoursSupplyDefault(0)
	bs.SetTotalCoinHoursSupplyDefault(0)

	// update info
	bs.updateInfo()
}

// updateInfo request the needed information
func (bs *BlockchainStatusModel) updateInfo() {
	// TODO: api work

	// block details
	println("backRequested")
	bs.SetNumberOfBlocks(2)
	bs.SetTimestampLastBlock("2000-01-01 00:00")
	bs.SetHashLastBlock("")

	// sky details
	bs.SetCurrentSkySupply(0)
	bs.SetTotalSkySupply(0)
	bs.SetCurrentCoinHoursSupply(0)
	bs.SetTotalCoinHoursSupply(0)
}
