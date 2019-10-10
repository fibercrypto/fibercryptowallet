package skycoin

import (
	"testing"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/stretchr/testify/require"
)

var genPublic, genSecret = cipher.GenerateKeyPair()

func makeAddress() cipher.Address {
	p, _ := cipher.GenerateKeyPair()
	return cipher.AddressFromPubKey(p)
}

func makeTransactionFromUxOuts(t *testing.T, uxs []coin.UxOut, secs []cipher.SecKey) coin.Transaction {
	require.Equal(t, len(uxs), len(secs))

	txn := coin.Transaction{}

	err := txn.PushOutput(makeAddress(), 1e6, 50)
	require.NoError(t, err)
	err = txn.PushOutput(makeAddress(), 5e6, 50)
	require.NoError(t, err)

	for _, ux := range uxs {
		err = txn.PushInput(ux.Hash())
		require.NoError(t, err)
	}

	txn.SignInputs(secs)

	err = txn.UpdateHeader()
	require.NoError(t, err)
	return txn
}

func makeTransactionFromUxOut(t *testing.T, ux coin.UxOut, s cipher.SecKey) coin.Transaction {
	return makeTransactionFromUxOuts(t, []coin.UxOut{ux}, []cipher.SecKey{s})
}

func makeUxBodyWithSecret(t *testing.T) (coin.UxBody, cipher.SecKey) {
	p, s := cipher.GenerateKeyPair()
	return coin.UxBody{
		SrcTransaction: testutil.RandSHA256(t),
		Address:        cipher.AddressFromPubKey(p),
		Coins:          1e6,
		Hours:          100,
	}, s
}

func makeUxOutWithSecret(t *testing.T) (coin.UxOut, cipher.SecKey) {
	body, sec := makeUxBodyWithSecret(t)
	return coin.UxOut{
		Head: coin.UxHead{
			Time:  100,
			BkSeq: 2,
		},
		Body: body,
	}, sec
}

func makeTransaction(t *testing.T) coin.Transaction {
	ux, s := makeUxOutWithSecret(t)
	return makeTransactionFromUxOut(t, ux, s)
}

func makeTransactionMultipleInputs(t *testing.T, n int) (coin.Transaction, []cipher.SecKey) {
	uxs := make([]coin.UxOut, n)
	secs := make([]cipher.SecKey, n)
	for i := 0; i < n; i++ {
		ux, s := makeUxOutWithSecret(t)
		uxs[i] = ux
		secs[i] = s
	}
	return makeTransactionFromUxOuts(t, uxs, secs), secs
}

func makeTransactions(t *testing.T, n int) coin.Transactions { //nolint:unparam
	txns := make(coin.Transactions, n)
	for i := range txns {
		txns[i] = makeTransaction(t)
	}
	return txns
}

func copyTransaction(txn coin.Transaction) coin.Transaction {
	txo := coin.Transaction{}
	txo.Length = txn.Length
	txo.Type = txn.Type
	txo.InnerHash = txn.InnerHash
	txo.Sigs = make([]cipher.Sig, len(txn.Sigs))
	copy(txo.Sigs, txn.Sigs)
	txo.In = make([]cipher.SHA256, len(txn.In))
	copy(txo.In, txn.In)
	txo.Out = make([]coin.TransactionOutput, len(txn.Out))
	copy(txo.Out, txn.Out)
	return txo
}
