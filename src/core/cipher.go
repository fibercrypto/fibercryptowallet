package core

// Address identifier for sending and receiving transactions
type Address interface {
	// IsBip32 flaqg shall be set if adddress generation complies to BIP 32
	IsBip32() bool
	// String return human-readable representation of this address
	String() string
	// GetCryptoAccount provides access to address transaction history
	GetCryptoAccount() CryptoAccount
}

// Iterate over addresses in a container
type AddressIterator interface {
	// Value of address at iterator pointer position
	Value() Address
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
}

// TxnSigner defines the contract enforced upon objects able to sin transacions.
type TxnSigner interface {
	// ReadyForTxn determines whether this signer instance can be used by wallet to sign given transaction
	ReadyForTxn(Wallet, Transaction) (bool, error)
	// SignTransaction partially or in full
	SignTransaction(Transaction, PasswordReader, []string) (Transaction, error)
	// GetSignerUID provides the key identifying this signer among peer strategies supported by an object
	GetSignerUID() UID
	// GetSignerDescription facilitates a human readable caption identifying this signing strategy
	GetSignerDescription() string
}

// TxnSignerIterator enumerates a set if TxSigner strategies
type TxnSignerIterator interface {
	// Value of signer at iterator pointer position
	Value() TxnSigner
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
	// Count total number of items in sequence
	Count() int
}
