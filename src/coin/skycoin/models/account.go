package skycoin //nolint goimports

import (
	"path/filepath"
	"strconv"

	"github.com/SkycoinProject/skycoin/src/cli"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/SkycoinProject/skycoin/src/util/droplet"
	"github.com/SkycoinProject/skycoin/src/wallet"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var log = logging.MustGetLogger("Skycoin Account")

func (addr *SkycoinAddress) GetBalance(ticker string) (uint64, error) {

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		log.WithError(err).Error("Couldn't get API client")
		return 0, err
	}
	defer ReturnSkycoinClient(c)
	log.Info("POST /api/v1/balance?addrs=xxx")
	bl, err := c.Balance([]string{addr.address})
	if err != nil {
		log.WithError(err).WithField("addrs", "addr.address").Error("Couldn't POST /api/v1/balance?addrs=xxx")
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
func (addr *SkycoinAddress) ListAssets() []string {
	return []string{Sky, CoinHour}
}
func (addr *SkycoinAddress) ScanUnspentOutputs() core.TransactionOutputIterator {
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		log.WithError(err).Error("Couldn't get API client")
		return nil
	}
	defer ReturnSkycoinClient(c)
	log.Info("POST /api/v1/outputs?addrs=xxx")
	outputSummary, err := c.OutputsForAddresses([]string{addr.String()})
	if err != nil {
		log.WithError(err).WithField("addrs", addr.String()).Error("Couldn't POST /api/v1/outputs?addrs=xxx")
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
			spent:           true,
			calculatedHours: out.CalculatedHours,
		})
	}

	return NewSkycoinTransactionOutputIterator(skyOutputs)
}
func (addr *SkycoinAddress) ListTransactions() core.TransactionIterator {

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		log.WithError(err).Error("Couldn't get API client")
		return nil
	}
	defer ReturnSkycoinClient(c)
	transactions := make([]core.Transaction, 0)
	txn, err := c.TransactionsVerbose([]string{addr.String()})
	if err != nil {
		log.WithError(err).WithField("addrs", addr.String()).Error("Couldn't POST /api/v1/transactions?verbose=1")
		return nil
	}

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
func (addr *SkycoinAddress) ListPendingTransactions() (core.TransactionIterator, error) { //------TODO
	return nil, nil
}

func (wlt *RemoteWallet) GetBalance(ticker string) (uint64, error) {
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		log.WithError(err).Error("Couldn't get API client")
		return 0, err
	}
	defer ReturnSkycoinClient(c)
	log.Info("GET /api/v1/wallet/balance")
	bl, err := c.WalletBalance(wlt.Id)
	if err != nil {
		log.WithError(err).WithField("id", wlt.Id).Error("Couldn't GET /api/v1/wallet/balance")
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

func (wlt *RemoteWallet) ListAssets() []string {
	return []string{Sky, CoinHour}
}

func (wlt *RemoteWallet) ScanUnspentOutputs() core.TransactionOutputIterator {
	log.Info("Calling RemoteWallet.GetLoadedAddresses()")
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		log.WithError(err).Error("RemoteWallet.GetLoadedAddresses() failed")
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

func (wlt *RemoteWallet) ListTransactions() core.TransactionIterator {
	log.Info("Calling RemoteWallet.GetLoadedAddresses()")
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		log.WithError(err).Error("RemoteWallet.GetLoadedAddresses() failed")
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

func (wlt *RemoteWallet) ListPendingTransactions() (core.TransactionIterator, error) {
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		log.WithError(err).Error("Couldn't get API client")
		return nil, err
	}
	defer ReturnSkycoinClient(c)
	log.Info("GET /api/v1/wallet/transactions&verbose=1")
	response, err2 := c.WalletUnconfirmedTransactionsVerbose(wlt.GetId())
	if err2 != nil {
		log.WithError(err).WithField("id", wlt.GetId()).Error("Couldn't GET /api/v1/wallet/transactions&verbose=1")
		return nil, err2
	}
	txns := make([]core.Transaction, 0)
	for _, ut := range response.Transactions {
		txns = append(txns, &SkycoinPendingTransaction{Transaction: ut})
	}
	return NewSkycoinTransactionIterator(txns), nil
}

func (wlt *LocalWallet) GetBalance(ticker string) (uint64, error) {
	walletName := filepath.Join(wlt.WalletDir, wlt.Id)
	log.WithField("walletName", walletName).Info("Calling wallet.Load(walletName)")
	walletLoaded, err := wallet.Load(walletName)
	if err != nil {
		log.WithError(err).Error("wallet.Load(walletName) failed")
		return 0, err
	}
	var addrs []string
	addresses := walletLoaded.GetAddresses()
	for _, addr := range addresses {
		addrs = append(addrs, addr.String())
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		log.WithError(err).Error("Couldn't get API client")
		return 0, err
	}
	defer ReturnSkycoinClient(c)
	log.Info("POST /api/v1/outputs?addrs=xxx")
	outs, err := c.OutputsForAddresses(addrs)
	if err != nil {
		log.WithError(err).WithField("Length of addrs", len(addrs)).Error("Couldn't POST /api/v1/outputs?addrs=xxx")
		return 0, err
	}

	bl, err := getBalanceOfAddresses(outs, addrs)
	if err != nil {
		log.WithError(err).Warn("getBalanceOfAddresses(outs, addrs) failed")
		return 0, err
	}

	if ticker == Sky {
		flSky, err := strconv.ParseFloat(bl.Confirmed.Coins, 64)
		if err != nil {
			log.WithError(err).WithField("bl.Confirmed.Coins", bl.Confirmed.Coins).Error("strconv.ParseFloat(bl.Confirmed.Coins, 64) failed")
			return 0, err
		}
		accuracy, err2 := util.AltcoinQuotient(Sky)
		if err2 != nil {
			log.WithError(err2).WithField("Sky", Sky).Error("util.AltcoinQuotient(Sky) failed")
			return 0, err2
		}
		return uint64(flSky * float64(accuracy)), nil
	} else if ticker == CoinHour {
		coinHours, err := strconv.ParseFloat(bl.Confirmed.Hours, 64)
		if err != nil {
			log.WithError(err).WithField("bl.Confirmed.Hours", bl.Confirmed.Hours).Error("strconv.ParseFloat(bl.Confirmed.Hours, 64) failed")
			return 0, err
		}
		accuracy, err2 := util.AltcoinQuotient(CoinHour)
		if err2 != nil {
			log.WithError(err2).WithField("CoinHour", CoinHour).Error("util.AltcoinQuotient(CoinHour) failed")
			return 0, err2
		}
		return uint64(coinHours * float64(accuracy)), nil
	} else {
		return 0, errorTickerInvalid{ticker}
	}

}

func (wlt *LocalWallet) ListAssets() []string {
	return []string{Sky, CoinHour}
}

func (wlt *LocalWallet) ScanUnspentOutputs() core.TransactionOutputIterator {
	log.Info("Calling LocalWallet.GetLoadedAddresses()")
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		log.WithError(err).Error("LocalWallet.GetLoadedAddresses() failed")
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

func (wlt *LocalWallet) ListTransactions() core.TransactionIterator {
	addressesIter, err := wlt.GetLoadedAddresses()
	if err != nil {
		log.WithError(err).Error("LocalWallet.GetLoadedAddresses() failed")
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

func (wlt *LocalWallet) ListPendingTransactions() (core.TransactionIterator, error) { //------TODO
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		log.WithError(err).Error("Couldn't get API client")
		return nil, err
	}
	defer ReturnSkycoinClient(c)
	log.Info("GET /api/v1/wallet/transactions&verbose=1")
	response, err2 := c.WalletUnconfirmedTransactionsVerbose(wlt.GetId())
	if err2 != nil {
		log.WithError(err2).WithField("id", wlt.GetId()).Error("Couldn't GET /api/v1/wallet/transactions&verbose=1")
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
			log.Errorf("Found address %s in GetUnspentOutputs result, but this address wasn't requested", o.Address)
			return nil, errors.ErrUnexpectedUxOutAddress
		}

		amt, err := droplet.FromString(o.Coins)
		if err != nil {
			log.WithError(err).Error("droplet.FromString failed")
			return nil, errors.ErrParseTxnCoins
		}

		b := addrBalances[o.Address]
		b.confirmed.Coins += amt
		b.confirmed.Hours += o.CalculatedHours

		addrBalances[o.Address] = b
	}

	// Count spendable balances
	for _, o := range outs.SpendableOutputs() {
		if _, ok := addrsMap[o.Address]; !ok {
			log.Errorf("Found address %s in GetUnspentOutputs result, but this address wasn't requested", o.Address)
			return nil, errors.ErrUnexpectedUxOutAddress
		}

		amt, err := droplet.FromString(o.Coins)
		if err != nil {
			log.WithError(err).Error("droplet.FromString failed")
			return nil, errors.ErrParseTxnCoins
		}

		b := addrBalances[o.Address]
		b.spendable.Coins += amt
		b.spendable.Hours += o.CalculatedHours

		addrBalances[o.Address] = b
	}

	// Count predicted balances
	for _, o := range outs.ExpectedOutputs() {
		if _, ok := addrsMap[o.Address]; !ok {
			log.Errorf("Found address %s in GetUnspentOutputs result, but this address wasn't requested", o.Address)
			return nil, errors.ErrUnexpectedUxOutAddress
		}

		amt, err := droplet.FromString(o.Coins)
		if err != nil {
			log.WithError(err).Error("droplet.FromString failed")
			return nil, errors.ErrParseTxnCoins
		}

		b := addrBalances[o.Address]
		b.expected.Coins += amt
		b.expected.Hours += o.CalculatedHours

		addrBalances[o.Address] = b
	}

	toBalance := func(b wallet.Balance) (cli.Balance, error) {
		coins, err := droplet.ToString(b.Coins)
		if err != nil {
			log.WithError(err).Error()
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
			log.WithError(err).Error()
			return nil, err
		}

		totalSpendable, err = totalSpendable.Add(b.spendable)
		if err != nil {
			log.WithError(err).Error()
			return nil, err
		}

		totalExpected, err = totalExpected.Add(b.expected)
		if err != nil {
			log.WithError(err).Error()
			return nil, err
		}

		balRlt.Addresses[i].Confirmed, err = toBalance(b.confirmed)
		if err != nil {
			log.WithError(err).Error()
			return nil, err
		}

		balRlt.Addresses[i].Spendable, err = toBalance(b.spendable)
		if err != nil {
			log.WithError(err).Error()
			return nil, err
		}

		balRlt.Addresses[i].Expected, err = toBalance(b.expected)
		if err != nil {
			log.WithError(err).Error()
			return nil, err
		}
	}

	var err error
	balRlt.Confirmed, err = toBalance(totalConfirmed)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	balRlt.Spendable, err = toBalance(totalSpendable)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	balRlt.Expected, err = toBalance(totalExpected)
	if err != nil {
		log.WithError(err).Error()
		return nil, err
	}

	return balRlt, nil
}
