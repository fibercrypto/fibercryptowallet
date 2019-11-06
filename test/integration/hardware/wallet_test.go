package hardware

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/hardware"
	"github.com/skycoin/hardware-wallet-go/src/skywallet"
	"testing"

	//MM "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	_ "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin"
	"github.com/stretchr/testify/require"
	//integrationtestutil "github.com/fibercrypto/FiberCryptoWallet/test/integration/util"
)



func TestTransactionSignInputFromDevice(t *testing.T) {
	dev := skywallet.NewDevice(skywallet.DeviceTypeEmulator)
	hs := hardware.NewSkyWallet(dev)
	require.NotNil(t, hs)

	//keyTestData, err := integrationtestutil.GenerateTestKeyPair(t)
	//require.NoError(t, err)
	//integrationtestutil.ForceSetMnemonic(t, dev, keyTestData.Mnemonic)
	//MM.TransactionSignInputTest(t, hs)
}