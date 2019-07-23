package walletsManager

import (
	"github.com/skycoin/skycoin/src/api"
	"github.com/therecipe/qt/core"
	"github.com/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/FiberCryptoWallet/src/util"

	
)





func init() {
	WalletModel_QmlRegisterType2("WalletsManager", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("WalletsManager", 1, 0, "QWallet")
	AddressesModel_QmlRegisterType2("WalletsManager", 1, 0, "AddressModel")
	QAddress_QmlRegisterType2("WalletsManager", 1, 0, "QAddress")
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")
	
}



type WalletManager struct {
	core.QObject
	_ func ()	`constructor:"init"`
	_ func(walletM *WalletModel, seed string, label string, password string, scanN int)	`slot:"createEncryptedWallet"`
	_ func(walletM *WalletModel, seed string, label string, scanN int) 	`slot:"createUnencryptedWallet"`
	_ func(entropy int) string `slot:"getNewSeed"`
	_ func(seed string) int  `slot:"verifySeed"`
	_ func(addressM *AddressesModel, id string, n int, password string)	`slot:"newWalletAddress"`
	_ func(walletM *WalletModel, id string, password string)	`slot:"encryptWallet"`
	_ func(walletM *WalletModel, id string, password string)	`slot:"decryptWallet"`

}

func (walletM *WalletManager) init(){
	walletM.ConnectCreateEncryptedWallet(createEncryptedWallet)
	walletM.ConnectCreateUnencryptedWallet(createUnencryptedWallet)
	walletM.ConnectGetNewSeed(getNewSeed)
	walletM.ConnectVerifySeed(verifySeed)
	walletM.ConnectNewWalletAddress(newWalletAddress)
	walletM.ConnectEncryptWallet(encryptWallet)
}


func createEncryptedWallet(walletModel *WalletModel, seed, label, password string, scanN int){
	c := util.newClient()
	_, err := c.CreateEncryptedWallet(seed, label, password, scanN)
	if err != nil{
		return
	}
}

func createUnencryptedWallet(walletModel *WalletModel ,seed, label string, scanN int){
	c := util.newClient()
	wlt, err := c.CreateUnencryptedWallet(seed, label, scanN)
	if err!= nil{
		return
	}
	qwallet, err := walletResponseToQWallet(wlt)
	if err != nil{
		return
	}
	walletModel.AddWallet(qwallet)
}

func getNewSeed(entropy int) string{
	c := util.newClient()
	seed, err := c.NewSeed(entropy)
	if err != nil {
		return ""
	}
	
	return seed
}

func verifySeed(seed string) int{
	c := util.newClient()
	ok, err := c.VerifySeed(seed)
	if err != nil{
		return 0
	}
	if !ok {
		return 0
	}

	return 1
}

func newWalletAddress(addressModel *AddressesModel, label string, n int, password string){
	c := util.newClient()
	_, err := c.NewWalletAddress(label, n, password)
	if err != nil{
		return
	}

}

func encryptWallet(walletModel *WalletModel, label, password string){
	c := util.newClient()
	_, err := c.EncryptWallet(label, password)
	if err != nil{
		return
	}
}

func decryptWallet(walletModel *WalletModel, label, password string){
	c := util.newClient()
	_, err := c.DecryptWallet(label, password)
	if err != nil {
		return
	}
}