package skycoin

import (
	"encoding/hex"
	"strings"

	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/skytypes"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logNetwork = logging.MustGetLogger("Skycoin network")

const (
	PoolSection = "skycoin"
)

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
	skytypes.SkycoinAPI
	pool core.MultiPoolSection
}

// nolint megacheck TODO: This functions is not used
func (sc *SkycoinApiClient) returnToPool() {
	sc.pool.Put(sc.SkycoinAPI)
}

func NewSkycoinApiClient(section string) (skytypes.SkycoinAPI, error) {
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

	skyApi, ok := obj.(skytypes.SkycoinAPI)
	if !ok {
		logNetwork.Errorf("There is no proper client in %s pool", section)
		return nil, errors.ErrInvalidPoolObjectType
	}
	return &SkycoinApiClient{
		SkycoinAPI: skyApi,
		pool:       pool,
	}, nil
}

func ReturnSkycoinClient(obj skytypes.SkycoinAPI) {
	poolObj, ok := obj.(*SkycoinApiClient)
	if !ok {
		return
	}
	poolObj.pool.Put(poolObj.SkycoinAPI)
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
	unTxn, ok := txn.(skytypes.SkycoinTxn)
	if !ok {
		return errors.ErrInvalidTxn
	}
	c, err := NewSkycoinApiClient(spex.poolSection)
	if err != nil {
		return err
	}
	defer ReturnSkycoinClient(c)
	txnBytes, err := unTxn.EncodeSkycoinTransaction()
	if err != nil {
		return err
	}
	_, err = c.InjectEncodedTransaction(hex.EncodeToString(txnBytes))
	if err != nil {
		return err
	}

	return nil
}

func (spex *SkycoinPEX) GetTxnPool() (core.TransactionIterator, error) {
	logNetwork.Info("Getting transaction pool")
	c, err := NewSkycoinApiClient(PoolSection)
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

type SkycoinPexNodeIterator struct {
	//Implements PexNodeIterator interface
	current  int
	networks []core.PexNode
}

func (it *SkycoinPexNodeIterator) Value() core.PexNode {
	return it.networks[it.current]
}

func (it *SkycoinPexNodeIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *SkycoinPexNodeIterator) HasNext() bool {
	return !((it.current + 1) >= len(it.networks))
}

func NewSkycoinPexNodeIterator(network []core.PexNode) *SkycoinPexNodeIterator {
	return &SkycoinPexNodeIterator{networks: network, current: -1}
}

type SkycoinNetworkConnections struct {
	//Implements NetworkSet interface
	nodeAddress string
}

func NewSkycoinRemoteNetwork(nodeAddress string) *SkycoinNetworkConnections {
	return &SkycoinNetworkConnections{nodeAddress}
}

func (remoteNetwork *SkycoinNetworkConnections) newClient() *api.Client {
	return api.NewClient(remoteNetwork.nodeAddress)
}

func (remoteNetwork *SkycoinNetworkConnections) ListPeers() core.PexNodeIterator {
	logNetwork.Info("Getting list of peers in Skycoin network connections")
	c := remoteNetwork.newClient()
	nets, err := c.NetworkConnections(nil)

	if err != nil {
		logNetwork.WithError(err).Warn("Couldn't get connections")
		return nil
	}
	var netIterator []core.PexNode
	for _, con := range nets.Connections {
		netIterator = append(netIterator, connectionsToNetwork(con))
	}

	return NewSkycoinPexNodeIterator(netIterator)
}

type SkycoinPexNode struct {
	Ip          string
	Port        uint16
	Source      bool
	Block       uint64
	LastSeenIn  int64
	LastSeenOut int64
}

func (network *SkycoinPexNode) GetIp() string {
	return network.Ip
}

func (network *SkycoinPexNode) GetPort() uint16 {
	return network.Port
}

func (network *SkycoinPexNode) GetBlockHeight() uint64 {
	return network.Block
}

func (network *SkycoinPexNode) IsTrusted() bool {
	return network.Source
}

func (network *SkycoinPexNode) GetLastSeenIn() int64 {
	return network.LastSeenIn
}

func (network *SkycoinPexNode) GetLastSeenOut() int64 {
	return network.LastSeenOut
}

func connectionsToNetwork(connection readable.Connection) *SkycoinPexNode {
	return &SkycoinPexNode{
		Ip:          strings.Split(connection.Addr, ":")[0],
		Port:        connection.ListenPort,
		LastSeenIn:  connection.LastSent,
		LastSeenOut: connection.LastReceived,
		Block:       connection.Height,
		Source:      connection.IsTrustedPeer,
	}
}
