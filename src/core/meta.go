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
