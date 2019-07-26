package walletsManager

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/blockchain/wallet"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/wallets"
	qtcore "github.com/therecipe/qt/core"
)

func init() {
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")
}

type WalletManager struct {
	qtcore.QObject
	WalletEnv     core.WalletEnv
	SeedGenerator core.SeedGenerator

	_ func()                                                                       `constructor:"init"`
	_ func(seed string, label string, password string, scanN int) *wallets.QWallet `slot:"createEncryptedWallet"`
	_ func(seed string, label string, scanN int) *wallets.QWallet                  `slot:"createUnencryptedWallet"`
	_ func(entropy int) string                                                     `slot:"getNewSeed"`
	_ func(seed string) int                                                        `slot:"verifySeed"`
	//_ func(id string, n int, password string)                                      `slot:"newWalletAddress"`
	_ func(id string, password string) `slot:"encryptWallet"`
	_ func(id string, password string) `slot:"decryptWallet"`
}

func (walletM *WalletManager) init() {
	walletM.ConnectCreateEncryptedWallet(walletM.createEncryptedWallet)
	walletM.ConnectCreateUnencryptedWallet(walletM.createUnencryptedWallet)
	walletM.ConnectGetNewSeed(walletM.getNewSeed)
	walletM.ConnectVerifySeed(walletM.verifySeed)
	//walletM.ConnectNewWalletAddress(newWalletAddress)
	walletM.ConnectEncryptWallet(walletM.encryptWallet)
	walletM.ConnectDecryptWallet(walletM.decryptWallet)

	walletM.WalletEnv = new(wallet.WalletNode)
	walletM.SeedGenerator = new(wallet.SeedService)

}

func (walletM *WalletManager) createEncryptedWallet(seed, label, password string, scanN int) *wallets.QWallet {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, true, pwd, scanN)
	if err != nil {
		return nil
	}

	return fromWalletToQWallet(wlt)
}

func (walletM *WalletManager) createUnencryptedWallet(seed, label string, scanN int) *wallets.QWallet {
	pwd := func(message string) (string, error) {
		return nil, nil
	}

	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, false, pwd, scanN)
	if err != nil {
		return err
	}
	return fromWalletToQWallet(wlt)
}

func (walletM *WalletManager) getNewSeed(entropy int) string {
	seed, err := walletM.SeedGenerator.GenerateMnemonic(entropy)
	if err != nil {
		return nil
	}
	return seed
}

func (walletM *WalletManager) verifySeed(seed string) int {
	ok, err := walletM.SeedGenerator.VerifyMnemonic(seed)
	if err != nil {
		return nil
	}
	if ok {
		return 1
	}
	return 0

}

func (walletM *WalletManager) encryptWallet(id, password string) {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Encrypt(id, pwd)
}

func (walletM *WalletManager) decryptWallet(id, password string) {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Decrypt(id, pwd)
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
