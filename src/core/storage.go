package core

type KeyValueStorage interface {
	GetValue(key string) interface{}
}

// AddressBook provides method to manage a contact database.
type AddressBook interface {
	Init(secType int, password string) error
	Authenticate(password string) error
	ChangeSecurity(NewSecType int, oldPassword, newPassword string) error
	GetContact(id uint64) (Contact, error)
	ListContact() ([]Contact, error)
	InsertContact(contact Contact) (uint64, error)
	DeleteContact(id uint64) error
	UpdateContact(id uint64, contact Contact) error
	GetStorage() Storage
	HasInit() bool
	IsOpen() bool
	GetSecType() int
	Close() error
}

type Storage interface {
	InsertValue(value interface{}) (uint64, error)
	GetValue(key uint64) (interface{}, error)
	ListValues() (map[uint64]interface{}, error)
	DeleteValue(key uint64) error
	UpdateValue(key uint64, newValue interface{}) error
	Path() string
	GetConfig() map[string]string
	InsertConfig(map[string]string) error
	Close() error
}

// Contact provides encrypt / decrypt data.
type Contact interface {
	GetID() uint64
	SetID(id uint64)
	GetAddresses() []StringAddress
	SetAddresses([]StringAddress)
	GetName() string
	SetName(string)
	IsValid() bool
}
type StringAddress interface {
	GetValue() []byte
	SetValue(val []byte)
	GetCoinType() []byte
	SetCoinType(val []byte)
	IsValid() bool
}
