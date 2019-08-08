package core

type LastBlockState struct {
	NumberOfBlocks uint64
	TimeStamp      TransactionTimestamp
	HashLastBlock  string
}

type BlockchainStatus interface {
	GetBalance(ticker string) (uint64, error)
	GetLastBlockStatus() (LastBlockState, error)
	SetCacheTimeout(time uint64)
}
