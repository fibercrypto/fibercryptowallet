package skycoin

import (
	"strconv"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/readable"
)

/*
SkycoinTransaction
*/
type SkycoinTransaction struct {
	skyTxn  readable.TransactionVerbose
	status  core.TransactionStatus
	inputs  []core.TransactionInput
	outputs []core.TransactionOutput
}

func (txn *SkycoinTransaction) SupportedAssets() []string {
	return []string{Sky, CoinHour}
}

func (txn *SkycoinTransaction) GetTimestamp() core.TransactionTimestamp {
	return core.TransactionTimestamp(txn.skyTxn.Timestamp)
}

func (txn *SkycoinTransaction) GetStatus() core.TransactionStatus {

	if txn.status == core.TXN_STATUS_CONFIRMED {
		return txn.status
	}

	c := util.NewClient()
	txnU, _ := c.Transaction(txn.skyTxn.Hash)
	if txnU.Status.Confirmed {
		txn.status = core.TXN_STATUS_CONFIRMED
		return txn.status
	}
	txn.status = core.TXN_STATUS_PENDING
	return txn.status
}

func (txn *SkycoinTransaction) GetInputs() []core.TransactionInput {
	if txn.inputs == nil {
		c := util.NewClient()
		transaction, err := c.TransactionVerbose(txn.skyTxn.Hash)
		if err != nil {
			return nil
		}
		txn.inputs = make([]core.TransactionInput, 0)
		for _, in := range transaction.Transaction.In {
			txn.inputs = append(txn.inputs, SkycoinTransactionInput{
				skyIn:       in,
				spentOutput: nil,
			})
		}

	}

	return txn.inputs
}

func (txn *SkycoinTransaction) GetOutputs() []core.TransactionOutput {
	if txn.outputs == nil {
		txn.outputs = make([]core.TransactionOutput, 0)
		for _, out := range txn.skyTxn.Out {
			txn.outputs = append(txn.outputs, SkycoinTransactionOutput{
				skyOut: out,
				spent:  false,
			})
		}
	}
	return txn.outputs
}

func (txn *SkycoinTransaction) GetId() string {
	return txn.skyTxn.Hash
}

func (txn *SkycoinTransaction) ComputeFee(ticker string) uint64 {
	if ticker == CoinHour {
		return txn.skyTxn.Fee
	}
	return 0
}

/**
 * SkycoinTransactionIterator
 */
type SkycoinTransactionIterator struct {
	current int
	data    []core.Transaction
}

func (iter *SkycoinTransactionIterator) Value() core.Transaction {
	return iter.data[iter.current]
}

func (iter *SkycoinTransactionIterator) Next() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func (iter *SkycoinTransactionIterator) HasNext() bool {
	if (iter.current + 1) >= len(iter.data) {
		return false
	}
	return true
}

func NewSkycoinTransactionIterator(txns []core.Transaction) *SkycoinTransactionIterator {
	return &SkycoinTransactionIterator{data: txns, current: -1}
}

type SkycoinTransactionInput struct {
	skyIn       readable.TransactionInput
	spentOutput *SkycoinTransactionOutput
}

func (in SkycoinTransactionInput) GetId() string {
	return in.skyIn.Hash
}

func (in SkycoinTransactionInput) GetSpentOutput() core.TransactionOutput {
	if in.spentOutput == nil {
		c := util.NewClient()

		out, err := c.UxOut(in.skyIn.Hash)
		if err != nil {
			return nil
		}

		skyOut := &SkycoinTransactionOutput{
			skyOut: readable.TransactionOutput{
				Address: out.OwnerAddress,
				Coins:   strconv.FormatFloat(float64(out.Coins/1e6), 'f', -1, 64),
				Hours:   out.Hours,
				Hash:    out.Uxid,
			},
			spent: true}
		in.spentOutput = skyOut

	}
	return in.spentOutput

}

func (in SkycoinTransactionInput) GetCoins(ticker string) uint64 {
	if ticker == Sky {
		skyF, _ := strconv.ParseFloat(in.skyIn.Coins, 64)
		return uint64(skyF * 1e6)
	}
	if ticker == CoinHour {
		return in.skyIn.CalculatedHours
	}
	return 0
}

/**
 * SkycoinTransactionInputIterator
 */
type SkycoinTransactionInputIterator struct {
	current int
	data    []*SkycoinTransactionInput
}

func (iter *SkycoinTransactionInputIterator) Value() core.TransactionInput {
	return iter.data[iter.current]
}

func (iter *SkycoinTransactionInputIterator) Next() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func (iter *SkycoinTransactionInputIterator) HasNext() bool {
	if (iter.current + 1) >= len(iter.data) {
		return false
	}
	return true
}

func NewSkycoinTransactioninputIterator(ins []*SkycoinTransactionInput) *SkycoinTransactionInputIterator {
	return &SkycoinTransactionInputIterator{data: ins, current: -1}
}

/**
 * SkycoinTransactionOutput
 */
type SkycoinTransactionOutput struct {
	skyOut readable.TransactionOutput
	spent  bool
}

func (out SkycoinTransactionOutput) GetId() string {
	return out.skyOut.Hash

}

func (out SkycoinTransactionOutput) GetAddress() core.Address {
	return SkycoinAddress{out.skyOut.Address}

}

func (out SkycoinTransactionOutput) GetCoins(ticker string) uint64 {
	if ticker == Sky {
		skyF, _ := strconv.ParseFloat(out.skyOut.Coins, 64)
		return uint64(skyF * 1e6)
	}
	if ticker == CoinHour {
		return out.skyOut.Hours
	}
	return 0
}

func (out SkycoinTransactionOutput) IsSpent() bool {
	if out.spent {
		return true
	}

	c := util.NewClient()
	ou, err := c.UxOut(out.skyOut.Hash)
	if err != nil {
		return false
	}
	if ou.SpentTxnID != "0000000000000000000000000000000000000000000000000000000000000000" {
		out.spent = true
		return true
	}
	return false
}

/**
 * SkycoinTransactionOutputIterator
 */
type SkycoinTransactionOutputIterator struct {
	current int
	data    []core.TransactionOutput
}

func (iter *SkycoinTransactionOutputIterator) Value() core.TransactionOutput {
	return iter.data[iter.current]
}

func (iter *SkycoinTransactionOutputIterator) Next() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func (iter *SkycoinTransactionOutputIterator) HasNext() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func NewSkycoinTransactionOutputIterator(outs []core.TransactionOutput) *SkycoinTransactionOutputIterator {
	return &SkycoinTransactionOutputIterator{data: outs, current: -1}
}
