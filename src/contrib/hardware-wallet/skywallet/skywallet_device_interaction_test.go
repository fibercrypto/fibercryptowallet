package hardware

import (
	"errors"
	hardware_wallet "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
	"github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet/skywallet/mocks"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func createDeviceInteraction2(dev skywallet.Devicer) hardware_wallet.DeviceInteraction {
	return &SkyWalletInteraction{
		dev: dev,
		initializeWasWarn: false,
		uploadFirmwareWasWarn: false,
		secureWasWarn: false,
	}
}

func TestAddressGenShouldWorkOk(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgAddrs := []string{"jhfjdhfjd", "dfd787fd8"}
	addrResp := messages.ResponseSkycoinAddress{Addresses: orgAddrs}
	data, err := proto.Marshal(&addrResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_ResponseSkycoinAddress),
		Data: data,
	}
	dev.On("AddressGen", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	addrs, err := di.AddressGen(uint32(1), uint32(1), false, "deterministic").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.NoError(t, err)
	addrStrs, ok := addrs.([]string)
	require.True(t, ok)
	require.Len(t, addrStrs, 2)
	require.Equal(t, orgAddrs[0], addrStrs[0])
	require.Equal(t, orgAddrs[1], addrStrs[1])
}

func TestAddressGenShouldHandleDeviceError(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	dev.On("AddressGen", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(wire.Message{}, errors.New(""))
	di := createDeviceInteraction2(dev)

	// When
	_, err := di.AddressGen(uint32(1), uint32(1), false, "deterministic").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
}

func TestAddressGenShouldHandleErrorDecodingResponse(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgAddrs := []string{"jhfjdhfjd", "dfd787fd8"}
	addrResp := messages.ResponseSkycoinAddress{Addresses: orgAddrs}
	data, err := proto.Marshal(&addrResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Success),
		Data: data,
	}
	dev.On("AddressGen", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	_, err = di.AddressGen(uint32(1), uint32(1), false, "deterministic").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
	require.Equal(t, err, errors.New("calling DecodeResponseSkycoinAddress with wrong message type: MessageType_Success"))
}

func TestApplySettingsShouldWorkOk(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgMsg := "great"
	applySettingResp := messages.Success{Message: proto.String(orgMsg)}
	data, err := proto.Marshal(&applySettingResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Success),
		Data: data,
	}
	dev.On("ApplySettings", mock.Anything, mock.Anything, mock.Anything).Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	response, err := di.ApplySettings(nil, "l", "l").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.NoError(t, err)
	responseStr, ok := response.(string)
	require.True(t, ok)
	require.Equal(t, orgMsg, responseStr)
}

func TestApplySettingsShouldHandleDeviceError(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	dev.On("ApplySettings", mock.Anything, mock.Anything, mock.Anything).Return(wire.Message{}, errors.New(""))
	di := createDeviceInteraction2(dev)

	// When
	_, err := di.ApplySettings(nil, "l", "l").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
}

func TestApplySettingShouldHandleErrorDecodingResponse(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgResp := "jhfjdhfjd"
	applySettingsResp := messages.Failure{Message: proto.String(orgResp)}
	data, err := proto.Marshal(&applySettingsResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Failure),
		Data: data,
	}
	dev.On("ApplySettings", mock.Anything, mock.Anything, mock.Anything).Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	_, err = di.ApplySettings(nil, "l","l").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
	require.Equal(t, err, errors.New("calling DecodeSuccessMsg with wrong message type: MessageType_Failure"))
}

func TestBackupShouldWorkOk(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgMsg := "great"
	backupResp := messages.Success{Message: proto.String(orgMsg)}
	data, err := proto.Marshal(&backupResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Success),
		Data: data,
	}
	dev.On("Backup").Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	response, err := di.Backup().Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.NoError(t, err)
	responseStr, ok := response.(string)
	require.True(t, ok)
	require.Equal(t, orgMsg, responseStr)
}

func TestBackupShouldHandleDeviceError(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	dev.On("Backup").Return(wire.Message{}, errors.New(""))
	di := createDeviceInteraction2(dev)

	// When
	_, err := di.Backup().Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
}

func TestBackupShouldHandleErrorDecodingResponse(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgResp := "jhfjdhfjd"
	backupResp := messages.Failure{Message: proto.String(orgResp)}
	data, err := proto.Marshal(&backupResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Failure),
		Data: data,
	}
	dev.On("Backup").Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	_, err = di.Backup().Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
	require.Equal(t, err, errors.New("calling DecodeSuccessMsg with wrong message type: MessageType_Failure"))
}

func TestCheckMessageSignatureShouldWorkOk(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgMsg := "great"
	checkMsgSignatureRespResp := messages.Success{Message: proto.String(orgMsg)}
	data, err := proto.Marshal(&checkMsgSignatureRespResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Success),
		Data: data,
	}
	dev.On("CheckMessageSignature", mock.Anything, mock.Anything, mock.Anything).Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	response, err := di.CheckMessageSignature("m", "s", "a").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.NoError(t, err)
	responseStr, ok := response.(string)
	require.True(t, ok)
	require.Equal(t, orgMsg, responseStr)
}

func TestCheckMessageSignatureShouldHandleDeviceError(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	dev.On("CheckMessageSignature", mock.Anything, mock.Anything, mock.Anything).Return(wire.Message{}, errors.New(""))
	di := createDeviceInteraction2(dev)

	// When
	_, err := di.CheckMessageSignature("m", "s", "a").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
}

func TestCheckMessageSignatureShouldHandleErrorDecodingResponse(t *testing.T) {
	// Giving
	dev := &mocks.Devicer{}
	orgResp := "jhfjdhfjd"
	checkMsgSignResp := messages.Failure{Message: proto.String(orgResp)}
	data, err := proto.Marshal(&checkMsgSignResp)
	require.NoError(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Failure),
		Data: data,
	}
	dev.On("CheckMessageSignature", mock.Anything, mock.Anything, mock.Anything).Return(msg, nil)
	di := createDeviceInteraction2(dev)

	// When
	_, err = di.CheckMessageSignature("m", "s", "a").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.Error(t, err)
	require.Equal(t, err, errors.New("calling DecodeSuccessMsg with wrong message type: MessageType_Failure"))
}