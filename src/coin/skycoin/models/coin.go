package skycoin

import (
	"fmt"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"strconv"
	"time"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/visor"
)

var logCoin = logging.MustGetLogger("Skycoin coin")

/*
	SkycoinPendingTransaction

	Implements Transaction interface
*/
type SkycoinPendingTransaction struct {
	Transaction readable.UnconfirmedTransactionVerbose
}

func (txn *SkycoinPendingTransaction) SupportedAssets() []string {
	logCoin.Info("Getting supported assets")
	return []string{Sky, CoinHour, CalculatedHour}
}

func (txn *SkycoinPendingTransaction) GetTimestamp() core.Timestamp {
	logCoin.Info("Getting timestamp")
	return core.Timestamp(txn.Transaction.Received.Unix())
}

func (txn *SkycoinPendingTransaction) GetStatus() core.TransactionStatus {
	logCoin.Info("Getting status")
	return core.TXN_STATUS_PENDING
}

func (txn *SkycoinPendingTransaction) GetInputs() []core.TransactionInput {
	logCoin.Info("Getting inputs from Skycoin pending transaction")
	inputs := make([]core.TransactionInput, 0)
	for _, input := range txn.Transaction.Transaction.In {
		inputs = append(inputs, &SkycoinPendingTransactionInput{Input: input})
	}
	return inputs
}

func (txn *SkycoinPendingTransaction) GetOutputs() []core.TransactionOutput {
	logCoin.Info("Getting outputs from Skycoin pending transaction")
	outputs := make([]core.TransactionOutput, 0)
	for _, output := range txn.Transaction.Transaction.Out {
		outputs = append(outputs, &SkycoinPendingTransactionOutput{Output: output})
	}
	return outputs
}

func (txn *SkycoinPendingTransaction) GetId() string {
	logCoin.Info("Getting id of pending transaction")
	return txn.Transaction.Transaction.Hash
}

func (txn *SkycoinPendingTransaction) ComputeFee(ticker string) (uint64, error) {
	logCoin.Info("Computing fee for " + ticker + " ticket")
	if ticker == CoinHour {
		return txn.Transaction.Transaction.Fee, nil
	} else if util.StringInList(ticker, txn.SupportedAssets()) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinTransactionIterator
 */
type SkycoinTransactionIterator struct { //Implements TransactionIterator interface
	Current      int
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

/**
 * SkycoinPendingTransactionInput
 */
type SkycoinPendingTransactionInput struct { //Implements TransactionInput interface
	Input readable.TransactionInput
}

func (in *SkycoinPendingTransactionInput) GetId() string {
	return ""
}

func (in *SkycoinPendingTransactionInput) IsSpent() bool {
	return true
}

func (in *SkycoinPendingTransactionInput) GetSpentOutput() core.TransactionOutput {
	return nil
}

func (in *SkycoinPendingTransactionInput) GetCoins(ticker string) (uint64, error) {
	logCoin.Info("Getting coins from " + ticker + " ticker")
	if util.StringInList(ticker, []string{Sky, CoinHour, CalculatedHour}) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinPendingTransactionOutput
 */
type SkycoinPendingTransactionOutput struct { //Implements TransactionOutput interface
	Output readable.TransactionOutput
}

func (sto *SkycoinPendingTransactionOutput) GetId() string {
	return sto.Output.Hash
}

func (sto *SkycoinPendingTransactionOutput) IsSpent() bool {
	return false
}

func (sto *SkycoinPendingTransactionOutput) GetAddress() core.Address {
	return &SkycoinAddress{address: sto.Output.Address}
}

func (sto *SkycoinPendingTransactionOutput) GetCoins(ticker string) (uint64, error) {
	logCoin.Info("Getting coins from " + ticker + " ticker")
	accuracy, err := util.AltcoinQuotient(ticker)
	if err != nil {
		return uint64(0), err
	}
	if ticker == Sky {
		coin, err2 := strconv.ParseFloat(sto.Output.Coins, 64)
		if err2 != nil {
			return uint64(0), err2
		}
		return uint64(coin * float64(accuracy)), nil
	} else if ticker == CoinHour {
		return sto.Output.Hours * accuracy, nil
	} else if ticker == CalculatedHour {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinTransactionOutputIterator
 */
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

func NewUninjectedTransaction(txn *coin.Transaction, fee uint64) (*SkycoinUninjectedTransaction, error) {
	return &SkycoinUninjectedTransaction{
		txn:     txn,
		inputs:  nil,
		outputs: nil,
		fee:     fee,
	}, nil
}

type SkycoinUninjectedTransaction struct {
	txn     *coin.Transaction
	inputs  []core.TransactionInput
	outputs []core.TransactionOutput
	fee     uint64
}

func (skyTxn *SkycoinUninjectedTransaction) SupportedAssets() []string {
	logCoin.Info("Getting supported assets from un injected transactions")
	return []string{Sky, CoinHour}
}

func (skyTxn *SkycoinUninjectedTransaction) GetTimestamp() core.Timestamp {
	logCoin.Info("Getting timestamp")
	return 0
}

func (skyTxn *SkycoinUninjectedTransaction) GetStatus() core.TransactionStatus {
	logCoin.Info("Getting status for un injected transaction")
	return core.TXN_STATUS_CREATED
}

func (skyTxn *SkycoinUninjectedTransaction) GetInputs() []core.TransactionInput {
	logCoin.Info("Getting inputs from un injected transactions")
	if skyTxn.inputs == nil {
		inputs, err := getSkycoinTransactionInputsFromInputsHashes(skyTxn.txn.In)
		if err != nil {
			return nil
		}
		skyTxn.inputs = inputs
	}
	return skyTxn.inputs
}

func (skyTxn *SkycoinUninjectedTransaction) GetOutputs() []core.TransactionOutput {
	logCoin.Info("Getting outputs from un injected transactions")
	if skyTxn.outputs == nil {
		outputs := make([]core.TransactionOutput, 0)
		for _, out := range skyTxn.txn.Out {
			rOut, err := readable.NewTransactionOutput(&out, skyTxn.txn.Hash())
			if err != nil {
				return nil
			}
			outputs = append(outputs, &SkycoinTransactionOutput{
				skyOut: *rOut,
				spent:  false,
			})
		}
		skyTxn.outputs = outputs
	}
	return skyTxn.outputs
}

func (skyTxn *SkycoinUninjectedTransaction) ComputeFee(ticker string) (uint64, error) {
	logCoin.Info("Computing fee for un injected transaction with" + ticker + " ticker")
	if ticker == CoinHour {
		return skyTxn.fee, nil
	}
	return 0, nil
}

func (skyTxn *SkycoinUninjectedTransaction) GetId() string {
	logCoin.Info("Getting if for un injected transaction")
	return skyTxn.txn.Hash().String()
}

/*
SkycoinTransaction
*/
type SkycoinTransaction struct {
	skyTxn readable.TransactionVerbose

	status  core.TransactionStatus
	inputs  []core.TransactionInput
	outputs []core.TransactionOutput
}

func (txn *SkycoinTransaction) SupportedAssets() []string {
	logCoin.Info("Getting supported assets from transactions")
	return []string{Sky, CoinHour, CalculatedHour}
}

func (txn *SkycoinTransaction) GetTimestamp() core.Timestamp {
	logCoin.Info("Getting timestamp transactions")
	return core.Timestamp(txn.skyTxn.Timestamp)
}

func (txn *SkycoinTransaction) GetStatus() core.TransactionStatus {
	logCoin.Info("Getting status for transactions")

	if txn.status == core.TXN_STATUS_CONFIRMED {
		return txn.status
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return 0
	}
	defer ReturnSkycoinClient(c)
	txnU, err := c.Transaction(txn.skyTxn.Hash)
	if err != nil {
		return 0
	}
	if txnU.Status.Confirmed {
		txn.status = core.TXN_STATUS_CONFIRMED
		return txn.status
	}
	txn.status = core.TXN_STATUS_PENDING
	return txn.status
}

func (txn *SkycoinTransaction) GetInputs() []core.TransactionInput {
	logCoin.Info("Getting inputs from transaction")

	if txn.inputs == nil {
		ins, err := getSkycoinTransactionInputsFromTxnHash(txn.skyTxn.Hash)
		if err != nil {
			return nil
		}
		txn.inputs = ins
	}
	return txn.inputs
}

func (txn *SkycoinTransaction) GetOutputs() []core.TransactionOutput {
	logCoin.Info("Getting outputs transactions")

	if txn.outputs == nil {
		txn.outputs = make([]core.TransactionOutput, 0)
		for _, out := range txn.skyTxn.Out {
			txn.outputs = append(txn.outputs, &SkycoinTransactionOutput{
				skyOut: out,
				spent:  false,
			})
		}
	}
	return txn.outputs
}

func (txn *SkycoinTransaction) GetId() string {
	logCoin.Info("Getting if from transaction")
	return txn.skyTxn.Hash
}

func (txn *SkycoinTransaction) ComputeFee(ticker string) (uint64, error) {
	logCoin.Info("Compute fee for transaction with " + ticker + "ticker")
	if ticker == CoinHour {
		return txn.skyTxn.Fee, nil
	} else if util.StringInList(ticker, txn.SupportedAssets()) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}
func getSkycoinTransactionInputsFromTxnHash(hash string) ([]core.TransactionInput, error) {
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer ReturnSkycoinClient(c)
	transaction, err := c.TransactionVerbose(hash)
	if err != nil {
		return nil, err
	}
	inputs := make([]core.TransactionInput, 0)
	for _, in := range transaction.Transaction.In {
		inputs = append(inputs, &SkycoinTransactionInput{
			skyIn:       in,
			spentOutput: nil,
		})
	}

	return inputs, nil
}

func getSkycoinTransactionInputsFromInputsHashes(inputsHashes []cipher.SHA256) ([]core.TransactionInput, error) {
	inputs := make([]core.TransactionInput, 0)
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer ReturnSkycoinClient(c)

	for _, in := range inputsHashes {
		ux, err := c.UxOut(in.String())
		if err != nil {
			return nil, err
		}
		addr, err := cipher.DecodeBase58Address(ux.OwnerAddress)
		if err != nil {
			return nil, err
		}
		srcTxn, err := cipher.SHA256FromHex(ux.SrcTx)
		if err != nil {
			return nil, err
		}
		cUx := coin.UxOut{
			Head: coin.UxHead{
				BkSeq: ux.SrcBkSeq,
				Time:  ux.Time,
			},
			Body: coin.UxBody{
				Address:        addr,
				Coins:          ux.Coins,
				Hours:          ux.Hours,
				SrcTransaction: srcTxn,
			},
		}

		visorInput, err := visor.NewTransactionInput(cUx, uint64(time.Now().UTC().Unix()))
		if err != nil {
			return nil, err
		}
		readInput, err := readable.NewTransactionInput(visorInput)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, &SkycoinTransactionInput{
			skyIn:       readInput,
			spentOutput: nil,
		})

	}
	return inputs, nil
}

type SkycoinTransactionInput struct {
	skyIn       readable.TransactionInput
	spentOutput *SkycoinTransactionOutput
}

func (in *SkycoinTransactionInput) GetId() string {
	logCoin.Info("Getting id for transaction input")
	return in.skyIn.Hash
}

func (in *SkycoinTransactionInput) GetSpentOutput() core.TransactionOutput {
	logCoin.Info("Getting spent outputs for transaction inputs")

	if in.spentOutput == nil {

		c, err := NewSkycoinApiClient(PoolSection)
		if err != nil {
			return nil
		}
		defer ReturnSkycoinClient(c)
		out, err := c.UxOut(in.skyIn.Hash)
		if err != nil {
			return nil
		}
		skyAccuracy, err := util.AltcoinQuotient("SKY")
		if err != nil {
			return nil
		}

		skyOut := &SkycoinTransactionOutput{
			skyOut: readable.TransactionOutput{
				Address: out.OwnerAddress,
				Coins:   strconv.FormatFloat(float64(out.Coins)/float64(skyAccuracy), 'f', -1, 64),
				Hours:   out.Hours,
				Hash:    out.Uxid,
			},
			spent: true}
		in.spentOutput = skyOut

	}
	return in.spentOutput

}

func (in *SkycoinTransactionInput) GetCoins(ticker string) (uint64, error) {
	logCoin.Info("Getting coins for transaction inputs using " + ticker + "ticker")

	accuracy, err2 := util.AltcoinQuotient(ticker)
	if err2 != nil {
		return uint64(0), err2
	}
	if ticker == Sky {
		skyf, err := strconv.ParseFloat(in.skyIn.Coins, 64)
		if err != nil {
			return 0, err
		}
		return uint64(skyf * float64(accuracy)), nil
	} else if ticker == CoinHour {
		return in.skyIn.Hours * accuracy, nil
	} else if ticker == CalculatedHour {
		return in.skyIn.CalculatedHours * accuracy, nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinTransactionInputIterator
 */
type SkycoinTransactionInputIterator struct {
	current int
	data    []core.TransactionInput
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
	return (iter.current + 1) < len(iter.data)
}

func NewSkycoinTransactioninputIterator(ins []core.TransactionInput) *SkycoinTransactionInputIterator {
	return &SkycoinTransactionInputIterator{data: ins, current: -1}
}

/**
 * SkycoinTransactionOutput
 */
type SkycoinTransactionOutput struct {
	skyOut          readable.TransactionOutput
	spent           bool
	calculatedHours uint64
}

func (out *SkycoinTransactionOutput) GetId() string {
	logCoin.Info("Getting if of transaction output")
	return out.skyOut.Hash

}

func (out *SkycoinTransactionOutput) GetAddress() core.Address {
	logCoin.Info("Getting address for transaction output")
	return &SkycoinAddress{address: out.skyOut.Address}
}

func (out *SkycoinTransactionOutput) GetCoins(ticker string) (uint64, error) {
	logCoin.Info("Getting coins for transaction outputs using " + ticker + "ticker")
	accuracy, err2 := util.AltcoinQuotient(ticker)
	if err2 != nil {
		return uint64(0), err2
	}
	if ticker == Sky {
		skyf, err := strconv.ParseFloat(out.skyOut.Coins, 64)
		if err != nil {
			return 0, err
		}
		return uint64(skyf * float64(accuracy)), nil
	} else if ticker == CoinHour {
		return out.skyOut.Hours * accuracy, nil
	} else if ticker == CalculatedHour {
		return out.calculatedHours * accuracy, nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

func (out *SkycoinTransactionOutput) IsSpent() bool {
	logCoin.Info("Checking if output is spent")
	if out.spent {
		return true
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return true
	}
	defer ReturnSkycoinClient(c)
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
