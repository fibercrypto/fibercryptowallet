package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
	"log"
)

type SkycoinNetworkIterator struct {
	//Implements NetworkIterator interface
	current  int
	networks []core.INetwork
}

func (it *SkycoinNetworkIterator) Value() core.INetwork {
	return it.networks[it.current]
}

func (it *SkycoinNetworkIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *SkycoinNetworkIterator) HasNext() bool {
	if (it.current + 1) >= len(it.networks) {
		return false
	}
	return true
}

func NewSkycoinNetworkIterator(network []core.INetwork) *SkycoinNetworkIterator {
	return &SkycoinNetworkIterator{networks: network, current: -1}
}

type SkycoinRemoteNetwork struct {
	//Implements WalletStorage and WalletSet interfaces
	nodeAddress string
}

func (remoteNetwork *SkycoinRemoteNetwork) newClient() *api.Client {
	return api.NewClient(remoteNetwork.nodeAddress)
}

func (remoteNetwork *SkycoinRemoteNetwork) ListNetworks() core.NetworkIterator {
	c := remoteNetwork.newClient()
	nets, err := c.NetworkConnections(nil)
	if err != nil {
		log.Print("Error getting connections")
		return nil
	}
	var netIterator []core.INetwork
	for _, con := range nets.Connections {
		netIterator = append(netIterator, connectionsToNetwork(con))
	}
	return NewSkycoinNetworkIterator(netIterator)
}

type Network struct {
	Ip          string
	Port        uint16
	Source      bool
	Block       uint64
	LastSeenIn  int64
	LastSeenOut int64
}

func (network Network)GetIp() string{
	return network.Ip
}

func (network Network)GetPort() uint16 {
	return network.Port
}

func (network Network)GetBlock() uint64 {
	return network.Block
}

func (network Network)IsTrusted() bool{
	return network.Source
}

func (network Network)GetLastSeenIn() int64{
	return network.LastSeenIn
}

func (network Network)GetLastSeenOut() int64{
	return network.LastSeenOut
}



func connectionsToNetwork(connection readable.Connection) Network {
	return Network{
		Ip:   connection.Addr,
		Port: connection.ListenPort,
		LastSeenIn: connection.LastSent,
		LastSeenOut: connection.LastReceived,
		Block: connection.Height,
		Source: connection.IsTrustedPeer,
	}
}
