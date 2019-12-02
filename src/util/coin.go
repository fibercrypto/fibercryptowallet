package util

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/errors"
	"github.com/sirupsen/logrus"
	stdErr "errors"
)

func NewGenericOutput(addr core.Address) GenericOutput {
	return GenericOutput{
		Address: addr,
		Balance: make(map[string]uint64),
	}
}

// GenericOutput is a transient editable transaction output
type GenericOutput struct {
	Address core.Address
	Balance map[string]uint64
}

// GetId provides transaction output ID
func (gOut *GenericOutput) GetId() string {
	return ""
}

// IsSpent determines whether there exists a confirmed transaction with an input spending this output
func (gOut *GenericOutput) IsSpent() bool {
	return false
}

// GetAddress returns the address of the party receiving funds
func (gOut *GenericOutput) GetAddress() core.Address {
	return gOut.Address
}

// SetCoins allocates coin for an asset given its ticker
func (gOut *GenericOutput) SetCoins(ticker string, coins uint64) {
	gOut.Balance[ticker] = coins
}

// PushCoins parses coins string and allocates them for an asset given its ticker
func (gOut *GenericOutput) PushCoins(ticker string, coinStr string) error {
	coins, err := GetCoinValue(coinStr, ticker)
	if err != nil {
		return err
	}
	gOut.Balance[ticker] = coins
	return nil
}

// GetCoins looks up coins for asset represented by ticker that have been transferred in this output
func (gOut *GenericOutput) GetCoins(ticker string) (uint64, error) {
	if coins, hasCoins := gOut.Balance[ticker]; hasCoins {
		return coins, nil
	}
	return 0, errors.ErrInvalidAltcoinTicker
}

// SupportedAssets enumerates tickers of crypto assets supported by this output
func (gOut *GenericOutput) SupportedAssets() (tickers []string) {
	for t := range gOut.Balance {
		tickers = append(tickers, t)
	}
	return
}

func (gOut *GenericOutput) Clone() (interface{}, error) {
	if gOut == nil {
		return gOut, nil
	}
	addr, err := gOut.Address.Clone()
	if err != nil {
		// TODO i18n
		logrus.WithError(err).Errorln("unable to clone Address")
		return nil, errors.ErrDeepCopyFailed
	}
	addrTyped, ok := addr.(core.Address)
	if !ok {
		// TODO i18n
		err = stdErr.New("can not get address from cloned object")
		logrus.Errorln(err.Error())
		return nil, err
	}
	balance := make(map[string]uint64, len(gOut.Balance))
	for k, v := range gOut.Balance {
		balance[k] = v
	}
	newOut := &GenericOutput{
		Address: addrTyped,
		Balance: balance,
	}
	return newOut, nil
}

// Type assertions
var (
	_ core.TransactionOutput = &GenericOutput{}
)
