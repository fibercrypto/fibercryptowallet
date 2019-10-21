package skycoin

import (
	"encoding/json"
	"fmt"

	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
)

const (
	LocalWallet = iota
	RemoteWallet
	SectionName               = "skycoin"
	SettingPathToNode         = "node"
	SettingPathToWalletSource = "walletSource"
)

var (
	sectionManager *local.SectionManager
)

func registerConfig() error {
	cm := local.GetConfigManager()
	node := local.NewOption(SettingPathToNode, false, "https://staging.node.skycoin.net")
	wltSrc := &walletSource{
		id:     1,
		Tp:     LocalWallet,
		Source: "/home/kid/.skycoin/wallets",
	}
	wltSrcBytes, err := json.Marshal(wltSrc)
	if err != nil {
		return err
	}

	wltOpt := local.NewOption(fmt.Sprintf("%s/%d", SettingPathToWalletSource, wltSrc.id), false, string(wltSrcBytes))

	sectionManager = cm.RegisterSection(SectionName, []*local.Option{node, wltOpt})
	return nil
}

func getOption(path string) (string, error) {
	return sectionManager.GetValue(path)
}

func getValues(prefix string) ([]string, error) {
	return sectionManager.GetValues(prefix)
}

func getWalletSources() ([]*walletSource, error) {
	wltsString, err := getValues(SettingPathToWalletSource)
	if err != nil {
		return nil, err
	}
	wltSrcs := make([]*walletSource, len(wltsString))
	for i, wlt := range wltsString {
		err = json.Unmarshal([]byte(wlt), wltSrcs[i])
		if err != nil {
			return nil, err
		}
	}

	return wltSrcs, nil
}

type walletSource struct {
	id     int
	Tp     int    `json:"SourceType"`
	Source string `json:"Source"`
}
