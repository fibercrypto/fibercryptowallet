package skycoin

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/visor"
)

/*
SkycoinPendingTransaction
*/
type SkycoinPendingTransaction struct { //Implements Transaction interface
	Transaction readable.UnconfirmedTransactionVerbose
}

func (txn *SkycoinPendingTransaction) SupportedAssets() []string {
	return []string{Sky, CoinHour, CalculatedHour}
}

func (txn *SkycoinPendingTransaction) GetTimestamp() core.Timestamp {
	return core.Timestamp(txn.Transaction.Received.Unix())
}

func (txn *SkycoinPendingTransaction) GetStatus() core.TransactionStatus {
	return core.TXN_STATUS_PENDING
}

func (txn *SkycoinPendingTransaction) GetInputs() []core.TransactionInput {
	inputs := make([]core.TransactionInput, 0)
	for _, input := range txn.Transaction.Transaction.In {
		inputs = append(inputs, &SkycoinTransactionInput{skyIn: input})
	}
	return inputs
}

func (txn *SkycoinPendingTransaction) GetOutputs() []core.TransactionOutput {
	outputs := make([]core.TransactionOutput, 0)
	for _, output := range txn.Transaction.Transaction.Out {
		outputs = append(outputs, &SkycoinTransactionOutput{skyOut: output, spent: false})
	}
	return outputs
}

func (txn *SkycoinPendingTransaction) GetId() string {
	return txn.Transaction.Transaction.Hash
}

func (txn *SkycoinPendingTransaction) ComputeFee(ticker string) (uint64, error) {
	if ticker == CoinHour {
		return txn.Transaction.Transaction.Fee, nil
	} else if util.StringInList(ticker, txn.SupportedAssets()) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

func newCreatedTransactionOutput(uxID, address, coins, hours string) api.CreatedTransactionOutput {
	return api.CreatedTransactionOutput{
		UxID:    uxID,
		Address: address,
		Coins:   coins,
		Hours:   hours,
	}
}

func newCreatedTransactionInput(uxID, address, coins, hours, calculatedHours string, time, block uint64, txID string) api.CreatedTransactionInput {
	return api.CreatedTransactionInput{
		UxID:            uxID,
		Address:         address,
		Coins:           coins,
		Hours:           hours,
		CalculatedHours: calculatedHours,
		Time:            time,
		Block:           block,
		TxID:            txID,
	}
}

func newCreatedTransaction(length uint32, txnType uint8, txID string, innerHash string, fee string, ins []api.CreatedTransactionInput, outs []api.CreatedTransactionOutput, sigs []string) *api.CreatedTransaction {
	rTxn := api.CreatedTransaction{
		Length:    length,
		Type:      txnType,
		TxID:      txID,
		InnerHash: innerHash,
		Fee:       fee,
		Sigs:      sigs,
		In:        ins,
		Out:       outs,
	}
	return &rTxn
}

func blockTxnToCreatedTxn(blockTxn readable.BlockTransactionVerbose, timestamp uint64) (*api.CreatedTransaction, error) {
	sigs := append([]string{}, blockTxn.Sigs...)
	ins := make([]api.CreatedTransactionInput, len(blockTxn.In))
	outs := make([]api.CreatedTransactionOutput, len(blockTxn.Out))
	for i, input := range blockTxn.In {
		ins[i] = newCreatedTransactionInput(
			input.Hash,
			input.Address,
			input.Coins,
			fmt.Sprint(input.Hours),
			fmt.Sprint(input.CalculatedHours),
			timestamp,
			// Unconfirmed transactions are not included in a block yet
			0,
			blockTxn.Hash,
		)
	}
	for i, output := range blockTxn.Out {
		outs[i] = newCreatedTransactionOutput(
			output.Hash,
			output.Address,
			output.Coins,
			fmt.Sprint(output.Hours),
		)
	}
	return newCreatedTransaction(
		blockTxn.Length,
		blockTxn.Type,
		blockTxn.Hash,
		blockTxn.InnerHash,
		fmt.Sprint(blockTxn.Fee),
		ins, outs, sigs,
	), nil
}

// ToCreatedTransaction return an instance of api.CreatedTransaction equivalent to he current transaction object
func (txn *SkycoinPendingTransaction) ToCreatedTransaction() (*api.CreatedTransaction, error) {
	return blockTxnToCreatedTxn(txn.Transaction.Transaction, uint64(txn.Transaction.Announced.UnixNano()))
}

func serializeCreatedTransaction(txn readableTxn) ([]byte, error) {
	rTxn, err := txn.ToCreatedTransaction()
	if err != nil {
		return nil, err
	}
	skyTxn, err := rTxn.ToTransaction()
	if err != nil {
		return nil, err
	}
	return skyTxn.Serialize()
}

// EncodeSkycoinTransaction serialize transaction data for subsequent broadcast through the peer-to-peer network
func (txn *SkycoinPendingTransaction) EncodeSkycoinTransaction() ([]byte, error) {
	return serializeCreatedTransaction(txn)
}

// skycoinTxn represents the common internal operations that should be applied upon Skycoin transaction wrapper types
type skycoinTxn interface {
	// EncodeSkycoinTransaction serialize transaction data for subsequent broadcast through the peer-to-peer network
	EncodeSkycoinTransaction() ([]byte, error)
}

// readableTxn expreses the contract to use Skycoin readable transactions for signing
type readableTxn interface {
	// ToCreatedTransaction return an instance of api.CreatedTransaction equivalent to he current transaction object
	ToCreatedTransaction() (*api.CreatedTransaction, error)
}

func verifyReadableTransaction(rTxn readableTxn, checkSigned bool) error {
	var createdTxn *api.CreatedTransaction
	if cTxn, err := rTxn.ToCreatedTransaction(); err != nil {
		createdTxn = cTxn
	} else {
		return err
	}
	txn, err := createdTxn.ToTransaction()
	if err != nil {
		return err
	}
	if checkSigned {
		return txn.Verify()
	}
	return txn.VerifyUnsigned()
}

// VerifyUnsigned checks for valid unsigned transaction
func (txn *SkycoinPendingTransaction) VerifyUnsigned() error {
	if !txn.Transaction.IsValid {
		// FIXME: Unique error object
		return errors.New("Invalid unconfirmed transaction")
	}
	return verifyReadableTransaction(txn, false)
}

// VerifySigned checks for valid unsigned transaction
func (txn *SkycoinPendingTransaction) VerifySigned() error {
	if !txn.Transaction.IsValid {
		// FIXME: Unique error object
		return errors.New("Invalid unconfirmed transaction")
	}
	return verifyReadableTransaction(txn, true)
}

func checkFullySigned(rTxn readableTxn) (bool, error) {
	cTxn, err := rTxn.ToCreatedTransaction()
	if err != nil {
		return false, err
	}
	txn, err2 := cTxn.ToTransaction()
	if err2 != nil {
		return false, err2
	}
	return txn.IsFullySigned(), nil
}

// IsFullySigned deermine whether all transaction elements have been signed
func (txn *SkycoinPendingTransaction) IsFullySigned() (bool, error) {
	return checkFullySigned(txn)
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
	return []string{Sky, CoinHour}
}

func (skyTxn *SkycoinUninjectedTransaction) GetTimestamp() core.Timestamp {
	return 0
}

func (skyTxn *SkycoinUninjectedTransaction) GetStatus() core.TransactionStatus {
	return core.TXN_STATUS_CREATED
}

func (skyTxn *SkycoinUninjectedTransaction) GetInputs() []core.TransactionInput {
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
	if ticker == CoinHour {
		return skyTxn.fee, nil
	}
	return 0, nil
}

func (skyTxn *SkycoinUninjectedTransaction) GetId() string {
	return skyTxn.txn.Hash().String()
}

// VerifyUnsigned checks for valid unsigned transaction
func (txn *SkycoinUninjectedTransaction) VerifyUnsigned() error {
	return txn.txn.VerifyUnsigned()
}

// VerifySigned checks for valid unsigned transaction
func (txn *SkycoinUninjectedTransaction) VerifySigned() error {
	return txn.txn.Verify()
}

// IsFullySigned deermine whether all transaction elements have been signed
func (txn *SkycoinUninjectedTransaction) IsFullySigned() (bool, error) {
	return txn.txn.IsFullySigned(), nil
}

func (txn *SkycoinUninjectedTransaction) EncodeSkycoinTransaction() ([]byte, error) {
	return txn.txn.Serialize()
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
	return []string{Sky, CoinHour, CalculatedHour}
}

func (txn *SkycoinTransaction) GetTimestamp() core.Timestamp {
	return core.Timestamp(txn.skyTxn.Timestamp)
}

func (txn *SkycoinTransaction) GetStatus() core.TransactionStatus {

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
	return txn.skyTxn.Hash
}

func (txn *SkycoinTransaction) ComputeFee(ticker string) (uint64, error) {
	if ticker == CoinHour {
		return txn.skyTxn.Fee, nil
	} else if util.StringInList(ticker, txn.SupportedAssets()) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

// EncodeSkycoinTransaction serialize transaction data for subsequent broadcast through the peer-to-peer network
func (txn *SkycoinTransaction) EncodeSkycoinTransaction() ([]byte, error) {
	return serializeCreatedTransaction(txn)
}

// ToCreatedTransaction retrieve the equivalent core.Transaction object
func (txn *SkycoinTransaction) ToCreatedTransaction() (*api.CreatedTransaction, error) {
	return blockTxnToCreatedTxn(txn.skyTxn.BlockTransactionVerbose, uint64(txn.skyTxn.Timestamp))
}

// VerifyUnsigned checks for valid unsigned transaction
func (txn *SkycoinTransaction) VerifyUnsigned() error {
	return verifyReadableTransaction(txn, false)
}

// VerifySigned checks for valid unsigned transaction
func (txn *SkycoinTransaction) VerifySigned() error {
	return verifyReadableTransaction(txn, true)
}

// IsFullySigned deermine whether all transaction elements have been signed
func (txn *SkycoinTransaction) IsFullySigned() (bool, error) {
	return checkFullySigned(txn)
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

/*
SkycoinTransacionInput wraps verbose readable transaction input data
*/
type SkycoinTransactionInput struct {
	skyIn       readable.TransactionInput
	spentOutput *SkycoinTransactionOutput
}

func (in *SkycoinTransactionInput) GetId() string {
	return in.skyIn.Hash
}

func (in *SkycoinTransactionInput) GetSpentOutput() core.TransactionOutput {
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
	return out.skyOut.Hash

}

func (out *SkycoinTransactionOutput) GetAddress() core.Address {
	return &SkycoinAddress{address: out.skyOut.Address}
}

func (out *SkycoinTransactionOutput) GetCoins(ticker string) (uint64, error) {
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

func newCreatedTransactionInputs(rIns []api.CreatedTransactionInput) []core.TransactionInput {
	ins := make([]core.TransactionInput, len(rIns))
	for i, rIn := range rIns {
		ins[i] = &SkycoinCreatedTransactionInput{
			skyIn: rIn,
		}
	}
	return ins
}

/*
SkycoinCreatedTransacionInput wraps created transaction input data
*/
type SkycoinCreatedTransactionInput struct {
	skyIn       api.CreatedTransactionInput
	spentOutput *SkycoinCreatedTransactionOutput
}

// GetId return transaction UXID
func (in *SkycoinCreatedTransactionInput) GetId() string {
	return in.skyIn.UxID
}

func (in *SkycoinCreatedTransactionInput) GetSpentOutput() core.TransactionOutput {
	if in.spentOutput == nil {

		calculatedHours, err := in.GetCoins(in.skyIn.CalculatedHours)
		if err != nil {
			calculatedHours = 0
		}
		skyOut := &SkycoinCreatedTransactionOutput{
			skyOut: api.CreatedTransactionOutput{
				Address: in.skyIn.Address,
				Coins:   in.skyIn.Coins,
				Hours:   in.skyIn.Hours,
				UxID:    in.skyIn.UxID,
			},
			calculatedHours: calculatedHours,
			spent:           false}
		in.spentOutput = skyOut

	}
	return in.spentOutput

}

// GetCoins return input balance in one of supported coins , or error
func (in *SkycoinCreatedTransactionInput) GetCoins(ticker string) (uint64, error) {
	accuracy, err2 := util.AltcoinQuotient(ticker)
	if err2 != nil {
		return uint64(0), err2
	}
	var result uint64
	var tmpResult int64
	var err error
	if ticker == Sky {
		var skyf float64
		skyf, err = strconv.ParseFloat(in.skyIn.Coins, 64)
		result = uint64(skyf * float64(accuracy))
	} else if ticker == CoinHour {
		tmpResult, err = strconv.ParseInt(in.skyIn.Hours, 10, 64)
		result = uint64(tmpResult)
	} else if ticker == CalculatedHour {
		tmpResult, err = strconv.ParseInt(in.skyIn.CalculatedHours, 10, 64)
		result = uint64(tmpResult)
	} else {
		return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
	}
	if err != nil {
		return 0, err
	}
	return result, nil
}

func newCreatedTransactionOutputs(rOuts []api.CreatedTransactionOutput) []core.TransactionOutput {
	ins := make([]core.TransactionOutput, len(rOuts))
	for i, rOut := range rOuts {
		ins[i] = &SkycoinCreatedTransactionOutput{
			skyOut: rOut,
		}
	}
	return ins
}

/**
 * SkycoinCreatedTransactionOutput
 */
type SkycoinCreatedTransactionOutput struct {
	skyOut          api.CreatedTransactionOutput
	spent           bool
	calculatedHours uint64
}

func (out *SkycoinCreatedTransactionOutput) GetId() string {
	return out.skyOut.UxID
}

func (out *SkycoinCreatedTransactionOutput) GetAddress() core.Address {
	return &SkycoinAddress{address: out.skyOut.Address}
}

func (out *SkycoinCreatedTransactionOutput) GetCoins(ticker string) (uint64, error) {
	accuracy, err2 := util.AltcoinQuotient(ticker)
	if err2 != nil {
		return uint64(0), err2
	}
	var tmpResult int64
	var result uint64
	var err error
	if ticker == Sky {
		var skyf float64
		skyf, err = strconv.ParseFloat(out.skyOut.Coins, 64)
		result = uint64(skyf * float64(accuracy))
	} else if ticker == CoinHour {
		tmpResult, err = strconv.ParseInt(out.skyOut.Hours, 10, 64)
		result = uint64(tmpResult)
	} else if ticker == CalculatedHour {
		result = out.calculatedHours
		err = nil
	} else {
		err = fmt.Errorf("Invalid ticker %v\n", ticker)
	}
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (out *SkycoinCreatedTransactionOutput) IsSpent() bool {
	if out.spent {
		return true
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return true
	}
	defer ReturnSkycoinClient(c)
	ou, err := c.UxOut(out.skyOut.UxID)
	if err != nil {
		return false
	}
	if ou.SpentTxnID != "0000000000000000000000000000000000000000000000000000000000000000" {
		out.spent = true
		return true
	}
	return false
}

// NewSkycoinCreatedTransaction return readable created transaction wrapper
func NewSkycoinCreatedTransaction(rTxn api.CreatedTransaction) *SkycoinCreatedTransaction {
	return &SkycoinCreatedTransaction{
		skyTxn: rTxn,
	}
}

/*
SkycoinCreatedTransaction wraps a readable created transaction to implement core.Transaction interface
*/
type SkycoinCreatedTransaction struct {
	skyTxn api.CreatedTransaction

	status  core.TransactionStatus
	inputs  []core.TransactionInput
	outputs []core.TransactionOutput
}

// SupportedAssets are SKY, SKYCH, and accumulated SKYCH
func (txn *SkycoinCreatedTransaction) SupportedAssets() []string {
	return []string{Sky, CoinHour, CalculatedHour}
}

// GetTimestamp will return zero
func (txn *SkycoinCreatedTransaction) GetTimestamp() core.Timestamp {
	return 0
}

func (txn *SkycoinCreatedTransaction) GetStatus() core.TransactionStatus {
	return core.TXN_STATUS_CREATED
}

// GetInputs return inputs spent by this transaction
func (txn *SkycoinCreatedTransaction) GetInputs() []core.TransactionInput {
	if txn.inputs == nil {
		txn.inputs = newCreatedTransactionInputs(txn.skyTxn.In)
	}
	return txn.inputs
}

// GetOuptuts return outputs owned by transaction receivers
func (txn *SkycoinCreatedTransaction) GetOutputs() []core.TransactionOutput {
	if txn.outputs == nil {
		txn.outputs = newCreatedTransactionOutputs(txn.skyTxn.Out)
	}
	return txn.outputs
}

func (txn *SkycoinCreatedTransaction) GetId() string {
	return txn.skyTxn.TxID
}

func (txn *SkycoinCreatedTransaction) ComputeFee(ticker string) (uint64, error) {
	if ticker == CoinHour {
		fee, err := strconv.ParseInt(txn.skyTxn.Fee, 10, 64)
		if err != nil {
			return uint64(0), err
		}
		return uint64(fee), nil
	} else if util.StringInList(ticker, txn.SupportedAssets()) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

// EncodeSkycoinTransaction serialize transaction data for subsequent broadcast through the peer-to-peer network
func (txn *SkycoinCreatedTransaction) EncodeSkycoinTransaction() ([]byte, error) {
	return serializeCreatedTransaction(txn)
}

// ToCreatedTransaction retrieve the equivalent core.Transaction object
func (txn *SkycoinCreatedTransaction) ToCreatedTransaction() (*api.CreatedTransaction, error) {
	return &txn.skyTxn, nil
}

// VerifyUnsigned checks for valid unsigned transaction
func (txn *SkycoinCreatedTransaction) VerifyUnsigned() error {
	return verifyReadableTransaction(txn, false)
}

// VerifySigned checks for valid unsigned transaction
func (txn *SkycoinCreatedTransaction) VerifySigned() error {
	return verifyReadableTransaction(txn, true)
}

// IsFullySigned deermine whether all transaction elements have been signed
func (txn *SkycoinCreatedTransaction) IsFullySigned() (bool, error) {
	return checkFullySigned(txn)
}

// Type assertions to abort compilation if contracts not satisfied
var (
	pendingTxn         SkycoinPendingTransaction
	skyTxnInIter       SkycoinTransactionInputIterator
	skyTxnIter         SkycoinTransactionIterator
	skyTxnOutIter      SkycoinTransactionOutputIterator
	uninjectedTxn      SkycoinUninjectedTransaction
	skyTxn             SkycoinTransaction
	skyTxnIn           SkycoinTransactionInput
	skyTxnOut          SkycoinTransactionOutput
	cTxn               SkycoinCreatedTransaction
	pendingTxnSky      skycoinTxn                     = &pendingTxn
	pendingTxnReadable readableTxn                    = &pendingTxn
	pendingTxnCore     core.Transaction               = &pendingTxn
	skyTxnIterator     core.TransactionIterator       = &skyTxnIter
	skyTxnInIterator   core.TransactionInputIterator  = &skyTxnInIter
	skyTxnOutIterator  core.TransactionOutputIterator = &skyTxnOutIter
	uninjectedTxnCore  core.Transaction               = &uninjectedTxn
	uninjectedTxnSky   skycoinTxn                     = &uninjectedTxn
	skyTxnSky          skycoinTxn                     = &skyTxn
	skyTxnReadable     readableTxn                    = &skyTxn
	skyTxnCore         core.Transaction               = &skyTxn
	skyTxnInCore       core.TransactionInput          = &skyTxnIn
	skyTxnOutCore      core.TransactionOutput         = &skyTxnOut
	cTxnSky            skycoinTxn                     = &cTxn
	cTxnReadable       readableTxn                    = &cTxn
	cTxnCore           core.Transaction               = &cTxn
)
