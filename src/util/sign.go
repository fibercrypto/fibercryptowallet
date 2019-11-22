package util

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
)

type signingKeyPair struct {
	wallet core.Wallet
	signer core.UID
}

// MultiWalletSign generic strategy for using multiple wallets to sign a transaction
func GenericMultiWalletSign(txn core.Transaction, signSpec []core.InputSignDescriptor, pwd core.PasswordReader) (signedTxn core.Transaction, err error) {
	groups := make(map[signingKeyPair][]string)
	// Aggregate inputs by wallet,signer combination
	for _, descriptor := range signSpec {
		key := signingKeyPair{
			wallet: descriptor.Wallet,
			signer: descriptor.SignerID,
		}
		inputs, isNotEmpty := groups[key]
		if !isNotEmpty {
			inputs = []string{}
		}
		groups[key] = append(inputs, descriptor.InputIndex)
	}
	signedTxn = txn

	for signPair, indices := range groups {

		signedTxn, err = signPair.wallet.Sign(signedTxn, signPair.signer, pwd, indices)
		if err != nil {
			logUtil.WithError(err).Errorf("Error signing inputs %v of wallet %v with signer %s", indices, signPair.wallet, string(signPair.signer))
			return nil, err
		}

	}
	return signedTxn, nil
}
