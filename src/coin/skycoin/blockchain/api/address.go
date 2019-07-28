package api

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

type Address struct {
	address string
}

func (addr Address) IsBip32() bool {
	return true
}

func (addr Address) String() string {
	return addr.address
}

func (addr Address) GetCryptoAccount() core.CryptoAccount {
	return addr
}

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
		return 0, errorTickerInvalid{}
	}
}
func (addr Address) ListAssets() []string {
	return []string{"Sky", "CoinHour"}
}
func (addr Address) ScanUnspentOutputs() core.TransactionOutputIterator {
	return nil
}
func (addr Address) ListTransactions() core.TransactionIterator {
	return nil
}
