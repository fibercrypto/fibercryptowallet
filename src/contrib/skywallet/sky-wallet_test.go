package hardware

import (
	"errors"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	fce "github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/contrib/skywallet/mocks"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"testing"
	messages "github.com/fibercrypto/skywallet-protob/go"
)



func TestGetSignerUIDShouldBeOk(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	expectedDevId := "c24a046d-7d7b-484d-baf6-abedd883023f"
	f := messages.Features{
		DeviceId: proto.String(expectedDevId),
	}
	fb, err := proto.Marshal(&f)
	require.Nil(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Features),
		Data: fb,
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil, &dev)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.NoError(t, err)
	require.Equal(t, core.UID(expectedDevId), devId)
}

func TestGetSignerUIDShouldFailForUninitializedDevice(t *testing.T) {
	// Giving
	sw := NewSkyWallet(nil, nil)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, err, fce.ErrHwUnexpected)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerUIDShouldFailOnDeviceError(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	dev.On("GetFeatures").Return(wire.Message{}, errors.New(""))
	sw := NewSkyWallet(nil, &dev)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerUIDShouldFailForFailResponse(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	f := messages.Failure{}
	fb, err := proto.Marshal(&f)
	require.Nil(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Failure),
		Data: fb,
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil, &dev)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerUIDShouldFailForInvalidMessageType(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_PinMatrixAck),
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil, &dev)

	// When
	devId, err := sw.GetSignerUID()

	// Then
	require.Error(t, err)
	require.Equal(t, core.UID(""), devId)
}

func TestGetSignerDescriptionShouldBeOk(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	expectedDevDescription := urnPrefix+"c24a046d-7d7b-484d-baf6-abedd883023f"
	f := messages.Features{
		Label: proto.String("c24a046d-7d7b-484d-baf6-abedd883023f"),
	}
	fb, err := proto.Marshal(&f)
	require.Nil(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Features),
		Data: fb,
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil, &dev)

	// When
	devDescription, err := sw.GetSignerDescription()

	// Then
	require.NoError(t, err)
	require.Equal(t, expectedDevDescription, devDescription)
}

func TestGetSignerDescriptionShouldFailForUninitializedDevice(t *testing.T) {
	// Giving
	sw := SkyWallet{}

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
	require.Equal(t, "", devId)
}

func TestGetSignerDescriptionShouldFailOnDeviceError(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	dev.On("GetFeatures").Return(wire.Message{}, errors.New(""))
	sw := NewSkyWallet(nil, &dev)

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
	require.Equal(t, "", devId)
}

func TestGetSignerDescriptionShouldFailForFailResponse(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	f := messages.Failure{}
	fb, err := proto.Marshal(&f)
	require.Nil(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Failure),
		Data: fb,
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil, &dev)

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
	require.Equal(t, "", devId)
}

func TestGetSignerDescriptionShouldFailForInvalidMessageType(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_PinMatrixAck),
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(nil, &dev)

	// When
	devId, err := sw.GetSignerDescription()

	// Then
	require.Error(t, err)
	require.Equal(t, "", devId)
}