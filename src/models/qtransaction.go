package models

import (
	"fmt"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/models/address"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	qtcore "github.com/therecipe/qt/core"
)

type QTransaction struct {
	qtcore.QObject
	txn *core.Transaction
	_   string               `property:"amount"`
	_   string               `property:"hoursTraspassed"`
	_   string               `property:"hoursBurned"`
	_   string               `property:"transactionId"`
	_   *address.AddressList `property:"addresses"`
	_   *address.AddressList `property:"inputs"`
	_   *address.AddressList `property:"outputs"`
}

func NewQTransactionFromTransaction(txn core.Transaction) (*QTransaction, error) {
	fmt.Println(3.1)
	qtxn := NewQTransaction(nil)
	qtxn.SetTransactionId(txn.GetId())
	//addresses := address.NewAddressList(nil)
	inputs := address.NewAddressList(nil)
	outputs := address.NewAddressList(nil)
	var hoursTraspassed uint64
	var skyTraspassed uint64
	inputsAddresses := make(map[string]struct{}, 0)
	quotient, err := util.AltcoinQuotient("SKYCH")
	if err != nil {
		return nil, err
	}
	fmt.Println(3.3)
	fee := util.FormatCoins(txn.ComputeFee("SKYCH"), quotient)
	qtxn.SetHoursBurned(fee)

	//Creating inputs
	ins := txn.GetInputs()
	fmt.Println("MyInputs")
	fmt.Println(len(ins))
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
		ch, err := in.GetCoins("SKYCH")
		if err != nil {
			return nil, err
		}
		qIn.SetAddressCoinHours(util.FormatCoins(ch, quotient))
		inputs.AddAddress(qIn)
	}
	qtxn.SetInputs(inputs)
	fmt.Println(3.4)
	//Creating Outputs
	outs := txn.GetOutputs()
	fmt.Println("OUTS")
	fmt.Println(len(outs))
	for _, out := range outs {
		qOu := address.NewAddressDetails(nil)
		addr := out.GetAddress().String()
		qOu.SetAddress(addr)
		quotient, err := util.AltcoinQuotient("SKY")
		sky, err := out.GetCoins("SKY")
		fmt.Println("sky ")
		fmt.Println(sky)
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
	fmt.Println(3.5)
	qtxn.SetOutputs(outputs)
	qtxn.SetHoursTraspassed(util.FormatCoins(hoursTraspassed, quotient))
	fmt.Println(hoursTraspassed)
	fmt.Println(3.6)
	quotient, err = util.AltcoinQuotient("SKY")
	if err != nil {
		return nil, err
	}
	fmt.Println(3.7)
	fmt.Println(skyTraspassed)
	strSky := util.FormatCoins(skyTraspassed, quotient)
	fmt.Println(3.8)
	qtxn.SetAmount(strSky)

	return qtxn, nil
}
