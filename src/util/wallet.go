package util

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

// SimpleWalletOutput put together transacion output with riginating wallet
type SimpleWalletOutput struct {
	Wallet core.Wallet
	UxOut  core.TransactionOutput
}

// GetWallet return wallet
func (wo *SimpleWalletOutput) GetWallet() core.Wallet {
	return wo.Wallet
}

// GetOutput return transaction output.
func (wo *SimpleWalletOutput) GetOutput() core.TransactionOutput {
	return wo.UxOut
}

// Type assertions
var (
	_ core.WalletOutput = &SimpleWalletOutput{}
)
