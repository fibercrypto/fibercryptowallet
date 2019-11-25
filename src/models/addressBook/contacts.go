package addressBook

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
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
	ID
)

const path = ".fiber/data.dt"

var isOpen = false
var db *data.DB
var logAddressBook = logging.MustGetLogger("AddressBook")

func init() { AddrsBookModel_QmlRegisterType2("AddrsBookManager", 1, 0, "AddrsBookModel") }

var addresses = make([]core.ReadableAddress, 0)

type AddrsBookModel struct {
	qtcore.QAbstractListModel

	_ map[int]*qtcore.QByteArray            `property:"roles"`
	_ []*QContact                           `property:"contacts"`
	_ int                                   `property:"count"`
	_ func()                                `constructor:"init"`
	_ func(row int, id uint64)              `slot:"removeContact,auto"`
	_ func(row int, id uint64, name string) `slot:"editContact,auto"`
	_ func(name string)                     `slot:"newContact"`
	_ func(string) bool                     `slot:"openAddrsBook"`
	_ func(string) bool                     `slot:"initAddrsBook"`
	_ func() bool                           `slot:"exist"`
	_ func(value, coinType string)          `slot:"addAddress"`
}

type QContact struct {
	qtcore.QObject
	_ uint64              `property:"id"`
	_ string              `property:"name"`
	_ AddrsBkAddressModel `property:"address"`
}

func (abm *AddrsBookModel) init() {
	logAddressBook.Info("Init addressBook model")
	abm.SetRoles(map[int]*qtcore.QByteArray{
		Name:    qtcore.NewQByteArray2("name", -1),
		Address: qtcore.NewQByteArray2("address", -1),
		ID:      qtcore.NewQByteArray2("id", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(abm, qml.QQmlEngine__CppOwnership)
	abm.ConnectRowCount(abm.rowCount)
	abm.ConnectData(abm.data)
	abm.ConnectColumnCount(abm.columnCount)
	abm.ConnectRoleNames(abm.roleNames)
	abm.ConnectNewContact(abm.newContact)
	abm.ConnectDestroyAddrsBookModel(abm.close)
	abm.ConnectOpenAddrsBook(abm.openAddrsBook)
	abm.ConnectInitAddrsBook(abm.initAddrsBook)
	// abm.ConnectEditContact(abm.editContact)
	// abm.ConnectRemoveContact(abm.removeContact)
	abm.ConnectExist(abm.exist)
	abm.ConnectAddAddress(abm.addAddress)
}

func (abm *AddrsBookModel) rowCount(*qtcore.QModelIndex) int {
	return len(abm.Contacts())
}

func (abm *AddrsBookModel) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	logAddressBook.Info("Loading data for index")
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}
	if index.Row() >= len(abm.Contacts()) {
		return qtcore.NewQVariant()
	}
	contact := abm.Contacts()[index.Row()]

	switch role {
	case Name:
		{
			return qtcore.NewQVariant1(contact.Name())
		}
	case Address:
		{
			return qtcore.NewQVariant1(contact.Address())
		}
	case ID:
		{
			return qtcore.NewQVariant1(contact.Id())
		}
	default:
		return qtcore.NewQVariant()
	}
}

func (abm *AddrsBookModel) roleNames() map[int]*qtcore.QByteArray {
	return abm.Roles()
}

func (abm *AddrsBookModel) columnCount(parent *qtcore.QModelIndex) int {
	return 1
}

func (abm *AddrsBookModel) removeContact(row int, id uint64) {
	logAddressBook.Infof("Remove contact with id %d", id)
	if row < 0 || row >= abm.Count() {
		return
	}
	if err := db.DeleteContact(id); err != nil {
		logAddressBook.Error(err)
		return
	}
	abm.BeginRemoveRows(qtcore.NewQModelIndex(), row, row)
	abm.SetContacts(append(abm.Contacts()[:row], abm.Contacts()[row+1:]...))
	abm.EndRemoveRows()
	abm.SetCount(abm.Count() - 1)

}

func (abm *AddrsBookModel) addContact(c *QContact) {
	logAddressBook.Info("Add Contact")
	var row = 0
	for row < len(abm.Contacts()) && c.Name() > abm.Contacts()[row].Name() {
		row++
	}
	abm.BeginInsertColumns(qtcore.NewQModelIndex(), row, row)
	qml.QQmlEngine_SetObjectOwnership(c, qml.QQmlEngine__CppOwnership)
	abm.SetContacts(append(append(abm.Contacts()[:row], c), abm.Contacts()[row:]...))
	abm.EndInsertRows()
	abm.SetCount(abm.Count() + 1)
}

func (abm *AddrsBookModel) editContact(row int, id uint64, name string) {
	logAddressBook.Info("Edit contact")
	if row < 0 || row >= abm.Count() {
		return
	}
	qc := NewQContact(nil)
	qc.SetName(name)
	qa := fromAddressToQAddress(addresses)
	am := NewAddrsBkAddressModel(nil)
	am.SetAddress(qa)
	qc.SetAddress(am)
	qc.SetId(id)
	var c = data.Contact{}
	c.SetAddresses(addresses)
	c.SetName(name)
	if err := db.UpdateContact(id, &c); err != nil {
		logAddressBook.Error(err)
		return
	}
	abm.BeginRemoveRows(qtcore.NewQModelIndex(), row, row)
	abm.SetContacts(append(abm.Contacts()[:row], abm.Contacts()[row+1:]...))
	abm.EndRemoveRows()
	abm.SetCount(abm.Count() - 1)
	abm.addContact(qc)
	addresses = []core.ReadableAddress{}
}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, path)
	return fileDir
}

func (abm *AddrsBookModel) loadContacts(contacts []*QContact) {
	logAddressBook.Info("loading contacts")
	for _, c := range contacts {
		abm.addContact(c)
	}
}

func (abm *AddrsBookModel) newContact(name string) {
	qc := NewQContact(nil)
	qc.SetName(name)
	qa := fromAddressToQAddress(addresses)
	am := NewAddrsBkAddressModel(nil)
	am.SetAddress(qa)
	qc.SetAddress(am)
	var contact data.Contact
	contact.SetName(name)
	contact.SetAddresses(addresses)
	if id, err := db.InsertContact(&contact); err != nil {
		logAddressBook.Error(err)
	} else {
		qc.SetId(id)
	}
	addresses = []core.ReadableAddress{}
	abm.addContact(qc)
}

func (*AddrsBookModel) close() {
	logAddressBook.Info("Closing address book")
	if isOpen {
		if err := db.Close(); err != nil {
			logAddressBook.Error(err)
		} else {
			isOpen = false
		}
	}
}

func (abm *AddrsBookModel) openAddrsBook(password string) bool {
	var err error
	if !abm.exist() {
		return false
	}

	if isOpen {
		abm.close()
	}
	abm.SetContacts([]*QContact{})
	logAddressBook.Info("Opening address book")
	if db, err = data.LoadFromFile(getConfigFileDir(), []byte(password)); err != nil {
		logAddressBook.Error(err)
		return false
	}

	contacts, err := db.ListContact()
	if err != nil {
		logAddressBook.Error(err)
	}
	qcontacts := fromContactToQContact(contacts)
	logAddressBook.Infof("%#v", qcontacts)
	abm.loadContacts(qcontacts)
	isOpen = true
	return true
}

func (abm *AddrsBookModel) initAddrsBook(password string) bool {
	var err error
	if abm.exist() {
		return false
	}
	logAddressBook.Info("Creating address book")

	if db, err = data.Init([]byte(password), getConfigFileDir()); err != nil {
		logAddressBook.Error(err)
	}

	isOpen = true
	return true
}

func (*AddrsBookModel) exist() bool {
	_, err := os.Stat(getConfigFileDir())
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		logAddressBook.Error(err)
		return false
	}
	return true
}

func fromContactToQContact(contacts []core.Contact) []*QContact {
	var qContacts = make([]*QContact, 0)
	for _, c := range contacts {
		qc := NewQContact(nil)
		qc.SetName(c.GetName())
		logAddressBook.Info(c.GetID())
		qc.SetId(c.GetID())
		qAddressModel := NewAddrsBkAddressModel(nil)
		qAddressModel.SetAddress(fromAddressToQAddress(c.GetAddresses()))
		qc.SetAddress(qAddressModel)
		qContacts = append(qContacts, qc)
	}
	return qContacts
}

func (*AddrsBookModel) addAddress(value, coinType string) {
	logAddressBook.Infof("%#v", addresses)
	logAddressBook.Infof("value: %#v, type: %#v", value, coinType)
	addresses = append(addresses, &data.Address{Value: []byte(value), Coin: []byte(coinType)})
}
