package core

import "github.com/fibercrypto/FiberCryptoWallet/src/util/logging"

var logPlugin = logging.MustGetLogger("Altcoin Plugin")

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
	LoadPEX(netType string) (PEX, error)
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
	logPlugin.Info("Register plugin to Altcoin manager")
	p.RegisterTo(m)
	m.registeredPlugins = append(m.registeredPlugins, p)
}

func (m *fibercryptoAltcoinManager) RegisterAltcoin(info AltcoinMetadata, plugin AltcoinPlugin) {
	logPlugin.Info("Register altcoin to Altcoin manager")
	m.altcoinMap[info.Ticker] = altcoinRecord{
		Manager:  plugin,
		Metadata: info,
	}
}

func (m *fibercryptoAltcoinManager) ListRegisteredPlugins() []AltcoinPlugin {
	logPlugin.Info("Listing registered plugins in Altcoin manager")
	return m.registeredPlugins
}

func (m *fibercryptoAltcoinManager) LookupAltcoinManager(ticker string) (AltcoinPlugin, bool) {
	logPlugin.Info("Looking up for registered altcoin's")
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Manager, true
	}
	return nil, false
}

func (m *fibercryptoAltcoinManager) DescribeAltcoin(ticker string) (AltcoinMetadata, bool) {
	logPlugin.Info("Describing Altcoin manager")
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Metadata, true
	}
	return AltcoinMetadata{}, false
}

func LoadAltcoinManager() AltcoinManager {
	logPlugin.Info("Loading Altcoin manager")
	if manager.altcoinMap == nil {
		manager.altcoinMap = make(map[string]altcoinRecord, 5)
	}
	return &manager
}
