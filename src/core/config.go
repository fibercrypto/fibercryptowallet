package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	pathToConfigFromHome         = ".fiber/config.json"
	pathToDefaultWalletsFromHome = ".skycoin/wallets"
	localWallet                  = iota
	remoteWallet
)

var (
	confManager *ConfigManager
	once        sync.Once
)

type ConfigManager struct {
	sourceList []*WalletSource `json:"SourceList"`
	node       string          `json:"Node"`
}

type WalletSource struct {
	sourceType int    `json:"type"`
	source     string `json:"source"`
}

func (ws *WalletSource) GetType() int {
	return ws.sourceType
}
func (ws *WalletSource) GetSource() string {
	return ws.source
}

func (cm *ConfigManager) GetSources() []*WalletSource {
	return cm.sourceList
}

func GetConfigManager() *ConfigManager {
	once.Do(func() {
		cm := new(ConfigManager)

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
	cm := new(ConfigManager)
	fileDir := getConfigFileDir()
	dat, err := ioutil.ReadFile(fileDir)
	if err != nil {
		return getDefaultConfigManager()
	}
	err = json.Unmarshal(dat, cm)
	if err != nil {
		return getDefaultConfigManager()
	}
	return cm

}

func getDefaultConfigManager() *ConfigManager {
	cm := new(ConfigManager)
	cm.node = "http://stagin.node.skycoin.net"
	cm.sourceList = []*WalletSource{getDefaultWalletSource()}
	jsonFormat, _ := json.Marshal(cm)
	ioutil.WriteFile(getConfigFileDir(), jsonFormat, 0644)
	return cm

}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, pathToConfigFromHome)
	return fileDir
}

func getDefaultWalletSource() *WalletSource {
	ws := new(WalletSource)
	ws.sourceType = localWallet
	walletsDir := filepath.Join(os.Getenv("HOME"), pathToDefaultWalletsFromHome)
	ws.source = walletsDir
	return ws

}
