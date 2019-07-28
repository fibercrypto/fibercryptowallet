package wallet

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
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

func (wlt Wallet) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {
	c := util.NewClient()
	password, _ := pwd("Insert password")
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil
	}
	addresses := make([]Address, 0)
	for _, entry := range wltR.Entries[startIndex:int(min(len(wltR.Entries), int(startIndex+count)))] {
		addresses = append(addresses, walletEntryToAddress(entry))
	}
	//Checking if all the neccesary addresses exists
	if uint32(len(wltR.Entries)) < (startIndex + count) {
		difference := (startIndex + count) - uint32(len(wltR.Entries))
		newAddrs, err := c.NewWalletAddress(wlt.Id, int(difference), password)
		if err != nil {
			return nil
		}
		for _, addr := range newAddrs {
			addresses = append(addresses, Address{addr})
		}
	}

	return NewSkycoinAddressIterator(addresses)

}

func (wlt Wallet) GetCryptoAccount() core.CryptoAccount {
	return wlt
}

func (wlt Wallet) GetLoadedAddresses() (core.AddressIterator, error) {
	c := util.NewClient()
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil, err
	}
	addresses := make([]Address, 0)
	for _, entry := range wltR.Entries {
		addresses = append(addresses, walletEntryToAddress(entry))
	}

	return NewSkycoinAddressIterator(addresses), nil
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

func walletEntryToAddress(wltE readable.WalletEntry) Address {
	return Address{wltE.Address}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
