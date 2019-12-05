package skycoin //nolint goimports

import (
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"

	util "github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/stretchr/testify/require"
)

func TestRegisterSkycoinPlugin(t *testing.T) {
	require.Equal(t, "Skycoin", util.AltcoinCaption(params.SkycoinTicker))
	require.Equal(t, "Coin Hours", util.AltcoinCaption(params.CoinHoursTicker))
	require.Equal(t, "Calculated Hours", util.AltcoinCaption(params.CalculatedHoursTicker))
}
