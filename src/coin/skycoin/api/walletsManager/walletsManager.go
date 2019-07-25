package walletsManager

import (
	"github.com/therecipe/qt/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/wallets"

	
)




func init() {
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")	
}



type WalletManager struct {
	core.QObject
	_ func ()	`constructor:"init"`
	_ func(seed string, label string, password string, scanN int) *wallets.QWallet	`slot:"createEncryptedWallet"`
	_ func(seed string, label string, scanN int) *wallets.QWallet	`slot:"createUnencryptedWallet"`
	_ func(entropy int) string `slot:"getNewSeed"`
	_ func(seed string) int  `slot:"verifySeed"`
	_ func(id string, n int, password string)	`slot:"newWalletAddress"`
	_ func(id string, password string)	`slot:"encryptWallet"`
	_ func(id string, password string)	`slot:"decryptWallet"`

}

func (walletM *WalletManager) init(){
	walletM.ConnectCreateEncryptedWallet(createEncryptedWallet)
	walletM.ConnectCreateUnencryptedWallet(createUnencryptedWallet)
	walletM.ConnectGetNewSeed(getNewSeed)
	walletM.ConnectVerifySeed(verifySeed)
	walletM.ConnectNewWalletAddress(newWalletAddress)
	walletM.ConnectEncryptWallet(encryptWallet)
}


func createEncryptedWallet(seed, label, password string, scanN int) {
	c := util.NewClient()
	_, err := c.CreateEncryptedWallet(seed, label, password, scanN)
	if err != nil{
		return
	}
	//wltr, err := wallets.WalletResponseToQWallet(wlt)
	//if err != nil{
	//	return
	//}
	//return qwallet
}

func createUnencryptedWallet(seed, label string, scanN int) *util.Wallet{
	c := util.NewClient()
	_, err := c.CreateUnencryptedWallet(seed, label, scanN)
	if err!= nil{
		return nil
	}
	//qwallet, err := wallets.WalletResponseToQWallet(wlt)
	//if err != nil{
	//	return nil
	//}
	//return qwallet
	w := new(util.Wallet)
	w.Sky = 20
	return w
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

func newWalletAddress(label string, n int, password string){
	c := util.NewClient()
	_, err := c.NewWalletAddress(label, n, password)
	if err != nil{
		return
	}

}

func encryptWallet(label, password string){
	c := util.NewClient()
	_, err := c.EncryptWallet(label, password)
	if err != nil{
		return
	}
}

func decryptWallet(label, password string){
	c := util.NewClient()
	_, err := c.DecryptWallet(label, password)
	if err != nil {
		return
	}
}