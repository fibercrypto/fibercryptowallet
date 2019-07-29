package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/readable"
)

const (
	Sky      = "Sky"
	CoinHour = "CoinHour"
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

type WalletService struct { //Implements WalletStorage and WalletSet interfaces
}

func (wltSrv *WalletService) ListWallets() core.WalletIterator {
	c := util.NewClient()
	wlts, err := c.Wallets()
	if err != nil {
		return nil
	}
	wallets := make([]Wallet, 0)
	for _, wlt := range wlts {
		wallets = append(wallets, walletResponseToWallet(wlt))
	}
	return NewSkycoinWalletIterator(wallets)
}

func (wltSrv *WalletService) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	if IsEncrypted {
		password, _ := pwd("Enter your password")
		return wltSrv.createEncryptedWallet(seed, label, password, scanAddressesN)
	} else {
		return wltSrv.createUnencryptedWallet(seed, label, scanAddressesN)
	}
}
func (wltSrv *WalletService) createEncryptedWallet(seed, label, password string, scanN int) (core.Wallet, error) {
	c := util.NewClient()
	wltR, err := c.CreateEncryptedWallet(seed, label, password, scanN)
	if err != nil {
		return nil, err
	}

	wlt := walletResponseToWallet(*wltR)
	return &wlt, nil

}

func (wltSrv *WalletService) createUnencryptedWallet(seed, label string, scanN int) (core.Wallet, error) {
	c := util.NewClient()
	wltR, err := c.CreateUnencryptedWallet(seed, label, scanN)
	if err != nil {
		return nil, err
	}
	wlt := walletResponseToWallet(*wltR)
	return &wlt, nil
}

func (wltSrv *WalletService) Encrypt(walletName string, pwd core.PasswordReader) {
	c := util.NewClient()
	password, _ := pwd("Insert password")
	_, err := c.EncryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *WalletService) Decrypt(walletName string, pwd core.PasswordReader) {
	c := util.NewClient()
	password, _ := pwd("Insert password")
	_, err := c.DecryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (wltSrv *WalletService) IsEncrypted(walletName string) (bool, error) {
	c := util.NewClient()
	wlt, err := c.Wallet(walletName)
	if err != nil {
		return false, err
	}
	return wlt.Meta.Encrypted, nil
}
func (wltSrv *WalletService) GetWallet(id string) core.Wallet {
	c := util.NewClient()
	wltR, err := c.Wallet(id)
	if err != nil {
		return nil
	}
	return walletResponseToWallet(*wltR)
}

type WalletNode struct { //Implements WallentEnv interface
	wltService *WalletService
}

func (wltEnv *WalletNode) GetStorage() core.WalletStorage {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(WalletService)
	}
	return wltEnv.wltService
}

func (wltEnv *WalletNode) GetWalletSet() core.WalletSet {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(WalletService)
	}
	return wltEnv.wltService
}

type SeedService struct{} //Implements SeedGenerator interface

func (seedService *SeedService) GenerateMnemonic(entropy int) (string, error) {
	c := util.NewClient()
	seed, err := c.NewSeed(entropy)
	if err != nil {
		return "", err
	}

	return seed, nil
}

func (seedService *SeedService) VerifyMnemonic(seed string) (bool, error) {
	c := util.NewClient()
	ok, err := c.VerifySeed(seed)
	if err != nil {
		return false, err
	}
	return ok, nil
}

type errorTickerInvalid struct{}

func (err errorTickerInvalid) Error() string {
	return "Ticker invalid"
}

type Wallet struct { //Implements Wallet and CryptoAccount interfaces
	Id        string
	Label     string
	CoinType  string
	Encrypted bool
}

func (wlt Wallet) GetLabel() string {
	return wlt.Label
}

func (wlt Wallet) SetLabel(name string) {
	c := util.NewClient()
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
	c := util.NewClient()
	password, _ := pwd("Insert password")
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil
	}
	addresses := make([]Address, 0)
	for _, entry := range wltR.Entries[startIndex:int(min(len(wltR.Entries), int(startIndex+count)))] {
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
			addresses = append(addresses, Address{addr})
		}
	}

	return NewSkycoinAddressIterator(addresses)

}

func (wlt Wallet) GetCryptoAccount() core.CryptoAccount {
	return wlt
}

func (wlt Wallet) GetLoadedAddresses() (core.AddressIterator, error) {
	c := util.NewClient()
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		return nil, err
	}
	addresses := make([]Address, 0)
	for _, entry := range wltR.Entries {
		addresses = append(addresses, walletEntryToAddress(entry))
	}

	return NewSkycoinAddressIterator(addresses), nil
}

func walletResponseToWallet(wltR api.WalletResponse) Wallet {
	wlt := Wallet{}
	wlt.CoinType = wltR.Meta.Coin
	wlt.Encrypted = wltR.Meta.Encrypted
	wlt.Label = wltR.Meta.Label
	wlt.Id = wltR.Meta.Filename
	return wlt
}

func walletEntryToAddress(wltE readable.WalletEntry) Address {
	return Address{wltE.Address}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
