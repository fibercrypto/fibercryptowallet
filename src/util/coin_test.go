package util

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
)

func TestNewGenericOutput(t *testing.T) {
	addr := &mocks.Address{}

	output := NewGenericOutput(addr, "id")
	require.Equal(t, GenericOutput{Address: addr, id: "id", Balance: make(map[string]uint64)}, output)

	output.SetCoins("some_coin", uint64(20))
	some_coin, err := output.GetCoins("some_coin")
	require.Equal(t, uint64(20), some_coin)
	require.NoError(t, err)

	require.Equal(t, addr, output.GetAddress())
	require.Equal(t, "id", output.GetId())
	require.False(t, output.IsSpent())

	assets := output.SupportedAssets()
	require.Equal(t, 1, len(assets))
	find := false
	for _, asset := range assets {
		if asset == "some_coin" {
			find = true
			break
		}
	}
	require.True(t, find)

	_, err2 := output.GetCoins("invalid_coin")
	require.Error(t, err2)

	//TODO: Find how to test PushCoin. There are problems registering coins

}
