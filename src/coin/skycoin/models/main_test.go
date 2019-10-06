package skycoin

import(
	"os"
	"testing"

	"github.com/skycoin/skycoin/src/api"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
)

var global_mock *SkycoinApiMock

//Prepare the mock API for all test
func TestMain(m *testing.M){
	cf := local.GetConfigManager()
	client, _ := NewSkycoinConnectionFactory(cf.GetNode()).Create()

	global_mock = &SkycoinApiMock{Node: client.(*api.Client)}
	core.GetMultiPool().CreateSection(PoolSection, global_mock)
	util.RegisterAltcoin(NewSkyFiberPlugin(SkycoinMainNetParams))
	os.Exit(m.Run())
}