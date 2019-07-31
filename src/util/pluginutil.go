package util

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
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
