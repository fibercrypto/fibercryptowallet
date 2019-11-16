package skytypes

import (
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
)

// SkycoinAPI contract for Skycoin REST API clients
type SkycoinAPI interface {
	// Transaction Get transaction info by id
	Transaction(txid string) (*readable.TransactionWithStatus, error)
	// Transactions Get transactions for addresses
	Transactions(addrs []string) ([]readable.TransactionWithStatus, error)
	// TransactionVerbose Get transaction info by id. Include spent input data
	TransactionVerbose(txid string) (*readable.TransactionWithStatusVerbose, error)
	// TransactionsVerbose Get transactions for addresses. Include spent input data
	TransactionsVerbose(addrs []string) ([]readable.TransactionWithStatusVerbose, error)
	// UxOut Get uxout
	UxOut(uxID string) (*readable.SpentOutput, error)
	// PendingTransactionsVerbose Get unconfirmed transactions
	PendingTransactionsVerbose() ([]readable.UnconfirmedTransactionVerbose, error)
	// CoinSupply Determine coin supply
	CoinSupply() (*api.CoinSupply, error)
	// LastBlocks Get last N blocks
	LastBlocks(n uint64) (*readable.Blocks, error)
	// BlockchainProgress Get blockchain progress
	BlockchainProgress() (*readable.BlockchainProgress, error)
	// Balance Get balance of addresses
	Balance(addrs []string) (*api.BalanceResponse, error)
	// OutputsForAddresses Get historical unspent outputs for an address
	OutputsForAddresses(addrs []string) (*readable.UnspentOutputsSummary, error)
	// Wallet Get wallet
	Wallet(id string) (*api.WalletResponse, error)
	// UpdateWallet Change wallet label
	UpdateWallet(id, label string) error
	// NewWalletAddress Generate new address in wallet
	NewWalletAddress(id string, n int, password string) ([]string, error)
	// Wallets Get wallets
	Wallets() ([]api.WalletResponse, error)
	// CreateWallet Create wallet
	CreateWallet(o api.CreateWalletOptions) (*api.WalletResponse, error)
	// EncryptWallet Encrypt wallet
	EncryptWallet(id, password string) (*api.WalletResponse, error)
	// DecryptWallet Decrypt wallet
	DecryptWallet(id, password string) (*api.WalletResponse, error)
	// WalletBalance Get wallet balance
	WalletBalance(id string) (*api.BalanceResponse, error)
	// WalletUnconfirmedTransactionsVerbose Get unconfirmed transactions of a wallet
	WalletUnconfirmedTransactionsVerbose(id string) (*api.UnconfirmedTxnsVerboseResponse, error)
	// NetworkConnections Get a list of all connections
	NetworkConnections(filters *api.NetworkConnectionsFilter) (*api.Connections, error)
	// InjectTransaction Inject transaction
	InjectTransaction(txn *coin.Transaction) (string, error)
	// InjectTransaction Inject raw transaction
	InjectEncodedTransaction(rawTxn string) (string, error)
	// WalletSignTransaction Sign transaction
	WalletSignTransaction(req api.WalletSignTransactionRequest) (*api.CreateTransactionResponse, error)
	// WalletCreateTransaction Create transaction from wallet addresses
	WalletCreateTransaction(req api.WalletCreateTransactionRequest) (*api.CreateTransactionResponse, error)
	// CreateTransaction Create transaction from unspent outputs or addresses
	CreateTransaction(req api.CreateTransactionRequest) (*api.CreateTransactionResponse, error)
}
