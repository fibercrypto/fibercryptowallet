package models

import (
	"errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/sirupsen/logrus"

	"github.com/therecipe/qt/qml"

	sky "github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/models"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	qtcore "github.com/therecipe/qt/core"
)

type WalletManager struct {
	qtcore.QObject
	WalletEnv     core.WalletEnv
	SeedGenerator core.SeedGenerator

	_ func()                                                                                                                                 `constructor:"init"`
	_ func(seed string, label string, password string, scanN int) *QWallet                                                                   `slot:"createEncryptedWallet"`
	_ func(seed string, label string, scanN int) *QWallet                                                                                    `slot:"createUnencryptedWallet"`
	_ func(entropy int) string                                                                                                               `slot:"getNewSeed"`
	_ func(seed string) int                                                                                                                  `slot:"verifySeed"`
	_ func(id string, n int, password string)                                                                                                `slot:"newWalletAddress"`
	_ func(id string, password string) int                                                                                                   `slot:"encryptWallet"`
	_ func(id string, password string) int                                                                                                   `slot:"decryptWallet"`
	_ func() []*QWallet                                                                                                                      `slot:"getWallets"`
	_ func(id string) []*QAddress                                                                                                            `slot:"getAddresses"`
	_ func(id string, source string, password string, index []int, qTxn *QTransaction) *QTransaction                                         `slot:"signTxn"`
	_ func(wltId string, destinationAddress string, amount string) *QTransaction                                                             `slot:"sendTo"`
	_ func(id, label string) *QWallet                                                                                                        `slot:"editWallet"`
	_ func(wltId, address string) []*QOutput                                                                                                 `slot:"getOutputs"`
	_ func(txn *QTransaction) bool                                                                                                           `slot:"broadcastTxn"`
	_ func(wltId string, from, addrTo, skyTo, coinHoursTo []string, change string, automaticCoinHours bool, burnFactor string) *QTransaction `slot:"sendFromAddresses"`
	//_ func(wltId string, outs, addrTo, skyTo, coinHoursTo []string, change, automaticCoinHours bool, burnFactor string, password string) `slot:"sendFromOutputs"`
}

func (walletM *WalletManager) init() {
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
	//walletM.ConnectSendFromOutputs(walletM.sendFromOutputs)
	walletM.ConnectBroadcastTxn(walletM.broadcastTxn)
	altManager := core.LoadAltcoinManager()
	walletsEnvs := make([]core.WalletEnv, 0)
	for _, plug := range altManager.ListRegisteredPlugins() {
		walletsEnvs = append(walletsEnvs, plug.LoadWalletEnvs()...)
	}

	walletM.WalletEnv = walletsEnvs[0]

	walletM.SeedGenerator = new(sky.SeedService)

}

func (walletM *WalletManager) broadcastTxn(txn *QTransaction) bool {
	altManager := core.LoadAltcoinManager()
	plug, _ := altManager.LookupAltcoinManager("SKY")
	pex, err := plug.LoadPEX("MainNet")
	if err != nil {
		return false
	}
	err = pex.BroadcastTxn(txn.txn)
	if err != nil {
		return false
	}
	logrus.Info("Txn Injected")
	return true
}

func (walletM *WalletManager) sendFromAddresses(wltId string, from, addrTo, skyTo, coinHoursTo []string, change string, automaticCoinHours bool, burnFactor string) *QTransaction {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
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
	opt.AddKeyValue("CoinHoursSelectionMode", automaticCoinHours)

	txn, err := wlt.SendFromAddress(addrsFrom, outputsTo, changeAddr, opt)
	if err != nil {
		return nil
	}

	qtxn, err := NewQTransactionFromTransaction(txn)
	if err != nil {
		return nil
	}
	return qtxn

}

func (walletM *WalletManager) getOutputs(wltId, address string) []*QOutput {
	addrIter, err := walletM.WalletEnv.GetWalletSet().GetWallet(wltId).GetLoadedAddresses()
	if err != nil {
		return nil
	}
	outs := make([]*QOutput, 0)
	var addr core.Address
	for addrIter.Next() {
		if addrIter.Value().String() == address {
			addr = addrIter.Value()
			break
		}
	}
	outsIter := addr.GetCryptoAccount().ScanUnspentOutputs()
	for outsIter.Next() {
		qout := NewQOutput(nil)
		qout.SetOutputID(outsIter.Value().GetId())
		skyV, err := outsIter.Value().GetCoins(sky.Sky)
		if err != nil {
			return nil
		}
		quotient, err := util.AltcoinQuotient(sky.Sky)
		if err != nil {
			return nil
		}
		sSky := util.FormatCoins(skyV, quotient)
		qout.SetAddressSky(sSky)
		ch, err := outsIter.Value().GetCoins(sky.CoinHour)
		if err != nil {
			return nil
		}
		quotient, err = util.AltcoinQuotient(sky.CoinHour)
		if err != nil {
			return nil
		}
		sCh := util.FormatCoins(ch, quotient)
		qout.SetAddressCoinHours(sCh)
		outs = append(outs, qout)
	}
	return outs
}

func (walletM *WalletManager) sendTo(wltId, destinationAddress, amount string) *QTransaction {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(wltId)
	addr := &GenericAddress{
		Address: destinationAddress,
	}
	coins, err := util.GetCoinValue(amount, "SKY")
	if err != nil {
		return nil
	}
	txn, err := wlt.Transfer(addr, coins, nil)

	if err != nil {
		return nil
	}
	qtxn, err := NewQTransactionFromTransaction(txn)
	if err != nil {
		return nil
	}
	return qtxn

}

func (walletM *WalletManager) signTxn(id, source, password string, index []int, qTxn *QTransaction) *QTransaction {
	// Get wallet
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)

	txn, err := wlt.Sign(qTxn.txn, source, func(message string) (string, error) {
		return password, nil
	}, nil) // TODO Get index for sign specific txn indexes
	if err != nil {
		logrus.Warn("Error signing txn", err)
		return nil
	}
	qTxn, err = NewQTransactionFromTransaction(txn)
	if err != nil {
		return nil
	}
	return qTxn
}

func (walletM *WalletManager) createEncryptedWallet(seed, label, password string, scanN int) *QWallet {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, true, pwd, scanN)
	if err != nil {
		return nil
	}

	return fromWalletToQWallet(wlt, true)

}

func (walletM *WalletManager) createUnencryptedWallet(seed, label string, scanN int) *QWallet {
	pwd := func(message string) (string, error) {
		return "", nil
	}

	wlt, err := walletM.WalletEnv.GetWalletSet().CreateWallet(label, seed, false, pwd, scanN)
	if err != nil {
		return nil
	}
	return fromWalletToQWallet(wlt, false)

}

func (walletM *WalletManager) getNewSeed(entropy int) string {
	seed, err := walletM.SeedGenerator.GenerateMnemonic(entropy)
	if err != nil {
		return ""
	}
	return seed
}

func (walletM *WalletManager) verifySeed(seed string) int {
	ok, err := walletM.SeedGenerator.VerifyMnemonic(seed)
	if err != nil {
		return 0
	}
	if ok {
		return 1
	}
	return 0

}

func (walletM *WalletManager) encryptWallet(id, password string) int {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Encrypt(id, pwd)
	ret, _ := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if ret {
		return 1
	}
	return 0
}

func (walletM *WalletManager) decryptWallet(id, password string) int {
	pwd := func(message string) (string, error) {
		return password, nil
	}
	walletM.WalletEnv.GetStorage().Decrypt(id, pwd)
	ret, _ := walletM.WalletEnv.GetStorage().IsEncrypted(id)
	if ret {
		return 1
	}
	return 0
}

func (walletM *WalletManager) newWalletAddress(id string, n int, password string) {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	pwd := func(message string) (string, error) {
		return password, nil
	}
	wltEntrieslen := 0
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		return
	}
	for it.Next() {
		wltEntrieslen++
	}
	wlt.GenAddresses(core.AccountAddress, uint32(wltEntrieslen), uint32(n), pwd)
}

func (walletM *WalletManager) getWallets() []*QWallet {

	qwallets := make([]*QWallet, 0)
	it := walletM.WalletEnv.GetWalletSet().ListWallets()

	for it.Next() {

		encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(it.Value().GetId())
		if err != nil {
			continue
		}
		if encrypted {
			qw := fromWalletToQWallet(it.Value(), true)
			qwallets = append(qwallets, qw)
		} else {
			qw := fromWalletToQWallet(it.Value(), false)
			qwallets = append(qwallets, qw)
		}

	}

	return qwallets

}

func (walletM *WalletManager) editWallet(id, label string) *QWallet {
	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(id)
	wlt.SetLabel(label)
	wlt = walletM.WalletEnv.GetWalletSet().GetWallet(id)
	encrypted, err := walletM.WalletEnv.GetStorage().IsEncrypted(wlt.GetId())
	if err != nil {
		return nil
	}
	qwallet := fromWalletToQWallet(wlt, encrypted)
	return qwallet
}

func (walletM *WalletManager) getAddresses(Id string) []*QAddress {

	wlt := walletM.WalletEnv.GetWalletSet().GetWallet(Id)
	qaddresses := make([]*QAddress, 0)
	it, err := wlt.GetLoadedAddresses()
	if err != nil {
		return nil
	}
	for it.Next() {
		addr := it.Value()
		qAddress := NewQAddress(nil)
		qml.QQmlEngine_SetObjectOwnership(qAddress, qml.QQmlEngine__CppOwnership)
		qAddress.SetAddress(addr.String())
		qAddress.SetMarked(0)
		coins, err := addr.GetCryptoAccount().GetBalance("SKY")
		if err != nil {

			continue
		}
		//TODO: report possible error
		accuracy, _ := util.AltcoinQuotient("SKY")
		qaddress.SetAddressSky(util.FormatCoins(sky, accuracy))
		coinH, err := addr.GetCryptoAccount().GetBalance("SKYCH")
		accuracy, _ = util.AltcoinQuotient("SKYCH")
		if err != nil {

			continue
		}
		qaddress.SetAddressCoinHours(util.FormatCoins(coinH, accuracy))
		qml.QQmlEngine_SetObjectOwnership(qAddress, qml.QQmlEngine__CppOwnership)

		qaddresses = append(qaddresses, qAddress)

	}

	return qaddresses
}

func fromWalletToQWallet(wlt core.Wallet, isEncrypted bool) *QWallet {

	qwallet := NewQWallet(nil)
	qml.QQmlEngine_SetObjectOwnership(qwallet, qml.QQmlEngine__CppOwnership)
	qwallet.SetName(wlt.GetLabel())

	qwallet.SetFileName(wlt.GetId())

	qwallet.SetEncryptionEnabled(0)
	if isEncrypted {
		qwallet.SetEncryptionEnabled(1)
	}

	bl, err := wlt.GetCryptoAccount().GetBalance(sky.Sky)
	if err != nil {
		bl = 0
	}

	//TODO: report possible error
	accuracy, _ := util.AltcoinQuotient(skycoin.SkycoinTicker)
	floatBl := float64(bl) / float64(accuracy)
	qwallet.SetSky(floatBl)

	bl, err = wlt.GetCryptoAccount().GetBalance(skycoin.CoinHoursTicker)
	accuracy, _ = util.AltcoinQuotient(skycoin.SkycoinTicker)
	if err != nil {
		bl = 0
	}
	qwallet.SetCoinHours(bl)

	return qwallet
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
		values: make(map[string]interface{},0),
	}
	tOptions.AddKeyValue("CoinHoursMode", "auto")
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
