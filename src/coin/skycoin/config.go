package skycoin

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"strings"

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

func getMultiPlatformUserDirectory() string {
	usr, err := user.Current()
	if err != nil {
		//TODO: Log error
		return ""
	}
	return filepath.Join(usr.HomeDir, string(os.PathSeparator), ".skycoin", string(os.PathSeparator), "wallets")
}

func registerConfig() error {
	cm := local.GetConfigManager()
	node := map[string]string{"node": "https://staging.node.skycoin.net"}
	nodeBytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	nodeOpt := local.NewOption(SettingPathToNode, []string{}, false, string(nodeBytes))
	walletsDefaultDirectory := getMultiPlatformUserDirectory()
	wltSrc := &walletSource{
		id:     1,
		Tp:     string(LocalWallet),
		Source: walletsDefaultDirectory,
	}
	wltSrcBytes, err := json.Marshal(wltSrc)
	if err != nil {
		return err
	}

	wltOpt := local.NewOption(string(wltSrc.id), []string{SettingPathToWalletSource}, false, string(wltSrcBytes))

	sectionManager = cm.RegisterSection(SectionName, []*local.Option{nodeOpt, wltOpt})
	return nil
}

func getOption(path string) (string, error) {
	stringList := strings.Split(path, "/")
	return sectionManager.GetValue(stringList[len(stringList)-1], stringList[:len(stringList)-1])
}

func getValues(prefix string) ([]string, error) {
	return sectionManager.GetValues(strings.Split(prefix, "/"))
}

func getWalletSources() ([]*walletSource, error) {
	wltsString, err := getValues(SettingPathToWalletSource)
	if err != nil {
		return nil, err
	}
	wltSrcs := make([]*walletSource, len(wltsString))
	for i, wlt := range wltsString {
		wltSrcs[i] = new(walletSource)
		err = json.Unmarshal([]byte(wlt), wltSrcs[i])
		if err != nil {
			return nil, err
		}
	}
	return wltSrcs, nil
}

type walletSource struct {
	id     int
	Tp     string `json:"SourceType"`
	Source string `json:"Source"`
}
