package history

import (
	"github.com/therecipe/qt/core"
)

const (
	transactionStatusConfirmed = iota
	transactionStatusPending
	transactionStatusPreview
)

const (
	transactionTypeSend = iota
	transactionTypeReceive
)

type Address struct {
	address          string
	addressSky       int64
	addressCoinHours int64
}

type TransactionDetails struct {
	date            string
	status          int8
	ttype           int8
	amount          int64
	hoursReceived   int64
	hoursBurned     int64
	transactionID   string
	sentAddress     string
	receivedAddress string
}

type QHistoryModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	transactions []TransactionDetails
}

func (hm *QHistoryModel) init() {

	hm.transactions = []TransactionDetails{
		TransactionDetails{"2019-03-02 10:19", transactionStatusConfirmed, transactionTypeSend, 9, 65, 6287, "ksuwet837iwetr27ietr", "734irwepaweygtawieta783cwyc", "uewyru63u3789jfrgpaldcxz4f6"},
		TransactionDetails{"2019-03-02 10:44", transactionStatusConfirmed, transactionTypeReceive, 10000, 9421, 218489, "hdk84iesurys29q8ikhf", "nvslkjoid98wemxsoqzmap50vah", "ps73729ssksof972jwkf83owhfi"},
		TransactionDetails{"2019-03-02 15:01", transactionStatusConfirmed, transactionTypeSend, 125, 4203, 7950, "p9438938qksaxjhsdkuq", "qwe98uimhfdkfu23y9hewi8qd02", "p2837djkdjvkauyfiawbw83h48f"},
		TransactionDetails{"2019-03-03 08:13", transactionStatusConfirmed, transactionTypeSend, 100, 3877, 6911, "73t4dunksxr8w7zuwe13", "89niweuq82zqur37izuwsklparu", "0j3726djf9f4w9sovgwuw9e4n9s"},
		TransactionDetails{"2019-03-03 12:23", transactionStatusConfirmed, transactionTypeReceive, 80, 3084, 4603, "9782nz87e7tquetyn83w", "t7ekwduf8045ogmcbq63nf9bm36", "aske8jdhbsmwq9204h49y25jrue"},
		TransactionDetails{"2019-03-03 13:33", transactionStatusConfirmed, transactionTypeReceive, 60, 2796, 2347, "9127enqiuyrt6tyejshd", "2imdiym9w7ji8s29ydk0djwd6wj", "0gndqy725k9fj3ewdardgbkd83l"},
	}

}
