package skycoin

import (
	"testing"

	"github.com/skycoin/skycoin/src/cipher/bip39"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/testutil"
	"github.com/stretchr/testify/require"
)

func generateRandomKeyData(t *testing.T) (*KeyData, error) {
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

	return kd, nil
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
