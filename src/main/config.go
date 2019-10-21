package local

import (
	"errors"
	"fmt"
	"sync"

	qtcore "github.com/therecipe/qt/core"
)

var (
	confManager         *ConfigManager
	once                sync.Once
	OptionNotFoundError error = errors.New("Option not found")
)

func init() {
	qs := qtcore.NewQSettings("Simelo", "FiberCrypto Wallet", nil)
	fmt.Println(qs.ApplicationName())
	confManager = &ConfigManager{
		setting: qs,
	}
}

type ConfigManager struct {
	setting *qtcore.QSettings
}

func (cm *ConfigManager) RegisterSection(name string, options []*Option) *SectionManager {
	cm.setting.BeginGroup(name)
	defer cm.setting.EndGroup()
	defer cm.setting.Sync()
	fmt.Println(name)

	for _, opt := range options {
		if !opt.optional && !cm.setting.Contains(opt.name) {
			fmt.Println("REGISTERING VALUE")
			cm.setting.SetValue(opt.name, qtcore.NewQVariant1(opt._default))
		}
	}

	return &SectionManager{
		name:     name,
		settings: cm.setting,
	}
}

type SectionManager struct {
	name     string
	settings *qtcore.QSettings
}

func (sm *SectionManager) GetValue(path string) (string, error) {
	sm.settings.BeginGroup(sm.name)
	defer sm.settings.EndGroup()
	val := sm.settings.Value(path, nil)
	if val.IsNull() {
		return "", OptionNotFoundError
	}

	return val.ToString(), nil
}

func (sm *SectionManager) Save(path string, value string) error {
	sm.settings.BeginGroup(sm.name)
	defer sm.settings.EndGroup()

	if !sm.settings.Contains(path) {
		return OptionNotFoundError
	}
	sm.settings.SetValue(path, qtcore.NewQVariant1(value))

	return nil
}

func (sm *SectionManager) GetValues(prefix string) ([]string, error) {
	sm.settings.BeginGroup(sm.name)
	defer sm.settings.EndGroup()
	groups := sm.settings.ChildGroups()
	finded := false
	for _, grp := range groups {
		if grp == prefix {
			finded = true
			break
		}
	}
	if !finded {
		return nil, OptionNotFoundError
	}

	sm.settings.BeginGroup(prefix)
	defer sm.settings.EndGroup()
	return sm.settings.ChildKeys(), nil
}

type Option struct {
	name        string
	sectionPath []string
	optional    bool
	_default    string
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

func GetConfigManager() *ConfigManager {
	return confManager
}

// type WalletSource struct {
// 	id         int
// 	sourceType int
// 	source     string
// }

// func (ws *WalletSource) GetType() int {
// 	return ws.sourceType
// }
// func (ws *WalletSource) GetSource() string {
// 	return ws.source
// }
// func (ws *WalletSource) GetId() int {
// 	return ws.id
// }

// func (ws *WalletSource) edit(source string, tp int) {
// 	ws.source = source
// 	ws.sourceType = tp
// }

// func (ws *WalletSource) getWalletSourceJson() *walletSourceJson {
// 	return &walletSourceJson{
// 		SourceType: ws.GetType(),
// 		Source:     ws.GetSource(),
// 	}
// }

// type configManagerJson struct {
// 	SourceList []*walletSourceJson `json:"SourceList"`
// 	Node       string              `json:"Node"`
// }

// func (cmJ *configManagerJson) getConfigManager() *ConfigManager {
// 	wltsSource := make([]*WalletSource, 0)
// 	for _, wltS := range cmJ.SourceList {
// 		wltsSource = append(wltsSource, wltS.getWalletSource())
// 	}

// 	return &ConfigManager{
// 		node:       cmJ.Node,
// 		sourceList: wltsSource,
// 	}
// }

// type walletSourceJson struct {
// 	SourceType int    `json:"Type"`
// 	Source     string `json:"Source"`
// }

// func (wsJ *walletSourceJson) getWalletSource() *WalletSource {
// 	return &WalletSource{
// 		source:     wsJ.Source,
// 		sourceType: wsJ.SourceType,
// 	}
// }

// func GetConfigManager() *ConfigManager {
// 	once.Do(func() {
// 		var cm *ConfigManager

// 		if configFileExist() {

// 			cm = loadConfigFromFile()
// 		} else {

// 			cm = getDefaultConfigManager()

// 		}
// 		confManager = cm
// 	})

// 	return confManager
// }

// func configFileExist() bool {
// 	fileDir := getConfigFileDir()
// 	if _, err := os.Stat(fileDir); err != nil {
// 		if os.IsNotExist(err) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func loadConfigFromFile() *ConfigManager {
// 	cm := new(configManagerJson)
// 	fileDir := getConfigFileDir()
// 	dat, err := ioutil.ReadFile(fileDir)

// 	if err != nil {

// 		return getDefaultConfigManager()
// 	}
// 	err = json.Unmarshal(dat, cm)
// 	if err != nil {

// 		return getDefaultConfigManager()
// 	}
// 	configM := cm.getConfigManager()
// 	cont := 1
// 	for _, ws := range configM.sourceList {
// 		ws.id = cont
// 		cont++
// 	}
// 	return configM

// }

// func getDefaultConfigManager() *ConfigManager {

// 	cm := new(ConfigManager)

// 	cm.node = "https://staging.node.skycoin.net"
// 	cm.sourceList = []*WalletSource{getDefaultWalletSource()}

// 	jsonFormat, _ := json.Marshal(cm.getConfigManagerJson())

// 	os.MkdirAll(filepath.Dir(getConfigFileDir()), 0755)

// 	ioutil.WriteFile(getConfigFileDir(), jsonFormat, 0644)

// 	return cm

// }

// func getConfigFileDir() string {
// 	homeDir := os.Getenv("HOME")
// 	fileDir := filepath.Join(homeDir, pathToConfigFromHome)
// 	return fileDir
// }

// func getDefaultWalletSource() *WalletSource {
// 	ws := new(WalletSource)
// 	ws.sourceType = LocalWallet
// 	ws.id = 1
// 	walletsDir := filepath.Join(os.Getenv("HOME"), pathToDefaultWalletsFromHome)
// 	ws.source = walletsDir
// 	return ws

// }
