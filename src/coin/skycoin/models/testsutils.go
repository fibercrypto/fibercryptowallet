package skycoin

import (
	"github.com/SkycoinProject/skycoin/src/wallet"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/testsuite"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/SkycoinProject/skycoin/src/cipher"
	skytestsuite "github.com/SkycoinProject/skycoin/src/cipher/testsuite"
	"github.com/SkycoinProject/skycoin/src/coin"
	"github.com/SkycoinProject/skycoin/src/testutil"
	"github.com/SkycoinProject/skycoin/src/util/file"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func makeUninjectedTransaction(t *testing.T, txn *coin.Transaction, fee uint64) *SkycoinUninjectedTransaction {
	if txn == nil {
		skyTxn, err := makeTransaction(t)
		require.NoError(t, err)
		txn = &skyTxn
	}
	uitxn, err := NewUninjectedTransaction(txn, fee)
	require.NoError(t, err)
	return uitxn
}

func makeLocalWallet(t *testing.T) core.Wallet {
	_, kd, err := makeUxOutWithSecret(t)
	require.NoError(t, err)
	require.NoError(t, err)
	wallets := makeLocalWalletsFromKeyData(t, []KeyData{*kd})
	wallet := wallets[0]
	seed, seckeys, err2 := cipher.GenerateDeterministicKeyPairsSeed([]byte(kd.Mnemonic), kd.AddressIndex+1)
	require.NoError(t, err2)
	require.Equal(t, kd.SecKey, seckeys[kd.AddressIndex])
	_, _, err3 := cipher.GenerateDeterministicKeyPair(seed)
	require.NoError(t, err3)
	wallet.GenAddresses(core.AccountAddress, uint32(kd.AddressIndex), 2, nil)
	return wallet
}

var (
	tmpWalletDir  string
	testWalletEnv core.WalletEnv
)

func loadTestWalletEnv(t *testing.T) core.WalletEnv {
	if tmpWalletDir == "" {
		tmpDir, err := ioutil.TempDir("", testsuite.TestIDToken)
		require.NoError(t, err)
		tmpWalletDir = tmpDir
	}
	if testWalletEnv == nil {
		testWalletEnv = NewWalletDirectory(tmpWalletDir)
	}
	return testWalletEnv
}

var whitespaceReplacer = strings.NewReplacer(" ", "-")

func makeLocalWalletsFromKeyData(t *testing.T, keysData []KeyData) []core.Wallet {
	walletsCache := make(map[string]core.Wallet)
	wallets := make([]core.Wallet, len(keysData))
	walletSet := loadTestWalletEnv(t).GetWalletSet()
	for i, kd := range keysData {
		walletID := whitespaceReplacer.Replace(kd.Mnemonic)
		var w core.Wallet
		var isFound bool
		var err error
		if w, isFound = walletsCache[kd.Mnemonic]; !isFound {
			if w = walletSet.GetWallet(walletID); w == nil {
				w, err = walletSet.CreateWallet(walletID, kd.Mnemonic, wallet.WalletTypeDeterministic, false, util.EmptyPassword, 0)
				require.NoError(t, err)
			}
			walletsCache[kd.Mnemonic] = w
		}
		wallets[i] = w
		w.GenAddresses(core.AccountAddress, 0, uint32(kd.AddressIndex+1), nil)
		w.GenAddresses(core.ChangeAddress, 0, uint32(kd.AddressIndex+1), nil)
	}
	return wallets
}

type TransferOptions struct {
	values map[string]interface{}
}

func (tOpt *TransferOptions) GetValue(key string) interface{} {
	return tOpt.values[key]
}

func (tOpt *TransferOptions) SetValue(key string, value interface{}) {
	tOpt.values[key] = value
}

// GenerateTestKeyPair provides deterministic sequence of test keys
// that can be recovered later inside a wallet
func GenerateTestKeyPair(t *testing.T) (*KeyData, error) {
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

var global_mock *SkycoinApiMock

var logModelTest = logging.MustGetLogger("Skycoin Model Test")
// API method used in other test with different values.

// CleanGlobalMock util when is needed to change the values of an
func CleanGlobalMock() {
	if global_mock == nil {
		global_mock = new(SkycoinApiMock)
	} else {
		global_mock.ExpectedCalls = []*mock.Call{}
	}
}

func TransactionSignInputTestSkyHwImpl(t *testing.T, signers []core.TxnSigner, setWallet func(*testing.T, core.TxnSigner, core.Wallet)) {
	_, keysData, _, err := makeTransactionMultipleInputs(t, 3)
	require.NoError(t, err)
	// Load local wallets
	wallets := makeLocalWalletsFromKeyData(t, keysData)
	for idx := range signers {
		setWallet(t, signers[idx], wallets[idx])
	}
	TransactionSignInputTestImpl(t, signers)
}

func TransactionSignInputTestImpl(t *testing.T, signers []core.TxnSigner) {
	txn, keysData, uxspent, err := makeTransactionMultipleInputs(t, 3)
	require.NoError(t, err)
	// Mock UxOut API calls
	for _, ux := range uxspent {
		mockSkyApiUxOut(global_mock, ux)
	}

	uiTxn := makeUninjectedTransaction(t, &txn, 1)
	var signedCoreTxn core.Transaction
	var isFullySigned bool
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)

	// Load local wallets
	wallets := makeLocalWalletsFromKeyData(t, keysData)
	for _, signer := range signers {
		err = util.AttachSignService(signer)
		require.NoError(t, err)
	}

	// Input is already signed
	var signer core.TxnSigner
	if len(signers) > 0 {
		signer = signers[0]
	}
	_, err = wallets[0].Sign(uiTxn, signer, util.EmptyPassword, []string{"#0"})
	testutil.RequireError(t, err, "Transaction is fully signed")
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)

	// Input is not signed
	txn.Sigs[1] = cipher.Sig{}
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	if len(signers) > 1 {
		signer = signers[1]
	}
	signedCoreTxn, err = wallets[1].Sign(uiTxn, signer, nil, []string{"#1"})
	require.NoError(t, err)
	signedTxn, isUninjected := signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	require.NotEqual(t, uiTxn.txn, signedTxn.txn)
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	isFullySigned, err = signedTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)
	if len(signers) > 2 {
		signer = signers[2]
	}
	_, err = wallets[1].Sign(signedTxn, signer, nil, []string{"#1"})
	testutil.RequireError(t, err, "Transaction is fully signed")
	// Repeat using UXID
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	uxId := txn.In[1].Hex()
	if len(signers) > 1 {
		signer = signers[1]
	}
	signedCoreTxn, err = wallets[1].Sign(uiTxn, signer, nil, []string{uxId})
	require.NoError(t, err)
	signedTxn, isUninjected = signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	require.NotEqual(t, uiTxn.txn, signedTxn.txn)
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	isFullySigned, err = signedTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)
	if len(signers) > 1 {
		signer = signers[1]
	}
	_, err = wallets[1].Sign(signedTxn, signer, nil, []string{"#1"})
	testutil.RequireError(t, err, "Transaction is fully signed")

	// Transaction has no sigs; sigs array is initialized
	txn.Sigs = nil
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	if len(signers) > 2 {
		signer = signers[2]
	}
	signedCoreTxn, err = wallets[2].Sign(uiTxn, signer, nil, []string{"#2"})
	require.NoError(t, err)
	signedTxn, isUninjected = signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	require.NotEqual(t, uiTxn.txn, signedTxn.txn)
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	require.Len(t, signedTxn.txn.Sigs, 3)
	require.True(t, signedTxn.txn.Sigs[0].Null())
	require.True(t, signedTxn.txn.Sigs[1].Null())
	require.False(t, signedTxn.txn.Sigs[2].Null())

	// Signing the rest of the inputs individually works
	if len(signers) > 1 {
		signer = signers[1]
	}
	signedCoreTxn, err = wallets[1].Sign(signedTxn, signer, nil, []string{"#1"})
	require.NoError(t, err)
	signedTxn, isUninjected = signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	isFullySigned, err = signedTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	if len(signers) > 0 {
		signer = signers[0]
	}
	signedCoreTxn, err = wallets[0].Sign(signedTxn, signer, nil, []string{"#0"})
	require.NoError(t, err)
	signedTxn, isUninjected = signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	isFullySigned, err = signedTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)
}
