package core

type CryptoAccount interface {
	GetBalance(ticker string) (uint64, error)
	ListAssets() []string
	ScanUnspentOutputs() TransactionOutputIterator
	ListTransactions() TransactionIterator
	ListPendingTransactions() (TransactionIterator, error)
}
