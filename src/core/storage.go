package core

type KeyValueStorage interface {
	GetValue(key string) interface{}
}

// AddressBook provides method to manage a contact database.
type AddressBook interface {
	Init(secType int, password string) error
	Authenticate(password string) error
	GetContact(id uint64) (Contact, error)
	ListContact() ([]Contact, error)
	InsertContact(contact Contact) (uint64, error)
	DeleteContact(id uint64) error
	UpdateContact(id uint64, contact Contact) error
	GetPath() string
	HasInit() bool
	GetSecType() int
	Close() error
	Exist(path string) bool
}

// Contact provides encrypt / decrypt data.
type Contact interface {
	GetID() uint64
	SetID(id uint64)
	GetAddresses() []StringAddress
	SetAddresses([]StringAddress)
	GetName() string
	SetName(string)
}
type StringAddress interface {
	GetValue() []byte
	SetValue(val []byte)
	GetCoinType() []byte
	SetCoinType(val []byte)
}
