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
}

func addWalletToQWalletFunction(func())
