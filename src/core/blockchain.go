package core

type LastBlockState struct {
	NumberOfBlocks uint64
	TimeStamp      TransactionTimestamp
	HashLastBlock  string
}

type SkycoinBlockchain interface {
	GetBalance(ticker string) (uint64, error)
	LastBlockStatus() (LastBlockState, error)
	SetCacheTimeout(time uint64)
}
