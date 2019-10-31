package data

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
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
		{name: "empty-Contact", fields: fields{}, want: []byte{}, wantErr: false},
		{name: "not-empty-Contact", fields: fields{
			ID:      1,
			Address: []Address{{Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"), Coin: []byte("skycoin")}},
			Name:    []byte("Foo"),
		}, want: []byte{0x8, 0x1, 0x12, 0x3, 0x46, 0x6f, 0x6f, 0x1a, 0x2e, 0xa, 0x23, 0x4a, 0x55, 0x64, 0x52, 0x75,
			0x54, 0x69, 0x71, 0x44, 0x31, 0x6d, 0x47, 0x63, 0x77, 0x33, 0x73, 0x35, 0x38, 0x74, 0x77, 0x4d, 0x67, 0x33,
			0x56, 0x50, 0x70, 0x58, 0x70, 0x7a, 0x62, 0x6b, 0x64, 0x52, 0x76, 0x4a, 0x12, 0x7, 0x73, 0x6b, 0x79, 0x63,
			0x6f, 0x69, 0x6e}, wantErr: false},
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
				t.Errorf("MarshalBinary() got = %#v, want %v", got, string(tt.want))
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
		{name: "empty", fields: fields{}, args: args{data: []byte{}}, wantErr: false},
		{name: "not-empty", fields: fields{
			ID:      1,
			Address: []Address{{Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"), Coin: []byte("skycoin")}},
			Name:    []byte("Foo"),
		}, args: args{data: []byte{0x8, 0x1, 0x12, 0x3, 0x46, 0x6f, 0x6f, 0x1a, 0x2e, 0xa, 0x23, 0x4a, 0x55, 0x64, 0x52, 0x75,
			0x54, 0x69, 0x71, 0x44, 0x31, 0x6d, 0x47, 0x63, 0x77, 0x33, 0x73, 0x35, 0x38, 0x74, 0x77, 0x4d, 0x67, 0x33,
			0x56, 0x50, 0x70, 0x58, 0x70, 0x7a, 0x62, 0x6b, 0x64, 0x52, 0x76, 0x4a, 0x12, 0x7, 0x73, 0x6b, 0x79, 0x63,
			0x6f, 0x69, 0x6e}}, wantErr: false},
		{name: "EOF", fields: fields{ID: 2}, args: args{data: []byte{0x12, 0x3, 0x46, 0x6f, 0x6f, 0x1a, 0x2e, 0xa, 0x23,
			0x4a, 0x55, 0x64, 0x52, 0x75, 0x54, 0x69, 0x71, 0x44, 0x31, 0x6d, 0x47, 0x63, 0x77, 0x33}}, wantErr: true},
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
		{name: "empty", field: fields{}, args: args{id: 1}, wantErr: false},
		{name: "one-contact", field: fields{
			Address: []Address{{
				Value: []byte("JUdRuTiqD1mGcw3s58twMg3VPpXpzbkdRvJ"),
				Coin:  []byte("skycoin"),
			}},
			Name: []byte("Contact1"),
		}, args: args{id: 2}, wantErr: false},
		{name: "two-address", field: fields{
			Address: []Address{{
				Value: []byte("25MP2EHPZyfEqUnXfapgUj1TQfZVXdn5RrZ"),
				Coin:  []byte("skycoin"),
			}, {
				Value: []byte("2TFC2Ktc6Y3UAUqo7WGA55X6mqoKZRaFp9s"),
				Coin:  []byte("BTC"),
			}},
			Name: []byte("Contact2"),
		}, args: args{id: 3}, wantErr: false},
		{name: "bad-id", field: fields{
			Name: []byte("Contact3"),
		}, args: args{id: 6}, wantErr: true},
	}

	ab := OpenAddrsBook(t)
	defer CloseTest(t, &ab)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ab.InsertContact(&Contact{
				Address: tt.field.Address,
				Name:    tt.field.Name,
			}, []byte(defaultPass)); err != nil {
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
		id       uint64
		password []byte
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
				id:       1,
				password: []byte(defaultPass),
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
			id:       2,
			password: []byte(defaultPass),
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
			id:       3,
			password: []byte(defaultPass),
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
		{name: "bad-id",
			fields: fields{
				Name: []byte("Contact3"),
			},
			args: args{
				id:       6,
				password: []byte(defaultPass),
			},
			want:    nil,
			wantErr: true,
		},
		{name: "bad-password",
			fields: fields{
				Name: []byte("Contact3"),
				Address: []Address{{
					Value: []byte("2TFC2Ktc6UAUqo7WGA55X6mqoKZRaFp9s"),
					Coin:  []byte("BTC"),
				}},
			},
			args: args{
				id:       5,
				password: []byte("wrongPass"),
			},
			want:    nil,
			wantErr: true,
		},
	}

	ab := OpenAddrsBook(t)
	defer CloseTest(t, &ab)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := ab.InsertContact(&Contact{
				Address: tt.fields.Address,
				Name:    tt.fields.Name,
			}, []byte(defaultPass)); err != nil {
				t.Errorf("Error inserting contact: %s", err)
			}

			got, err := ab.GetContact(tt.args.id, tt.args.password)
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

// func Test_addressBook_GetHashFromConfig(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.GetHashFromConfig(); (err != nil) != tt.wantErr {
// 				t.Errorf("GetHashFromConfig() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_GetPath(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if got := ab.GetPath(); got != tt.want {
// 				t.Errorf("GetPath() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_InsertContact(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		c        core.Contact
// 		password []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.InsertContact(tt.args.c, tt.args.password); (err != nil) != tt.wantErr {
// 				t.Errorf("InsertContact() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_ListContact(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		password []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    []core.Contact
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			got, err := ab.ListContact(tt.args.password)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ListContact() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ListContact() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_New(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		password []byte
// 		path     string
// 		mnemonic string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.Init(tt.args.password, tt.args.path, tt.args.mnemonic); (err != nil) != tt.wantErr {
// 				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_Open(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.Open(); (err != nil) != tt.wantErr {
// 				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_SearchAddress(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		address  []byte
// 		coin     []byte
// 		password []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.SearchAddress(tt.args.address, tt.args.coin, tt.args.password); (err != nil) != tt.wantErr {
// 				t.Errorf("SearchAddress() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_UpdateContact(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		id         uint64
// 		newContact core.Contact
// 		password   []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.UpdateContact(tt.args.id, tt.args.newContact, tt.args.password); (err != nil) != tt.wantErr {
// 				t.Errorf("UpdateContact() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_decryptAESGCM(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		cipherMsg []byte
// 		password  []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    core.Contact
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			got, err := ab.decryptAESGCM(tt.args.cipherMsg, tt.args.password)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("decryptAESGCM() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("decryptAESGCM() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_encryptAESGCM(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		c        *Contact
// 		password []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			got, err := ab.encryptAESGCM(tt.args.c, tt.args.password)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("encryptAESGCM() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("encryptAESGCM() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_genEntropy(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		mnemonic string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.genEntropy(tt.args.mnemonic); (err != nil) != tt.wantErr {
// 				t.Errorf("genEntropy() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_addressBook_verifyHash(t *testing.T) {
// 	type fields struct {
// 		db      *bolt.DB
// 		dbPath  string
// 		hash    []byte
// 		entropy []byte
// 	}
// 	type args struct {
// 		password []byte
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ab := &AddressBookWithBolt{
// 				db:      tt.fields.db,
// 				dbPath:  tt.fields.dbPath,
// 				hash:    tt.fields.hash,
// 				entropy: tt.fields.entropy,
// 			}
// 			if err := ab.verifyHash(tt.args.password); (err != nil) != tt.wantErr {
// 				t.Errorf("verifyHash() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_derivePassphrase(t *testing.T) {
// 	type args struct {
// 		entropy  []byte
// 		password []byte
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []byte
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := derivePassphrase(tt.args.entropy, tt.args.password); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("derivePassphrase() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// Generate a temporal file and return its path.
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

// Open a address book using a test file.
func OpenAddrsBook(t *testing.T) AddressBookWithBolt {
	AddrsBook := AddressBookWithBolt{}
	path := GetFilePath(t)
	mnemonic := "mandate ride tide eternal laundry stem prison era calm topic rate remain"
	if err := AddrsBook.Init([]byte(defaultPass), path, mnemonic); err != nil {
		t.Error(err)
	}
	return AddrsBook
}

func CloseTest(t *testing.T, ab *AddressBookWithBolt) {
	if err := ab.Close(); err != nil {
		t.Errorf("Error closing db: %s", err)
	}
	if err := os.Remove(ab.GetPath()); err != nil {
		t.Error(err)
	}
}
