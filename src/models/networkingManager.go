package models

import (
	skycoin "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	qtcore "github.com/therecipe/qt/core"
)

type NetworkingManager struct {
	qtcore.QObject
	Networks *core.NetworkSet
	_        func()                `constructor:"init"`
	_        func() []*QNetworking `slot:"getNetworks"`
}

func (net *NetworkingManager) init() {
	net.ConnectGetNetworks(net.getNetworks)
	net.Networks = skycoin.NewSkycoinRemoteNetwork("http://127.0.0.1:6420")

}

func (net *NetworkingManager) getNetworks() []*QNetworking {
	networks := make([]*QNetworking, 0)
	netIterator := net.Networks.ListNetworks()
	for netIterator.Next() {
		networks = append(networks, INetworkToQNetworking(netIterator.Value()))
	}
	return networks
}
