package core

type PEX interface {
	getTxnPool() TransactionIterator
	getConnections() 
	broadcastTxn(txn Transaction)
}
