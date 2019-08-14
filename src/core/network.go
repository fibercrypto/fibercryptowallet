package core

type PEX interface {
	GetTxnPool() TransactionIterator
	GetConnections()
	Broadcast(txn Transaction)
}

type PexNodeIterator interface {
	Value() PexNode
	Next() bool
	HasNext() bool
}

type NetworkSet interface {
	ListNetworks() PexNodeIterator
}

type PexNode interface {
	GetIp() string
	GetPort() uint16
	GetBlockHeight() uint64
	IsTrusted() bool
	GetLastSeenIn() int64
	GetLastSeenOut() int64
}

