package local

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	qtcore "github.com/therecipe/qt/core"
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

func init() {
	qs := qtcore.NewQSettings("Simelo", "FiberCrypto Wallet", nil)
	confManager = &ConfigManager{
		setting: qs,
	}
}

type ConfigManager struct {
	setting *qtcore.QSettings
}

func (cm *ConfigManager) RegisterSection(section *ConfigSection) error {
	sections := cm.setting.ChildGroups()
	for _, sect := range sections {
		if section.name == sect {
			return errors.New("Invalid section name")
		}
	}

	cm.setting.BeginGroup(section.name)
	for key, value := range section.options {
		if !key.optional {
			cm.setting.SetValue(key.name, qtcore.NewQVariant1(key._default))
		}
	}
	cm.setting.EndGroup()

	return nil
}

type ConfigSection struct {
	name    string
	options map[Option]string
}

func NewConfigSection(name string, options map[Option]string) *ConfigSection {
	return &ConfigSection{
		name:    name,
		options: options,
	}
}

type Option struct {
	name     string
	optional bool
	_default string
}

func NewOption(name string, optional bool, _default string) *Option {
	if !optional && _default == "" {
		return nil
	}
	return &Option{
		name:     name,
		optional: optional,
		_default: _default,
	}
}

type WalletSource struct {
	id         int
	sourceType int
	source     string
}

func (ws *WalletSource) GetType() int {
	return ws.sourceType
}
func (ws *WalletSource) GetSource() string {
	return ws.source
}
func (ws *WalletSource) GetId() int {
	return ws.id
}

func (ws *WalletSource) edit(source string, tp int) {
	ws.source = source
	ws.sourceType = tp
}

func (ws *WalletSource) getWalletSourceJson() *walletSourceJson {
	return &walletSourceJson{
		SourceType: ws.GetType(),
		Source:     ws.GetSource(),
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
	dat, err := ioutil.ReadFile(fileDir)

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

	jsonFormat, _ := json.Marshal(cm.getConfigManagerJson())

	os.MkdirAll(filepath.Dir(getConfigFileDir()), 0755)

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
	ws.sourceType = LocalWallet
	ws.id = 1
	walletsDir := filepath.Join(os.Getenv("HOME"), pathToDefaultWalletsFromHome)
	ws.source = walletsDir
	return ws

}
