package skycoin

import (
	"encoding/hex"
	goerrors "errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/SkycoinProject/skycoin/src/testutil"
	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/requirethat"
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
			OwnerAddress: "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8",
			Coins:        uint64(1000000),
			Hours:        uint64(20),
			Uxid:         "out1",
		},
		nil,
	)

	input := &SkycoinTransactionInput{skyIn: readable.TransactionInput{Hash: "in1"}}
	output := input.GetSpentOutput()

	t.Logf("%#v", output)
	require.Equal(t, output.GetId(), "out1")
	require.Equal(t, output.GetAddress().String(), "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8")
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

func TestSupportedAssets(t *testing.T) {
	pendTxn := new(SkycoinPendingTransaction)
	assets := pendTxn.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour}, assets)

	coreTxn := new(SkycoinUninjectedTransaction)
	assets = coreTxn.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour}, assets)

	skyTxn := new(SkycoinTransaction)
	assets = skyTxn.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour}, assets)

	cTxn := new(SkycoinCreatedTransaction)
	assets = cTxn.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour}, assets)

	skyTxnOut := new(SkycoinTransactionOutput)
	assets = skyTxnOut.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour, CalculatedHour}, assets)

	skyTxnIn := new(SkycoinTransactionInput)
	assets = skyTxnIn.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour, CalculatedHour}, assets)

	cTxnOut := new(SkycoinCreatedTransactionOutput)
	assets = cTxnOut.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour, CalculatedHour}, assets)

	cTxnIn := new(SkycoinCreatedTransactionInput)
	assets = cTxnIn.SupportedAssets()
	requirethat.ElementsMatch(t, []string{Sky, CoinHour, CalculatedHour}, assets)
}

func TestPendingTxnStatus(t *testing.T) {
	pendTxn := new(SkycoinPendingTransaction)
	require.Equal(t, core.TXN_STATUS_PENDING, pendTxn.GetStatus())
}

func TestUninjectedTxnTimestamp(t *testing.T) {
	coreTxn := new(SkycoinUninjectedTransaction)
	require.Equal(t, core.Timestamp(0), coreTxn.GetTimestamp())
}

func TestUninjectedTxnStatus(t *testing.T) {
	coreTxn := new(SkycoinUninjectedTransaction)
	require.Equal(t, core.TXN_STATUS_CREATED, coreTxn.GetStatus())
}

func TestUninjectedTxnFee(t *testing.T) {
	coreTxn := new(SkycoinUninjectedTransaction)

	fee, err := coreTxn.ComputeFee(Sky)
	require.NoError(t, err)
	require.Equal(t, uint64(0), fee)

	_, err = coreTxn.ComputeFee(CalculatedHour)
	testutil.RequireError(t, err, "Feature not implemented")

	coreTxn.fee = 64
	fee, err = coreTxn.ComputeFee(CoinHour)
	require.NoError(t, err)
	require.Equal(t, uint64(64), fee)

	_, err = coreTxn.ComputeFee("NOCOINATALL")
	testutil.RequireError(t, err, "Invalid ticker")
}

func TestSkycoinTxnFee(t *testing.T) {
	skyTxn := new(SkycoinTransaction)

	fee, err := skyTxn.ComputeFee(Sky)
	require.NoError(t, err)
	require.Equal(t, uint64(0), fee)

	_, err = skyTxn.ComputeFee(CalculatedHour)
	testutil.RequireError(t, err, "Feature not implemented")

	_, err = skyTxn.ComputeFee("NOCOINATALL")
	testutil.RequireError(t, err, "Invalid ticker")
}

func TestSkycoinCreatedTxnFee(t *testing.T) {
	cTxn := new(SkycoinCreatedTransaction)

	fee, err := cTxn.ComputeFee(Sky)
	require.NoError(t, err)
	require.Equal(t, uint64(0), fee)

	_, err = cTxn.ComputeFee(CalculatedHour)
	testutil.RequireError(t, err, "Feature not implemented")

	_, err = cTxn.ComputeFee("NOCOINATALL")
	testutil.RequireError(t, err, "Invalid ticker")
}

func TestPendingTxnGetTimestamp(t *testing.T) {
	cur := time.Now()
	sTxn := new(SkycoinPendingTransaction)
	sTxn.Transaction = &readable.UnconfirmedTransactionVerbose{Received: cur}

	require.Equal(t, core.Timestamp(cur.Unix()), sTxn.GetTimestamp())
}

func TestPendingTxnGetInputs(t *testing.T) {
	hashes := make([]string, 0)
	for i := 0; i < 10; i++ {
		hashes = append(hashes, fmt.Sprintf("hash%d", i))
	}
	sTxn := new(SkycoinPendingTransaction)
	inputs := make([]readable.TransactionInput, 0)
	for _, hash := range hashes {
		inputs = append(inputs, readable.TransactionInput{Hash: hash})
	}
	sTxn.Transaction = &readable.UnconfirmedTransactionVerbose{
		Transaction: readable.BlockTransactionVerbose{
			In: inputs,
		},
	}
	inHashes := make([]string, 0)
	for _, input := range sTxn.GetInputs() {
		inHashes = append(inHashes, input.GetId())
	}
	requirethat.ElementsMatch(t, hashes, inHashes)
}

func TestPendingTxnGetOutputs(t *testing.T) {
	hashes := make([]string, 0)
	for i := 0; i < 10; i++ {
		hashes = append(hashes, fmt.Sprintf("hash%d", i))
	}
	sTxn := new(SkycoinPendingTransaction)
	outputs := make([]readable.TransactionOutput, 0)
	for _, hash := range hashes {
		outputs = append(outputs, readable.TransactionOutput{Hash: hash})
	}
	sTxn.Transaction = &readable.UnconfirmedTransactionVerbose{
		Transaction: readable.BlockTransactionVerbose{
			Out: outputs,
		},
	}
	outHashes := make([]string, 0)
	for _, output := range sTxn.GetOutputs() {
		outHashes = append(outHashes, output.GetId())
	}
	requirethat.ElementsMatch(t, hashes, outHashes)
}

func TestPendingTxnGetId(t *testing.T) {
	hashes := make([]string, 0)
	for i := 0; i < 10; i++ {
		hashes = append(hashes, fmt.Sprintf("hash%d", i))
	}
	for _, hash := range hashes {
		t.Run(hash, func(t *testing.T) {
			sTxn := new(SkycoinPendingTransaction)
			sTxn.Transaction = &readable.UnconfirmedTransactionVerbose{
				Transaction: readable.BlockTransactionVerbose{
					Hash: hash,
				},
			}
			require.Equal(t, hash, sTxn.GetId())
		})
	}
}

func TestPendingTxnComputeFee(t *testing.T) {
	tests := []struct {
		ticker      string
		fee         uint64
		amount      uint64
		wantedError error
	}{
		{ticker: Sky, wantedError: nil},
		{ticker: CoinHour, fee: 20, amount: 20, wantedError: nil},
		{ticker: CoinHour, fee: 42, amount: 42, wantedError: nil},
		{ticker: CalculatedHour, wantedError: errors.ErrNotImplemented},
		{ticker: "INVALIDTICKER", wantedError: errors.ErrInvalidAltcoinTicker},
	}

	for _, tt := range tests {
		t.Run(tt.ticker, func(t *testing.T) {
			sTxn := new(SkycoinPendingTransaction)
			sTxn.Transaction = &readable.UnconfirmedTransactionVerbose{
				Transaction: readable.BlockTransactionVerbose{
					Fee: tt.fee,
				},
			}
			amount, err := sTxn.ComputeFee(tt.ticker)
			require.Equal(t, tt.amount, amount)
			require.Equal(t, tt.wantedError, err)
		})
	}
}

func Test_newCreatedTransactionOutput(t *testing.T) {
	tests := []struct {
		uxID    string
		address string
		coins   string
		hours   string
	}{
		{uxID: "uxId1", address: "addr1", coins: "coins1", hours: "hours1"},
		{uxID: "uxId2", address: "addr2", coins: "coins2", hours: "hours2"},
		{uxID: "uxId3", address: "addr3", coins: "coins3", hours: "hours3"},
		{uxID: "uxId4", address: "addr4", coins: "coins4", hours: "hours4"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			txn := newCreatedTransactionOutput(tt.uxID, tt.address, tt.coins, tt.hours)
			require.Equal(t, tt.uxID, txn.UxID)
			require.Equal(t, tt.address, txn.Address)
			require.Equal(t, tt.coins, txn.Coins)
			require.Equal(t, tt.hours, txn.Hours)
		})
	}
}

func Test_newCreatedTransactionInput(t *testing.T) {
	tests := []struct {
		uxID            string
		address         string
		coins           string
		hours           string
		calculatedHours string
		time            uint64
		block           uint64
		txID            string
	}{
		{
			uxID:            "uxId1",
			address:         "addr1",
			coins:           "coins1",
			hours:           "hours1",
			calculatedHours: "cH1",
			txID:            "id1",
			time:            1,
			block:           1,
		},
		{
			uxID:            "uxId2",
			address:         "addr2",
			coins:           "coins2",
			hours:           "hours2",
			calculatedHours: "cH2",
			txID:            "id2",
			time:            2,
			block:           2,
		},
		{
			uxID:            "uxId3",
			address:         "addr3",
			coins:           "coins3",
			hours:           "hours3",
			calculatedHours: "cH3",
			txID:            "id3",
			time:            3,
			block:           3,
		},
		{
			uxID:            "uxId4",
			address:         "addr4",
			coins:           "coins4",
			hours:           "hours4",
			calculatedHours: "cH4",
			txID:            "id4",
			time:            4,
			block:           4,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			txn := newCreatedTransactionInput(
				tt.uxID, tt.address, tt.coins, tt.hours, tt.calculatedHours, tt.time, tt.block, tt.txID,
			)
			require.Equal(t, tt.uxID, txn.UxID)
			require.Equal(t, tt.address, txn.Address)
			require.Equal(t, tt.coins, txn.Coins)
			require.Equal(t, tt.hours, txn.Hours)
			require.Equal(t, tt.calculatedHours, txn.CalculatedHours)
			require.Equal(t, tt.time, txn.Time)
			require.Equal(t, tt.block, txn.Block)
			require.Equal(t, tt.txID, txn.TxID)
		})
	}
}

func Test_newCreatedTransaction(t *testing.T) {
	tests := []struct {
		length    uint32
		txnType   uint8
		txID      string
		innerHash string
		fee       string
		ins       []api.CreatedTransactionInput
		outs      []api.CreatedTransactionOutput
		sigs      []string
	}{
		{
			length:    1,
			txnType:   1,
			txID:      "txID1",
			innerHash: "hash1",
			fee:       "fee1",
			ins: []api.CreatedTransactionInput{
				api.CreatedTransactionInput{UxID: "UxID1"},
			},
			outs: []api.CreatedTransactionOutput{
				api.CreatedTransactionOutput{UxID: "UxID1"},
			},
			sigs: []string{"first1", "second1"},
		},
		{
			length:    2,
			txnType:   2,
			txID:      "txID2",
			innerHash: "hash2",
			fee:       "fee2",
			ins: []api.CreatedTransactionInput{
				api.CreatedTransactionInput{UxID: "UxID2"},
			},
			outs: []api.CreatedTransactionOutput{
				api.CreatedTransactionOutput{UxID: "UxID2"},
			},
			sigs: []string{"first2", "second2"},
		},
		{
			length:    3,
			txnType:   3,
			txID:      "txID3",
			innerHash: "hash3",
			fee:       "fee3",
			ins: []api.CreatedTransactionInput{
				api.CreatedTransactionInput{UxID: "UxID3"},
			},
			outs: []api.CreatedTransactionOutput{
				api.CreatedTransactionOutput{UxID: "UxID3"},
			},
			sigs: []string{"first3", "second3"},
		},
		{
			length:    4,
			txnType:   4,
			txID:      "txID4",
			innerHash: "hash4",
			fee:       "fee4",
			ins: []api.CreatedTransactionInput{
				api.CreatedTransactionInput{UxID: "UxID4"},
			},
			outs: []api.CreatedTransactionOutput{
				api.CreatedTransactionOutput{UxID: "UxID4"},
			},
			sigs: []string{"first4", "second4"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			txn := newCreatedTransaction(
				tt.length,
				tt.txnType,
				tt.txID,
				tt.innerHash,
				tt.fee,
				tt.ins,
				tt.outs,
				tt.sigs,
			)
			require.Equal(t, tt.length, txn.Length)
			require.Equal(t, tt.txnType, txn.Type)
			require.Equal(t, tt.txID, txn.TxID)
			require.Equal(t, tt.innerHash, txn.InnerHash)
			require.Equal(t, tt.fee, txn.Fee)
			requirethat.ElementsMatch(t, tt.ins, txn.In)
			requirethat.ElementsMatch(t, tt.outs, txn.Out)
			requirethat.ElementsMatch(t, tt.sigs, txn.Sigs)
		})
	}
}

func Test_blockTxnToCreatedTxn(t *testing.T) {
	tests := []struct {
		Address         string
		Coins           string
		Hours           uint64
		CalculatedHours uint64
		Time            uint64
		Block           uint64
		Hash            string
		Length          uint32
		Type            uint8
		InnerHash       string
		Fee             uint64
		sigs            []string
	}{
		{
			Address:         "addr1",
			Coins:           "coins1",
			Hours:           1,
			CalculatedHours: 1,
			Time:            1,
			Hash:            "hash1",
			Length:          1,
			Type:            1,
			InnerHash:       "inner1",
			Fee:             1,
			sigs:            []string{"first1", "second1"},
		},
		{
			Address:         "addr2",
			Coins:           "coins2",
			Hours:           2,
			CalculatedHours: 2,
			Time:            2,
			Hash:            "hash2",
			Length:          2,
			Type:            2,
			InnerHash:       "inner2",
			Fee:             2,
			sigs:            []string{"first2", "second2"},
		},
		{
			Address:         "addr3",
			Coins:           "coins3",
			Hours:           3,
			CalculatedHours: 3,
			Time:            3,
			Hash:            "hash3",
			Length:          3,
			Type:            3,
			InnerHash:       "inner3",
			Fee:             3,
			sigs:            []string{"first3", "second3"},
		},
		{
			Address:         "addr4",
			Coins:           "coins4",
			Hours:           4,
			CalculatedHours: 4,
			Time:            4,
			Hash:            "hash4",
			Length:          4,
			Type:            4,
			InnerHash:       "inner4",
			Fee:             4,
			sigs:            []string{"first4", "second4"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(*testing.T) {
			block := readable.BlockTransactionVerbose{
				Length:    tt.Length,
				Type:      tt.Type,
				Hash:      tt.Hash,
				InnerHash: tt.InnerHash,
				Fee:       tt.Fee,
				In: []readable.TransactionInput{
					readable.TransactionInput{
						Hash:            tt.Hash,
						Address:         tt.Address,
						Coins:           tt.Coins,
						Hours:           tt.Hours,
						CalculatedHours: tt.CalculatedHours,
					},
				},
				Out: []readable.TransactionOutput{
					readable.TransactionOutput{
						Hash:    tt.Hash,
						Address: tt.Address,
						Coins:   tt.Coins,
						Hours:   tt.Hours,
					},
				},
				Sigs: tt.sigs,
			}
			txn, err := blockTxnToCreatedTxn(block, tt.Time)
			require.NoError(t, err)
			require.Equal(t, tt.Length, txn.Length)
			require.Equal(t, tt.Type, txn.Type)
			require.Equal(t, tt.Hash, txn.TxID)
			require.Equal(t, tt.InnerHash, txn.InnerHash)
			require.Equal(t, fmt.Sprint(tt.Fee), txn.Fee)
			require.Equal(t, 1, len(txn.In))
			require.Equal(t, 1, len(txn.Out))
			requirethat.ElementsMatch(t,
				[]api.CreatedTransactionInput{
					newCreatedTransactionInput(
						tt.Hash,
						tt.Address,
						tt.Coins,
						fmt.Sprint(tt.Hours),
						fmt.Sprint(tt.CalculatedHours),
						tt.Time,
						tt.Block,
						tt.Hash,
					),
				},
				txn.In,
			)
			requirethat.ElementsMatch(t,
				[]api.CreatedTransactionOutput{
					newCreatedTransactionOutput(
						tt.Hash, tt.Address, tt.Coins, fmt.Sprint(tt.Hours),
					),
				},
				txn.Out,
			)
			requirethat.ElementsMatch(t, tt.sigs, txn.Sigs)
		})
	}
}

func TestPendingTxnToCreatedTransaction(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			txn := &SkycoinPendingTransaction{
				Transaction: &readable.UnconfirmedTransactionVerbose{
					Transaction: readable.BlockTransactionVerbose{
						Hash: fmt.Sprintf("hash%d", i),
					},
					Announced: time.Now(),
				},
			}
			expected, err := blockTxnToCreatedTxn(
				txn.Transaction.Transaction,
				uint64(txn.Transaction.Announced.UnixNano()),
			)
			require.NoError(t, err)
			created, err1 := txn.ToCreatedTransaction()
			require.NoError(t, err1)
			require.Equal(t, expected, created)
		})
	}
}

func Test_serializeCreatedTransaction(t *testing.T) {
	mockTxn := new(mocks.ReadableTxn)
	id := "0000000000000000000000000000000000000000000000000000000000000000"
	created := &api.CreatedTransaction{
		TxID:      "78877fa898f0b4c45c9c33ae941e40617ad7c8657a307db62bc5691f92f4f60e",
		InnerHash: "No-match",
	}
	mockTxn.On("ToCreatedTransaction").Return(nil, goerrors.New("failure")).Once()
	mockTxn.On("ToCreatedTransaction").Return(created, nil)

	_, err := serializeCreatedTransaction(mockTxn)
	require.Error(t, err)

	_, err = serializeCreatedTransaction(mockTxn)
	require.Error(t, err)

	created.InnerHash = id
	ser, err := serializeCreatedTransaction(mockTxn)
	require.NoError(t, err)
	txn, err := created.ToTransaction()
	require.NoError(t, err)
	expected, err := txn.Serialize()
	require.NoError(t, err)
	require.Equal(t, expected, ser)
}

func TestSkycoinPendingTransactionEncodeSkycoinTransaction(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	require.NoError(t, err)
	sTxn := &SkycoinPendingTransaction{
		Transaction: &readable.UnconfirmedTransactionVerbose{
			Transaction: readable.BlockTransactionVerbose{
				InnerHash: "0000000000000000000000000000000000000000000000000000000000000000",
				Hash:      "78877fa898f0b4c45c9c33ae941e40617ad7c8657a307db62bc5691f92f4f60e",
			},
			Announced: date,
		},
	}
	ser, err := sTxn.EncodeSkycoinTransaction()
	require.NoError(t, err)
	exp, err := serializeCreatedTransaction(sTxn)
	require.NoError(t, err)
	require.Equal(t, exp, ser)
}
