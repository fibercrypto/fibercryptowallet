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
	Encrypt(walletName, password string)
	Decrypt(walletName, password string)
	IsEncrypted(walletName string) bool
}

type Wallet interface {
	ScanAddresses(startIndex, count int) AddressIterator
	GetId() string
	GetLabel() string
	SetLabel(wltName string)
	Transfer(to Address, amount uint64)
	SendFromAddress(from, to Address, amount uint64)
	Spend(unspent, new []TransactionOutput)
	GenChangeAddress() Address
}

type SeedGenerator interface {
	GenerateMnemonic(wordCount int) (string, error)
	VerifyMnemonic(seed string) error
}

type WalletEnv interface {
	GetStorage() WalletStorage
	GetWalletSet() WalletSet
}
