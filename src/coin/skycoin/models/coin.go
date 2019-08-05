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
  if ticker == "SKYCH" { 
    return 1 
  } 
  return 2
} 
 
/** 
 * SkycoinTransactionIterator 
 */ 
type SkycoinTransactionIterator struct { //Implements TransactionIterator interface
	current int
	transactions []SkycoinTransaction
} 
 
func (it *SkycoinTransactionIterator) Value() core.Transaction { 
	return &it.transactions[it.current]
} 
 
func (it *SkycoinTransactionIterator) Next() bool { 
	if it.HasNext() {
		it.current++
		return true
	}
	return false
} 
 
func (it *SkycoinTransactionIterator) HasNext() bool { 
	return (it.current + 1) < len(it.transactions)
} 
 
func NewSkycoinTransactionIterator(transactions []SkycoinTransaction) *SkycoinTransactionIterator {
	return &SkycoinTransactionIterator{transactions: transactions, current: -1}
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
	 Coins float64
	 Hours uint64
	 Id    string
} 
 
func (sto *SkycoinTransactionOutput) GetId() string { 
  return "249iEtdhniFppBTpkzq7syoiBaydLi6USnU" 
} 
 
func (*SkycoinTransactionOutput) GetAddress() core.Address { 
  return &SkycoinAddress{} 
} 
 
func (*SkycoinTransactionOutput) GetCoins(ticker string) uint64 { 
  if ticker == "SKY" { 
    return 20000000 
  } 
  if ticker == "SKYCH" { 
    return 10000000 
  } 
  return 0 
} 
 
func (*SkycoinTransactionOutput) GetSpentInput() core.TransactionInput { 
  return &SkycoinTransactionInput{} 
} 
 
/** 
 * SkycoinTransactionOutputIterator 
 */ 
type SkycoinTransactionOutputIterator struct { 
  index int 
  value *SkycoinTransactionOutput 
} 
 
func (iter *SkycoinTransactionOutputIterator) Value() core.TransactionOutput { 
  if iter.index == 0 && iter.value == nil { 
    iter.value = &SkycoinTransactionOutput{} 
  } 
  return iter.value 
} 
 
func (iter *SkycoinTransactionOutputIterator) Next() bool { 
  if iter.index < 3 { 
    iter.value = &SkycoinTransactionOutput{} 
    iter.index++ 
    return true 
  } 
  return false 
} 
 
func (iter *SkycoinTransactionOutputIterator) HasNext() bool { 
  return iter.index < 3 
} 
