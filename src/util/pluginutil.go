package util

import (
	"errors"
	"github.com/sirupsen/logrus"
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

// LookupSignerByUID search for signer matching given ID
func LookupSignerByUID(wlt core.Wallet, id core.UID) core.TxnSigner {
	wltSigner, isSigner := wlt.(core.TxnSigner)
	// Reference to self
	if id == core.UID("") {
		if isSigner {
			return wltSigner
		}
		return nil
	}
	// Wallet matches ID
	uid, err := wltSigner.GetSignerUID()
	if err != nil {
		logrus.WithError(err).Errorln("unable to get signer uid")
		return nil
	}
	if isSigner && uid == id {
		return wltSigner
	}
	// Lookup global signers
	return local.LoadAltcoinManager().LookupSignService(id)
}
