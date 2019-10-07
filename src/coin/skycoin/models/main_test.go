package skycoin

import(
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
)

var global_mock *SkycoinApiMock

//util when is needed to change the values of an 
//API method used in other test with diferent values.
func CleanGlobalMock(){
	global_mock.ExpectedCalls = []*mock.Call{}
}

//Prepare the mock API for all test
func TestMain(m *testing.M){
	global_mock = new(SkycoinApiMock)
	core.GetMultiPool().CreateSection(PoolSection, global_mock)
	util.RegisterAltcoin(NewSkyFiberPlugin(SkycoinMainNetParams))
	os.Exit(m.Run())
}