package models

import (
	"github.com/SkycoinProject/skycoin/src/util/logging"
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	qtCore "github.com/therecipe/qt/core"
)

var logNetworkingManager = logging.MustGetLogger("modelsNetworkingManager")

type NetworkingManager struct {
	qtCore.QObject
	Networks core.PexNodeSet
	_        func()                `constructor:"init"`
	_        func() []*QNetworking `slot:"getNetworks"`
}

func (net *NetworkingManager) init() {
	net.ConnectGetNetworks(net.getNetworks)
	net.Networks = skycoin.NewSkycoinRemoteNetwork("http://127.0.0.1:6420")

}

func (net *NetworkingManager) getNetworks() []*QNetworking {
	logNetworkingManager.Info("Getting networks")
	networks := make([]*QNetworking, 0)

	netIterator := net.Networks.ListPeers()

	if netIterator == nil {
		logNetworkingManager.WithError(nil).Error("Couldn't load networks")
		return networks
	}
	for netIterator.Next() {

		networks = append(networks, INetworkToQNetworking(netIterator.Value()))
	}

	return networks
}
