package skytypes

import (
	"github.com/skycoin/skycoin/src/api"
)

// SkycoinTxn represents the common internal operations that should be applied upon Skycoin transaction wrapper types
type SkycoinTxn interface {
	// EncodeSkycoinTransaction serialize transaction data for subsequent broadcast through the peer-to-peer network
	EncodeSkycoinTransaction() ([]byte, error)
}

// ReadableTxn expreses the contract to use Skycoin readable transactions for signing
type ReadableTxn interface {
	// ToCreatedTransaction return an instance of api.CreatedTransaction equivalent to he current transaction object
	ToCreatedTransaction() (*api.CreatedTransaction, error)
}

// SkycoinWallet internal contract to be satisfied by Skycoin wallets
type SkycoinWallet interface {
	GetSkycoinWalletType() string
}
