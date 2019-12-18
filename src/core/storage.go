package core

// KeyValueStore provides read / write access to values given a key
type KeyValueStore interface {
	// GetValue lookup value for key
	GetValue(key string) interface{}
	// SetValue bind value o known key
	SetValue(key string, value interface{})
}

const (
	// StrSignerID option key for storing signer ID
	StrSignerID = "signer.id"
	// StrMethodType option key for method type name
	StrTypeName = "api.typename"
	// StrMethodName option key for method name
	StrMethodName = "api.method"
	// StrWalletLabel option key for wallet label
	StrWalletLabel = "wallet.label"
	// StrWalletName option key for wallet name
	StrWalletName = "wallet.id"
	// StrCoinTicker option key for coin ticker ID
	StrCoinTicker = "coin.ticker"
	// StrSenderObject option key for object that triggered an action
	StrSenderObject = "call.self"

	// TypeNameAddress Address type name
	TypeNameAddress = "Address"
	// TypeNameAddressIterator AddressIterator type name
	TypeNameAddressIterator = "AddressIterator"
	// TypeNameAltcoinManager AltcoinManager type name
	TypeNameAltcoinManager = "AltcoinManager"
	// TypeNameAltcoinPlugin AltcoinPlugin type name
	TypeNameAltcoinPlugin = "AltcoinPlugin"
	// TypeNameBlock Block type name
	TypeNameBlock = "Block"
	// TypeNameBlockchainSignService BlockchainSignService type name
	TypeNameBlockchainSignService = "BlockchainSignService"
	// TypeNameBlockchainStatus BlockchainStatus type name
	TypeNameBlockchainStatus = "BlockchainStatus"
	// TypeNameBlockchainTransactionAPI BlockchainTransactionAPI type name
	TypeNameBlockchainTransactionAPI = "BlockchainTransactionAPI"
	// TypeNameCryptoAccount CryptoAccount type name
	TypeNameCryptoAccount = "CryptoAccount"
	// TypeNameKeyValueStore KeyValueStore type name
	TypeNameKeyValueStore = "KeyValueStore"
	// TypeNameMultiPool MultiPool type name
	TypeNameMultiPool = "MultiPool"
	// TypeNameMultiPoolSection MultiPoolSection type name
	TypeNameMultiPoolSection = "MultiPoolSection"
	// TypeNamePEX PEX type name
	TypeNamePEX = "PEX"
	// TypeNamePexNode PexNode type name
	TypeNamePexNode = "PexNode"
	// TypeNamePexNodeIterator PexNodeIterator type name
	TypeNamePexNodeIterator = "PexNodeIterator"
	// TypeNamePexNodeSet PexNodeSet type name
	TypeNamePexNodeSet = "PexNodeSet"
	// TypeNamePooledObjectFactory PooledObjectFactory type name
	TypeNamePooledObjectFactory = "PooledObjectFactory"
	// TypeNameSeedGenerator SeedGenerator type name
	TypeNameSeedGenerator = "SeedGenerator"
	// TypeNameTransaction Transaction type name
	TypeNameTransaction = "Transaction"
	// TypeNameTransactionInput TransactionInput type name
	TypeNameTransactionInput = "TransactionInput"
	// TypeNameTransactionInputIterator TransactionInputIterator type name
	TypeNameTransactionInputIterator = "TransactionInputIterator"
	// TypeNameTransactionIterator TransactionIterator type name
	TypeNameTransactionIterator = "TransactionIterator"
	// TypeNameTransactionOutput TransactionOutput type name
	TypeNameTransactionOutput = "TransactionOutput"
	// TypeNameTransactionOutputIterator TransactionOutputIterator type name
	TypeNameTransactionOutputIterator = "TransactionOutputIterator"
	// TypeNameTxnSigner TxnSigner type name
	TypeNameTxnSigner = "TxnSigner"
	// TypeNameTxnSignerIterator TxnSignerIterator type name
	TypeNameTxnSignerIterator = "TxnSignerIterator"
	// TypeNameWallet Wallet type name
	TypeNameWallet = "Wallet"
	// TypeNameWalletAddress WalletAddress type name
	TypeNameWalletAddress = "WalletAddress"
	// TypeNameWalletEnv WalletEnv type name
	TypeNameWalletEnv = "WalletEnv"
	// TypeNameWalletIterator WalletIterator type name
	TypeNameWalletIterator = "WalletIterator"
	// TypeNameWalletOutput WalletOutput type name
	TypeNameWalletOutput = "WalletOutput"
	// TypeNameWalletSet WalletSet type name
	TypeNameWalletSet = "WalletSet"
	// TypeNameWalletStorage WalletStorage type name
	TypeNameWalletStorage = "WalletStorage"
)

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
	Encrypt(secType int, key []byte) ([]byte, error)
	Decrypt(secType int, data, key []byte) error
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
