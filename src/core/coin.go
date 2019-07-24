package core

type Transaction interface {
	GetInputs() []TransactionInput
	GetOutputs() []TransactionOutput
	GetId() string
}

type TransactionIterator interface {
	Value() Transaction
	Next() bool
	HasNext() bool
}

type TransactionInput interface {
	GetId() string
	IsSpent() bool
}

type TransactionInputIterator interface {
	Value() TransactionInput
	Next() bool
	HasNext() bool
}

type TransactionOutput interface {
	GetId() string
	GetAddress() Address
	GetCoins(ticker string) uint64
	GetSpentInput() TransactionInput
}

type TransactionOutputIterator interface {
	Value() TransactionOutput
	Next() bool
	HasNext() bool
}
