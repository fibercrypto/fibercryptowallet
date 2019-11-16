package skycoin

import (
	"path/filepath"
	"testing"

	"github.com/skycoin/skycoin/src/cipher/bip39"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/testsuite"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/cipher"
	skytestsuite "github.com/skycoin/skycoin/src/cipher/testsuite"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/skycoin/skycoin/src/util/file"
	"github.com/stretchr/testify/require"
)

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
	_, genSecret     = cipher.GenerateKeyPair()
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

func generateRandomKeyData(t *testing.T) (*KeyData, error) {
	entropy, err := bip39.NewEntropy(128)
	require.NoError(t, err)
	mnemonic, err := bip39.NewMnemonic(entropy)
	require.NoError(t, err)
	seed, err := bip39.NewSeed(mnemonic, "")
	require.NoError(t, err)
	pubKey, secKey, err := cipher.GenerateDeterministicKeyPair(seed)
	require.NoError(t, err)

	kd := &KeyData{
		AddressIndex: 0,
		Entropy:      entropy,
		Mnemonic:     mnemonic,
		PubKey:       pubKey,
		SecKey:       secKey,
	}

	return kd, nil
}

func generateTestKeyPair(t *testing.T) (*KeyData, error) {
	var err error
	if seedEntropy == nil {
		// Load suite test data
		fn := filepath.Join(testsuite.GetSkycoinCipherTestDataDir(), "seed-0000.golden") //testsuite.ManyAddressesFilename)

		var dataJSON skytestsuite.SeedTestDataJSON
		err := file.LoadJSON(fn, &dataJSON)
		require.NoError(t, err)

		data, err := skytestsuite.SeedTestDataFromJSON(&dataJSON)
		require.NoError(t, err)

		// Initialize internal test state
		seedEntropy = []byte(data.Seed)
		seedContinuation = seedEntropy
		seedMnemonic = string(data.Seed)
		seedData = data
		seedPairIndex = 0
	}

	var keytestData KeyData
	seedContinuation, keytestData.PubKey, keytestData.SecKey, err = cipher.DeterministicKeyPairIterator(seedContinuation)
	if err != nil {
		return nil, err
	}
	keytestData.Mnemonic = seedMnemonic
	keytestData.Entropy = seedEntropy
	keytestData.AddressIndex = seedPairIndex
	seedPairIndex++
	if keytestData.AddressIndex < len(seedData.Keys) {
		// Confirm that deterministic address sequence is correct
		require.Equal(t, seedData.Keys[keytestData.AddressIndex].Public, keytestData.PubKey)
		require.Equal(t, seedData.Keys[keytestData.AddressIndex].Secret, keytestData.SecKey)
	}
	return &keytestData, nil
}

func makeUxBodyWithRandomSecret(t *testing.T) (coin.UxBody, *KeyData, error) {
	keydata, err := generateRandomKeyData(t)
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

func makeUxWithRandomSecret(t *testing.T) (coin.UxOut, *KeyData, error) {
	body, kd, err := makeUxBodyWithRandomSecret(t)
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

func makeTransactionFromMultipleWallets(t *testing.T, n int) (coin.Transaction, []KeyData, []coin.UxOut, error) {
	uxs := make([]coin.UxOut, n)
	keysdata := make([]KeyData, n)
	secs := make([]cipher.SecKey, n)

	for i := 0; i < n; i++ {
		ux, kd, err := makeUxWithRandomSecret(t)
		require.NoError(t, err)
		uxs[i] = ux
		secs[i] = kd.SecKey
		keysdata[i] = *kd
	}

	return makeTransactionFromUxOuts(t, uxs, secs), keysdata, uxs, nil
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

func makeTransactions(t *testing.T, n int) (coin.Transactions, error) { //nolint:unparam,megacheck
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

func makeSimpleWalletAddress(wallet core.Wallet, address core.Address) core.WalletAddress {
	return &util.SimpleWalletAddress{
		Wallet: wallet,
		UxOut:  address,
	}
}

func makeSimpleWalletOutput(wallet core.Wallet, out core.TransactionOutput) core.WalletOutput {
	return &util.SimpleWalletOutput{
		Wallet: wallet,
		UxOut:  out,
	}
}
