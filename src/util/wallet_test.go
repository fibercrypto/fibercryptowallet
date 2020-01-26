package util

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
)

func TestSimpleWalletAddress(t *testing.T) {
	wllt, addr := &mocks.Wallet{}, &mocks.Address{}
	wa := &SimpleWalletAddress{Wallet: wllt, UxOut: addr}

	require.Equal(t, wllt, wa.GetWallet())
	require.Equal(t, addr, wa.GetAddress())
}

func TestSimpleWalletOutput(t *testing.T) {
	wllt, out := &mocks.Wallet{}, &mocks.TransactionOutput{}
	wa := &SimpleWalletOutput{Wallet: wllt, UxOut: out}

	require.Equal(t, wllt, wa.GetWallet())
	require.Equal(t, out, wa.GetOutput())
}
