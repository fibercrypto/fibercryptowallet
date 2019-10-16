package skycoin //nolint goimports

import (
	"errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"

	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"

	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

var logFiberPlugin = logging.MustGetLogger("Skycoin Plugin")

type SkyFiberPlugin struct {
	Params params.SkyFiberParams
}

func (p *SkyFiberPlugin) ListSupportedAltcoins() []core.AltcoinMetadata {
	logFiberPlugin.Info("Listing supported Altcoin's for SkyFiber plugin")
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
	logFiberPlugin.Info("Listing supported families for SkyFiber plugin")
	return []string{params.SkycoinFamily}
}

func (p *SkyFiberPlugin) RegisterTo(manager core.AltcoinManager) {
	logFiberPlugin.Info("Registering SkyFiber plugin to Altcoin Manager")
	for _, info := range p.ListSupportedAltcoins() {
		manager.RegisterAltcoin(info, p)
	}
}

func (p *SkyFiberPlugin) GetName() string {
	logFiberPlugin.Info("Getting SkyFiber plugin name")
	return "SkyFiber"
}

func (p *SkyFiberPlugin) GetDescription() string {
	logFiberPlugin.Info("Getting SkyFiber plugin description")
	return "FiberCrypto wallet connector for Skycoin and SkyFiber altcoins"
}

func (p *SkyFiberPlugin) LoadWalletEnvs() []core.WalletEnv {
	logFiberPlugin.Info("Loading wallets env for SkyFiber plugin")
	config := local.GetConfigManager()
	wltSources := config.GetSources()

	wltEnvs := make([]core.WalletEnv, 0)
	for _, wltS := range wltSources {
		tp := wltS.GetType()
		source := wltS.GetSource()
		var wltEnv core.WalletEnv
		if tp == local.LocalWallet {
			wltEnv = &sky.WalletDirectory{WalletDir: source}
		} else if tp == local.RemoteWallet {
			wltEnv = sky.NewWalletNode(source)
		}
		wltEnvs = append(wltEnvs, wltEnv)
	}

	return wltEnvs
}

func (p *SkyFiberPlugin) LoadPEX(netType string) (core.PEX, error) {
	logFiberPlugin.Info("Loading PEX for SkyFiber plugin")
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
	cf := local.GetConfigManager()
	err := core.GetMultiPool().CreateSection(sky.PoolSection, sky.NewSkycoinConnectionFactory(cf.GetNode()))
	if err != nil {
		return
	}
	util.RegisterAltcoin(NewSkyFiberPlugin(params.SkycoinMainNetParams))
}
