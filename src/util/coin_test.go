package util

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

func TestNewGenericOutput(t *testing.T) {
	fakeTicker := "MOCKSCOIN"
	fakeDesc := "Fake coin"
	fakeExp := 3
	fakeMeta := core.AltcoinMetadata{
		Name:          fakeDesc,
		Ticker:        fakeTicker,
		Family:        fakeTicker,
		HasBip44:      false,
		Bip44CoinType: 0,
		Accuracy:      int32(fakeExp),
	}
	mockPlugin := new(mocks.AltcoinPlugin)
	mockPlugin.On("RegisterTo", mock.Anything).Return().Run(func(args mock.Arguments) {
		manager := args.Get(0).(core.AltcoinManager)
		manager.RegisterAltcoin(fakeMeta, mockPlugin)
	})
	mockPlugin.On("GetName").Return(fakeDesc)
	RegisterAltcoin(mockPlugin)

	addr := new(mocks.Address)

	notRegisteredCoin := "some_coin"
	output := NewGenericOutput(addr, "id")
	require.Equal(t, GenericOutput{Address: addr, id: "id", Balance: make(map[string]uint64)}, output)

	pushError := output.PushCoins(fakeTicker, "42")
	require.Nil(t, pushError)
	pushError = output.PushCoins(notRegisteredCoin, "42")
	require.NotNil(t, pushError)

	output.SetCoins(notRegisteredCoin, uint64(20))
	amount, err := output.GetCoins(notRegisteredCoin)
	require.Equal(t, uint64(20), amount)
	require.Nil(t, err)
	amount, err = output.GetCoins(fakeTicker)
	require.Equal(t, uint64(42000), amount)
	require.Nil(t, err)

	require.Equal(t, addr, output.GetAddress())
	require.Equal(t, "id", output.GetId())
	require.False(t, output.IsSpent())

	testAssets := []string{fakeTicker, notRegisteredCoin}
	sort.Strings(testAssets)
	assets := output.SupportedAssets()
	sort.Strings(assets)
	require.Equal(t, testAssets, assets)

	_, err2 := output.GetCoins("other_coin")
	require.Error(t, err2)
}
