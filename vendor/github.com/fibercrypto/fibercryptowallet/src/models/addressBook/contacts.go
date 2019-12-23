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

	_ map[int]*qtcore.QByteArray                                 `property:"roles"`
	_ []*QContact                                                `property:"contacts"`
	_ int                                                        `property:"count"`
	_ func()                                                     `constructor:"init"`
	_ func(row int, id uint64)                                   `slot:"removeContact,auto"`
	_ func(row int, id uint64, name string)                      `slot:"editContact,auto"`
	_ func(name string)                                          `slot:"newContact"`
	_ func()                                                     `slot:"loadContacts"`
	_ func(int, string)                                          `slot:"initAddrsBook"`
	_ func() int                                                 `slot:"getSecType"`
	_ func(password string) bool                                 `slot:"authenticate"`
	_ func() bool                                                `slot:"hasInit"`
	_ func(value, coinType string)                               `slot:"addAddress"`
	_ func(newSecType int, oldPassword, newPassword string) bool `slot:"changeSecType"`
	_ func(value string) bool                                    `slot:"addressIsValid"`
	_ func(row int, name string) bool                            `slot:"nameExist"`
	_ func(row int, address string, coinType string) bool        `slot:"addressExist"`
}

type QContact struct {
	qtcore.QObject
	_ uint64              `property:"id"`
	_ string              `property:"name"`
	_ AddrsBkAddressModel `property:"address"`
}

func (addrBookModel *AddrsBookModel) init() {
	logAddressBook.Info("Init addressBook model")
	addrBookModel.SetRoles(map[int]*qtcore.QByteArray{
		Name:    qtcore.NewQByteArray2("name", -1),
		Address: qtcore.NewQByteArray2("address", -1),
		ID:      qtcore.NewQByteArray2("id", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(addrBookModel, qml.QQmlEngine__CppOwnership)
	addrBookModel.ConnectRowCount(addrBookModel.rowCount)
	addrBookModel.ConnectData(addrBookModel.data)
	addrBookModel.ConnectColumnCount(addrBookModel.columnCount)
	addrBookModel.ConnectRoleNames(addrBookModel.roleNames)
	addrBookModel.ConnectNewContact(addrBookModel.newContact)
	addrBookModel.ConnectGetSecType(addrBookModel.getSecType)
	addrBookModel.ConnectAuthenticate(addrBookModel.authenticate)
	addrBookModel.ConnectAddressIsValid(addrBookModel.addressIsValid)
	addrBookModel.ConnectAddressExist(addrBookModel.addressExist)
	addrBookModel.ConnectNameExist(addrBookModel.nameExist)
	addrBookModel.ConnectLoadContacts(addrBookModel.loadContacts)
	addrBookModel.ConnectInitAddrsBook(addrBookModel.initAddrsBook)
	// addrBookModel.ConnectEditContact(addrBookModel.editContact)
	// addrBookModel.ConnectRemoveContact(addrBookModel.removeContact)
	addrBookModel.ConnectHasInit(addrBookModel.hasInit)
	addrBookModel.ConnectChangeSecType(addrBookModel.changeSecType)
	addrBookModel.ConnectAddAddress(addrBookModel.addAddress)
	if addrsBook == nil {
		db, err := data.GetBoltStorage(getConfigFileDir())
		if err != nil {
			logAddressBook.Error(err)
		}
		addrsBook = data.NewAddressBook(db)
	}

}

func (addrBookModel *AddrsBookModel) rowCount(*qtcore.QModelIndex) int {
	return len(addrBookModel.Contacts())
}

func (addrBookModel *AddrsBookModel) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	logAddressBook.Info("Loading data for index")
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}
	if index.Row() >= len(addrBookModel.Contacts()) {
		return qtcore.NewQVariant()
	}
	contact := addrBookModel.Contacts()[index.Row()]

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

func (addrBookModel *AddrsBookModel) roleNames() map[int]*qtcore.QByteArray {
	return addrBookModel.Roles()
}

func (addrBookModel *AddrsBookModel) columnCount(parent *qtcore.QModelIndex) int {
	return 1
}

func (addrBookModel *AddrsBookModel) removeContact(row int, id uint64) {
	logAddressBook.Infof("Remove contact with id %d", id)
	if row < 0 || row >= addrBookModel.Count() {
		return
	}
	if err := addrsBook.DeleteContact(id); err != nil {
		logAddressBook.Error(err)
		return
	}
	addrBookModel.BeginRemoveRows(qtcore.NewQModelIndex(), row, row)
	addrBookModel.SetContacts(append(addrBookModel.Contacts()[:row], addrBookModel.Contacts()[row+1:]...))
	addrBookModel.EndRemoveRows()
	addrBookModel.SetCount(addrBookModel.Count() - 1)

}

func (addrBookModel *AddrsBookModel) addContact(c *QContact) {
	logAddressBook.Info("Add Contact")
	var row = 0
	for row < len(addrBookModel.Contacts()) && c.Name() > addrBookModel.Contacts()[row].Name() {
		row++
	}
	addrBookModel.BeginInsertColumns(qtcore.NewQModelIndex(), row, row)
	qml.QQmlEngine_SetObjectOwnership(c, qml.QQmlEngine__CppOwnership)
	addrBookModel.SetContacts(append(append(addrBookModel.Contacts()[:row], c), addrBookModel.Contacts()[row:]...))
	addrBookModel.EndInsertRows()
	addrBookModel.SetCount(addrBookModel.Count() + 1)
}

func (addrBookModel *AddrsBookModel) editContact(row int, id uint64, name string) {
	logAddressBook.Info("Edit contact")
	if row < 0 || row >= addrBookModel.Count() {
		return
	}
	qContact := NewQContact(nil)
	qContact.SetName(name)
	qAddresses := fromAddressToQAddress(addresses)
	addrsBkAddressModel := NewAddrsBkAddressModel(nil)
	addrsBkAddressModel.SetAddress(qAddresses)
	qContact.SetAddress(addrsBkAddressModel)
	qContact.SetId(id)
	var contact = data.Contact{}
	contact.SetAddresses(addresses)
	contact.SetName(name)
	if err := addrsBook.UpdateContact(id, &contact); err != nil {
		logAddressBook.Error(err)
		return
	}
	addrBookModel.BeginRemoveRows(qtcore.NewQModelIndex(), row, row)
	addrBookModel.SetContacts(append(addrBookModel.Contacts()[:row], addrBookModel.Contacts()[row+1:]...))
	addrBookModel.EndRemoveRows()
	addrBookModel.SetCount(addrBookModel.Count() - 1)
	addrBookModel.addContact(qContact)
	addresses = []core.StringAddress{}
}

func getConfigFileDir() string {
	homeDir := os.Getenv("HOME")
	fileDir := filepath.Join(homeDir, ".fiber/data.dt")
	return fileDir
}

func (addrBookModel *AddrsBookModel) loadContacts() {
	logAddressBook.Info("loading contactsList")
	addrBookModel.SetContacts([]*QContact{})
	contactsList, err := addrsBook.ListContact()
	if err != nil {
		logAddressBook.Error(err)
	}
	qContactsList := fromContactToQContact(contactsList)

	for _, qContact := range qContactsList {
		addrBookModel.addContact(qContact)
	}
}

func (addrBookModel *AddrsBookModel) getSecType() int {
	return addrsBook.GetSecType()
}

func (addrBookModel *AddrsBookModel) authenticate(password string) bool {

	if err := addrsBook.Authenticate(password); err != nil {
		logAddressBook.Error(err)
		return false
	}
	return true
}

func (addrBookModel *AddrsBookModel) newContact(name string) {
	qContact := NewQContact(nil)
	qContact.SetName(name)
	qAddresses := fromAddressToQAddress(addresses)
	addrsBkAddressModel := NewAddrsBkAddressModel(nil)
	addrsBkAddressModel.SetAddress(qAddresses)
	qContact.SetAddress(addrsBkAddressModel)
	var contact data.Contact
	contact.SetName(name)
	contact.SetAddresses(addresses)
	if id, err := addrsBook.InsertContact(&contact); err != nil {
		logAddressBook.Error(err)
	} else {
		qContact.SetId(id)
	}
	addresses = []core.StringAddress{}
	addrBookModel.addContact(qContact)
}

func (*AddrsBookModel) close() {
	logAddressBook.Info("Closing address book")
	if addrsBook.IsOpen() {
		if err := addrsBook.Close(); err != nil {
			logAddressBook.Error(err)
		}
	}
}

func (addrBookModel *AddrsBookModel) initAddrsBook(secType int, password string) {
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

func fromContactToQContact(contactsList []core.Contact) []*QContact {
	var qContactsList = make([]*QContact, 0)
	for _, contact := range contactsList {
		qContact := NewQContact(nil)
		qContact.SetName(contact.GetName())
		logAddressBook.Info(contact.GetID())
		qContact.SetId(contact.GetID())
		qAddressModel := NewAddrsBkAddressModel(nil)
		qAddressModel.SetAddress(fromAddressToQAddress(contact.GetAddresses()))
		qContact.SetAddress(qAddressModel)
		qContactsList = append(qContactsList, qContact)
	}
	return qContactsList
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

func (addrBookModel *AddrsBookModel) nameExist(row int, name string) bool {
	for e := range addrBookModel.Contacts() {
		if row != e && addrBookModel.Contacts()[e].Name() == name {
			return true
		}
	}
	return false
}

func (addrBookModel *AddrsBookModel) addressExist(row int, address string, coinType string) bool {
	for e := range addrBookModel.Contacts() {
		for i := range addrBookModel.Contacts()[e].Address().Address() {
			if row != e && addrBookModel.Contacts()[e].Address().Address()[i].Value() == address &&
				addrBookModel.Contacts()[e].Address().Address()[i].CoinType() == coinType {
				return true
			}
		}
	}
	return false
}

func (addrBookModel *AddrsBookModel) changeSecType(secType int, oldPassword string, newPassword string) bool {
	if err := addrsBook.ChangeSecurity(secType, oldPassword, newPassword); err != nil {
		logAddressBook.Error(err)
		return false
	}
	return true
}
