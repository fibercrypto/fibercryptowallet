package explorer

import (
// skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
// "github.com/fibercrypto/fibercryptowallet/src/core"
// "github.com/fibercrypto/fibercryptowallet/src/params"
// "github.com/fibercrypto/fibercryptowallet/src/util/logging"
// "github.com/fibercrypto/fibercryptowallet/src/core"
// qtCore "github.com/therecipe/qt/core"
)

// var logExplorer = logging.MustGetLogger("Explorer")

func init() {
	BlocksModel_QmlRegisterType2("ExplorerModels", 1, 0, "QBlocks")
}

//
// type ExploreManager struct {
// 	qtCore.QObject
// 	// Blockchain core.BlockchainStatus
// 	BlockList  BlocksModel
// 	_          int    `property:"currentPage"`
// 	_          func() `constructor:"init"`
// 	_          func() `slot:"loadModel"`
// 	// _        func() []*QNetworking `slot:"getNetworks"`
// }
//
// func (explorer *ExploreManager) init() {
// 	// logExplorer.Info("Init explorerManager")
// 	// explorer.Blockchain = skycoin.NewSkycoinBlockchain(params.DataRefreshTimeout)
// 	// explorer.BlockList.init()
// }
//
// func (explorer *ExploreManager) loadModel() {
// 	// explorer.Blockchain = skycoin.NewSkycoinBlockchain(params.DataRefreshTimeout)
// }
