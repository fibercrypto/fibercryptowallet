package skycoin

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	//local "github.com/fibercrypto/fibercryptowallet/src/main"
	appParams "github.com/fibercrypto/fibercryptowallet/src/params"
)

// SkyFiberPlugin provide support for SkyFiber coins
type SkyFiberPlugin struct {
	Params params.SkyFiberParams
}

// ListSupportedAltcoins to enumerate supported assets and related metadata
func (p *SkyFiberPlugin) ListSupportedAltcoins() []core.AltcoinMetadata {
	return []core.AltcoinMetadata{
		core.AltcoinMetadata{
			Name:     SkycoinName,
			Ticker:   SkycoinTicker,
			Family:   SkycoinFamily,
			HasBip44: false,
			Accuracy: 6,
		},
		core.AltcoinMetadata{
			Name:     CoinHoursName,
			Ticker:   CoinHoursTicker,
			Family:   SkycoinFamily,
			HasBip44: false,
			Accuracy: 0,
		},
		core.AltcoinMetadata{
			Name:     CalculatedHoursName,
			Ticker:   CalculatedHoursTicker,
			Family:   SkycoinFamily,
			HasBip44: false,
			Accuracy: 0,
		},
	}
}

// ListSupportedFamilies classifies similar cryptocurrencies into a family
func (p *SkyFiberPlugin) ListSupportedFamilies() []string {
	return []string{SkycoinFamily}
}

// RegisterTo boilerplate to register this plugin against an altcoin manager and enable it
func (p *SkyFiberPlugin) RegisterTo(manager core.AltcoinManager) {
	for _, info := range p.ListSupportedAltcoins() {
		manager.RegisterAltcoin(info, p)
	}
}

// GetName provides concise human-readable caption o identify this plugin
func (p *SkyFiberPlugin) GetName() string {
	return "SkyFiber"
}

// GetDescription describes plugin and its features
func (p *SkyFiberPlugin) GetDescription() string {
	return "FiberCrypto wallet connector for Skycoin and SkyFiber altcoins"
}

// LoadWalletEnvs loads wallet environments to lookup and create wallets
func (p *SkyFiberPlugin) LoadWalletEnvs() []core.WalletEnv {

	wltSources, err := config.GetWalletSources()
	if err != nil {
		return nil
	}
	wltEnvs := make([]core.WalletEnv, 0)
	for _, wltS := range wltSources {

		tp := wltS.Tp
		source := wltS.Source
		var wltEnv core.WalletEnv
		if tp == string(config.LocalWallet) {
			wltEnv = &WalletDirectory{WalletDir: source}
		} else if tp == string(config.RemoteWallet) {
			wltEnv = NewWalletNode(source)
		}
		wltEnvs = append(wltEnvs, wltEnv)
	}

	return wltEnvs
}

// LoadPEX instantiates proxy object to interact with nodes nodes of the P2P network
func (p *SkyFiberPlugin) LoadPEX(netType string) (core.PEX, error) {
	if netType != "MainNet" {
		return nil, errors.ErrInvalidNetworkType
	}
	return NewSkycoinPEX(PoolSection), nil
}

// LoadTransactionAPI blockchain transaction API entry poiny
func (p *SkyFiberPlugin) LoadTransactionAPI(netType string) (core.BlockchainTransactionAPI, error) {
	if netType != "MainNet" {
		return nil, errors.ErrInvalidNetworkType
	}
	return NewSkycoinBlockchain(appParams.DataRefreshTimeout), nil
}

// LoadSignService sign service entry point
func (p *SkyFiberPlugin) LoadSignService() (core.BlockchainSignService, error) {
	return &SkycoinSignService{}, nil
}

// AddressFromString retrieves address corresponding to readable representation
func (p *SkyFiberPlugin) AddressFromString(addrStr string) (core.Address, error) {
	addr, err := NewSkycoinAddress(addrStr)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

// PubKeyFromBytes retrieves address corresponding to readable representation
func (p *SkyFiberPlugin) PubKeyFromBytes(b []byte) (core.PubKey, error) {
	return skyPubKeyFromBytes(b)
}

// SecKeyFromBytes retrieves address corresponding to readable representation
func (p *SkyFiberPlugin) SecKeyFromBytes(b []byte) (core.SecKey, error) {
	return skySecKeyFromBytes(b)
}

// NewSkyFiberPlugin instantiate SkyFiber plugin entry point
func NewSkyFiberPlugin(params params.SkyFiberParams) core.AltcoinPlugin {
	return &SkyFiberPlugin{
		Params: params,
	}
}

// Type assertions
var (
	_ core.AltcoinPlugin = &SkyFiberPlugin{}
)
