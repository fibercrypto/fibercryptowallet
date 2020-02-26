package hardware

import (
	"github.com/chebyrash/promise"
	mocks2 "github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	hardware_wallet "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func createDeviceHelper(di hardware_wallet.DeviceInteraction) hardware_wallet.DeviceHelper {
	return &SkyWalletHelper{di: di}
}

func TestFirstAddressShouldWorkOk(t *testing.T) {
	// Giving
	di := &mocks2.DeviceInteraction{}
	orgAddrs := []string{"jhfjdhfjd", "dfd787fd8"}
	prm := promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve(orgAddrs)
	})
	di.On("AddressGen", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(prm, nil)
	dh := createDeviceHelper(di)

	// When
	addr, err := dh.FirstAddress("deterministic").Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.NoError(t, err)
	addrStr, ok := addr.(string)
	require.True(t, ok)
	require.Equal(t, orgAddrs[0], addrStr)
}

func TestMatcherShouldFindItIfFound(t *testing.T) {
	// Giving
	orgAddrs := []string{"jhfjdhfjd", "dfd787fd8"}
	addr := &mocks2.Address{}
	addr.On("String").Return(orgAddrs[0])
	addrIt := &mocks2.AddressIterator{}
	addrIt.On("Next").Return(true)
	addrIt.On("Value").Return(addr)
	wlt := &mocks2.Wallet{}
	wlt.On("GenAddresses", core.AccountAddress, uint32(0), uint32(1), mock.AnythingOfType("core.PasswordReader")).Return(addrIt)

	// When
	found := matcher(wlt, orgAddrs[0])

	// Then
	require.True(t, found)
}

func TestMatcherShouldNotFindItIfNotFound(t *testing.T) {
	// Giving
	orgAddrs := []string{"jhfjdhfjd", "dfd787fd8"}
	addr := &mocks2.Address{}
	addr.On("String").Return(orgAddrs[0])
	addrIt := &mocks2.AddressIterator{}
	addrIt.On("Next").Return(true)
	addrIt.On("Value").Return(addr)
	wlt := &mocks2.Wallet{}
	wlt.On("GenAddresses", core.AccountAddress, uint32(0), uint32(1), mock.AnythingOfType("core.PasswordReader")).Return(addrIt)

	// When
	found := matcher(wlt, "orgAddrs[0]")

	// Then
	require.False(t, found)
}

func TestMatcherShouldNotFindItIfWltCanNotGenerateAdrrs(t *testing.T) {
	// Giving
	addrIt := &mocks2.AddressIterator{}
	addrIt.On("Next").Return(false)
	wlt := &mocks2.Wallet{}
	wlt.On("GenAddresses", core.AccountAddress, uint32(0), uint32(1), mock.AnythingOfType("core.PasswordReader")).Return(addrIt)

	// When
	found := matcher(wlt, "")

	// Then
	require.False(t, found)
}

func TestDeviceMatchShouldWorkOk4DeterministicWlt(t *testing.T) {
	// Giving
	orgAddrs := []string{"jhfjdhfjd", "dfd787fd8"}
	addr := &mocks2.Address{}
	addr.On("String").Return(orgAddrs[0])
	addrIt := &mocks2.AddressIterator{}
	addrIt.On("Next").Return(true)
	addrIt.On("Value").Return(addr)
	wlt := &mocks2.Wallet{}
	wlt.On("GenAddresses", core.AccountAddress, uint32(0), uint32(1), mock.AnythingOfType("core.PasswordReader")).Return(addrIt)
	di := &mocks2.DeviceInteraction{}
	prm := promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve(orgAddrs)
	})
	di.On("AddressGen", uint32(1), uint32(0), false, skyWallet.WalletTypeDeterministic).Return(prm)
	dh := createDeviceHelper(di)

	// When
	match, err := dh.DeviceMatch(wlt).Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.NoError(t, err)
	matchBool, ok := match.(bool)
	require.True(t, ok)
	require.True(t, matchBool)
}

func TestDeviceMatchShouldWorkOk4Bip44Wlt(t *testing.T) {
	// Giving
	orgAddrs := []string{"jhfjdhfjd", "dfd787fd8"}
	addr := &mocks2.Address{}
	addr.On("String").Return(orgAddrs[0])
	addrIt := &mocks2.AddressIterator{}
	addrIt.On("Next").Return(true)
	addrIt.On("Value").Return(addr)
	wlt := &mocks2.Wallet{}
	wlt.On("GenAddresses", core.AccountAddress, uint32(0), uint32(1), mock.AnythingOfType("core.PasswordReader")).Return(addrIt)
	di := &mocks2.DeviceInteraction{}
	prm := promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve(orgAddrs)
	})
	invalidpPrm := promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve([]string{""})
	})
	di.On("AddressGen", uint32(1), uint32(0), false, skyWallet.WalletTypeDeterministic).Return(invalidpPrm)
	di.On("AddressGen", uint32(1), uint32(0), false, skyWallet.WalletTypeBip44).Return(prm)
	dh := createDeviceHelper(di)

	// When
	match, err := dh.DeviceMatch(wlt).Then(func(data interface{}) interface{} {
		return data
	}).Await()

	// Then
	require.NoError(t, err)
	matchBool, ok := match.(bool)
	require.True(t, ok)
	require.True(t, matchBool)
}

func TestShouldBeSecuredShouldReturnFalseAfterFirstTime(t *testing.T) {
	// Giving
	di := &mocks2.DeviceInteraction{}
	di.On("SecureWasWarn").Return(true)
	dh := createDeviceHelper(di)

	// When
	val, err := dh.ShouldBeSecured().Await()

	// Then
	require.NoError(t, err)
	matchBool, ok := val.(bool)
	require.True(t, ok)
	require.False(t, matchBool)
}

func TestShouldBeSecuredShouldReturnTrueForPinProtection(t *testing.T) {
	// Giving
	di := &mocks2.DeviceInteraction{}
	di.On("SecureWasWarn").Return(false)
	prm := promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve(messages.Features{PinProtection: proto.Bool(false)})
	})
	di.On("Features").Return(prm)
	dh := createDeviceHelper(di)

	// When
	val, err := dh.ShouldBeSecured().Await()

	// Then
	require.NoError(t, err)
	matchBool, ok := val.(bool)
	require.True(t, ok)
	require.True(t, matchBool)
}

func TestShouldBeSecuredShouldReturnTrueForNeedBackup(t *testing.T) {
	// Giving
	di := &mocks2.DeviceInteraction{}
	di.On("SecureWasWarn").Return(false)
	prm := promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve(
			messages.Features{
				PinProtection: proto.Bool(true),
				NeedsBackup: proto.Bool(true)})
	})
	di.On("Features").Return(prm)
	dh := createDeviceHelper(di)

	// When
	val, err := dh.ShouldBeSecured().Await()

	// Then
	require.NoError(t, err)
	matchBool, ok := val.(bool)
	require.True(t, ok)
	require.True(t, matchBool)
}

func TestShouldBeSecuredShouldReturnFalseForNoNeedBackupAndHavePin(t *testing.T) {
	// Giving
	di := &mocks2.DeviceInteraction{}
	di.On("SecureWasWarn").Return(false)
	prm := promise.New(func(resolve func(interface{}), reject func(error)) {
		resolve(
			messages.Features{
				PinProtection: proto.Bool(true),
				NeedsBackup: proto.Bool(false)})
	})
	di.On("Features").Return(prm)
	dh := createDeviceHelper(di)

	// When
	val, err := dh.ShouldBeSecured().Await()

	// Then
	require.NoError(t, err)
	matchBool, ok := val.(bool)
	require.True(t, ok)
	require.False(t, matchBool)
}