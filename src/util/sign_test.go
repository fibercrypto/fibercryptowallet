package util

import (
	"fmt"
	"testing"

	"github.com/SkycoinProject/skycoin/src/testutil"
	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignerMethods(t *testing.T) {
	signerID1 := core.UID("TestAddLookupSigner#1")
	signer1 := new(mocks.TxnSigner)
	signer1.On("GetSignerUID").Return(signerID1, nil)
	signer1.On("GetSignerDescription").Return(string(signerID1), nil)
	uid, err := signer1.GetSignerUID()
	require.NoError(t, err)
	require.Equal(t, uid, signerID1)
	desc, err := signer1.GetSignerDescription()
	require.NoError(t, err)
	require.Equal(t, desc, string(signerID1))
	err = AttachSignService(signer1)
	require.NoError(t, err)
	defer RemoveSignService(signerID1) // nolint gosec

	signer := LookupSignService(signerID1)
	require.NotNil(t, signer)
	require.Equal(t, signer, signer1)
	desc, err = GetSignerDescription(signerID1)
	require.NoError(t, err)
	require.Equal(t, string(signerID1), desc)

	err = RemoveSignService(signerID1)
	require.NoError(t, err)
	signer = LookupSignService(signerID1)
	require.Nil(t, signer)
	_, err = GetSignerDescription(signerID1)
	testutil.RequireError(t, err, "Invalid Id")
}

func TestSignersEnum(t *testing.T) {
	signerIDs := []core.UID{
		core.UID("TestSignersEnum#0"),
		core.UID("TestSignersEnum#1"),
		core.UID("TestSignersEnum#2"),
	}
	signers := make([]*mocks.TxnSigner, len(signerIDs))
	// Install global signers
	for i, uid := range signerIDs {
		signer := new(mocks.TxnSigner)
		signer.On("GetSignerUID").Return(uid, nil)
		uid, err := signer.GetSignerUID()
		require.NoError(t, err)
		require.Equal(t, uid, uid)
		err = AttachSignService(signer)
		require.NoError(t, err)
		signers[i] = signer
	}
	defer RemoveSignService(signerIDs[0]) // nolint gosec
	defer RemoveSignService(signerIDs[1]) // nolint gosec
	defer RemoveSignService(signerIDs[2]) // nolint gosec

	allSigners := EnumerateSignServices()
	signersFound := make(map[core.UID]struct{})
	for allSigners.Next() {
		uid, err := allSigners.Value().GetSignerUID()
		require.NoError(t, err)
		signerID := uid
		_, wasFound := signersFound[signerID]
		require.False(t, wasFound)
		signersFound[signerID] = struct{}{}
	}
	_, wasFound := signersFound[signerIDs[0]]
	require.True(t, wasFound)
	_, wasFound = signersFound[signerIDs[1]]
	require.True(t, wasFound)
	_, wasFound = signersFound[signerIDs[2]]
	require.True(t, wasFound)
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
		signer.On("GetSignerUID").Return(uid, nil)
		dUid, err := signer.GetSignerUID()
		require.NoError(t, err)
		require.Equal(t, dUid, uid)
		err = AttachSignService(signer)
		require.NoError(t, err)
		signers[i] = signer
	}
	defer RemoveSignService(signerIDs[0]) // nolint gosec
	defer RemoveSignService(signerIDs[1]) // nolint gosec
	defer RemoveSignService(signerIDs[2]) // nolint gosec
	defer RemoveSignService(signerIDs[3]) // nolint gosec
	defer RemoveSignService(signerIDs[4]) // nolint gosec
	defer RemoveSignService(signerIDs[5]) // nolint gosec

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
	for supportedSigners.Next() {
		signerID, err := supportedSigners.Value().GetSignerUID()
		require.NoError(t, err)
		_, wasFound := signersFound[signerID]
		require.False(t, wasFound)
		signersFound[signerID] = struct{}{}
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

	ready, err := ReadyForTxn(signerIDs[0], wlt, txn)
	require.True(t, ready)
	require.Nil(t, err)
	ready, err = ReadyForTxn(signerIDs[1], wlt, txn)
	require.False(t, ready)
	require.Nil(t, err)
	ready, err = ReadyForTxn(signerIDs[2], wlt, txn)
	require.False(t, ready)
	require.NotNil(t, err)
	ready, err = ReadyForTxn(signerIDs[3], wlt, txn)
	require.True(t, ready)
	require.Nil(t, err)
	ready, err = ReadyForTxn(signerIDs[4], wlt, txn)
	require.True(t, ready)
	require.NotNil(t, err)
	ready, err = ReadyForTxn(signerIDs[5], wlt, txn)
	require.True(t, ready)
	require.Nil(t, err)
	_, err = ReadyForTxn(core.UID("unknown_id"), wlt, txn)
	require.NotNil(t, err)
}

func TestLookupSignServiceForWallet(t *testing.T) {
	type WalletSigner struct {
		mocks.Wallet
		mocks.TxnSigner
	}

	emptyUID := core.UID("")
	uid := core.UID("walletid")
	other := core.UID("otherid")

	ws := new(WalletSigner)
	ws.TxnSigner.On("GetSignerUID").Return(uid)
	var signer core.TxnSigner = ws
	err := AttachSignService(signer)
	defer RemoveSignService(uid) // nolint gosec
	require.NoError(t, err)

	tests := []struct {
		wallet core.Wallet
		id     core.UID
		want   core.TxnSigner
	}{
		{wallet: new(mocks.Wallet), id: emptyUID, want: nil},
		{wallet: ws, id: emptyUID, want: signer},
		{wallet: ws, id: uid, want: signer},
		{wallet: new(mocks.Wallet), id: uid, want: signer},
		{wallet: new(mocks.Wallet), id: other, want: nil},
		{wallet: ws, id: other, want: nil},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("lookup%d", i), func(t *testing.T) {
			_signer, err := LookupSignServiceForWallet(tt.wallet, tt.id)
			require.Equal(t, tt.want, _signer)
			if tt.want == nil {
				require.NotNil(t, err)
			}
		})
	}
}

func TestSignTransaction(t *testing.T) {
	pwd := func(s string, kvs core.KeyValueStore) (string, error) {
		return s, nil
	}

	uid := core.UID("signer_id")
	signer := new(mocks.TxnSigner)
	signer.On("GetSignerUID").Return(uid)
	attachErr := AttachSignService(signer)
	require.Nil(t, attachErr)
	defer RemoveSignService(uid) // nolint gosec

	ind := make([]string, 0)
	txn := new(mocks.Transaction)
	ctx := new(mocks.KeyValueStore)

	signer.On("SignTransaction", txn, mock.Anything, ind).Return(txn, nil).Run(func(args mock.Arguments) {
		pwdReader := args.Get(1).(core.PasswordReader)
		msg := "message"
		_msg, err2 := pwdReader(msg, ctx)
		require.Nil(t, err2)
		require.Equal(t, msg, _msg)
	})
	signedTxn, err := SignTransaction(uid, txn, pwd, ind)
	require.Nil(t, err)
	require.Equal(t, txn, signedTxn)

	_, err = SignTransaction(core.UID(""), txn, pwd, ind)
	require.NotNil(t, err)
}

func TestGenericMultiWalletSign(t *testing.T) {
	var pwd core.PasswordReader = func(s string, kvs core.KeyValueStore) (string, error){
		return s, nil
	}

	inputsIdx := []string{"index1", "index2"}
	wlt, txn := new(mocks.Wallet), new(mocks.Transaction)

	uid := core.UID("signer_id")
	signer := new(mocks.TxnSigner)
	signer.On("GetSignerUID").Return(uid)
	attachErr := AttachSignService(signer)
	require.Nil(t, attachErr)
	defer RemoveSignService(uid) // nolint gosec

	spec := make([]core.InputSignDescriptor, 0)
	for _, input := range inputsIdx {
		inputSpec := core.InputSignDescriptor{
			InputIndex: input,
			SignerID:   uid,
			Wallet:     wlt,
		}
		spec = append(spec, inputSpec)
	}

	badTxn := new(mocks.Transaction)
	badTxn.On("GetId").Return("bad_txn_id")
	wlt.On("Sign", txn, signer, mock.Anything, inputsIdx).Return(txn, nil)
	wlt.On("Sign", badTxn, signer, mock.Anything, inputsIdx).Return(txn, errors.ErrInvalidTxn)
	_txn, err := GenericMultiWalletSign(txn, spec, pwd)
	require.Equal(t, txn, _txn)
	require.Nil(t, err)

	_, err = GenericMultiWalletSign(badTxn, spec, pwd)
	require.NotNil(t, err)

	spec = []core.InputSignDescriptor{
		core.InputSignDescriptor{
			InputIndex: "",
			SignerID:   core.UID(""),
			Wallet:     wlt,
		},
	}
	_, err = GenericMultiWalletSign(txn, spec, pwd)
	require.NotNil(t, err)
}
