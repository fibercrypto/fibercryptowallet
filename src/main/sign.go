package local

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/signutil"
)

func (m *fibercryptoAltcoinManager) AttachSignService(signSrv core.TxnSigner) error {
	if signSrv != nil {
		uid, err := signSrv.GetSignerUID()
		if err != nil {
			return err
		}
		m.signers[uid] = signSrv
	}
	return nil
}

func (m *fibercryptoAltcoinManager) EnumerateSignServices() core.TxnSignerIterator {
	return signutil.NewTxnSignerIteratorFromMap(m.signers)
}

func (m *fibercryptoAltcoinManager) LookupSignService(id core.UID) core.TxnSigner {
	if signSrv, isFound := m.signers[id]; isFound {
		return signSrv
	}
	return nil
}

func (m *fibercryptoAltcoinManager) RemoveSignService(signSrv core.TxnSigner) error {
	uid, err := signSrv.GetSignerUID()
	if err != nil {
		return err
	}
	if _, isBound := m.signers[uid]; isBound {
		delete(m.signers, uid)
		return nil
	}
	return errors.ErrInvalidValue
}

// SignServicesForTxn returns an object to iterate over signing srategies supported for a given transaction
func (m *fibercryptoAltcoinManager) SignServicesForTxn(wlt core.Wallet, txn core.Transaction) core.TxnSignerIterator {
	return signutil.FilterSignersFromMap(
		m.signers,
		func(signer core.TxnSigner) bool {
			canSign, err := signer.ReadyForTxn(wlt, txn)
			return err == nil && canSign
		})
}
