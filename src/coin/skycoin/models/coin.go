package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

type SkycoinTransactionOutput struct { //Implements TransactionOutput interface
	Id 		  string
	Address   SkycoinAddress
	Sky 	  uint64
	CoinHours uint64
}

func (sto *SkycoinTransactionOutput) GetId() string {
	return sto.Id
}

func (sto *SkycoinTransactionOutput) GetAddress() core.Address {
	return sto.Address
}

func (sto *SkycoinTransactionOutput) GetCoins(ticker string) uint64 {
	if ticker == "SKY" {
		return sto.Sky
	}
	return sto.CoinHours
}

func (sto *SkycoinTransactionOutput) IsSpent() bool {
	return true
}

type SkycoinTransactionOutputIterator struct { //Implements TransactionOutputIterator interface
	Current int
	Outputs []core.TransactionOutput
}

func (it *SkycoinTransactionOutputIterator) Value() core.TransactionOutput {
	return it.Outputs[it.Current]
}

func (it *SkycoinTransactionOutputIterator) Next() bool {
	if it.HasNext() {
		it.Current++
		return true
	}
	return false
}

func (it *SkycoinTransactionOutputIterator) HasNext() bool {
	return (it.Current + 1) < len(it.Outputs)
}

func NewSkycoinTransactionOutputIterator(outputs []core.TransactionOutput) *SkycoinTransactionOutputIterator {
	return &SkycoinTransactionOutputIterator{Outputs: outputs, Current: -1}
}