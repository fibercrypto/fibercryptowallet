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
	err := config.RegisterConfig()
	if err != nil {
		logSkycoin.Warn("Couldn't register Skycoin configuration")
	}

	// logStr, err := config.GetOption(config.SettingPathToNode)
	// if err != nil {
	// 	logSkycoin.Warn("Couldn't get log options")
	// 	logSkycoin.WithError(err).Error("AA")
	// }
	// log := make(map[string]string)
	// err = json.Unmarshal([]byte(logStr), &log)
	// if err != nil {
	// 	logSkycoin.Warn("Couldn't unmarshal from options")
	// }
	// level, err := logging.LevelFromString(log["level"])
	// if err != nil {
	// 	logSkycoin.Warn("Couldn't get level from logging " + log["level"])
	// }
	// logging.SetLevel(level)

	nodeStr, err := config.GetOption(config.SettingPathToNode)
	if err != nil {
		logSkycoin.Warn("Couldn't get node options")
	}
	node := make(map[string]string)
	err = json.Unmarshal([]byte(nodeStr), &node)
	if err != nil {
		logSkycoin.Warn("Couldn't unmarshal from options")
	}
	level, err := logging.LevelFromString(node["level"])
	if err != nil {
		logSkycoin.Warn("Couldn't get level from logging")
		logSkycoin.WithError(err).WithField("string", node["level"]).Error()
	}
	logging.SetLevel(level)
	err = core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(node["address"]))
	if err != nil {
		logSkycoin.Warn("Couldn't create section for Skycoin")
	}
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))

}
