package hardware

import (
	"errors"
	"github.com/chebyrash/promise"
	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	fce "github.com/fibercrypto/fibercryptowallet/src/errors"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func createMockedDeviceInteraction(t *testing.T) *mocks.DeviceInteraction {
	skyWltInteraction = &mocks.DeviceInteraction{}
	require.NotNil(t, SkyWltInteractionInstance())
	require.Equal(t, skyWltInteraction, SkyWltInteractionInstance())
	gi, ok := skyWltInteraction.(*mocks.DeviceInteraction)
	require.True(t, ok)
	return gi
}

func TestGetSignerUIDShouldBeOk(t *testing.T) {
	// Giving
	dev := createMockedDeviceInteraction(t)
	expectedDevId := "c24a046d-7d7b-484d-baf6-abedd883023f"
	dev.On("Features").Return(promise.New(func(resolve func(interface{}), reject func(error)) {
		f := messages.Features{
			DeviceId: proto.String(expectedDevId),
		}
		resolve(f)
	}), nil)
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.NoError(t, err)
	require.Equal(t, core.UID(expectedDevId), devId)
}

func TestGetSignerUIDShouldFailOnNullId(t *testing.T) {
	// Giving
	dev := createMockedDeviceInteraction(t)
	dev.On("Features").Return(promise.New(func(resolve func(interface{}), reject func(error)) {
		f := messages.Features{
			DeviceId: nil,
		}
		resolve(f)
	}), nil)
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Equal(t, err, fce.ErrNilValue)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerUIDShouldFailOnDeviceError(t *testing.T) {
	// Giving
	dev := createMockedDeviceInteraction(t)
	dev.On("Features").Return(promise.New(func(resolve func(interface{}), reject func(error)) {
		reject(errors.New("any error"))
	}))
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerDescriptionShouldBeOk(t *testing.T) {
	// Giving
	dev := createMockedDeviceInteraction(t)
	expectedDevDescription := urnPrefix +"c24a046d-7d7b-484d-baf6-abedd883023f"
	dev.On("Features").Return(promise.New(func(resolve func(interface{}), reject func(error)) {
		f := messages.Features{
			Label: proto.String("c24a046d-7d7b-484d-baf6-abedd883023f"),
		}
		resolve(f)
	}))
	sw := NewSkyWallet(nil)

	// When
	devDescription, err := sw.GetSignerDescription()

	// Then
	require.NoError(t, err)
	require.Equal(t, expectedDevDescription, devDescription)
}

func TestGetSignerDescriptionShouldFailOnDeviceError(t *testing.T) {
	// Giving
	dev := createMockedDeviceInteraction(t)
	dev.On("Features").Return(promise.New(func(resolve func(interface{}), reject func(error)) {
		reject(errors.New(""))
	}))
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
	require.Equal(t, "", devId)
}

func TestGetSignerDescriptionShouldFailOnNullLabel(t *testing.T) {
	// Giving
	dev := createMockedDeviceInteraction(t)
	dev.On("Features").Return(promise.New(func(resolve func(interface{}), reject func(error)) {
		f := messages.Features{
			Label: nil,
		}
		resolve(f)
	}))
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Equal(t, err, fce.ErrNilValue)
	require.Equal(t, "", devId)
}

func TestGetFeaturesShouldHandleInvalidMsgResponse(t *testing.T) {
	// Giving
	dev := createMockedDeviceInteraction(t)
	dev.On("Features").Return(promise.New(func(resolve func(interface{}), reject func(error)) {
		f := messages.ResponseTransactionSign{}
		resolve(f)
	}))
	sw := NewSkyWallet(nil)

	// When
	_, err := sw.getDeviceFeatures()

	// Then
	require.Error(t, err)
}

func TestVerifyDerivationTypeBip44(t *testing.T) {
	// Giving
	addr := &mocks.Address{}
	addr.On("IsBip32").Return(true)
	output := &mocks.TransactionOutput{}
	output.On("GetAddress").Return(addr, nil)
	input := &mocks.TransactionInput{}
	input.On("GetSpentOutput").Return(output, nil)
	inputs := []core.TransactionInput{input}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)

	// When
	dt, err := derivationType(txn)

	// Then
	require.NoError(t, err)
	require.Equal(t, dt, skyWallet.WalletTypeBip44)
}

func TestVerifyDerivationTypeDeterministic(t *testing.T) {
	// Giving
	addr := &mocks.Address{}
	addr.On("IsBip32").Return(false)
	output := &mocks.TransactionOutput{}
	output.On("GetAddress").Return(addr, nil)
	input := &mocks.TransactionInput{}
	input.On("GetSpentOutput").Return(output, nil)
	inputs := []core.TransactionInput{input}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)

	// When
	dt, err := derivationType(txn)

	// Then
	require.NoError(t, err)
	require.Equal(t, dt, skyWallet.WalletTypeDeterministic)
}

func TestVerifyInputsGroupingShouldWorkOk1(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr2 := &mocks.Address{}
	addr1.On("IsBip32").Return(true)
	addr2.On("IsBip32").Return(true)
	output1 := &mocks.TransactionOutput{}
	output2 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	output2.On("GetAddress").Return(addr2, nil)
	input1 := &mocks.TransactionInput{}
	input2 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	input2.On("GetSpentOutput").Return(output2, nil)
	inputs := []core.TransactionInput{input1, input2}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)

	// When
	err := verifyInputsGrouping(txn)

	// Then
	require.NoError(t, err)
}

func TestVerifyInputsGroupingShouldWorkOk2(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr2 := &mocks.Address{}
	addr1.On("IsBip32").Return(false)
	addr2.On("IsBip32").Return(false)
	output1 := &mocks.TransactionOutput{}
	output2 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	output2.On("GetAddress").Return(addr2, nil)
	input1 := &mocks.TransactionInput{}
	input2 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	input2.On("GetSpentOutput").Return(output2, nil)
	inputs := []core.TransactionInput{input1, input2}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)

	// When
	err := verifyInputsGrouping(txn)

	// Then
	require.NoError(t, err)
}

func TestVerifyInputsGroupingShouldDetectErr1(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr2 := &mocks.Address{}
	addr1.On("IsBip32").Return(false)
	addr2.On("IsBip32").Return(true)
	output1 := &mocks.TransactionOutput{}
	output2 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	output2.On("GetAddress").Return(addr2, nil)
	input1 := &mocks.TransactionInput{}
	input2 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	input2.On("GetSpentOutput").Return(output2, nil)
	inputs := []core.TransactionInput{input1, input2}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)

	// When
	err := verifyInputsGrouping(txn)

	// Then
	require.Error(t, err)
}

func TestVerifyInputsGroupingShouldDetectErr2(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr2 := &mocks.Address{}
	addr1.On("IsBip32").Return(true)
	addr2.On("IsBip32").Return(false)
	output1 := &mocks.TransactionOutput{}
	output2 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	output2.On("GetAddress").Return(addr2, nil)
	input1 := &mocks.TransactionInput{}
	input2 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	input2.On("GetSpentOutput").Return(output2, nil)
	inputs := []core.TransactionInput{input1, input2}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)

	// When
	err := verifyInputsGrouping(txn)

	// Then
	require.Error(t, err)
}

func TestSignTransactionShouldDetectInvalidInputsGrouping(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr2 := &mocks.Address{}
	addr1.On("IsBip32").Return(true)
	addr2.On("IsBip32").Return(false)
	output1 := &mocks.TransactionOutput{}
	output2 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	output2.On("GetAddress").Return(addr2, nil)
	input1 := &mocks.TransactionInput{}
	input2 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	input2.On("GetSpentOutput").Return(output2, nil)
	inputs := []core.TransactionInput{input1, input2}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)
	wlt := NewSkyWallet(&mocks.Wallet{})

	// When
	_, err := wlt.SignTransaction(
		txn, func(string, core.KeyValueStore) (string, error){return "", nil}, []string{})

	// Then
	require.Error(t, err)
	require.Equal(t, fce.ErrTxnSignFailure, err)
}

func TestSignTransactionShouldHandleErrorFromIsFullySigned(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr1.On("IsBip32").Return(true)
	output1 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	input1 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	inputs := []core.TransactionInput{input1}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)
	txn.On("IsFullySigned").Return(false, errors.New("asd"))
	wlt := NewSkyWallet(&mocks.Wallet{})

	// When
	_, err := wlt.SignTransaction(
		txn, func(string, core.KeyValueStore) (string, error){return "", nil}, []string{})

	// Then
	require.Error(t, err)
	require.Equal(t, errors.New("asd"), err)
}

func TestSignTransactionShouldHandleFailIfIsFullySigned(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr1.On("IsBip32").Return(true)
	output1 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	input1 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	inputs := []core.TransactionInput{input1}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)
	txn.On("IsFullySigned").Return(true, nil)
	wlt := NewSkyWallet(&mocks.Wallet{})

	// When
	_, err := wlt.SignTransaction(
		txn, func(string, core.KeyValueStore) (string, error){return "", nil}, []string{})

	// Then
	require.Error(t, err)
	require.Equal(t, errors.New("Transaction is fully signed"), err)
}

func TestSignTransactionShouldHandleErrorFromGetHashIndices(t *testing.T) {
	//t.Skip("ddd")
	// Giving
	addr1 := &mocks.Address{}
	addr1.On("IsBip32").Return(true)
	output1 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	input1 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	inputs := []core.TransactionInput{input1}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)
	txn.On("IsFullySigned").Return(false, nil)
	wlt := NewSkyWallet(&mocks.Wallet{})

	// When
	_, err := wlt.SignTransaction(
		txn, func(string, core.KeyValueStore) (string, error){return "", nil}, []string{"#asd"})

	// Then
	require.Error(t, err)
	require.Equal(t, fce.ErrTxnSignFailure, err)
}

func TestSignTransactionShouldHandleErrorFromsignTransaction(t *testing.T) {
	// Giving
	addr1 := &mocks.Address{}
	addr1.On("IsBip32").Return(true)
	output1 := &mocks.TransactionOutput{}
	output1.On("GetAddress").Return(addr1, nil)
	input1 := &mocks.TransactionInput{}
	input1.On("GetSpentOutput").Return(output1, nil)
	inputs := []core.TransactionInput{input1}
	txn := &mocks.Transaction{}
	txn.On("GetInputs").Return(inputs)
	txn.On("IsFullySigned").Return(false, nil)
	txn.On("ComputeFee", mock.Anything).Return(uint64(0), errors.New("error computing fee"))
	wlt := NewSkyWallet(&mocks.Wallet{})

	// When
	_, err := wlt.SignTransaction(
		txn, func(string, core.KeyValueStore) (string, error){return "", nil}, nil)

	// Then
	require.Error(t, err)
	require.Equal(t, fce.ErrTxnSignFailure, err)
}
