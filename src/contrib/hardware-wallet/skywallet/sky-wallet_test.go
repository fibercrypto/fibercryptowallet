package hardware

import (
	"errors"
	"github.com/chebyrash/promise"
	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	fce "github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
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

func TestShouldFailForFailResponse(t *testing.T) {
	t.Skip("TODO: This test should be moved to SkyWalletInteraction")
	// Giving
	dev := createMockedDeviceInteraction(t)
	f := messages.Failure{}
	fb, err := proto.Marshal(&f)
	require.Nil(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Failure),
		Data: fb,
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
	require.Equal(t, "", devId)
}

func TestGetSignerDescriptionShouldFailForInvalidMessageType(t *testing.T) {
	t.Skip("TODO: This test should be moved to SkyWalletInteraction")
	// Giving
	dev := createMockedDeviceInteraction(t)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_PinMatrixAck),
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
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

func TestGetSignerUIDShouldFailForUninitializedDevice(t *testing.T) {
	t.Skip("TODO: This test should be moved to SkyWalletInteraction")
	// Giving
	createMockedDeviceInteraction(t)
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, err, fce.ErrHwUnexpected)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerUIDShouldFailForFailResponse(t *testing.T) {
	t.Skip("TODO: This test should be moved to SkyWalletInteraction")
	// Giving
	dev := createMockedDeviceInteraction(t)
	f := messages.Failure{}
	fb, err := proto.Marshal(&f)
	require.Nil(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Failure),
		Data: fb,
	}
	dev.On("Features").Return(msg, nil)
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerUIDShouldFailForInvalidMessageType(t *testing.T) {
	t.Skip("TODO: This test should be moved to SkyWalletInteraction")
	// Giving
	dev := createMockedDeviceInteraction(t)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_PinMatrixAck),
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerDescriptionShouldFailForUninitializedDevice(t *testing.T) {
	t.Skip("lllllllll7")
	// Giving
	createMockedDeviceInteraction(t)
	sw := SkyWallet{}

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
	require.Equal(t, "", devId)
}