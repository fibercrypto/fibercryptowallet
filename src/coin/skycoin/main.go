package skycoin //nolint goimports

import (
	"encoding/json"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/config"
	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

func init() {
	registerConfig()
	nodeStr, _ := GetOption(config.SettingPathToNode)
	node := make(map[string]string, 0)
	json.Unmarshal([]byte(nodeStr), &node)
	core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(node["node"]))

}
