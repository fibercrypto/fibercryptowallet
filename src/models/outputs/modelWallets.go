package outputs

import (
	qtcore "github.com/therecipe/qt/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models" //callable as skycoin
)

const (
	Name = int(qtcore.Qt__UserRole) + 1
	QAddresses = int(qtcore.Qt__UserRole) + 2
	//Set the correct NodeAddress
	ADDR = "http://127.0.0.1:46445"
)

type ModelWallets struct {
	qtcore.QAbstractListModel

	WalletEnv core.WalletEnv
	_ func()                      `constructor:"init"`

	_ map[int]*qtcore.QByteArray  `property:"roles"`
	_ []*ModelAddresses        	  `property:"addresses"`

	_ func()				      `slot:"loadModel"`
	_ func([]*ModelAddresses) 	  `slot:"addAddresses"`
}

func (m *ModelWallets) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Name: 			 qtcore.NewQByteArray2("name", -1),
		QAddresses:		 qtcore.NewQByteArray2("addresses", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectAddAddresses(m.addAddresses)

	m.WalletEnv = &skycoin.WalletNode{NodeAddress: ADDR}

	m.loadModel()
}

func (m *ModelWallets) rowCount(*qtcore.QModelIndex) int {
	println("Row count:", len(m.Addresses()))
	return len(m.Addresses())
}

func (m *ModelWallets) roleNames() map[int]*qtcore.QByteArray {
	return m.Roles()
}

func (m *ModelWallets) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	println("data requested")
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
	case QAddresses:
		{
			return qtcore.NewQVariant1(m.Addresses())
		}
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *ModelWallets) insertRows(row int, count int) bool {
	m.BeginInsertRows(qtcore.NewQModelIndex(), row, row + count)
	m.EndInsertRows()
	return true
}

func (m *ModelWallets) loadModel() {
	aModels := make([]*ModelAddresses, 0)
	println("loadModel")

	wallets := m.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		return
	}
	for wallets.Next() {
		addresses, err := wallets.Value().GetLoadedAddresses()
		if err != nil {
			println(err)
			return
		}
		ma := NewModelAddresses(nil)
		ma.SetName(wallets.Value().GetLabel())
		oModels := make([]*ModelOutputs, 0)

		for addresses.Next() {
			a := addresses.Value()
			outputs := a.GetCryptoAccount().ScanUnspentOutputs()
			mo := NewModelOutputs(nil)
			mo.SetAddress(a.String())
			qOutputs := make([]*QOutput, 0)

			for outputs.Next() {
				to := outputs.Value()
				qo := NewQOutput(nil)
				qo.SetOutputID(to.GetId())
				qo.SetAddressSky(to.GetCoins("Sky"))
				qo.SetAddressCoinHours(to.GetCoins(""))
				qOutputs = append(qOutputs, qo)
			}
			mo.addOutputs(qOutputs)
			//println("Address of ModelOutputs ", a.String(), " --> ", &mo)
			oModels = append(oModels, mo)
		}
		ma.addOutputs(oModels)
		aModels = append(aModels, ma)
		//println("Address of ModelAddresses on append --> ", &ma)
	}
	m.addAddresses(aModels)
}

func (m *ModelWallets) addAddresses(ma []*ModelAddresses) {
	m.SetAddresses(ma)
	m.insertRows(len(m.Addresses()), len(ma))
}
