package skycoin

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

func TestSkycoinBlockStructure(t *testing.T) {
	rBlock := &readable.Block{}
	skyBlock := &SkycoinBlock{}

	_, err := skyBlock.GetHash()
	require.Error(t, err)
	_, err = skyBlock.GetPrevHash()
	require.Error(t, err)
	_, err = skyBlock.GetVersion()
	require.Error(t, err)
	_, err = skyBlock.GetTime()
	require.Error(t, err)
	_, err = skyBlock.GetHeight()
	require.Error(t, err)
	_, err = skyBlock.GetFee("coin")
	require.Error(t, err)
	_, err = skyBlock.IsGenesisBlock()
	require.Error(t, err)

	skyfee := uint64(0)
	version := uint32(42)
	time, chfee := uint64(100), uint64(10000)
	hash, prev := "block_hash", "block_prev_hash"
	rBlock.Head.Time = time
	rBlock.Head.Fee = chfee
	rBlock.Head.Hash = hash
	rBlock.Head.Version = version
	rBlock.Head.PreviousHash = prev

	skyBlock.Block = rBlock
	bhash, err1 := skyBlock.GetHash()
	require.NoError(t, err1)
	require.Equal(t, []byte(hash), bhash)
	bprev, err2 := skyBlock.GetPrevHash()
	require.NoError(t, err2)
	require.Equal(t, []byte(prev), bprev)
	bversion, err3 := skyBlock.GetVersion()
	require.NoError(t, err3)
	require.Equal(t, version, bversion)
	btime, err4 := skyBlock.GetTime()
	require.NoError(t, err4)
	require.Equal(t, core.Timestamp(time), btime)
	bheigth, err5 := skyBlock.GetHeight()
	require.NoError(t, err5)
	//TODO: When the behavior of the GetHeight function is decided, the test must be modified accordingly
	require.Equal(t, uint64(0), bheigth)
	bfee, err6 := skyBlock.GetFee(CoinHour)
	require.NoError(t, err6)
	require.Equal(t, chfee, bfee)
	bfee, err6 = skyBlock.GetFee(Sky)
	require.NoError(t, err6)
	require.Equal(t, skyfee, bfee)
	genesis, err7 := skyBlock.IsGenesisBlock()
	require.NoError(t, err7)
	require.True(t, genesis)
}

func TestSkycoinBlockchainStatusGetCoinValue(t *testing.T) {
	global_mock.On("CoinSupply").Return(
		&api.CoinSupply{
			CurrentCoinHourSupply: "200",
			TotalCoinHourSupply:   "300",
			CurrentSupply:         "200.111111",
			TotalSupply:           "300",
		},
		nil,
	)
	global_mock.On("BlockchainProgress").Return(&readable.BlockchainProgress{}, nil)

	block := &SkycoinBlockchain{CacheTime: 20}
	val, err := block.GetCoinValue(core.CoinCurrentSupply, Sky)
	require.NoError(t, err)
	require.Equal(t, val, uint64(200111111))
	val, err = block.GetCoinValue(core.CoinCurrentSupply, CoinHour)
	require.NoError(t, err)
	require.Equal(t, val, uint64(200))
	val, err = block.GetCoinValue(core.CoinTotalSupply, Sky)
	require.NoError(t, err)
	require.Equal(t, val, uint64(300000000))
	val, err = block.GetCoinValue(core.CoinTotalSupply, CoinHour)
	require.NoError(t, err)
	require.Equal(t, val, uint64(300))
}

func TestSkycoinBlockchainStatusGetNumberOfBlocks(t *testing.T) {
	CleanGlobalMock()
	global_mock.On("LastBlocks", uint64(1)).Return(
		&readable.Blocks{
			Blocks: []readable.Block{
				readable.Block{},
			},
		},
		nil,
	)
	global_mock.On("BlockchainProgress").Return(
		&readable.BlockchainProgress{
			Current: uint64(20),
			Highest: uint64(42),
			Peers:   []readable.PeerBlockchainHeight{},
		},
		nil,
	)

	block := &SkycoinBlockchain{CacheTime: 20}
	val, err := block.GetNumberOfBlocks()
	require.NoError(t, err)
	require.Equal(t, val, uint64(20))
}

func TestSkycoinBlockchainStatusGetLastBlock(t *testing.T) {
	CleanGlobalMock()
	global_mock.On("LastBlocks", uint64(1)).Return(
		&readable.Blocks{
			Blocks: []readable.Block{
				readable.Block{
					Head: readable.BlockHeader{Version: uint32(3)},
				},
			},
		},
		nil,
	)
	global_mock.On("BlockchainProgress").Return(&readable.BlockchainProgress{}, nil)

	status := &SkycoinBlockchain{CacheTime: 20}
	block, err := status.GetLastBlock()
	require.NoError(t, err)
	val, err2 := block.GetVersion()
	require.NoError(t, err2)
	require.Equal(t, val, uint32(3))
}

func TestSkycoinBlockchainSetCacheTime(t *testing.T) {
	bchn := &SkycoinBlockchain{}
	time := uint64(345)
	bchn.SetCacheTime(time)
	require.Equal(t, time, bchn.CacheTime)
}
