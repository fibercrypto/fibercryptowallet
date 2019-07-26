package wallet

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type WalletService struct {
}

func (wltSrv *WalletService) ListWallets() core.WalletIterator {
	c := util.NewClient()
	wlts, err := c.Wallets()
	if err != nil {
		return nil
	}
	wallets := make([]Wallet, 0)
	for _, wlt := range wlts {
		wallets = append(wallets, walletResponseToWallet(wlt))
	}
	return NewSkycoinWalletIterator(wallets)
}

func (wltSrv *WalletService) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	if IsEncrypted {
		password, _ := pwd("Enter your password")
		return wltSrv.createEncryptedWallet(seed, label, password, scanAddressesN)
	} else {
		return wltSrv.createUnencryptedWallet(seed, label, scanAddressesN)
	}
}
func (wltSrv *WalletService) createEncryptedWallet(seed, label, password string, scanN int) (core.Wallet, error) {
	c := util.NewClient()
	wltR, err := c.CreateEncryptedWallet(seed, label, password, scanN)
	if err != nil {
		return nil, err
	}

	wlt := walletResponseToWallet(*wltR)
	return &wlt, nil

}

func (wltSrv *WalletService) createUnencryptedWallet(seed, label string, scanN int) (core.Wallet, error) {
	c := util.NewClient()
	wltR, err := c.CreateUnencryptedWallet(seed, label, scanN)
	if err != nil {
		return nil, err
	}
	wlt := walletResponseToWallet(*wltR)
	return &wlt, nil
}

func (wltSrv *WalletService) Encrypt(walletName, password string) {
	c := util.NewClient()
	_, err := c.EncryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *WalletService) Decrypt(walletName, password string) {
	c := util.NewClient()
	_, err := c.DecryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *WalletService) IsEncrypted(walletName string) (bool, error) {
	c := util.NewClient()
	wlt, err := c.Wallet(walletName)
	if err != nil {
		return false, err
	}
	return wlt.Meta.Encrypted, nil
}
