package addressBook

import (
	"fmt"
	cor "github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/data"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"github.com/skycoin/skycoin/src/util/file"
	"github.com/therecipe/qt/core"
	"os"
	"path/filepath"
)

const path = ".fiber/data.dt"

var logAddressBook = logging.MustGetLogger("AddressBook")
var db cor.AddressBook

func init() { AddrsBookManager_QmlRegisterType2("AddrsBookModels", 1, 0, "AddrsBookManager") }

type AddrsBookManager struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ func()                                  `signal:"remove,auto"`
	_ func(obj []*core.QVariant)              `signal:"add,auto"`
	_ func(firstName string, lastName string) `signal:"edit,auto"`

	modelData []cor.Contact
}

func (m *AddrsBookManager) init() {
	m.modelData = []cor.Contact{}
	var err error
	if ok, _ := file.Exists(getConfigFileDir()); !ok {
		db, err = data.Init([]byte(""), getConfigFileDir(), "")
		if err != nil {
			logAddressBook.Error(err)
		}
	} else {
		db, err = data.LoadFromFile(getConfigFileDir(), []byte(""))
		if err != nil {
			logAddressBook.Error(err)
		}
	}

	cc, err := db.ListContact()
	if err != nil {
		logAddressBook.Error(err)
	}
	m.modelData = append(m.modelData, cc...)
	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
}

func (m *AddrsBookManager) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *AddrsBookManager) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}

	item := m.modelData[index.Row()]
	return core.NewQVariant1(fmt.Sprintf("%v %v", string(item.(*data.Contact).Name), string(item.GetAddresses()[0].GetValue())))
}

func (m *AddrsBookManager) remove() {

}

func (m *AddrsBookManager) add(item []*core.QVariant) {

}

func (m *AddrsBookManager) edit(firstName string, lastName string) {
}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, path)
	return fileDir
}
