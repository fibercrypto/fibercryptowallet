package core

import(
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/api"
)

type SkycoinAPI interface {
	Transaction(txid string) 									(*readable.TransactionWithStatus, error)
	Transactions(addrs []string) 								([]readable.TransactionWithStatus, error)
	TransactionVerbose(txid string) 							(*readable.TransactionWithStatusVerbose, error) 
	TransactionsVerbose(addrs []string) 						([]readable.TransactionWithStatusVerbose, error)
	UxOut(uxID string) 											(*readable.SpentOutput, error)
	PendingTransactionsVerbose() 								([]readable.UnconfirmedTransactionVerbose, error)
	CoinSupply() 												(*api.CoinSupply, error)
	LastBlocks(n uint64) 										(*readable.Blocks, error)
	BlockchainProgress() 										(*readable.BlockchainProgress, error)
	Balance(addrs []string) 									(*api.BalanceResponse, error)
	OutputsForAddresses(addrs []string) 						(*readable.UnspentOutputsSummary, error)
	Wallet(id string) 											(*api.WalletResponse, error)
	UpdateWallet(id, label string) 								error
	NewWalletAddress(id string, n int, password string)			([]string, error)
	Wallets() 													([]api.WalletResponse, error)
	CreateWallet(o api.CreateWalletOptions) 					(*api.WalletResponse, error)
	EncryptWallet(id, password string)		 					(*api.WalletResponse, error)
	DecryptWallet(id, password string) 							(*api.WalletResponse, error)
	WalletBalance(id string) 									(*api.BalanceResponse, error)
	WalletUnconfirmedTransactionsVerbose(id string)		 		(*api.UnconfirmedTxnsVerboseResponse, error)
	NetworkConnections(filters *api.NetworkConnectionsFilter) 	(*api.Connections, error)
}
