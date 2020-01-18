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

	logSettingStr, err := config.GetOption(config.SettingPathToLog)
	if err != nil {
		logSkycoin.Warn("Couldn't get log options")
	}
	logSetting := make(map[string]string)
	err = json.Unmarshal([]byte(logSettingStr), &logSetting)
	if err != nil {
		logSkycoin.Warn("Couldn't unmarshal from options")
	}
	level, err := logging.LevelFromString(logSetting["level"])
	if err != nil {
		logSkycoin.Warn("Couldn't get level from logging")
		logSkycoin.WithError(err).WithField("string", logSetting["level"]).Error()
	}
	logging.SetLevel(level)

	nodeSettingStr, err := config.GetOption(config.SettingPathToNode)
	if err != nil {
		logSkycoin.Warn("Couldn't get node settings")
	}
	node := make(map[string]string)
	err = json.Unmarshal([]byte(nodeSettingStr), &node)

	err = core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(node["address"]))
	if err != nil {
		logSkycoin.Warn("Couldn't create section for Skycoin")
	}
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))

}
