package outputs

import (
	"github.com/therecipe/qt/core"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/simelo/FiberCryptoWallet/src/util"
)

const (
	Name = int(core.Qt__UserRole) + 1
	Entries = int(core.Qt__UserRole) + 2
)

type OutputsList struct {
	core.QAbstractListModel

	_ func()                    `constructor:"init"`

	_ map[int]*core.QByteArray  `property:"roles"`
	_ []*WalletsOutputs         `property:"wallets"`
}

type WalletsOutputs struct {
	core.QObject
	
	_ string              `property:"name"`
	_ []*QEntries         `property:"entries"`
}

type QEntries struct {
	core.QObject
	
	_ string              `property:"address"`
	_ []*QOutput          `property:"outputs"`
}

type QOutput struct {
	core.QObject
	
	_ string       `property:"hash"`
	_ int          `property:"sky"`
	_ int          `property:"coinHours"`
}

func (m *OutputsList) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Name: 			 core.NewQByteArray2("name", -1),
		Entries:             core.NewQByteArray2("entries", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)

}

func (m *OutputsList) rowCount(*core.QModelIndex) int {
	return len(m.Transactions())
}

func (m *OutputsList) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *OutputsList) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Wallets()){
		return core.NewQVariant()
	}

	wo := m.Wallets()[index.Row()]

	switch role{
	case Name:
		{
			return core.NewQVariant1(wo.Name())
		}
	case Entries:
		{
			return core.NewQVariant1(wo.Entries())
		}
	default:
		{
			return core.NewQVariant()
		}
	}
}

func getData() ([]*WalletsOutputs, error) {
	c := util.newClient()
	wallets, err := c.Wallets()
	if err != nil {
		return nil, err
	}
	woModels := make([]*WalletsOutputs, 0)
	for _, w := range wallets {
		wo := NewWalletsOutputs(nil)
		wo.SetName(w.Meta.Label)
		entries := make([]*QEntries, 0)
		for _, e := range w.Entries {
			entry := NewQEntries(nil)
			entry.SetAddress(e.Address.Key)
			outputs, err := getOutputs(e.Address.Key)
			if err != nil {
				return nil, err
			}
			entries = append(entries, entry)
		}
		woModels = append(woModels, wo)
	}
	return woModels, nil
}

func getOutputs(string address) ([]*QOutput, error) {
	c := util.newClient()
	outputs, err := c.Outputs([address])
	if err != nil {
		return nil, err
	}
	qOutputs := make([]*QOutput, 0)
	for _, o := range outputs {
		/*o --> UnspentOutputsSummary
				* HeadOutputs are unspent outputs confirmed in the blockchain
  					HeadOutputs UnspentOutputs json:"head_outputs"
  				* OutgoingOutputs are unspent outputs being spent in unconfirmed transactions
  					OutgoingOutputs UnspentOutputs json:"outgoing_outputs"
  				* IncomingOutputs are unspent outputs being created by unconfirmed transactions
  					IncomingOutputs UnspentOutputs json:"incoming_outputs"
		*/
		for _, unspentOutput := range o.HeadOutputs {
			qo := NewQOutput(nil)
			qo.SetHash(unspentOutput.Hash)
			qo.SetSky(unspentOutput.Coins)
			qo.SetCoinHours(unspentOutput.Hours)
			qOutputs = append(qOutputs, qo)
		}
	}
	return qOutputs, nil
} 