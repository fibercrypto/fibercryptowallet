package hardware

import (
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func assertText(t *testing.T, msg interface{}, text string) {
	msgStr, ok := msg.(string)
	require.True(t, ok)
	require.Equal(t, msgStr, text)
}

func forceWipe(t *testing.T) {
	dev := SkyWltInteractionInstance()
	msg, err := dev.Wipe().Then(func(data interface{}) interface{} {
		return data
	}).Await()
	require.NoError(t, err)
	assertText(t, msg, "Device wiped")
}

func forceSetMnemonic(t *testing.T, mnemonic string) {
	forceWipe(t)
	dev := SkyWltInteractionInstance()
	msg, err := dev.SetMnemonic(mnemonic).Then(func(data interface{}) interface{} {
		return data
	}).Await()
	require.NoError(t, err)
	assertText(t, msg, "wire junk original sword bread bottom armor dog snow accident inform rigid")
}

func createEmulatorDevice() {
	CreateSkyWltInteractionInstanceOnce(
		true,
		skywallet.DeviceTypeEmulator,
		func(requestKind skywallet.InputRequestKind, title, message string)(string, error){return "", nil})
}

func setUpHardwareWallet(t *testing.T) {
	util.RegisterAltcoin(skycoin.NewSkyFiberPlugin(skycoin.SkycoinMainNetParams))
	createEmulatorDevice()
	keyTestData, err := skycoin.GenerateTestKeyPair(t)
	require.NoError(t, err)
	forceSetMnemonic(t, keyTestData.Mnemonic)
}

var logModelTest = logging.MustGetLogger("Skycoin Hardware Wallet Test")

//Prepare the mock API for all test
func TestMain(m *testing.M) {
	mock := new(skycoin.SkycoinApiMock)
	skycoin.SetGlobalMock(mock)
	err := core.GetMultiPool().CreateSection(skycoin.PoolSection, mock)
	if err != nil {
		logModelTest.WithError(err).Warn("Error creating pool section")
		return
	}
	util.RegisterAltcoin(skycoin.NewSkyFiberPlugin(skycoin.SkycoinMainNetParams))
	os.Exit(m.Run())
}

func TestTransactionSignInputFromDevice(t *testing.T) {
	skycoin.CleanGlobalMock()
	skyApiMock := skycoin.GetGlobalMock()
	err := core.GetMultiPool().CreateSection(skycoin.PoolSection, skyApiMock)
	if err != nil {
		logModelTest.WithError(err).Warn("Error creating pool section")
		return
	}
	setUpHardwareWallet(t)
	hs1 := NewSkyWallet(nil)
	hs2 := NewSkyWallet(nil)
	hs3 := NewSkyWallet(nil)
	setWallet := func(t *testing.T, signer core.TxnSigner, wlt core.Wallet) {
		hs, ok := signer.(*SkyWallet)
		require.True(t, ok)
		hs.wlt = wlt
	}
	signers := make([]core.TxnSigner, 3)
	signers[0] = hs1
	signers[1] = hs2
	signers[2] = hs3
	skycoin.TransactionSignInputTestSkyHwImpl(t, signers, setWallet)
}
