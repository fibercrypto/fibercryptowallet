package models

import (
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	"sync"
	"time"

	//"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/config"
	coin "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/models"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	local "github.com/fibercrypto/fibercryptowallet/src/main"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	qtcore "github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	//"time"
)

var logWalletModel = logging.MustGetLogger("Wallet Model")

var Helper = NewHelper(nil)

type helper struct {
	qtcore.QObject

	_ func(f func()) `signal:"runInMain,auto"`
}

func (h *helper) runInMain(f func()) { f() }

const (
	QName = int(qtcore.Qt__UserRole) + iota + 1
	QAddresses
)

type ModelWallets struct {
	qtcore.QAbstractListModel
	addresses      []*ModelAddresses `property:"addresses"`
	addressesMutex sync.Mutex
	WalletEnv      core.WalletEnv
	_              func() `constructor:"init"`

	_ map[int]*qtcore.QByteArray `property:"roles"`
	_ bool                       `property:"loading"`

	_ func()                  `slot:"loadModel"`
	_ func(int)               `slot:"removeWallet"`
	_ func()                  `slot:"cleanModel"`
	_ func([]*ModelAddresses) `slot:"addAddresses"`
}

func (m *ModelWallets) init() {
	m.SetRoles(map[int]*qtcore.QByteArray{
		QName:      qtcore.NewQByteArray2("name", -1),
		QAddresses: qtcore.NewQByteArray2("qaddresses", -1),
	})

	m.ConnectRowCount(m.rowCount)
	m.ConnectCleanModel(m.cleanModel)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectData(m.data)
	m.ConnectLoadModel(m.loadModel)
	m.ConnectAddAddresses(m.addAddresses)
	m.ConnectRemoveWallet(m.removeWallet)
	m.SetLoading(true)
	m.addresses = make([]*ModelAddresses, 0)
	altManager := local.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}

	m.WalletEnv = walletsEnvs[0]
	go func() {
		uptimeTicker := time.NewTicker(time.Duration(config.GetDataUpdateTime()) * time.Microsecond)

		for range uptimeTicker.C {
			go m.loadModel()
		}
	}()
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

	logWalletModel.Info("Loading Model")
	m.SetLoading(true)
	fullyLoad := true
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
		qml.QQmlEngine_SetObjectOwnership(ma, qml.QQmlEngine__CppOwnership)
		ma.SetName(wallets.Value().GetLabel())
		ma.SetId(wallets.Value().GetId())
		oModels := make([]*ModelOutputs, 0)

		for addresses.Next() {
			a := addresses.Value()
			outputs := a.GetCryptoAccount().ScanUnspentOutputs()
			if outputs == nil {
				logWalletModel.WithField("address", a.String()).Warn("Couldn't get unspent outputs")
				fullyLoad = false
				continue
			}
			mo := NewModelOutputs(nil)
			qml.QQmlEngine_SetObjectOwnership(mo, qml.QQmlEngine__CppOwnership)
			mo.SetAddress(a.String())
			qOutputs := make([]*QOutput, 0)

			for outputs.Next() {
				to := outputs.Value()
				qo := NewQOutput(nil)
				qml.QQmlEngine_SetObjectOwnership(qo, qml.QQmlEngine__CppOwnership)
				qo.SetOutputID(to.GetId())
				val, err := to.GetCoins(coin.Sky)
				if err != nil {
					logWalletModel.WithError(nil).Warn("Couldn't get " + coin.Sky + " coins")
					fullyLoad = false
					continue
				}
				accuracy, err := util.AltcoinQuotient(coin.Sky)
				if err != nil {
					logWalletModel.WithError(err).Warn("Couldn't get " + coin.Sky + " coins quotient")
					fullyLoad = false
					continue
				}
				coins := util.FormatCoins(val, accuracy)
				qo.SetAddressSky(coins)
				val, err = to.GetCoins(coin.CoinHoursTicker)
				if err != nil {
					logWalletModel.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins")
					fullyLoad = false
					continue
				}
				accuracy, err = util.AltcoinQuotient(coin.CoinHoursTicker)
				if err != nil {
					logWalletModel.WithError(err).Warn("Couldn't get " + coin.CoinHoursTicker + " coins quotient")
					fullyLoad = false
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
	Helper.RunInMain(func() {
		m.addAddresses(aModels)
	})
	if fullyLoad {
		m.SetLoading(false)
	}
}

func (m *ModelWallets) addAddresses(ma []*ModelAddresses) {
	for _, modelAddresses := range ma {
		row := 0
		find := false
		for _, modelASet := range m.addresses {
			if modelAddresses.Id() == modelASet.Id() {
				modelASet.addOutputs(modelAddresses.outputs)
				find = true
				break
			}
			row++
		}
		if !find {
			m.BeginInsertRows(qtcore.NewQModelIndex(), row, row)
			m.addressesMutex.Lock()
			m.addresses = append(m.addresses, modelAddresses)
			m.addressesMutex.Unlock()
			m.EndInsertRows()
		} else {
			m.DataChanged(m.Index(len(m.addresses)-1, row, qtcore.NewQModelIndex()), m.Index(len(m.addresses)-1, row+1, qtcore.NewQModelIndex()), []int{int(qtcore.Qt__DisplayRole)})
		}
	}
	for i := len(m.addresses) - 1; i >= 0; i-- {
		if len(m.addresses[i].outputs) < 1 {
			m.removeWallet(i)
		}
	}
}

func (m *ModelWallets) removeWallet(row int) {
	m.BeginRemoveRows(qtcore.NewQModelIndex(), row, row)
	m.addressesMutex.Lock()
	m.addresses = append(m.addresses[:row], m.addresses[row+1:]...)
	m.addressesMutex.Unlock()
	m.EndRemoveRows()
}
