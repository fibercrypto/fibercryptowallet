package models

import (
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"strconv"

	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models" //callable as skycoin
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"

	qtcore "github.com/therecipe/qt/core"
)

var logBlockchain = logging.MustGetLogger("modelsBlockchain")

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

func (blockchainStatus *BlockchainStatusModel) init() {
	// block details
	blockchainStatus.SetNumberOfBlocksDefault("0")
	blockchainStatus.SetTimestampLastBlockDefault(qtcore.NewQDateTime3(qtcore.NewQDate3(2000, 1, 1), qtcore.NewQTime3(0, 0, 0, 0), qtcore.Qt__LocalTime))
	blockchainStatus.SetHashLastBlockDefault("")
	// sky details
	blockchainStatus.SetCurrentSkySupplyDefault("0")
	blockchainStatus.SetTotalSkySupplyDefault("0")
	blockchainStatus.SetCurrentCoinHoursSupplyDefault("0")
	blockchainStatus.SetTotalCoinHoursSupplyDefault("0")
	blockchainStatus.SetLoading(true)
	blockchainStatus.infoRequester = skycoin.NewSkycoinBlockchainStatus(1000000) //FIXME: set correct value
}

func (blockchainStatus *BlockchainStatusModel) update() {
	// update info
	if err := blockchainStatus.updateInfo(); err != nil {
		logBlockchain.WithError(err).Warn("Couldn't update blockchain Info")
		return
	}
	return
}

// updateInfo request the needed information
func (blockchainStatus *BlockchainStatusModel) updateInfo() error {
	logBlockchain.Info("Updating Blockchain Status")
	blockchainStatus.SetLoading(true)

	block, err := blockchainStatus.infoRequester.GetLastBlock()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get last block")
		return err
	}

	lastBlockHash, err := block.GetHash()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get the hash of the last block")
		return err
	}
	numberOfBlocks, err := blockchainStatus.infoRequester.GetNumberOfBlocks()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get the number of blocks")
		return err
	}
	timestamp, err := block.GetTime()
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get block time")
		return err
	}
	year, month, day, h, m, s := util.ParseDate(int64(timestamp))

	currentSkySupply, err := blockchainStatus.infoRequester.GetCoinValue(core.CoinCurrentSupply, skycoin.Sky)
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get current coin supply of Skycoins")
		return err
	}
	totalSkySupply, err := blockchainStatus.infoRequester.GetCoinValue(core.CoinTotalSupply, skycoin.Sky)
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get total coin supply of Skycoins")
		return err
	}
	currentCoinHoursSupply, err := blockchainStatus.infoRequester.GetCoinValue(core.CoinCurrentSupply, skycoin.CoinHour)
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get current coin supply of Coin Hours")
		return err
	}
	totalCoinHoursSupply, err := blockchainStatus.infoRequester.GetCoinValue(core.CoinTotalSupply, skycoin.CoinHour)
	if err != nil {
		logBlockchain.WithError(err).Warn("Couldn't get total coin supply of coin hours")
		return err
	}

	// block details
	blockchainStatus.SetNumberOfBlocks(strconv.FormatUint(numberOfBlocks, 10))
	blockchainStatus.SetTimestampLastBlock(qtcore.NewQDateTime3(qtcore.NewQDate3(year, month, day), qtcore.NewQTime3(h, m, s, 0), qtcore.Qt__LocalTime))
	blockchainStatus.SetHashLastBlock(string(lastBlockHash))

	// sky details
	blockchainStatus.SetCurrentSkySupply(strconv.FormatUint(currentSkySupply, 10))
	blockchainStatus.SetTotalSkySupply(strconv.FormatUint(totalSkySupply, 10))
	blockchainStatus.SetCurrentCoinHoursSupply(strconv.FormatUint(currentCoinHoursSupply, 10))
	blockchainStatus.SetTotalCoinHoursSupply(strconv.FormatUint(totalCoinHoursSupply, 10))
	blockchainStatus.SetLoading(false)

	return nil
}
