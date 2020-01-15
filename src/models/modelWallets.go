package models

import (
	coin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	local "github.com/fibercrypto/fibercryptowallet/src/main"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtcore "github.com/therecipe/qt/core"
)

var logWalletModel = logging.MustGetLogger("Wallet Model")

const (
	QName = int(qtcore.Qt__UserRole) + iota + 1
	QAddresses
)

type ModelWallets struct {
	qtcore.QAbstractListModel
	addresses []*ModelAddresses
	WalletEnv core.WalletEnv
	_         func() `constructor:"init"`

	_ map[int]*qtcore.QByteArray `property:"roles"`
	_ bool                       `property:"loading"`

	_ func()                  `slot:"loadModel"`
	_ func()                  `slot:"cleanModel"`
	_ func([]*ModelAddresses) `slot:"addAddresses"`
}

func (m *ModelWallets) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		Name:       qtcore.NewQByteArray2("name", -1),
		QAddresses: qtcore.NewQByteArray2("qaddresses", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectCleanModel(m.cleanModel)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectAddAddresses(m.addAddresses)
	m.SetLoading(true)
	m.addresses = make([]*ModelAddresses, 0)
	altManager := local.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}

	m.WalletEnv = walletsEnvs[0]

}

func (m *ModelWallets) rowCount(*qtcore.QModelIndex) int {
	return len(m.addresses)
}

func (m *ModelWallets) roleNames() map[int]*qtcore.QByteArray {
	return m.Roles()
}

func (m *ModelWallets) data(index *qtcore.QModelIndex, role int) *qtcore.QVariant {
	if !index.IsValid() {
		return qtcore.NewQVariant()
	}

	if index.Row() >= len(m.addresses) {
		return qtcore.NewQVariant()
	}

	w := m.addresses[index.Row()]

	switch role {
	case QName:
		{
			return qtcore.NewQVariant1(w.Name())
		}
	case QAddresses:
		{
			return qtcore.NewQVariant1(w)
		}
	default:
		{
			return qtcore.NewQVariant()
		}
	}
}

func (m *ModelWallets) insertRows(row int, count int) bool {
	m.BeginInsertRows(qtcore.NewQModelIndex(), row, row+count)
	m.EndInsertRows()
	return true
}

func (m *ModelWallets) cleanModel() {
	m.SetLoading(false)
	m.addresses = make([]*ModelAddresses, 0)
}

func (m *ModelWallets) loadModel() {
	m.BeginResetModel()
	logWalletModel.Info("Loading Model")
	m.SetLoading(true)
	aModels := make([]*ModelAddresses, 0)
	wallets := m.WalletEnv.GetWalletSet().ListWallets()
	if wallets == nil {
		logWalletModel.WithError(nil).Warn("Couldn't load wallet")
		return
	}
	for wallets.Next() {

		addresses, err := wallets.Value().GetLoadedAddresses()
		if err != nil {
			logWalletModel.WithError(nil).Warn("Couldn't get loaded address")
			return
		}
		ma := NewModelAddresses(nil)
		ma.SetName(wallets.Value().GetLabel())
		ma.SetId(wallets.Value().GetId())
		oModels := make([]*ModelOutputs, 0)

		for addresses.Next() {
			a := addresses.Value()
			outputs := a.GetCryptoAccount().ScanUnspentOutputs()
			if outputs == nil {
				logWalletModel.WithField("address", a.String()).Warn("Couldn't get unspent outputs")
				continue
			}
			mo := NewModelOutputs(nil)
			mo.SetAddress(a.String())
			qOutputs := make([]*QOutput, 0)

			for outputs.Next() {
				to := outputs.Value()
				qo := NewQOutput(nil)
				qo.SetOutputID(to.GetId())
				val, err := to.GetCoins(coin.Sky)
				if err != nil {
					logWalletModel.WithError(nil).Warn("Couldn't get " + coin.Sky + " coins")
					continue
				}
				accuracy, err := util.AltcoinQuotient(coin.Sky)
				if err != nil {
					logWalletModel.WithError(err).Warn("Couldn't get " + coin.Sky + " coins quotient")
					continue
				}
				coins := util.FormatCoins(val, accuracy)
				qo.SetAddressSky(coins)
				val, err = to.GetCoins(coin.CoinHoursTicker)
				if err != nil {
					logWalletModel.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins")
					continue
				}
				accuracy, err = util.AltcoinQuotient(coin.CoinHoursTicker)
				if err != nil {
					logWalletModel.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
					continue
				}
				coinsH := util.FormatCoins(val, accuracy)
				qo.SetAddressCoinHours(coinsH)
				qOutputs = append(qOutputs, qo)
			}
			if len(qOutputs) != 0 {
				mo.addOutputs(qOutputs)
				oModels = append(oModels, mo)
			}
		}
		ma.addOutputs(oModels)
		aModels = append(aModels, ma)
	}
	logWalletModel.Info("Model loaded")
	m.addAddresses(aModels)
	m.SetLoading(false)
	m.EndResetModel()
}

func (m *ModelWallets) addAddresses(ma []*ModelAddresses) {
	for _, modelAddresses := range ma {
		find := false
		for _, modelASet := range m.addresses {
			if modelAddresses.Id() == modelASet.Id() {
				modelASet.addOutputs(modelAddresses.outputs)
				find = true
				break
			}
		}
		if !find {
			m.addresses = append(m.addresses, modelAddresses)
		}
	}
	m.insertRows(len(m.addresses), len(ma))
}
