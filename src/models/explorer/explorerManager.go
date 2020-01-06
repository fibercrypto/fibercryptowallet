package explorer

// import (
// 	qtCore "github.com/therecipe/qt/core"
// )
//
func init() {
	BlocksModel_QmlRegisterType2("ExplorerModels", 1, 0, "QBlocks")
	// BlocksModel_QmlRegisterType2("ExplorerModels", 1, 0, "ExplorerManager")
}

//
// type ExploreManager struct {
// 	qtCore.QObject
// 	Blockchain core.BlockchainStatus
// _ func() `constructor:"init"`
//
// _ func(hash string) `slot:"loadBlockByHash"`
// }
//
// func (explorer *ExploreManager) init() {
// 	logExplorer.Info("Init explorerManager")
//
// }
//
