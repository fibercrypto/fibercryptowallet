package skycoin

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/readable"
)

func TestSkycoinTransactionGetStatus(t *testing.T) {
	global_mock.On("Transaction", "hash1").Return(
		&readable.TransactionWithStatus{
			Status: readable.TransactionStatus{
				Confirmed: true,
			},
		},
		nil,
	)
	global_mock.On("Transaction", "hash2").Return(
		&readable.TransactionWithStatus{
			Status: readable.TransactionStatus{
				Confirmed: false,
			},
		},
		nil,
	)

	thx1 := &SkycoinTransaction{skyTxn: readable.TransactionVerbose{}}
	thx1.skyTxn.Hash = "hash1"
	thx2 := &SkycoinTransaction{skyTxn: readable.TransactionVerbose{}}
	thx2.skyTxn.Hash = "hash2"

	require.Equal(t, thx1.GetStatus(), core.TXN_STATUS_CONFIRMED)
	require.Equal(t, thx2.GetStatus(), core.TXN_STATUS_PENDING)
}

func TestSkycoinTransactionGetInputs(t *testing.T) {
	//set correct return value
	response := &readable.TransactionWithStatusVerbose{
		Transaction: readable.TransactionVerbose{},
	}
	response.Transaction.In = []readable.TransactionInput{
		readable.TransactionInput{
			Hash:            "I1",
			Coins:           "20",
			Hours:           uint64(20),
			CalculatedHours: uint64(20),
		},
		readable.TransactionInput{
			Hash:            "I2",
			Coins:           "20",
			Hours:           uint64(20),
			CalculatedHours: uint64(20),
		},
	}
	global_mock.On("TransactionVerbose", "hash1").Return(response, nil)

	thx1 := &SkycoinTransaction{skyTxn: readable.TransactionVerbose{}}
	thx1.skyTxn.Hash = "hash1"

	inputs := thx1.GetInputs()
	require.Equal(t, inputs[0].GetId(), "I1")
	require.Equal(t, inputs[1].GetId(), "I2")
	it := NewSkycoinTransactioninputIterator(inputs)
	for it.Next() {
		sky, err := it.Value().GetCoins(Sky)
		require.NoError(t, err)
		require.Equal(t, sky, uint64(20000000))
		hours, err1 := it.Value().GetCoins(CoinHour)
		require.NoError(t, err1)
		require.Equal(t, hours, uint64(20))
	}
}

func TestSkycoinTransactionInputGetSpentOutput(t *testing.T) {
	global_mock.On("UxOut", "in1").Return(
		&readable.SpentOutput{
			OwnerAddress: "dir",
			Coins:        uint64(1000000),
			Hours:        uint64(20),
			Uxid:         "out1",
		},
		nil,
	)

	input := &SkycoinTransactionInput{skyIn: readable.TransactionInput{Hash: "in1"}}
	output := input.GetSpentOutput()

	require.Equal(t, output.GetId(), "out1")
	require.Equal(t, output.GetAddress().String(), "dir")
	sky, err := output.GetCoins(Sky)
	require.NoError(t, err)
	require.Equal(t, sky, uint64(1000000))
	hours, err1 := output.GetCoins(CoinHour)
	require.NoError(t, err1)
	require.Equal(t, hours, uint64(20))
}

func TestSkycoinTransactionOutputIsSpent(t *testing.T) {
	global_mock.On("UxOut", "out1").Return(
		&readable.SpentOutput{
			SpentTxnID: "0000000000000000000000000000000000000000000000000000000000000000",
		},
		nil,
	)
	global_mock.On("UxOut", "out2").Return(
		&readable.SpentOutput{
			SpentTxnID: "0",
		},
		nil,
	)

	output1 := &SkycoinTransactionOutput{skyOut: readable.TransactionOutput{Hash: "out1"}}
	output2 := &SkycoinTransactionOutput{skyOut: readable.TransactionOutput{Hash: "out2"}}

	require.Equal(t, output1.IsSpent(), false)
	require.Equal(t, output2.IsSpent(), true)
}

func TestUninjectedTransactionSignedUnsigned(t *testing.T) {
	txn, _, _, err1 := makeTransactionMultipleInputs(t, 2)
	require.NoError(t, err1)
	uiTxn, err := NewUninjectedTransaction(&txn, 1000)
	require.NoError(t, err)
	isFullySigned, err := uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.True(t, isFullySigned)

	txn.Sigs[1] = cipher.Sig{}
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)

	txn.Sigs[0] = cipher.Sig{}
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)

	txn.Sigs = nil
	isFullySigned, err = uiTxn.IsFullySigned()
	require.NoError(t, err)
	require.False(t, isFullySigned)
}

func TestSkycoinUninjectedTransactionGetInputs(t *testing.T) {
	CleanGlobalMock()

	fee := uint64(500)
	txn, err := makeTransaction(t)
	require.NoError(t, err)
	ut, err := NewUninjectedTransaction(&txn, fee)
	require.NoError(t, err)

	b := randBytes(t, 32)
	h, err := cipher.SHA256FromBytes(b)
	require.NoError(t, err)

	ut.txn.In = []cipher.SHA256{h}

	addr := makeAddress()

	global_mock.On("UxOut", h.String()).Return(
		&readable.SpentOutput{
			OwnerAddress: addr.String(),
			Coins:        uint64(1000000),
			Hours:        uint64(20),
			Uxid:         "out1",
			SrcTx:        hex.EncodeToString(h[:]),
		},
		nil,
	)

	tiList := ut.GetInputs()
	ti := tiList[0]
	require.Equal(t, 1, len(tiList))
	sky, err := ti.GetCoins(Sky)
	require.NoError(t, err)
	require.Equal(t, sky, uint64(1000000))
	hours, err := ti.GetCoins(CoinHour)
	require.NoError(t, err)
	require.Equal(t, hours, uint64(20))
	//TODO: Find a way to get the calculatedHours for the expected value
	val, err := ti.GetCoins("INVALID_TICKER")
	require.Error(t, err)
	require.Equal(t, uint64(0), val)
}

func TestSkycoinCreatedTransactionOutputIsSpent(t *testing.T) {
	global_mock.On("UxOut", "out1").Return(
		&readable.SpentOutput{
			SpentTxnID: "0000000000000000000000000000000000000000000000000000000000000000",
		},
		nil,
	)
	global_mock.On("UxOut", "out2").Return(
		&readable.SpentOutput{
			SpentTxnID: "0",
		},
		nil,
	)

	output1 := &SkycoinCreatedTransactionOutput{skyOut: api.CreatedTransactionOutput{UxID: "out1"}}
	output2 := &SkycoinCreatedTransactionOutput{skyOut: api.CreatedTransactionOutput{UxID: "out2"}}

	require.Equal(t, output1.IsSpent(), false)
	require.Equal(t, output2.IsSpent(), true)
	require.Equal(t, output2.IsSpent(), true)

}
