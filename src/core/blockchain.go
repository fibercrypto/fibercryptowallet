package core

// CoinValueMetric enumerates all possible values of blockchain metrics
type CoinValueMetric uint32

const (
	// CoinCurrentSupply to retrieve amount of coins distributed to tenants
	CoinCurrentSupply CoinValueMetric = iota
	// CoinTotalSupply to retrieve total amount of coins
	CoinTotalSupply
)

// BlockchainStatus measure blockchain metrics in real time
type BlockchainStatus interface {
	// GeCoinValue retrieves value of a blockchain metric
	GetCoinValue(coinvalue CoinValueMetric, ticker string) (uint64, error)
	// GetLastBlock retrieves block at the tip of he block chain
	GetLastBlock() (Block, error)
	// GetNumberOfBlocks determine number of blocks in the blockchain
	GetNumberOfBlocks() (uint64, error)
}

// BlockchainAPI abstract interface for transactions management and utility functions for specific blockchain.
// The service should use the blockchain node to implement given interface.
type BlockchainAPI interface {
	// Transfer instantiates a transaction to send funds from wallet to known address account owner
	// If null address is specified for unspent wallet output then any wallet address may be chosen to build transaction
	Transfer(uxOuts []WalletOutput, to, change Address, options KeyValueStorage) (Transaction, error)
	// Spend instantiate a transaction that spends specific outputs to send to multiple destination addresses
	Spend(unspent, new []WalletOutput, change Address, options KeyValueStorage) (Transaction, error)
}
