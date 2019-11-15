package local

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/fibercrypto/FiberCryptoWallet/src/errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
)

var logConfigManager = logging.MustGetLogger("ConfigManager")

const (
	pathToConfigFromHome         = ".fiber/config.json"
	pathToDefaultWalletsFromHome = ".skycoin/wallets"
	LocalWallet                  = iota
	RemoteWallet
)

var (
	confManager *ConfigManager
	once        sync.Once
)

type ConfigManager struct {
	sourceList []*WalletSource
	node       string
}

type WalletSource struct {
	id         int
	sourceType int
	source     string
}

func (ws *WalletSource) GetType() int {
	logConfigManager.Info("Getting wallet type")
	return ws.sourceType
}
func (ws *WalletSource) GetSource() string {
	logConfigManager.Info("Getting wallet source")
	return ws.source
}
func (ws *WalletSource) GetId() int {
	logConfigManager.Info("Getting wallet id")
	return ws.id
}

func (ws *WalletSource) edit(source string, tp int) {
	logConfigManager.Info("Editing wallet source")
	ws.source = source
	ws.sourceType = tp
}

func (ws *WalletSource) getWalletSourceJson() *walletSourceJson {
	return &walletSourceJson{
		SourceType: ws.GetType(),
		Source:     ws.GetSource(),
	}
}

func (cm *ConfigManager) GetSources() []*WalletSource {
	return cm.sourceList
}

func (cm *ConfigManager) GetNode() string {
	return cm.node
}

func (cm *ConfigManager) EditWalletSource(id int, source string, tp int) error {
	var src *WalletSource
	for _, wltSrc := range cm.sourceList {
		if wltSrc.id == id {
			src = wltSrc
			break
		}
	}
	if src == nil {
		return errors.ErrInvalidID
	}

	if tp != LocalWallet && tp != RemoteWallet {
		tp = src.sourceType
	}

	if source == "" {
		source = src.source
	}

	src.edit(source, tp)
	return nil

}

func (cm *ConfigManager) EditNode(node string) {
	cm.node = node
}

func (cm *ConfigManager) Save() error {
	logConfigManager.Info("Saving configuration")

	jsonFormat, err := json.Marshal(cm.getConfigManagerJson())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(getConfigFileDir(), jsonFormat, 0644)
}

func (cm *ConfigManager) getConfigManagerJson() *configManagerJson {
	logConfigManager.Info("Getting configuration from JSON")

	wltSources := make([]*walletSourceJson, 0)
	for _, wltS := range cm.sourceList {
		wltSources = append(wltSources, wltS.getWalletSourceJson())
	}
	return &configManagerJson{
		SourceList: wltSources,
		Node:       cm.node,
	}
}

type configManagerJson struct {
	SourceList []*walletSourceJson `json:"SourceList"`
	Node       string              `json:"Node"`
}

func (cmJ *configManagerJson) getConfigManager() *ConfigManager {
	wltsSource := make([]*WalletSource, 0)
	for _, wltS := range cmJ.SourceList {
		wltsSource = append(wltsSource, wltS.getWalletSource())
	}

	return &ConfigManager{
		node:       cmJ.Node,
		sourceList: wltsSource,
	}
}

type walletSourceJson struct {
	SourceType int    `json:"Type"`
	Source     string `json:"Source"`
}

func (wsJ *walletSourceJson) getWalletSource() *WalletSource {
	return &WalletSource{
		source:     wsJ.Source,
		sourceType: wsJ.SourceType,
	}
}

func GetConfigManager() *ConfigManager {
	once.Do(func() {
		var cm *ConfigManager

		if configFileExist() {

			cm = loadConfigFromFile()
		} else {

			cm = getDefaultConfigManager()

		}
		confManager = cm
	})

	return confManager
}

func configFileExist() bool {
	fileDir := getConfigFileDir()
	if _, err := os.Stat(fileDir); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func loadConfigFromFile() *ConfigManager {
	cm := new(configManagerJson)
	fileDir := getConfigFileDir()
	dat, err := ioutil.ReadFile(fileDir) //nolint gosec

	if err != nil {

		return getDefaultConfigManager()
	}
	err = json.Unmarshal(dat, cm)
	if err != nil {

		return getDefaultConfigManager()
	}
	configM := cm.getConfigManager()
	cont := 1
	for _, ws := range configM.sourceList {
		ws.id = cont
		cont++
	}
	return configM

}

func getDefaultConfigManager() *ConfigManager {

	cm := new(ConfigManager)

	cm.node = "https://staging.node.skycoin.net"
	cm.sourceList = []*WalletSource{getDefaultWalletSource()}

	jsonFormat, err := json.Marshal(cm.getConfigManagerJson())
	if err != nil {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(getConfigFileDir()), 0750)
	if err != nil {
		return nil
	}

	err = ioutil.WriteFile(getConfigFileDir(), jsonFormat, 0644)
	if err != nil {
		return nil
	}

	return cm

}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, pathToConfigFromHome)
	return fileDir
}

func getDefaultWalletSource() *WalletSource {
	ws := new(WalletSource)
	ws.sourceType = LocalWallet
	ws.id = 1
	walletsDir := filepath.Join(os.Getenv("HOME"), pathToDefaultWalletsFromHome)
	ws.source = walletsDir
	return ws

}
