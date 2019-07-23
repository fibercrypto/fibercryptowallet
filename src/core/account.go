package core

type CryptoAccount interface {
	GetBalance(ticker string)
	ListAssets() []string
	ScanUnspentOutputs() TransactionOutputIterator
	ListTransactions() TransactionIterator
}
