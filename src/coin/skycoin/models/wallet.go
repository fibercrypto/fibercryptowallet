package models

import (
	"errors"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	"github.com/skycoin/skycoin/src/readable"
)

const (
	Sky                     = skycoin.SkycoinTicker
	CoinHour                = skycoin.CoinHoursTicker
	WalletTypeDeterministic = "deterministic"

	WalletTypeCollection = "collection"

	WalletTypeBip44 = "bip44"

	WalletTypeXPub = "xpub"
)

type SkycoinWalletIterator struct { //Implements WalletIterator interface
	current int
	wallets []Wallet
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

func NewSkycoinWalletIterator(wallets []Wallet) *SkycoinWalletIterator {
	return &SkycoinWalletIterator{wallets: wallets, current: -1}
}

type SkycoinRemoteWallet struct { //Implements WalletStorage and WalletSet interfaces
	nodeAddress string
}

func (wltSrv *SkycoinRemoteWallet) newClient() *api.Client {
	return api.NewClient(wltSrv.nodeAddress)
}

func (wltSrv *SkycoinRemoteWallet) ListWallets() core.WalletIterator {
	c := wltSrv.newClient()
	wlts, err := c.Wallets()
	if err != nil {
		return nil
	}
	wallets := make([]Wallet, 0)
	for _, wlt := range wlts {
		nwlt := walletResponseToWallet(wlt)
		nwlt.nodeAddress = wltSrv.nodeAddress
		wallets = append(wallets, nwlt)
	}
	return NewSkycoinWalletIterator(wallets)
}

func (wltSrv *SkycoinRemoteWallet) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	wlt := Wallet{}
	if IsEncrypted {
		password, _ := pwd("Enter your password")
		wltOpt := api.CreateWalletOptions{}
		wltOpt.Type = WalletTypeDeterministic
		wltOpt.Seed = seed
		wltOpt.Password = password
		wltOpt.Encrypt = true
		wltOpt.Label = label
		wltOpt.ScanN = scanAddressesN
		c := wltSrv.newClient()
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
		c := wltSrv.newClient()
		wltR, err := c.CreateWallet(wltOpt)
		if err != nil {
			return nil, err
		}
		wlt = walletResponseToWallet(*wltR)
	}

	return &wlt, nil
}

func (wltSrv *SkycoinRemoteWallet) Encrypt(walletName string, pwd core.PasswordReader) {
	c := wltSrv.newClient()
	password, _ := pwd("Insert password")
	_, err := c.EncryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) Decrypt(walletName string, pwd core.PasswordReader) {
	c := wltSrv.newClient()
	password, _ := pwd("Insert password")
	_, err := c.DecryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) IsEncrypted(walletName string) (bool, error) {
	c := wltSrv.newClient()
	wlt, err := c.Wallet(walletName)
	if err != nil {
		return false, err
	}
	return wlt.Meta.Encrypted, nil
}
func (wltSrv *SkycoinRemoteWallet) GetWallet(id string) core.Wallet {
	c := wltSrv.newClient()
	wltR, err := c.Wallet(id)
	if err != nil {
		return nil
	}
	nwlt := walletResponseToWallet(*wltR)
	nwlt.nodeAddress = wltSrv.nodeAddress
	return nwlt
}

type WalletNode struct { //Implements WallentEnv interface
	wltService  *SkycoinRemoteWallet
	nodeAddress string
}

func (wltEnv *WalletNode) GetStorage() core.WalletStorage {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteWallet)
		wltEnv.wltService.nodeAddress = wltEnv.nodeAddress
	}
	return wltEnv.wltService
}

func (wltEnv *WalletNode) GetWalletSet() core.WalletSet {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteWallet)
		wltEnv.wltService.nodeAddress = wltEnv.nodeAddress
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

type Wallet struct { //Implements Wallet and CryptoAccount interfaces
	Id          string
	Label       string
	CoinType    string
	Encrypted   bool
	nodeAddress string
}

func (wlt Wallet) newClient() *api.Client {
	return api.NewClient(wlt.nodeAddress)
}
func (wlt Wallet) GetLabel() string {
	return wlt.Label
}

func (wlt Wallet) SetLabel(name string) {
	c := wlt.newClient()
	_ = c.UpdateWallet(wlt.Id, name)
}

func (wlt Wallet) GetId() string {
	return wlt.Id
}

func (wlt Wallet) Transfer(to core.Address, amount uint64) { //------TODO
	return
}

func (wlt Wallet) SendFromAddress(from, to core.Address, amount uint64) { //------TODO
	return
}

func (wlt Wallet) Spend(unspent, new []core.TransactionOutput) { //------TODO
	return
}

func (wlt Wallet) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {
	c := wlt.newClient()
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

func (wlt Wallet) GetCryptoAccount() core.CryptoAccount {
	return wlt
}

func (wlt Wallet) GetLoadedAddresses() (core.AddressIterator, error) {
	c := wlt.newClient()
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

func walletResponseToWallet(wltR api.WalletResponse) Wallet {
	wlt := Wallet{}
	wlt.CoinType = string(wltR.Meta.Coin)
	wlt.Encrypted = wltR.Meta.Encrypted
	wlt.Label = wltR.Meta.Label
	wlt.Id = wltR.Meta.Filename
	return wlt
}

func walletEntryToAddress(wltE readable.WalletEntry) SkycoinAddress {
	return SkycoinAddress{wltE.Address}
}
