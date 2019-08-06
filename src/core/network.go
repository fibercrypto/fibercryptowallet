package core

type PEX interface {
	getTxnPool() TransactionIterator
	getConnections() 
	broadcast(txn Transaction)
}
