package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	qtcore "github.com/therecipe/qt/core"
)

type QTransaction struct {
	qtcore.QObject
	txn core.Transaction
	_   string               `property:"amount"`
	_   string               `property:"hoursTraspassed"`
	_   string               `property:"hoursBurned"`
	_   string               `property:"transactionId"`
	_   *address.AddressList `property:"inputs"`
	_   *address.AddressList `property:"outputs"`
}

func NewQTransactionFromTransaction(txn core.Transaction) (*QTransaction, error) {
	qtxn := NewQTransaction(nil)
	qtxn.txn = txn
	qtxn.SetTransactionId(txn.GetId())
	inputs := address.NewAddressList(nil)
	outputs := address.NewAddressList(nil)
	var hoursTraspassed uint64
	var skyTraspassed uint64
	inputsAddresses := make(map[string]struct{}, 0)
	quotient, err := util.AltcoinQuotient("SKYCH")
	if err != nil {
		return nil, err
	}
	ch, err := txn.ComputeFee("SKYCH")
	if err != nil {
		return nil, nil
	}
	fee := util.FormatCoins(ch, quotient)
	qtxn.SetHoursBurned(fee)

	//Creating inputs
	ins := txn.GetInputs()
	for _, in := range ins {
		qIn := address.NewAddressDetails(nil)
		addr := in.GetSpentOutput().GetAddress().String()
		inputsAddresses[addr] = struct{}{}
		qIn.SetAddress(addr)
		quotient, err := util.AltcoinQuotient("SKY")
		if err != nil {
			return nil, err
		}
		sky, err := in.GetCoins("SKY")
		if err != nil {
			return nil, err
		}
		qIn.SetAddressSky(util.FormatCoins(sky, quotient))
		quotient, err = util.AltcoinQuotient("SKYCH")
		if err != nil {
			return nil, err
		}
		ch, err := in.GetCoins("SKYCHC")
		if err != nil {
			return nil, err
		}
		qIn.SetAddressCoinHours(util.FormatCoins(ch, quotient))
		inputs.AddAddress(qIn)
	}
	qtxn.SetInputs(inputs)

	//Creating Outputs
	outs := txn.GetOutputs()
	for _, out := range outs {
		qOu := address.NewAddressDetails(nil)
		addr := out.GetAddress().String()
		qOu.SetAddress(addr)
		quotient, err := util.AltcoinQuotient("SKY")
		sky, err := out.GetCoins("SKY")
		if err != nil {
			return nil, err
		}
		qOu.SetAddressSky(util.FormatCoins(sky, quotient))
		quotient, err = util.AltcoinQuotient("SKYCH")
		if err != nil {
			return nil, err
		}
		ch, err := out.GetCoins("SKYCH")
		if err != nil {
			return nil, err
		}
		qOu.SetAddressCoinHours(util.FormatCoins(ch, quotient))
		outputs.AddAddress(qOu)
		_, ok := inputsAddresses[addr]
		if !ok {
			skyTraspassed += sky
			hoursTraspassed += ch
		}
	}
	qtxn.SetOutputs(outputs)
	qtxn.SetHoursTraspassed(util.FormatCoins(hoursTraspassed, quotient))
	quotient, err = util.AltcoinQuotient("SKY")
	if err != nil {
		return nil, err
	}
	strSky := util.FormatCoins(skyTraspassed, quotient)
	qtxn.SetAmount(strSky)

	return qtxn, nil
}
