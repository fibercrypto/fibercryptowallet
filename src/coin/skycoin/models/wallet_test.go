package skycoin

import (
	"strconv"
	"testing"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/stretchr/testify/assert"

	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/testutil"
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

	global_mock.On("Transactions", []string{}).Return(nil, nil)
	global_mock.On("Transactions", []string{addressesN[0]}).Return(
		[]readable.TransactionWithStatus{},
		nil)
	global_mock.On("Transactions", []string{addressesN[1]}).Return(
		[]readable.TransactionWithStatus{
			readable.TransactionWithStatus{
				Status: readable.TransactionStatus{
					Confirmed: true,
				},
			},
		},
		nil)
	global_mock.On("Transactions", []string{addressesN[2]}).Return(
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

	thxF := &TransactionFinder{}

	mask, err := thxF.AddressesActivity([]cipher.Address{})
	assert.Nil(t, err)
	assert.Equal(t, 0, len(mask))
	assert.Equal(t, []bool{}, mask)

	mask, err = thxF.AddressesActivity(addresses)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(mask))
	assert.Equal(t, false, mask[0])
	for i := 1; i < 3; i++ {
		assert.Equal(t, true, mask[i])
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
		assert.Equal(t, "wallet", wlt.GetLabel())
		assert.Equal(t, "FiberCrypto", wlt.GetId())
	}
}

func TestSkycoinRemoteWalletCreateWallet(t *testing.T) {

	seed, label, pwd, scanN := "seed", "label", "pwd", 666

	wltOpt1 := api.CreateWalletOptions{
		Type:     WalletTypeDeterministic,
		Seed:     seed,
		Label:    label,
		Password: pwd,
		ScanN:    scanN,
		Encrypt:  true,
	}
	wltOpt2 := api.CreateWalletOptions{
		Type:    WalletTypeDeterministic,
		Seed:    seed,
		Label:   label,
		ScanN:   scanN,
		Encrypt: false,
	}

	global_mock.On("CreateWallet", wltOpt1).Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "walletEncrypted",
				Encrypted: true,
			},
		},
		nil)
	global_mock.On("CreateWallet", wltOpt2).Return(
		&api.WalletResponse{
			Meta: readable.WalletMeta{
				Coin:      "Sky",
				Filename:  "FiberCrypto",
				Label:     "walletNonEncrypted",
				Encrypted: false,
			},
		},
		nil)

	wltSrv := &SkycoinRemoteWallet{poolSection: PoolSection}
	pwdReader := func(message string) (string, error) {
		return "pwd", nil
	}

	wlt1, err := wltSrv.CreateWallet(label, seed, true, pwdReader, scanN)
	assert.Nil(t, err)
	assert.Equal(t, "walletEncrypted", wlt1.GetLabel())
	assert.Equal(t, "FiberCrypto", wlt1.GetId())

	wlt2, err := wltSrv.CreateWallet(label, seed, false, pwdReader, scanN)
	assert.Nil(t, err)
	assert.Equal(t, "walletNonEncrypted", wlt2.GetLabel())
	assert.Equal(t, "FiberCrypto", wlt2.GetId())
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
	assert.Nil(t, err)
	assert.Equal(t, true, encrypted)

	encrypted, err = wltSrv.IsEncrypted("nonEncrypted")
	assert.Nil(t, err)
	assert.Equal(t, false, encrypted)
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
	assert.Equal(t, "wallet", wlt.GetLabel())
	assert.Equal(t, "FiberCrypto", wlt.GetId())
}

func TestRemoteWalletSign(t *testing.T) {
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
	assert.Nil(t, err)

	walletSignTxn := api.WalletSignTransactionRequest{
		EncodedTransaction: encodedResponse,
		WalletID:           "wallet",
		Password:           "password",
		SignIndexes:        nil,
	}

	crtTxn := &api.CreateTransactionResponse{
		Transaction: api.CreatedTransaction{
			Fee:       "100",
			InnerHash: hash.Hex(),
		},
	}
	tx := &coin.Transaction{
		InnerHash: hash,
	}
	b, _ := tx.Serialize()
	crtTxn.Transaction.TxID = cipher.SumSHA256(b).Hex()

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

	ret, err := wlt.Sign(&unTxn, "source", pwdReader, nil)
	assert.Nil(t, err)
	value, err := ret.ComputeFee(CoinHour)
	assert.Nil(t, err)
	assert.Equal(t, uint64(100), value)
	assert.Equal(t, crtTxn.Transaction.TxID, ret.GetId())
}

type TransferOptions struct {
	values map[string]interface{}
}

func (tOpt *TransferOptions) GetValue(key string) interface{} {
	return tOpt.values[key]
}

func (tOpt *TransferOptions) AddKeyValue(key string, value interface{}) {
	tOpt.values[key] = value
}

func NewTransferOptions() *TransferOptions {
	tOptions := TransferOptions{
		values: make(map[string]interface{}, 0),
	}
	return &tOptions
}

func TestRemoteWalletTransfer(t *testing.T) {
	CleanGlobalMock()
	destinationAddress := testutil.MakeAddress()
	sky := 500
	hash := testutil.RandSHA256(t)

	addr := &SkycoinAddress{
		address: destinationAddress.String(),
	}

	opt := NewTransferOptions()
	opt.AddKeyValue("BurnFactor", "0.5")
	opt.AddKeyValue("CoinHoursSelectionType", "auto")

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

	crtTxn := &api.CreateTransactionResponse{
		Transaction: api.CreatedTransaction{
			Fee:       "500",
			InnerHash: hash.Hex(),
		},
	}
	tx := &coin.Transaction{
		InnerHash: hash,
	}
	b, _ := tx.Serialize()
	crtTxn.Transaction.TxID = cipher.SumSHA256(b).Hex()

	global_mock.On("WalletCreateTransaction", wreq).Return(
		crtTxn,
		nil)

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}

	txn, err := wlt.Transfer(addr, uint64(sky*1e6), opt)
	assert.Nil(t, err)
	val, err := txn.ComputeFee(CoinHour)
	assert.Nil(t, err)
	assert.Equal(t, uint64(sky), val)
	assert.Equal(t, crtTxn.Transaction.TxID, txn.GetId())

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
	fromAddr := &SkycoinAddress{
		address: startAddress.String(),
	}
	chgAddr := &SkycoinAddress{
		address: changeAddress,
	}

	opt1 := NewTransferOptions()
	opt1.AddKeyValue("BurnFactor", "0.5")
	opt1.AddKeyValue("CoinHoursSelectionType", "auto")

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

	crtTxn1 := &api.CreateTransactionResponse{
		Transaction: api.CreatedTransaction{
			Fee:       strconv.Itoa(sky),
			InnerHash: hash.Hex(),
		},
	}

	opt2 := NewTransferOptions()
	opt2.AddKeyValue("BurnFactor", "0.5")
	opt2.AddKeyValue("CoinHoursSelectionType", "manual")

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

	crtTxn2 := &api.CreateTransactionResponse{
		Transaction: api.CreatedTransaction{
			Fee:       strconv.Itoa(sky),
			InnerHash: hash.Hex(),
		},
	}

	tx := &coin.Transaction{
		InnerHash: hash,
	}
	b, _ := tx.Serialize()
	crtTxn1.Transaction.TxID = cipher.SumSHA256(b).Hex()
	crtTxn2.Transaction.TxID = crtTxn1.Transaction.TxID

	global_mock.On("WalletCreateTransaction", wreq1).Return(
		crtTxn1,
		nil)
	global_mock.On("WalletCreateTransaction", wreq2).Return(
		crtTxn2,
		nil)

	wlt1 := &RemoteWallet{
		Id:          "wallet1",
		poolSection: PoolSection,
	}

	txn, err := wlt1.SendFromAddress([]core.Address{fromAddr}, []core.TransactionOutput{toAddr}, chgAddr, opt1)
	assert.Nil(t, err)
	val, err := txn.ComputeFee(CoinHour)
	assert.Nil(t, err)
	assert.Equal(t, util.FormatCoins(uint64(sky), 10), util.FormatCoins(uint64(val), 10))
	assert.Equal(t, crtTxn1.Transaction.TxID, txn.GetId())

	wlt2 := &RemoteWallet{
		Id:          "wallet2",
		poolSection: PoolSection,
	}

	txn, err = wlt2.SendFromAddress([]core.Address{fromAddr}, []core.TransactionOutput{toAddr}, chgAddr, opt2)
	assert.Nil(t, err)
	val, err = txn.ComputeFee(CoinHour)
	assert.Nil(t, err)
	assert.Equal(t, util.FormatCoins(uint64(sky), 10), util.FormatCoins(uint64(val), 10))
	assert.Equal(t, crtTxn2.Transaction.TxID, txn.GetId())
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
				readable.WalletEntry{Address: "addr"},
			},
		},
		nil)

	global_mock.On("NewWalletAddress", "wallet", 1, pwd).Return(
		[]string{"addr", "addr"},
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
		assert.Equal(t, "addr", a.String())
	}
}

func TestRemoteWalletGetLoadedAddresses(t *testing.T) {

	wlt := &RemoteWallet{
		Id:          "wallet",
		poolSection: PoolSection,
	}
	iter, err := wlt.GetLoadedAddresses()
	assert.Nil(t, err)
	items := 0
	for iter.Next() {
		a := iter.Value()
		items++
		assert.Equal(t, "addr", a.String())
	}
	assert.Equal(t, 1, items)
}
