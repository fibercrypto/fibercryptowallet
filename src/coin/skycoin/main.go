package skycoin

import (
	"errors"

	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

func init() {
	cf := local.GetConfigManager()
	core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(cf.GetNode()))
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))
}
