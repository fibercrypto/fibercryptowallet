package explorer

import (
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtcore "github.com/therecipe/qt/core"
)

var logExplorer = logging.MustGetLogger("Explorer")

type BlocksModel struct {
	qtcore.QAbstractListModel
	_ int               `property:"currentPage"`
	_ func()            `constructor:"init"`
	_ func(pageNum int) `signal:"loadPage"`
	_ func()            `signal:"update,auto"`
}

func (blocksM *BlocksModel) init() {
}

func (blocksM *BlocksModel) update() {
	// update info
	if err := blocksM.updateInfo(); err != nil {
		logExplorer.WithError(err).Warn("Couldn't update blockchain Info")
		return
	}
	return
}

// updateInfo request the needed information
func (blocksM *BlocksModel) updateInfo() error {
	return nil
}

func (blocksM *BlocksModel) loadPage(pageNum int) {

}
