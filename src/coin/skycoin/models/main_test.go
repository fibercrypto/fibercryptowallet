package skycoin

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/mock"
	"os"
	"sort"
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/params"
	util "github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Prepare the mock API for all test
func TestMain(m *testing.M) {
	if global_mock == nil {
		global_mock = new(SkycoinApiMock)
	}
	err := core.GetMultiPool().CreateSection(PoolSection, global_mock)
	if err != nil {
		logModelTest.WithError(err).Warn("Error creating pool section")
		return
	}
	util.RegisterAltcoin(NewSkyFiberPlugin(SkycoinMainNetParams))
	os.Exit(m.Run())
}

func TestSkyFiberPluginFunctions(t *testing.T) {
	name := "SkyFiber"
	family := []string{SkycoinFamily}
	description := "FiberCrypto wallet connector for Skycoin and SkyFiber altcoins"
	altcoins := []core.AltcoinMetadata{
		core.AltcoinMetadata{
			Name:     SkycoinName,
			Ticker:   SkycoinTicker,
			Family:   SkycoinFamily,
			HasBip44: false,
			Accuracy: 6,
		},
		core.AltcoinMetadata{
			Name:     CoinHoursName,
			Ticker:   CoinHoursTicker,
			Family:   SkycoinFamily,
			HasBip44: false,
			Accuracy: 0,
		},
		core.AltcoinMetadata{
			Name:     CalculatedHoursName,
			Ticker:   CalculatedHoursTicker,
			Family:   SkycoinFamily,
			HasBip44: false,
			Accuracy: 0,
		},
	}
	cmp := func(altcoinList []core.AltcoinMetadata) {
		sort.SliceStable(altcoinList, func(i, j int) bool {
			return altcoinList[i].Name < altcoinList[j].Name
		})
	}

	plugin := NewSkyFiberPlugin(SkycoinMainNetParams)
	require.NotNil(t, plugin)
	sPlugin := plugin.(*SkyFiberPlugin)
	require.NotNil(t, sPlugin)
	require.Equal(t, SkycoinMainNetParams, sPlugin.Params)
	require.Equal(t, name, plugin.GetName())
	require.Equal(t, description, plugin.GetDescription())

	pluginFamily := plugin.ListSupportedFamilies()
	sort.Strings(pluginFamily)
	sort.Strings(family)
	require.Equal(t, family, pluginFamily)

	pluginAltcoins := plugin.ListSupportedAltcoins()
	cmp(pluginAltcoins)
	cmp(altcoins)
	require.Equal(t, altcoins, pluginAltcoins)

	signer, err := plugin.LoadSignService()
	require.Nil(t, err)
	require.NotNil(t, signer)
	require.IsType(t, new(SkycoinSignService), signer)
}

func TestSkyFiberPluginRegisterTo(t *testing.T) {
	plugin := NewSkyFiberPlugin(SkycoinMainNetParams)
	altcoins := plugin.ListSupportedAltcoins()

	manager := new(mocks.AltcoinManager)
	manager.On("RegisterAltcoin", mock.AnythingOfType("core.AltcoinMetadata"), plugin).Return().Run(
		func(args mock.Arguments) {
			meta := args.Get(0).(core.AltcoinMetadata)
			idx := -1
			for i, altcoin := range altcoins {
				if altcoin == meta {
					idx = i
					break
				}
			}
			require.NotEqual(t, -1, idx)
			altcoins[idx] = altcoins[len(altcoins)-1]
			altcoins = altcoins[:len(altcoins)-1]
		},
	)
	plugin.RegisterTo(manager)
	require.Equal(t, 0, len(altcoins))
}

func TestSkyFiberPluginNetOperations(t *testing.T) {
	plugin := NewSkyFiberPlugin(SkycoinMainNetParams)
	net, invalidNet := "MainNet", "custom-net%d"

	for i := 10; i < 20; i++ {
		name := fmt.Sprintf(invalidNet, i)
		t.Run(name, func(t *testing.T) {
			pex, err := plugin.LoadPEX(name)
			require.Nil(t, pex)
			require.NotNil(t, err)

			api, err1 := plugin.LoadTransactionAPI(name)
			require.Nil(t, api)
			require.NotNil(t, err1)
		})
	}
	t.Run(net, func(t *testing.T) {
		pex, err := plugin.LoadPEX(net)
		require.Nil(t, err)
		require.IsType(t, new(SkycoinPEX), pex)
		spex := pex.(*SkycoinPEX)
		require.Equal(t, PoolSection, spex.poolSection)

		api, err1 := plugin.LoadTransactionAPI(net)
		require.Nil(t, err1)
		require.IsType(t, new(SkycoinBlockchain), api)
		sAPI := api.(*SkycoinBlockchain)
		require.Equal(t, uint64(params.DataRefreshTimeout*1000000000), sAPI.CacheTime)
	})
}

func TestSkyFiberPluginAddressFromString(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		wantErr bool
	}{
		{
			name:    "address1",
			args:    "R6aHqKWSQfvpdo2fGSrq4F1RYXkBWR9HHJ",
			wantErr: false,
		},
		{
			name:    "address2",
			args:    "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt",
			wantErr: false,
		},
		{
			name:    "empty",
			args:    "",
			wantErr: true,
		},
		{
			name:    "invalid character",
			args:    "701d23fd513bad325938ba56869f9faba19384a8ec3dd41833aff147eac53947",
			wantErr: true,
		},
		{
			name:    "invalid checksum",
			args:    "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKk",
			wantErr: true,
		},
	}

	plugin := NewSkyFiberPlugin(SkycoinMainNetParams)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr, err := plugin.AddressFromString(tt.args)
			if tt.wantErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
				require.IsType(t, new(SkycoinAddress), addr)
			}
		})
	}
}

func TestSkyFiberPluginPubKeyFromBytes(t *testing.T) {
	pubkeyFromHex := func(s string) []byte {
		b, err := hex.DecodeString(s)
		require.NoError(t, err)
		return b
	}
	tests := []struct {
		name    string
		pubkHex string
		wantErr bool
	}{
		{name: "OK", pubkHex: "034f1e3f2391bd3670151fd4fa3accc6a0273885984404089e5b846871db4c5304", wantErr: false},
		{name: "wrong", pubkHex: "0213c9273d9f944c3d907bfe844090ecef3d9504c88c5165cb690de98125a4e4", wantErr: true},
		{name: "OK2", pubkHex: "0304eb48d7c0b3a915d0f53c6d966f4d9fa75df645e63dfb91d589592790943613", wantErr: false},
	}

	plugin := NewSkyFiberPlugin(SkycoinMainNetParams)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sPk, err := plugin.PubKeyFromBytes(pubkeyFromHex(tt.pubkHex))
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.IsType(t, new(SkycoinPubKey), sPk)
			got := sPk.(*SkycoinPubKey)
			require.NoError(t, err)
			assert.Equal(t, pubkeyFromHex(tt.pubkHex), got.Bytes())
			assert.Equal(t, tt.pubkHex, got.pubkey.Hex())
			require.NoError(t, got.Verify())
			assert.False(t, got.Null())
			spk, err := toSkycoinPubKey(got)
			require.NoError(t, err)
			assert.Equal(t, spk.Hex(), tt.pubkHex)
			require.NoError(t, spk.Verify())
			assert.False(t, spk.Null())
		})
	}
}

func TestSkyFiberPluginSecKeyFromBytes(t *testing.T) {
	pubkeyFromHex := func(s string) []byte {
		b, err := hex.DecodeString(s)
		require.NoError(t, err)
		return b
	}
	tests := []struct {
		name    string
		pubkHex string
		wantErr bool
	}{
		{name: "OK", pubkHex: "c9135a2b667eb0847fb7ad3d1ae58a1ea2d0c38526c8948b520417dcab618563", wantErr: false},
		{name: "wrong", pubkHex: "0213c9273d9f944c3d907bfe844090ecef3d9504c88c5165cb690de985a4e4", wantErr: true},
		{name: "OK2", pubkHex: "408ea9aef71391071d275f8255bd9b6d22d5d5a22e6ab2bfc54307fb273d468c", wantErr: false},
	}

	plugin := NewSkyFiberPlugin(SkycoinMainNetParams)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sSk, err := plugin.SecKeyFromBytes(pubkeyFromHex(tt.pubkHex))
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.IsType(t, new(SkycoinSecKey), sSk)
			got := sSk.(*SkycoinSecKey)
			require.NoError(t, err)
			assert.Equal(t, pubkeyFromHex(tt.pubkHex), got.Bytes())
			assert.Equal(t, tt.pubkHex, got.seckey.Hex())
			require.NoError(t, got.Verify())
			assert.False(t, got.Null())
			spk, err := toSkycoinSecKey(got)
			require.NoError(t, err)
			assert.Equal(t, spk.Hex(), tt.pubkHex)
			require.NoError(t, spk.Verify())
			assert.False(t, spk.Null())
		})
	}
}
