package core

type AltcoinMetadata struct {
	Name          string
	Ticker        string
	Family        string
	HasBip44      bool
	Bip44CoinType int32
	Accuracy      int32
}

type AltcoinPlugin interface {
	ListSupportedAltcoins() []AltcoinMetadata
	ListSupportedFamilies() []string
	RegisterTo(manager AltcoinManager)
	GetName() string
	GetDescription() string
	LoadWalletEnvs() []WalletEnv
}

type AltcoinManager interface {
	RegisterPlugin(p AltcoinPlugin)
	RegisterAltcoin(info AltcoinMetadata, plugin AltcoinPlugin)
	ListRegisteredPlugins() []AltcoinPlugin
	LookupAltcoinManager(ticker string) (AltcoinPlugin, bool)
	DescribeAltcoin(ticker string) (AltcoinMetadata, bool)
}

type altcoinRecord struct {
	Manager  AltcoinPlugin
	Metadata AltcoinMetadata
}

// fibercoinAltcoinManager is a singleton class
type fibercryptoAltcoinManager struct {
	registeredPlugins []AltcoinPlugin
	altcoinMap        map[string]altcoinRecord
}

var (
	manager fibercryptoAltcoinManager
)

func (m *fibercryptoAltcoinManager) RegisterPlugin(p AltcoinPlugin) {
	p.RegisterTo(m)
	m.registeredPlugins = append(m.registeredPlugins, p)
}

func (m *fibercryptoAltcoinManager) RegisterAltcoin(info AltcoinMetadata, plugin AltcoinPlugin) {
	m.altcoinMap[info.Ticker] = altcoinRecord{
		Manager:  plugin,
		Metadata: info,
	}
}

func (m *fibercryptoAltcoinManager) ListRegisteredPlugins() []AltcoinPlugin {
	return m.registeredPlugins
}

func (m *fibercryptoAltcoinManager) LookupAltcoinManager(ticker string) (AltcoinPlugin, bool) {
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Manager, true
	}
	return nil, false
}

func (m *fibercryptoAltcoinManager) DescribeAltcoin(ticker string) (AltcoinMetadata, bool) {
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Metadata, true
	}
	return AltcoinMetadata{}, false
}

func LoadAltcoinManager() AltcoinManager {
	if manager.altcoinMap == nil {
		manager.altcoinMap = make(map[string]altcoinRecord, 5)
	}
	return &manager
}
