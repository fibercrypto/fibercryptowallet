package util

import (
	"errors"
	"math"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
)

func AltcoinCaption(ticker string) string {
	if info, isRegistered := local.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return info.Name
	}
	return ticker + " <Unregistered>"
}

func AltcoinQuotient(ticker string) (uint64, error) {
	if info, isRegistered := local.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return uint64(math.Pow(float64(10), float64(info.Accuracy))), nil
	}
	return uint64(0), errors.New(ticker + " <Unregistered>")
}

func RegisterAltcoin(p core.AltcoinPlugin) {
	local.LoadAltcoinManager().RegisterPlugin(p)
}

// LoadAltcoinManager duplicated to avoid recursive import loop
func LoadAltcoinManager() core.AltcoinManager {
	return local.LoadAltcoinManager()
}
