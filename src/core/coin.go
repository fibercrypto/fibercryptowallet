package core

import (
	"github.com/skycoin/skycoin/src/cipher"
)

type Timestamp uint64
type TransactionStatus uint32

const (
	TXN_STATUS_CREATED TransactionStatus = iota
	TXN_STATUS_PENDING
	TXN_STATUS_CONFIRMED
)

type Transaction interface {
	SupportedAssets() []string
	GetTimestamp() Timestamp
	GetStatus() TransactionStatus
	GetInputs() []TransactionInput
	GetOutputs() []TransactionOutput
	GetId() string
	ComputeFee(ticker string) uint64
}

type TransactionIterator interface {
	Value() Transaction
	Next() bool
	HasNext() bool
}

type TransactionInput interface {
	GetId() string
	GetSpentOutput() TransactionOutput
}

type TransactionInputIterator interface {
	Value() TransactionInput
	Next() bool
	HasNext() bool
}

type TransactionOutput interface {
	GetId() string
	IsSpent() bool
	GetAddress() Address
	GetCoins(ticker string) uint64
}

type TransactionOutputIterator interface {
	Value() TransactionOutput
	Next() bool
	HasNext() bool
}

type Block interface {
	GetHash() ([]byte, error)
	GetPrevHash() ([]byte, error)
	GetVersion() uint32
	GetTime() Timestamp
	GetHeight() uint64
	GetFee(ticker string) (uint64, error)
	IsGenesisBlock() (bool, error)
}
