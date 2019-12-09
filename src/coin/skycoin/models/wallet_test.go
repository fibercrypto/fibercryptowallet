package skycoin

import (
	mocks2 "github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"math"
	"os"
	"strconv"
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/skycoin/skycoin/src/wallet"
	"github.com/stretchr/testify/require"
)

func TestTransactionFinderAddressesActivity(t *testing.T) {
	CleanGlobalMock()

	addresses := make([]cipher.Address, 0)
	addressesN := make([]string, 0)

	for i := 0; i < 3; i++ {
		p, _ := cipher.GenerateKeyPair()
		a := cipher.AddressFromPubKey(p)
		s := a.String()
		addresses = append(addresses, a)
		addressesN = append(addressesN, s)
	}

	mockSkyApiTransactions(global_mock, addressesN)

	thxF := &TransactionFinder{}

	mask, err := thxF.AddressesActivity([]cipher.Address{})
	require.NoError(t, err)
	require.Equal(t, 0, len(mask))
	require.Equal(t, []bool{}, mask)

	mask, err = thxF.AddressesActivity(addresses)
	require.NoError(t, err)
	require.Equal(t, 3, len(mask))
	require.Equal(t, false, mask[0])
	for i := 1; i < 3; i++ {
		require.Equal(t, true, mask[i])
	}
}

func TestSkycoinRemoteWalletListWallets(t *testing.T) {

	global_mock.On("Wallets").Return(
		[]api.WalletResponse{
			api.WalletResponse{
				Meta: readable.WalletMeta{
					Coin:      "Sky",
					Filename:  "FiberCrypto",
					Label:     "wallet",
					Encrypted: true,
				},
			},
			api.WalletResponse{
				Meta: readable.WalletMeta{
					Coin:      "Sky",
					Filename:  "FiberCrypto",
					Label:     "wallet",
					Encrypted: true,
				},
			},
		},
		nil)

	wltSrv := &SkycoinRemoteWallet{poolSection: PoolSection}
	iter := wltSrv.ListWallets()
	for iter.Next() {
		wlt := iter.Value()
		require.Equal(t, "wallet", wlt.GetLabel())
		require.Equal(t, "FiberCrypto", wlt.GetId())
	}
}

func TestSkycoinRemoteWalletCreateWallet(t *testing.T) {

	seed, label, pwd, scanN := "seed", "label", "pwd", 666

	wltOpt1 := api.CreateWalletOptions{
		Type:     wallet.WalletTypeDeterministic,
		Seed:     seed,
		Label:    label,
		Password: pwd,
		ScanN:    scanN,
		Encrypt:  true,
	}
	wltOpt2 := api.CreateWalletOptions{
		Type:    wallet.WalletTypeDeterministic,
		Seed:    seed,
		Label:   label,
		ScanN:   scanN,
		Encrypt: false,
	}

	mockSkyApiCreateWallet(global_mock, &wltOpt1, "walletEncrypted", true)
	mockSkyApiCreateWallet(global_mock, &wltOpt2, "walletNonEncrypted", false)

	wltSrv := &SkycoinRemoteWallet{poolSection: PoolSection}
	pwdReader := func(message string) (string, error) {
		return "pwd", nil
	}

	wlt1, err := wltSrv.CreateWallet(label, seed, wallet.WalletTypeDeterministic, true, pwdReader, scanN)
	require.NoError(t, err)
	require.Equal(t, "walletEncrypted", wlt1.GetLabel())
	require.Equal(t, "FiberCrypto", wlt1.GetId())

	wlt2, err := wltSrv.CreateWallet(label, seed, wallet.WalletTypeDeterministic, false, pwdReader, scanN)
	require.NoError(t, err)
	require.Equal(t, "walletNonEncrypted", wlt2.GetLabel())
	require.Equal(t, "FiberCrypto", wlt2.GetId())
}

func TestSkycoinRemoteWalletEncrypt(t *testing.T) {
	CleanGlobalMock()
	global_mock.On("EncryptWallet", "wallet", "pwd").Return(&api.WalletResponse{}, nil)

	wltSrv := &SkycoinRemoteWallet{poolSection: PoolSection}
	pwdReader := func(message string) (string, error) {
		return "pwd", nil
	}

	wltSrv.Encrypt("wallet", pwdReader)
}

func TestSkycoinRemoteWalletDecrypt(t *testing.T) {
	CleanGlobalMock()
	global_mock.On("DecryptWallet", "wallet", "pwd").Return(&api.WalletResponse{}, nil)

	wltSrv := &SkycoinRemoteWallet{poolSection: PoolSection}
	pwdReader := func(message string) (string, error) {
		return "pwd", nil
	}

	wltSrv.Decrypt("wallet", pwdReader)
}

func TestSkycoinRemoteWalletIsEncrypted(t *testing.T) {

	global_mock.On("Wallet", "encrypted").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Encrypted: true,
			},
		},
		nil)
	global_mock.On("Wallet", "nonEncrypted").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Encrypted: false,
			},
		},
		nil)

	wltSrv := &SkycoinRemoteWallet{poolSection: PoolSection}

	encrypted, err := wltSrv.IsEncrypted("encrypted")
	require.NoError(t, err)
	require.Equal(t, true, encrypted)

	encrypted, err = wltSrv.IsEncrypted("nonEncrypted")
	require.NoError(t, err)
	require.Equal(t, false, encrypted)
}

func TestSkycoinRemoteWalletGetWallet(t *testing.T) {
	CleanGlobalMock()

	global_mock.On("Wallet", "wallet").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "wallet",
				Encrypted: true,
			},
			Entries: []readable.WalletEntry{
				readable.WalletEntry{Address: "addr"},
			},
		},
		nil)

	wltSrv := &SkycoinRemoteWallet{poolSection: PoolSection}
	wlt := wltSrv.GetWallet("wallet")
	require.Equal(t, "wallet", wlt.GetLabel())
	require.Equal(t, "FiberCrypto", wlt.GetId())
}

func TestRemoteWalletSignSkycoinTxn(t *testing.T) {
	hash := testutil.RandSHA256(t)
	txn := coin.Transaction{
		Length:    100,
		Type:      0,
		InnerHash: hash,
	}
	unTxn := SkycoinUninjectedTransaction{
		txn:     &txn,
		inputs:  nil,
		outputs: nil,
		fee:     100,
	}
	encodedResponse, err := unTxn.txn.SerializeHex()
	require.NoError(t, err)

	walletSignTxn := api.WalletSignTransactionRequest{
		EncodedTransaction: encodedResponse,
		WalletID:           "wallet",
		Password:           "password",
		SignIndexes:        nil,
	}

	crtTxn, err := api.NewCreateTransactionResponse(&txn, nil)
	crtTxn.Transaction.Fee = "100"
	require.NoError(t, err)

	global_mock.On("WalletSignTransaction", walletSignTxn).Return(
		crtTxn,
		nil)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	pwdReader := func(message string) (string, error) {
		return "password", nil
	}
	ret, err := wlt.Sign(&unTxn, nil, pwdReader, nil)
	require.NoError(t, err)
	require.NotNil(t, ret)
	value, err := ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(100), value)
}

func TestRemoteWalletSetLabel(t *testing.T) {
	CleanGlobalMock()
	global_mock.On("UpdateWallet", "walletId", "wallet").Return(nil)

	wlt := &RemoteWallet{
		Id:          "walletId",
		poolSection: PoolSection,
	}

	wlt.SetLabel("wallet")
}

func NewTransferOptions() *TransferOptions {
	tOptions := TransferOptions{
		values: make(map[string]interface{}),
	}
	return &tOptions
}

func TestRemoteWalletTransfer(t *testing.T) {
	CleanGlobalMock()
	destinationAddress := testutil.MakeAddress()
	sky := 500
	hash := testutil.RandSHA256(t)

	addr, err := NewSkycoinAddress(destinationAddress.String())
	assert.NoError(t, err)
	opt := NewTransferOptions()
	opt.SetValue("BurnFactor", "0.5")
	opt.SetValue("CoinHoursSelectionType", "auto")

	req := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type:        "auto",
			Mode:        "share",
			ShareFactor: "0.5",
		},
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
			},
		},
	}

	wreq := api.WalletCreateTransactionRequest{
		Unsigned:                 true,
		WalletID:                 "wallet",
		CreateTransactionRequest: req,
	}

	txn := coin.Transaction{
		Length:    100,
		Type:      0,
		InnerHash: hash,
	}
	crtTxn, err := api.NewCreateTransactionResponse(&txn, nil)
	require.NoError(t, err)
	crtTxn.Transaction.Fee = "500"

	mockSkyApiWalletCreateTransaction(global_mock, &wreq, crtTxn)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	quot, err := util.AltcoinQuotient(params.SkycoinTicker)
	require.NoError(t, err)

	destination := &SkycoinTransactionOutput{
		skyOut: readable.TransactionOutput{
			Address: addr.String(),
			Coins:   util.FormatCoins(uint64(sky*1e6), quot),
		}}

	ret, err := wlt.Transfer(destination, opt)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err := ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(sky), val)
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())

}

func TestRemoteWalletSendFromAddress(t *testing.T) {
	CleanGlobalMock()
	startAddress := testutil.MakeAddress()
	destinationAddress := testutil.MakeAddress()
	changeAddress := (testutil.MakeAddress()).String()
	sky := 500
	hash := testutil.RandSHA256(t)

	toAddr := &SkycoinTransactionOutput{
		skyOut: readable.TransactionOutput{
			Address: destinationAddress.String(),
			Coins:   strconv.Itoa(sky),
			Hours:   uint64(250),
		},
	}
	fromAddr, err := NewSkycoinAddress(startAddress.String())
	assert.NoError(t, err)
	chgAddr, err := NewSkycoinAddress(changeAddress)
	assert.NoError(t, err)

	opt1 := NewTransferOptions()
	opt1.SetValue("BurnFactor", "0.5")
	opt1.SetValue("CoinHoursSelectionType", "auto")

	req1 := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type:        "auto",
			Mode:        "share",
			ShareFactor: "0.5",
		},
		ChangeAddress: &changeAddress,
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
			},
		},
		Addresses: []string{startAddress.String()},
	}

	wreq1 := api.WalletCreateTransactionRequest{
		Unsigned:                 true,
		WalletID:                 "wallet1",
		CreateTransactionRequest: req1,
	}

	opt2 := NewTransferOptions()
	opt2.SetValue("BurnFactor", "0.5")
	opt2.SetValue("CoinHoursSelectionType", "manual")

	req2 := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type: "manual",
		},
		ChangeAddress: &changeAddress,
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
				Hours:   "250",
			},
		},
		Addresses: []string{startAddress.String()},
	}

	wreq2 := api.WalletCreateTransactionRequest{
		Unsigned:                 true,
		WalletID:                 "wallet2",
		CreateTransactionRequest: req2,
	}

	txn := coin.Transaction{
		Length:    100,
		Type:      0,
		InnerHash: hash,
	}
	crtTxn, err := api.NewCreateTransactionResponse(&txn, nil)
	require.NoError(t, err)
	crtTxn.Transaction.Fee = strconv.Itoa(sky)

	mockSkyApiWalletCreateTransaction(global_mock, &wreq1, crtTxn)
	mockSkyApiWalletCreateTransaction(global_mock, &wreq2, crtTxn)

	// Testing HoursSelection to auto
	wlt1 := &RemoteWallet{
		Id:          "wallet1",
		poolSection: PoolSection,
	}

	ret, err := wlt1.SendFromAddress([]core.Address{fromAddr}, []core.TransactionOutput{toAddr}, chgAddr, opt1)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err := ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, util.FormatCoins(uint64(sky), 10), util.FormatCoins(uint64(val), 10))
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())

	// Testing HoursSelection to manual
	wlt2 := &RemoteWallet{
		Id:          "wallet2",
		poolSection: PoolSection,
	}

	ret, err = wlt2.SendFromAddress([]core.Address{fromAddr}, []core.TransactionOutput{toAddr}, chgAddr, opt2)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err = ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, util.FormatCoins(uint64(sky), 10), util.FormatCoins(uint64(val), 10))
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())
}

func TestRemoteWalletSpend(t *testing.T) {
	CleanGlobalMock()
	destinationAddress := testutil.MakeAddress()
	changeAddress := (testutil.MakeAddress()).String()
	sky := 500
	hash := testutil.RandSHA256(t)

	toAddr := &SkycoinTransactionOutput{
		skyOut: readable.TransactionOutput{
			Address: destinationAddress.String(),
			Coins:   strconv.Itoa(sky),
			Hours:   uint64(250),
		},
	}
	chgAddr, err := NewSkycoinAddress(changeAddress)
	assert.NoError(t, err)
	opt := NewTransferOptions()
	opt.SetValue("BurnFactor", "0.5")
	opt.SetValue("CoinHoursSelectionType", "auto")

	req := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type:        "auto",
			Mode:        "share",
			ShareFactor: "0.5",
		},
		ChangeAddress: &changeAddress,
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
			},
		},
	}

	wreq := api.WalletCreateTransactionRequest{
		Unsigned:                 true,
		WalletID:                 "wallet",
		CreateTransactionRequest: req,
	}

	txn := coin.Transaction{
		Length:    100,
		Type:      0,
		InnerHash: hash,
	}
	crtTxn, err := api.NewCreateTransactionResponse(&txn, nil)
	require.NoError(t, err)
	crtTxn.Transaction.Fee = "500"

	mockSkyApiWalletCreateTransaction(global_mock, &wreq, crtTxn)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}

	ret, err := wlt.Spend(nil, []core.TransactionOutput{toAddr}, chgAddr, opt)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err := ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(sky), val)
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())
}

func TestRemoteWalletGenAddresses(t *testing.T) {
	CleanGlobalMock()
	pwd := "pwd"

	global_mock.On("Wallet", "wallet").Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "wallet",
				Encrypted: true,
			},
			Entries: []readable.WalletEntry{
				readable.WalletEntry{Address: "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"},
			},
		},
		nil)

	global_mock.On("NewWalletAddress", "wallet", 1, pwd).Return(
		[]string{"2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8", "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8"},
		nil)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	pwdReader := func(message string) (string, error) {
		return "pwd", nil
	}
	iter := wlt.GenAddresses(0, 0, 2, pwdReader)
	for iter.Next() {
		a := iter.Value()
		require.Equal(t, "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8", a.String())
	}
}

func TestRemoteWalletGetLoadedAddresses(t *testing.T) {

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	iter, err := wlt.GetLoadedAddresses()
	require.NoError(t, err)
	items := 0
	for iter.Next() {
		a := iter.Value()
		items++
		require.Equal(t, "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8", a.String())
	}
	require.Equal(t, 1, items)
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
	require.NoError(t, err)
	to := txn.Out[0]
	err = txn.PushOutput(to.Address, to.Coins, to.Hours)
	require.NoError(t, err)
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Duplicate output in transaction")

	// Invalid signature, empty
	txn, err = makeTransaction(t)
	require.NoError(t, err)
	txn.Sigs[0] = cipher.Sig{}
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Unsigned input in transaction")

	// Invalid signature, not empty
	// A stable invalid signature must be used because random signatures could appear valid
	// Note: Transaction.Verify() only checks that the signature is a minimally valid signature
	badSig := "9a0f86874a4d9541f58a1de4db1c1b58765a868dc6f027445d0a2a8a7bddd1c45ea559fcd7bef45e1b76ccdaf8e50bbebd952acbbea87d1cb3f7a964bc89bf1ed5"
	txn, err = makeTransaction(t)
	require.NoError(t, err)
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
	require.NoError(t, err)
	txn.Out[0].Coins = 0
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Zero coin output")

	// Output coin overflow
	txn, err = makeTransaction(t)
	require.NoError(t, err)
	txn.Out[0].Coins = math.MaxUint64 - 3e6
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	testutil.RequireError(t, uiTxn.VerifySigned(), "Output coins overflow")

	// Output coins are not multiples of 1e6 (valid, decimal restriction is not enforced here)
	txn, err = makeTransaction(t)
	require.NoError(t, err)
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
	require.NoError(t, err)
	txn.Out[0].Coins = 10e6
	txn.Out[1].Coins = 1e6
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn = makeUninjectedTransaction(t, &txn, 0)
	require.NoError(t, uiTxn.VerifySigned())
}

func TestUninjectedTransactionVerifyUnsigned(t *testing.T) {
	txn, _, _, err := makeTransactionMultipleInputs(t, 2)
	require.NoError(t, err)
	uiTxn := makeUninjectedTransaction(t, &txn, 0)
	err = uiTxn.VerifyUnsigned()
	testutil.RequireError(t, err, "Unsigned transaction must contain a null signature")

	// Invalid signature, not empty
	// A stable invalid signature must be used because random signatures could appear valid
	// Note: Transaction.Verify() only checks that the signature is a minimally valid signature
	badSig := "9a0f86874a4d9541f58a1de4db1c1b58765a868dc6f027445d0a2a8a7bddd1c45ea559fcd7bef45e1b76ccdaf8e50bbebd952acbbea87d1cb3f7a964bc89bf1ed5"
	txn, _, _, err = makeTransactionMultipleInputs(t, 2)
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
	txn, _, _, err = makeTransactionMultipleInputs(t, 3)
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

// w, err = walletSet.CreateWallet(walletID, kd.Mnemonic, wallet.WalletTypeDeterministic, false, func(string) (string, error) { return "", nil }, 0)
func TestTransactionSignInput(t *testing.T) {
	TransactionSignInputTestImpl(t, nil)
}

func TestTransactionSignInputs(t *testing.T) {
	// Build transaction step by step
	txn := &coin.Transaction{}
	ux, kd, err := makeUxOutWithSecret(t)
	require.NoError(t, err)
	err = txn.PushInput(ux.Hash())
	require.NoError(t, err)
	wallets := makeLocalWalletsFromKeyData(t, []KeyData{*kd})
	wallet := wallets[0]
	seed, seckeys, err2 := cipher.GenerateDeterministicKeyPairsSeed([]byte(kd.Mnemonic), kd.AddressIndex+1)
	require.NoError(t, err2)
	require.Equal(t, kd.SecKey, seckeys[kd.AddressIndex])
	p2, _, err3 := cipher.GenerateDeterministicKeyPair(seed)
	require.NoError(t, err3)
	wallet.GenAddresses(core.AccountAddress, uint32(kd.AddressIndex), 2, nil)
	ux2 := coin.UxOut{
		Head: coin.UxHead{
			Time:  100,
			BkSeq: 2,
		},
		Body: coin.UxBody{
			SrcTransaction: testutil.RandSHA256(t),
			Address:        cipher.AddressFromPubKey(p2),
			Coins:          1e6,
			Hours:          100,
		},
	}
	err = txn.PushInput(ux2.Hash())
	require.NoError(t, err)
	err = txn.PushOutput(makeAddress(), 40, 80)
	require.NoError(t, err)
	require.Equal(t, len(txn.Sigs), 0)
	err = txn.UpdateHeader()
	require.NoError(t, err)
	uiTxn := makeUninjectedTransaction(t, txn, 0)
	isFullySigned, err := uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)

	// Mock Skycoin API calls
	mockSkyApiUxOut(global_mock, ux)
	mockSkyApiUxOut(global_mock, ux2)

	// Valid signing
	h := txn.HashInner()
	signedCoreTxn, err := wallet.Sign(uiTxn, nil, util.EmptyPassword, []string{"0", "1"})
	require.NoError(t, err)
	signedTxn, isUninjected := signedCoreTxn.(*SkycoinUninjectedTransaction)
	require.True(t, isUninjected)
	isFullySigned, err = signedTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)
	require.Equal(t, len(signedTxn.txn.Sigs), 2)
	h2 := signedTxn.txn.HashInner()
	require.Equal(t, h2, h)
	p := kd.PubKey
	a := cipher.AddressFromPubKey(p)
	a2 := cipher.AddressFromPubKey(p2)
	require.NoError(t, cipher.VerifyAddressSignedHash(a, signedTxn.txn.Sigs[0], cipher.AddSHA256(h, signedTxn.txn.In[0])))
	require.NoError(t, cipher.VerifyAddressSignedHash(a2, signedTxn.txn.Sigs[1], cipher.AddSHA256(h, signedTxn.txn.In[1])))
	require.Error(t, cipher.VerifyAddressSignedHash(a, signedTxn.txn.Sigs[1], cipher.AddSHA256(h, signedTxn.txn.In[0])))
	require.Error(t, cipher.VerifyAddressSignedHash(a2, signedTxn.txn.Sigs[0], cipher.AddSHA256(h, signedTxn.txn.In[1])))
}

func TestLocalWalletTransfer(t *testing.T) {
	CleanGlobalMock()
	destinationAddress := testutil.MakeAddress()
	sky := 500
	wlt := makeLocalWallet(t)

	addr, err := NewSkycoinAddress(destinationAddress.String())
	assert.NoError(t, err)
	loadedAddrs, err := wlt.GetLoadedAddresses()
	require.NoError(t, err)
	addrs := make([]string, 0)
	for loadedAddrs.Next() {
		addrs = append(addrs, loadedAddrs.Value().String())
	}

	opt := NewTransferOptions()
	opt.SetValue("BurnFactor", "0.5")
	opt.SetValue("CoinHoursSelectionType", "auto")

	req := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type:        "auto",
			Mode:        "share",
			ShareFactor: "0.5",
		},
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
			},
		},
		Addresses: addrs,
	}

	hash := testutil.RandSHA256(t)
	txn := coin.Transaction{
		Length:    100,
		Type:      0,
		InnerHash: hash,
	}
	crtTxn, err := api.NewCreateTransactionResponse(&txn, nil)
	require.NoError(t, err)
	crtTxn.Transaction.Fee = "500"

	mockSkyApiCreateTransaction(global_mock, &req, crtTxn)

	quot, err := util.AltcoinQuotient(params.SkycoinTicker)
	require.NoError(t, err)

	destination := &SkycoinTransactionOutput{
		skyOut: readable.TransactionOutput{
			Address: addr.String(),
			Coins:   util.FormatCoins(uint64(sky*1e6), quot),
		}}

	ret, err := wlt.Transfer(destination, opt)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err := ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(sky), val)
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())
}

func TestLocalWalletSendFromAddress(t *testing.T) {
	CleanGlobalMock()
	startAddress := testutil.MakeAddress()
	destinationAddress := testutil.MakeAddress()
	changeAddress := (testutil.MakeAddress()).String()
	sky := 500
	wlt := makeLocalWallet(t)

	toAddr := &SkycoinTransactionOutput{
		skyOut: readable.TransactionOutput{
			Address: destinationAddress.String(),
			Coins:   strconv.Itoa(sky),
			Hours:   uint64(250),
		},
	}
	fromAddr, err := NewSkycoinAddress(startAddress.String())
	assert.NoError(t, err)
	chgAddr, err := NewSkycoinAddress(changeAddress)
	assert.NoError(t, err)

	opt1 := NewTransferOptions()
	opt1.SetValue("BurnFactor", "0.5")
	opt1.SetValue("CoinHoursSelectionType", "auto")

	req1 := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type:        "auto",
			Mode:        "share",
			ShareFactor: "0.5",
		},
		ChangeAddress: &changeAddress,
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
			},
		},
		Addresses: []string{startAddress.String()},
	}

	opt2 := NewTransferOptions()
	opt2.SetValue("BurnFactor", "0.5")
	opt2.SetValue("CoinHoursSelectionType", "manual")

	req2 := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type: "manual",
		},
		ChangeAddress: &changeAddress,
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
				Hours:   "250",
			},
		},
		Addresses: []string{startAddress.String()},
	}

	hash := testutil.RandSHA256(t)
	txn := coin.Transaction{
		Length:    100,
		Type:      0,
		InnerHash: hash,
	}
	crtTxn, err := api.NewCreateTransactionResponse(&txn, nil)
	require.NoError(t, err)
	crtTxn.Transaction.Fee = strconv.Itoa(sky)

	mockSkyApiCreateTransaction(global_mock, &req1, crtTxn)
	mockSkyApiCreateTransaction(global_mock, &req2, crtTxn)

	// Testing HoursSelection to auto
	ret, err := wlt.SendFromAddress([]core.Address{fromAddr}, []core.TransactionOutput{toAddr}, chgAddr, opt1)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err := ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, util.FormatCoins(uint64(sky), 10), util.FormatCoins(uint64(val), 10))
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())

	// Testing HoursSelection to manual
	ret, err = wlt.SendFromAddress([]core.Address{fromAddr}, []core.TransactionOutput{toAddr}, chgAddr, opt2)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err = ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, util.FormatCoins(uint64(sky), 10), util.FormatCoins(uint64(val), 10))
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())
}

func TestLocalWalletSpend(t *testing.T) {
	CleanGlobalMock()
	destinationAddress := testutil.MakeAddress()
	changeAddress := (testutil.MakeAddress()).String()
	sky := 500
	wlt := makeLocalWallet(t)

	toAddr := &SkycoinTransactionOutput{
		skyOut: readable.TransactionOutput{
			Address: destinationAddress.String(),
			Coins:   strconv.Itoa(sky),
			Hours:   uint64(250),
		},
	}
	chgAddr, err := NewSkycoinAddress(changeAddress)
	assert.NoError(t, err)
	opt := NewTransferOptions()
	opt.SetValue("BurnFactor", "0.5")
	opt.SetValue("CoinHoursSelectionType", "auto")

	req := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		HoursSelection: api.HoursSelection{
			Type:        "auto",
			Mode:        "share",
			ShareFactor: "0.5",
		},
		ChangeAddress: &changeAddress,
		To: []api.Receiver{
			api.Receiver{
				Address: destinationAddress.String(),
				Coins:   strconv.Itoa(sky),
			},
		},
	}

	hash := testutil.RandSHA256(t)
	txn := coin.Transaction{
		Length:    100,
		Type:      0,
		InnerHash: hash,
	}
	crtTxn, err := api.NewCreateTransactionResponse(&txn, nil)
	require.NoError(t, err)
	crtTxn.Transaction.Fee = "500"

	mockSkyApiCreateTransaction(global_mock, &req, crtTxn)

	ret, err := wlt.Spend(nil, []core.TransactionOutput{toAddr}, chgAddr, opt)
	require.NoError(t, err)
	require.NotNil(t, ret)
	val, err := ret.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(sky), val)
	require.Equal(t, crtTxn.Transaction.TxID, ret.GetId())
}

func TestSkycoinWalletTypes(t *testing.T) {
	var wltSet core.WalletSet = &SkycoinRemoteWallet{}
	require.Equal(t, wallet.WalletTypeBip44, wltSet.DefaultWalletType())
	require.Equal(t, []string{wallet.WalletTypeDeterministic, wallet.WalletTypeBip44}, wltSet.SupportedWalletTypes())

	wltSet = &SkycoinLocalWallet{}
	require.Equal(t, wallet.WalletTypeBip44, wltSet.DefaultWalletType())
	require.Equal(t, []string{wallet.WalletTypeDeterministic, wallet.WalletTypeBip44}, wltSet.SupportedWalletTypes())
}

func TestLookupWalletOk(t *testing.T) {
	// Giving
	ws := &mocks2.PersistibleSet{}
	_, kd, err := makeUxOutWithSecret(t)
	require.NoError(t, err)
	wallets := makeLocalWalletsFromKeyData(t, []KeyData{*kd})
	ws.On("ListWallets").Return(NewSkycoinWalletIterator(wallets))
	we := NewWalletDirectory(os.TempDir())
	we.SetWltService(ws)
	addrs := wallets[0].GenAddresses(core.AccountAddress, 0, 1, nil)
	require.True(t, addrs.Next())
	addr1 := addrs.Value()

	// When
	wltFound, err := we.LookupWallet(addr1.String())

	// Then
	require.NoError(t, err)
	require.Equal(t, wallets[0], wltFound)
}

func TestLookupWalletFailAsExpected(t *testing.T) {
	// Giving
	ws := &mocks2.PersistibleSet{}
	_, kd, err := makeUxOutWithSecret(t)
	require.NoError(t, err)
	wallets := makeLocalWalletsFromKeyData(t, []KeyData{*kd})
	ws.On("ListWallets").Return(NewSkycoinWalletIterator(wallets))
	we := NewWalletDirectory(os.TempDir())
	we.SetWltService(ws)
	addrs := wallets[0].GenAddresses(core.AccountAddress, 0, 1, nil)
	require.True(t, addrs.Next())
	addr1 := addrs.Value()
	strAddr := addr1.String()
	strAddr = strAddr[len(strAddr)/2:] + strAddr[:len(strAddr)/2]

	// When
	wltFound, err := we.LookupWallet(strAddr)

	// Then
	require.Error(t, err)
	require.Equal(t, errors.ErrWltFromAddrNotFound, err)
	require.Equal(t, nil, wltFound)
}