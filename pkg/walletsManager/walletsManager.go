package walletsManager

import (
	"github.com/skycoin/skycoin/src/api"
	"github.com/therecipe/qt/core"
)





func init() {
	WalletModel_QmlRegisterType2("WalletsManager", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("WalletsManager", 1, 0, "QWallet")
	AddressesModel_QmlRegisterType2("WalletsManager", 1, 0, "AddressModel")
	QAddress_QmlRegisterType2("WalletsManager", 1, 0, "QAddress")
	WalletManager_QmlRegisterType2("WalletsManager", 1, 0, "WalletManager")
	
}




//Return address of daemon node----INCOMPLETE
func nodeAddress() string {
	return "http://127.0.0.1:38391"
}

//Return username of daemon node----INCOMPLETE
func nodeUsername() string {
	return "Kid"
}

//Return password of daemon node-----INCOMPLETE
func nodePassword() string {
	return "P@ssw0rd!"
}

func newClient() *api.Client {
	c := api.NewClient(nodeAddress())
	//c.SetAuth(nodeUsername(), nodePassword())
	return c
}




type WalletManager struct {
	core.QObject
	
	_ func(seed string, label string, password string, scanN int)	`slot:"createEncryptedWallet"`
	_ func(seed string, label string, scanN int) 	`slot:"createUnencryptedWallet"`
	_ func(entropy int) string `slot:"getNewSeed"`
	_ func(seed string) int  `slot:"verifySeed"`
	_ func(id string, n int, password string)	`slot:"newWalletAddress"`
	_ func(id string, password string)	`slot:"encryptWallet"`
	_ func(id string, password string)	`slot:"decryptWallet"`

}

func initWalletManager(){
	walletM := NewWalletManager(nil)
	walletM.ConnectCreateEncryptedWallet(createEncryptedWallet)
	walletM.ConnectCreateUnencryptedWallet(createUnencryptedWallet)
	walletM.ConnectGetNewSeed(getNewSeed)
	walletM.ConnectVerifySeed(verifySeed)
	walletM.ConnectNewWalletAddress(newWalletAddress)
	walletM.ConnectEncryptWallet(encryptWallet)
}

func createEncryptedWallet(seed, label, password string, scanN int){
	c := newClient()
	_, err := c.CreateEncryptedWallet(seed, label, password, scanN)
	if err != nil{
		return
	}
}

func createUnencryptedWallet(seed, label string, scanN int){
	c := newClient()
	_, err := c.CreateUnencryptedWallet(seed, label, scanN)
	if err!= nil{
		return
	}
}

func getNewSeed(entropy int) string{
	c := newClient()
	seed, err := c.NewSeed(entropy)
	if err != nil {
		return ""
	}

	return seed
}

func verifySeed(seed string) int{
	c := newClient()
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
	c := newClient()
	_, err := c.NewWalletAddress(label, n, password)
	if err != nil{
		return
	}

}

func encryptWallet(label, password string){
	c := newClient()
	_, err := c.EncryptWallet(label, password)
	if err != nil{
		return
	}
}

func decryptWallet(label, password string){
	c := newClient()
	_, err := c.DecryptWallet(label, password)
	if err != nil {
		return
	}
}