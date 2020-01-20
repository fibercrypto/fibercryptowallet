package skycoin

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	util "github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/stretchr/testify/mock"
)

var global_mock *SkycoinApiMock

var logModelTest = logging.MustGetLogger("Skycoin Model Test")

// CleanGlobalMock util when is needed to change the values of an

// API method used in other test with different values.
func CleanGlobalMock() {
	global_mock.ExpectedCalls = []*mock.Call{}
}

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
