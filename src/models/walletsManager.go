package models

import (
	"errors"
	"github.com/fibercrypto/F/src/coin/skycoin"
	"sync"

	"github.com/fibercrypto/FiberCryptoWallet/src/util"

	"github.com/therecipe/qt/qml"

	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	qtCore "github.com/therecipe/qt/core"
)

var logWalletManager = logging.MustGetLogger("modelsWalletManager")

var once sync.Once
var walletManager *WalletManager

type WalletManager struct {
	qtCore.QObject
	WalletEnv           core.WalletEnv
	SeedGenerator       core.SeedGenerator
	wallets             []*QWallet
	addresseseByWallets map[string][]*QAddress
	outputsByAddress    map[string][]*QOutput
	altManager          core.AltcoinManager
	_                   func()                                                                                                                                 `slot:"updateWalletEnvs"`
	_                   func(wltId, address string)                                                                                                            `slot:"updateOutputs"`
	_                   func(string)                                                                                                                           `slot:"updateAddresses"`
	_                   func()                                                                                                                                 `slot:"updateWallets"`
	_                   func()                                                                                                                                 `constructor:"init"`
	_                   func(seed string, label string, password string, scanN int) *QWallet                                                                   `slot:"createEncryptedWallet"`
	_                   func(seed string, label string, scanN int) *QWallet                                                                                    `slot:"createUnencryptedWallet"`
	_                   func(entropy int) string                                                                                                               `slot:"getNewSeed"`
	_                   func(seed string) int                                                                                                                  `slot:"verifySeed"`
	_                   func(id string, n int, password string)                                                                                                `slot:"newWalletAddress"`
	_                   func(id string, password string) int                                                                                                   `slot:"encryptWallet"`
	_                   func(id string, password string) int                                                                                                   `slot:"decryptWallet"`
	_                   func() []*QWallet                                                                                                                      `slot:"getWallets"`
	_                   func(id string) []*QAddress                                                                                                            `slot:"getAddresses"`
	_                   func(id string, source string, password string, index []int, qTxn *QTransaction) *QTransaction                                         `slot:"signTxn"`
	_                   func(wltId string, destinationAddress string, amount string) *QTransaction                                                             `slot:"sendTo"`
	_                   func(id, label string) *QWallet                                                                                                        `slot:"editWallet"`
	_                   func(wltId, address string) []*QOutput                                                                                                 `slot:"getOutputs"`
	_                   func(txn *QTransaction) bool                                                                                                           `slot:"broadcastTxn"`
	_                   func(wltId string, from, addrTo, skyTo, coinHoursTo []string, change string, automaticCoinHours bool, burnFactor string) *QTransaction `slot:"sendFromAddresses"`
	_                   func(wltId string, outs, addrTo, skyTo, coinHoursTo []string, change string, automaticCoinHours bool, burnFactor string) *QTransaction `slot:"sendFromOutputs"`
	_                   func() []*QAddress                                                                                                                     `slot:"getAllAddresses"`
	_                   func(wltId string) []*QOutput                                                                                                          `slot:"getOutputsFromWallet"`
}

func (walletM *WalletManager) init() {
	logWalletManager.Info("Initializing WalletManager")
	once.Do(func() {
		qml.QQmlEngine_SetObjectOwnership(walletM, qml.QQmlEngine__CppOwnership)
		walletM.ConnectCreateEncryptedWallet(walletM.createEncryptedWallet)
		walletM.ConnectCreateUnencryptedWallet(walletM.createUnencryptedWallet)
		walletM.ConnectGetNewSeed(walletM.getNewSeed)
		walletM.ConnectVerifySeed(walletM.verifySeed)
		walletM.ConnectNewWalletAddress(walletM.newWalletAddress)
		walletM.ConnectEncryptWallet(walletM.encryptWallet)
		walletM.ConnectDecryptWallet(walletM.decryptWallet)
		walletM.ConnectGetWallets(walletM.getWallets)
		walletM.ConnectGetAddresses(walletM.getAddresses)
		walletM.ConnectSendTo(walletM.sendTo)
		walletM.ConnectSignTxn(walletM.signTxn)
		walletM.ConnectGetOutputs(walletM.getOutputs)
		walletM.ConnectSendFromAddresses(walletM.sendFromAddresses)
		walletM.ConnectSendFromOutputs(walletM.sendFromOutputs)
		walletM.ConnectBroadcastTxn(walletM.broadcastTxn)
		walletM.ConnectGetAllAddresses(walletM.getAllAddresses)
		walletM.ConnectGetOutputsFromWallet(walletM.getOutputsFromWallet)
		walletM.ConnectUpdateWalletEnvs(walletM.updateWalletEnvs)
		walletM.ConnectUpdateWallets(walletM.updateWallets)
		walletM.ConnectUpdateAddresses(walletM.updateAddresses)
		walletM.ConnectUpdateOutputs(walletM.updateOutputs)
		walletM.addresseseByWallets = make(map[string][]*QAddress, 0)
		walletM.outputsByAddress = make(map[string][]*QOutput, 0)
		walletM.altManager = core.LoadAltcoinManager()

		walletM.SeedGenerator = new(sky.SeedService)
		walletManager = walletM
		walletM.updateWalletEnvs()
	})

	walletM = walletManager

}

func GetWalletEnv() core.WalletEnv {
	logWalletManager.Info("Getting WalletEnv")
	wm := GetWalletManager()
	if wm == nil {
		return nil
	}

	return wm.WalletEnv
}

func GetWalletManager() *WalletManager {
	return walletManager
}
func (walletM *WalletManager) updateWalletEnvs() {
	logWalletManager.Info("Updating WalletEnvs")
	walletsEnvs := make([]core.WalletEnv, 0)

	for _, plug := range walletM.altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}
	walletM.WalletEnv = walletsEnvs[0]
}

func (walletM *WalletManager) updateAddresses(wltId string) {
	logWalletManager.Info("Updating Addresses")
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
	qAddresses := make([]*QAddress, 0)
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		logWalletManager.WithError(err).Warn("Couldn't loaded addresses")
		walletM.addresseseByWallets[wltId] = qAddresses
		return
	}
	for it.Next() {
		addr := it.Value()
		qAddress := NewQAddress(nil)
		qml.QQmlEngine_SetObjectOwnership(qAddress, qml.QQmlEngine__CppOwnership)
		qAddress.SetAddress(addr.String())
		qAddress.SetMarked(0)
		qAddress.SetWallet(wlt.GetLabel())
		qAddress.SetWalletId(wlt.GetId())
		skyFl, err := addr.GetCryptoAccount().GetBalance(skycoin.SkycoinTicker)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't load " + skycoin.SkycoinTicker + " Balance")
			continue
		}
		accuracy, err := util.AltcoinQuotient(skycoin.SkycoinTicker)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't load " + skycoin.SkycoinTicker + " quotient")
			continue
		}
		qAddress.SetAddressSky(util.FormatCoins(skyFl, accuracy))
		coinH, err := addr.GetCryptoAccount().GetBalance(skycoin.CoinHoursTicker)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't load " + skycoin.CoinHoursTicker + " Balance")
			continue
		}
		accuracy, err = util.AltcoinQuotient(skycoin.CoinHoursTicker)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't load " + skycoin.CoinHoursTicker + " quotient")
			continue
		}
		qAddress.SetAddressCoinHours(util.FormatCoins(coinH, accuracy))
		qml.QQmlEngine_SetObjectOwnership(qAddress, qml.QQmlEngine__CppOwnership)

		qAddresses = append(qAddresses, qAddress)

	}

	walletM.addresseseByWallets[wltId] = qAddresses
}

func (walletM *WalletManager) updateOutputs(wltId, address string) {
	outs := make([]*QOutput, 0)
	addressIterator, err := walletM.WalletEnv.GetWalletSet().GetWallet(wltId).GetLoadedAddresses()
	if err != nil {
		logWalletManager.WithError(err).Warn("Couldn't get an addresses iterator")
		walletM.outputsByAddress[address] = outs
		return
	}
	var addr core.Address
	for addressIterator.Next() {
		if addressIterator.Value().String() == address {
			addr = addressIterator.Value()
			break
		}
	}
	if addr == nil {
		logWalletManager.WithError(err).Warn("Couldn't get address")
		walletM.outputsByAddress[address] = outs
		return
	}
	outsIter := addr.GetCryptoAccount().ScanUnspentOutputs()
	if outsIter == nil {
		logWalletManager.WithError(err).Warn("Couldn't scan unspent outputs")
		walletM.outputsByAddress[address] = outs
		return
	}
	for outsIter.Next() {
		qout := NewQOutput(nil)
		qml.QQmlEngine_SetObjectOwnership(qout, qml.QQmlEngine__CppOwnership)
		qout.SetOutputID(outsIter.Value().GetId())
		skyV, err := outsIter.Value().GetCoins(sky.Sky)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't get " + sky.Sky + " coins")
			continue
		}
		quotient, err := util.AltcoinQuotient(sky.Sky)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't get " + sky.Sky + " quotient")
			continue
		}
		sSky := util.FormatCoins(skyV, quotient)
		qout.SetAddressSky(sSky)
		ch, err := outsIter.Value().GetCoins(sky.CoinHour)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't get " + sky.CoinHour + " coins")
			continue
		}
		quotient, err = util.AltcoinQuotient(sky.CoinHour)
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't get " + sky.Sky + " quotient")
			continue
		}
		sCh := util.FormatCoins(ch, quotient)
		qout.SetAddressCoinHours(sCh)
		qout.SetAddressOwner(addr.String())
		qout.SetWalletOwner(wltId)
		outs = append(outs, qout)
	}

	walletM.outputsByAddress[address] = outs
}

func (walletM *WalletManager) updateWallets() {
	qWallets := make([]*QWallet, 0)
	it := walletM.WalletEnv.GetWalletSet().ListWallets()
	if it == nil {
		logWalletManager.WithError(nil).Warn("Couldn't get a wallet iterator")
		return
	}
	for it.Next() {

		encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(it.Value().GetId())
		if err != nil {
			logWalletManager.WithError(err).Warn("Couldn't get wallet by id")
			continue
		}
		if encrypted {
			qw := fromWalletToQWallet(it.Value(), true)
			qWallets = append(qWallets, qw)
		} else {
			qw := fromWalletToQWallet(it.Value(), false)
			qWallets = append(qWallets, qw)
		}

	}
	walletM.wallets = qWallets
}

func (walletM *WalletManager) getAllAddresses() []*QAddress {
	logWalletManager.Info("Getting all addresses")
	qAddresses := make([]*QAddress, 0)
	wltIter := walletM.WalletEnv.GetWalletSet().ListWallets()
	if wltIter == nil {
		logWalletManager.Warn("Error getting the list of wallets")
		return nil
	}
	for wltIter.Next() {
		qAddresses = append(qAddresses, walletM.getAddresses(wltIter.Value().GetId())...)
	}
	return qAddresses
}
func (walletM *WalletManager) broadcastTxn(txn *QTransaction) bool {
	logWalletManager.Info("Broadcasting transaction")
	altManager := core.LoadAltcoinManager()
	plug, _ := altManager.LookupAltcoinManager(skycoin.SkycoinTicker)
	pex, err := plug.LoadPEX("MainNet")
	if err != nil {
		logWalletManager.WithError(err).Warn("Error loading PEX")
		return false
	}
	err = pex.BroadcastTxn(txn.txn)
	if err != nil {
		logWalletManager.WithError(err).Warn("Error broadcasting transaction")
		return false
	}
	logWalletManager.Info("Transaction Injected")
	return true
}

func (walletM *WalletManager) sendFromOutputs(wltId string, from, addrTo, skyTo, coinHoursTo []string, change string, automaticCoinHours bool, burnFactor string) *QTransaction {
	logWalletManager.Info("Creating transaction")
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
	if wlt == nil {
		logWalletManager.Warn("Couldn't load wallet to create transaction")
		return nil
	}
	outputsFrom := make([]core.TransactionOutput, 0)
	for _, out := range from {
		outputsFrom = append(outputsFrom, &GenericOutput{
			addr: out,
		})
	}
	outputsTo := make([]core.TransactionOutput, 0)
	for i := 0; i < len(addrTo); i++ {
		ch := ""
		if !automaticCoinHours {
			ch = coinHoursTo[i]
		}
		outputsTo = append(outputsTo, &GenericOutput{
			addr: addrTo[i],
			sky:  skyTo[i],
			ch:   ch,
		})
	}
	changeAddr := &GenericAddress{change}
	opt := NewTransferOptions()
	opt.AddKeyValue("BurnFactor", burnFactor)
	if automaticCoinHours {
		opt.AddKeyValue("CoinHoursSelectionType", "auto")
	} else {
		opt.AddKeyValue("CoinHoursSelectionType", "manual")
	}

	txn, err := wlt.Spend(outputsFrom, outputsTo, changeAddr, opt)
	if err != nil {
		logWalletManager.WithError(err).Info("Error creating transaction")
		return nil
	}

	qTransaction, err := NewQTransactionFromTransaction(txn)
	if err != nil {
		logWalletManager.WithError(err).Info("Error converting transaction")
		return nil
	}
	return qTransaction
}
func (walletM *WalletManager) sendFromAddresses(wltId string, from, addrTo, skyTo, coinHoursTo []string, change string, automaticCoinHours bool, burnFactor string) *QTransaction {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
	if wlt == nil {
		logWalletManager.Warn("Couldn't load wallet to create transaction")
		return nil
	}
	addrsFrom := make([]core.Address, 0)
	for _, addr := range from {

		addrsFrom = append(addrsFrom, &GenericAddress{addr})
	}
	outputsTo := make([]core.TransactionOutput, 0)
	for i := 0; i < len(addrTo); i++ {
		ch := ""
		if !automaticCoinHours {
			ch = coinHoursTo[i]
		}
		outputsTo = append(outputsTo, &GenericOutput{
			addr: addrTo[i],
			sky:  skyTo[i],
			ch:   ch,
		})
	}
	changeAddr := &GenericAddress{change}

	opt := NewTransferOptions()
	opt.AddKeyValue("BurnFactor", burnFactor)
	if automaticCoinHours {
		opt.AddKeyValue("CoinHoursSelectionType", "auto")
	} else {
		opt.AddKeyValue("CoinHoursSelectionType", "manual")
	}

	txn, err := wlt.SendFromAddress(addrsFrom, outputsTo, changeAddr, opt)
	if err != nil {
		logWalletManager.WithError(err).Info("Error creating transaction")
		return nil
	}

	qtxn, err := NewQTransactionFromTransaction(txn)
	if err != nil {
		logWalletManager.WithError(err).Info("Error converting transaction")
		return nil
	}
	logWalletManager.Info("Transaction created")
	return qtxn

}

func (walletM *WalletManager) getOutputs(wltId, address string) []*QOutput {
	outs, ok := walletM.outputsByAddress[address]
	if !ok {
		walletM.updateOutputs(wltId, address)
	}
	return outs
}

func (walletM *WalletManager) getOutputsFromWallet(wltId string) []*QOutput {
	logWalletManager.Info("Getting Outputs from wallet by Id")
	outs := make([]*QOutput, 0)
	addrIter, err := walletM.WalletEnv.GetWalletSet().GetWallet(wltId).GetLoadedAddresses()
	if err != nil {
		logWalletManager.WithError(err).Warn("Couldn't load addresses iterator")
		return nil
	}
	for addrIter.Next() {
		outs = append(outs, walletM.getOutputs(wltId, addrIter.Value().String())...)
	}
	logWalletManager.Info("Loaded all outputs")
	return outs
}

func (walletM *WalletManager) sendTo(wltId, destinationAddress, amount string) *QTransaction {
	logWalletManager.Info("Creating Transaction")
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
	addr := &GenericAddress{
		Address: destinationAddress,
	}
	coins, err := util.GetCoinValue(amount, skycoin.SkycoinTicker)
	if err != nil {
		logWalletManager.WithError(err).Warn("Couldn't get Skycoin's")
		return nil
	}
	opt := NewTransferOptions()
	opt.AddKeyValue("BurnFactor", "0.5")
	opt.AddKeyValue("CoinHoursSelectionType", "auto")
	if wlt == nil {
		logWalletManager.Warn("Couldn't load wallet to create transaction")
		return nil
	}
	txn, err := wlt.Transfer(addr, coins, opt)
	if err != nil {
		logWalletManager.WithError(err).Warn("Couldn't create transaction")
		return nil
	}
	qTxn, err := NewQTransactionFromTransaction(txn)
	if err != nil {
		logWalletManager.WithError(err).Warn("Couldn't convert transaction")
		return nil
	}
	logWalletManager.Info("Transaction created")
	return qTxn

}

func (walletM *WalletManager) signTxn(id, source, password string, index []int, qTxn *QTransaction) *QTransaction {
	logWalletManager.Info("Signig transaction")
	// Get wallet
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	if wlt == nil {
		logWalletManager.Warn("Couldn't load wallet to Sign transaction")
		return nil
	}
	txn, err := wlt.Sign(qTxn.txn, source, func(message string) (string, error) {
		return password, nil
	}, nil) // TODO Get index for sign specific txn indexes
	if err != nil {
		logWalletManager.WithError(err).Warn("Error signing txn")
		return nil
	}
	qTxn, err = NewQTransactionFromTransaction(txn)
	if err != nil {
		logWalletManager.WithError(err).Warn("Error converting transaction")
		return nil
	}
	return qTxn
}

func (walletM *WalletManager) createEncryptedWallet(seed, label, password string, scanN int) *QWallet {
	logWalletManager.Info("Creating encrypted wallet")
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, true, pwd, scanN)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't create encrypted wallet")
		return nil
	}

	logWalletManager.Info("Created encrypted wallet")
	return fromWalletToQWallet(wlt, true)

}

func (walletM *WalletManager) createUnencryptedWallet(seed, label string, scanN int) *QWallet {
	logWalletManager.Info("Creating encrypted wallet")
	pwd := func(message string) (string, error) {
		return "", nil
	}

	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, false, pwd, scanN)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't create unencrypted wallet")
		return nil
	}
	logWalletManager.Info("Created encrypted wallet")
	return fromWalletToQWallet(wlt, false)

}

func (walletM *WalletManager) getNewSeed(entropy int) string {
	logWalletManager.Info("Getting new seed")
	seed, err := walletM.SeedGenerator.GenerateMnemonic(entropy)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get new seed")
		return ""
	}
	logWalletManager.Info("Generated new seed")
	return seed
}

func (walletM *WalletManager) verifySeed(seed string) int {
	logWalletManager.Info("Verifying seed")
	ok, err := walletM.SeedGenerator.VerifyMnemonic(seed)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't verify seed")
		return 0
	}
	logWalletManager.Info("Verified seed")
	if ok {
		return 1
	}
	return 0

}

func (walletM *WalletManager) encryptWallet(id, password string) int {
	logWalletManager.Info("Encrypting wallet")
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Encrypt(id, pwd)
	ret, err := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't create encrypted wallets")
	}
	logWalletManager.Info("Wallet encrypted")
	if ret {
		return 1
	}
	return 0
}

func (walletM *WalletManager) decryptWallet(id, password string) int {
	logWalletManager.Info("Decrypt wallet")
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Decrypt(id, pwd)
	ret, err := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't decrypt wallet")
	}
	logWalletManager.Info("Wallet decrypted")
	if ret {
		return 1
	}
	return 0
}

func (walletM *WalletManager) newWalletAddress(id string, n int, password string) {
	logWalletManager.Info("Creating new wallet addresses")
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wltEntriesLen := 0
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't load addresses")
		return
	}
	for it.Next() {
		wltEntriesLen++
	}
	wlt.GenAddresses(core.AccountAddress, uint32(wltEntriesLen), uint32(n), pwd)
	logWalletManager.Info("New addresses created")
}

func (walletM *WalletManager) getWallets() []*QWallet {
	if walletM.wallets == nil {
		walletM.updateWallets()
	}
	logWalletManager.Info("Getting wallets")
	qWallets := make([]*QWallet, 0)
	if walletM.WalletEnv == nil {
		walletM.UpdateWalletEnvs()
	}
	it := walletM.WalletEnv.GetWalletSet().ListWallets()

	if it == nil {
		logWalletManager.WithError(nil).Error("Couldn't load wallets")
		return qWallets

	}

	for it.Next() {

		encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(it.Value().GetId())
		if err != nil {
			logWalletManager.WithError(err).Error("Couldn't get wallets")
			return qWallets
		}
		if encrypted {
			qw := fromWalletToQWallet(it.Value(), true)
			qWallets = append(qWallets, qw)
		} else {
			qw := fromWalletToQWallet(it.Value(), false)
			qWallets = append(qWallets, qw)
		}

	}

	logWalletManager.Info("Wallets obtained")
	return walletM.wallets
}

func (walletM *WalletManager) editWallet(id, label string) *QWallet {
	logWalletManager.Info("Editing wallet")

	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	wlt.SetLabel(label)
	wlt = walletM.WalletEnv.GetWalletSet().GetWallet(id)
	encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(wlt.GetId())
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't edit wallet")
		return nil
	}
	qWallet := fromWalletToQWallet(wlt, encrypted)
	logWalletManager.Info("Wallet edited")
	return qWallet
}

func (walletM *WalletManager) getAddresses(Id string) []*QAddress {
	addrs, ok := walletM.addresseseByWallets[Id]
	if !ok {
		walletM.updateAddresses(Id)
		addrs = walletM.addresseseByWallets[Id]
	}
	return addrs
}

func fromWalletToQWallet(wlt core.Wallet, isEncrypted bool) *QWallet {

	qWallet := NewQWallet(nil)
	qml.QQmlEngine_SetObjectOwnership(qWallet, qml.QQmlEngine__CppOwnership)
	qWallet.SetName(wlt.GetLabel())

	qWallet.SetFileName(wlt.GetId())

	qWallet.SetEncryptionEnabled(0)
	if isEncrypted {
		qWallet.SetEncryptionEnabled(1)
	}

	bl, err := wlt.GetCryptoAccount().GetBalance(sky.Sky)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Skycoin balance")
		return qWallet
	}

	accuracy, err := util.AltcoinQuotient(params.SkycoinTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Skycoin Altcoin quotient")
		return qWallet
	}

	accuracy, err = util.AltcoinQuotient(params.SkycoinTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Skycoin Altcoin quotient")
		return qWallet
	}
	floatBl := float64(bl) / float64(accuracy)
	qWallet.SetSky(floatBl)

	bl, err = wlt.GetCryptoAccount().GetBalance(params.CoinHoursTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Coin Hours balance")
		return qWallet
	}
	accuracy, err = util.AltcoinQuotient(params.SkycoinTicker)
	if err != nil {
		logWalletManager.WithError(err).Error("Couldn't get Coin Hours Altcoin quotient")
		return qWallet
	}
	qWallet.SetCoinHours(bl)

	return qWallet
}

type GenericOutput struct {
	addr string
	sky  string
	ch   string
}

func (gOut *GenericOutput) GetId() string {
	return ""
}
func (gOut *GenericOutput) IsSpent() bool {
	return false
}
func (gOut *GenericOutput) GetAddress() core.Address {
	return &GenericAddress{Address: gOut.addr}
}
func (gOut *GenericOutput) GetCoins(ticker string) (uint64, error) {
	if ticker == sky.Sky {
		val, err := util.GetCoinValue(gOut.sky, ticker)
		if err != nil {
			return 0, err
		}
		return val, nil
	}
	if ticker == sky.CoinHour {
		val, err := util.GetCoinValue(gOut.ch, ticker)
		if err != nil {
			return 0, err
		}
		return val, nil
	}

	return 0, errors.New("invalid ticker")
}

type TransferOptions struct {
	values map[string]interface{}
}

func (tOpt *TransferOptions) GetValue(key string) interface{} {
	return tOpt.values[key]
}

func (tOpt *TransferOptions) AddKeyValue(key string, value interface{}) {
	tOpt.values[key] = value
}

func NewTransferOptions() *TransferOptions {
	tOptions := TransferOptions{
		values: make(map[string]interface{}, 0),
	}
	return &tOptions
}

type GenericAddress struct {
	Address string
}

func (ga *GenericAddress) IsBip32() bool {
	return true
}

func (ga *GenericAddress) String() string {
	return ga.Address
}

func (ga *GenericAddress) GetCryptoAccount() core.CryptoAccount {
	return nil
}
