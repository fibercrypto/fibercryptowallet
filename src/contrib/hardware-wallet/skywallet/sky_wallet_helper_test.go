package hardware

import (
	"github.com/chebyrash/promise"
	mocks2 "github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	hardware_wallet "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	skyWallet "github.com/fibercrypto/skywallet-go/src/skywallet"
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