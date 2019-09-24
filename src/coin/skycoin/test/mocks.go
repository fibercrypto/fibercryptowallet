package test

import(
	// "testing"
	// "github.com/stretchr/testify/mock"
	core "github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

type SkycoinConnectionFactoryMock struct {
	node *api.Client
}

func (m *SkycoinConnectionFactoryMock) Create() (core.PooledObject, error) {
	return &SkycoinApiMock{node: m.node}, nil
}

func (m *SkycoinConnectionFactoryMock) SetNode(client *api.Client) {
	m.node = client
}

type SkycoinApiMock struct {
	// mock.Mock
	node *api.Client
}
func (m *SkycoinApiMock) Transaction(txid string) (*readable.TransactionWithStatus, error){
	return m.node.Transaction(txid)
}
func (m *SkycoinApiMock) Transactions(addrs []string) ([]readable.TransactionWithStatus, error){
	return m.node.Transactions(addrs)
}
func (m *SkycoinApiMock) TransactionVerbose(txid string) (*readable.TransactionWithStatusVerbose, error){
	return m.node.TransactionVerbose(txid)
}
func (m *SkycoinApiMock) TransactionsVerbose(addrs []string) ([]readable.TransactionWithStatusVerbose, error){
	return m.node.TransactionsVerbose(addrs)
}
func (m *SkycoinApiMock) UxOut(uxID string) (*readable.SpentOutput, error){
	return m.node.UxOut(uxID)
}
func (m *SkycoinApiMock) PendingTransactionsVerbose() ([]readable.UnconfirmedTransactionVerbose, error){
	return m.node.PendingTransactionsVerbose()
}
func (m *SkycoinApiMock) CoinSupply() (*api.CoinSupply, error){
	return m.node.CoinSupply()
}
func (m *SkycoinApiMock) LastBlocks(n uint64) (*readable.Blocks, error){
	return m.node.LastBlocks(n)
}
func (m *SkycoinApiMock) BlockchainProgress() (*readable.BlockchainProgress, error){
	return m.node.BlockchainProgress()
}
func (m *SkycoinApiMock) Balance(addrs []string) (*api.BalanceResponse, error){
	return m.node.Balance(addrs)
}
func (m *SkycoinApiMock) OutputsForAddresses(addrs []string) (*readable.UnspentOutputsSummary, error){
	return m.node.OutputsForAddresses(addrs)
}
func (m *SkycoinApiMock) Wallet(id string) (*api.WalletResponse, error){
	return m.node.Wallet(id)
}
func (m *SkycoinApiMock) UpdateWallet(id, label string) error{
	return m.node.UpdateWallet(id, label)
}
func (m *SkycoinApiMock) NewWalletAddress(id string, n int, password string) ([]string, error){
	return m.node.NewWalletAddress(id, n, password)
}
func (m *SkycoinApiMock) Wallets() ([]api.WalletResponse, error){
	return m.node.Wallets()
}
func (m *SkycoinApiMock) CreateWallet(o api.CreateWalletOptions) (*api.WalletResponse, error){
	return m.node.CreateWallet(o)
}
func (m *SkycoinApiMock) EncryptWallet(id, password string) (*api.WalletResponse, error){
	return m.node.EncryptWallet(id, password)
}
func (m *SkycoinApiMock) DecryptWallet(id, password string) (*api.WalletResponse, error){
	return m.node.DecryptWallet(id, password)
}
func (m *SkycoinApiMock) WalletBalance(id string) (*api.BalanceResponse, error){
	return m.node.WalletBalance(id)
}
func (m *SkycoinApiMock) WalletUnconfirmedTransactionsVerbose(id string) (*api.UnconfirmedTxnsVerboseResponse, error){
	return m.node.WalletUnconfirmedTransactionsVerbose(id)
}
func (m *SkycoinApiMock) NetworkConnections(filters *api.NetworkConnectionsFilter) (*api.Connections, error){
	return m.node.NetworkConnections(filters)
}
