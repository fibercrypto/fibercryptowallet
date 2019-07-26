package core

type CryptoAccount interface {
	GetBalance(ticker string) uint64
	ListAssets() []string
	ScanUnspentOutputs() TransactionOutputIterator
	ListTransactions() TransactionIterator
}
