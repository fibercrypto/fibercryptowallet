package data

import (
	"encoding/json"
	"testing"
)

func TestDbWrite(t *testing.T) {
	t.Log("Init test dbWrite:")

	t.Log("Creting AddressBook:")
	AddrBook := []Contact{
		{
			Address: []Address{{
				Value: []byte("JUdRuTiqD1mGcw358twMg3VPpXpzbkdRvJ"),
				Coin:  []byte("skycoin"),
			}},
			WalletName: []byte("hsequeda"),
		}, {
			Address: []Address{{
				Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
				Coin:  []byte("skycoin"),
			}, {
				Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
				Coin:  []byte("BTC"),
			}},
			WalletName: []byte("maria"),
		},
	}

	t.Log("Writing in the DB.")
	for _, v := range AddrBook {
		b, err := json.Marshal(v.Address)
		if err != nil {
			t.Fatal("Error encoding address book to json.", err)
			return
		}
		if err = dbWrite(v.WalletName, b); err != nil {
			t.Fatal("Error writing in DB", err)
			return
		}

	}

	t.Log("Success")
}

func TestDbRead(t *testing.T) {
	t.Log("Init test dbRead:")

	var cases = []string{"hsequeda", "maria"}
	var AddressBook []Contact
	for _, v := range cases {
		AddressBookByte, err := dbRead(v)
		if err != nil {
			t.Fatal("Error reading from DB: ", err)
			return
		}
		var C Contact
		if err = json.Unmarshal(AddressBookByte, &C.Address); err != nil {
			t.Fatal("Error decoding json: ", err)
			return
		}
		C.WalletName = []byte(v)
		AddressBook = append(AddressBook, C)
	}
	for k, v := range AddressBook {
		t.Logf("Contact #%d: ContactName: %s, Addresses:", k, v.WalletName)
		for ak, av := range v.Address {
			t.Logf("Address #%d: value: %s, CoinType: %s", ak, av.Value, av.Coin)

		}
	}
	t.Log("Success")
}
