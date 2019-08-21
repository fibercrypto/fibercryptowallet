package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	qtcore "github.com/therecipe/qt/core"
)

const (
	pathToConfigFromHome         = ".fiber/config.json"
	pathToDefaultWalletsFromHome = ".skycoin/wallets"
	localWallet                  = iota
	remoteWallet
)

type walletSource struct {
	sourceType int    `json:"type"`
	source     string `json:"source"`
}

type WalletSource struct {
	qtcore.QObject
	_ int    `property:"sourceType"`
	_ string `property:"source"`
}

func getDefaultWalletSource() *walletSource {
	ws := new(walletSource)
	ws.sourceType = localWallet
	walletsDir := filepath.Join(os.Getenv("HOME"), pathToDefaultWalletsFromHome)
	ws.source = walletsDir
	return ws

}

type ConfigManager struct {
	qtcore.QObject
	sourceList []*walletSource        `json:"SourceList"`
	node       string                 `json:"Node"`
	_          func()                 `constructor:"init"`
	_          func() []*WalletSource `slot:"getSources"`
}

func (cm *ConfigManager) init() {
	cm.ConnectGetSources(cm.getSources)
	if configFileExist() {
		cm = loadConfigFromFile()
	} else {
		cm = getDefaultConfigManager()
	}

}

func (cm *ConfigManager) getSources() []*WalletSource {
	wltsSource := make([]*WalletSource, 0)
	for _, wltS := range cm.sourceList {
		newWltSource := NewWalletSource(nil)
		newWltSource.SetSourceType(wltS.sourceType)
		newWltSource.SetSource(wltS.source)
		wltsSource = append(wltsSource, newWltSource)
	}
	return wltsSource
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
	cm := NewConfigManager(nil)
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
	cm := NewConfigManager(nil)
	cm.node = "http://stagin.node.skycoin.net"
	cm.sourceList = []*walletSource{getDefaultWalletSource()}
	jsonFormat, _ := json.Marshal(cm)
	ioutil.WriteFile(getConfigFileDir(), jsonFormat, 0644)
	return cm

}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, pathToConfigFromHome)
	return fileDir
}
