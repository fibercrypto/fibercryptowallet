package addressBook

import (
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/data"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
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

var addrsBook core.AddressBook
var logAddressBook = logging.MustGetLogger("Address Book Model")

func init() { AddrsBookModel_QmlRegisterType2("AddrsBookManager", 1, 0, "AddrsBookModel") }

var addresses = make([]core.StringAddress, 0)

type AddrsBookModel struct {
	qtcore.QAbstractListModel

	_ map[int]*qtcore.QByteArray                          `property:"roles"`
	_ []*QContact                                         `property:"contacts"`
	_ int                                                 `property:"count"`
	_ func()                                              `constructor:"init"`
	_ func(row int, id uint64)                            `slot:"removeContact,auto"`
	_ func(row int, id uint64, name string)               `slot:"editContact,auto"`
	_ func(name string)                                   `slot:"newContact"`
	_ func()                                              `slot:"loadContacts"`
	_ func(int, string)                                   `slot:"initAddrsBook"`
	_ func() int                                          `slot:"getSecType"`
	_ func(string) bool                                   `slot:"authenticate"`
	_ func() bool                                         `slot:"hasInit"`
	_ func(value, coinType string)                        `slot:"addAddress"`
	_ func(value string) bool                             `slot:"addressIsValid"`
	_ func(row int, name string) bool                     `slot:"nameExist"`
	_ func(row int, address string, coinType string) bool `slot:"addressExist"`
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
	abm.ConnectGetSecType(abm.getSecType)
	abm.ConnectAuthenticate(abm.authenticate)
	abm.ConnectAddressIsValid(abm.addressIsValid)
	abm.ConnectAddressExist(abm.addressExist)
	abm.ConnectNameExist(abm.nameExist)
	// abm.ConnectDestroyAddrsBookModel(abm.close)
	abm.ConnectLoadContacts(abm.loadContacts)
	abm.ConnectInitAddrsBook(abm.initAddrsBook)
	// abm.ConnectEditContact(abm.editContact)
	// abm.ConnectRemoveContact(abm.removeContact)
	abm.ConnectHasInit(abm.hasInit)
	abm.ConnectAddAddress(abm.addAddress)
	if addrsBook == nil {
		db, err := data.GetBoltStorage(getConfigFileDir())
		if err != nil {
			logAddressBook.Error(err)
		}
		addrsBook = data.NewAddressBook(db)
	}

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
	if err := addrsBook.DeleteContact(id); err != nil {
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
	if err := addrsBook.UpdateContact(id, &c); err != nil {
		logAddressBook.Error(err)
		return
	}
	abm.BeginRemoveRows(qtcore.NewQModelIndex(), row, row)
	abm.SetContacts(append(abm.Contacts()[:row], abm.Contacts()[row+1:]...))
	abm.EndRemoveRows()
	abm.SetCount(abm.Count() - 1)
	abm.addContact(qc)
	addresses = []core.StringAddress{}
}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, ".fiber/data.dt")
	return fileDir
}

func (abm *AddrsBookModel) loadContacts() {
	logAddressBook.Info("loading contacts")
	abm.SetContacts([]*QContact{})
	contacts, err := addrsBook.ListContact()
	if err != nil {
		logAddressBook.Error(err)
	}
	qContacts := fromContactToQContact(contacts)

	for _, c := range qContacts {
		abm.addContact(c)
	}
}

func (abm *AddrsBookModel) getSecType() int {
	secType, err := addrsBook.GetSecType()
	if err != nil {
		logAddressBook.Error(err)
	}
	return secType
}

func (abm *AddrsBookModel) authenticate(password string) bool {
	if err := addrsBook.Authenticate(password); err != nil {
		logAddressBook.Error(err)
		return false
	}
	return true
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
	if id, err := addrsBook.InsertContact(&contact); err != nil {
		logAddressBook.Error(err)
	} else {
		qc.SetId(id)
	}
	addresses = []core.StringAddress{}
	abm.addContact(qc)
}

func (*AddrsBookModel) close() {
	logAddressBook.Info("Closing address book")
	if addrsBook.IsOpen() {
		if err := addrsBook.Close(); err != nil {
			logAddressBook.Error(err)
		}
	}
}

func (abm *AddrsBookModel) initAddrsBook(secType int, password string) {
	var err error
	if addrsBook.HasInit() {
		return
	}
	logAddressBook.Info("Creating address book")
	if err = addrsBook.Init(secType, password); err != nil {
		logAddressBook.Error(err)
	}
}

func (*AddrsBookModel) hasInit() bool {
	return addrsBook.HasInit()
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
	for e := range addresses {
		if string(addresses[e].GetValue()) == value && string(addresses[e].GetCoinType()) == coinType {
			return
		}
	}
	addresses = append(addresses, &data.Address{Value: []byte(value), Coin: []byte(coinType)})
}

func (*AddrsBookModel) addressIsValid(value string) bool {
	if _, err := skycoin.NewSkycoinAddress(value); err != nil {
		return false
	}
	return true
}

func (abm *AddrsBookModel) nameExist(row int, name string) bool {
	for e := range abm.Contacts() {
		if row != e && abm.Contacts()[e].Name() == name {
			return true
		}
	}
	return false
}

func (abm *AddrsBookModel) addressExist(row int, address string, coinType string) bool {
	for e := range abm.Contacts() {
		for f := range abm.Contacts()[e].Address().Address() {
			if row != e && abm.Contacts()[e].Address().Address()[f].Value() == address &&
				abm.Contacts()[e].Address().Address()[f].CoinType() == coinType {
				return true
			}
		}
	}
	return false
}
