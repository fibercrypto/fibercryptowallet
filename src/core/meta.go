package core

type CoinValueKey uint32

const (
	CoinCurrentSupply CoinValueKey = iota
	CoinTotalSupply
)

type BlockchainStatus interface {
	GetCoinValue(coinvalue CoinValueKey, ticker string) (uint64, error)
	GetLastBlock() (Block, error)
	GetNumberOfBlocks() (uint64, error)
}
