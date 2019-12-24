package skycoin

import (
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"strings"

	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

type SkycoinPexNodeIterator struct {
	core.Iterator
}

func NewSkycoinPexNodeIterator(network []core.PexNode) *SkycoinPexNodeIterator {
	return &SkycoinPexNodeIterator{Iterator: util.NewGenericIterator(network)}
}

type SkycoinNetworkConnections struct {
	// Implements NetworkSet interface
	nodeAddress string
}

func NewSkycoinRemoteNetwork(nodeAddress string) *SkycoinNetworkConnections {
	return &SkycoinNetworkConnections{nodeAddress}
}

func (remoteNetwork *SkycoinNetworkConnections) newClient() *api.Client {
	return api.NewClient(remoteNetwork.nodeAddress)
}

func (remoteNetwork *SkycoinNetworkConnections) ListPeers() core.Iterator {
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

var _ core.Iterator = &SkycoinPexNodeIterator{}
