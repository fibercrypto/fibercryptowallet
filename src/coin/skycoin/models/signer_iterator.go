package skycoin

import "github.com/fibercrypto/FiberCryptoWallet/src/core"

type SignerIterator struct { //Implements TxnSignerIterator interfaces
	current   int
	signers []core.TxnSigner
}

func (it *SignerIterator) Value() core.TxnSigner {
	return it.signers[it.current]
}

func (it *SignerIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *SignerIterator) HasNext() bool {
	return (it.current + 1) < len(it.signers)
}

func (it *SignerIterator) Count() int {
	return len(it.signers)
}

func NewSignerIterator(signers []core.TxnSigner) *SignerIterator {
	return &SignerIterator{signers: signers, current: -1}
}
