package data

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	// "github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const defaultPass = "Qwerty12345678"

func TestContact_MarshalBinary(t *testing.T) {
	type fields struct {
		ID      uint64
		Address []Address
		Name    []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{name: "empty-Contact", fields: fields{}, want: []byte(`{"ID":0,"Address":null,"Name":null}`), wantErr: false},
		{name: "not-empty-Contact", fields: fields{
			ID:      1,
			Address: []Address{{Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"), Coin: []byte("skycoin")}},
			Name:    []byte("Foo"),
		}, want: []byte(`{"ID":1,"Address":[{"Value":"SlVkUnVUaXFEMW1HY3czczU4dHdNZzNWUHBYcHpia2RSdko=",` +
			`"Coin":"c2t5Y29pbg=="}],"Name":"Rm9v"}`), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contact{
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
				Name:    tt.fields.Name,
			}
			got, err := c.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalBinary() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestContact_UnmarshalBinary(t *testing.T) {
	type fields struct {
		ID      uint64
		Address []Address
		Name    []byte
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "empty", fields: fields{}, args: args{data: []byte{}}, wantErr: true},
		{name: "not-empty", fields: fields{
			ID:      1,
			Address: []Address{{Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"), Coin: []byte("skycoin")}},
			Name:    []byte("Foo"),
		}, args: args{data: []byte(`{"Address":[{"Value":null}]}`)}, wantErr: false},
		{name: "EOF", fields: fields{ID: 2}, args: args{data: []byte(`{"ID":2,"Address":null,"Name":null}`)}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Contact{
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
				Name:    tt.fields.Name,
			}
			if err := c.UnmarshalBinary(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addressBook_DeleteContact(t *testing.T) {
	type args struct {
		id uint64
	}
	type fields struct {
		Address []Address
		Name    []byte
	}
	tests := []struct {
		name    string
		field   fields
		args    args
		wantErr bool
	}{
		{name: "empty",
			field: fields{},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{name: "one-contact",
			field: fields{
				Address: []Address{{
					Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
					Coin:  []byte("skycoin"),
				}},
				Name: []byte("Contact1"),
			}, args: args{
				id: 2},
			wantErr: false,
		},
		{name: "two-address",
			field: fields{
				Address: []Address{{
					Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
					Coin:  []byte("skycoin"),
				}, {
					Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
					Coin:  []byte("BTC"),
				}},
				Name: []byte("Contact2"),
			}, args: args{
				id: 3},
			wantErr: false,
		},
		{name: "bad-ID",
			field: fields{
				Name: []byte("Contact3"),
			}, args: args{
				id: 6},
			wantErr: true,
		},
	}

	ab := InitAddrsBook(t)
	defer CloseTest(t, ab)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ab.InsertContact(&Contact{
				Address: tt.field.Address,
				Name:    tt.field.Name,
			}); err != nil {
				t.Errorf("Error inserting contact: %s", err)
			}
			if err := ab.DeleteContact(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addressBook_GetContact(t *testing.T) {
	type fields struct {
		Address []Address
		Name    []byte
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.Contact
		wantErr bool
	}{
		{name: "empty",
			fields: fields{},
			args: args{
				id: 1,
			},
			want: &Contact{
				ID: 1,
			},
			wantErr: false},
		{name: "one-address",
			fields: fields{
				Address: []Address{{
					Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
					Coin:  []byte("skycoin"),
				}},
				Name: []byte("Contact1"),
			}, args: args{
				id: 2,
			},
			want: &Contact{
				ID: 2,
				Address: []Address{{
					Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
					Coin:  []byte("skycoin"),
				}},
				Name: []byte("Contact1"),
			},
			wantErr: false,
		},
		{name: "two-address", fields: fields{
			Address: []Address{{
				Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
				Coin:  []byte("skycoin"),
			}, {
				Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
				Coin:  []byte("BTC"),
			}},
			Name: []byte("Contact2"),
		}, args: args{
			id: 3,
		},
			want: &Contact{
				ID: 3,
				Address: []Address{{
					Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
					Coin:  []byte("skycoin"),
				}, {
					Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
					Coin:  []byte("BTC"),
				}},
				Name: []byte("Contact2"),
			},
			wantErr: false,
		},
		{name: "bad-ID",
			fields: fields{
				Name: []byte("Contact3"),
			},
			args: args{
				id: 6,
			},
			want:    nil,
			wantErr: true,
		},
	}

	ab := InitAddrsBook(t)
	defer CloseTest(t, ab)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if _, err := ab.InsertContact(&Contact{
				Address: tt.fields.Address,
				Name:    tt.fields.Name,
			}); err != nil {
				t.Errorf("Error inserting contact: %s", err)
			}

			got, err := ab.GetContact(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContact() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addressBook_InsertContact(t *testing.T) {
	type args struct {
		c core.Contact
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "empty",
			args: args{
				c: &Contact{},
			},
			wantErr: false,
		}, {name: "one-address",
			args: args{
				c: &Contact{
					Address: []Address{{
						Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				},
			},
			wantErr: false,
		}, {name: "two-address",
			args: args{
				c: &Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
						Coin:  []byte("BTC"),
					}},
					Name: []byte("Contact2"),
				},
			},
			wantErr: false,
		}, {name: "repeat-address-diff-type",
			args: args{
				c: &Contact{
					Address: []Address{{
						Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact3"),
				},
			},
			wantErr: false,
		}, {name: "repeat-address-same-type",
			args: args{
				c: &Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact4"),
				},
			},
			wantErr: true,
		},
		{name: "repeat-name",
			args: args{
				c: &Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact4"),
				},
			},
			wantErr: true,
		},
	}
	ab := InitAddrsBook(t)
	defer CloseTest(t, ab)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ab.InsertContact(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("InsertContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_addressBook_ListContact(t *testing.T) {
	type fields struct {
		Contacts []Contact
	}
	type args struct {
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []core.Contact
		wantErr bool
	}{
		{name: "empty",
			fields:  fields{Contacts: []Contact(nil)},
			want:    []core.Contact(nil),
			wantErr: false},
		{name: "one-contact",
			fields: fields{Contacts: []Contact{
				{Address: []Address{{
					Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
					Coin:  []byte("skycoin"),
				}},
					Name: []byte("contact1"),
				},
			}},
			want: []core.Contact{&Contact{
				ID: 1,
				Address: []Address{{
					Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
					Coin:  []byte("skycoin"),
				}},
				Name: []byte("contact1"),
			}},
			wantErr: false},
		{name: "multiple-contacts",
			fields: fields{Contacts: []Contact{
				{Address: []Address{{
					Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
					Coin:  []byte("skycoin"),
				}},
					Name: []byte("contact1"),
				},
				{Address: []Address{{
					Value: []byte("9BSEAEE3XGtQ2X43BCT2XCYgheGLQQigEG"),
					Coin:  []byte("skycoin"),
				}, {
					Value: []byte("29cnQPHuWHCRF26LEAb2gR83ywnF3F9HduW"),
					Coin:  []byte("skycoin")}},
					Name: []byte("contact2"),
				},
				{Address: []Address{{
					Value: []byte("2ymjULRdbiFoUNJKNhWbQ3JqdE8TXnZkyU"),
					Coin:  []byte("BTC"),
				}},
					Name: []byte("contact3"),
				}, {
					Address: []Address{{
						Value: []byte("oHvj7oy8maES9HJiQHJTp4GvcUcpz3voDq"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2SGMfTFV2zbQzGw7aJm1D5EeEPgych5ixuC"),
						Coin:  []byte("skycoin")}},
					Name: []byte("contact4"),
				}, {
					Address: []Address{{
						Value: []byte("2DpeofcsamDfanrRz34qjYvskRzKqzNKMcj"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2EVNa4CK9SKosT4j1GEn8SuuUUEAXaHAMbM"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("n5SteDkkYdR3VJtMnVYcQ45L16rDDrseG8"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("contact5"),
				}, {Address: []Address{{
					Value: []byte("25MP2EHddPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
					Coin:  []byte("skycoin"),
				}},
					Name: []byte("contact6"),
				},
				{Address: []Address{{
					Value: []byte("9BSEAEEr3XddGtQ2X43BCT2XCYgheGLQQigEG"),
					Coin:  []byte("skycoin"),
				}, {
					Value: []byte("29cnQPHuweWHCRF26LEAb2gR83ywnF3F9HduW"),
					Coin:  []byte("skycoin")}},
					Name: []byte("contact7"),
				},
				{Address: []Address{{
					Value: []byte("2ymjULRdbiasdFoUNJKNhWbQ3JqdE8TXnZkyU"),
					Coin:  []byte("BTC"),
				}},
					Name: []byte("contact8"),
				}, {
					Address: []Address{{
						Value: []byte("oHvrwfj7oy8maES9HJiQHJTp4GvcUcpz3voDq"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2SGMfTsaazFV2zbQzGw7aJm1D5EeEPgych5ixuC"),
						Coin:  []byte("skycoin")}},
					Name: []byte("contact9"),
				}, {
					Address: []Address{{
						Value: []byte("2DprrreofcsamDfanrRz34qjYvskRzKqzNKMcj"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2EVNa4esdfsCK9SKosT4j1GEn8SuuUUEAXaHAMbM"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("n5SteDkkYdRfsf3VJtMnVYcQ45L16rDDrseG8"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("contact10"),
				},
			}},
			want: []core.Contact{
				&Contact{
					ID: 1,
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("contact1"),
				},
				&Contact{
					ID: 2,
					Address: []Address{{
						Value: []byte("9BSEAEE3XGtQ2X43BCT2XCYgheGLQQigEG"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("29cnQPHuWHCRF26LEAb2gR83ywnF3F9HduW"),
						Coin:  []byte("skycoin")}},
					Name: []byte("contact2"),
				},
				&Contact{
					ID: 3,
					Address: []Address{{
						Value: []byte("2ymjULRdbiFoUNJKNhWbQ3JqdE8TXnZkyU"),
						Coin:  []byte("BTC"),
					}},
					Name: []byte("contact3"),
				},
				&Contact{
					ID: 4,
					Address: []Address{{
						Value: []byte("oHvj7oy8maES9HJiQHJTp4GvcUcpz3voDq"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2SGMfTFV2zbQzGw7aJm1D5EeEPgych5ixuC"),
						Coin:  []byte("skycoin")}},
					Name: []byte("contact4"),
				},
				&Contact{
					ID: 5,
					Address: []Address{{
						Value: []byte("2DpeofcsamDfanrRz34qjYvskRzKqzNKMcj"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2EVNa4CK9SKosT4j1GEn8SuuUUEAXaHAMbM"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("n5SteDkkYdR3VJtMnVYcQ45L16rDDrseG8"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("contact5"),
				}, &Contact{
					ID: 6,
					Address: []Address{{
						Value: []byte("25MP2EHddPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("contact6"),
				},
				&Contact{
					ID: 7,
					Address: []Address{{
						Value: []byte("9BSEAEEr3XddGtQ2X43BCT2XCYgheGLQQigEG"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("29cnQPHuweWHCRF26LEAb2gR83ywnF3F9HduW"),
						Coin:  []byte("skycoin")}},
					Name: []byte("contact7"),
				},
				&Contact{
					ID: 8,
					Address: []Address{{
						Value: []byte("2ymjULRdbiasdFoUNJKNhWbQ3JqdE8TXnZkyU"),
						Coin:  []byte("BTC"),
					}},
					Name: []byte("contact8"),
				}, &Contact{
					ID: 9,
					Address: []Address{{
						Value: []byte("oHvrwfj7oy8maES9HJiQHJTp4GvcUcpz3voDq"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2SGMfTsaazFV2zbQzGw7aJm1D5EeEPgych5ixuC"),
						Coin:  []byte("skycoin")}},
					Name: []byte("contact9"),
				}, &Contact{
					ID: 10,
					Address: []Address{{
						Value: []byte("2DprrreofcsamDfanrRz34qjYvskRzKqzNKMcj"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("2EVNa4esdfsCK9SKosT4j1GEn8SuuUUEAXaHAMbM"),
						Coin:  []byte("skycoin"),
					}, {
						Value: []byte("n5SteDkkYdRfsf3VJtMnVYcQ45L16rDDrseG8"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("contact10"),
				},
			},
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ab := InitAddrsBook(t)
			defer CloseTest(t, ab)
			for _, contact := range tt.fields.Contacts {
				if _, err := ab.InsertContact(&contact); err != nil {
					t.Error(err)
					return
				}
			}
			got, err := ab.ListContact()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListContact() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdateContact(t *testing.T) {
	type args struct {
		id         uint64
		newContact core.Contact
	}
	type insertArgs struct {
		contacts []core.Contact
	}
	tests := []struct {
		name       string
		args       args
		insertArgs insertArgs
		wantErr    bool
	}{
		{name: "empty-insert",
			args: args{
				id: 1,
				newContact: &Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				},
			},
			insertArgs: insertArgs{contacts: []core.Contact{&Contact{}}},
			wantErr:    false,
		},
		{name: "empty-update",
			args: args{
				id:         1,
				newContact: &Contact{},
			},
			insertArgs: insertArgs{
				contacts: []core.Contact{&Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				}}},
			wantErr: false,
		},
		{name: "update-coinType",
			args: args{
				id: 1,
				newContact: &Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("BTC"),
					}},
					Name: []byte("Contact1"),
				},
			},
			insertArgs: insertArgs{
				contacts: []core.Contact{&Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				}}},
			wantErr: false,
		},
		{name: "update-address",
			args: args{
				id: 1,
				newContact: &Contact{
					Address: []Address{{
						Value: []byte("n5SteDkkYdR3VJtMnVYcQ45L16rDDrseG8"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				},
			},
			insertArgs: insertArgs{
				contacts: []core.Contact{&Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				}}},
			wantErr: false,
		},
		{name: "same-contact",
			args: args{
				id: 1,
				newContact: &Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				},
			},
			insertArgs: insertArgs{
				contacts: []core.Contact{&Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				}}},
			wantErr: false,
		}, {name: "repeat-address",
			args: args{
				id: 1,
				newContact: &Contact{
					Address: []Address{{
						Value: []byte("n5SteDkkYdR3VJtMnVYcQ45L16rDDrseG8"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				},
			},
			insertArgs: insertArgs{
				contacts: []core.Contact{&Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				},
					&Contact{
						Address: []Address{{
							Value: []byte("n5SteDkkYdR3VJtMnVYcQ45L16rDDrseG8"),
							Coin:  []byte("skycoin"),
						}},
						Name: []byte("Contact2"),
					}}},
			wantErr: true,
		},
		{name: "repeat-name",
			args: args{
				id: 1,
				newContact: &Contact{
					Address: []Address{{
						Value: []byte("9BSEAEE3XGtQ2X43BCT2XCYgheGLQQigEG"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact2"),
				},
			},
			insertArgs: insertArgs{
				contacts: []core.Contact{&Contact{
					Address: []Address{{
						Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
						Coin:  []byte("skycoin"),
					}},
					Name: []byte("Contact1"),
				},
					&Contact{
						Address: []Address{{
							Value: []byte("n5SteDkkYdR3VJtMnVYcQ45L16rDDrseG8"),
							Coin:  []byte("skycoin"),
						}},
						Name: []byte("Contact2"),
					}}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ab := InitAddrsBook(t)
			defer CloseTest(t, ab)
			for e := range tt.insertArgs.contacts {
				if _, err := ab.InsertContact(tt.insertArgs.contacts[e]); err != nil {
					t.Error(err)
					return
				}
			}

			if err := ab.UpdateContact(tt.args.id, tt.args.newContact); (err != nil) != tt.wantErr {
				t.Errorf("UpdateContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Init(t *testing.T) {
	type args struct {
		secType  int
		password string
	}
	tests := []struct {
		name string
		// fields  fields
		args    args
		wantErr bool
	}{
		{name: "Type 1", args: args{
			secType:  Type1,
			password: "",
		}, wantErr: false},
		{name: "Type 2", args: args{
			secType:  Type2,
			password: "",
		}, wantErr: false},
		{name: "Type 3", args: args{
			secType:  Type3,
			password: defaultPass,
		}, wantErr: false},
		{name: "Type wrong", args: args{
			secType:  -1,
			password: "",
		}, wantErr: true},
		{name: "Two Init", args: args{
			secType:  Type3,
			password: defaultPass,
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			addressBook, err := NewAddressBook(GetFilePath(t))
			if err != nil {
				t.Error(addressBook)
			}

			if tt.name == "Two Init" {
				_ = addressBook.Init(tt.args.secType, tt.args.password)
			}

			defer CloseTest(t, addressBook)

			if err := addressBook.Init(tt.args.secType, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Generate a temporal file and return its path.
func GetFilePath(t *testing.T) string {
	home := os.Getenv("HOME")

	f, err := ioutil.TempFile(home+"/temp", "testaddressbook-")
	if err != nil {
		t.Error(err)
	}

	if err := f.Close(); err != nil {
		t.Error(err)
	}
	return f.Name()
}

// Open a address book using a test file.
func InitAddrsBook(t *testing.T) core.AddressBook {
	path := GetFilePath(t)
	AddrsBook, err := NewAddressBook(path)
	if err != nil {
		t.Error(err)
	}
	err = AddrsBook.Init(Type3, defaultPass)
	if err != nil {
		t.Error(err)
	}

	return AddrsBook
}

func CloseTest(t *testing.T, ab core.AddressBook) {
	path := ab.GetPath()
	if err := ab.Close(); err != nil {
		t.Errorf("Error closing db: %s", err)
	}
	if err := os.Remove(path); err != nil {
		t.Error(err)
	}
}