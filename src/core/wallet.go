package core

// WalletIteratr iterates over sequences of wallets
type WalletIterator interface {
	// Value of wallet at iterator pointer position
	Value() Wallet
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
}

// WalletSet allows for creating wallets and listing those in a set
type WalletSet interface {
	// ListWallets returns an iterator over wallets in the set
	ListWallets() WalletIterator
	// GetWallet to lookup wallet by ID
	GetWallet(id string) Wallet
	// CreateWallet instantiates a new wallet given account seed
	CreateWallet(name string, seed string, isEncryptrd bool, pwd PasswordReader, scanAddressesN int) (Wallet, error)
}

// WalletStorage provides access to the underlying wallets data store
type WalletStorage interface {
	// Encrypt protects wallet data using cryptography
	Encrypt(walletName string, password PasswordReader)
	// Decrypt unlocks wallet for accessing internal data
	Decrypt(walletName string, password PasswordReader)
	// IsEncrypted queries whether wallet data is encrypted or not
	IsEncrypted(walletName string) (bool, error)
}

// AddressType differentiates between BIP44 public external and internal change addresses
type AddressType uint32

const (
	// AccountAddress refers to public external address for sharing with peers
	AccountAddress AddressType = iota
	// ChangeAddress refers to internal change address
	ChangeAddress
)

// Wallet defines the contract that must be satisfied by altcoin crypto wallets
type Wallet interface {
	// GetId returns wallet local identifier
	GetId() string
	// GetLabel provides a human-readable name for this wallet to be shown in GUI controls
	GetLabel() string
	// SetLabel establishes a label for this wallet
	SetLabel(wltName string)
	// Transfer instantiates a transaction to send funds from any wallet address to known address account owner
	Transfer(to Address, amount uint64, options KeyValueStorage) (Transaction, error)
	// SendFromAddress instantiates a transaction to send funds from specific source addresses
	// to multiple destination addresses
	SendFromAddress(from []Address, to []TransactionOutput, change Address, options KeyValueStorage) (Transaction, error)
	// Spend instantiate a transaction that spends specific outputs to send to multiple destination addresses
	Spend(unspent, new []TransactionOutput, change Address, options KeyValueStorage) (Transaction, error)
	// GenAddresses discover new addresses based on default hierarchically deterministic derivation sequences
	// FIXME: Support account index to be fully compatible with BIP44
	GenAddresses(addrType AddressType, startIndex, count uint32, pwd PasswordReader) AddressIterator
	// GetCryptoAccount instantiate object to determine wallet balance and transaction history
	GetCryptoAccount() CryptoAccount
	// GetLoadedAddresses iterates over wallet addresses discovered and known to have previous history and coins
	GetLoadedAddresses() (AddressIterator, error)
	// Sign creates a new transaction by (fully or partially) choosing a strategy to sign another transaction
	Sign(txn Transaction, source UID, pwd PasswordReader, index []string) (Transaction, error)
	// AttachSignService binds a signing strategy to this wallet
	AttachSignService(TxnSigner) error
	// RemoveSignService detaches a signing strategy from this wallet
	RemoveSignService(TxnSigner) error
	// EnumerateSignServices returns an object to iterate over signing srategies attached to this wallet
	EnumerateSignServices() TxnSignerIterator
}

// WalletOutput binds transaction output to originating wallet
type WalletOutput interface {
	// GetWallet return wallet
	GetWallet() Wallet
	// GetOutput return transaction output.
	GetOutput() TransactionOutput
}

// SeedGenerator establishes the contract for generating BIP39-compatible mnemonics
type SeedGenerator interface {
	// GenerateMnemonic generates a valid BIP-39 mnemonic phrase
	GenerateMnemonic(wordCount int) (string, error)
	// VerifyMnemonic shall determine whether a given mnemonic phrase complies to BIP39
	VerifyMnemonic(seed string) (bool, error)
}

// WalletEnvironment is the entry point to manage wallets
type WalletEnv interface {
	// GetStorage provides access to wallet data store
	GetStorage() WalletStorage
	// GetWalletSet loads wallets in this environment
	GetWalletSet() WalletSet
}
