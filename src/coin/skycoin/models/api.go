package skycoin

import(
	// "testing"
	"github.com/stretchr/testify/mock"
	core "github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

type SkycoinApiMock struct {
	mock.Mock
	Node *api.Client
}

func (m *SkycoinApiMock) Create() (core.PooledObject, error) {
	return m, nil
}

func (m *SkycoinApiMock) Transaction(txid string) (*readable.TransactionWithStatus, error){
	args := m.Called(txid)
	return args.Get(0).(*readable.TransactionWithStatus), args.Error(1)
}
func (m *SkycoinApiMock) Transactions(addrs []string) ([]readable.TransactionWithStatus, error){
	return m.Node.Transactions(addrs)
}
func (m *SkycoinApiMock) TransactionVerbose(txid string) (*readable.TransactionWithStatusVerbose, error){
	args := m.Called(txid)
	return args.Get(0).(*readable.TransactionWithStatusVerbose), args.Error(1)
}
func (m *SkycoinApiMock) TransactionsVerbose(addrs []string) ([]readable.TransactionWithStatusVerbose, error){
	return m.Node.TransactionsVerbose(addrs)
}
func (m *SkycoinApiMock) UxOut(uxID string) (*readable.SpentOutput, error){
	args := m.Called(uxID)
	return args.Get(0).(*readable.SpentOutput), args.Error(1)
}
func (m *SkycoinApiMock) PendingTransactionsVerbose() ([]readable.UnconfirmedTransactionVerbose, error){
	args := m.Called();
	return args.Get(0).([]readable.UnconfirmedTransactionVerbose), args.Error(1)
}
func (m *SkycoinApiMock) CoinSupply() (*api.CoinSupply, error){
	return m.Node.CoinSupply()
}
func (m *SkycoinApiMock) LastBlocks(n uint64) (*readable.Blocks, error){
	return m.Node.LastBlocks(n)
}
func (m *SkycoinApiMock) BlockchainProgress() (*readable.BlockchainProgress, error){
	return m.Node.BlockchainProgress()
}
func (m *SkycoinApiMock) Balance(addrs []string) (*api.BalanceResponse, error){
	return m.Node.Balance(addrs)
}
func (m *SkycoinApiMock) OutputsForAddresses(addrs []string) (*readable.UnspentOutputsSummary, error){
	return m.Node.OutputsForAddresses(addrs)
}
func (m *SkycoinApiMock) Wallet(id string) (*api.WalletResponse, error){
	return m.Node.Wallet(id)
}
func (m *SkycoinApiMock) UpdateWallet(id, label string) error{
	return m.Node.UpdateWallet(id, label)
}
func (m *SkycoinApiMock) NewWalletAddress(id string, n int, password string) ([]string, error){
	return m.Node.NewWalletAddress(id, n, password)
}
func (m *SkycoinApiMock) Wallets() ([]api.WalletResponse, error){
	return m.Node.Wallets()
}
func (m *SkycoinApiMock) CreateWallet(o api.CreateWalletOptions) (*api.WalletResponse, error){
	return m.Node.CreateWallet(o)
}
func (m *SkycoinApiMock) EncryptWallet(id, password string) (*api.WalletResponse, error){
	return m.Node.EncryptWallet(id, password)
}
func (m *SkycoinApiMock) DecryptWallet(id, password string) (*api.WalletResponse, error){
	return m.Node.DecryptWallet(id, password)
}
func (m *SkycoinApiMock) WalletBalance(id string) (*api.BalanceResponse, error){
	return m.Node.WalletBalance(id)
}
func (m *SkycoinApiMock) WalletUnconfirmedTransactionsVerbose(id string) (*api.UnconfirmedTxnsVerboseResponse, error){
	args := m.Called(id);
	return args.Get(0).(*api.UnconfirmedTxnsVerboseResponse), args.Error(1)
}
func (m *SkycoinApiMock) NetworkConnections(filters *api.NetworkConnectionsFilter) (*api.Connections, error){
	return m.Node.NetworkConnections(filters)
}