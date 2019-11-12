package skywallet

import (
	"bytes"
	"sync"
	"testing"

	messages "github.com/fibercrypto/skywallet-protob/go"

	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type devicerSuit struct {
	suite.Suite
}

func (suite *devicerSuit) SetupTest() {
}

func TestDevicerSuitSuit(t *testing.T) {
	suite.Run(t, new(devicerSuit))
}

type testHelperCloseableBuffer struct {
	bytes.Buffer
}

func (cwr testHelperCloseableBuffer) Read(p []byte) (n int, err error) {
	return 0, nil
}
func (cwr testHelperCloseableBuffer) Write(p []byte) (n int, err error) {
	return 0, nil
}
func (cwr testHelperCloseableBuffer) Close(disconnect bool) error {
	return nil
}

func (suite *devicerSuit) TestAddressGen() {
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	tt := []struct {
		name       string
		addressN   uint32
		startIndex uint32
		err        error
		msgKind    uint16
	}{
		{
			name:       "addressN zero",
			addressN:   0,
			startIndex: 0,
			err:        ErrAddressNZero,
			msgKind:    0,
		},

		{
			name:       "no error",
			addressN:   1,
			startIndex: 0,
			msgKind:    2,
		},
	}

	for _, tc := range tt {
		msg, err := device.AddressGen(tc.addressN, tc.startIndex, false)
		suite.Equal(err, tc.err)
		suite.Equal(msg.Kind, tc.msgKind)
	}

	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
}

func (suite *devicerSuit) TestApplySettings() {
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	tt := []struct {
		name          string
		usePassphrase *bool
		label         string
		language      string
		err           error
		msgKind       uint16
	}{
		{
			name:          "no error",
			usePassphrase: new(bool),
			msgKind:       2,
		},
	}

	for _, tc := range tt {
		msg, err := device.ApplySettings(tc.usePassphrase, tc.label, tc.language)
		suite.Equal(err, tc.err)
		suite.Equal(msg.Kind, tc.msgKind)
	}

	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
}

func (suite *devicerSuit) TestBackup() {
	// NOTE(denisacostaq@gmail.com): Giving
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	// NOTE(denisacostaq@gmail.com): When
	msg, err := device.Backup()

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Nil(err)
	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
	require.Equal(suite.T(), msg.Kind, uint16(messages.MessageType_MessageType_Success))
}

func (suite *devicerSuit) TestCancel() {
	// NOTE(denisacostaq@gmail.com): Giving
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	// NOTE(denisacostaq@gmail.com): When
	msg, err := device.Cancel()

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Nil(err)
	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
	require.Equal(suite.T(), msg.Kind, uint16(messages.MessageType_MessageType_Success))
}

func (suite *devicerSuit) TestCheckMessageSignature() {
	// NOTE(denisacostaq@gmail.com): Giving
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	// NOTE(denisacostaq@gmail.com): When
	msg, err := device.CheckMessageSignature("", "", "")

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Nil(err)
	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
	require.Equal(suite.T(), msg.Kind, uint16(messages.MessageType_MessageType_Success))
}

func (suite *devicerSuit) TestFirmwareUpload() {
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	device := getMockDevice(driverMock)

	tt := []struct {
		name       string
		deviceType DeviceType
		err        error
	}{
		{
			name:       "emulator",
			deviceType: DeviceTypeEmulator,
			err:        ErrDeviceTypeEmulator,
		},
	}

	for _, tc := range tt {
		driverMock.On("DeviceType").Return(tc.deviceType)
		err := device.FirmwareUpload([]byte{}, [32]byte{})
		suite.Equal(err, tc.err)
	}

	driverMock.AssertCalled(suite.T(), "DeviceType")
}

func (suite *devicerSuit) TestGenerateMnemonic() {
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	tt := []struct {
		name      string
		wordCount uint32
		err       error
		msgKind   uint16
	}{
		{
			name:      "invalid word count",
			wordCount: 36,
			err:       ErrInvalidWordCount,
			msgKind:   0,
		},

		{
			name:      "no error",
			wordCount: 12,
			msgKind:   2,
		},
	}

	for _, tc := range tt {
		msg, err := device.GenerateMnemonic(tc.wordCount, false)
		suite.Equal(err, tc.err)
		suite.Equal(msg.Kind, tc.msgKind)
	}

	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
}

func (suite *devicerSuit) TestRecovery() {
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	tt := []struct {
		name          string
		wordCount     uint32
		usePassphrase bool
		dryRun        bool
		err           error
		msgKind       uint16
	}{
		{
			name:      "invalid word count",
			wordCount: 36,
			err:       ErrInvalidWordCount,
			msgKind:   0,
		},

		{
			name:      "no error",
			wordCount: 12,
			msgKind:   2,
		},
	}

	for _, tc := range tt {
		msg, err := device.Recovery(tc.wordCount, &tc.usePassphrase, tc.dryRun)
		suite.Equal(err, tc.err)
		suite.Equal(msg.Kind, tc.msgKind)
	}

	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
}

func (suite *devicerSuit) TestSetMnemonic() {
	// NOTE(denisacostaq@gmail.com): Giving
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	// NOTE(denisacostaq@gmail.com): When
	msg, err := device.SetMnemonic("cloud flower upset remain green metal below cup stem infant art thank")

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Nil(err)
	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
	require.Equal(suite.T(), msg.Kind, uint16(messages.MessageType_MessageType_Success))
}

func (suite *devicerSuit) TestRemovePinCode() {
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	tt := []struct {
		name      string
		removePin *bool
		err       error
		msgKind   uint16
	}{
		{
			name:    "removePin nil",
			err:     ErrRemovePinNil,
			msgKind: 0,
		},

		{
			name:      "no error",
			removePin: new(bool),
			msgKind:   2,
		},
	}

	for _, tc := range tt {
		msg, err := device.ChangePin(tc.removePin)
		suite.Equal(err, tc.err)
		suite.Equal(msg.Kind, tc.msgKind)
	}

	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
}

func (suite *devicerSuit) TestTransactionSign() {
	// NOTE(denisacostaq@gmail.com): Giving
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	// NOTE(denisacostaq@gmail.com): When
	msg, err := device.TransactionSign(nil, nil)

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Nil(err)
	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
	require.Equal(suite.T(), msg.Kind, uint16(messages.MessageType_MessageType_Success))
}

func (suite *devicerSuit) TestWipe() {
	// NOTE(denisacostaq@gmail.com): Giving
	driverMock := &MockDeviceDriver{}
	driverMock.On("GetDevice").Return(&testHelperCloseableBuffer{}, nil)
	driverMock.On("SendToDevice", mock.Anything, mock.Anything).Return(
		wire.Message{Kind: uint16(messages.MessageType_MessageType_Success), Data: nil}, nil)
	device := getMockDevice(driverMock)

	// NOTE(denisacostaq@gmail.com): When
	msg, err := device.Wipe()
	// NOTE: Assert
	suite.Nil(err)
	driverMock.AssertCalled(suite.T(), "GetDevice")
	driverMock.AssertNumberOfCalls(suite.T(), "SendToDevice", 1)
	mock.AssertExpectationsForObjects(suite.T(), driverMock)
	require.Equal(suite.T(), msg.Kind, uint16(messages.MessageType_MessageType_Success))
}

func getMockDevice(mock *MockDeviceDriver) Device {
	return Device{mock, sync.Mutex{}, nil, false, false, ButtonType(-1)}
}
