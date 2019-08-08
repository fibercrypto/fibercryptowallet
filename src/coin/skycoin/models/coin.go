package skycoin 
 
import (
	"strconv" 
	"github.com/skycoin/skycoin/src/readable"
  	"github.com/fibercrypto/FiberCryptoWallet/src/core" 
) 
 
/* 
SkycoinPendingTransaction 
*/ 
type SkycoinPendingTransaction struct{ //Implements Transaction interface
	Transaction readable.UnconfirmedTransactionVerbose
	// Timestamp core.TransactionTimestamp
	// Status    core.TransactionStatus
	// Inputs    []core.TransactionInput
	// Outputs   []core.TransactionOutput
	// Id        string
	// Fee       uint64
} 
 
func (txn *SkycoinPendingTransaction) SupportedAssets() []string { 
  	return []string{Sky, CoinHour} 
} 
 
func (txn *SkycoinPendingTransaction) GetTimestamp() core.TransactionTimestamp { 
  	return core.TransactionTimestamp(txn.Transaction.Received.Unix())
} 
 
func (txn *SkycoinPendingTransaction) GetStatus() core.TransactionStatus { 
  	return core.TXN_STATUS_PENDING
} 
 
func (txn *SkycoinPendingTransaction) GetInputs() []core.TransactionInput { 
	inputs := make([]core.TransactionInput, 0)
	for _ , input := range txn.Transaction.Transaction.In {
		inputs = append(inputs, &SkycoinTransactionInput{Input: input})
	}
	return inputs
} 
 
func (txn *SkycoinPendingTransaction) GetOutputs() []core.TransactionOutput { 
	outputs := make([]core.TransactionOutput, 0)
	for _ , output := range txn.Transaction.Transaction.Out {
		outputs = append(outputs, &SkycoinTransactionOutput{Output: output})
	}
	return outputs
} 
 
func (txn *SkycoinPendingTransaction) GetId() string { 
  	return txn.Transaction.Transaction.Hash
} 
 
func (txn *SkycoinPendingTransaction) ComputeFee(ticker string) uint64 { 
	if ticker == Sky {
		return uint64(0);
	}
	return txn.Transaction.Transaction.Fee
} 
 
/** 
 * SkycoinTransactionIterator 
 */ 
type SkycoinTransactionIterator struct { //Implements TransactionIterator interface
	Current int
	Transactions []core.Transaction
} 
 
func (it *SkycoinTransactionIterator) Value() core.Transaction { 
	return it.Transactions[it.Current]
} 
 
func (it *SkycoinTransactionIterator) Next() bool { 
	if it.HasNext() {
		it.Current++
		return true
	}
	return false
} 
 
func (it *SkycoinTransactionIterator) HasNext() bool { 
	return (it.Current + 1) < len(it.Transactions)
} 
 
func NewSkycoinTransactionIterator(transactions []core.Transaction) *SkycoinTransactionIterator {
	return &SkycoinTransactionIterator{Transactions: transactions, Current: -1}
}

type SkycoinTransactionInput struct { //Implements TransactionInput interface
	Input readable.TransactionInput
} 
 
func (in *SkycoinTransactionInput) GetId() string { 
  	return "f1a61f49cef012e4822b314ca6657d66fdbe3c4d46110a079052a064b9a51e66" 
} 
 
func (in *SkycoinTransactionInput) IsSpent() bool { 
  	return true 
} 
 
func (in *SkycoinTransactionInput) GetSpentOutput() core.TransactionOutput { 
  	return &SkycoinTransactionOutput{} 
} 
 
/** 
 * SkycoinTransactionInputIterator 
 */ 
type SkycoinTransactionInputIterator struct { 
	index int 
	value *SkycoinTransactionInput 
} 
 
func (iter *SkycoinTransactionInputIterator) Value() core.TransactionInput { 
	if iter.index == 0 && iter.value == nil { 
		iter.value = &SkycoinTransactionInput{} 
	} 
	return iter.value 
} 
 
func (iter *SkycoinTransactionInputIterator) Next() bool { 
	if iter.index < 3 { 
		iter.value = &SkycoinTransactionInput{} 
		iter.index++ 
		return true 
	} 
	return false 
} 
 
func (iter *SkycoinTransactionInputIterator) HasNext() bool { 
  	return iter.index < 3 
} 
 
/** 
 * SkycoinTransactionOutput 
 */ 
type SkycoinTransactionOutput struct { //Implements TransactionOutput interface 
	Output readable.TransactionOutput
	// Id        string 
	// Address   SkycoinAddress 
	// Sky       uint64 
	// CoinHours uint64 
} 

func (sto *SkycoinTransactionOutput) GetId() string { 
	return sto.Output.Hash 
} 

func (sto *SkycoinTransactionOutput) IsSpent() bool { 
	//TODO:
	return false 
} 

func (sto *SkycoinTransactionOutput) GetAddress() core.Address { 
	return SkycoinAddress{address: sto.Output.Address}
} 

func (sto *SkycoinTransactionOutput) GetCoins(ticker string) uint64 { 
	if ticker == Sky { 
		coin, _ := strconv.ParseFloat(sto.Output.Coins, 64) 
		return uint64(coin * 1000000)
	} 
	return sto.Output.Hours 
} 

type SkycoinTransactionOutputIterator struct { //Implements TransactionOutputIterator interface 
	Current int 
	Outputs []core.TransactionOutput 
} 

func (it *SkycoinTransactionOutputIterator) Value() core.TransactionOutput { 
	return it.Outputs[it.Current] 
} 

func (it *SkycoinTransactionOutputIterator) Next() bool { 
	if it.HasNext() { 
		it.Current++ 
		return true 
	} 
	return false 
} 

func (it *SkycoinTransactionOutputIterator) HasNext() bool { 
	return (it.Current + 1) < len(it.Outputs) 
}

func NewSkycoinTransactionOutputIterator(outputs []core.TransactionOutput) *SkycoinTransactionOutputIterator {
	return &SkycoinTransactionOutputIterator{Outputs: outputs, Current: -1}
}

