package walletsManager

import (
	"github.com/skycoin/skycoin/src/api"
	"github.com/therecipe/qt/core"
)


const (
	Name = int(core.Qt__UserRole) + 1
	EncryptionEnabled = int(core.Qt__UserRole) + 2
	Sky = int(core.Qt__UserRole) + 3
	CoinHours = int(core.Qt__UserRole) + 4
)

// Endpoints
const (
	GetWalletsEndpoint = "/api/v1/wallets"
)

func init() {
	WalletModel_QmlRegisterType2("WalletsManager", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("WalletsManager", 1, 0, "QWallet")
	AddressesModel_QmlRegisterType2("WalletsManager", 1, 0, "AddressModel")
	QAddress_QmlRegisterType2("WalletsManager", 1, 0, "QAddress")
	
}




//Return address of daemon node----INCOMPLETE
func nodeAddress() string {
	return "https://127.0.0.1"
}

//Return username of daemon node----INCOMPLETE
func nodeUsername() string {
	return "Kid"
}

//Return password of daemon node-----INCOMPLETE
func nodePassword() string {
	return "P@ssw0rd!"
}

func newClient() *api.Client {
	c := api.NewClient(nodeAddress())
	c.SetAuth(nodeUsername(), nodePassword())
	return c
}

type WalletModel struct {
	core.QAbstractListModel
	
	_ func() `constructor:"init"`
	
	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QWallet	`property:"wallets"`

	_ func(*QWallet) `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, sky, coinHours int)	`slot:"editWallet"`
	_ func(row int)	`slot:"removeWallet"`
	_ func()	`slot:"loadModel"`
	
}

type QWallet struct {
	core.QObject
	
	_ string `property:"name"`
	_ int 	 `property:"encryptionEnabled"`
	_ int `property:"sky"`
	_ int `property:"coinHours"`
	
	
}



func (m *WalletModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name: core.NewQByteArray2("name", -1),
		EncryptionEnabled: core.NewQByteArray2("encryptionEnabled", -1),
		Sky: core.NewQByteArray2("sky", -1),
		CoinHours: core.NewQByteArray2("coinHours", -1),
	})
	
	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddWallet(m.addWallet)
	m.ConnectEditWallet(m.editWallet)
	m.ConnectRemoveWallet(m.removeWallet)
	m.ConnectLoadModel(m.loadModel)
	
	

	m.loadModel()
	
	
}



func (m *WalletModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Wallets()){
		return core.NewQVariant()
	}

	var w = m.Wallets()[index.Row()]

	switch role{
	case Name:
		{
			return core.NewQVariant1(w.Name())
		}

	case EncryptionEnabled:
		{
			return core.NewQVariant1(w.EncryptionEnabled())
		} 

	case Sky:
		{
			return core.NewQVariant1(w.Sky())
		}
		
	case CoinHours:
		{
			return core.NewQVariant1(w.CoinHours())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *WalletModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Wallets())
}

func (m *WalletModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *WalletModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *WalletModel) addWallet(w *QWallet){
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Wallets()), len(m.Wallets()))
	m.SetWallets(append(m.Wallets(), w))
	m.EndInsertRows()
} 

func (m *WalletModel) editWallet(row int, name string, encrypted bool, sky, coinHours int) {
	w := m.Wallets()[row]
	w.SetName(name)
	w.SetEncryptionEnabled(0)
	if encrypted {
		w.SetEncryptionEnabled(1)
	}
	w.SetSky(sky)
	w.SetCoinHours(coinHours)

	pIndex := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{Name, EncryptionEnabled, Sky, CoinHours})
}


func (m *WalletModel) removeWallet(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetWallets(append(m.Wallets()[:row], m.Wallets()[row+1:]...))
	m.EndRemoveRows()
}


func getWalletsModel() ([]*QWallet, error) {
	c := newClient()
	wallets, err := c.Wallets()
	if err != nil {
		return nil, err
	}
	walletsModels := make([]*QWallet, 0)
	for _, w := range wallets {
		wlt := NewQWallet(nil)
		wlt.SetName(w.Meta.Label)
		wlt.SetEncryptionEnabled(0)
		if w.Meta.Encrypted {
			wlt.SetEncryptionEnabled(1)
		}
		bl, err := c.WalletBalance(w.Meta.Filename)
		if err != nil {
			return nil, err
		}
		wlt.SetSky(int(bl.Confirmed.Coins))
		wlt.SetCoinHours(int(bl.Confirmed.Hours))
		walletsModels = append(walletsModels, wlt)
	}
	return walletsModels, nil

	
}

func getWalletsModelTest() ([]*QWallet, error){

	walletsModel := make([]*QWallet, 0)
	wlt := NewQWallet(nil)
	wlt.SetName("Kid's Wallet")
	wlt.SetEncryptionEnabled(0)
	wlt.SetSky(200)
	wlt.SetCoinHours(1000)
	walletsModel  = append(walletsModel, wlt)

	wlt2 := NewQWallet(nil)
	wlt2.SetName("Adri's Wallet")
	wlt2.SetEncryptionEnabled(1)
	wlt2.SetSky(2000)
	wlt2.SetCoinHours(1000)
	walletsModel  = append(walletsModel, wlt2)



	return walletsModel, nil
}

func (m *WalletModel) loadModel() {
	wltsModels, err := getWalletsModelTest()
	if err != nil {
		return 
	}
	
	m.SetWallets(wltsModels)

}