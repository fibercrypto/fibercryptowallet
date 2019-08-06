package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

/*
SkycoinTransaction
*/
type SkycoinTransaction struct {
	timeStamp uint64
	status    core.TransactionStatus
	inputs    []SkycoinTransactionInput
	outputs   []SkycoinTransactionOutput
	fee       uint64
	id        string
}

func (txn *SkycoinTransaction) SupportedAssets() []string {
	return []string{Sky, CoinHour}
}

func (txn *SkycoinTransaction) GetTimestamp() core.TransactionTimestamp {
	return core.TransactionTimestamp(txn.timeStamp)
}

func (txn *SkycoinTransaction) GetStatus() core.TransactionStatus {
	if txn.status == core.TXN_STATUS_CONFIRMED {
		return txn.status
	}

	c := util.NewClient()
	txnU, err := c.Transaction(txn.id)
	if txnU.Status.Confirmed {
		txn.status = core.TXN_STATUS_CONFIRMED
		return txn.status
	}
	txn.status = core.TXN_STATUS_PENDING
	return txn.status
}

func (txn *SkycoinTransaction) GetInputs() []core.TransactionInput {
	return txn.inputs
}

func (txn *SkycoinTransaction) GetOutputs() []core.TransactionOutput {
	return txn.outputs
}

func (txn *SkycoinTransaction) GetId() string {
	return txn.id
}

func (txn *SkycoinTransaction) ComputeFee(ticker string) uint64 {
	if ticker == CoinHour {
		return txn.fee
	}
	return 0
}

/**
 * SkycoinTransactionIterator
 */
type SkycoinTransactionIterator struct {
	current int
	data    []core.Transaction
}

func (iter *SkycoinTransactionIterator) Value() core.Transaction {
	return iter.data[iter.current]
}

func (iter *SkycoinTransactionIterator) Next() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func (iter *SkycoinTransactionIterator) HasNext() bool {
	if (iter.current + 1) >= len(iter.data) {
		return false
	}
	return true
}

func NewSkycoinTransactionIterator(txns []core.Transaction) *SkycoinTransactionIterator {
	return &SkycoinTransactionIterator{data: txns, current: -1}
}

type SkycoinTransactionInput struct {
	id          string
	spentOutput *SkycoinTransactionOutput
}

func (in *SkycoinTransactionInput) GetId() string {
	return in.id
}

func (in *SkycoinTransactionInput) GetSpentOutput() core.TransactionOutput {
	return in.spentOutput
}

/**
 * SkycoinTransactionInputIterator
 */
type SkycoinTransactionInputIterator struct {
	current int
	data    []*SkycoinTransactionInput
}

func (iter *SkycoinTransactionInputIterator) Value() core.TransactionInput {
	return iter.data[iter.current]
}

func (iter *SkycoinTransactionInputIterator) Next() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func (iter *SkycoinTransactionInputIterator) HasNext() bool {
	if (iter.current + 1) >= len(iter.data) {
		return false
	}
	return true
}

func NewSkycoinTransactioninputIterator(ins []*SkycoinTransactionInput) *SkycoinTransactionInputIterator {
	return &SkycoinTransactionInputIterator{data: ins, current: -1}
}

/**
 * SkycoinTransactionOutput
 */
type SkycoinTransactionOutput struct {
	id              string
	amountSky       uint64
	amountCoinHours uint64
	address         SkycoinAddress
	isSpent         bool
}

func (out *SkycoinTransactionOutput) GetId() string {
	return out.id
}

func (out *SkycoinTransactionOutput) GetAddress() core.Address {
	return out.address
}

func (out *SkycoinTransactionOutput) GetCoins(ticker string) uint64 {
	if ticker == Sky {
		return out.amountSky
	}
	if ticker == CoinHour {
		return out.amountCoinHours
	}
	return 0
}

func (out *SkycoinTransactionOutput) IsSpent() bool {
	return out.isSpent
}

/**
 * SkycoinTransactionOutputIterator
 */
type SkycoinTransactionOutputIterator struct {
	current int
	data    []*SkycoinTransactionOutput
}

func (iter *SkycoinTransactionOutputIterator) Value() core.TransactionOutput {
	return iter.data[iter.current]
}

func (iter *SkycoinTransactionOutputIterator) Next() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func (iter *SkycoinTransactionOutputIterator) HasNext() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func NewSkycoinTransactionOutputIterator(outs []*SkycoinTransactionOutput) *SkycoinTransactionOutputIterator {
	return &SkycoinTransactionOutputIterator{data: outs, current: -1}
}
