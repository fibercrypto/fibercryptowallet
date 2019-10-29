package data_test

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var AddressBookTestCases = []data.Contact{
	{
		Address: []data.Address{{
			Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
			Coin:  []byte("skycoin"),
		}},
		Name: []byte("hsequeda"),
	}, {
		Address: []data.Address{{
			Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
			Coin:  []byte("skycoin"),
		}},
		Name: []byte("hsequeda"),
	},
	{
		Address: []data.Address{{
			Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
			Coin:  []byte("skycoin"),
		}, {
			Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
			Coin:  []byte("BTC"),
		}},
		Name: []byte("maria"),
	},
}

var pass = []byte("qwerty")

func TestAddressBook_Insert(t *testing.T) {
	t.Log("New test Insert:")
	ab := data.GetAddressBook()
	OpenAddrsBook(t)
	t.Log("Writing in the addressBook.")
	for k := range AddressBookTestCases {
		if err := ab.InsertContact(&AddressBookTestCases[k], pass); err != nil {
			t.Errorf("Error inserting value into database: %s", err)
		}
	}
	defer func() {
		if err := ab.Close(); err != nil {
			t.Errorf("Error closing db: %s", err)
		}
		if err := os.Remove(ab.GetPath()); err != nil {
			t.Error(err)
		}
	}()
	// Verify user can be retrived.
	other, err := ab.GetContact(1, pass)
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(other, &AddressBookTestCases[0]) {
		t.Errorf("unexpected user: %#v \n and have:%#v", other, AddressBookTestCases[0])
	}

	t.Log("Success")
}

func TestContact_MarshalBinary(t *testing.T) {
	var c = data.Contact{
		ID: 1,
		Address: []data.Address{{
			Value: []byte("myaddress"),
			Coin:  []byte("sky"),
		}},
		Name: []byte("qwerty"),
	}
	b, err := c.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, b)
	var newC data.Contact
	if err := newC.UnmarshalBinary(b); err != nil {
		t.Error(err)
	}
	assert.ObjectsAreEqual(c, newC)

	t.Log("Success")
}

func TestAddressBook_List(t *testing.T) {
	t.Log("New test Insert:")
	ab := data.GetAddressBook()
	OpenAddrsBook(t)
	for e := range AddressBookTestCases {
		if err := ab.InsertContact(&AddressBookTestCases[e], pass); err != nil {
			t.Error(err)
		}
	}
	t.Log("Reading addressBook.")
	contactList, err := ab.ListContact(pass)
	if err != nil {
		t.Error(err)
	}
	for _, v := range contactList {
		t.Logf("ID:%d/3 \n", v.GetID())
		// t.Logf("Name:%s \n", v.Name)
	}
	defer func() {
		if err := ab.Close(); err != nil {
			t.Errorf("Error closing db: %s", err)
		}
		if err := os.Remove(ab.GetPath()); err != nil {
			t.Error(err)
		}
	}()
	t.Log("Success")
}

func GetFilePath(t *testing.T) string {
	f, err := ioutil.TempFile("/home/hsequeda/temp", "testaddressbook-")
	if err != nil {
		t.Error(err)
	}

	if err := f.Close(); err != nil {
		t.Error(err)
	}
	return f.Name()
}

func OpenAddrsBook(t *testing.T) {
	AddrsBook := data.GetAddressBook()
	path := GetFilePath(t)
	mnemonic := "mandate ride tide eternal laundry stem prison era calm topic rate remain"
	if err := AddrsBook.New(pass, path, mnemonic); err != nil {
		t.Error(err)
	}
}
