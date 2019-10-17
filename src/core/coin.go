package core

// Timestamp number of
type Timestamp uint64

// TransactionStatus enumerates possible states of transaction lifecycle
type TransactionStatus uint32

const (
	// TXN_STATUS_CREATED Transaction created and not broadcasted
	TXN_STATUS_CREATED TransactionStatus = iota
	// TXN_STATUS_PENDING Transaction broadcasted but pending for confirmation
	TXN_STATUS_PENDING
	// TXN_STATUS_CONFIRMED Transaction confirmed
	TXN_STATUS_CONFIRMED
)

// Transaction encapsulates the contract for atomic transfers of coins
type Transaction interface {
	// Crypto assets involved in or supported by this transaction
	SupportedAssets() []string
	// GetTimestamp at the moment of creation
	GetTimestamp() Timestamp
	// GetStatus to retrieve transaction status
	GetStatus() TransactionStatus
	// GetInputs to list transaction inputs for spent transactions
	GetInputs() []TransactionInput
	// GetOutputs to list transaction outputs for coins distributed to participants
	GetOutputs() []TransactionOutput
	// GetId o retrieve transaction ID
	GetId() string
	// ComputeFee calculates transaction fee expressed in coins of asset represented by ticker
	ComputeFee(ticker string) (uint64, error)
	// VerifyUnsigned checks for valid unsigned transaction
	VerifyUnsigned() error
	// VerifySigned checks for valid unsigned transaction
	VerifySigned() error
	// IsFullySigned deermine whether all transaction elements have been signed
	IsFullySigned() (bool, error)
}

// TransactionIterator iterates over a sequence of transactions
type TransactionIterator interface {
	// Value of transaction at iterator pointer position
	Value() Transaction
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
}

// TransactionInput provides cryptographic proof of spending funds
type TransactionInput interface {
	// GetId provides transaction input ID
	GetId() string
	// GetSpentOutput looks up the output spent by this input
	GetSpentOutput() TransactionOutput
	// GetCoins looks up coins for asset represented by ticker
	// that have been spent by this input
	GetCoins(ticker string) (uint64, error)
}

// TransactionInputIterator iterates over a sequence of transaction inputs
type TransactionInputIterator interface {
	// Value of transaction input at iterator pointer position
	Value() TransactionInput
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
}

// TransactionOutput provides cryptographic proof of funds transfer
type TransactionOutput interface {
	// GetId provides transaction output ID
	GetId() string
	// IsSpent determines whether there exists a confirmed transaction with an input spending this output
	IsSpent() bool
	// GetAddress returns the address of the party receiving funds
	GetAddress() Address
	// GetCoins looks up coins for asset represented by ticker
	// that have been transferred in this output
	GetCoins(ticker string) (uint64, error)
}

// TransactionOutputIterator iterates over a sequence of transaction outputs
type TransactionOutputIterator interface {
	// Value of transaction output at iterator pointer position
	Value() TransactionOutput
	// Next discards current value and moves iteration pointer up to next item
	Next() bool
	// HasNext may be used to query whether more items are to be expected in the sequence
	HasNext() bool
}

// Block included in the blockchain
type Block interface {
	// GetHash returns block ID
	GetHash() ([]byte, error)
	// GetPrevHash retrieve previous blockID
	GetPrevHash() ([]byte, error)
	// GetVersion provides block version
	GetVersion() (uint32, error)
	// GetTime provides block creation timestamp
	GetTime() (Timestamp, error)
	// GetHeight computes the number of blocks preceding this block in the blockchain
	GetHeight() (uint64, error)
	// GetFee computes block fee expressed in coins of asset identified by ticker
	GetFee(ticker string) (uint64, error)
	// IsGenesisBlock determines whether this block starts blockchain sequence
	IsGenesisBlock() (bool, error)
}
