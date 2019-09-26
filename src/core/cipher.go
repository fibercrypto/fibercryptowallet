package core

type Address interface {
	IsBip32() bool
	String() string
	GetCryptoAccount() CryptoAccount
}

type AddressIterator interface {
	Value() Address
	Next() bool
	HasNext() bool
}

// TxnSigner defines the contract enforced upon objects able to sin transacions.
type TxnSigner interface {
	// SignTransaction partially or in full
	SignTransaction(Transaction, PasswordReader, []string) (Transaction, error)
	// GetSignerUID provides the key identifying this signer among peer strategies supported by an object
	GetSignerUID() UID
	// GetSignerDescription facilitates a human readable caption identifying this signing strategy
	GetSignerDescription() string
}

type TxnSignerIterator interface {
	Value() TxnSigner
	Next() bool
	HasNext() bool
	Count() int
}
