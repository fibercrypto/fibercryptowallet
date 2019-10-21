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

	for _, opt := range options {
		for _, sect := range opt.sectionPath {
			cm.setting.BeginGroup(sect)
			defer cm.setting.EndGroup()
		}
		if !opt.optional && !cm.setting.Contains(opt.name) {
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

func (sm *SectionManager) GetValue(name string, sectionPath []string) (string, error) {
	sm.settings.BeginGroup(sm.name)
	defer sm.settings.EndGroup()
	for _, sect := range sectionPath {
		groups := sm.settings.ChildGroups()
		finded := false
		for _, grp := range groups {
			if grp == sect {
				finded = true
				break
			}
		}
		if !finded {
			return "", OptionNotFoundError
		}
		sm.settings.BeginGroup(sect)
		defer sm.settings.EndGroup()
	}
	val := sm.settings.Value(name, qtcore.NewQVariant())
	if val.IsNull() {
		return "", OptionNotFoundError
	}
	return val.ToString(), nil
}

func (sm *SectionManager) Save(name string, sectionPath []string, value string) error {
	sm.settings.BeginGroup(sm.name)
	defer sm.settings.EndGroup()
	for _, sect := range sectionPath {
		groups := sm.settings.ChildGroups()
		finded := false
		for _, grp := range groups {
			if grp == sect {
				finded = true
				break
			}
		}
		if !finded {
			return OptionNotFoundError
		}
		sm.settings.BeginGroup(sect)
		defer sm.settings.EndGroup()
	}
	if !sm.settings.Contains(name) {
		return OptionNotFoundError
	}
	sm.settings.SetValue(name, qtcore.NewQVariant1(value))
	return nil
}

func (sm *SectionManager) GetValues(sectionPath []string) ([]string, error) {
	sm.settings.BeginGroup(sm.name)
	defer sm.settings.EndGroup()

	for _, sect := range sectionPath {
		groups := sm.settings.ChildGroups()
		finded := false
		for _, grp := range groups {
			if grp == sect {
				fmt.Println("FINDED ", grp)
				finded = true
				break
			}
		}
		if !finded {
			return nil, OptionNotFoundError
		}
		sm.settings.BeginGroup(sect)
		defer sm.settings.EndGroup()
	}
	keys := sm.settings.ChildKeys()
	values := make([]string, 0)
	for _, key := range keys {
		values = append(values, sm.settings.Value(key, qtcore.NewQVariant()).ToString())
	}
	return values, nil
}

type Option struct {
	name        string
	sectionPath []string
	optional    bool
	_default    string
}

func NewOption(name string, sectionPath []string, optional bool, _default string) *Option {
	if !optional && _default == "" {
		return nil
	}
	return &Option{
		name:        name,
		sectionPath: sectionPath,
		optional:    optional,
		_default:    _default,
	}
}

func GetConfigManager() *ConfigManager {
	return confManager
}
