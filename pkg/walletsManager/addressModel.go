package walletsManager

import (
		"github.com/therecipe/qt/core"
		"fmt"
)

const (
	Address = int(core.Qt__UserRole) + 1
	ASky = int(core.Qt__UserRole) + 2
	ACoinHours = int(core.Qt__UserRole) + 3
)



type AddressesModel struct{
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ [] *QAddress	`property:"addresses"`

	_ int	`property:"loaded"`

	_ func(*QAddress)	`slot:"addAddress"`
	_ func(int)	`slot:"removeAddress"`
	_ func(int, string, int, int)	`slot:"editAddress"`
	_ func(string)	`slot:"loadModel"`


}


type QAddress struct{
	core.QObject

	_ string `property:"address"`
	_ int	`property:"addressSky"`
	_ int	`property:"addressCoinHours"`
}

func (m *AddressesModel) init(){
	m.SetRoles(map[int]*core.QByteArray{
		Address: core.NewQByteArray2("address", -1),
		ASky: core.NewQByteArray2("addressSky", -1),
		ACoinHours: core.NewQByteArray2("addressCoinHours", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddAddress(m.addAddress)
	m.ConnectEditAddress(m.editAddress)
	m.ConnectRemoveAddress(m.removeAddress)
	m.ConnectLoadModel(m.loadModel)

	m.SetLoaded(0)


}


func (m *AddressesModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Addresses()){
		return core.NewQVariant()
	}

	var a = m.Addresses()[index.Row()]

	switch role{
	
	case Address:
		{
			return core.NewQVariant1(a.Address())
		}
	case ASky:
		{
			return core.NewQVariant1(a.AddressSky())
		}
	case ACoinHours:
		{
			return core.NewQVariant1(a.AddressCoinHours())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *AddressesModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Addresses())
}

func (m *AddressesModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *AddressesModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *AddressesModel) addAddress(address *QAddress) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Addresses()), len(m.Addresses()))
	m.SetAddresses(append(m.Addresses(), address))
	m.EndInsertRows()
}

func (m *AddressesModel) removeAddress(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetAddresses(append(m.Addresses()[:row], m.Addresses()[row+1:]...))
	m.EndRemoveRows()
}

func (m *AddressesModel) editAddress(row int, address string, sky, coinHours int){
	a := m.Addresses()[row]
	a.SetAddress(address)
	a.SetAddressSky(sky)
	a.SetAddressCoinHours(coinHours)
	
	pIndex := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(pIndex, pIndex, []int {Address, ASky, ACoinHours})
}

func (m *AddressesModel) loadModel(wallet string){
	fmt.Println(wallet)
	Qaddresses, err := getQAddresses(wallet)
	fmt.Println("Cargando direcciones")
	if err != nil {
		return
	}
	
	addresses := make([]*QAddress, 0)
	address := NewQAddress(nil)
	address.SetAddress("--------------------------")
	address.SetAddressSky(0)
	address.SetAddressCoinHours(0)
	addresses = append(addresses, address)
	addresses = append(addresses, Qaddresses...)
	fmt.Println(len(addresses))
	m.SetAddresses(addresses)
	fmt.Println("Direcciones cargadas")
	m.SetLoaded(1)
}

func getQAddresses(wallet string) ([]*QAddress, error) {
	c := newClient()
	wlt, err := c.Wallet(wallet)
	if err != nil {
		return nil, err
	}
	Qaddresses := make([]*QAddress, 0)
	entries := wlt.Entries
	for _, entry := range entries {
		address := NewQAddress(nil)
		address.SetAddress(entry.Address)
		addresses := make([]string, 0)
		addresses = append(addresses, entry.Address)
		bl, err := c.Balance(addresses)
		if err != nil {
			return nil, err
		}
		address.SetAddressSky(int(bl.Confirmed.Coins))
		address.SetAddressCoinHours(int(bl.Confirmed.Hours))
		Qaddresses = append(Qaddresses, address)
	}

	return Qaddresses, nil
}

