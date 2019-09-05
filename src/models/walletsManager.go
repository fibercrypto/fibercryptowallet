package models

import (
	"strconv"

	skycoin "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"

	//"github.com/fibercrypto/FiberCryptoWallet/src/models/history"
	qtcore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type WalletManager struct {
	qtcore.QObject
	WalletEnv     core.WalletEnv
	SeedGenerator core.SeedGenerator

	_ func()                                                               `constructor:"init"`
	_ func(seed string, label string, password string, scanN int) *QWallet `slot:"createEncryptedWallet"`
	_ func(seed string, label string, scanN int) *QWallet                  `slot:"createUnencryptedWallet"`
	_ func(entropy int) string                                             `slot:"getNewSeed"`
	_ func(seed string) int                                                `slot:"verifySeed"`
	_ func(id string, n int, password string)                              `slot:"newWalletAddress"`
	_ func(id string, password string)                                     `slot:"encryptWallet"`
	_ func(id string, password string)                                     `slot:"decryptWallet"`
	_ func() []*QWallet                                                    `slot:"getWallets"`
	_ func(id string) []*QAddress                                          `slot:"getAddresses"`
	_ func(wltId, destinationAddress, amount string)                       `slot:"sendTo"`
	//_ func() *history.Tes                                                  `slot:"getAddressesWithWallets"`
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
	walletM.ConnectSendTo(walletM.sendTo)
	altManager := core.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}

	walletM.WalletEnv = walletsEnvs[0]

	walletM.SeedGenerator = new(skycoin.SeedService)

}

func (walletM *WalletManager) sendTo(wltId, destinationAddress, amount string) {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
	addr := &GenericAdddress{destinationAddress}
	skyF, _ := strconv.ParseFloat(amount, 64)
	sky := uint64(skyF * 1e6)
	// TODO --> Add get password method
	wlt.Transfer(addr, sky, "")
}

func (walletM *WalletManager) createEncryptedWallet(seed, label, password string, scanN int) *QWallet {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, true, pwd, scanN)
	if err != nil {
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
		return nil
	}
	return fromWalletToQWallet(wlt, false)

}

func (walletM *WalletManager) getNewSeed(entropy int) string {
	seed, err := walletM.SeedGenerator.GenerateMnemonic(entropy)
	if err != nil {
		return ""
	}
	return seed
}

func (walletM *WalletManager) verifySeed(seed string) int {
	ok, err := walletM.SeedGenerator.VerifyMnemonic(seed)
	if err != nil {
		return 0
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

func (walletM *WalletManager) newWalletAddress(id string, n int, password string) {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wltEntrieslen := 0
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		return
	}
	for it.Next() {
		wltEntrieslen++
	}
	wlt.GenAddresses(core.AccountAddress, uint32(wltEntrieslen), uint32(n), pwd)
}

func (walletM *WalletManager) getWallets() []*QWallet {

	qwallets := make([]*QWallet, 0)
	it := walletM.WalletEnv.GetWalletSet().ListWallets()

	for it.Next() {

		encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(it.Value().GetId())
		if err != nil {
			continue
		}
		if encrypted {
			qw := fromWalletToQWallet(it.Value(), true)
			qwallets = append(qwallets, qw)
		} else {
			qw := fromWalletToQWallet(it.Value(), false)
			qwallets = append(qwallets, qw)
		}

	}

	return qwallets

}

func (walletM *WalletManager) getAddresses(Id string) []*QAddress {

	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(Id)
	qaddresses := make([]*QAddress, 0)
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		return nil
	}
	for it.Next() {
		addr := it.Value()
		qaddress := NewQAddress(nil)
		qaddress.SetAddress(addr.String())
		qaddress.SetMarked(0)
		sky, err := addr.GetCryptoAccount().GetBalance("SKY")
		if err != nil {

			continue
		}
		qaddress.SetAddressSky(sky)
		coinH, err := addr.GetCryptoAccount().GetBalance("SKYCH")
		if err != nil {

			continue
		}
		qaddress.SetAddressCoinHours(coinH)
		qml.QQmlEngine_SetObjectOwnership(qaddress, qml.QQmlEngine__CppOwnership)

		qaddresses = append(qaddresses, qaddress)

	}

	return qaddresses
}

//func (walletM *WalletManager) getAddressesWithWallets() *history.Tes {
//	response := make(map[string]string, 0)
//	it := walletM.WalletEnv.GetWalletSet().ListWallets()
//	for it.Next() {
//		wlt := it.Value()
//		addrs, _ := wlt.GetLoadedAddresses()
//
//		for addrs.Next() {
//			response[addrs.Value().String()] = wlt.GetId()
//		}
//
//	}
//	respon := history.NewTes(nil)
//	respon.Data = response
//	return respon
//}

func fromWalletToQWallet(wlt core.Wallet, isEncrypted bool) *QWallet {

	qwallet := NewQWallet(nil)
	qwallet.SetName(wlt.GetLabel())

	qwallet.SetFileName(wlt.GetId())

	qwallet.SetEncryptionEnabled(0)
	if isEncrypted {
		qwallet.SetEncryptionEnabled(1)
	}

	bl, err := wlt.GetCryptoAccount().GetBalance("Sky")
	if err != nil {
		bl = 0
	}
	qwallet.SetSky(bl)

	bl, err = wlt.GetCryptoAccount().GetBalance("CoinHour")
	if err != nil {
		bl = 0
	}
	qwallet.SetCoinHours(bl)

	return qwallet
}

type GenericAdddress struct {
	addr string
}

func (ga *GenericAdddress) IsBip32() bool {
	return true
}

func (ga *GenericAdddress) String() string {
	return ga.addr
}

func (ga *GenericAdddress) GetCryptoAccount() core.CryptoAccount {
	return nil
}
