package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
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
	if ticker == "Sky" {
		return sto.Sky
	}
	return sto.CoinHours
}

type SkycoinTransactionOutputIterator struct { //Implements TransactionOutputIterator interface
	Current interface
	Outputs []SkycoinTransactionOutput
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