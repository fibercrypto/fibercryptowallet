package skycoin

import (
	"errors"
	"time"

	"github.com/skycoin/skycoin/src/readable"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type SkycoinBlock struct { //implements core.Block interface
	Block *readable.Block
}

func (sb *SkycoinBlock) GetHash() ([]byte, error) {
	if sb.Block == nil {
		return nil, errors.New("Block not setted or nil")
	}
	return []byte(sb.Block.Head.Hash), nil
}

func (sb *SkycoinBlock) GetPrevHash() ([]byte, error) {
	if sb.Block == nil {
		return nil, errors.New("Block not setted or nil")
	}
	return []byte(sb.Block.Head.PreviousHash), nil
}

func (sb *SkycoinBlock) GetVersion() (uint32, error) {
	if sb.Block == nil {
		return 0, errors.New("Block not setted or nil")
	}
	return sb.Block.Head.Version, nil
}

func (sb *SkycoinBlock) GetTime() (core.Timestamp, error) {
	if sb.Block == nil {
		return 0, errors.New("Block not setted or nil")
	}
	return core.Timestamp(sb.Block.Head.Time), nil
}
func (sb *SkycoinBlock) GetHeight() (uint64, error) {
	if sb.Block == nil {
		return 0, errors.New("Block not setted or nil")
	}
	return 0, nil //TODO ???
}

func (sb *SkycoinBlock) GetFee(ticker string) (uint64, error) {
	if sb.Block == nil {
		return 0, errors.New("Block not setted or nil")
	}
	if ticker == CoinHour {
		return sb.Block.Head.Fee, nil // FIXME
	}
	return 0, nil
}

func (sb *SkycoinBlock) IsGenesisBlock() (bool, error) {
	if sb.Block == nil {
		return true, errors.New("Block not setted or nil")
	}
	return true, nil //TODO ???
}

type SkycoinBlockchainInfo struct {
	LastBlockInfo         *SkycoinBlock
	CurrentSkySupply      uint64
	TotalSkySupply        uint64
	CurrentCoinHourSupply uint64
	TotalCoinHourSupply   uint64
	NumberOfBlocks        *readable.BlockchainProgress
}

type SkycoinBlockchainStatus struct { //Implements BlockchainStatus interface
	lastTimeStatusRequested uint64
	lastTimeSupplyRequested uint64
	CacheTime               uint64
	cachedStatus            *SkycoinBlockchainInfo
}

func NewSkycoinBlockchainStatus(invalidCacheTime uint64) *SkycoinBlockchainStatus {
	return &SkycoinBlockchainStatus{CacheTime: invalidCacheTime}
}
func (ss SkycoinBlockchainStatus) GetCoinValue(coinvalue core.CoinValueKey, ticker string) (uint64, error) {
	elapsed := uint64(time.Now().UTC().UnixNano()) - ss.lastTimeSupplyRequested
	if elapsed > ss.CacheTime || ss.cachedStatus == nil {
		if ss.cachedStatus == nil {
			ss.cachedStatus = new(SkycoinBlockchainInfo)
		}
		if err := ss.requestSupplyInfo(); err != nil {
			return 0, err
		}
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
		return 0, errorTickerInvalid{} //TODO: Customize error
	}
}

func (ss SkycoinBlockchainStatus) GetLastBlock() (core.Block, error) {
	elapsed := uint64(time.Now().UTC().UnixNano()) - ss.lastTimeSupplyRequested
	if elapsed > ss.CacheTime || ss.cachedStatus == nil {
		if ss.cachedStatus == nil {
			ss.cachedStatus = new(SkycoinBlockchainInfo)
		}
		if err := ss.requestStatusInfo(); err != nil {
			return nil, err
		}
	}
	return ss.cachedStatus.LastBlockInfo, nil
}

func (ss SkycoinBlockchainStatus) GetNumberOfBlocks() (uint64, error) {
	if ss.cachedStatus == nil {
		if ss.cachedStatus == nil {
			ss.cachedStatus = new(SkycoinBlockchainInfo)
		}
		if err := ss.requestStatusInfo(); err != nil {
			return 0, errors.New("Number of Blocks unset ")
		}
	}

	return ss.cachedStatus.NumberOfBlocks.Current, nil
}

func (ss SkycoinBlockchainStatus) SetCacheTime(time uint64) {
	ss.CacheTime = time
}

func (ss SkycoinBlockchainStatus) requestSupplyInfo() error {

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	coinSupply, err := c.CoinSupply()
	if err != nil {
		return err
	}

	ss.cachedStatus.CurrentCoinHourSupply, err = util.GetCoinValue(coinSupply.CurrentCoinHourSupply)
	if err != nil {
		return err
	}

	ss.cachedStatus.TotalCoinHourSupply, err = util.GetCoinValue(coinSupply.TotalCoinHourSupply)
	if err != nil {
		return err
	}

	ss.cachedStatus.CurrentSkySupply, err = util.GetCoinValue(coinSupply.CurrentSupply)
	if err != nil {
		return err
	}

	ss.cachedStatus.TotalSkySupply, err = util.GetCoinValue(coinSupply.TotalSupply)
	if err != nil {
		return err
	}

	ss.cachedStatus.NumberOfBlocks, err = c.BlockchainProgress()
	if err != nil {
		return err
	}

	return nil
}

func (ss SkycoinBlockchainStatus) requestStatusInfo() error {

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	blocks, err := c.LastBlocks(1)

	if err != nil {
		return err
	}
	lastBlock := blocks.Blocks[len(blocks.Blocks)-1]
	ss.cachedStatus.LastBlockInfo = &SkycoinBlock{Block: &lastBlock}

	progress, err := c.BlockchainProgress()
	if err != nil {
		return err
	}
	ss.cachedStatus.NumberOfBlocks = &readable.BlockchainProgress{
		Current: progress.Current,
		Highest: progress.Highest,
		Peers:   progress.Peers,
	}

	return nil
}
