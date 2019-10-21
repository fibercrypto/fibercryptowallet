package skycoin //nolint goimports

import (
	"github.com/fibercrypto/F/src/coin/skycoin"
	"testing"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/stretchr/testify/require"
)

func TestRegisterSkycoinPlugin(t *testing.T) {
	require.Equal(t, "Skycoin", util.AltcoinCaption(skycoin.SkycoinTicker))
	require.Equal(t, "Coin Hours", util.AltcoinCaption(skycoin.CoinHoursTicker))
}
