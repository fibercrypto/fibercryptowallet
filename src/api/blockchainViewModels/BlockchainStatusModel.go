package blockchainViewModels

import (
	"time"

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

// updateInfo request the needed information
func (bs *BlockchainStatusModel) updateInfo() error {
	// TODO: api work
	dt := time.Now()
	year, _, day := dt.Date() //month problem
	month := 6
	h, m, _ := dt.Clock()

	// block details
	println("updated-status")
	bs.SetNumberOfBlocks(2)
	bs.SetTimestampLastBlock(core.NewQDateTime3(core.NewQDate3(year, month, day), core.NewQTime3(h, m, 0, 0), core.Qt__LocalTime)) //TODO: datetime
	bs.SetHashLastBlock("12312312323123123123")

	// sky details
	bs.SetCurrentSkySupply(123)
	bs.SetTotalSkySupply(211)
	bs.SetCurrentCoinHoursSupply(12)
	bs.SetTotalCoinHoursSupply(123)

	return nil
}
