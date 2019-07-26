package wallet

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"
)

const (
	Sky      = "Sky"
	CoinHour = "CoinHour"
)

type errorTickerInvalid struct{}

func (err errorTickerInvalid) Error() string {
	return "Ticker invalid"
}

type Wallet struct {
	Id        string
	Label     string
	CoinType  string
	Encrypted bool
}

func (wlt Wallet) GetLabel() string {
	return wlt.Label
}

func (wlt Wallet) SetLabel(name string) {
	c := util.NewClient()
	_ = c.UpdateWallet(wlt.Id, name)
}

func (wlt Wallet) GetId() string {
	return wlt.Id
}

func (wlt Wallet) Transfer(to core.Address, amount uint64) {
	return
}

func (wlt Wallet) SendFromAddress(from, to core.Address, amount uint64) {
	return
}

func (wlt Wallet) Spend(unspent, new []core.TransactionOutput) {
	return
}

func (wlt Wallet) GenChangeAddress() core.Address {
	return nil
}

func (wlt Wallet) GetBalance(ticker string) (uint64, error) {
	c := util.NewClient()
	bl, err := c.WalletBalance(wlt.Id)
	if err != nil {
		return 0, err
	}

	if ticker == Sky {
		return bl.Confirmed.Coins, nil
	} else if ticker == CoinHour {
		return bl.Confirmed.Hours, nil
	} else {
		return 0, errorTickerInvalid{}
	}

}

func (wlt Wallet) ListAssets() []string {
	return []string{"Sky", "CoinHour"}
}

func (wlt Wallet) ScanUnspentOutputs() core.TransactionOutputIterator {
	return nil
}

func (wlt Wallet) ListTransactions() core.TransactionIterator {
	return nil
}

func (wlt Wallet) GenAddresses(addrType core.AddressType, startIndex, count uint32) core.AddressIterator {
	return nil
}

func (wlt Wallet) GetCryptoAccount() core.CryptoAccount {
	return wlt
}

func newWalletAddress(label string, n int, password string) {
	c := util.NewClient()
	_, err := c.NewWalletAddress(label, n, password)
	if err != nil {
		return
	}

}

func walletResponseToWallet(wltR api.WalletResponse) Wallet {
	wlt := Wallet{}
	wlt.CoinType = wltR.Meta.Coin
	wlt.Encrypted = wltR.Meta.Encrypted
	wlt.Label = wltR.Meta.Label
	wlt.Id = wltR.Meta.Filename
	return wlt
}
