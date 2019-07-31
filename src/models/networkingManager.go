package models

import (
"github.com/fibercrypto/F/src/coin/skycoin/models"
"github.com/fibercrypto/F/src/core"
qtcore "github.com/therecipe/qt/core"
)

type NetworkingManager struct {
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
	_ func() []*QNetworking                                                    `slot:"getNetworks"`
	_ func(id string) []*QAddress                                          `slot:"getAddresses"`
}

func (net *NetworkingManager) init() {
	net.ConnectCreateEncryptedWallet(walletM.createEncryptedWallet)
	net.ConnectCreateUnencryptedWallet(walletM.createUnencryptedWallet)
	net.ConnectGetNewSeed(walletM.getNewSeed)
	net.ConnectVerifySeed(walletM.verifySeed)
	net.ConnectNewWalletAddress(walletM.newWalletAddress)
	net.ConnectEncryptWallet(walletM.encryptWallet)
	net.ConnectDecryptWallet(walletM.decryptWallet)
	net.ConnectGetWallets(walletM.getWallets)
	net.ConnectGetAddresses(walletM.getAddresses)

	net.WalletEnv = new(models.WalletNode) //Set the nodeAddress field to WalletNode type
	net.SeedGenerator = new(models.SeedService)

}

func (walletM *NetworkingManager) createEncryptedWallet(seed, label, password string, scanN int) *QWallet {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, true, pwd, scanN)
	if err != nil {
		return nil
	}

	return fromWalletToQWallet(wlt, true)

}

func (walletM *NetworkingManager) createUnencryptedWallet(seed, label string, scanN int) *QWallet {
	pwd := func(message string) (string, error) {
		return "", nil
	}

	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, false, pwd, scanN)
	if err != nil {
		return nil
	}
	return fromWalletToQWallet(wlt, false)

}

func (walletM *NetworkingManager) getNewSeed(entropy int) string {
	seed, err := walletM.SeedGenerator.GenerateMnemonic(entropy)
	if err != nil {
		return ""
	}
	return seed
}

func (walletM *NetworkingManager) verifySeed(seed string) int {
	ok, err := walletM.SeedGenerator.VerifyMnemonic(seed)
	if err != nil {
		return 0
	}
	if ok {
		return 1
	}
	return 0

}

func (walletM *NetworkingManager) encryptWallet(id, password string) {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Encrypt(id, pwd)
}

func (walletM *NetworkingManager) decryptWallet(id, password string) {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Decrypt(id, pwd)
}

func (walletM *NetworkingManager) newWalletAddress(id string, n int, password string) {
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

func (walletM *NetworkingManager) getNetworks() []*QNetworking {
	networks := make([]*QNetworking, 0)
	it := walletM.WalletEnv.GetWalletSet().ListWallets()
	for it.Next() {
		encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(it.Value().GetId())
		if err != nil {
			continue
		}
		if encrypted {
			qwallets = append(qwallets, fromWalletToQWallet(it.Value(), true))
		} else {
			qwallets = append(qwallets, fromWalletToQWallet(it.Value(), false))
		}

	}
	return qwallets
}
