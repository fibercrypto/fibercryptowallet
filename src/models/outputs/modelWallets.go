package outputs

import (
	qtcore "github.com/therecipe/qt/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models" //callable as skycoin
)

const (
	Name = int(qtcore.Qt__UserRole) + 1
)

type ModelWallets struct {
	qtcore.QAbstractListModel

	WalletEnv core.WalletEnv
	_ func()                      `constructor:"init"`

	_ map[int]*qtcore.QByteArray  `property:"roles"`
	_ []*ModelAddresses        	  `property:"addresses"`
	_ func()				      `slot:"loadModel"`
}

func (m *ModelWallets) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Name: 			 qtcore.NewQByteArray2("name", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectLoadModel(m.loadModel)

	//Set the correct NodeAddress
	addr := "http://127.0.0.1:37039"
	m.WalletEnv = &skycoin.WalletNode{NodeAddress: addr}

	m.loadModel()
}

func (m *ModelWallets) rowCount(*qtcore.QModelIndex) int {
	return len(m.Transactions())
}

func (m *ModelWallets) roleNames() map[int]*qtcore.QByteArray {
	return m.Roles()
}

func (m *ModelWallets) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}

	if index.Row() >= len(m.Addresses()){
		return qtcore.NewQVariant()
	}

	w := m.Addresses()[index.Row()]

	switch role{
	case Name:
		{
			return qtcore.NewQVariant1(w.Name())
		}
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *ModelWallets) loadModel() {
	aModels = make([]*ModelAddresses, 0)

	wallets := m.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		return
	}
	for wallets.Next() {
		addresses, err = wallets.Value().GetLoadedAddresses()
		if err != nil {
			println(err)
			return
		}
		ma = NewModelAddresses(nil)
		ma.SetName(wallets.Value().GetLabel())
		oModels = make([]*ModelOutputs, 0)

		for addresses.Next() {
			outputs = addresses.Value().GetCryptoAccount().ScanUnspentOutputs()
			mo = NewModelOutputs(nil)
			mo.SetAddress(addresses.Value().String())
			qOutputs = make([]*QOutput)

			for outputs.Next() {
				to = outputs.Value()
				qo = NewQOutput(nil)
				qo.SetOutputID(to.GetId())
				qo.SetAddressSky(to.GetCoins("Sky"))
				qo.SetAddressCoinHours(to.GetCoins(""))
				qOutputs = append(qOutputs, qo)
			}
			mo.SetOutputs(qOutputs)
			oModels = append(oModels, mo)
		}
		ma.SetOutputs(oModels)
		aModels = append(aModels, ma)
	}
	m.SetAddresses(aModels)
}
