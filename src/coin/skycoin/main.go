package skycoin //nolint goimports

import (
	"encoding/json"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/config"
	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

func init() {
	config.RegisterConfig()
	nodeStr, _ := config.GetOption(config.SettingPathToNode)
	node := make(map[string]string, 0)
	json.Unmarshal([]byte(nodeStr), &node)
	core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(node["address"]))
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))

}
