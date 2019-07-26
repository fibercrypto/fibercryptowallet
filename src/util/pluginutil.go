package pluginutil

import (
	"github.github.com/FiberCrypto/FibercryptoWallet/src/core"
)

func AltcoinCaption(ticker string) string {
	if info, isRegistered := core.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return info.Name
	}
	return ticker + " <Unregistered>"
}
