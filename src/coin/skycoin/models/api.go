package skycoin

import (
	// "testing"

	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/stretchr/testify/mock"
)

type SkycoinApiMock struct {
	mock.Mock
}

func (m *SkycoinApiMock) Create() (interface{}, error) {
	return m, nil
}

func (m *SkycoinApiMock) Transaction(txid string) (*readable.TransactionWithStatus, error) {
	args := m.Called(txid)
	return args.Get(0).(*readable.TransactionWithStatus), args.Error(1)
}
func (m *SkycoinApiMock) Transactions(addrs []string) ([]readable.TransactionWithStatus, error) {
	args := m.Called(addrs)
	return args.Get(0).([]readable.TransactionWithStatus), args.Error(1)
}
func (m *SkycoinApiMock) TransactionVerbose(txid string) (*readable.TransactionWithStatusVerbose, error) {
	args := m.Called(txid)
	return args.Get(0).(*readable.TransactionWithStatusVerbose), args.Error(1)
}
func (m *SkycoinApiMock) TransactionsVerbose(addrs []string) ([]readable.TransactionWithStatusVerbose, error) {
	args := m.Called(addrs)
	return args.Get(0).([]readable.TransactionWithStatusVerbose), args.Error(1)
}
func (m *SkycoinApiMock) UxOut(uxID string) (*readable.SpentOutput, error) {
	args := m.Called(uxID)
	return args.Get(0).(*readable.SpentOutput), args.Error(1)
}
func (m *SkycoinApiMock) PendingTransactionsVerbose() ([]readable.UnconfirmedTransactionVerbose, error) {
	args := m.Called()
	return args.Get(0).([]readable.UnconfirmedTransactionVerbose), args.Error(1)
}
func (m *SkycoinApiMock) CoinSupply() (*api.CoinSupply, error) {
	args := m.Called()
	return args.Get(0).(*api.CoinSupply), args.Error(1)
}
func (m *SkycoinApiMock) LastBlocks(n uint64) (*readable.Blocks, error) {
	args := m.Called(n)
	return args.Get(0).(*readable.Blocks), args.Error(1)
}
func (m *SkycoinApiMock) BlockchainProgress() (*readable.BlockchainProgress, error) {
	args := m.Called()
	return args.Get(0).(*readable.BlockchainProgress), args.Error(1)
}
func (m *SkycoinApiMock) Balance(addrs []string) (*api.BalanceResponse, error) {
	args := m.Called(addrs)
	return args.Get(0).(*api.BalanceResponse), args.Error(1)
}
func (m *SkycoinApiMock) OutputsForAddresses(addrs []string) (*readable.UnspentOutputsSummary, error) {
	args := m.Called(addrs)
	return args.Get(0).(*readable.UnspentOutputsSummary), args.Error(1)
}
func (m *SkycoinApiMock) Wallet(id string) (*api.WalletResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*api.WalletResponse), args.Error(1)
}
func (m *SkycoinApiMock) UpdateWallet(id, label string) error {
	args := m.Called(id, label)
	return args.Error(0)
}
func (m *SkycoinApiMock) NewWalletAddress(id string, n int, password string) ([]string, error) {
	args := m.Called(id, n, password)
	return args.Get(0).([]string), args.Error(1)
}
func (m *SkycoinApiMock) Wallets() ([]api.WalletResponse, error) {
	args := m.Called()
	return args.Get(0).([]api.WalletResponse), args.Error(1)

}
func (m *SkycoinApiMock) CreateWallet(o api.CreateWalletOptions) (*api.WalletResponse, error) {
	args := m.Called(o)
	return args.Get(0).(*api.WalletResponse), args.Error(1)
}
func (m *SkycoinApiMock) EncryptWallet(id, password string) (*api.WalletResponse, error) {
	args := m.Called(id, password)
	return args.Get(0).(*api.WalletResponse), args.Error(1)
}
func (m *SkycoinApiMock) DecryptWallet(id, password string) (*api.WalletResponse, error) {
	args := m.Called(id, password)
	return args.Get(0).(*api.WalletResponse), args.Error(1)
}
func (m *SkycoinApiMock) WalletBalance(id string) (*api.BalanceResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*api.BalanceResponse), args.Error(1)
}
func (m *SkycoinApiMock) WalletUnconfirmedTransactionsVerbose(id string) (*api.UnconfirmedTxnsVerboseResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*api.UnconfirmedTxnsVerboseResponse), args.Error(1)
}
func (m *SkycoinApiMock) NetworkConnections(filters *api.NetworkConnectionsFilter) (*api.Connections, error) {
	args := m.Called(filters)
	return args.Get(0).(*api.Connections), args.Error(1)
}
func (m *SkycoinApiMock) InjectTransaction(txn *coin.Transaction) (string, error) {
	args := m.Called(txn)
	return args.Get(0).(string), args.Error(1)
}
func (m *SkycoinApiMock) WalletSignTransaction(req api.WalletSignTransactionRequest) (*api.CreateTransactionResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*api.CreateTransactionResponse), args.Error(1)
}
func (m *SkycoinApiMock) WalletCreateTransaction(req api.WalletCreateTransactionRequest) (*api.CreateTransactionResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*api.CreateTransactionResponse), args.Error(1)
}
func (m *SkycoinApiMock) CreateTransaction(req api.CreateTransactionRequest) (*api.CreateTransactionResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*api.CreateTransactionResponse), args.Error(1)
}

func (m *SkycoinApiMock) InjectEncodedTransaction(rawTxn string) (string, error) {
	args := m.Called(rawTxn)
	return args.Get(0).(string), args.Error(1)
}
