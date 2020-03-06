// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import messages "github.com/fibercrypto/skywallet-protob/go"
import mock "github.com/stretchr/testify/mock"
import promise "github.com/chebyrash/promise"

// DeviceInteraction is an autogenerated mock type for the DeviceInteraction type
type DeviceInteraction struct {
	mock.Mock
}

// AddressGen provides a mock function with given fields: addressN, startIndex, confirmAddress, walletType
func (_m *DeviceInteraction) AddressGen(addressN uint32, startIndex uint32, confirmAddress bool, walletType string) *promise.Promise {
	ret := _m.Called(addressN, startIndex, confirmAddress, walletType)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(uint32, uint32, bool, string) *promise.Promise); ok {
		r0 = rf(addressN, startIndex, confirmAddress, walletType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// ApplySettings provides a mock function with given fields: usePassphrase, label, language
func (_m *DeviceInteraction) ApplySettings(usePassphrase *bool, label string, language string) *promise.Promise {
	ret := _m.Called(usePassphrase, label, language)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(*bool, string, string) *promise.Promise); ok {
		r0 = rf(usePassphrase, label, language)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// Backup provides a mock function with given fields:
func (_m *DeviceInteraction) Backup() *promise.Promise {
	ret := _m.Called()

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func() *promise.Promise); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// Cancel provides a mock function with given fields:
func (_m *DeviceInteraction) Cancel() *promise.Promise {
	ret := _m.Called()

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func() *promise.Promise); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// ChangePin provides a mock function with given fields: removePin
func (_m *DeviceInteraction) ChangePin(removePin *bool) *promise.Promise {
	ret := _m.Called(removePin)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(*bool) *promise.Promise); ok {
		r0 = rf(removePin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// CheckMessageSignature provides a mock function with given fields: message, signature, address
func (_m *DeviceInteraction) CheckMessageSignature(message string, signature string, address string) *promise.Promise {
	ret := _m.Called(message, signature, address)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(string, string, string) *promise.Promise); ok {
		r0 = rf(message, signature, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// ClearSingleTimeOperationsCache provides a mock function with given fields:
func (_m *DeviceInteraction) ClearSingleTimeOperationsCache() {
	_m.Called()
}

// Features provides a mock function with given fields:
func (_m *DeviceInteraction) Features() *promise.Promise {
	ret := _m.Called()

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func() *promise.Promise); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// GenerateMnemonic provides a mock function with given fields: wordCount, usePassphrase
func (_m *DeviceInteraction) GenerateMnemonic(wordCount uint32, usePassphrase bool) *promise.Promise {
	ret := _m.Called(wordCount, usePassphrase)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(uint32, bool) *promise.Promise); ok {
		r0 = rf(wordCount, usePassphrase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// InitializeWasWarn provides a mock function with given fields:
func (_m *DeviceInteraction) InitializeWasWarn() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsAvailable provides a mock function with given fields:
func (_m *DeviceInteraction) IsAvailable() *promise.Promise {
	ret := _m.Called()

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func() *promise.Promise); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// Recovery provides a mock function with given fields: wordCount, usePassphrase, dryRun
func (_m *DeviceInteraction) Recovery(wordCount uint32, usePassphrase *bool, dryRun bool) *promise.Promise {
	ret := _m.Called(wordCount, usePassphrase, dryRun)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(uint32, *bool, bool) *promise.Promise); ok {
		r0 = rf(wordCount, usePassphrase, dryRun)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// SecureWasWarn provides a mock function with given fields:
func (_m *DeviceInteraction) SecureWasWarn() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetInitializeWasWarn provides a mock function with given fields:
func (_m *DeviceInteraction) SetInitializeWasWarn() {
	_m.Called()
}

// SetMnemonic provides a mock function with given fields: mnemonic
func (_m *DeviceInteraction) SetMnemonic(mnemonic string) *promise.Promise {
	ret := _m.Called(mnemonic)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(string) *promise.Promise); ok {
		r0 = rf(mnemonic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// SetSecureWasWarn provides a mock function with given fields:
func (_m *DeviceInteraction) SetSecureWasWarn() {
	_m.Called()
}

// SetUploadFirmwareWasWarn provides a mock function with given fields:
func (_m *DeviceInteraction) SetUploadFirmwareWasWarn() {
	_m.Called()
}

// SignMessage provides a mock function with given fields: addressN, addressIndex, message, walletType
func (_m *DeviceInteraction) SignMessage(addressN int, addressIndex int, message string, walletType string) *promise.Promise {
	ret := _m.Called(addressN, addressIndex, message, walletType)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func(int, int, string, string) *promise.Promise); ok {
		r0 = rf(addressN, addressIndex, message, walletType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// TransactionSign provides a mock function with given fields: inputs, outputs, walletType
func (_m *DeviceInteraction) TransactionSign(inputs []*messages.SkycoinTransactionInput, outputs []*messages.SkycoinTransactionOutput, walletType string) *promise.Promise {
	ret := _m.Called(inputs, outputs, walletType)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func([]*messages.SkycoinTransactionInput, []*messages.SkycoinTransactionOutput, string) *promise.Promise); ok {
		r0 = rf(inputs, outputs, walletType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// UploadFirmware provides a mock function with given fields: payload, hash
func (_m *DeviceInteraction) UploadFirmware(payload []byte, hash [32]byte) *promise.Promise {
	ret := _m.Called(payload, hash)

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func([]byte, [32]byte) *promise.Promise); ok {
		r0 = rf(payload, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}

// UploadFirmwareWasWarn provides a mock function with given fields:
func (_m *DeviceInteraction) UploadFirmwareWasWarn() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Wipe provides a mock function with given fields:
func (_m *DeviceInteraction) Wipe() *promise.Promise {
	ret := _m.Called()

	var r0 *promise.Promise
	if rf, ok := ret.Get(0).(func() *promise.Promise); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*promise.Promise)
		}
	}

	return r0
}
