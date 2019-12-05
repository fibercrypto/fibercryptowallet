package skycoin //nolint goimports

import (
	"encoding/json"

	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	sky "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	local "github.com/fibercrypto/fibercryptowallet/src/main"

	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

func init() {
	config.RegisterConfig()
	nodeStr, _ := config.GetOption(config.SettingPathToNode)
	node := make(map[string]string, 0)
	json.Unmarshal([]byte(nodeStr), &node)
	core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(node["address"]))
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))

}
