package skycoin

import (
	"io/ioutil"
	"math"
	"strings"
	"testing"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/testsuite"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/stretchr/testify/require"
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

func TestUninjectedTransactionVerifySigned(t *testing.T) {
	// Mismatch header hash
	txn, err := makeTransaction(t)
	require.NoError(t, err)
	txn.InnerHash = cipher.SHA256{}
	uiTxn := makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "InnerHash does not match computed hash")

	// No inputs
	txn, err = makeTransaction(t)
	require.NoError(t, err)
	txn.In = make([]cipher.SHA256, 0)
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "No inputs")

	// No outputs
	txn, err = makeTransaction(t)
	require.NoError(t, err)
	txn.Out = make([]coin.TransactionOutput, 0)
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "No outputs")

	// Invalid number of sigs
	txn, err = makeTransaction(t)
	require.NoError(t, err)
	txn.Sigs = make([]cipher.Sig, 0)
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Invalid number of signatures")
	txn.Sigs = make([]cipher.Sig, 20)
	err = txn.UpdateHeader()
	require.NoError(t, err)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Invalid number of signatures")

	// Too many sigs & inputs
	txn, err = makeTransaction(t)
	require.NoError(t, err)
	txn.Sigs = make([]cipher.Sig, math.MaxUint16+1)
	txn.In = make([]cipher.SHA256, math.MaxUint16+1)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Too many signatures and inputs")

	// Duplicate inputs
	ux, kd, err1 := makeUxOutWithSecret(t)
	require.NoError(t, err1)
	txn = makeTransactionFromUxOut(t, ux, kd.SecKey)
	err = txn.PushInput(txn.In[0])
	require.NoError(t, err)
	txn.Sigs = nil
	txn.SignInputs([]cipher.SecKey{kd.SecKey, kd.SecKey})
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Duplicate spend")

	// Duplicate outputs
	txn, err = makeTransaction(t)
	to := txn.Out[0]
	err = txn.PushOutput(to.Address, to.Coins, to.Hours)
	require.NoError(t, err)
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Duplicate output in transaction")

	// Invalid signature, empty
	txn, err = makeTransaction(t)
	txn.Sigs[0] = cipher.Sig{}
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Unsigned input in transaction")

	// Invalid signature, not empty
	// A stable invalid signature must be used because random signatures could appear valid
	// Note: Transaction.Verify() only checks that the signature is a minimally valid signature
	badSig := "9a0f86874a4d9541f58a1de4db1c1b58765a868dc6f027445d0a2a8a7bddd1c45ea559fcd7bef45e1b76ccdaf8e50bbebd952acbbea87d1cb3f7a964bc89bf1ed5"
	txn, err = makeTransaction(t)
	txn.Sigs[0] = cipher.MustSigFromHex(badSig)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Failed to recover pubkey from signature")

	// We can't check here for other invalid signatures:
	//      - Signatures signed by someone else, spending coins they don't own
	//      - Signatures signing a different message
	// This must be done by blockchain tests, because we need the address
	// from the unspent being spent
	// The verification here only checks that the signature is valid at all

	// Output coins are 0
	txn, err = makeTransaction(t)
	txn.Out[0].Coins = 0
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Zero coin output")

	// Output coin overflow
	txn, err = makeTransaction(t)
	txn.Out[0].Coins = math.MaxUint64 - 3e6
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Output coins overflow")

	// Output coins are not multiples of 1e6 (valid, decimal restriction is not enforced here)
	txn, err = makeTransaction(t)
	txn.Out[0].Coins += 10
	err = txn.UpdateHeader()
	require.NoError(t, err)
	txn.Sigs = nil
	txn.SignInputs([]cipher.SecKey{genSecret})
	require.NotEqual(t, txn.Out[0].Coins%1e6, uint64(0))
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	require.NoError(t, uiTxn.VerifySigned())

	// Valid
	txn, err = makeTransaction(t)
	txn.Out[0].Coins = 10e6
	txn.Out[1].Coins = 1e6
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	require.NoError(t, uiTxn.VerifySigned())
}

func TestUninjectedTransactionVerifyUnsigned(t *testing.T) {
	txn, _, err := makeTransactionMultipleInputs(t, 2)
	require.NoError(t, err)
	uiTxn := makeUninjectedTransaction(t, &txn, 0)
	err = uiTxn.VerifyUnsigned()
	testutil.RequireError(t, err, "Unsigned transaction must contain a null signature")

	// Invalid signature, not empty
	// A stable invalid signature must be used because random signatures could appear valid
	// Note: Transaction.Verify() only checks that the signature is a minimally valid signature
	badSig := "9a0f86874a4d9541f58a1de4db1c1b58765a868dc6f027445d0a2a8a7bddd1c45ea559fcd7bef45e1b76ccdaf8e50bbebd952acbbea87d1cb3f7a964bc89bf1ed5"
	txn, _, err = makeTransactionMultipleInputs(t, 2)
	require.NoError(t, err)
	txn.Sigs[0] = cipher.Sig{}
	txn.Sigs[1] = cipher.MustSigFromHex(badSig)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifyUnsigned(), "Failed to recover pubkey from signature")

	txn.Sigs = nil
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	err = uiTxn.VerifyUnsigned()
	testutil.RequireError(t, err, "Invalid number of signatures")

	// Transaction is unsigned if at least 1 signature is null
	txn, _, err = makeTransactionMultipleInputs(t, 3)
	require.NoError(t, err)
	require.True(t, len(txn.Sigs) > 1)
	txn.Sigs[0] = cipher.Sig{}
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	err = uiTxn.VerifyUnsigned()
	require.NoError(t, err)

	// Transaction is unsigned if all signatures are null
	for i := range txn.Sigs {
		txn.Sigs[i] = cipher.Sig{}
	}
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	err = uiTxn.VerifyUnsigned()
	require.NoError(t, err)
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

func makeLocalWalletsFromKeyData(t *testing.T, keysData []KeyData) ([]core.Wallet, error) {
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
				w, err = walletSet.CreateWallet(walletID, kd.Mnemonic, false, func(string) (string, error) { return "", nil }, 1)
				require.NoError(t, err)
			}
			walletsCache[kd.Mnemonic] = w
		}
		wallets[i] = w
	}
	return wallets, nil
}

func TestTransactionSignInput(t *testing.T) {
	txn, keysData, err := makeTransactionMultipleInputs(t, 3)
	require.NoError(t, err)
	uiTxn := makeUninjectedTransaction(t, &txn, 0)
	var signedCoreTxn core.Transaction
	var isFullySigned bool
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)

	// Input is already signed
	wallets, err1 := makeLocalWalletsFromKeyData(t, keysData)
	require.NoError(t, err1)
	signedCoreTxn, err = wallets[0].Sign(uiTxn, SignerIDLocalWallet, nil, []string{"0"})
	testutil.RequireError(t, err, "Input already signed")
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)

	// Input is not signed
	txn.Sigs[1] = cipher.Sig{}
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	signedCoreTxn, err = wallets[1].Sign(uiTxn, SignerIDLocalWallet, nil, []string{"1"})
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
	signedCoreTxn, err = wallets[1].Sign(uiTxn, SignerIDLocalWallet, nil, []string{"1"})
	testutil.RequireError(t, err, "Input already signed")

	// Transaction has no sigs; sigs array is initialized
	txn.Sigs = nil
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	signedCoreTxn, err = wallets[2].Sign(uiTxn, SignerIDLocalWallet, nil, []string{"2"})
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

	/*
		// SignInputs on a partially signed transaction fails
		require.Panics(t, func() {
			txn.SignInputs(seckeys)
		})

		// Signing the rest of the inputs individually works
		err = txn.SignInput(seckeys[1], 1)
		require.NoError(t, err)
		require.False(t, txn.IsFullySigned())
		err = txn.SignInput(seckeys[0], 0)
		require.NoError(t, err)
		require.True(t, txn.IsFullySigned())

		// Can use SignInputs on allocated array of empty sigs
		txn.Sigs = make([]cipher.Sig, 3)
		txn.SignInputs(seckeys)
		require.True(t, txn.IsFullySigned())
	*/
}

/*
func TestTransactionSignInputs(t *testing.T) {
	txn := &Transaction{}
	// Panics if txns already signed
	txn.Sigs = append(txn.Sigs, cipher.Sig{})
	require.Panics(t, func() { txn.SignInputs([]cipher.SecKey{}) })
	// Panics if not enough keys
	txn = &Transaction{}
	ux, s := makeUxOutWithSecret(t)
	err := txn.PushInput(ux.Hash())
	require.NoError(t, err)
	ux2, s2 := makeUxOutWithSecret(t)
	err = txn.PushInput(ux2.Hash())
	require.NoError(t, err)
	err = txn.PushOutput(makeAddress(), 40, 80)
	require.NoError(t, err)
	require.Equal(t, len(txn.Sigs), 0)
	require.Panics(t, func() { txn.SignInputs([]cipher.SecKey{s}) })
	require.Equal(t, len(txn.Sigs), 0)
	// Valid signing
	h := txn.HashInner()
	require.NotPanics(t, func() { txn.SignInputs([]cipher.SecKey{s, s2}) })
	require.Equal(t, len(txn.Sigs), 2)
	h2 := txn.HashInner()
	require.Equal(t, h2, h)
	p := cipher.MustPubKeyFromSecKey(s)
	a := cipher.AddressFromPubKey(p)
	p = cipher.MustPubKeyFromSecKey(s2)
	a2 := cipher.AddressFromPubKey(p)
	require.NoError(t, cipher.VerifyAddressSignedHash(a, txn.Sigs[0], cipher.AddSHA256(h, txn.In[0])))
	require.NoError(t, cipher.VerifyAddressSignedHash(a2, txn.Sigs[1], cipher.AddSHA256(h, txn.In[1])))
	require.Error(t, cipher.VerifyAddressSignedHash(a, txn.Sigs[1], h))
	require.Error(t, cipher.VerifyAddressSignedHash(a2, txn.Sigs[0], h))
}
*/
