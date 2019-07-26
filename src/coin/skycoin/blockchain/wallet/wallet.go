package wallet

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util/connection.go"
	"github.com/skycoin/skycoin/src/api"
)

type Wallet struct {
	Id        string
	Name      string
	Label     string
	CoinType  string
	Encrypted bool
}

func (wlt *Wallet) GetName() string {
	return wlt.Name
}

func (wlt *Wallet) SetName(name string) {
	c := util.NewClient()
	err := c.UpdateWallet(wlt.Id, name)
}

func (wlt *Wallet) Transfer(to Address, amount uint64) {
	return
}

func (wlt *Wallet) SpendFormAddress(from, to Address, amount uint64) {
	return
}

func (wlt *Wallet) Spend(unspent, new []TransactionOutput) {
	return
}

func (wlt *Wallet) GenChangeAddress() Address {
	return
}





//func newWalletAddress(label string, n int, password string) {
	c := util.NewClient()
	_, err := c.NewWalletAddress(label, n, password)
	if err != nil {
		return
	}

//}
//
//func walletResponseToWallet(wltR api.WalletResponse) Wallet {
	wlt := new(Wallet)
	wlt.CoinType = wltR.Meta.Coin
	wlt.Encrypted = wltR.Meta.Encrypted
	wlt.Label = wltR.Meta.Label
	wlt.Id = wltR.Meta.Filename
	return wlt
//}
//
