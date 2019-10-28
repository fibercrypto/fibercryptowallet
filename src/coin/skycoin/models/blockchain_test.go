package skycoin

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

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

	block := &SkycoinBlockchainStatus{CacheTime: 20}
	val, err := block.GetCoinValue(core.CoinCurrentSupply, Sky)
	require.Nil(t, err)
	require.Equal(t, val, uint64(200111111))
	val, err = block.GetCoinValue(core.CoinCurrentSupply, CoinHour)
	require.Nil(t, err)
	require.Equal(t, val, uint64(200))
	val, err = block.GetCoinValue(core.CoinTotalSupply, Sky)
	require.Nil(t, err)
	require.Equal(t, val, uint64(300000000))
	val, err = block.GetCoinValue(core.CoinTotalSupply, CoinHour)
	require.Nil(t, err)
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

	block := &SkycoinBlockchainStatus{CacheTime: 20}
	val, err := block.GetNumberOfBlocks()
	require.Nil(t, err)
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

	status := &SkycoinBlockchainStatus{CacheTime: 20}
	block, err := status.GetLastBlock()
	require.Nil(t, err)
	val, err2 := block.GetVersion()
	require.Nil(t, err2)
	require.Equal(t, val, uint32(3))
}