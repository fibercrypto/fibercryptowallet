package explorer

func init() {
	BlocksModel_QmlRegisterType2("ExplorerModels", 1, 0, "QBlocks")
	// BlocksModel_QmlRegisterType2("ExplorerModels", 1, 0, "ExplorerManager")
}

// type ExplorerManager struct {
// 	qtCore.QObject
// 	_ func() `constructor:"init"`

// _         func() []*transactions.TransactionDetails `slot:"loadHistoryWithFilters"`
// _         func() []*transactions.TransactionDetails `slot:"loadHistory"`
// _         func(string)                              `slot:"addFilter"`
// _         func(string)                              `slot:"removeFilter"`
// walletEnv core.WalletEnv
// }
//
// func (em *ExplorerManager) init() {
//
// }
//
// func (em *ExplorerManager) GetTransactionDetails(txId string) *transactions.TransactionDetails {
//
// }
