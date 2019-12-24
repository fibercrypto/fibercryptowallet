package local

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logMainPlugin = logging.MustGetLogger("Altcoin plugin entry point")

type altcoinRecord struct {
	Manager  core.AltcoinPlugin
	Metadata core.AltcoinMetadata
}

// fibercoinAltcoinManager is a singleton class
type fibercryptoAltcoinManager struct {
	registeredPlugins []core.AltcoinPlugin
	altcoinMap        map[string]altcoinRecord
	signers           map[core.UID]core.TxnSigner
}

var (
	manager fibercryptoAltcoinManager
)

func (m *fibercryptoAltcoinManager) RegisterPlugin(p core.AltcoinPlugin) {
	logMainPlugin.Info("Register plugin to Altcoin manager")
	p.RegisterTo(m)
	m.registeredPlugins = append(m.registeredPlugins, p)
}

func (m *fibercryptoAltcoinManager) RegisterAltcoin(info core.AltcoinMetadata, plugin core.AltcoinPlugin) {
	logMainPlugin.Info("Register altcoin to Altcoin manager")
	m.altcoinMap[info.Ticker] = altcoinRecord{
		Manager:  plugin,
		Metadata: info,
	}
}

func (m *fibercryptoAltcoinManager) ListRegisteredPlugins() []core.AltcoinPlugin {
	logMainPlugin.Info("Listing registered plugins in Altcoin manager")
	return m.registeredPlugins
}

func (m *fibercryptoAltcoinManager) LookupAltcoinPlugin(ticker string) (core.AltcoinPlugin, bool) {
	logMainPlugin.Info("Looking up for registered altcoin's")
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Manager, true
	}
	return nil, false
}

func (m *fibercryptoAltcoinManager) DescribeAltcoin(ticker string) (core.AltcoinMetadata, bool) {
	logMainPlugin.Info("Describing Altcoin manager")
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Metadata, true
	}
	return core.AltcoinMetadata{}, false
}

// LoadAltcoinManager load altcoin manager singleton instance
func LoadAltcoinManager() core.AltcoinManager {
	logMainPlugin.Info("Loading Altcoin manager")
	if manager.altcoinMap == nil {
		manager.altcoinMap = make(map[string]altcoinRecord, 5)
		manager.signers = make(map[core.UID]core.TxnSigner, 5)
	}
	return &manager
}

// Type assertions
var (
	_ core.AltcoinManager = &manager
)
