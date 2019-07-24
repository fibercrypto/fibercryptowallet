package history

func init() {
	TransactionList_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionList")
	TransactionDetails_QmlRegisterType2("HistoryModels", 1, 0, "QTransactionDetail")
	AddressDetails_QmlRegisterType2("HistoryModels", 1, 0, "QAddress")
	AddressList_QmlRegisterType2("HistoryModels", 1, 0, "QAddressList")
}
