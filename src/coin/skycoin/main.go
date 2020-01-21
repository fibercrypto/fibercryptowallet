package skycoin //nolint goimports

import (
	"encoding/json"

	skylog "github.com/SkycoinProject/skycoin/src/util/logging"
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
	} else {
		logging.SetLevel(level)
	}
	writer, err := logging.GetOutputWriter(logSetting["output"])
	if err != nil {
		logSkycoin.WithError(err).Error("Error opening file: ", logSetting["output"])
	} else {
		logging.SetOutputTo(writer)
		skylog.SetOutputTo(writer)
	}

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
