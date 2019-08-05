package skycoin 
 
import ( 
  	"github.com/fibercrypto/FiberCryptoWallet/src/core" 
) 
 
/* 
SkycoinTransaction 
*/ 
type SkycoinTransaction struct{ //Implements Transaction interface
	Timestamp core.TransactionTimestamp
	Status    core.TransactionStatus
	Inputs    []core.TransactionInput
	Outputs   []core.TransactionOutput
	Id        string
} 
 
func (txn *SkycoinTransaction) SupportedAssets() []string { 
  	return []string{"SKY", "SKYCH"} 
} 
 
func (txn *SkycoinTransaction) GetTimestamp() core.TransactionTimestamp { 
  	return txn.Timestamp 
} 
 
func (txn *SkycoinTransaction) GetStatus() core.TransactionStatus { 
  	return txn.Status
} 
 
func (txn *SkycoinTransaction) GetInputs() []core.TransactionInput { 
	return txn.Inputs
} 
 
func (txn *SkycoinTransaction) GetOutputs() []core.TransactionOutput { 
 	return txn.Outputs
} 
 
func (txn *SkycoinTransaction) GetId() string { 
  	return txn.Id
} 
 
func (txn *SkycoinTransaction) ComputeFee(ticker string) uint64 { 
	iter := NewSkycoinTransactionOutputIterator(txn.Outputs)
	amount := uint64(0)
	for iter.Next() {
		amount = amount + iter.Value().GetCoins(ticker)
	}
	return amount
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
	Id        string 
	Address   SkycoinAddress 
	Sky       uint64 
	CoinHours uint64 
} 

func (sto *SkycoinTransactionOutput) GetId() string { 
	return sto.Id 
} 

func (sto *SkycoinTransactionOutput) GetAddress() core.Address { 
	return sto.Address 
} 

func (sto *SkycoinTransactionOutput) GetCoins(ticker string) uint64 { 
	if ticker == "SKY" { 
		return sto.Sky 
	} 
	return sto.CoinHours 
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

