package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin"
	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

var logWalletManager = logging.MustGetLogger("modelsWalletManager")

type WalletManager struct {
	qtCore.QObject
	WalletEnv     core.WalletEnv
	SeedGenerator core.SeedGenerator

	_ func()                                                               `constructor:"init"`
	_ func(seed string, label string, password string, scanN int) *QWallet `slot:"createEncryptedWallet"`
	_ func(seed string, label string, scanN int) *QWallet                  `slot:"createUnencryptedWallet"`
	_ func(entropy int) string                                             `slot:"getNewSeed"`
	_ func(seed string) int                                                `slot:"verifySeed"`
	_ func(id string, n int, password string)                              `slot:"newWalletAddress"`
	_ func(id string, password string) int                                 `slot:"encryptWallet"`
	_ func(id string, password string) int                                 `slot:"decryptWallet"`
	_ func() []*QWallet                                                    `slot:"getWallets"`
	_ func(id string) []*QAddress                                          `slot:"getAddresses"`
	_ func(id, label string) *QWallet                                      `slot:"editWallet"`
}

func (walletM *WalletManager) init() {
	qml.QQmlEngine_SetObjectOwnership(walletM, qml.QQmlEngine__CppOwnership)
	walletM.ConnectCreateEncryptedWallet(walletM.createEncryptedWallet)
	walletM.ConnectCreateUnencryptedWallet(walletM.createUnencryptedWallet)
	walletM.ConnectGetNewSeed(walletM.getNewSeed)
	walletM.ConnectVerifySeed(walletM.verifySeed)
	walletM.ConnectNewWalletAddress(walletM.newWalletAddress)
	walletM.ConnectEncryptWallet(walletM.encryptWallet)
	walletM.ConnectDecryptWallet(walletM.decryptWallet)
	walletM.ConnectGetWallets(walletM.getWallets)
	walletM.ConnectGetAddresses(walletM.getAddresses)
	altManager := core.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}

	walletM.WalletEnv = walletsEnvs[0]

	walletM.SeedGenerator = new(sky.SeedService)

}

func (walletM *WalletManager) createEncryptedWallet(seed, label, password string, scanN int) *QWallet {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, true, pwd, scanN)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't create encrypted wallet")
		return nil
	}

	return fromWalletToQWallet(wlt, true)

}

func (walletM *WalletManager) createUnencryptedWallet(seed, label string, scanN int) *QWallet {
	pwd := func(message string) (string, error) {
		return "", nil
	}

	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, false, pwd, scanN)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't create unencrypted wallet")
		return nil
	}
	return fromWalletToQWallet(wlt, false)

}

func (walletM *WalletManager) getNewSeed(entropy int) string {
	seed, err := walletM.SeedGenerator.GenerateMnemonic(entropy)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get new seed")
		return ""
	}
	return seed
}

func (walletM *WalletManager) verifySeed(seed string) int {
	logWalletManager.Info("Verifying seed")
	ok, err := walletM.SeedGenerator.VerifyMnemonic(seed)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't verify seed")
		return 0
	}
	if ok {
		return 1
	}
	return 0

}

func (walletM *WalletManager) encryptWallet(id, password string) int {
	logWalletManager.Info("Encrypting wallet")
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Encrypt(id, pwd)
	ret, err := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't create encrypted wallets")
	}
	if ret {
		return 1
	}
	return 0
}

func (walletM *WalletManager) decryptWallet(id, password string) int {
	logWalletManager.Info("Decrypt wallet")
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Decrypt(id, pwd)
	ret, err := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't decrypt wallet")
	}
	if ret {
		return 1
	}
	return 0
}

func (walletM *WalletManager) newWalletAddress(id string, n int, password string) {
	logWalletManager.Info("Creating new wallet addresses")
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wltEntriesLen := 0
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't load addresses")
		return
	}
	for it.Next() {
		wltEntriesLen++
	}
	wlt.GenAddresses(core.AccountAddress, uint32(wltEntriesLen), uint32(n), pwd)
}

func (walletM *WalletManager) getWallets() []*QWallet {
	logWalletManager.Info("Getting wallets")

	qWallets := make([]*QWallet, 0)
	it := walletM.WalletEnv.GetWalletSet().ListWallets()

	if it == nil {
		logWalletManager.WithError(nil).Error("Couldn't load wallets")
		return qWallets

	}

	for it.Next() {

		encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(it.Value().GetId())
		if err != nil {
			logWalletManager.WithError(err).Error("Couldn't get wallets")
			return qWallets
		}
		if encrypted {
			qw := fromWalletToQWallet(it.Value(), true)
			qWallets = append(qWallets, qw)
		} else {
			qw := fromWalletToQWallet(it.Value(), false)
			qWallets = append(qWallets, qw)
		}

	}

	return qWallets

}

func (walletM *WalletManager) editWallet(id, label string) *QWallet {
	logWalletManager.Info("Editing wallet")

	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	wlt.SetLabel(label)
	wlt = walletM.WalletEnv.GetWalletSet().GetWallet(id)
	encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(wlt.GetId())
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't edit wallet")
		return nil
	}
	qWallet := fromWalletToQWallet(wlt, encrypted)
	return qWallet
}

func (walletM *WalletManager) getAddresses(Id string) []*QAddress {
	logWalletManager.Info("Get addresses")

	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(Id)
	qAddresses := make([]*QAddress, 0)
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get loaded addresses")
		return nil
	}
	for it.Next() {
		addr := it.Value()
		qaddress := NewQAddress(nil)
		qml.QQmlEngine_SetObjectOwnership(qaddress, qml.QQmlEngine__CppOwnership)
		qaddress.SetAddress(addr.String())
		qaddress.SetMarked(0)
		skyBalance, err := addr.GetCryptoAccount().GetBalance("SKY")
		if err != nil {
			logWalletManager.WithError(err).Error("Couldn't load wallet " + addr.String())
			continue
		}
		accuracy, _ := util.AltcoinQuotient("SKY")
		qaddress.SetAddressSky(util.FormatCoins(skyBalance, accuracy))
		coinH, err := addr.GetCryptoAccount().GetBalance("SKYCH")
		if err != nil {
			logWalletManager.WithError(err).Error("Couldn't get cryptographic account")
			continue
		}
		accuracy, err = util.AltcoinQuotient("SKYCH")
		if err != nil {
			logWalletManager.WithError(err).Error("Couldn't get Altcoin quotient")
			continue
		}
		qaddress.SetAddressCoinHours(util.FormatCoins(coinH, accuracy))
		qml.QQmlEngine_SetObjectOwnership(qaddress, qml.QQmlEngine__CppOwnership)

		qAddresses = append(qAddresses, qaddress)

	}

	return qAddresses
}

func fromWalletToQWallet(wlt core.Wallet, isEncrypted bool) *QWallet {

	qWallet := NewQWallet(nil)
	qml.QQmlEngine_SetObjectOwnership(qWallet, qml.QQmlEngine__CppOwnership)
	qWallet.SetName(wlt.GetLabel())

	qWallet.SetFileName(wlt.GetId())

	qWallet.SetEncryptionEnabled(0)
	if isEncrypted {
		qWallet.SetEncryptionEnabled(1)
	}

	bl, err := wlt.GetCryptoAccount().GetBalance(skycoin.SkycoinTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Skycoin balance")
		return qWallet
	}

	//TODO: report possible error
	accuracy, err := util.AltcoinQuotient(skycoin.SkycoinTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Skycoin Altcoin quotient")
		return qWallet
	}
	floatBl := float64(bl) / float64(accuracy)
	qWallet.SetSky(floatBl)

	bl, err = wlt.GetCryptoAccount().GetBalance(skycoin.CoinHoursTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Coin Hours balance")
		return qWallet
	}
	accuracy, _ = util.AltcoinQuotient(skycoin.SkycoinTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Coin Hours Altcoin quotient")
		return qWallet
	}
	qWallet.SetCoinHours(bl)

	return qWallet
}
