package models

import (
	"strconv"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models" //callable as skycoin
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"

	qtcore "github.com/therecipe/qt/core"
)

// BlockchainStatusModel Contains info about the blockchain to be show.
type BlockchainStatusModel struct {
	qtcore.QObject
	_ func() `constructor:"init"`

	infoRequester core.BlockchainStatus

	_ string            `property:"numberOfBlocks"`
	_ bool              `property:"loading"`
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
	bs.SetLoading(true)
	bs.infoRequester = skycoin.NewSkycoinBlockchainStatus(1000000) //FIXME: set correct value
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
	bs.SetLoading(true)

	block, err := bs.infoRequester.GetLastBlock()
	if err != nil {
		return err
	}

	lastBlockHash, err := block.GetHash()
	if err != nil {
		return err
	}
	numberOfBlocks, err := bs.infoRequester.GetNumberOfBlocks()
	if err != nil {
		return err
	}
	tstamp, err := block.GetTime()
	if err != nil {
		return err
	}
	year, month, day, h, m, s := util.ParseDate(int64(tstamp))

	currentSkySupply, err := bs.infoRequester.GetCoinValue(core.CoinCurrentSupply, skycoin.Sky)
	if err != nil {
		return err
	}
	totalSkySupply, err := bs.infoRequester.GetCoinValue(core.CoinTotalSupply, skycoin.Sky)
	if err != nil {
		return err
	}
	currentCoinHoursSupply, err := bs.infoRequester.GetCoinValue(core.CoinCurrentSupply, skycoin.CoinHour)
	if err != nil {
		return err
	}
	totalCoinHoursSupply, err := bs.infoRequester.GetCoinValue(core.CoinTotalSupply, skycoin.CoinHour)
	if err != nil {
		return err
	}

	// block details
	bs.SetNumberOfBlocks(strconv.FormatUint(numberOfBlocks, 10))
	bs.SetTimestampLastBlock(qtcore.NewQDateTime3(qtcore.NewQDate3(year, month, day), qtcore.NewQTime3(h, m, s, 0), qtcore.Qt__LocalTime))
	bs.SetHashLastBlock(string(lastBlockHash))

	// sky details
	bs.SetCurrentSkySupply(strconv.FormatUint(currentSkySupply, 10))
	bs.SetTotalSkySupply(strconv.FormatUint(totalSkySupply, 10))
	bs.SetCurrentCoinHoursSupply(strconv.FormatUint(currentCoinHoursSupply, 10))
	bs.SetTotalCoinHoursSupply(strconv.FormatUint(totalCoinHoursSupply, 10))
	bs.SetLoading(false)

	return nil
}
