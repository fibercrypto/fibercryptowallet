package skycoin

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	local "github.com/fibercrypto/fibercryptowallet/src/main"
)

type SkyFiberPlugin struct {
	Params params.SkyFiberParams
}

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
			Family:   CoinHoursFamily,
			HasBip44: false,
			Accuracy: 0,
		},
		core.AltcoinMetadata{
			Name:     CalculatedHoursName,
			Ticker:   CalculatedHoursTicker,
			Family:   CalculatedHoursFamily,
			HasBip44: false,
			Accuracy: 0,
		},
	}
}

func (p *SkyFiberPlugin) ListSupportedFamilies() []string {
	return []string{SkycoinFamily}
}

func (p *SkyFiberPlugin) RegisterTo(manager core.AltcoinManager) {
	for _, info := range p.ListSupportedAltcoins() {
		manager.RegisterAltcoin(info, p)
	}
}

func (p *SkyFiberPlugin) GetName() string {
	return "SkyFiber"
}

func (p *SkyFiberPlugin) GetDescription() string {
	return "FiberCrypto wallet connector for Skycoin and SkyFiber altcoins"
}

func (p *SkyFiberPlugin) LoadWalletEnvs() []core.WalletEnv {

	config := local.GetConfigManager()
	wltSources := config.GetSources()

	wltEnvs := make([]core.WalletEnv, 0)
	for _, wltS := range wltSources {
		tp := wltS.GetType()
		source := wltS.GetSource()
		var wltEnv core.WalletEnv
		if tp == local.LocalWallet {
			wltEnv = &WalletDirectory{WalletDir: source}
		} else if tp == local.RemoteWallet {
			wltEnv = NewWalletNode(source)
		}
		wltEnvs = append(wltEnvs, wltEnv)
	}

	return wltEnvs
}

func (p *SkyFiberPlugin) LoadPEX(netType string) (core.PEX, error) {
	var poolSection string
	if netType == "MainNet" {
		poolSection = PoolSection
	} else {
		return nil, errors.ErrInvalidNetworkType
	}
	return NewSkycoinPEX(poolSection), nil

}

// AddressFromString retrieves address corresponding to readable representation
func (p *SkyFiberPlugin) AddressFromString(addrStr string) (core.Address, error) {
	return NewSkycoinAddress(addrStr)
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
