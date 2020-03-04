package local

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

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
	p.RegisterTo(m)
	m.registeredPlugins = append(m.registeredPlugins, p)
}

func (m *fibercryptoAltcoinManager) RegisterAltcoin(info core.AltcoinMetadata, plugin core.AltcoinPlugin) {
	m.altcoinMap[info.Ticker] = altcoinRecord{
		Manager:  plugin,
		Metadata: info,
	}
}

func (m *fibercryptoAltcoinManager) ListRegisteredPlugins() []core.AltcoinPlugin {
	return m.registeredPlugins
}

func (m *fibercryptoAltcoinManager) LookupAltcoinPlugin(ticker string) (core.AltcoinPlugin, bool) {
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Manager, true
	}
	return nil, false
}

func (m *fibercryptoAltcoinManager) DescribeAltcoin(ticker string) (core.AltcoinMetadata, bool) {
	if r, isRegistered := m.altcoinMap[ticker]; isRegistered {
		return r.Metadata, true
	}
	return core.AltcoinMetadata{}, false
}

// LoadAltcoinManager load altcoin manager singleton instance
func LoadAltcoinManager() core.AltcoinManager {
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
