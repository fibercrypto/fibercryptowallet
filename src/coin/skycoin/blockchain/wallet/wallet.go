package wallet

import (
	util "github.com/fibercrypto/FiberCryptoWallet/src/util/connection.go"
	"github.com/skycoin/skycoin/src/api"
	"github.com/fibercrypt/FiberCryptoWallet/src/core"
)

const (
	Sky = "Sky"
	CoinHour = "CoinHour"
)

type errorTickerInvalid{}
func (err errorTickerInvalid) Error() string{
	return "Ticker invalid"
}

type Wallet struct {
	Id        string
	Label     string
	CoinType  string
	Encrypted bool
}

func (wlt *Wallet) GetLabel() string {
	return wlt.Label
}

func (wlt *Wallet) SetLabel(name string) {
	c := util.NewClient()
	err := c.UpdateWallet(wlt.Id, name)
}

func (wlt *Wallet) GetId() string {
	return wlt.Id
}

func (wlt *Wallet) Transfer(to Address, amount uint64) {
	return
}

func (wlt *Wallet) SpendFormAddress(from, to core.Address, amount uint64) {
	return
}

func (wlt *Wallet) Spend(unspent, new []core.TransactionOutput) {
	return
}

func (wlt *Wallet) GenChangeAddress() Address {
	return
}

func (wlt *Wallet) GetBalance(ticker string) uint64{
	c := util.NewClient()
	bl, err := c.WalletBalance(wlt.Id)
	if err != nil {
		return err
	}

	if ticker == Sky{
		return bl.Confirmed.Coins
	} else if ticker == CoinHour{
		return bl.Confirmed.Hours
	} else{
		return errorTickerInvalid{}
	}
	

}

func (wlt *Wallet) ListAssets() []string {
	return []string{"Sky", "CoinHour"}
}

func (wlt *Wallet) ScanUnspentOutpus() core.TransanctionOutputIterator{
	return
}

func (wlt *Wallet) ListTransactions() core.TransactionIterator{
	return
}

func newWalletAddress(label string, n int, password string) {
	c := util.NewClient()
	_, err := c.NewWalletAddress(label, n, password)
	if err != nil {
		return
	}

}

func walletResponseToWallet(wltR api.WalletResponse) Wallet {
	wlt := new(Wallet)
	wlt.CoinType = wltR.Meta.Coin
	wlt.Encrypted = wltR.Meta.Encrypted
	wlt.Label = wltR.Meta.Label
	wlt.Id = wltR.Meta.Filename
	return wlt
}
