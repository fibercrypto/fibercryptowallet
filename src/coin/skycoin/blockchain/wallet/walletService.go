package wallet

import "github.com/fibercrypto/FiberCryptoWallet/src/core"

type WalletService struct {
}

func (wltSrv *WalletService) ListWallets() WalletIterator {
	c := util.NewClient()
	wlts, err = c.Wallets()
	if err != nil {
		return nil
	}
	wallets = make([]Wallet, 0)
	for _, wlt := range wlts {
		wallets = append(wallets, walletResponseToWallet(wlt))
	}
	return NewSkycoinWalletIterator(wallets)
}

func (wltSrv *WalletService) CreateWallet(name string, seed string, IsEncrypted bool, pwd PasswordReader, scanAddressesN int) (core.Wallet, error) {
	if IsEncrypted {
		password := pwd("Enter your password")
		return createEncryptedWallet(seed, name, password, scanAddressesN)
	} else {
		return createUnencryptedWallet(seed, name, scanAddressesN)
	}
}
func (wltSrv *WalletService) createEncryptedWallet(seed, label, password string, scanN int) (core.Wallet, error) {
	c := util.NewClient()
	wltR, err := c.CreateEncryptedWallet(seed, label, password, scanN)
	if err != nil {
		return nil, err
	}

	wlt := walletResponseToWallet(wltR)
	return wlt, nil

}

func (wltSrv *WalletService) createUnencryptedWallet(seed, label string, scanN int) (core.Wallet, error) {
	c := util.NewClient()
	_, err := c.CreateUnencryptedWallet(seed, label, scanN)
	if err != nil {
		return nil
	}

}

func (wltSrv *WalletService) Encrypt(walletName, password string) {
	c := util.NewClient()
	_, err := c.EncryptWallet(label, password)
	if err != nil {
		return
	}
}

func (wltSrv *WalletService) Decrypt(walletName, password string) {
	c := util.NewClient()
	_, err := c.DecryptWallet(label, password)
	if err != nil {
		return
	}
}

func (wltSrv *WalletService) IsEncrypted(walletName string) bool {
	c := util.NewClient()
	wlt, err := c.Wallet(walletName)
	if err != nil {
		return err
	}
	return wlt.Meta.Encrypted
}

func (wltSrv *WalletService) GenerateMnemonic(entropy int) (string, error) {
	c := util.NewClient()
	seed, err := c.NewSeed(entropy)
	if err != nil {
		return nil, err
	}

	return seed, nil
}

func (wltSrv *WalletService) VerifyMnemonic(seed string) (bool, error) {
	c := util.NewClient()
	ok, err := c.VerifySeed(seed)
	if err != nil {
		return nil, err
	}
	return ok, nil
}
