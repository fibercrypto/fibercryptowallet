package util

import (
	"errors"
	"math"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

func AltcoinCaption(ticker string) string {
	if info, isRegistered := core.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return info.Name
	}
	return ticker + " <Unregistered>"
}

func AltcoinQuotient(ticker string) (uint64, error) {
	if info, isRegistered := core.LoadAltcoinManager().DescribeAltcoin(ticker); isRegistered {
		return uint64(math.Pow(float64(10), float64(info.Accuracy))), nil
	}
	return uint64(0), errors.New(ticker + " <Unregistered>")
}

func RegisterAltcoin(p core.AltcoinPlugin) {
	core.LoadAltcoinManager().RegisterPlugin(p)
}
