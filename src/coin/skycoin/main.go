package skycoin

import (
	skycoin "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type SkyFiberPlugin struct {
	Params SkyFiberParams
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

	config := core.GetConfigManager()
	wltSources := config.GetSources()

	wltEnvs := make([]core.WalletEnv, 0)
	for _, wltS := range wltSources {
		tp := wltS.GetType()
		source := wltS.GetSource()
		var wltEnv core.WalletEnv
		if tp == core.LocalWallet {
			wltEnv = &skycoin.WalletDirectory{WalletDir: source}
		} else if tp == core.RemoteWallet {
			wltEnv = &skycoin.WalletNode{NodeAddress: source}
		}
		wltEnvs = append(wltEnvs, wltEnv)
	}

	return wltEnvs
}

func NewSkyFiberPlugin(params SkyFiberParams) core.AltcoinPlugin {
	return &SkyFiberPlugin{
		Params: params,
	}
}

func init() {
	util.RegisterAltcoin(NewSkyFiberPlugin(SkycoinMainNetParams))
}
