package skycoin

import (
	"encoding/json"

	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
)

const (
	LocalWallet = iota
	RemoteWallet
)

var (
	sectionManager *local.SectionManager
)

func registerConfig() error {
	cm := local.GetConfigManager()
	node := local.NewOption("node", false, "https://staging.node.skycoin.net")
	wltSrc := &walletSource{
		tp:     LocalWallet,
		source: "/home/kid/.skycoin/wallets",
	}
	wltSrcBytes, err := json.Marshal(wltSrc)
	if err != nil {
		return err
	}

	wltOpt := local.NewOption("walletSource", false, string(wltSrcBytes))

	sectionManager = cm.RegisterSection()

}

type walletSource struct {
	tp     int    `json:"SourceType"`
	source string `json:"Source"`
}
