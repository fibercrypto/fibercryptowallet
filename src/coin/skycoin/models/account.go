package skycoin

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/skycoin/skycoin/src/cli"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/util/droplet"
	"github.com/skycoin/skycoin/src/wallet"
)

func (addr SkycoinAddress) GetBalance(ticker string) (uint64, error) {

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return 0, err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	bl, err := c.Balance([]string{addr.address})

	if err != nil {
		return 0, err
	}

	if ticker == Sky {
		return bl.Confirmed.Coins, nil
	} else if ticker == CoinHour {
		return bl.Confirmed.Hours, nil
	} else {
		return 0, errorTickerInvalid{ticker}
	}
}
func (addr SkycoinAddress) ListAssets() []string {
	return []string{Sky, CoinHour}
}
func (addr SkycoinAddress) ScanUnspentOutputs() core.TransactionOutputIterator { //------TODO
	return nil
}
func (addr SkycoinAddress) ListTransactions() core.TransactionIterator { //------TODO
	return nil
}
func (addr SkycoinAddress) ListPendingTransactions() core.TransactionIterator { //------TODO
	return nil
}

func (wlt RemoteWallet) GetBalance(ticker string) (uint64, error) {
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		return 0, err
	}
	defer core.GetMultiPool().Return(wlt.poolSection, c)
	bl, err := c.WalletBalance(wlt.Id)
	if err != nil {
		return 0, err
	}

	if ticker == Sky {
		return bl.Confirmed.Coins, nil
	} else if ticker == CoinHour {
		return bl.Confirmed.Hours, nil
	} else {
		return 0, errorTickerInvalid{ticker}
	}

}

func (wlt RemoteWallet) ListAssets() []string {
	return []string{Sky, CoinHour}
}

func (wlt RemoteWallet) ScanUnspentOutputs() core.TransactionOutputIterator { //------TODO
	return nil
}

func (wlt RemoteWallet) ListTransactions() core.TransactionIterator { //------TODO
	return nil
}
func (wlt RemoteWallet) ListPendingTransactions() core.TransactionIterator { //------TODO
	return nil
}

func (wlt LocalWallet) GetBalance(ticker string) (uint64, error) {
	walletName := filepath.Join(wlt.WalletDir, wlt.Id)
	walletLoaded, err := wallet.Load(walletName)
	if err != nil {
		return 0, err
	}
	var addrs []string
	addresses := walletLoaded.GetAddresses()
	for _, addr := range addresses {
		addrs = append(addrs, addr.String())
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return 0, err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	outs, err := c.OutputsForAddresses(addrs)

	if err != nil {
		return 0, err
	}

	bl, err := getBalanceOfAddresses(outs, addrs)

	if ticker == Sky {
		return strconv.ParseUint(bl.Confirmed.Coins, 10, 64)
	} else if ticker == CoinHour {
		return strconv.ParseUint(bl.Confirmed.Hours, 10, 64)
	} else {
		return 0, errorTickerInvalid{ticker}
	}

}

func (wlt LocalWallet) ListAssets() []string {
	return []string{Sky, CoinHour}
}

func (wlt LocalWallet) ScanUnspentOutputs() core.TransactionOutputIterator { //------TODO
	return nil
}

func (wlt LocalWallet) ListTransactions() core.TransactionIterator { //------TODO
	return nil
}
func (wlt LocalWallet) ListPendingTransactions() core.TransactionIterator { //------TODO
	return nil
}

func getBalanceOfAddresses(outs *readable.UnspentOutputsSummary, addrs []string) (*cli.BalanceResult, error) {
	addrsMap := make(map[string]struct{}, len(addrs))
	for _, a := range addrs {
		addrsMap[a] = struct{}{}
	}

	addrBalances := make(map[string]struct {
		confirmed, spendable, expected wallet.Balance
	}, len(addrs))

	// Count confirmed balances
	for _, o := range outs.HeadOutputs {
		if _, ok := addrsMap[o.Address]; !ok {
			return nil, fmt.Errorf("Found address %s in GetUnspentOutputs result, but this address wasn't requested", o.Address)
		}

		amt, err := droplet.FromString(o.Coins)
		if err != nil {
			return nil, fmt.Errorf("droplet.FromString failed: %v", err)
		}

		b := addrBalances[o.Address]
		b.confirmed.Coins += amt
		b.confirmed.Hours += o.CalculatedHours

		addrBalances[o.Address] = b
	}

	// Count spendable balances
	for _, o := range outs.SpendableOutputs() {
		if _, ok := addrsMap[o.Address]; !ok {
			return nil, fmt.Errorf("Found address %s in GetUnspentOutputs result, but this address wasn't requested", o.Address)
		}

		amt, err := droplet.FromString(o.Coins)
		if err != nil {
			return nil, fmt.Errorf("droplet.FromString failed: %v", err)
		}

		b := addrBalances[o.Address]
		b.spendable.Coins += amt
		b.spendable.Hours += o.CalculatedHours

		addrBalances[o.Address] = b
	}

	// Count predicted balances
	for _, o := range outs.ExpectedOutputs() {
		if _, ok := addrsMap[o.Address]; !ok {
			return nil, fmt.Errorf("Found address %s in GetUnspentOutputs result, but this address wasn't requested", o.Address)
		}

		amt, err := droplet.FromString(o.Coins)
		if err != nil {
			return nil, fmt.Errorf("droplet.FromString failed: %v", err)
		}

		b := addrBalances[o.Address]
		b.expected.Coins += amt
		b.expected.Hours += o.CalculatedHours

		addrBalances[o.Address] = b
	}

	toBalance := func(b wallet.Balance) (cli.Balance, error) {
		coins, err := droplet.ToString(b.Coins)
		if err != nil {
			return cli.Balance{}, err
		}

		return cli.Balance{
			Coins: coins,
			Hours: strconv.FormatUint(b.Hours, 10),
		}, nil
	}

	var totalConfirmed, totalSpendable, totalExpected wallet.Balance
	balRlt := &cli.BalanceResult{
		Addresses: make([]cli.AddressBalances, len(addrs)),
	}

	for i, a := range addrs {
		b := addrBalances[a]
		var err error

		balRlt.Addresses[i].Address = a

		totalConfirmed, err = totalConfirmed.Add(b.confirmed)
		if err != nil {
			return nil, err
		}

		totalSpendable, err = totalSpendable.Add(b.spendable)
		if err != nil {
			return nil, err
		}

		totalExpected, err = totalExpected.Add(b.expected)
		if err != nil {
			return nil, err
		}

		balRlt.Addresses[i].Confirmed, err = toBalance(b.confirmed)
		if err != nil {
			return nil, err
		}

		balRlt.Addresses[i].Spendable, err = toBalance(b.spendable)
		if err != nil {
			return nil, err
		}

		balRlt.Addresses[i].Expected, err = toBalance(b.expected)
		if err != nil {
			return nil, err
		}
	}

	var err error
	balRlt.Confirmed, err = toBalance(totalConfirmed)
	if err != nil {
		return nil, err
	}

	balRlt.Spendable, err = toBalance(totalSpendable)
	if err != nil {
		return nil, err
	}

	balRlt.Expected, err = toBalance(totalExpected)
	if err != nil {
		return nil, err
	}

	return balRlt, nil
}
