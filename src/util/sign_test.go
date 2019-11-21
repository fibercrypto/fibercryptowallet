package util

import (
	"testing"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/mocks"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/errors"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/stretchr/testify/require"
)

func TestSignerMethods(t *testing.T) {
	signerID1 := core.UID("TestAddLookupSigner#1")
	signer1 := new(mocks.TxnSigner)
	signer1.On("GetSignerUID").Return(signerID1)
	signer1.On("GetSignerDescription").Return(string(signerID1))
	require.Equal(t, signer1.GetSignerUID(), signerID1)
	require.Equal(t, signer1.GetSignerDescription(), string(signerID1))
	err := AttachSignService(signer1)
	require.NoError(t, err)
	defer RemoveSignService(signerID1)

	signer := LookupSignService(signerID1)
	require.NotNil(t, signer)
	require.Equal(t, signer, signer1)
	desc, err := GetSignerDescription(signerID1)
	require.NoError(t, err)
	require.Equal(t, string(signerID1), desc)

	err = RemoveSignService(signerID1)
	signer = LookupSignService(signerID1)
	require.Nil(t, signer)
	desc, err = GetSignerDescription(signerID1)
	testutil.RequireError(t, err, "Invalid Id")
}

func TestSignersReadyForTxn(t *testing.T) {
	signerIDs := []core.UID{
		core.UID("TestSignersReadyForTxn#0"),
		core.UID("TestSignersReadyForTxn#1"),
		core.UID("TestSignersReadyForTxn#2"),
		core.UID("TestSignersReadyForTxn#3"),
		core.UID("TestSignersReadyForTxn#4"),
		core.UID("TestSignersReadyForTxn#5"),
	}
	signers := make([]*mocks.TxnSigner, len(signerIDs))
	// Install global signers
	for i, uid := range signerIDs {
		signer := new(mocks.TxnSigner)
		signer.On("GetSignerUID").Return(uid)
		require.Equal(t, signer.GetSignerUID(), uid)
		err := AttachSignService(signer)
		require.NoError(t, err)
		signers[i] = signer
	}
	defer RemoveSignService(signerIDs[0])
	defer RemoveSignService(signerIDs[1])
	defer RemoveSignService(signerIDs[2])
	defer RemoveSignService(signerIDs[3])
	defer RemoveSignService(signerIDs[4])
	defer RemoveSignService(signerIDs[5])

	var txn core.Transaction = new(mocks.Transaction)
	var wlt core.Wallet = new(mocks.Wallet)
	signers[0].On("ReadyForTxn", wlt, txn).Return(true, nil)
	signers[1].On("ReadyForTxn", wlt, txn).Return(false, nil)
	signers[2].On("ReadyForTxn", wlt, txn).Return(false, errors.ErrInvalidID)
	signers[3].On("ReadyForTxn", wlt, txn).Return(true, nil)
	signers[4].On("ReadyForTxn", wlt, txn).Return(true, errors.ErrInvalidID)
	signers[5].On("ReadyForTxn", wlt, txn).Return(true, nil)

	supportedSigners := SignServicesForTxn(wlt, txn)
	signersFound := make(map[core.UID]struct{})
	for supportedSigners.HasNext() {
		signerID := supportedSigners.Value().GetSignerUID()
		_, wasFound := signersFound[signerID]
		require.False(t, wasFound)
		signersFound[signerID] = struct{}{}
		supportedSigners.Next()
	}
	_, wasFound := signersFound[signerIDs[0]]
	require.True(t, wasFound)
	_, wasFound = signersFound[signerIDs[1]]
	require.False(t, wasFound)
	_, wasFound = signersFound[signerIDs[2]]
	require.False(t, wasFound)
	_, wasFound = signersFound[signerIDs[3]]
	require.True(t, wasFound)
	_, wasFound = signersFound[signerIDs[4]]
	require.False(t, wasFound)
	_, wasFound = signersFound[signerIDs[5]]
	require.True(t, wasFound)
}
