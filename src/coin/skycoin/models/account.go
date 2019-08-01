package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

func (addr SkycoinAddress) GetBalance(ticker string) (uint64, error) {
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
		return 0, errorTickerInvalid{ticker}
	}
}
func (addr SkycoinAddress) ListAssets() []string {
	return []string{Sky, CoinHour}
}
func (addr SkycoinAddress) ScanUnspentOutputs() core.TransactionOutputIterator {
	c := util.newClient()
	outputs, err := c.OutputsForAddresses([addr.address])
	if err != nil {
		return nil//, err
	}
	skyOutputs := make([]SkycoinTransactionOutput, 0)
	for _, o := range outputs {
		for _, unspentOutput := range o.HeadOutputs {
			so := UnspentOutputToSkycoinTransactionOutput(unspentOutput)
			so.Address = addr
			skyOutputs = append(skyOutputs, so)
		}
	}
	return NewSkycoinTransactionOutputIterator(skyOutputs)//, nil
}

func (addr SkycoinAddress) ListTransactions() core.TransactionIterator { //------TODO
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
		return 0, errorTickerInvalid{ticker}
	}

}

func (wlt Wallet) ListAssets() []string {
	return []string{Sky, CoinHour}
}

func (wlt Wallet) ScanUnspentOutputs() core.TransactionOutputIterator { //------TODO
	return nil
}

func (wlt Wallet) ListTransactions() core.TransactionIterator { //------TODO
	return nil
}

func UnspentOutputToSkycoinTransactionOutput(uo UnspentOutput) *SkycoinTransactionOutput {
	return &SkycoinTransactionOutput{
		Id: uo.Hash,
		Sky: uint64(uo.Coins), //TODO: Check if this correct i.e uint64(string)
		CoinHours: uo.Hours
	}
}

func NewSkycoinTransactionOutputIterator(outputs []SkycoinTransactionOutput) *SkycoinTransactionOutputIterator {
	return &SkycoinTransactionOutputIterator{Outputs: outputs, Current: -1}
}