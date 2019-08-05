package core

type PEX interface {
	GetTxnPool() (TransactionIterator, error)
	GetConnections()
	BroadcastTxn(txn Transaction)
}
