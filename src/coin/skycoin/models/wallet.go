package skycoin

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/wallet"
)

const (
	Sky                     = params.SkycoinTicker
	CoinHour                = params.CoinHoursTicker
	WalletTypeDeterministic = "deterministic"

	WalletTypeCollection = "collection"

	WalletTypeBip44 = "bip44"

	WalletTypeXPub        = "xpub"
	walletExt             = ".wlt"
	WalletTimestampFormat = "2006_01_02"
)

type SkycoinWalletIterator struct {
	//Implements WalletIterator interface
	current int
	wallets []core.Wallet
}

func (it *SkycoinWalletIterator) Value() core.Wallet {
	return it.wallets[it.current]
}

func (it *SkycoinWalletIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *SkycoinWalletIterator) HasNext() bool {
	if (it.current + 1) >= len(it.wallets) {
		return false
	}
	return true
}

func NewSkycoinWalletIterator(wallets []core.Wallet) *SkycoinWalletIterator {
	return &SkycoinWalletIterator{wallets: wallets, current: -1}
}

type SkycoinRemoteWallet struct {
	//Implements WalletStorage and WalletSet interfaces
	poolSection string
}

func (wltSrv *SkycoinRemoteWallet) ListWallets() core.WalletIterator {

	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		return nil
	}
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)

	wlts, err := c.Wallets()
	if err != nil {
		return nil
	}
	wallets := make([]core.Wallet, 0)
	for _, wlt := range wlts {
		nwlt := walletResponseToWallet(wlt)
		nwlt.poolSection = wltSrv.poolSection
		wallets = append(wallets, nwlt)
	}

	return NewSkycoinWalletIterator(wallets)
}

func (wltSrv *SkycoinRemoteWallet) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	wlt := RemoteWallet{}
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	if IsEncrypted {
		password, _ := pwd("Enter your password")
		wltOpt := api.CreateWalletOptions{}
		wltOpt.Type = WalletTypeDeterministic
		wltOpt.Seed = seed
		wltOpt.Password = password
		wltOpt.Encrypt = true
		wltOpt.Label = label
		wltOpt.ScanN = scanAddressesN

		wltR, err := c.CreateWallet(wltOpt)
		if err != nil {
			return nil, err
		}
		wlt = walletResponseToWallet(*wltR)
	} else {
		wltOpt := api.CreateWalletOptions{}
		wltOpt.Type = WalletTypeDeterministic
		wltOpt.Seed = seed
		wltOpt.Encrypt = false
		wltOpt.Label = label
		wltOpt.ScanN = scanAddressesN

		wltR, err := c.CreateWallet(wltOpt)
		if err != nil {
			return nil, err
		}
		wlt = walletResponseToWallet(*wltR)
	}
	wlt.poolSection = wltSrv.poolSection
	return &wlt, nil
}

func (wltSrv *SkycoinRemoteWallet) Encrypt(walletName string, pwd core.PasswordReader) {
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		return
	}
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	password, _ := pwd("Insert password")
	_, err = c.EncryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) Decrypt(walletName string, pwd core.PasswordReader) {
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		return
	}
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	password, _ := pwd("Insert password")
	_, err = c.DecryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) IsEncrypted(walletName string) (bool, error) {
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		return false, err
	}
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	wlt, err := c.Wallet(walletName)
	if err != nil {
		return false, err
	}
	return wlt.Meta.Encrypted, nil
}
func (wltSrv *SkycoinRemoteWallet) GetWallet(id string) core.Wallet {
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		return nil
	}
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	wltR, err := c.Wallet(id)
	if err != nil {
		return nil
	}
	nwlt := walletResponseToWallet(*wltR)
	nwlt.poolSection = wltSrv.poolSection
	return nwlt
}

func NewWalletNode(nodeAddress string) *WalletNode {

	pool := core.GetMultiPool()
	sections := pool.ListSections()
	cont := 1
	var sect string
	for true {
		find := false
		sect = fmt.Sprintf("skycoin-%d", cont)
		for _, sec := range sections {
			if sec == sect {
				find = true
				break
			}
		}
		cont++
		if !find {
			break
		}
	}

	pool.CreateSection(sect, NewSkycoinConnectionFactory(nodeAddress))
	return &WalletNode{
		NodeAddress: nodeAddress,
		poolSection: sect,
	}

}

type WalletNode struct {
	//Implements WallentEnv interface
	wltService  *SkycoinRemoteWallet
	NodeAddress string
	poolSection string
}

func (wltEnv *WalletNode) GetStorage() core.WalletStorage {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteWallet)
		wltEnv.wltService.poolSection = wltEnv.poolSection
	}
	return wltEnv.wltService
}

func (wltEnv *WalletNode) GetWalletSet() core.WalletSet {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteWallet)
		wltEnv.wltService.poolSection = wltEnv.poolSection
	}
	return wltEnv.wltService
}

type SeedService struct{} //Implements SeedGenerator interface

func (seedService *SeedService) GenerateMnemonic(entropyBits int) (string, error) {
	if entropyBits != 128 && entropyBits != 256 {
		return "", errors.New("Entropy must be 128 or 256")
	}

	entropy, err := bip39.NewEntropy(entropyBits)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

func (seedService *SeedService) VerifyMnemonic(seed string) (bool, error) {
	err := bip39.ValidateMnemonic(seed)
	if err != nil {
		return false, err
	}
	return true, nil
}

type errorTickerInvalid struct {
	tickerUsed string
}

func (err errorTickerInvalid) Error() string {
	return (err.tickerUsed + " is an invalid ticker. Use " + Sky + " or " + CoinHour)
}

type RemoteWallet struct {
	//Implements Wallet and CryptoAccount interfaces
	Id          string
	Label       string
	CoinType    string
	Encrypted   bool
	poolSection string
}

func (wlt RemoteWallet) Sign(txn core.Transaction, source string, pwd core.PasswordReader, index []int) (core.Transaction, error) {
	client, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, client)
	unInjectedTransaction, ok := txn.(*SkycoinUninjectedTransaction)
	if !ok {
		logrus.Warn(err)
		return nil, err
	}

	password, err := pwd("Encryted")
	if err != nil {
		logrus.Warn("Error getting password")
		return nil, err
	}
	encodedResponse, err := unInjectedTransaction.txn.SerializeHex()
	if err != nil {
		logrus.Warn("Couldn't get Transaction Encoded")
		return nil, err
	}
	walletSignTxn := api.WalletSignTransactionRequest{
		EncodedTransaction: encodedResponse,
		WalletID:           wlt.Id,
		Password:           password,
		SignIndexes:        index,
	}
	switch source {
	// Add selector for sign with a specific kind of wallet
	}
	txnResponse, err := client.WalletSignTransaction(walletSignTxn)
	if err != nil {
		logrus.Warn("Error signing transaction")
		return nil, err
	}
	skyTxn, err := txnResponse.Transaction.ToTransaction()
	if err != nil {
		return nil, err
	}
	value, err := util.GetCoinValue(txnResponse.Transaction.Fee, CoinHour)
	if err != nil {
		logrus.Warn("Error getting fee")
		return nil, nil
	}
	unTxn, err := NewUninjectedTransaction(skyTxn, value)
	if err != nil {
		return nil, err
	}
	return unTxn, nil
}

func (wlt RemoteWallet) Inject(rawTxn string) error {
	client, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logrus.Warn(err)
		return err
	}
	defer core.GetMultiPool().Return(PoolSection, client)

	_, err = client.InjectEncodedTransaction(rawTxn)
	if err != nil {
		logrus.Warn("Error Injecting transaction")
		return err
	}
	return nil
}

func (wlt RemoteWallet) GetLabel() string {
	return wlt.Label
}

func (wlt RemoteWallet) SetLabel(name string) {
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		return
	}

	defer core.GetMultiPool().Return(wlt.poolSection, c)
	_ = c.UpdateWallet(wlt.Id, name)
}

func (wlt RemoteWallet) GetId() string {
	return wlt.Id
}

func (wlt RemoteWallet) Transfer(to core.Address, amount uint64, options core.KeyValueStorage) (core.Transaction, error) {
	logrus.Info("Transfer from remote wallet")
	client, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, client)
	wltR, err := client.Wallet(wlt.Id)

	if err != nil {
		logrus.Warn("Error getting remote wallet")
		return nil, err
	}
	var txnOutput SkycoinTransactionOutput
	txnOutput.skyOut.Address = to.String()
	txnOutput.skyOut.Coins = strconv.FormatUint(amount/1e6, 10)

	req, err := wlt.createTransaction(nil, []core.TransactionOutput{&txnOutput}, nil, nil, client, wltR)
	return req, nil
}

func (wlt RemoteWallet) createTransaction(from []core.Address, to, uxOut []core.TransactionOutput, change core.Address, client *api.Client, wltR *api.WalletResponse) (core.Transaction, error) {
	var req api.WalletCreateTransactionRequest
	if wltR.Meta.Encrypted {
		req = api.WalletCreateTransactionRequest{
			Unsigned: true,
			WalletID: wltR.Meta.Filename,
		}
	} else {
		req = api.WalletCreateTransactionRequest{
			Unsigned: true,
			WalletID: wltR.Meta.Filename,
		}
	}
	if from != nil && len(from) > 0 {
		req.Addresses = make([]string, 0)
		for _, str := range from {
			req.Addresses = append(req.Addresses, str.String())
		}
	}
	if change != nil && len(change.String()) > 0 {
		st := change.String()
		req.ChangeAddress = &st
	}
	if uxOut != nil && len(uxOut) > 0 {
		req.UxOuts = make([]string, 0)
		for _, out := range uxOut {
			req.UxOuts = append(req.UxOuts, out.GetAddress().String())
		}

	}
	req.To = make([]api.Receiver, 0)
	for _, toTxn := range to {
		coins, err := toTxn.GetCoins(Sky)
		if err != nil {
			logrus.Warn(err)
			return nil, err
		}
		strAmount := strconv.FormatFloat(float64(coins)/1e6, 'f', -1, 64)
		req.To = append(req.To, api.Receiver{
			Address: toTxn.GetAddress().String(),
			Coins:   strAmount,
		})
	}

	req.HoursSelection = api.HoursSelection{
		Mode:        "share",
		Type:        "auto",
		ShareFactor: "0.5",
	}
	response, err := client.WalletCreateTransaction(req)
	if err != nil {
		logrus.Warn("Error creating transaction request")
		return nil, err
	}
	fee, err := util.GetCoinValue(response.Transaction.Fee, CoinHour)
	if err != nil {
		return nil, err
	}
	skyTxn, err := response.Transaction.ToTransaction()
	if err != nil {
		return nil, err
	}
	unTxn, err := NewUninjectedTransaction(skyTxn, fee)
	if err != nil {
		return nil, err
	}
	return unTxn, nil

}

func (wlt RemoteWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
	logrus.Info("Transfer from remote wallet")
	client, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, client)
	wltR, err := client.Wallet(wlt.Id)

	if err != nil {
		logrus.Warn("Error getting remote wallet")
		return nil, err
	}

	req, err := wlt.createTransaction(from, to, nil, change, client, wltR)

	if err != nil {
		logrus.Warn("Error creating transaction response")
		return nil, err
	}
	return req, nil
}

func (wlt RemoteWallet) Spend(unspent, new []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
	logrus.Info("Transfer from remote wallet")
	client, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, client)
	wltR, err := client.Wallet(wlt.Id)

	if err != nil {
		logrus.Warn("Error getting remote wallet")
		return nil, err
	}
	req, err := wlt.createTransaction(nil, nil, unspent, change, client, wltR)
	if err != nil {
		logrus.Warn("Error creating transaction request")
		return nil, err
	}

	return req, nil
}

func (wlt RemoteWallet) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		return nil
	}
	defer core.GetMultiPool().Return(wlt.poolSection, c)
	password, _ := pwd("Insert password")
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil
	}
	addresses := make([]SkycoinAddress, 0)
	for _, entry := range wltR.Entries[startIndex:int(util.Min(len(wltR.Entries), int(startIndex+count)))] {
		addresses = append(addresses, walletEntryToAddress(entry, wlt.poolSection))
	}
	//Checking if all the neccesary addresses exists
	if uint32(len(wltR.Entries)) < (startIndex + count) {
		difference := (startIndex + count) - uint32(len(wltR.Entries))
		newAddrs, err := c.NewWalletAddress(wlt.Id, int(difference), password)
		if err != nil {
			return nil
		}
		for _, addr := range newAddrs {
			addresses = append(addresses, SkycoinAddress{address: addr})
		}
	}

	return NewSkycoinAddressIterator(addresses)

}

func (wlt RemoteWallet) GetCryptoAccount() core.CryptoAccount {
	return wlt
}

func (wlt RemoteWallet) GetLoadedAddresses() (core.AddressIterator, error) {
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(wlt.poolSection, c)
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil, err
	}
	addresses := make([]SkycoinAddress, 0)
	for _, entry := range wltR.Entries {
		addresses = append(addresses, walletEntryToAddress(entry, wlt.poolSection))
	}

	return NewSkycoinAddressIterator(addresses), nil
}

func walletResponseToWallet(wltR api.WalletResponse) RemoteWallet {
	wlt := RemoteWallet{}
	wlt.CoinType = string(wltR.Meta.Coin)
	wlt.Encrypted = wltR.Meta.Encrypted
	wlt.Label = wltR.Meta.Label
	wlt.Id = wltR.Meta.Filename
	return wlt
}

func walletEntryToAddress(wltE readable.WalletEntry, poolSection string) SkycoinAddress {
	return SkycoinAddress{address: wltE.Address, poolSection: poolSection}
}

type WalletDirectory struct {
	//Implements WallentEnv interface
	WalletDir  string
	wltService *SkycoinLocalWallet
}

func (wltDir *WalletDirectory) GetStorage() core.WalletStorage {
	if wltDir.wltService == nil {
		wltDir.wltService = &SkycoinLocalWallet{wltDir.WalletDir}
	}
	return wltDir.wltService
}

func (wltDir *WalletDirectory) GetWalletSet() core.WalletSet {
	if wltDir.wltService == nil {
		wltDir.wltService = &SkycoinLocalWallet{wltDir.WalletDir}
	}
	return wltDir.wltService
}

type SkycoinLocalWallet struct {
	//Implements WalletStorage and WalletSet interfaces
	walletDir string
}

func (wltSrv *SkycoinLocalWallet) ListWallets() core.WalletIterator {
	wallets := make([]core.Wallet, 0)
	entries, err := ioutil.ReadDir(wltSrv.walletDir)
	if err != nil {
		return nil
	}

	for _, e := range entries {
		if e.Mode().IsRegular() {
			name := e.Name()
			if !strings.HasSuffix(name, walletExt) {
				continue
			}

			path := filepath.Join(wltSrv.walletDir, name)
			w, err := wallet.Load(path)
			if err != nil {
				return nil
			}
			wallets = append(wallets, LocalWallet{
				Id:        name,
				Label:     w.Label(),
				Encrypted: w.IsEncrypted(),
				Type:      w.Type(),
				CoinType:  string(w.Coin()),
				WalletDir: wltSrv.walletDir,
			})
		}
	}

	return NewSkycoinWalletIterator(wallets)
}

func (wltSrv *SkycoinLocalWallet) GetWallet(id string) core.Wallet {
	path := filepath.Join(wltSrv.walletDir, id)
	w, err := wallet.Load(path)
	if err != nil {
		return nil
	}
	return LocalWallet{
		Id:        id,
		Label:     w.Label(),
		Encrypted: w.IsEncrypted(),
		Type:      w.Type(),
		CoinType:  string(w.Coin()),
		WalletDir: wltSrv.walletDir,
	}
}

func (wltSrv *SkycoinLocalWallet) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	password, _ := pwd("Insert Password")

	passwordByte := []byte(password)
	opts := wallet.Options{
		Label:    label,
		Seed:     seed,
		Encrypt:  IsEncrypted,
		Type:     WalletTypeDeterministic,
		Password: passwordByte,
	}
	wltName := wltSrv.newUnicWalletFilename()
	var wlt wallet.Wallet
	var err error
	if scanAddressesN > 0 {
		wlt, err = wallet.NewWalletScanAhead(wltName, opts, &TransactionFinder{})
		if err != nil {

			return nil, err
		}
	} else {
		wlt, err = wallet.NewWallet(wltName, opts)
		if err != nil {

			return nil, err
		}
	}

	if err := wallet.Save(wlt, wltSrv.walletDir); err != nil {
		return nil, err
	}

	return LocalWallet{
		Id:        wltName,
		Label:     wlt.Label(),
		Encrypted: wlt.IsEncrypted(),
		Type:      wlt.Type(),
		CoinType:  string(wlt.Coin()),
		WalletDir: wltSrv.walletDir,
	}, nil
}

func (wltSrv *SkycoinLocalWallet) newUnicWalletFilename() string {
	name := ""
	for {
		timestamp := time.Now().Format(WalletTimestampFormat)
		padding := hex.EncodeToString((cipher.RandByte(2)))
		name = fmt.Sprintf("%s_%s.%s", timestamp, padding, walletExt[1:])
		if w := wltSrv.GetWallet(name); w == nil {
			break
		}
	}
	return name

}

func (wltSrv *SkycoinLocalWallet) Encrypt(walletName string, password core.PasswordReader) {
	wltName := filepath.Join(wltSrv.walletDir, walletName)
	wlt, err := wallet.Load(wltName)
	if err != nil {
		return
	}

	if wlt.IsEncrypted() {
		return
	}

	pwd, _ := password("Insert Password")
	pwdBytes := []byte(pwd)

	if err := wallet.Lock(wlt, pwdBytes, "scrypt-chacha20poly1305"); err != nil {
		return
	}

	if err := wallet.Save(wlt, wltSrv.walletDir); err != nil {
		return
	}

}

func (wltSrv *SkycoinLocalWallet) Decrypt(walletName string, password core.PasswordReader) {
	wltName := filepath.Join(wltSrv.walletDir, walletName)
	wlt, err := wallet.Load(wltName)
	if err != nil {
		return
	}
	if !wlt.IsEncrypted() {
		return
	}
	pwd, err := password("Insert Password")
	pwdBytes := []byte(pwd)

	unlockedWallet, err := wallet.Unlock(wlt, pwdBytes)
	if err != nil {
		return
	}
	if err := wallet.Save(unlockedWallet, wltSrv.walletDir); err != nil {
		return
	}
	return

}

func (wltSrv *SkycoinLocalWallet) IsEncrypted(walletName string) (bool, error) {
	wltName := filepath.Join(wltSrv.walletDir, walletName)

	wlt, err := wallet.Load(wltName)
	if err != nil {
		return false, err
	}
	return wlt.IsEncrypted(), nil
}

type TransactionFinder struct {
}

func (tf *TransactionFinder) AddressesActivity(addresses []cipher.Address) ([]bool, error) {
	addrs := make([]string, 0)
	for _, addr := range addresses {
		addrs = append(addrs, addr.String())
	}
	answ := make([]bool, len(addrs))
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, c)

	for i := 0; i < len(addrs); i++ {
		txn, err := c.Transactions([]string{addrs[i]})
		if err != nil {
			return nil, err
		}
		if len(txn) != 0 {
			answ[i] = true
		}
	}
	return answ, nil
}

type LocalWallet struct {
	Id        string
	Label     string
	CoinType  string
	Encrypted bool
	Type      string
	WalletDir string
}

func (wlt LocalWallet) Sign(encodedTxn core.Transaction, source string, pwd core.PasswordReader, index []int) (core.Transaction, error) {
	return nil, nil
}

func (wlt LocalWallet) Inject(txn string) error {
	return nil
}

func (wlt LocalWallet) GetId() string {
	return wlt.Id
}
func (wlt LocalWallet) GetLabel() string {
	return wlt.Label
}
func (wlt LocalWallet) SetLabel(wltName string) {
	wltFile, err := wallet.Load(filepath.Join(wlt.WalletDir, wlt.GetId()))
	if err != nil {
		return
	}
	wltFile.SetLabel(wltName)

	err = wallet.Save(wltFile, wlt.WalletDir)
	if err != nil {
		return
	}
	wlt.Label = wltName

}
func (wlt LocalWallet) Transfer(to core.Address, amount uint64, options core.KeyValueStorage) (core.Transaction, error) {
	strAmount := strconv.FormatFloat(float64(amount/1e6), 'f', -1, 64)
	bl, err := wlt.GetBalance(Sky)
	if err != nil {
		return nil, err
	}
	if bl < amount {
		return nil, errors.New("Don't have enough sky to make the transaction")
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	addresses := make([]string, 0)
	iterAddr, err := wlt.GetLoadedAddresses()
	if err != nil {
		return nil, err
	}
	for iterAddr.Next() {
		addresses = append(addresses, iterAddr.Value().String())
	}
	dst := api.Receiver{
		Address: to.String(),
		Coins:   strAmount,
	}
	hS := api.HoursSelection{
		Mode:        "share",
		Type:        "auto",
		ShareFactor: "0.5",
	}
	rTxn := api.CreateTransactionRequest{
		IgnoreUnconfirmed: false,
		Addresses:         addresses,
		To:                []api.Receiver{dst},
		HoursSelection:    hS,
	}
	txnResponse, err := c.CreateTransaction(rTxn)

	if err != nil {
		return nil, err
	}
	fee, err := util.GetCoinValue(txnResponse.Transaction.Fee, CoinHour)
	if err != nil {
		return nil, err
	}
	skyTxn, err := txnResponse.Transaction.ToTransaction()

	if err != nil {
		return nil, err
	}
	txn, err := NewUninjectedTransaction(skyTxn, fee)
	if err != nil {
		return nil, err
	}
	return txn, nil

}
func (wlt LocalWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
	clt, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, clt)
	addrs := make([]string, 0)
	for _, addr := range from {
		addrs = append(addrs, addr.String())
	}
	chAddr := ""
	if change != nil {
		chAddr = change.String()
	}

	obj := options.GetValue("CoinHoursMode")
	coinHoursMode, ok := obj.(string)
	if !ok {
		return nil, errors.New("Invalid options")
	}
	obj = options.GetValue("BurnFactor")
	burnFactor, ok := obj.(string)
	if !ok {
		return nil, errors.New("Invalid options")
	}
	coinHoursSelection := api.HoursSelection{
		Mode: "manual",
	}
	if coinHoursMode == "auto" {
		coinHoursSelection.Mode = "auto"
		coinHoursSelection.Type = "share"
		coinHoursSelection.ShareFactor = burnFactor
	}

	destination := make([]api.Receiver, 0)
	for _, out := range to {
		skyV, err := out.GetCoins(Sky)
		if err != nil {
			return nil, err
		}
		quotient, err := util.AltcoinQuotient(Sky)
		if err != nil {
			return nil, err
		}
		strAmount := strconv.FormatFloat(float64(skyV/quotient), 'f', -1, 64)
		recv := api.Receiver{}
		recv.Address = out.GetAddress().String()
		recv.Coins = strAmount
		if coinHoursSelection.Mode == "manual" {
			chV, err := out.GetCoins(CoinHour)
			if err != nil {
				return nil, err
			}
			quotient, err = util.AltcoinQuotient(CoinHour)
			if err != nil {
				return nil, err
			}
			recv.Hours = strconv.FormatFloat(float64(chV/quotient), 'f', -1, 64)
		}
		destination = append(destination, recv)
	}

	req := api.CreateTransactionRequest{
		Addresses:         addrs,
		ChangeAddress:     &chAddr,
		HoursSelection:    coinHoursSelection,
		IgnoreUnconfirmed: false,
		To:                destination,
	}

	txnResponse, err := clt.CreateTransaction(req)
	if err != nil {
		return nil, err
	}
	txn, err := txnResponse.Transaction.ToTransaction()
	if err != nil {
		return nil, err
	}

	//skyWlt, err := wallet.Load(filepath.Join(wlt.WalletDir, wlt.Id))
	//if err != nil {
	//	return nil, err
	//}
	//pass, err := password("Insert your password")
	//if err != nil {
	//	return nil, err
	//}
	//
	//if skyWlt.IsEncrypted() {
	//	skyWlt, err = wallet.Unlock(skyWlt, []byte(pass))
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	uxouts := make([]coin.UxOut, 0)
	for _, in := range txn.In {
		ux, err := clt.UxOut(in.String())
		if err != nil {
			return nil, err
		}
		addr, err := cipher.DecodeBase58Address(ux.OwnerAddress)
		if err != nil {
			return nil, err
		}
		srctxn, err := cipher.SHA256FromHex(ux.SrcTx)
		if err != nil {
			return nil, err
		}
		uxouts = append(uxouts, coin.UxOut{
			Head: coin.UxHead{
				BkSeq: ux.SrcBkSeq,
				Time:  ux.Time,
			},
			Body: coin.UxBody{
				Address:        addr,
				Coins:          ux.Coins,
				Hours:          ux.Hours,
				SrcTransaction: srctxn,
			},
		})
	}

	return nil, nil

}
func (wlt LocalWallet) Spend(unspent, new []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
	clt, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return nil, err
	}
	defer core.GetMultiPool().Return(PoolSection, clt)
	uxOuts := make([]string, 0)
	for _, uxOut := range unspent {
		uxOuts = append(uxOuts, uxOut.GetId())
	}
	chAddr := ""
	if change != nil {
		chAddr = change.String()
	}

	obj := options.GetValue("CoinHoursMode")
	coinHoursMode, ok := obj.(string)
	if !ok {
		return nil, errors.New("Invalid options")
	}
	obj = options.GetValue("BurnFactor")
	burnFactor, ok := obj.(string)
	if !ok {
		return nil, errors.New("Invalid options")
	}
	coinHoursSelection := api.HoursSelection{
		Mode: "manual",
	}
	if coinHoursMode == "auto" {
		coinHoursSelection.Mode = "auto"
		coinHoursSelection.Type = "share"
		coinHoursSelection.ShareFactor = burnFactor
	}

	destination := make([]api.Receiver, 0)
	for _, out := range new {
		skyV, err := out.GetCoins(Sky)
		if err != nil {
			return nil, err
		}
		quotient, err := util.AltcoinQuotient(Sky)
		if err != nil {
			return nil, err
		}
		strAmount := strconv.FormatFloat(float64(skyV/quotient), 'f', -1, 64)
		recv := api.Receiver{}
		recv.Address = out.GetAddress().String()
		recv.Coins = strAmount
		if coinHoursSelection.Mode == "manual" {
			chV, err := out.GetCoins(CoinHour)
			if err != nil {
				return nil, err
			}
			quotient, err = util.AltcoinQuotient(CoinHour)
			if err != nil {
				return nil, err
			}
			recv.Hours = strconv.FormatFloat(float64(chV/quotient), 'f', -1, 64)
		}
		destination = append(destination, recv)
	}

	req := api.CreateTransactionRequest{
		UxOuts:            uxOuts,
		ChangeAddress:     &chAddr,
		HoursSelection:    coinHoursSelection,
		IgnoreUnconfirmed: false,
		To:                destination,
	}
	txnResponse, err := clt.CreateTransaction(req)
	if err != nil {
		return nil, err
	}
	txn, err := txnResponse.Transaction.ToTransaction()
	if err != nil {
		return nil, err
	}

	//skyWlt, err := wallet.Load(filepath.Join(wlt.WalletDir, wlt.Id))
	//if err != nil {
	//	return nil, err
	//}
	//pass, err := password("Insert your password")
	//if err != nil {
	//	return nil, err
	//}
	//
	//if skyWlt.IsEncrypted() {
	//	skyWlt, err = wallet.Unlock(skyWlt, []byte(pass))
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	uxouts := make([]coin.UxOut, 0)
	for _, in := range txn.In {
		ux, err := clt.UxOut(in.String())
		if err != nil {
			return nil, err
		}
		addr, err := cipher.DecodeBase58Address(ux.OwnerAddress)
		if err != nil {
			return nil, err
		}
		srctxn, err := cipher.SHA256FromHex(ux.SrcTx)
		if err != nil {
			return nil, err
		}
		uxouts = append(uxouts, coin.UxOut{
			Head: coin.UxHead{
				BkSeq: ux.SrcBkSeq,
				Time:  ux.Time,
			},
			Body: coin.UxBody{
				Address:        addr,
				Coins:          ux.Coins,
				Hours:          ux.Hours,
				SrcTransaction: srctxn,
			},
		})
	}

	return nil, nil
}
func (wlt LocalWallet) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {
	walletName := filepath.Join(wlt.WalletDir, wlt.Id)
	walletLoaded, err := wallet.Load(walletName)
	if err != nil {
		return nil
	}
	addressCount := walletLoaded.EntriesLen()
	if uint32(addressCount) < (startIndex + count) {
		diff := (startIndex + count) - uint32(addressCount)
		genAddressesInFile := func(w wallet.Wallet, n uint64) ([]cipher.Addresser, error) {
			return w.GenerateAddresses(n)
		}

		if walletLoaded.IsEncrypted() {
			genAddressesInFile = func(w wallet.Wallet, n uint64) ([]cipher.Addresser, error) {
				password, err := pwd("Insert Password")
				if err != nil {
					return nil, nil
				}
				passwordBytes := []byte(password)
				var addrs []cipher.Addresser
				if err := wallet.GuardUpdate(w, passwordBytes, func(wlt wallet.Wallet) error {
					var err error
					addrs, err = wlt.GenerateAddresses(n)
					return err
				}); err != nil {
					return nil, err
				}

				return addrs, nil
			}
		}
		_, err := genAddressesInFile(walletLoaded, uint64(diff))

		if err != nil {
			return nil
		}

		if err := wallet.Save(walletLoaded, wlt.WalletDir); err != nil {
			return nil
		}
	}

	walletLoaded, err = wallet.Load(walletName)
	if err != nil {
		return nil
	}

	addrs := walletLoaded.GetAddresses()[startIndex : startIndex+count]
	skyAddrs := make([]SkycoinAddress, 0)
	for _, addr := range addrs {
		skyAddrs = append(skyAddrs, SkycoinAddress{address: addr.String()})
	}
	return NewSkycoinAddressIterator(skyAddrs)

}
func (wlt LocalWallet) GetCryptoAccount() core.CryptoAccount {
	return wlt
}
func (wlt LocalWallet) GetLoadedAddresses() (core.AddressIterator, error) {
	walletName := filepath.Join(wlt.WalletDir, wlt.Id)
	walletLoaded, err := wallet.Load(walletName)
	if err != nil {
		return nil, err
	}
	addrs := make([]SkycoinAddress, 0)
	addresses := walletLoaded.GetAddresses()
	for _, addr := range addresses {
		addrs = append(addrs, SkycoinAddress{address: addr.String()})
	}

	return NewSkycoinAddressIterator(addrs), nil

}
