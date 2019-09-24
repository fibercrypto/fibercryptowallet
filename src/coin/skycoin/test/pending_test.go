package test

import(
	"testing"
	//_ "github.com/stretchr/testify/mock"

	mocks "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/test"
	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin"
	models "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	core "github.com/fibercrypto/FiberCryptoWallet/src/core"
	local "github.com/fibercrypto/FiberCryptoWallet/src/main"
	util "github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"

)

func TestInit(t *testing.T){
	println("testing if init run")
	cf := local.GetConfigManager()
	client, _ := models.NewSkycoinConnectionFactory(cf.GetNode()).Create()
	fac := &mocks.SkycoinConnectionFactoryMock{}
	fac.SetNode(client.(*api.Client))
	core.GetMultiPool().CreateSection(models.PoolSection, fac)
	util.RegisterAltcoin(sky.NewSkyFiberPlugin(sky.SkycoinMainNetParams))


}
