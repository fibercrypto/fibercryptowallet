package core

type PEX interface {
	getTxnPool() TransactionIterator
	getConnections() 
	broadcast(txn Transaction)

type NetworkIterator interface {
	Value() INetwork
	Next() bool
	HasNext() bool
}

type NetworkSet interface {
	ListNetworks() NetworkIterator
}

type INetwork interface {
	GetIp() string
	GetPort() uint16
	GetBlock() uint64
	IsTrusted() bool
	GetLastSeenIn() int64
	GetLastSeenOut() int64
}

