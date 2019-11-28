package skycoin

import (
	"testing"

	"github.com/skycoin/skycoin/src/cipher/bip39"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/stretchr/testify/require"
)

func generateRandomKeyData(t *testing.T) *KeyData {
	entropy, err := bip39.NewEntropy(128)
	require.NoError(t, err)
	mnemonic, err := bip39.NewMnemonic(entropy)
	require.NoError(t, err)
	require.NoError(t, err)
	pubKey, secKey, err := cipher.GenerateDeterministicKeyPair([]byte(mnemonic))
	require.NoError(t, err)

	kd := &KeyData{
		AddressIndex: 0,
		Entropy:      entropy,
		Mnemonic:     mnemonic,
		PubKey:       pubKey,
		SecKey:       secKey,
	}

	return kd
}

func makeUxBodyWithRandomSecret(t *testing.T) (coin.UxBody, *KeyData) {
	keydata := generateRandomKeyData(t)
	return coin.UxBody{
		SrcTransaction: testutil.RandSHA256(t),
		Address:        cipher.AddressFromPubKey(keydata.PubKey),
		Coins:          1e6,
		Hours:          100,
	}, keydata
}

func makeUxWithRandomSecret(t *testing.T) (coin.UxOut, *KeyData) {
	body, kd := makeUxBodyWithRandomSecret(t)
	return coin.UxOut{
		Head: coin.UxHead{
			Time:  100,
			BkSeq: 2,
		},
		Body: body,
	}, kd
}

func makeTransactionFromMultipleWallets(t *testing.T, n int) (coin.Transaction, []KeyData, []coin.UxOut) {
	uxs := make([]coin.UxOut, n)
	keysdata := make([]KeyData, n)
	secs := make([]cipher.SecKey, n)

	for i := 0; i < n; i++ {
		ux, kd := makeUxWithRandomSecret(t)
		uxs[i] = ux
		secs[i] = kd.SecKey
		keysdata[i] = *kd
	}

	return makeTransactionFromUxOuts(t, uxs, secs), keysdata, uxs
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
