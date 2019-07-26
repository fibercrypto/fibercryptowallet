package wallet

type SkycoinWalletIterator struct {
	current int
	wallets []Wallet
}

func (it *SkycoinWalletIterator) Value() Wallet {
	return it.wallets[it.current]
}

func (it *SkycoinWalletIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *SkycoinWalletIterator) HasNext() bool {
	if (it.current + 1) >= len(it.wallets) {
		return false
	}
	return true
}

func NewSkycoinWalletIterator(wallets []Wallet) *NewSkycoinWalletIterator {
	return &SkycoinWalletIterator{wallets: wallets, current: -1}
}
