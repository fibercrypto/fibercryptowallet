package skycoin

import (
	"github.com/FiberCrypto/FiberCryptoWallet/src/core"
)

/*
SkycoinTransaction
*/
type SkycoinTransaction struct{}

func (txn *SkycoinTransaction) SupportedAssets() []string {
	return []string{"SKY", "SKYCH"}
}

func (txn *SkycoinTransaction) GetTimestamp() core.TransactionTimestamp {
	return 1551626612
}

func (txn *SkycoinTransaction) GetStatus() core.TransactionStatus {
	return core.TXN_STATUS_CONFIRMED
}

func (txn *SkycoinTransaction) GetInputs() []core.TransactionInput {
	return []core.TransactionInput{
		&SkycoinTransactionInput{},
		&SkycoinTransactionInput{},
		&SkycoinTransactionInput{},
	}
}

func (txn *SkycoinTransaction) GetOutputs() []core.TransactionOutput {
	return []core.TransactionOutput{
		&SkycoinTransactionOutput{},
		&SkycoinTransactionOutput{},
		&SkycoinTransactionOutput{},
	}
}

func (txn *SkycoinTransaction) GetId() string {
	return "32dae4040c49bfdca006e887fe8212a4f6a1b3a03665c21595f8f385dd9aefa8"
}

func (txn *SkycoinTransaction) ComputeFee(ticker string) uint64 {
	if ticker == "SKYCH" {
		return 12345
	}
	return 0
}

/**
 * SkycoinTransactionIterator
 */
type SkycoinTransactionIterator struct {
	index int
	value *SkycoinTransaction
}

func (iter *SkycoinTransactionIterator) Value() core.Transaction {
	if iter.index == 0 && iter.value == nil {
		iter.value = &SkycoinTransaction{}
	}
	return iter.value
}

func (iter *SkycoinTransactionIterator) Next() bool {
	if iter.index < 3 {
		iter.value = &SkycoinTransaction{}
		iter.index++
		return true
	}
	return false
}

func (iter *SkycoinTransactionIterator) HasNext() bool {
	return iter.index < 3
}

type SkycoinTransactionInput struct {
}

func (in *SkycoinTransactionInput) GetId() string {
	return "f1a61f49cef012e4822b314ca6657d66fdbe3c4d46110a079052a064b9a51e66"
}

func (in *SkycoinTransactionInput) IsSpent() bool {
	return true
}

func (in *SkycoinTransactionInput) GetSpentOutput() core.TransactionOutput {
	return &SkycoinTransactionOutput{}
}

/**
 * SkycoinTransactionInputIterator
 */
type SkycoinTransactionInputIterator struct {
	index int
	value *SkycoinTransactionInput
}

func (iter *SkycoinTransactionInputIterator) Value() core.TransactionInput {
	if iter.index == 0 && iter.value == nil {
		iter.value = &SkycoinTransactionInput{}
	}
	return iter.value
}

func (iter *SkycoinTransactionInputIterator) Next() bool {
	if iter.index < 3 {
		iter.value = &SkycoinTransactionInput{}
		iter.index++
		return true
	}
	return false
}

func (iter *SkycoinTransactionInputIterator) HasNext() bool {
	return iter.index < 3
}

/**
 * SkycoinTransactionOutput
 */
type SkycoinTransactionOutput struct {
}

func (*SkycoinTransactionOutput) GetId() string {
	return "249iEtdhniFppBTpkzq7syoiBaydLi6USnU"
}

func (*SkycoinTransactionOutput) GetAddress() core.Address {
	return &SkycoinAddress{}
}

func (*SkycoinTransactionOutput) GetCoins(ticker string) uint64 {
	if ticker == "SKY" {
		return 20000000
	}
	if ticker == "SKYCH" {
		return 10000000
	}
	return 0
}

func (*SkycoinTransactionOutput) GetSpentInput() core.TransactionInput {
	return &SkycoinTransactionInput{}
}

/**
 * SkycoinTransactionOutputIterator
 */
type SkycoinTransactionOutputIterator struct {
	index int
	value *SkycoinTransactionOutput
}

func (iter *SkycoinTransactionOutputIterator) Value() core.TransactionOutput {
	if iter.index == 0 && iter.value == nil {
		iter.value = &SkycoinTransactionOutput{}
	}
	return iter.value
}

func (iter *SkycoinTransactionOutputIterator) Next() bool {
	if iter.index < 3 {
		iter.value = &SkycoinTransactionOutput{}
		iter.index++
		return true
	}
	return false
}

func (iter *SkycoinTransactionOutputIterator) HasNext() bool {
	return iter.index < 3
}
