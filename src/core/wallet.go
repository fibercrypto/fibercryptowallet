package core

type WalletIterator interface {
	Value() Wallet
	Next() bool
	HasNext() bool
}

type WalletSet interface {
	ListWallets() WalletIterator
	CreateWallet(name string, seed string, isEncryptrd bool, pwd PasswordReader, scanAddressesN int) (Wallet, error)
}

type WalletStorage interface {
	Encrypt(walletName, password PasswordReader)
	Decrypt(walletName, password PasswordReader)
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
	Transfer(to Address, amount uint64)
	SendFromAddress(from, to Address, amount uint64)
	Spend(unspent, new []TransactionOutput)
	GenAddresses(addrType AddressType, startIndex, count uint32) AddressIterator
	GetCryptoAccount() CryptoAccount
}

type SeedGenerator interface {
	GenerateMnemonic(wordCount int) (string, error)
	VerifyMnemonic(seed string) (bool, error)
}

type WalletEnv interface {
	GetStorage() WalletStorage
	GetWalletSet() WalletSet
}
