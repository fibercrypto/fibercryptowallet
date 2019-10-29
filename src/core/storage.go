package core

// KeyValueStorage provides read / write access to values given a key
type KeyValueStorage interface {
	GetValue(key string) interface{}
}

type AddressBook interface {
	Open() error
	GetContact(id uint64, password []byte) (Contact, error)
	ListContact(password []byte) ([]Contact, error)
	InsertContact(contact Contact, password []byte) error
	DeleteContact(id uint64) error
	// UpdateContact() ()
}

type Contact interface {
	EncryptContact(password, entropy []byte) ([]byte, error)
	DecryptContact(encryptMsg, password, entropy []byte) error
	GetID() uint64
	SetID(id uint64)
}
