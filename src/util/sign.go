package util

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/errors"
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
)

// AttachSignService registers a signing strategy for use by wallets
func AttachSignService(signer core.TxnSigner) error {
	return local.LoadAltcoinManager().AttachSignService(signer)
}

// LookupSignService instantiate global signing straegy identified by UID
func LookupSignService(signerID core.UID) core.TxnSigner {
	return local.LoadAltcoinManager().LookupSignService(signerID)
}

// RemoveSignService detaches a signing strategy
func RemoveSignService(signerID core.UID) error {
	signer := LookupSignService(signerID)
	if signer == nil {
		return errors.ErrInvalidID
	}
	return local.LoadAltcoinManager().RemoveSignService(signer)
}

// EnumerateSignServices returns an object to iterate over global signing srategies
func EnumerateSignServices() core.TxnSignerIterator {
	return local.LoadAltcoinManager().EnumerateSignServices()
}

// SignServicesForTxn returns an object to iterate over strategies supported to sign a given transaction on behalf of a wallet
func SignServicesForTxn(wlt core.Wallet, txn core.Transaction) core.TxnSignerIterator {
	return local.LoadAltcoinManager().SignServicesForTxn(wlt, txn)
}

// ReadyForTxn determines whether global signer identified by UID can be used by wallet to sign given transaction
func ReadyForTxn(signerID core.UID, wallet core.Wallet, txn core.Transaction) (bool, error) {
	signer := LookupSignService(signerID)
	if signer == nil {
		return false, errors.ErrInvalidID
	}
	return signer.ReadyForTxn(wallet, txn)
}

// SignTransaction sign transaction partially or in full with signer identified by UID
func SignTransaction(signerID core.UID, txn core.Transaction, pwd core.PasswordReader, indices []string) (core.Transaction, error) {
	signer := LookupSignService(signerID)
	if signer == nil {
		return nil, errors.ErrInvalidID
	}
	return signer.SignTransaction(txn, pwd, indices)
}

// GetSignerDescription human readable caption for signing strategy identified by UID
func GetSignerDescription(signerID core.UID) (string, error) {
	signer := LookupSignService(signerID)
	if signer == nil {
		return "", errors.ErrInvalidID
	}
	return signer.GetSignerDescription(), nil
}

// LookupSignServiceForWallet instantiate signing straegy identified by UID. Fall back to wallet if empty
func LookupSignServiceForWallet(wlt core.Wallet, signerID core.UID) (core.TxnSigner, error) {
	if signerID == "" {
		wltSigner, isTxnSigner := wlt.(core.TxnSigner)
		if !isTxnSigner {
			logUtil.WithError(errors.ErrInvalidID).Errorf("Wallet %v can not sign transactions", wlt)
			return nil, errors.ErrWalletCantSign
		}
		return wltSigner, nil
	}
	if signer := LookupSignService(signerID); signer != nil {
		return signer, nil
	}
	return nil, errors.ErrInvalidID
}

type signingKeyPair struct {
	wallet core.Wallet
	signer core.UID
}

// MultiWalletSign generic strategy for using multiple wallets to sign a transaction
func GenericMultiWalletSign(txn core.Transaction, signSpec []core.InputSignDescriptor, pwds map[string]core.PasswordReader) (signedTxn core.Transaction, err error) {
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

		signer, err := LookupSignServiceForWallet(signPair.wallet, signPair.signer)
		if err != nil {
			logUtil.WithError(err).Errorf("Unknown signer %s specified for signing inputs %v of wallet %v", string(signPair.signer), indices, signPair.wallet)
			return nil, errors.ErrInvalidID
		}
		pwd := pwds[signPair.wallet.GetId()]
		signedTxn, err = signPair.wallet.Sign(signedTxn, signer, pwd, indices)
		if err != nil {
			logUtil.WithError(err).Errorf("Error signing inputs %v of wallet %v with signer %s", indices, signPair.wallet, string(signPair.signer))
			return nil, err
		}

	}
	return signedTxn, nil
}
