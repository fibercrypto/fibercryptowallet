package walletsManager

import (
	"github.com/therecipe/qt/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/wallets"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"

	
)




func init() {
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")	
}



type WalletManager struct {
	core.QObject
	_ func ()	`constructor:"init"`
	_ func(walletM *wallets.WalletModel, seed string, label string, password string, scanN int)	`slot:"createEncryptedWallet"`
	_ func(walletM *wallets.WalletModel, seed string, label string, scanN int) 	`slot:"createUnencryptedWallet"`
	_ func(entropy int) string `slot:"getNewSeed"`
	_ func(seed string) int  `slot:"verifySeed"`
	_ func(addressM *wallets.AddressesModel, id string, n int, password string)	`slot:"newWalletAddress"`
	_ func(walletM *wallets.WalletModel, id string, password string)	`slot:"encryptWallet"`
	_ func(walletM *wallets.WalletModel, id string, password string)	`slot:"decryptWallet"`

}

func (walletM *WalletManager) init(){
	walletM.ConnectCreateEncryptedWallet(createEncryptedWallet)
	walletM.ConnectCreateUnencryptedWallet(createUnencryptedWallet)
	walletM.ConnectGetNewSeed(getNewSeed)
	walletM.ConnectVerifySeed(verifySeed)
	walletM.ConnectNewWalletAddress(newWalletAddress)
	walletM.ConnectEncryptWallet(encryptWallet)
}


func createEncryptedWallet(walletModel *wallets.WalletModel, seed, label, password string, scanN int){
	c := util.NewClient()
	_, err := c.CreateEncryptedWallet(seed, label, password, scanN)
	if err != nil{
		return
	}
}

func createUnencryptedWallet(walletModel *wallets.WalletModel ,seed, label string, scanN int){
	c := util.NewClient()
	wlt, err := c.CreateUnencryptedWallet(seed, label, scanN)
	if err!= nil{
		return
	}
	qwallet, err := wallets.WalletResponseToQWallet(wlt)
	if err != nil{
		return
	}
	walletModel.AddWallet(qwallet)
}

func getNewSeed(entropy int) string{
	c := util.NewClient()
	seed, err := c.NewSeed(entropy)
	if err != nil {
		return ""
	}
	
	return seed
}

func verifySeed(seed string) int{
	c := util.NewClient()
	ok, err := c.VerifySeed(seed)
	if err != nil{
		return 0
	}
	if !ok {
		return 0
	}

	return 1
}

func newWalletAddress(addressesM *wallets.AddressesModel, label string, n int, password string){
	c := util.NewClient()
	_, err := c.NewWalletAddress(label, n, password)
	if err != nil{
		return
	}

}

func encryptWallet(walletM *wallets.WalletModel, label, password string){
	c := util.NewClient()
	_, err := c.EncryptWallet(label, password)
	if err != nil{
		return
	}
}

func decryptWallet(walletM *wallets.WalletModel, label, password string){
	c := util.NewClient()
	_, err := c.DecryptWallet(label, password)
	if err != nil {
		return
	}
}