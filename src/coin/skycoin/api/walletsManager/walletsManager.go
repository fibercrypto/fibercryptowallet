package walletsManager

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/wallets"
	qtcore "github.com/therecipe/qt/core"
)

func init() {
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")
}

type WalletManager struct {
	qtcore.QObject
	core.WalletEnv
	core.SeedGenerator

	_ func()                                                                       `constructor:"init"`
	_ func(seed string, label string, password string, scanN int) *wallets.QWallet `slot:"createEncryptedWallet"`
	_ func(seed string, label string, scanN int) *wallets.QWallet                  `slot:"createUnencryptedWallet"`
	_ func(entropy int) string                                                     `slot:"getNewSeed"`
	_ func(seed string) int                                                        `slot:"verifySeed"`
	_ func(id string, n int, password string)                                      `slot:"newWalletAddress"`
	_ func(id string, password string)                                             `slot:"encryptWallet"`
	_ func(id string, password string)                                             `slot:"decryptWallet"`
}

func (walletM *WalletManager) init() {
	walletM.ConnectCreateEncryptedWallet(createEncryptedWallet)
	walletM.ConnectCreateUnencryptedWallet(createUnencryptedWallet)
	walletM.ConnectGetNewSeed(getNewSeed)
	walletM.ConnectVerifySeed(verifySeed)
	walletM.ConnectNewWalletAddress(newWalletAddress)
	walletM.ConnectEncryptWallet(encryptWallet)
	walletM.ConnectDescryptWallet(descryptWallet)
}

func createEncryptedWallet(seed, label, password string, scanN int) *wallets.QWallet {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wlt, err := core.WalletEnv.GetWalletSet().CreateWallet(label, seed, true, pwd, scanN)
	if err != nil {
		return nil
	}

	return fromWalletToQWallet(wlt)
}

func createUnencryptedWallet(seed, label string, scanN int) *wallets.QWallet {
	pwd := func(message string) (string, error) {
		return nil, nil
	}

	wlt, err := core.WalletEnv.GetWalletSet().CreateWallet(label, seed, false, pwd, scanN)
	if err != nil {
		return err
	}
	return fromWalletToQWallet(wlt)
}

func getNewSeed(entropy int) string {
	seed, err := core.SeedGenerator.GenerateMnemonic(entropy)
	if err != nil {
		return nil
	}
	return seed
}

func verifySeed(seed string) int {
	ok, err := core.SeedGenerator.VerifyMnemonic(seed)
	if err != nil {
		return nil
	}
	if ok {
		return 1
	}
	return 0

}

func encryptWallet(id, password string) {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	core.WalletEnv.GetStorage().Encrypt(id, pwd)
}

func decryptWallet(id, password string) {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	core.WalletEnv.GetStorage().Decrypt(id, pwd)
}

func fromWalletToQWallet(wlt *walletWithBalance) {
	qwallet = wallets.NewQWallet(nil)
	qwallet.SetName(wlt.GetLabel())
	qwallet.SetFileName(wlt.GetId())
	qwallet.SetEncryptionEnabled(wlt.IsEncrypted())
	qwallet.SetSky(wlt.GetBalance("Sky"))
	qwallet.SetCoinHours(wlt.GetBalance("CoinHours"))
	return qwallet
}

type walletWithBalance interface {
	core.Wallet
	core.CryptoAccount
}
