package models

import (
	"fmt"
	"strconv"

	"github.com/fibercrypto/FiberCryptoWallet/src/util"

	"github.com/therecipe/qt/qml"

	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"

	qtcore "github.com/therecipe/qt/core"
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
	_ func(id string, password string) int                                 `slot:"encryptWallet"`
	_ func(id string, password string) int                                 `slot:"decryptWallet"`
	_ func() []*QWallet                                                    `slot:"getWallets"`
	_ func(id string) []*QAddress                                          `slot:"getAddresses"`
	_ func(wltId, destinationAddress, amount, password string)             `slot:"sendTo"`
	_ func(id, label string) *QWallet                                      `slot:"editWallet"`
	_ func(wltId, address string) []*QOutput                               `slot:"getOutputs"`
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
	walletM.ConnectGetOutputs(walletM.getOutputs)
	altManager := core.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}

	walletM.WalletEnv = walletsEnvs[0]

	walletM.SeedGenerator = new(sky.SeedService)

}

func (walletM *WalletManager) getOutputs(wltId, address string) []*QOutput {
	fmt.Println(wltId)
	fmt.Println(address)
	addrIter, err := walletM.WalletEnv.GetWalletSet().GetWallet(wltId).GetLoadedAddresses()
	if err != nil {
		return nil
	}
	outs := make([]*QOutput, 0)
	var addr core.Address
	fmt.Println(0)
	for addrIter.Next() {
		if addrIter.Value().String() == address {
			addr = addrIter.Value()
			break
		}
	}
	fmt.Println(addr.String())
	outsIter := addr.GetCryptoAccount().ScanUnspentOutputs()
	fmt.Println("1")
	for outsIter.Next() {
		fmt.Println("2")
		qout := NewQOutput(nil)
		qout.SetOutputID(outsIter.Value().GetId())
		skyV := outsIter.Value().GetCoins(sky.Sky)
		quotient, err := util.GetCoinValue(sky.Sky)
		if err != nil {
			return nil
		}
		sSky := util.FormatCoins(skyV, quotient)
		qout.SetAddressSky(sSky)
		ch := outsIter.Value().GetCoins(sky.CoinHour)
		quotient, err = util.GetCoinValue(sky.CoinHour)
		if err != nil {
			return nil
		}

		sCh := util.FormatCoins(ch, quotient)
		qout.SetAddressCoinHours(sCh)
		outs = append(outs, qout)
	}

	return outs
}

func (walletM *WalletManager) sendTo(wltId, destinationAddress, amount, password string) {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
	addr := &GenericAdddress{destinationAddress}
	skyF, _ := strconv.ParseFloat(amount, 64)
	sky := uint64(skyF * 1e6)

	wlt.Transfer(addr, sky, func(text string) (string, error) {
		return password, nil
	}, nil)
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

func (walletM *WalletManager) encryptWallet(id, password string) int {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Encrypt(id, pwd)
	ret, _ := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if ret {
		return 1
	}
	return 0
}

func (walletM *WalletManager) decryptWallet(id, password string) int {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Decrypt(id, pwd)
	ret, _ := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if ret {
		return 1
	}
	return 0
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

func (walletM *WalletManager) editWallet(id, label string) *QWallet {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	wlt.SetLabel(label)
	wlt = walletM.WalletEnv.GetWalletSet().GetWallet(id)
	encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(wlt.GetId())
	if err != nil {
		return nil
	}
	qwallet := fromWalletToQWallet(wlt, encrypted)
	return qwallet
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
		qml.QQmlEngine_SetObjectOwnership(qaddress, qml.QQmlEngine__CppOwnership)
		qaddress.SetAddress(addr.String())
		qaddress.SetMarked(0)
		sky, err := addr.GetCryptoAccount().GetBalance("SKY")
		if err != nil {

			continue
		}
		flSky := float64(sky / 1e6)
		qaddress.SetAddressSky(strconv.FormatFloat(flSky, 'f', -1, 64))
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

func fromWalletToQWallet(wlt core.Wallet, isEncrypted bool) *QWallet {

	qwallet := NewQWallet(nil)
	qml.QQmlEngine_SetObjectOwnership(qwallet, qml.QQmlEngine__CppOwnership)
	qwallet.SetName(wlt.GetLabel())

	qwallet.SetFileName(wlt.GetId())

	qwallet.SetEncryptionEnabled(0)
	if isEncrypted {
		qwallet.SetEncryptionEnabled(1)
	}

	bl, err := wlt.GetCryptoAccount().GetBalance(sky.Sky)
	if err != nil {
		bl = 0
	}

	floatBl := float64(bl / 1e6)
	qwallet.SetSky(floatBl)

	bl, err = wlt.GetCryptoAccount().GetBalance(sky.CoinHour)
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
