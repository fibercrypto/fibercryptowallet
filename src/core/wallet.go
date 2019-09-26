package core

type WalletIterator interface {
	Value() Wallet
	Next() bool
	HasNext() bool
}

type WalletSet interface {
	ListWallets() WalletIterator
	GetWallet(id string) Wallet
	CreateWallet(name string, seed string, isEncryptrd bool, pwd PasswordReader, scanAddressesN int) (Wallet, error)
}

type WalletStorage interface {
	Encrypt(walletName string, password PasswordReader)
	Decrypt(walletName string, password PasswordReader)
	IsEncrypted(walletName string) (bool, error)
}

type AddressType uint32

const (
	AccountAddress AddressType = iota
	ChangeAddress
)

type Wallet interface {
	GetId() string
	GetLabel() string
	SetLabel(wltName string)
	Transfer(to Address, amount uint64, options KeyValueStorage) (Transaction, error)
	SendFromAddress(from []Address, to []TransactionOutput, change Address, options KeyValueStorage) (Transaction, error)
	Spend(unspent, new []TransactionOutput, change Address, options KeyValueStorage) (Transaction, error)
	GenAddresses(addrType AddressType, startIndex, count uint32, pwd PasswordReader) AddressIterator
	GetCryptoAccount() CryptoAccount
	GetLoadedAddresses() (AddressIterator, error)
	Sign(txn Transaction, source string, pwd PasswordReader, index []int) (Transaction, error)
	AttachSignService(TxnSignStrategy) error
	RemoveSignService(TxnSignStrategy) error
	EnumerateSignServices() TxnSignStrategyIterator
}

type SeedGenerator interface {
	GenerateMnemonic(wordCount int) (string, error)
	VerifyMnemonic(seed string) (bool, error)
}

type WalletEnv interface {
	GetStorage() WalletStorage
	GetWalletSet() WalletSet
}
