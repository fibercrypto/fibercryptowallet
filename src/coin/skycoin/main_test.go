package skycoin //nolint goimports

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"testing"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/stretchr/testify/require"
)

func TestRegisterSkycoinPlugin(t *testing.T) {
	require.Equal(t, "Skycoin", util.AltcoinCaption(params.SkycoinTicker))
	require.Equal(t, "Coin Hours", util.AltcoinCaption(params.CoinHoursTicker))
}
