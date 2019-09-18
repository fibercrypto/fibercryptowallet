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
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
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
func (addr SkycoinAddress) ScanUnspentOutputs() core.TransactionOutputIterator {
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		println(err.Error())
		return nil
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	outputSummary, err := c.OutputsForAddresses([]string{addr.String()})
	if err != nil {
		println(err.Error())
		return nil
	}

	outs := outputSummary.SpendableOutputs()
	skyOutputs := make([]core.TransactionOutput, 0)
	for _, out := range outs {
		skyOutputs = append(skyOutputs, &SkycoinTransactionOutput{

			skyOut: readable.TransactionOutput{
				Address: out.Address,
				Coins:   out.Coins,
				Hours:   out.Hours,
				Hash:    out.Hash,
			},
			spent: true,
			calculatedHours: out.CalculatedHours,
		})
	}

	return NewSkycoinTransactionOutputIterator(skyOutputs)
}
func (addr SkycoinAddress) ListTransactions() core.TransactionIterator {

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	transactions := make([]core.Transaction, 0)
	txn, _ := c.TransactionsVerbose([]string{addr.String()})

	for _, tx := range txn {
		st := core.TXN_STATUS_PENDING
		if tx.Status.Confirmed {
			st = core.TXN_STATUS_CONFIRMED
		}

		transactions = append(transactions, &SkycoinTransaction{
			skyTxn: tx.Transaction,
			status: st,
		})

	}

	return NewSkycoinTransactionIterator(transactions)

}
func (addr SkycoinAddress) ListPendingTransactions() (core.TransactionIterator, error) { //------TODO
	return nil,nil
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

func (wlt RemoteWallet) ScanUnspentOutputs() core.TransactionOutputIterator {
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		return nil
	}
	unOuts := make([]core.TransactionOutput, 0)
	for addressesIter.Next() {
		outsIter := addressesIter.Value().GetCryptoAccount().ScanUnspentOutputs()
		for outsIter.Next() {
			unOuts = append(unOuts, outsIter.Value())
		}
	}
	return NewSkycoinTransactionOutputIterator(unOuts)
}

func (wlt RemoteWallet) ListTransactions() core.TransactionIterator {
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		return nil
	}
	txns := make([]core.Transaction, 0)
	for addressesIter.Next() {
		txnsIter := addressesIter.Value().GetCryptoAccount().ListTransactions()
		for txnsIter.Next() {
			txns = append(txns, txnsIter.Value())
		}
	}

	return NewSkycoinTransactionIterator(txns)
}

func (wlt RemoteWallet) ListPendingTransactions() (core.TransactionIterator, error) { 
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	response, err2 := c.WalletUnconfirmedTransactionsVerbose(wlt.GetId())
	if err2 != nil {
		return nil, err2
	}
	txns := make([]core.Transaction, 0)
	for _, ut := range response.Transactions {
		txns = append(txns, &SkycoinPendingTransaction{Transaction: ut})
	}
	return NewSkycoinTransactionIterator(txns), nil
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
		skyf, err := strconv.ParseFloat(bl.Confirmed.Coins, 64)
		if err != nil {
			return 0, err
		}
		accuracy, err2 := util.AltcoinQuotient(Sky)
		if err2 != nil {
			return 0, err2
		}
		return uint64(skyf * float64(accuracy)), nil
	} else if ticker == CoinHour {
		coinHours, err := strconv.ParseFloat(bl.Confirmed.Hours, 64)
		if err != nil {
			return 0, err
		}
		accuracy, err2 := util.AltcoinQuotient(CoinHour)
		if err2 != nil {
			return 0, err2
		}
		return uint64(coinHours * float64(accuracy)), nil
	} else {
		return 0, errorTickerInvalid{ticker}
	}

}

func (wlt LocalWallet) ListAssets() []string {
	return []string{Sky, CoinHour}
}

func (wlt LocalWallet) ScanUnspentOutputs() core.TransactionOutputIterator {
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		return nil
	}
	unOuts := make([]core.TransactionOutput, 0)
	for addressesIter.Next() {
		outsIter := addressesIter.Value().GetCryptoAccount().ScanUnspentOutputs()
		for outsIter.Next() {
			unOuts = append(unOuts, outsIter.Value())
		}
	}
	return NewSkycoinTransactionOutputIterator(unOuts)
}

func (wlt LocalWallet) ListTransactions() core.TransactionIterator {
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		return nil
	}
	txns := make([]core.Transaction, 0)
	for addressesIter.Next() {
		txnsIter := addressesIter.Value().GetCryptoAccount().ListTransactions()
		for txnsIter.Next() {
			txns = append(txns, txnsIter.Value())
		}
	}

	return NewSkycoinTransactionIterator(txns)
}

func (wlt LocalWallet) ListPendingTransactions() (core.TransactionIterator, error) { //------TODO
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	response, err2 := c.WalletUnconfirmedTransactionsVerbose(wlt.GetId())
	if err2 != nil {
		return nil, err2
	}
	txns := make([]core.Transaction, 0)
	for _, ut := range response.Transactions {
		txns = append(txns, &SkycoinPendingTransaction{Transaction: ut})
	}
	return NewSkycoinTransactionIterator(txns), nil
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
