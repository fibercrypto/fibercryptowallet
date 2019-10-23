package skycoin

import (
	"path/filepath"
	"testing"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/testsuite"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	skytestsuite "github.com/skycoin/skycoin/src/cipher/testsuite"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/skycoin/skycoin/src/util/file"
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

var (
	seedPairIndex    = 0
	seedContinuation []byte
	seedMnemonic     string
	seedEntropy      []byte
	seedData         *skytestsuite.SeedTestData
)

type KeyData struct {
	SecKey       cipher.SecKey
	PubKey       cipher.PubKey
	Mnemonic     string
	Entropy      []byte
	AddressIndex int
}

// generateTestKeyPair provides deterministic sequence of test keys
// that can be recovered later inside a wallet
func generateTestKeyPair(t *testing.T) (*KeyData, error) {
	var err error
	if seedEntropy == nil {
		// Load suite test data
		fn := filepath.Join(testsuite.GetSkycoinCipherTestDataDir(), testsuite.ManyAddressesFilename)

		var dataJSON skytestsuite.SeedTestDataJSON
		err := file.LoadJSON(fn, &dataJSON)
		require.NoError(t, err)

		data, err := skytestsuite.SeedTestDataFromJSON(&dataJSON)
		require.NoError(t, err)

		// Initialize internal test state
		seedEntropy, err = bip39.EntropyFromMnemonic(string(data.Seed))
		require.NoError(t, err)
		seedContinuation = seedEntropy
		seedMnemonic = string(data.Seed)
		seedData = data
	}

	var keytestData KeyData
	seedContinuation, keytestData.PubKey, keytestData.SecKey, err = cipher.DeterministicKeyPairIterator(seedContinuation)
	if err != nil {
		return nil, err
	}
	// TODO: Verify wallet crypto pair
	keytestData.Mnemonic = seedMnemonic
	keytestData.Entropy = seedEntropy
	keytestData.AddressIndex = seedPairIndex
	seedPairIndex++
	return &keytestData, nil
}

func makeUxBodyWithSecret(t *testing.T) (coin.UxBody, *KeyData, error) {
	keydata, err := generateTestKeyPair(t)
	if err != nil {
		return coin.UxBody{}, nil, err
	}
	return coin.UxBody{
		SrcTransaction: testutil.RandSHA256(t),
		Address:        cipher.AddressFromPubKey(keydata.PubKey),
		Coins:          1e6,
		Hours:          100,
	}, keydata, nil
}

func makeUxOutWithSecret(t *testing.T) (coin.UxOut, *KeyData, error) {
	body, kd, err := makeUxBodyWithSecret(t)
	if err != nil {
		return coin.UxOut{}, nil, err
	}
	return coin.UxOut{
		Head: coin.UxHead{
			Time:  100,
			BkSeq: 2,
		},
		Body: body,
	}, kd, nil
}

func makeTransaction(t *testing.T) (coin.Transaction, error) {
	ux, kd, err := makeUxOutWithSecret(t)
	if err != nil {
		return coin.Transaction{}, err
	}
	return makeTransactionFromUxOut(t, ux, kd.SecKey), nil
}

func makeTransactionMultipleInputs(t *testing.T, n int) (coin.Transaction, []KeyData, []coin.UxOut, error) {
	uxs := make([]coin.UxOut, n)
	keysdata := make([]KeyData, n)
	secs := make([]cipher.SecKey, n)
	for i := 0; i < n; i++ {
		ux, kd, err := makeUxOutWithSecret(t)
		if err != nil {
			return coin.Transaction{}, nil, nil, err
		}
		uxs[i] = ux
		secs[i] = kd.SecKey
		keysdata[i] = *kd
	}
	return makeTransactionFromUxOuts(t, uxs, secs), keysdata, uxs, nil
}

func makeTransactions(t *testing.T, n int) (coin.Transactions, error) { //nolint:unparam
	txns := make(coin.Transactions, n)
	for i := range txns {
		var err error
		txns[i], err = makeTransaction(t)
		if err != nil {
			return nil, err
		}
	}
	return txns, nil
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

func makeSpentOutput(uxout coin.UxOut, spentBkSeq uint64, spentTxId cipher.SHA256) (rOut readable.SpentOutput) {
	rOut.Uxid = uxout.Hash().Hex()
	rOut.Time = uxout.Head.Time
	rOut.SrcBkSeq = uxout.Head.BkSeq
	rOut.SrcTx = uxout.Body.SrcTransaction.Hex()
	rOut.OwnerAddress = uxout.Body.Address.String()
	rOut.Coins = uxout.Body.Coins
	rOut.Hours = uxout.Body.Hours
	rOut.SpentBlockSeq = spentBkSeq
	rOut.SpentTxnID = spentTxId.Hex()
	return
}
