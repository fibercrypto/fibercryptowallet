package core

// InputSignDescriptor specifies how to sign a specific transaction input
type InputSignDescriptor struct {
	// InputIndex absolute (e.g. UXID) input identifier or relative (e.g. array index) in transaction context
	InputIndex string
	// SignerID selects a given signing strategy. If empty, default strategy will be chosen
	SignerID UID
	// Wallet placeholder containing private keys to sign transaction input
	Wallet Wallet
}

// BlockchainSignService implement multi-wallet transaction signing for the blockchain
type BlockchainSignService interface {
	// Sign creates a new transaction by (fully or partially) signing a given transaction
	Sign(txn Transaction, signSpec []InputSignDescriptor, pwd PasswordReader) (Transaction, error)
}
