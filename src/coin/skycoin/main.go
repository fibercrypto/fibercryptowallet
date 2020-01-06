package skycoin //nolint goimports

import (
	"encoding/json"

	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	sky "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"

	//local "github.com/fibercrypto/fibercryptowallet/src/main"

	util "github.com/fibercrypto/fibercryptowallet/src/util"
)

var logSkycoin = logging.MustGetLogger("Skycoin Altcoin")

func init() {
	UpdateAltcoin()
}

// Refresh Skycoin Altcoin node settings
func UpdateAltcoin() {
	err := config.RegisterConfig()
	if err != nil {
		logSkycoin.Warn("Couldn't register Skycoin configuration")
	}
	nodeStr, err := config.GetOption(config.SettingPathToNode)
	if err != nil {
		logSkycoin.Warn("Couldn't get node options")
	}
	node := make(map[string]string)
	err = json.Unmarshal([]byte(nodeStr), &node)
	if err != nil {
		logSkycoin.Warn("Couldn't unmarshal from options")
	}
	err = core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(node["address"]))
	if err != nil {
		logSkycoin.Warn("Couldn't create section for Skycoin")
	}
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))
}
