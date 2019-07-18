package WalletManager

import (
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/wallet"
	"github.com/therecipe/qt/core"
)


const (
	Name = int(core.Qt__UserRole) + 1<<iota
	EncryptionEnabled
	Sky
	CoinHours
)

// Endpoints
const (
	GetWalletsEndpoint = "/api/v1/wallets"
)

func init() {
	WalletModel_QmlRegisterType2("WalletsManager", 1, 0, "WalletModel")
	QWallet_QmlRegisterType2("WalletsManager", 1, 0, "QWallet")
	
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
	core.QAbstracListModel
	_ func() `constructor:"init"`
	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QWallet	`property:"wallets"`

}

type QWallet struct {
	core.QObject
	_ string `property:"name"`
	_ bool `property:"encryptionEnabled"`
	_ int `property:"sky"`
	_ int `property:"coinHours"`
	
	_ func() `constructor:"init"`

	_ func(*QWallet) `slot:"addWallet"`
	_ func(row int, name string, encryptionEnabled bool, sky, coinHours int)	`slot:"editWallet"`
	_ func(row int)	`slot:"removeWallet"`
	_ func()	`slot:"loadModel"`

}



func (m *WalletModel) init() {
	m.setRoles(map[int]*core.QByteArray{
		Name: core.NewQByteArray2("name", -1),
		EncryptionEnabled : core.NewQByteArray2("encryptionEnabled", -1),
		Sky: core.NewQByteArray2("sky", -1),
		CoinHours : core.NewQByteArray2("coinHours", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRolesNames(m.roleNames)

	m.ConnectAddWallet(m.addWallet)
	m.ConnectEditWallet(m.editWallet)
	m.ConnectRemoveWallet(m.removeWallet)


	
}

func (m *WalletModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Wallets()){
		return core.NewQVariant
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
	m.SetWallet(append(m.Wallets(), w))
	m.EndInsertRows()
} 

func (m *WalletModel) editWallet(row int, name string, encryptionEnabled bool, sky, coinHours int) {
	w := m.Wallets()[row]
	w.SetName(name)
	w.SetEncryptionEnabled(encryptionEnabled)
	w.SetSky(sky)
	w.SetCoinHours(coinHours)

	pIndex := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int{Name, EncryptionEnabled, Sky, CoinHours})
}


func (m *WalletModel) removeWallet(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetWallets(append(m.Wallets()[:row], m.Wallest()[row+1:]...))
	m.EndRemoveRows()
}

func (m *WalletModel) loadModel() {
	wltsModels, err := getWalletsModel()
	if err != nil {
		return err
	}
	m.SetWallets(wltsModels)

}



func (walletM *WalletManager) getWalletsModel() []*QWallet {
	c := newClient()
	wallets, err := c.Wallets()
	if err != nil {
		return nil, err
	}
	walletsModels := make([]*QWallet, len(wallets))
	for _, w in range wallets {
		wlt := NewQWallet(nil)
		wlt.SetName(w.Meta["label"])
		wlt.SetEncryptionEnabled(w.Meta["encrypted"])
		wlt.SetSky(2)
		wlt.SetCoinHours(10)
		append(walletsModels, &wlt)
	}
	return walletsModels, nil

	
}
