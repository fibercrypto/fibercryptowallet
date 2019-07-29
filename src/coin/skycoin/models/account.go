package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

func (addr Address) GetBalance(ticker string) (uint64, error) {
	c := util.NewClient()
	bl, err := c.Balance([]string{addr.address})
	if err != nil {
		return 0, err
	}

	if ticker == Sky {
		return bl.Confirmed.Coins, nil
	} else if ticker == CoinHour {
		return bl.Confirmed.Hours, nil
	} else {
		return 0, errorTickerInvalid{ticker, addr.ListAssets()}
	}
}
func (addr Address) ListAssets() []string {
	return []string{"Sky", "CoinHour"}
}
func (addr Address) ScanUnspentOutputs() core.TransactionOutputIterator { //------TODO
	return nil
}
func (addr Address) ListTransactions() core.TransactionIterator { //------TODO
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
		return 0, errorTickerInvalid{ticker, wlt.ListAssets()}
	}

}

func (wlt Wallet) ListAssets() []string {
	return []string{"Sky", "CoinHour"}
}

func (wlt Wallet) ScanUnspentOutputs() core.TransactionOutputIterator { //------TODO
	return nil
}

func (wlt Wallet) ListTransactions() core.TransactionIterator { //------TODO
	return nil
}
