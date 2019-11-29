package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"github.com/therecipe/qt/core"
	qtcore "github.com/therecipe/qt/core"
)

// Properties:

// For blocks:
// - Time
// - Block number
// - Transactions (number of Transactions)
// - Blockhash

// For block's details:
// - Height
// - Timestamp
// - Size
// - Hash
// - Parent Hash
// - Total Amount

// - Next block
// - Previous block

// For Transactions:
// - Inputs:
// 	* Coins
// 	* Initial hours
// 	* Final hours

// - Outputs:
// 	* Coins
// 	* Hours

// - Transaction fee (in hours)

// - Links for the transactions

var logExplorer = logging.MustGetLogger("modelsExplorer")

const (
	QBlocks = int(qtcore.Qt__UserRole) + iota + 1
)

// ModelBlocks List of Blocks to be show.
type ModelBlocks struct {
	qtcore.QObject
	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QBlock                `property:"outputs"`

	_ func([]*QBlock) `slot:"addBlocks"`
	_ func([]*QBlock) `slot:"insertBlocks"`
	_ func([]*QBlock) `slot:"loadModel"`
	_ func()          `slot:"cleanModel"`
}

// QBlock Contains info about the block to be show.
type QBlock struct {
	core.QObject

	_ string `property:"time"`
	_ string `property:"blockNumber"`
	_ string `property:"Transactions"`
	_ string `property:"Blockhash"`
}

func (m *ModelBlocks) init() {
	m.SetRoles(map[int]*core.QByteArray{
		QBlocks: core.NewQByteArray2("qblocks", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectAddBlocks(m.addBlocks)
	m.ConnectInsertBlocks(m.insertBlocks)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectCleanModel(m.cleanModel)

}
