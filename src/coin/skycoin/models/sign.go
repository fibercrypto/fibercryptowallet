package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
)

// SkycoinSignService implements BlockchainSignService for multi-wallet transaction signing
type SkycoinSignService struct{}

// Sign creates a new transaction by (fully or partially) signing a given transaction
func (sss *SkycoinSignService) Sign(txn core.Transaction, signSpec []core.InputSignDescriptor, pwd core.PasswordReader) (core.Transaction, error) {
	return util.GenericMultiWalletSign(txn, signSpec, pwd)
}

// Type assertions
var (
	_ core.BlockchainSignService = &SkycoinSignService{}
)
