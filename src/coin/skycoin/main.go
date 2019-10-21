package skycoin //nolint goimports

import (
	"errors"

	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type SkyFiberPlugin struct {
	Params params.SkyFiberParams
}

func (p *SkyFiberPlugin) ListSupportedAltcoins() []core.AltcoinMetadata {
	return []core.AltcoinMetadata{
		core.AltcoinMetadata{
			Name:     params.SkycoinName,
			Ticker:   params.SkycoinTicker,
			Family:   params.SkycoinFamily,
			HasBip44: false,
			Accuracy: 6,
		},
		core.AltcoinMetadata{
			Name:     params.CoinHoursName,
			Ticker:   params.CoinHoursTicker,
			Family:   params.CoinHoursFamily,
			HasBip44: false,
			Accuracy: 0,
		},
		core.AltcoinMetadata{
			Name:     params.CalculatedHoursName,
			Ticker:   params.CalculatedHoursTicker,
			Family:   params.CalculatedHoursFamily,
			HasBip44: false,
			Accuracy: 0,
		},
	}
}

func (p *SkyFiberPlugin) ListSupportedFamilies() []string {
	return []string{params.SkycoinFamily}
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

	wltSources, _ := getWalletSources()
	//TODO: Log if errors happens

	wltEnvs := make([]core.WalletEnv, 0)
	for _, wltS := range wltSources {
		tp := wltS.Tp
		source := wltS.Source
		var wltEnv core.WalletEnv
		if tp == LocalWallet {
			wltEnv = &sky.WalletDirectory{WalletDir: source}
		} else if tp == RemoteWallet {
			wltEnv = sky.NewWalletNode(source)
		}
		wltEnvs = append(wltEnvs, wltEnv)
	}

	return wltEnvs
}

func (p *SkyFiberPlugin) LoadPEX(netType string) (core.PEX, error) {
	var poolSection string
	if netType == "MainNet" {
		poolSection = sky.PoolSection
	} else {
		return nil, errors.New("Invalid netType")
	}
	return sky.NewSkycoinPEX(poolSection), nil

}
func NewSkyFiberPlugin(params params.SkyFiberParams) core.AltcoinPlugin {
	return &SkyFiberPlugin{
		Params: params,
	}
}

func init() {
	registerConfig()
	node, _ := getOption(SettingPathToNode)
	core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(node))
	util.RegisterAltcoin(NewSkyFiberPlugin(params.SkycoinMainNetParams))
}
