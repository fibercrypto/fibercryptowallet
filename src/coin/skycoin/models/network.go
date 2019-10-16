package skycoin

import (
	"errors"
	"fmt"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
)

var logNetwork = logging.MustGetLogger("Skycoin network")

const (
	PoolSection = "skycoin"
)

type SkycoinAPI interface {
	Transaction(txid string) (*readable.TransactionWithStatus, error)
	Transactions(addrs []string) ([]readable.TransactionWithStatus, error)
	TransactionVerbose(txid string) (*readable.TransactionWithStatusVerbose, error)
	TransactionsVerbose(addrs []string) ([]readable.TransactionWithStatusVerbose, error)
	UxOut(uxID string) (*readable.SpentOutput, error)
	PendingTransactionsVerbose() ([]readable.UnconfirmedTransactionVerbose, error)
	CoinSupply() (*api.CoinSupply, error)
	LastBlocks(n uint64) (*readable.Blocks, error)
	BlockchainProgress() (*readable.BlockchainProgress, error)
	Balance(addrs []string) (*api.BalanceResponse, error)
	OutputsForAddresses(addrs []string) (*readable.UnspentOutputsSummary, error)
	Wallet(id string) (*api.WalletResponse, error)
	UpdateWallet(id, label string) error
	NewWalletAddress(id string, n int, password string) ([]string, error)
	Wallets() ([]api.WalletResponse, error)
	CreateWallet(o api.CreateWalletOptions) (*api.WalletResponse, error)
	EncryptWallet(id, password string) (*api.WalletResponse, error)
	DecryptWallet(id, password string) (*api.WalletResponse, error)
	WalletBalance(id string) (*api.BalanceResponse, error)
	WalletUnconfirmedTransactionsVerbose(id string) (*api.UnconfirmedTxnsVerboseResponse, error)
	NetworkConnections(filters *api.NetworkConnectionsFilter) (*api.Connections, error)
	InjectTransaction(txn *coin.Transaction) (string, error)
	WalletSignTransaction(req api.WalletSignTransactionRequest) (*api.CreateTransactionResponse, error)
	WalletCreateTransaction(req api.WalletCreateTransactionRequest) (*api.CreateTransactionResponse, error)
	CreateTransaction(req api.CreateTransactionRequest) (*api.CreateTransactionResponse, error)
}

type SkycoinConnectionFactory struct {
	url string
}

func (cf *SkycoinConnectionFactory) Create() (interface{}, error) {

	return api.NewClient(cf.url), nil
}

func NewSkycoinConnectionFactory(url string) *SkycoinConnectionFactory {

	return &SkycoinConnectionFactory{
		url: url,
	}
}

type SkycoinApiClient struct {
	*api.Client
	pool core.MultiPoolSection
}

// nolint megacheck TODO: This functions is not used
func (sc *SkycoinApiClient) returnToPool() {
	sc.pool.Put(sc.Client)
}

func NewSkycoinApiClient(section string) (SkycoinAPI, error) {
	logNetwork.Info("Creating Skycoin api client")
	mpool := core.GetMultiPool()
	pool, err := mpool.GetSection(section)
	if err != nil {
		return nil, err
	}

	obj := pool.Get()

	if err != nil {
		for _, ok := err.(core.NotAvailableObjectsError); ok; _, ok = err.(core.NotAvailableObjectsError) {
			if err == nil {
				break
			}
		}
		return nil, err
	}

	skyApi, ok := obj.(*api.Client)
	if !ok {
		return nil, fmt.Errorf("There is not propers client in %s pool", section)
	}
	return &SkycoinApiClient{
		Client: skyApi,
		pool:   pool,
	}, nil
}

func ReturnSkycoinClient(obj SkycoinAPI) {
	poolObj, ok := obj.(*SkycoinApiClient)
	if !ok {
		return
	}
	poolObj.pool.Put(poolObj.Client)
}

func NewSkycoinPEX(poolSection string) *SkycoinPEX {
	logNetwork.Info("Creating new Skycoin PEX")
	return &SkycoinPEX{poolSection}
}

//Implements PEX interface
type SkycoinPEX struct {
	poolSection string
}

func (spex *SkycoinPEX) GetConnections() (core.PexNodeSet, error) {
	//TODO
	return nil, nil
}

func (spex *SkycoinPEX) BroadcastTxn(txn core.Transaction) error {
	logNetwork.Info("Broadcasting transaction")
	unTxn, ok := txn.(*SkycoinUninjectedTransaction)
	if !ok {
		return errors.New("Invalid Transaction")
	}
	c, err := NewSkycoinApiClient(spex.poolSection)
	if err != nil {
		return err
	}
	defer ReturnSkycoinClient(c)
	_, err = c.InjectTransaction(unTxn.txn)
	if err != nil {
		return err
	}

	return nil
}

func (spex *SkycoinPEX) GetTxnPool() (core.TransactionIterator, error) {
	logNetwork.Info("Getting transaction pool")
	c, err := NewSkycoinApiClient(spex.poolSection)
	if err != nil {
		return nil, err
	}
	defer ReturnSkycoinClient(c)
	txns, err2 := c.PendingTransactionsVerbose()
	if err2 != nil {
		return nil, err2
	}
	skycoinTxns := make([]core.Transaction, 0)
	for _, txn := range txns {
		skycoinTxns = append(skycoinTxns, &SkycoinPendingTransaction{Transaction: txn})
	}
	return NewSkycoinTransactionIterator(skycoinTxns), nil
}
