package util

import (
	"testing"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/mocks"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/stretchr/testify/require"
)

func TestAddLookupSigner(t *testing.T) {
	signerID1 := core.UID("TestAddLookupSigner#1")
	signer1 := new(mocks.TxnSigner)
	signer1.On("GetSignerUID").Return(signerID1)
	require.Equal(t, signer1.GetSignerUID(), signerID1)
	err := AttachSignService(signer1)
	require.NoError(t, err)

	signer := LookupSignService(signerID1)
	require.NotNil(t, signer)
	require.Equal(t, signer, signer1)
}
