package hardware

import (
	"github.com/chebyrash/promise"
	mocks2 "github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	hardware_wallet "github.com/fibercrypto/fibercryptowallet/src/contrib/hardware-wallet"
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
	addrStr, ok := addr.(string)
	require.True(t, ok)
	require.NoError(t, err)
	require.Equal(t, orgAddrs[0], addrStr)
}