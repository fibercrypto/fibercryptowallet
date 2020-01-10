package hardware_wallet

import (
	"github.com/chebyrash/promise"
	hardware "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet/skywallet"
)

type SkyWalletHelper struct {
}

func (wlt *SkyWalletHelper) FirstAddress(walletType string) *promise.Promise {
	dev := hardware.NewSkyWalletInteraction()
	prm := dev.AddressGen(1, 0, false, walletType)
	return prm.Then(func(data interface{}) interface{} {
		addresses := data.([]string)
		return addresses[0]
	}).Catch(func(err error) error {
		return err
	})
}
