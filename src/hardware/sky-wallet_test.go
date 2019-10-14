package hardware

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/hardware/mocks"
	"github.com/micro/protobuf/proto"
	"github.com/skycoin/hardware-wallet-go/src/skywallet/wire"
	"github.com/skycoin/hardware-wallet-protob/go"
	"github.com/stretchrcom/testify/require"
	"testing"
	"errors"
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
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerUID()

	// Then
	require.Equal(t, core.UID(expectedDevId), devId)
}

func TestGetSignerUIDShouldFailForUninitializedDevice(t *testing.T) {
	// Giving
	sw := SkyWallet{}

	// When
	devId := sw.GetSignerUID()

	// Then
	require.Equal(t, core.UID("undefined"), devId)
}

func TestGetSignerUIDShouldFailOnDeviceError(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	dev.On("GetFeatures").Return(wire.Message{}, errors.New(""))
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerUID()

	// Then
	require.Equal(t, core.UID("undefined"), devId)
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
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerUID()

	// Then
	require.Equal(t, core.UID("undefined"), devId)
}

func TestGetSignerUIDShouldFailForInvalidMessageType(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_PinMatrixAck),
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerUID()

	// Then
	require.Equal(t, core.UID("undefined"), devId)
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
	sw := NewSkyWallet(&dev)

	// When
	devDescription := sw.GetSignerDescription()

	// Then
	require.Equal(t, expectedDevDescription, devDescription)
}

func TestGetSignerDescriptionShouldFailForUninitializedDevice(t *testing.T) {
	// Giving
	sw := SkyWallet{}

	// When
	devId := sw.GetSignerDescription()

	// Then
	require.Equal(t, "undefined", devId)
}

func TestGetSignerDescriptionShouldFailOnDeviceError(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	dev.On("GetFeatures").Return(wire.Message{}, errors.New(""))
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerDescription()

	// Then
	require.Equal(t, "undefined", devId)
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
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerDescription()

	// Then
	require.Equal(t, "undefined", devId)
}

func TestGetSignerDescriptionShouldFailForInvalidMessageType(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_PinMatrixAck),
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerDescription()

	// Then
	require.Equal(t, "undefined", devId)
}