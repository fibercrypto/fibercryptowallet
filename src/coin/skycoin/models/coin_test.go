package skycoin

import (
	"encoding/hex"
	goerrors "errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/coin"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/skytypes"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util/requirethat"
)

func TestTransactionGetStatus(t *testing.T) {
	global_mock.On("Transaction", "hash1").Return(
		&readable.TransactionWithStatus{
			Status: readable.TransactionStatus{
				Confirmed: true,
			},
		},
		nil,
	).Once()
	global_mock.On("Transaction", "hash2").Return(
		&readable.TransactionWithStatus{
			Status: readable.TransactionStatus{
				Confirmed: false,
			},
		},
		nil,
	).Once()
	global_mock.On("Transaction", "hash3").Return(nil, goerrors.New("failure")).Once()

	tests := []struct {
		name string
		txn  core.Transaction
		want core.TransactionStatus
	}{
		{
			name: "SkycoinTransaction-Confirmed",
			txn: &SkycoinTransaction{
				skyTxn: readable.TransactionVerbose{
					BlockTransactionVerbose: readable.BlockTransactionVerbose{
						Hash: "hash1",
					},
				},
			},
			want: core.TXN_STATUS_CONFIRMED,
		},
		{
			name: "SkycoinTransaction-Unconfirmed",
			txn: &SkycoinTransaction{
				skyTxn: readable.TransactionVerbose{
					BlockTransactionVerbose: readable.BlockTransactionVerbose{
						Hash: "hash2",
					},
				},
			},
			want: core.TXN_STATUS_PENDING,
		},
		{
			name: "SkycoinTransaction-ApiError",
			txn: &SkycoinTransaction{
				skyTxn: readable.TransactionVerbose{
					BlockTransactionVerbose: readable.BlockTransactionVerbose{
						Hash: "hash3",
					},
				},
			},
			want: core.TXN_STATUS_CREATED,
		},
		{
			name: "SkycoinPendingTransaction",
			txn:  new(SkycoinPendingTransaction),
			want: core.TXN_STATUS_PENDING,
		},
		{
			name: "SkycoinUnjectedTransaction",
			txn:  new(SkycoinUninjectedTransaction),
			want: core.TXN_STATUS_CREATED,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.txn.GetStatus())
		})
	}
}

func TestSkycoinTransactionGetInputs(t *testing.T) {
	skyAmount := uint64(20000000)
	chAmount := uint64(20)

	response := &readable.TransactionWithStatusVerbose{
		Transaction: readable.TransactionVerbose{
			BlockTransactionVerbose: readable.BlockTransactionVerbose{
				In: []readable.TransactionInput{
					readable.TransactionInput{
						Hash:  "I1",
						Coins: "20",
						Hours: chAmount,
					},
					readable.TransactionInput{
						Hash:  "I2",
						Coins: "20",
						Hours: uint64(20),
					},
				},
			},
		},
	}
	global_mock.On("TransactionVerbose", "hash1").Return(nil, goerrors.New("failure")).Once()
	global_mock.On("TransactionVerbose", "hash1").Return(response, nil).Once()

	st := new(SkycoinTransaction)
	st.skyTxn.Hash = "hash1"

	tests := []struct {
		name    string
		txn     core.Transaction
		ids     []string
		wantNil bool
	}{
		{
			name:    "SkycoinTransaction-ApiError",
			txn:     st,
			wantNil: true,
		},
		{
			name: "SkycoinTransaction",
			txn:  st,
			ids:  []string{"I1", "I2"},
		},
		{
			name: "SkycoinTransaction-InputsSaved",
			txn:  st,
			ids:  []string{"I1", "I2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs := tt.txn.GetInputs()
			if tt.wantNil {
				require.Nil(t, inputs)
			} else {
				ids := make([]string, 0)
				it := NewSkycoinTransactioninputIterator(inputs)
				for it.Next() {
					input := it.Value()
					ids = append(ids, input.GetId())
					sky, err := input.GetCoins(Sky)
					require.NoError(t, err)
					require.Equal(t, skyAmount, sky)
					hours, err1 := input.GetCoins(CoinHour)
					require.NoError(t, err1)
					require.Equal(t, chAmount, hours)
				}
				requirethat.ElementsMatch(t, tt.ids, ids)
			}
		})
	}
}

func TestSkycoinTransactionInputGetSpentOutput(t *testing.T) {
	CleanGlobalMock()

	input := &SkycoinTransactionInput{skyIn: readable.TransactionInput{Hash: "in1"}}
	global_mock.On("UxOut", "in1").Return(nil, goerrors.New("failure")).Once()
	output := input.GetSpentOutput()
	require.Nil(t, output)

	global_mock.On("UxOut", "in1").Return(
		&readable.SpentOutput{
			OwnerAddress: "2JJ8pgq8EDAnrzf9xxBJapE2qkYLefW4uF8",
			Coins:        uint64(1000000),
			Hours:        uint64(20),
			Uxid:         "out1",
		},
		nil,
	).Once()

	output = input.GetSpentOutput()

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
	badID := "0000000000000000000000000000000000000000000000000000000000000000"
	global_mock.On("UxOut", "out").Return(nil, goerrors.New("failure")).Once()
	global_mock.On("UxOut", "out").Return(&readable.SpentOutput{SpentTxnID: badID}, nil).Once()
	global_mock.On("UxOut", "out").Return(&readable.SpentOutput{SpentTxnID: "42"}, nil).Once()

	output := &SkycoinTransactionOutput{skyOut: readable.TransactionOutput{Hash: "out"}}

	require.Equal(t, output.IsSpent(), false)
	require.Equal(t, output.IsSpent(), false)
	require.Equal(t, output.IsSpent(), true)
	require.Equal(t, output.IsSpent(), true)
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

	global_mock.On("UxOut", h.String()).Return(nil, goerrors.New("failure")).Once()
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
	require.Nil(t, tiList)

	tiList = ut.GetInputs()
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

func TestTransactionsGetTimestamp(t *testing.T) {
	cur := time.Now()
	tests := []struct {
		name string
		txn  core.Transaction
		want uint64
	}{
		{
			name: "SkycoinUninjectTransaction",
			txn:  new(SkycoinUninjectedTransaction),
			want: 0,
		},
		{
			name: "SkycoinTransaction",
			txn: &SkycoinTransaction{
				skyTxn: readable.TransactionVerbose{
					Timestamp: 42,
				},
			},
			want: 42,
		},
		{
			name: "SkycoinPendingTransaction",
			txn: &SkycoinPendingTransaction{
				Transaction: &readable.UnconfirmedTransactionVerbose{
					Received: cur,
				},
			},
			want: uint64(cur.Unix()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, core.Timestamp(tt.want), tt.txn.GetTimestamp())
		})
	}
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

func Test_verifyReadableTransaction(t *testing.T) {
	mockTxn := new(mocks.ReadableTxn)
	id := "0000000000000000000000000000000000000000000000000000000000000000"
	created := &api.CreatedTransaction{
		TxID:      "78877fa898f0b4c45c9c33ae941e40617ad7c8657a307db62bc5691f92f4f60e",
		InnerHash: "No-match",
	}
	mockTxn.On("ToCreatedTransaction").Return(nil, goerrors.New("failure")).Once()
	mockTxn.On("ToCreatedTransaction").Return(created, nil)

	// ReadableTxn error
	err := verifyReadableTransaction(mockTxn, false)
	require.Error(t, err)

	// transaction hash error
	err = verifyReadableTransaction(mockTxn, false)
	require.Error(t, err)

	// Verify fail
	created.InnerHash = id
	err = verifyReadableTransaction(mockTxn, false)
	require.Error(t, err)

	// VerifyUnsigned fail
	err = verifyReadableTransaction(mockTxn, true)
	require.Error(t, err)

	//TODO: add a case that not raise an error
}

func TestPendingTxnVerifySignature(t *testing.T) {
	tests := []struct {
		name string
		sTxn *SkycoinPendingTransaction
	}{
		{
			name: "empty transaction",
			sTxn: &SkycoinPendingTransaction{Transaction: new(readable.UnconfirmedTransactionVerbose)},
		},
		{
			name: "cero outputs",
			sTxn: &SkycoinPendingTransaction{
				Transaction: &readable.UnconfirmedTransactionVerbose{
					Transaction: readable.BlockTransactionVerbose{
						In: []readable.TransactionInput{readable.TransactionInput{}},
					},
				},
			},
		},
		//TODO: add valid tests
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.sTxn.Transaction.IsValid {
				require.Error(t, tt.sTxn.VerifySigned())
				require.Error(t, tt.sTxn.VerifyUnsigned())
				tt.sTxn.Transaction.IsValid = true
			}
			require.Equal(t, verifyReadableTransaction(tt.sTxn, false), tt.sTxn.VerifyUnsigned())
			require.Equal(t, verifyReadableTransaction(tt.sTxn, true), tt.sTxn.VerifySigned())
		})
	}
}

func Test_checkFullySigned(t *testing.T) {
	mockTxn := new(mocks.ReadableTxn)
	id := "0000000000000000000000000000000000000000000000000000000000000000"
	created := &api.CreatedTransaction{
		TxID:      "78877fa898f0b4c45c9c33ae941e40617ad7c8657a307db62bc5691f92f4f60e",
		InnerHash: "No-match",
	}
	mockTxn.On("ToCreatedTransaction").Return(nil, goerrors.New("failure")).Once()
	mockTxn.On("ToCreatedTransaction").Return(created, nil)

	_, err := checkFullySigned(mockTxn)
	require.Error(t, err)

	_, err = checkFullySigned(mockTxn)
	require.Error(t, err)

	// false
	created.InnerHash = id
	val, err := checkFullySigned(mockTxn)
	require.NoError(t, err)
	require.False(t, val)

	//TODO: add valid test
}

func TestTransactionIsFullySigned(t *testing.T) {
	type skyTxn interface {
		skytypes.ReadableTxn
		core.Transaction
	}
	tests := []struct {
		name string
		txn  skyTxn
	}{
		{
			name: "empty PendingTxn",
			txn:  &SkycoinPendingTransaction{Transaction: new(readable.UnconfirmedTransactionVerbose)},
		},
		{
			name: "cero outputs in PendingTxn",
			txn: &SkycoinPendingTransaction{
				Transaction: &readable.UnconfirmedTransactionVerbose{
					Transaction: readable.BlockTransactionVerbose{
						In: []readable.TransactionInput{readable.TransactionInput{}},
					},
				},
			},
		},
		{
			name: "empty Txn",
			txn:  new(SkycoinTransaction),
		},
		//TODO: add valid tests
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fully, err := checkFullySigned(tt.txn)
			fully1, err1 := tt.txn.IsFullySigned()
			require.Equal(t, err, err1)
			require.Equal(t, fully, fully1)
		})
	}
}

func TestSkycoinUninjectedTransactionGetOutputs(t *testing.T) {
	addr := makeAddress()
	addr1 := makeAddress()
	tests := []struct {
		name    string
		ujTxn   *SkycoinUninjectedTransaction
		addrs   []string
		wantNil bool
	}{
		{
			name: "empty transaction",
			ujTxn: &SkycoinUninjectedTransaction{
				outputs: []core.TransactionOutput{},
			},
			addrs: make([]string, 0),
		},
		{
			name: "nil outputs",
			ujTxn: &SkycoinUninjectedTransaction{
				txn: &coin.Transaction{},
			},
			addrs: make([]string, 0),
		},
		{
			name: "incorrect output",
			ujTxn: &SkycoinUninjectedTransaction{
				txn: &coin.Transaction{
					Out: []coin.TransactionOutput{
						coin.TransactionOutput{
							Coins: 9223372036854775808,
						},
					},
				},
			},
			wantNil: true,
		},
		{
			name: "some outputs",
			ujTxn: &SkycoinUninjectedTransaction{
				txn: &coin.Transaction{
					Out: []coin.TransactionOutput{
						coin.TransactionOutput{
							Address: addr,
						},
						coin.TransactionOutput{
							Address: addr1,
						},
					},
				},
			},
			addrs: []string{addr.String(), addr1.String()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputs := tt.ujTxn.GetOutputs()
			if tt.wantNil {
				require.Nil(t, outputs)
			} else {
				require.Equal(t, len(tt.addrs), len(outputs))
				hashes := make([]string, len(tt.addrs))
				for i, out := range outputs {
					hashes[i] = out.GetAddress().String()
				}
				requirethat.ElementsMatch(t, tt.addrs, hashes)
			}
		})
	}
}

func TestGetId(t *testing.T) {
	type ObjectWithID interface {
		GetId() string
	}
	tests := []struct {
		obj  ObjectWithID
		want string
	}{
		{
			obj: &SkycoinUninjectedTransaction{
				txn: new(coin.Transaction),
			},
			want: "78877fa898f0b4c45c9c33ae941e40617ad7c8657a307db62bc5691f92f4f60e",
		},
		{
			obj: &SkycoinUninjectedTransaction{
				txn: &coin.Transaction{
					Length: 5,
				},
			},
			want: "03e228f59704bc30de09f76fe9db0981ca77b6421aaa997e227d39bcc317174e",
		},
		{
			obj: &SkycoinUninjectedTransaction{
				txn: &coin.Transaction{
					Type: 2,
				},
			},
			want: "bb5e828965130b51e627725f6fea3247124da6799d28ccac81c247fd78b34621",
		},
		{
			obj: &SkycoinPendingTransaction{
				Transaction: &readable.UnconfirmedTransactionVerbose{
					Transaction: readable.BlockTransactionVerbose{
						Hash: "hash1",
					},
				},
			},
			want: "hash1",
		},
		{
			obj: &SkycoinPendingTransaction{
				Transaction: &readable.UnconfirmedTransactionVerbose{
					Transaction: readable.BlockTransactionVerbose{
						Hash: "hash2",
					},
				},
			},
			want: "hash2",
		},
		{
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					UxID: "uxid1",
				},
			},
			want: "uxid1",
		},
		{
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					UxID: "uxid2",
				},
			},
			want: "uxid2",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("ID%d", i), func(t *testing.T) {
			require.Equal(t, tt.want, tt.obj.GetId())
		})
	}
}

func TestSkycoinTransactionGetOutputs(t *testing.T) {
	skyAmount := uint64(20000000)
	chAmount := uint64(20)

	st := new(SkycoinTransaction)
	outs := []readable.TransactionOutput{
		readable.TransactionOutput{
			Hash:  "O1",
			Coins: "20",
			Hours: chAmount,
		},
		readable.TransactionOutput{
			Hash:  "O2",
			Coins: "20",
			Hours: chAmount,
		},
	}
	st.skyTxn.Out = outs

	tests := []struct {
		name    string
		txn     core.Transaction
		ids     []string
		wantNil bool
	}{
		{
			name: "SkycoinTransaction1",
			txn:  st,
			ids:  []string{"O1", "O2"},
		},
		{
			name: "SkycoinTransaction1-OutputsSaved",
			txn:  st,
			ids:  []string{"O1", "O2"},
		},
		{
			name: "SkycoinTransaction1-OutputsSaved",
			txn: &SkycoinTransaction{
				skyTxn: readable.TransactionVerbose{
					BlockTransactionVerbose: readable.BlockTransactionVerbose{
						Out: outs[:1],
					},
				},
			},
			ids: []string{"O1"},
		},
		{
			name: "SkycoinTransaction1-NoOutputs",
			txn:  new(SkycoinTransaction),
			ids:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputs := tt.txn.GetOutputs()
			if tt.wantNil {
				require.Nil(t, outputs)
			} else {
				ids := make([]string, 0)
				it := NewSkycoinTransactionOutputIterator(outputs)
				for it.Next() {
					output := it.Value()
					ids = append(ids, output.GetId())
					sky, err := output.GetCoins(Sky)
					require.NoError(t, err)
					require.Equal(t, skyAmount, sky)
					hours, err1 := output.GetCoins(CoinHour)
					require.NoError(t, err1)
					require.Equal(t, chAmount, hours)
				}
				requirethat.ElementsMatch(t, tt.ids, ids)
			}
		})
	}
}

func TestTransactionComputeFee(t *testing.T) {
	expectedError := func(ticker string) error {
		if ticker == CalculatedHour {
			return errors.ErrNotImplemented
		}
		if ticker == Sky || ticker == CoinHour {
			return nil
		}
		return errors.ErrInvalidAltcoinTicker
	}
	expectedAmount := func(ticker string, amount uint64) uint64 {
		if ticker == CoinHour {
			return amount
		}
		return 0
	}
	pendingTxnWithFee := func(ticker string, fee string) (core.Transaction, uint64, bool, error) {
		val, err := strconv.ParseUint(fee, 10, 64)
		if err != nil {
			return nil, 0, false, nil
		}
		return &SkycoinPendingTransaction{
			Transaction: &readable.UnconfirmedTransactionVerbose{
				Transaction: readable.BlockTransactionVerbose{
					Fee: val,
				},
			},
		}, expectedAmount(ticker, val), true, expectedError(ticker)
	}
	uninjectedTxnWithFee := func(ticker string, fee string) (core.Transaction, uint64, bool, error) {
		val, err := strconv.ParseUint(fee, 10, 64)
		if err != nil {
			return nil, 0, false, nil
		}
		return &SkycoinUninjectedTransaction{fee: val}, expectedAmount(ticker, val), true, expectedError(ticker)
	}
	skycoinTxnWithFee := func(ticker string, fee string) (core.Transaction, uint64, bool, error) {
		val, err := strconv.ParseUint(fee, 10, 64)
		if err != nil {
			return nil, 0, false, nil
		}
		return &SkycoinTransaction{
			skyTxn: readable.TransactionVerbose{
				BlockTransactionVerbose: readable.BlockTransactionVerbose{
					Fee: val,
				},
			},
		}, expectedAmount(ticker, val), true, expectedError(ticker)
	}
	createdTxnWithFee := func(ticker string, fee string) (core.Transaction, uint64, bool, error) {
		val, err := strconv.ParseInt(fee, 10, 64)
		var expError error
		if err != nil && ticker == CoinHour {
			expError = err
		} else {
			expError = expectedError(ticker)
		}
		return &SkycoinCreatedTransaction{
			skyTxn: api.CreatedTransaction{
				Fee: fee,
			},
		}, expectedAmount(ticker, uint64(val)), true, expError
	}

	tests := []struct {
		name      string
		generator func(string, string) (core.Transaction, uint64, bool, error)
	}{
		{
			name:      "SkycoinPendingTransaction",
			generator: pendingTxnWithFee,
		},
		{
			name:      "SkycoinUninjectedTransaction",
			generator: uninjectedTxnWithFee,
		},
		{
			name:      "SkycoinTransaction",
			generator: skycoinTxnWithFee,
		},
		{
			name:      "SkycoinCreatedTransaction",
			generator: createdTxnWithFee,
		},
	}
	tickers := []string{Sky, CoinHour, CalculatedHour, "INVALIDTICKER"}
	amounts := []string{"1", "2", "42", "100", "42,42", "1,1"}

	for _, tt := range tests {
		for _, ticker := range tickers {
			for _, amount := range amounts {
				t.Run(tt.name+"-"+ticker, func(t *testing.T) {
					thx, expected, valid, err := tt.generator(ticker, amount)
					if valid {
						val, err1 := thx.ComputeFee(ticker)
						if err != nil {
							require.Equal(t, err, err1)
						} else {
							require.NoError(t, err1)
							require.Equal(t, expected, val)
						}
					}
				})
			}
		}
	}
}

func TestTransactionVerifyUnsigned(t *testing.T) {
	type skyTxn interface {
		skytypes.ReadableTxn
		core.Transaction
	}
	tests := []struct {
		name string
		txn  skyTxn
	}{
		{
			name: "empty Txn",
			txn:  new(SkycoinTransaction),
		},
		//TODO: add better tests
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyReadableTransaction(tt.txn, false)
			err1 := tt.txn.VerifyUnsigned()
			require.Equal(t, err, err1)
		})
	}
}

func TestTransactionVerifySigned(t *testing.T) {
	type skyTxn interface {
		skytypes.ReadableTxn
		core.Transaction
	}
	tests := []struct {
		name string
		txn  skyTxn
	}{
		{
			name: "empty Txn",
			txn:  new(SkycoinTransaction),
		},
		//TODO: add better tests
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyReadableTransaction(tt.txn, true)
			err1 := tt.txn.VerifySigned()
			require.Equal(t, err, err1)
		})
	}
}

func Test_getSkycoinTransactionInputsFromTxnHash(t *testing.T) {
	CleanGlobalMock()

	inputs := []readable.TransactionInput{
		readable.TransactionInput{
			Hash: "hash1",
		},
		readable.TransactionInput{
			Hash: "hash2",
		},
		readable.TransactionInput{
			Hash: "hash3",
		},
		readable.TransactionInput{
			Hash: "hash4",
		},
	}

	for len(inputs) > 0 {
		global_mock.On("TransactionVerbose", "hash").Return(
			&readable.TransactionWithStatusVerbose{
				Transaction: readable.TransactionVerbose{
					BlockTransactionVerbose: readable.BlockTransactionVerbose{
						In: inputs,
					},
				},
			}, nil,
		).Once()
		t.Run("InputsFromTxnHash", func(t *testing.T) {
			txnInputs, err := getSkycoinTransactionInputsFromTxnHash("hash")
			require.NoError(t, err)
			rawInputs := make([]readable.TransactionInput, len(txnInputs))
			for i, in := range txnInputs {
				skyIn, valid := in.(*SkycoinTransactionInput)
				require.True(t, valid)
				require.Nil(t, skyIn.spentOutput)
				rawInputs[i] = skyIn.skyIn
			}
			requirethat.ElementsMatch(t, inputs, rawInputs)
		})
		inputs = inputs[:len(inputs)-1]
	}
	global_mock.On("TransactionVerbose", "hash").Return(nil, goerrors.New("failure"))
	_, err := getSkycoinTransactionInputsFromTxnHash("hash")
	require.Error(t, err)
}

func TestGetCoins(t *testing.T) {
	type ObjectWithCoins interface {
		GetCoins(string) (uint64, error)
	}

	invalidTicker := "INVALIDTICKER"
	tests := []struct {
		name   string
		obj    ObjectWithCoins
		ticker string
		want   uint64
		err    bool
	}{
		// SkycoinCreatedTransactionOutput
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: invalidTicker,
			obj:    new(SkycoinCreatedTransactionOutput),
			err:    true,
		},
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: Sky,
			obj: &SkycoinCreatedTransactionOutput{
				skyOut: api.CreatedTransactionOutput{
					Coins: "20",
				},
			},
			want: 20000000,
		},
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: Sky,
			obj: &SkycoinCreatedTransactionOutput{
				skyOut: api.CreatedTransactionOutput{
					Coins: "20.1",
				},
			},
			want: 20100000,
		},
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: Sky,
			obj: &SkycoinCreatedTransactionOutput{
				skyOut: api.CreatedTransactionOutput{
					Coins: "20,1a",
				},
			},
			err: true,
		},
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: CoinHour,
			obj: &SkycoinCreatedTransactionOutput{
				skyOut: api.CreatedTransactionOutput{
					Hours: "42",
				},
			},
			want: 42,
		},
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: CoinHour,
			obj: &SkycoinCreatedTransactionOutput{
				skyOut: api.CreatedTransactionOutput{
					Hours: "42.1",
				},
			},
			err: true,
		},
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: CalculatedHour,
			obj: &SkycoinCreatedTransactionOutput{
				calculatedHours: 42,
			},
			want: 42,
		},
		{
			name:   "SkycoinCreatedTransactionOutput",
			ticker: CalculatedHour,
			obj: &SkycoinCreatedTransactionOutput{
				calculatedHours: 50,
			},
			want: 50,
		},
		// SkycoinCreatedTransactionInput
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: invalidTicker,
			obj:    new(SkycoinCreatedTransactionInput),
			err:    true,
		},
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: Sky,
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					Coins: "20",
				},
			},
			want: 20000000,
		},
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: Sky,
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					Coins: "20.1",
				},
			},
			want: 20100000,
		},
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: Sky,
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					Coins: "20,1a",
				},
			},
			err: true,
		},
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: CoinHour,
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					Hours: "42",
				},
			},
			want: 42,
		},
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: CoinHour,
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					Hours: "42.1",
				},
			},
			err: true,
		},
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: CalculatedHour,
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					CalculatedHours: "42",
				},
			},
			want: 42,
		},
		{
			name:   "SkycoinCreatedTransactionInput",
			ticker: CalculatedHour,
			obj: &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					CalculatedHours: "42.1",
				},
			},
			err: true,
		},
		// SkycoinTransactionInput
		{
			name:   "SkycoinTransactionInput",
			ticker: invalidTicker,
			obj:    new(SkycoinTransactionInput),
			err:    true,
		},
		{
			name:   "SkycoinTransactionInput",
			ticker: Sky,
			obj: &SkycoinTransactionInput{
				skyIn: readable.TransactionInput{
					Coins: "20",
				},
			},
			want: 20000000,
		},
		{
			name:   "SkycoinTransactionInput",
			ticker: Sky,
			obj: &SkycoinTransactionInput{
				skyIn: readable.TransactionInput{
					Coins: "20.1",
				},
			},
			want: 20100000,
		},
		{
			name:   "SkycoinTransactionInput",
			ticker: Sky,
			obj: &SkycoinTransactionInput{
				skyIn: readable.TransactionInput{
					Coins: "20,1a",
				},
			},
			err: true,
		},
		{
			name:   "SkycoinTransactionInput",
			ticker: CoinHour,
			obj: &SkycoinTransactionInput{
				skyIn: readable.TransactionInput{
					Hours: 42,
				},
			},
			want: 42,
		},
		{
			name:   "SkycoinTransactionInput",
			ticker: CalculatedHour,
			obj: &SkycoinTransactionInput{
				skyIn: readable.TransactionInput{
					CalculatedHours: 42,
				},
			},
			want: 42,
		},
		// TransactionOutput
		{
			name:   "SkycoinTransactionOutput",
			ticker: invalidTicker,
			obj:    new(SkycoinTransactionOutput),
			err:    true,
		},
		{
			name:   "SkycoinTransactionOutput",
			ticker: Sky,
			obj: &SkycoinTransactionOutput{
				skyOut: readable.TransactionOutput{
					Coins: "20",
				},
			},
			want: 20000000,
		},
		{
			name:   "SkycoinTransactionOutput",
			ticker: Sky,
			obj: &SkycoinTransactionOutput{
				skyOut: readable.TransactionOutput{
					Coins: "20.1",
				},
			},
			want: 20100000,
		},
		{
			name:   "SkycoinTransactionOutput",
			ticker: Sky,
			obj: &SkycoinTransactionOutput{
				skyOut: readable.TransactionOutput{
					Coins: "20,1a",
				},
			},
			err: true,
		},
		{
			name:   "SkycoinTransactionOutput",
			ticker: CoinHour,
			obj: &SkycoinTransactionOutput{
				skyOut: readable.TransactionOutput{
					Hours: 42,
				},
			},
			want: 42,
		},
		{
			name:   "SkycoinTransactionOutput",
			ticker: CalculatedHour,
			obj: &SkycoinTransactionOutput{
				calculatedHours: 42,
			},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"-"+tt.ticker, func(t *testing.T) {
			amount, err := tt.obj.GetCoins(tt.ticker)
			if tt.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, amount)
			}
		})
	}
}

func TestSkycoinTransactionOutputGetAddress(t *testing.T) {
	strAddr := makeAddress().String()

	tests := []struct {
		name   string
		output core.TransactionOutput
		err    bool
		want   string
		isNil  bool
	}{
		{
			name:   "SkycoinTransactionOutput-empty",
			output: new(SkycoinTransactionOutput),
			isNil:  true,
		},
		{
			name: "SkycoinTransactionOutput-empty",
			output: &SkycoinTransactionOutput{
				skyOut: readable.TransactionOutput{
					Address: strAddr,
				},
			},
			want: strAddr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr := tt.output.GetAddress()
			if tt.isNil {
				require.Nil(t, addr)
			} else {
				require.NotNil(t, addr)
				require.Equal(t, tt.want, addr.String())
			}
		})
	}
}

func TestSkycoinCreatedTransactionInputGetSpentOutput(t *testing.T) {
	tests := []struct {
		addr    string
		coins   string
		hours   string
		uxid    string
		ccHours string
	}{
		{addr: makeAddress().String(), coins: "1", hours: "1", uxid: "uxid1", ccHours: "1"},
		{addr: makeAddress().String(), coins: "2", hours: "2", uxid: "uxid2", ccHours: "2"},
		{addr: makeAddress().String(), coins: "3", hours: "3", uxid: "uxid3", ccHours: "3"},
		{addr: makeAddress().String(), coins: "4", hours: "4", uxid: "uxid4", ccHours: "4"},
		{addr: makeAddress().String(), coins: "5", hours: "5", uxid: "uxid5", ccHours: "5"},
		{addr: makeAddress().String(), coins: "6", hours: "6", uxid: "uxid6", ccHours: "6"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("GetSepent%d", i), func(t *testing.T) {
			createdIn := &SkycoinCreatedTransactionInput{
				skyIn: api.CreatedTransactionInput{
					Address:         tt.addr,
					Coins:           tt.coins,
					Hours:           tt.hours,
					UxID:            tt.uxid,
					CalculatedHours: tt.ccHours,
				},
			}
			output := createdIn.GetSpentOutput()
			createdOut, valid := output.(*SkycoinCreatedTransactionOutput)
			require.True(t, valid)
			for _, asset := range createdIn.SupportedAssets() {
				cur, err := createdIn.GetCoins(asset)
				require.NoError(t, err)
				val, err := createdOut.GetCoins(asset)
				require.NoError(t, err)
				require.Equal(t, cur, val)
			}
			require.Equal(t, createdIn.GetId(), createdOut.GetId())
			require.Equal(t, tt.addr, createdOut.GetAddress().String())
		})
	}
}

func Test_newCreatedTransactionOutputs(t *testing.T) {
	outputs := []api.CreatedTransactionOutput{
		api.CreatedTransactionOutput{
			Address: "addr1",
		},
		api.CreatedTransactionOutput{
			Address: "addr2",
		},
		api.CreatedTransactionOutput{
			Address: "addr3",
		},
		api.CreatedTransactionOutput{
			Address: "addr4",
		},
		api.CreatedTransactionOutput{
			Address: "addr5",
		},
	}

	for len(outputs) > 0 {
		outs := newCreatedTransactionOutputs(outputs)
		rawOutputs := make([]api.CreatedTransactionOutput, len(outputs))
		for i, out := range outs {
			createdOut, valid := out.(*SkycoinCreatedTransactionOutput)
			require.True(t, valid)
			rawOutputs[i] = createdOut.skyOut
		}
		requirethat.ElementsMatch(t, outputs, rawOutputs)
		outputs = outputs[:len(outputs)-1]
	}
}
