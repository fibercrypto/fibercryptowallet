package skycoin

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/sirupsen/logrus"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/wallet"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

//Implements WalletIterator interface
type SkycoinWalletIterator struct {
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

type SkycoinRemoteWallet struct { //Implements WalletStorage and WalletSet interfaces
	poolSection string
}

func (wltSrv *SkycoinRemoteWallet) newClient() (*api.Client, error) {
	pool := core.GetMultiPool()
	conn, err := WaitForPooledObject(pool, wltSrv.poolSection)

	if err != nil {
		return nil, err
	}
	c, ok := conn.(*api.Client)
	if !ok {
		return nil, errors.New(fmt.Sprintf("There is not propers client in %s pool", wltSrv.poolSection))
	}

	return c, nil
}

func (wltSrv *SkycoinRemoteWallet) ListWallets() core.WalletIterator {
	c, err := wltSrv.newClient()
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
	var c *api.Client
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
		c, err := wltSrv.newClient()

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
		c, err := wltSrv.newClient()
		wltR, err := c.CreateWallet(wltOpt)
		if err != nil {
			return nil, err
		}
		wlt = walletResponseToWallet(*wltR)
	}

	return &wlt, nil
}

func (wltSrv *SkycoinRemoteWallet) Encrypt(walletName string, pwd core.PasswordReader) {
	c, err := wltSrv.newClient()
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	password, _ := pwd("Insert password")
	_, err = c.EncryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) Decrypt(walletName string, pwd core.PasswordReader) {
	c, err := wltSrv.newClient()
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	password, _ := pwd("Insert password")
	_, err = c.DecryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) IsEncrypted(walletName string) (bool, error) {
	c, err := wltSrv.newClient()
	defer core.GetMultiPool().Return(wltSrv.poolSection, c)
	wlt, err := c.Wallet(walletName)
	if err != nil {
		return false, err
	}
	return wlt.Meta.Encrypted, nil
}
func (wltSrv *SkycoinRemoteWallet) GetWallet(id string) core.Wallet {
	c, err := wltSrv.newClient()
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

type WalletNode struct { //Implements WallentEnv interface
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

//Implements Wallet and CryptoAccount interfaces
type RemoteWallet struct {
	Id          string
	Label       string
	CoinType    string
	Encrypted   bool
	poolSection string
}

func (wlt RemoteWallet) newClient() (*api.Client, error) {
	pool := core.GetMultiPool()
	conn, err := WaitForPooledObject(pool, wlt.poolSection)
	if err != nil {
		return nil, err
	}

	c, ok := conn.(*api.Client)
	if !ok {
		return nil, errors.New(fmt.Sprintf("There is not propers client in %s pool", wlt.poolSection))
	}
	return c, nil
}
func (wlt RemoteWallet) GetLabel() string {
	return wlt.Label
}

func (wlt RemoteWallet) SetLabel(name string) {
	c, _ := wlt.newClient()

	defer core.GetMultiPool().Return(wlt.poolSection, c)
	_ = c.UpdateWallet(wlt.Id, name)
}

func (wlt RemoteWallet) GetId() string {
	return wlt.Id
}

func (wlt RemoteWallet) Transfer(to core.Address, amount uint64, password string) error {
	client, err := wlt.newClient()
	defer core.GetMultiPool().Return(wlt.poolSection, client)
	wltR, err := client.Wallet(wlt.Id)
	if err != nil {
		return err
	}

	balance, err := wlt.GetBalance(Sky)
	if err != nil {
		return err
	}

	if balance < amount {
		return errors.New("Don't have enough sky to make the transaction")
	}

	addr := SkycoinAddress{address: to.String()}
	destination, err := addr.ToSkycoinCipherAddress()

	if err != nil {
		return errors.New("Destination address invalid")
	}

	txn := api.Receiver {
		Address: destination.String(),
		Coins:   strconv.FormatUint(amount, 10),
	}
	var req api.WalletCreateTransactionRequest
	if wltR.Meta.Encrypted {
		req = api.WalletCreateTransactionRequest {
			Unsigned: false,
			WalletID: wltR.Meta.Filename,
			Password: password,
		}
	} else {

		req = api.WalletCreateTransactionRequest {
			Unsigned: false,
			WalletID: wltR.Meta.Filename,
		}
	}
	req.To = []api.Receiver{txn}

	err = json.Unmarshal([]byte("{\"type\": \"auto\", \"mode\": \"share\", \"shareFactor\": \"0.5\"}"), &req.HoursSelection)

	transactionResponse, err := client.WalletCreateTransaction(req)

	txid, err := client.InjectEncodedTransaction(transactionResponse.EncodedTransaction)
	if err != nil {
		return err
	}
	logrus.Info("Transaction " + txid + " Injected")
	return nil
}

func (wlt RemoteWallet) SendFromAddress(from, to, change core.Address, amount uint64, password string) error { //------TODO
	return nil
}

func (wlt RemoteWallet) Spend(unspent, new []core.TransactionOutput, password string) error { //------TODO
	return nil
}

func (wlt RemoteWallet) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {
	c, err := wlt.newClient()
	defer core.GetMultiPool().Return(wlt.poolSection, c)
	password, _ := pwd("Insert password")
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil
	}
	addresses := make([]SkycoinAddress, 0)
	for _, entry := range wltR.Entries[startIndex:int(util.Min(len(wltR.Entries), int(startIndex+count)))] {
		addresses = append(addresses, walletEntryToAddress(entry))
	}
	//Checking if all the neccesary addresses exists
	if uint32(len(wltR.Entries)) < (startIndex + count) {
		difference := (startIndex + count) - uint32(len(wltR.Entries))
		newAddrs, err := c.NewWalletAddress(wlt.Id, int(difference), password)
		if err != nil {
			return nil
		}
		for _, addr := range newAddrs {
			addresses = append(addresses, SkycoinAddress{addr})
		}
	}

	return NewSkycoinAddressIterator(addresses)

}

func (wlt RemoteWallet) GetCryptoAccount() core.CryptoAccount {
	return wlt
}

func (wlt RemoteWallet) GetLoadedAddresses() (core.AddressIterator, error) {
	c, err := wlt.newClient()
	defer core.GetMultiPool().Return(wlt.poolSection, c)
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil, err
	}
	addresses := make([]SkycoinAddress, 0)
	for _, entry := range wltR.Entries {
		addresses = append(addresses, walletEntryToAddress(entry))
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

func walletEntryToAddress(wltE readable.WalletEntry) SkycoinAddress {
	return SkycoinAddress{wltE.Address}
}

type WalletDirectory struct { //Implements WallentEnv interface
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

type SkycoinLocalWallet struct { //Implements WalletStorage and WalletSet interfaces
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
		ScanN:    uint64(scanAddressesN),
	}
	wltName := wltSrv.newUnicWalletFilename()
	wlt, err := wallet.NewWallet(wltName, opts)
	if err != nil {
		return nil, err
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

type LocalWallet struct {
	Id        string
	Label     string
	CoinType  string
	Encrypted bool
	Type      string
	WalletDir string
}

func (wlt LocalWallet) GetId() string {
	return wlt.Id
}
func (wlt LocalWallet) GetLabel() string {
	return wlt.Label
}
func (wlt LocalWallet) SetLabel(wltName string) {
	wlt.Label = wltName
}
func (wlt LocalWallet) Transfer(to core.Address, amount uint64, password string) error {
	//addr := SkycoinAddress{address: to.String()}
	//
	//bl, err := wlt.GetBalance(Sky)
	//if err != nil {
	//	return err
	//}
	//
	//if bl < amount {
	//	return errors.New("Don't have enough sky to make the transaction")
	//}
	//shareFactor, _ := decimal.NewFromString("0.5")
	//destination, err := addr.ToSkycoinCipherAddress()
	//
	//if err != nil {
	//	return errors.New("Destination address invalid")
	//}
	//txn := coin.TransactionOutput{
	//	Address: *destination,
	//	Coins:   amount,
	//}
	//var hoursSelection transaction.HoursSelection
	//err = json.Unmarshal([]byte("{\"type\": \"auto\", \"mode\": \"share\", \"shareFactor\": \"0.5\"}"), &hoursSelection)
	//
	//p := transaction.Params{
	//	HoursSelection: hoursSelection,
	//	To: []coin.TransactionOutput{txn},
	//}
	return nil
}
func (wlt LocalWallet) SendFromAddress(from, to, change core.Address, amount uint64, password string) error { //------TODO
	return nil
}

func (wlt LocalWallet) Spend(unspent, new []core.TransactionOutput, password string) error { //------TODO
	return nil
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
		skyAddrs = append(skyAddrs, SkycoinAddress{addr.String()})
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
		addrs = append(addrs, SkycoinAddress{addr.String()})
	}

	return NewSkycoinAddressIterator(addrs), nil

}
