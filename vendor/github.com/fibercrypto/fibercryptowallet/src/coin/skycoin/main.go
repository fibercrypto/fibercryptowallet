package skycoin //nolint goimports

import (
	sky "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	local "github.com/fibercrypto/fibercryptowallet/src/main"

	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

func init() {
	cf := local.GetConfigManager()
	err := core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(cf.GetNode()))
	if err != nil {
		return
	}
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))
}
