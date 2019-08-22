package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

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
	sourceType int
	source     string
}

func (ws *WalletSource) GetType() int {
	return ws.sourceType
}
func (ws *WalletSource) GetSource() string {
	return ws.source
}

type configManagerJson struct {
	SourceList []*walletSourceJson `json:"SourceList"`
	Node       string              `json:"Node"`
}

type walletSourceJson struct {
	SourceType int    `json:"Type"`
	Source     string `json:"Source"`
}

func (cm *ConfigManager) GetSources() []*WalletSource {
	return cm.sourceList
}

func GetConfigManager() *ConfigManager {
	once.Do(func() {
		var cm *ConfigManager

		if configFileExist() {
			fmt.Println("CONFIG-EXIST")
			cm = loadConfigFromFile()
		} else {
			fmt.Println("CONFIG-NOT-EXIST")
			cm = getDefaultConfigManager()
			fmt.Println("CONFIG-OBTAINED")
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
	fmt.Println(3)
	jsonFormat, err := json.Marshal(cm)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsonFormat))
	os.MkdirAll(filepath.Dir(getConfigFileDir()), 0755)
	//fmt.Println(err.Error())
	fmt.Println(jsonFormat)
	ioutil.WriteFile(getConfigFileDir(), jsonFormat, 0644)
	//fmt.Println(err.Error())
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
	walletsDir := filepath.Join(os.Getenv("HOME"), pathToDefaultWalletsFromHome)
	ws.source = walletsDir
	return ws

}
