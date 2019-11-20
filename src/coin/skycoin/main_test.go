package skycoin //nolint goimports

import (
	"testing"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/stretchr/testify/require"

	"os"

	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"

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

func TestRegisterSkycoinPlugin(t *testing.T) {
	require.Equal(t, "Skycoin", util.AltcoinCaption(params.SkycoinTicker))
	require.Equal(t, "Coin Hours", util.AltcoinCaption(params.CoinHoursTicker))
	require.Equal(t, "Calculated Hours", util.AltcoinCaption(params.CalculatedHoursTicker))
}
