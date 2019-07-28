package skycoin

import (
	"testing"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/stretchr/testify/require"
)

func TestRegisterSkycoinPlugin(t *testing.T) {
	require.Equal(t, "Skycoin", util.AltcoinCaption("SKY"))
	require.Equal(t, "Coin Hours", util.AltcoinCaption("SKYCH"))
}
