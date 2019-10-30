package core

type KeyValueStorage interface {
	GetValue(key string) interface{}
}

// AddressBook provides method to manage a contact database.
type AddressBook interface {
	Open() error
	GetContact(id uint64, password []byte) (Contact, error)
	ListContact(password []byte) ([]Contact, error)
	InsertContact(contact Contact, password []byte) error
	DeleteContact(id uint64) error
	// UpdateContact() ()
}

// Contact provides Encrypt / Decrypt data.
type Contact interface {
	GetID() uint64
	SetID(id uint64)
	GetAddress(pos int64) ReadableAddress
	SetAddress(ReadableAddress)
}

type ReadableAddress interface {
	GetValue() []byte
	SetValue(val []byte)
	GetCoinType() []byte
	SetType(val []byte)
}
