package pluginutil

import (
	"github.com/FiberCrypto/FibercryptoWallet/src/core"
)

func AltcoinCaption(ticker string) string {
	if info, isRegistered := core.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return info.Name
	}
	return ticker + " <Unregistered>"
}

func RegisterAltcoin(p core.AltcoinPlugin) {
	core.LoadAltcoinManager().RegisterPlugin(p)
}
