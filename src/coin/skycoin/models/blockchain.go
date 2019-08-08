package models

import (
	"time"

	"github.com/skycoin/skycoin/src/readable"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

const (
	Sky      = "Sky"
	CoinHour = "CoinHour"
)

type SkycoinBlock struct { //implements core.Block interface
	Block *readable.Block
}

type SkycoinBlockchianInfo struct {
	LastBlockInfo         *SkycoinBlock
	CurrentSkySupply      string
	TotalSkySupply        string
	CurrentCoinHourSupply string
	TotalCoinHourSupply   string
}

type SkycoinBlockchainStatus struct { //Implements BlockchainStatus interface
	lastTimeStatusRequested uint64
	lastTimeSupplyRequested uint64
	cacheTime               uint64
	cachedStatus            *SkycoinBlockchianInfo
}

func (ss *SkycoinBlockchainStatus) GetCoinValue(coinvalue core.CoinValueKey, ticker string) (string, error) {
	elapsed := uint64(time.Now().UTC().UnixNano()) - ss.lastTimeSupplyRequested
	if elapsed > ss.cacheTime {
		err := ss.requestSupplyInfo()
		return "", err
	}

	switch ticker {
	case Sky:
		if coinvalue == core.CoinCurrentSupply {
			return ss.cachedStatus.CurrentSkySupply, nil
		}
		return ss.cachedStatus.TotalSkySupply, nil
	case CoinHour:
		if coinvalue == core.CoinCurrentSupply {
			return ss.cachedStatus.CurrentCoinHourSupply, nil
		}
		return ss.cachedStatus.TotalCoinHourSupply, nil
	default:
		return "", errorTickerInvalid{} //TODO: Customize error
	}
}

func (ss *SkycoinBlockchainStatus) GetLastBlock() (core.Block, error) {
	elapsed := uint64(time.Now().UTC().UnixNano()) - ss.lastTimeSupplyRequested

	if elapsed > ss.cacheTime {
		err := ss.requestSupplyInfo()
		return nil, err
	}
	return nil, nil
	// FIXME: return *ss.cachedStatus.LastBlockInfo, nil
}

func (ss *SkycoinBlockchainStatus) SetCacheTime(time uint64) {
	ss.cacheTime = time
}

func (ss *SkycoinBlockchainStatus) requestSupplyInfo() error {
	c := util.NewClient()

	coinSupply, err := c.CoinSupply()
	if err != nil {
		return err
	}

	ss.cachedStatus.CurrentCoinHourSupply = coinSupply.CurrentCoinHourSupply
	ss.cachedStatus.TotalCoinHourSupply = coinSupply.TotalCoinHourSupply
	ss.cachedStatus.CurrentSkySupply = coinSupply.CurrentSupply
	ss.cachedStatus.TotalSkySupply = coinSupply.TotalSupply

	return nil
}

func (ss *SkycoinBlockchainStatus) requestStatusInfo() error {
	c := util.NewClient()

	blocks, err := c.LastBlocks(1)

	if err != nil {
		return err
	}
	lastBlock := blocks.Blocks[len(blocks.Blocks)-1]
	ss.cachedStatus.LastBlockInfo = &SkycoinBlock{Block: &lastBlock}

	// ss.cachedStatus.LastBlockInfo.HashLastBlock = lastBlock.Head.Hash
	// ss.cachedStatus.LastBlockInfo.NumberOfBlocks = lastBlock.Head.BkSeq
	// ss.cachedStatus.LastBlockInfo.TimeStamp = core.Timestamp(lastBlock.Head.Time)

	return nil
}
