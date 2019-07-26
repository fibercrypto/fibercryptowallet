package core

type WalletIterator interface {
	// TODO: Define
}

type WalletSet interface {
	ListWallets() WalletIterator
	CreateWallet(name string, seed string, isEncryptrd bool, pwd PasswordReader, scanAddressesN int)
}

type WalletStorage interface {
	Encrypt(walletName, password string)
	Decrypt(walletName, password string)
	IsEncrypted(walletName string) bool
}

type Wallet interface {
	ScanAddresses(startIndex, count int) AddressIterator
	GetName() string
	SetName(wltName string)
	Transfer(to Address, amount uint64)
	SendFromAddress(from, to Address, amount uint64)
	Spend(unspent, new []TransactionOutput)
	GenChangeAddress() Address
}
