package skycoin

import (
	"crypto/rand"
	"github.com/davecgh/go-spew/spew"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/testsuite"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	skytestsuite "github.com/skycoin/skycoin/src/cipher/testsuite"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/skycoin/skycoin/src/util/file"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
	//"crypto/rand"
	//"testing"

	//"github.com/skycoin/skycoin/src/cipher"
	//"github.com/skycoin/skycoin/src/coin"
	//"github.com/skycoin/skycoin/src/readable"
	//"github.com/skycoin/skycoin/src/testutil"
	//"github.com/stretchr/testify/require"
)

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

func makeUxBodyWithSecret(t *testing.T) (coin.UxBody, *KeyData, error) {
	keydata, err := GenerateTestKeyPair(t)
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
				w, err = walletSet.CreateWallet(walletID, kd.Mnemonic, skywallet.WalletTypeDeterministic, false, func(string) (string, error) { return "", nil }, 0)
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

func randBytes(t *testing.T, n int) []byte {
	b := make([]byte, n)
	x, err := rand.Read(b)
	require.Equal(t, n, x)
	require.Nil(t, err)
	return b
}

// GenerateTestKeyPair provides deterministic sequence of test keys
//// that can be recovered later inside a wallet
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

func mockSkyApiUxOut(mock *SkycoinApiMock, ux coin.UxOut) {
	rUxOut := makeSpentOutput(ux, 0, cipher.SHA256{})
	mock.On("UxOut", ux.Hash().Hex()).Return(&rUxOut, nil)
}

func mockSkyApiTransactions(mock *SkycoinApiMock, addressesN []string) {
	mock.On("Transactions", []string{}).Return(nil, nil)
	mock.On("Transactions", []string{addressesN[0]}).Return(
		[]readable.TransactionWithStatus{},
		nil)
	mock.On("Transactions", []string{addressesN[1]}).Return(
		[]readable.TransactionWithStatus{
			readable.TransactionWithStatus{
				Status: readable.TransactionStatus{
					Confirmed: true,
				},
			},
		},
		nil)
	mock.On("Transactions", []string{addressesN[2]}).Return(
		[]readable.TransactionWithStatus{
			readable.TransactionWithStatus{
				Status: readable.TransactionStatus{
					Confirmed: true,
				},
			},
			readable.TransactionWithStatus{
				Status: readable.TransactionStatus{
					Confirmed: false,
				},
			},
		},
		nil)
}

func mockSkyApiCreateWallet(mock *SkycoinApiMock, wltOpt *api.CreateWalletOptions, label string, encrypted bool) {
	mock.On("CreateWallet", *wltOpt).Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     label,
				Encrypted: encrypted,
			},
		},
		nil)
}

func mockSkyApiWalletCreateTransaction(mock *SkycoinApiMock, wreq *api.WalletCreateTransactionRequest, crtTxn *api.CreateTransactionResponse) {
	mock.On("WalletCreateTransaction", *wreq).Return(
		crtTxn,
		nil)
}

func mockSkyApiCreateTransaction(mock *SkycoinApiMock, req *api.CreateTransactionRequest, crtTxn *api.CreateTransactionResponse) {
	mock.On("CreateTransaction", *req).Return(
		crtTxn,
		nil)
}

func mockSkyApiOutputsForAddresses(mock *SkycoinApiMock, addresses []string) {
	usOut := readable.UnspentOutput{
		Hash:            "hash1",
		Coins:           "42",
		Hours:           uint64(42),
		CalculatedHours: uint64(42),
		Address:         "2HPiZkMTD2pB9FZ6HbCxFSXa1FGeNkLeEbP",
	}
	response := &readable.UnspentOutputsSummary{
		HeadOutputs: readable.UnspentOutputs{usOut, usOut},
	}

	for _, addr := range addresses {
		mock.On(
			"OutputsForAddresses",
			[]string{
				addr,
			},
		).Return(response, nil)
	}
}

func mockSkyApiTransactionsVerbose(mock *SkycoinApiMock, addresses []string) {
	response := readable.TransactionWithStatusVerbose{
		Status: readable.TransactionStatus{
			Confirmed: false,
		},
	}
	response.Transaction.Hash = "hash1"

	for _, addr := range addresses {
		mock.On("TransactionsVerbose", []string{addr}).Return(
			[]readable.TransactionWithStatusVerbose{
				response,
			},
			nil,
		)
	}
}

var global_mock *SkycoinApiMock

var logModelTest = logging.MustGetLogger("Skycoin Model Test")
// CleanGlobalMock util when is needed to change the values of an

// API method used in other test with different values.
func CleanGlobalMock() {
	if global_mock == nil {
		global_mock = new(SkycoinApiMock)
	} else {
		global_mock.ExpectedCalls = []*mock.Call{}
	}
}

func TransactionSignInputTestImpl(t *testing.T, signer core.TxnSigner) {
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
	if signer != nil {
		err = util.AttachSignService(signer)
		require.NoError(t, err)
	}

	// Input is already signed
	_, err = wallets[0].Sign(uiTxn, signer, util.EmptyPassword, []string{"0"})
	testutil.RequireError(t, err, "Transaction is fully signed")
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)

	// Input is not signed
	txn.Sigs[1] = cipher.Sig{}
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	signedCoreTxn, err = wallets[1].Sign(uiTxn, signer, nil, []string{"1"})
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
	_, err = wallets[1].Sign(signedTxn, signer, nil, []string{"1"})
	// FIXME use a named var from errors
	testutil.RequireError(t, err, "Transaction is fully signed")

	// Transaction has no sigs; sigs array is initialized
	txn.Sigs = nil
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	spew.Dump("{{{{{uiTxn", uiTxn)
	signedCoreTxn, err = wallets[2].Sign(uiTxn, signer, nil, []string{"2"})
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
	signedCoreTxn, err = wallets[1].Sign(signedTxn, signer, nil, []string{"1"})
	require.NoError(t, err)
	signedTxn, isUninjected = signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	isFullySigned, err = signedTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
	signedCoreTxn, err = wallets[0].Sign(signedTxn, signer, nil, []string{"0"})
	require.NoError(t, err)
	signedTxn, isUninjected = signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	isFullySigned, err = signedTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)
}