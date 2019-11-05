package addressBook

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/data"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	qtcore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"os"
	"path/filepath"
)

const (
	Name = int(qtcore.Qt__UserRole + (iota + 1))
	Address
)

const path = ".fiber/data.dt"

var logAddressBook = logging.MustGetLogger("AddressBook")

func init() { AddrsBookModel_QmlRegisterType2("AddrsBookManager", 1, 0, "AddrsBookModel") }

type AddrsBookModel struct {
	qtcore.QAbstractListModel

	_ map[int]*qtcore.QByteArray               `property:"roles"`
	_ []*QContact                              `property:"contacts"`
	_ func()                                   `constructor:"init"`
	_ func(row int)                            `slot:"removeContact,auto"`
	_ func(*QContact)                          `slot:"addContact,auto"`
	_ func(row int, name string, addrs string) `slot:"editContact,auto"`
	_ func([]*QContact)                        `slot:"loadContacts,auto"`
	_ func(name, address string)               `slot:"newContact"`
	_ int                                      `property:"count"`
}

type QContact struct {
	qtcore.QObject
	_ string `property:"name"`
	_ string `property:"address"`
}

// type QAddress struct {
// 	_ string `property:"value"`
// 	_ string `property:"coinType"`
// }

func (adm *AddrsBookModel) init() {
	logAddressBook.Info("Init addressBook model")
	if _, err := data.Init([]byte(""), getConfigFileDir(), ""); err != nil {
		logAddressBook.Error(err)
	}
	adm.SetRoles(map[int]*qtcore.QByteArray{
		Name:    qtcore.NewQByteArray2("name", -1),
		Address: qtcore.NewQByteArray2("address", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(adm, qml.QQmlEngine__CppOwnership)
	adm.ConnectRowCount(adm.rowCount)
	adm.ConnectData(adm.data)
	adm.ConnectColumnCount(adm.columnCount)
	adm.ConnectRoleNames(adm.roleNames)

	adm.ConnectEditContact(adm.editContact)
	adm.ConnectRemoveContact(adm.removeContact)
	adm.ConnectAddContact(adm.addContact)
	adm.ConnectLoadContacts(adm.loadContacts)
	adm.ConnectNewContact(adm.newContact)
}

func (adm *AddrsBookModel) rowCount(*qtcore.QModelIndex) int {
	return len(adm.Contacts())
}

func (adm *AddrsBookModel) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	logAddressBook.Info("Loading data for index")
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}
	if index.Row() >= len(adm.Contacts()) {
		return qtcore.NewQVariant()
	}
	contact := adm.Contacts()[index.Row()]
	logAddressBook.Info("Outside of swith")

	switch role {
	case Name:
		{
			logAddressBook.Info("Inside of Name case.")
			return qtcore.NewQVariant1(contact.Name())
		}
	case Address:
		{
			return qtcore.NewQVariant1(contact.Address())
		}
	default:
		return qtcore.NewQVariant()
	}
}

func (adm *AddrsBookModel) roleNames() map[int]*qtcore.QByteArray {
	return adm.Roles()
}

func (adm *AddrsBookModel) columnCount(parent *qtcore.QModelIndex) int {
	return 1
}

func (adm *AddrsBookModel) removeContact(row int) {
	logAddressBook.Info("Remove contact for index")
	adm.BeginRemoveRows(qtcore.NewQModelIndex(), row, row)
	adm.SetContacts(append(adm.Contacts()[:row], adm.Contacts()[row+1:]...))
	adm.EndRemoveRows()
	adm.SetCount(adm.Count() - 1)

}

func (adm *AddrsBookModel) addContact(c *QContact) {
	logAddressBook.Info("Add Contact")
	adm.BeginInsertColumns(qtcore.NewQModelIndex(), len(adm.Contacts()), len(adm.Contacts()))
	// db, err := data.LoadFromFile(getConfigFileDir(), []byte(""))
	// if err != nil {
	// 	logAddressBook.Error(err)
	// }
	// defer db.Close()
	qml.QQmlEngine_SetObjectOwnership(c, qml.QQmlEngine__CppOwnership)
	adm.SetContacts(append(adm.Contacts(), c))

	// if err := db.InsertContact(&data.Contact{}); err != nil {
	// 	logAddressBook.Error(err)
	// }
	// adm.ConnectData(adm.data)
	adm.EndInsertRows()
	adm.SetCount(adm.Count() + 1)
}

func (adm *AddrsBookModel) editContact(row int, name string, addrs string) {}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, path)
	return fileDir
}

func (adm *AddrsBookModel) loadContacts(contacts []*QContact) {
	logAddressBook.Info("loading contacts")
	for _, c := range contacts {
		qml.QQmlEngine_SetObjectOwnership(c, qml.QQmlEngine__CppOwnership)
	}
	adm.BeginResetModel()
	adm.SetContacts(contacts)
	adm.EndResetModel()
	adm.SetCount(len(adm.Contacts()))
}

func (adm *AddrsBookModel) newContact(name, address string) {
	qc := NewQContact(nil)
	qc.SetName(name)
	qc.SetAddress(address)
	logAddressBook.Infof("inside of new Contact: name: %s \t address: %s", name, address)
	adm.addContact(qc)
}
