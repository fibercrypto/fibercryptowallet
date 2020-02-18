package hardware

import (
	"github.com/chebyrash/promise"
	hardware_wallet "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
)

type SkyWalletHelper struct {
	di hardware_wallet.DeviceInteraction
}

func NewSkyWalletHelper() hardware_wallet.DeviceHelper{
	return &SkyWalletHelper{di: SkyWltInteractionInstance()}
}

func (dev *SkyWalletHelper) FirstAddress(walletType string) *promise.Promise {
	prm := dev.di.AddressGen(1, 0, false, walletType)
	return prm.Then(func(data interface{}) interface{} {
		addresses := data.([]string)
		return addresses[0]
	}).Catch(func(err error) error {
		return err
	})
}

func (dev *SkyWalletHelper) DeviceMatch(wlt core.Wallet) *promise.Promise {
	matcher := func(firstAddr string) bool {
		addrs := wlt.GenAddresses(core.AccountAddress, 0, 1, nil)
		if addrs.Next() {
			addr := addrs.Value()
			return addr.String() == firstAddr
		}
		return false
	}
	prm := dev.FirstAddress(skyWallet.WalletTypeDeterministic)
	return prm.Then(func(data interface{}) interface{} {
		return matcher(data.(string))
	}).Then(func(data interface{}) interface{} {
		if !data.(bool) {
			return dev.FirstAddress(skyWallet.WalletTypeBip44)
		}
		return true
	}).Then(func(data interface{}) interface{} {
		if val, ok := data.(bool); ok {
			return val
		}
		return matcher(data.(string))
	}).Catch(func(err error) error {
		return err
	})
}

func (dev *SkyWalletHelper) ShouldBeSecured() *promise.Promise {
	return dev.di.Features().Then(func(data interface{}) interface{} {
		features := data.(messages.Features)
		if !*features.PinProtection {
			return true
		}
		if *features.NeedsBackup {
			return true
		}
		return false
	}).Catch(func(err error) error {
		return err
	})
}

func (dev *SkyWalletHelper) ShouldBeInitialized() *promise.Promise {
	return dev.di.Features().Then(func(data interface{}) interface{} {
		features := data.(messages.Features)
		if !*features.Initialized {
			return true
		}
		return false
	}).Catch(func(err error) error {
		return err
	})
}
