package core

// CryptoAccount defines the contract for objects owning coins used in transactions
type CryptoAccount interface {
	// GetBalance retrieves total number of coins for asset represented by ticker that may be spent by this account
	// Expect a non null error object if and only if an error is detected
	GetBalance(ticker string) (uint64, error)
	// ListAssets to enumerate the tickers of all assets supported by this account
	ListAssets() []string
	// ScanUnspentOutputs to determine the outputs that can participate in a transaction
	// without incurring in double spending
	ScanUnspentOutputs() TransactionOutputIterator
	// ListTransactions to show account history
	ListTransactions() TransactionIterator
	// ListPendingTransactions to obtain details of transactions pending for confirmation in the memory
	ListPendingTransactions() (TransactionIterator, error)
}
