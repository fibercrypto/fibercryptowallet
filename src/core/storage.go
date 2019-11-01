package core

// KeyValueStorage provides read / write access to values given a key
type KeyValueStorage interface {
	GetValue(key string) interface{}
}

// AddressBook provides method to manage a contact database.
type AddressBook interface {
	GetContact(id uint64) (Contact, error)
	ListContact() ([]Contact, error)
	InsertContact(contact Contact) error
	DeleteContact(id uint64) error
	UpdateContact(id uint64, contact Contact) error
}

// Contact provides encrypt / decrypt data.
type Contact interface {
	GetID() uint64
	SetID(id uint64)
	GetAddresses() []ReadableAddress
	SetAddresses([]ReadableAddress)
}

type ReadableAddress interface {
	GetValue() []byte
	SetValue(val []byte)
	GetCoinType() []byte
	SetType(val []byte)
}
