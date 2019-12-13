package hardware

import (
	skycoin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	integrationtestutil "github.com/fibercrypto/fibercryptowallet/test/integration/util"
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func setUpHardwareWallet(t *testing.T) {
	util.RegisterAltcoin(skycoin.NewSkyFiberPlugin(skycoin.SkycoinMainNetParams))
	dev := skywallet.NewDevice(skywallet.DeviceTypeEmulator)
	keyTestData, err := skycoin.GenerateTestKeyPair(t)
	require.NoError(t, err)
	integrationtestutil.ForceSetMnemonic(t, dev, keyTestData.Mnemonic)
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
	setUpHardwareWallet(t)
	callback := func(dev skywallet.Devicer, prvMsg wire.Message, outsLen int) (wire.Message, error) {
		var msg wire.Message
		for outsLen > 0 {
			integrationtestutil.PressAcceptButton(t, dev, prvMsg, messages.MessageType_MessageType_ButtonRequest)
			if outsLen == 1 {
				msg = integrationtestutil.PressAcceptButton(t, dev, prvMsg, messages.MessageType_MessageType_ResponseTransactionSign)
			} else {
				msg = integrationtestutil.PressAcceptButton(t, dev, prvMsg, messages.MessageType_MessageType_ButtonRequest)
			}
			outsLen--
		}
		return msg, nil
	}
	dev := skywallet.NewDevice(skywallet.DeviceTypeEmulator)
	hs1 := NewSkyWallet(nil, dev, callback)
	hs2 := NewSkyWallet(nil, dev, callback)
	hs3 := NewSkyWallet(nil, dev, callback)
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
