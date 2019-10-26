package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var AddressBookTestCases = []Contact{
	{
		Address: []Address{{
			Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
			Coin:  []byte("skycoin"),
		}},
		Name: []byte("hsequeda"),
	}, {
		Address: []Address{{
			Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
			Coin:  []byte("skycoin"),
		}},
		Name: []byte("hsequeda"),
	},
	{
		Address: []Address{{
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

func TestAddressBook_generateHash(t *testing.T) {
	t.Log("Init test generateHash")
	ab := AddressBook{}

	testCases := make(map[string]string)
	testCases["empty-address"] = ""
	testCases["ok"] = "qwerty"

	for k, v := range testCases {
		t.Logf("Generating hash for %s...", k)
		// Will be generated a hash.
		if err := ab.generateHash([]byte(v)); err != nil {
			t.Error(err)
		}

		t.Logf("Verify hash for %s...", k)
		// Will be verify the hash
		if err := ab.verifyHash([]byte(v)); err != nil {
			t.Error(err)
		}
	}

	t.Log("Success")
}

func TestAddressBook_Insert(t *testing.T) {
	t.Log("Init test Insert:")
	ab := OpenAddrsBook(t)
	t.Log("Writing in the AddressBook.")
	for k := range AddressBookTestCases {
		if err := ab.Insert(&AddressBookTestCases[k], pass); err != nil {
			t.Errorf("Error inserting value into database: %s", err)
		}
	}
	defer func() {
		if err := ab.Close(); err != nil {
			t.Errorf("Error closing db: %s", err)
		}
	}()

	// Verify user can be retrived.
	other, err := ab.Get(1, pass)
	t.Log(AddressBookTestCases[0].ID)
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(other, &AddressBookTestCases[0]) {
		t.Errorf("unexpected user: %#v \n and have:%#v", other, AddressBookTestCases[0])
	}

	t.Log("Success")
}

func TestContact_MarshalBinary(t *testing.T) {
	var c = Contact{
		ID: 1,
		Address: []Address{{
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
	var newC Contact
	if err := newC.UnmarshalBinary(b); err != nil {
		t.Error(err)
	}
	assert.ObjectsAreEqual(c, newC)

	t.Log("Success")
}

func TestAddressBook_MarshalBinary(t *testing.T) {
	var ab = AddressBook{
		hasPassword: false,
		hash:        []byte("myhash"),
		entropy:     []byte("entropy"),
	}

	b, err := ab.MarshalBinary()
	t.Log(string(b))
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, b)
	var newAB AddressBook
	if err := newAB.UnmarshalBinary(b); err != nil {
		t.Error(err)
	}
	assert.ObjectsAreEqual(ab, newAB)

	t.Log("Success")

}

func TestAddressBook_List(t *testing.T) {
	t.Log("Init test Insert:")
	ab := OpenAddrsBook(t)
	t.Log("Reading AddressBook.")
	contactList, err := ab.List(pass)
	if err != nil {
		t.Error(err)
	}
	for _, v := range contactList {
		t.Logf("ID:%d \n", v.ID)
		t.Logf("Name:%s \n", v.Name)
	}
	defer func() {
		if err := ab.Close(); err != nil {
			t.Errorf("Error closing db: %s", err)
		}
	}()
}

type TestAddressBook struct {
	*AddressBook
}

func NewTestAddressBook(t *testing.T) *TestAddressBook {
	f, err := ioutil.TempFile("/home/hsequeda/temp", "testaddressbook-")
	if err != nil {
		t.Error(err)
	}

	if err := f.Close(); err != nil {
		t.Error(err)
	}

	return &TestAddressBook{
		AddressBook: &AddressBook{
			dbPath: f.Name(),
		},
	}
}

func OpenAddrsBook(t *testing.T) *TestAddressBook {
	AddrsBook := NewTestAddressBook(t)
	if err := AddrsBook.InsertPass(pass); err != nil {
		t.Error(err)
	}
	if err := AddrsBook.GenerateMnemonic("mandate ride tide eternal laundry stem prison era calm topic rate remain"); err != nil {
		t.Error(err)
	}

	if err := AddrsBook.Open(); err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v \n", string(AddrsBook.AddressBook.entropy))
	return AddrsBook
}

func (s *TestAddressBook) Close() error {
	defer func() {
		if err := os.Remove(s.dbPath); err != nil {
			panic(err)
		}
	}()
	return s.AddressBook.Close()
}
