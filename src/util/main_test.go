package util

import (
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPubKeyFromBytes(t *testing.T) {
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
	bytes := []byte{1, 0, 1}
	mockPubKey := new(mocks.PubKey)
	mockPlugin.On("PubKeyFromBytes", bytes).Return(mockPubKey, nil)
	RegisterAltcoin(mockPlugin)

	pubKey, err := PubKeyFromBytes(fakeTicker, bytes)
	require.Nil(t, err)
	require.Equal(t, pubKey, mockPubKey)

	pubKey, err = PubKeyFromBytes(`CUSTOMTICKER`, bytes)
	require.NotNil(t, err)
}

func TestSecKeyFromBytes(t *testing.T) {
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
	bytes := []byte{1, 0, 1}
	mockSecKey := new(mocks.SecKey)
	mockPlugin.On("SecKeyFromBytes", bytes).Return(mockSecKey, nil)
	RegisterAltcoin(mockPlugin)

	pubKey, err := SecKeyFromBytes(fakeTicker, bytes)
	require.Nil(t, err)
	require.Equal(t, pubKey, mockSecKey)

	pubKey, err = SecKeyFromBytes(`CUSTOMTICKER`, bytes)
	require.NotNil(t, err)
}
