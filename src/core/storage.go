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
