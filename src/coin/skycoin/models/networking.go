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

type SkycoinNetworkIterator struct {//Implements WalletIterator interface
	current int
	networks []core.Network
}

func (it *SkycoinNetworkIterator) Value() core.Network {
	return it.networks[it.current]
}

func (it *SkycoinNetworkIterator) Next() bool {
	if it.HasNext() {
		it.current++
		return true
	}
	return false
}

func (it *SkycoinNetworkIterator) HasNext() bool {
	if (it.current + 1) >= len(it.networks) {
		return false
	}
	return true
}

func NewSkycoinWalletIterator(network []core.Network) *SkycoinNetworkIterator {
	return &SkycoinNetworkIterator{networks: network, current: -1}
}

type SkycoinRemoteNetwork struct {//Implements WalletStorage and WalletSet interfaces
	nodeAddress string
}

func (remoteNetwork *SkycoinRemoteNetwork) newClient() *api.Client {
	return api.NewClient(remoteNetwork.nodeAddress)
}

func (remoteNetwork *SkycoinRemoteNetwork) ListNetworks() core.NetworkIterator {
	c := remoteNetwork.newClient()
	wlts, err := c.Wallets()
	if err != nil {
		return nil
	}
	wallets := make([]Network, 0)
	for _, wlt := range wlts {
		nwlt := walletResponseToWallet(wlt)
		nwlt.nodeAddress = remoteNetwork.nodeAddress
		wallets = append(wallets, nwlt)
	}
	return NewSkycoinWalletIterator(wallets)
}

func (remoteNetwork *SkycoinRemoteNetwork) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	wlt := Network{}
	if IsEncrypted {
		password, _ := pwd("Enter your password")
		wltOpt := api.CreateWalletOptions{}
		wltOpt.Type = WalletTypeDeterministic
		wltOpt.Seed = seed
		wltOpt.Password = password
		wltOpt.Encrypt = true
		wltOpt.Label = label
		wltOpt.ScanN = scanAddressesN
		c := remoteNetwork.newClient()
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
		c := remoteNetwork.newClient()
		wltR, err := c.CreateWallet(wltOpt)
		if err != nil {
			return nil, err
		}
		wlt = walletResponseToWallet(*wltR)
	}

	return &wlt, nil
}

func (remoteNetwork *SkycoinRemoteNetwork) Encrypt(walletName string, pwd core.PasswordReader) {
	c := remoteNetwork.newClient()
	password, _ := pwd("Insert password")
	_, err := c.EncryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (remoteNetwork *SkycoinRemoteNetwork) Decrypt(walletName string, pwd core.PasswordReader) {
	c := remoteNetwork.newClient()
	password, _ := pwd("Insert password")
	_, err := c.DecryptWallet(walletName, password)
	if err != nil {
		return
	}
}

func (remoteNetwork *SkycoinRemoteNetwork) IsEncrypted(walletName string) (bool, error) {
	c := remoteNetwork.newClient()
	wlt, err := c.Wallet(walletName)
	if err != nil {
		return false, err
	}
	return wlt.Meta.Encrypted, nil
}
func (remoteNetwork *SkycoinRemoteNetwork) GetWallet(id string) core.Wallet {
	c := remoteNetwork.newClient()
	wltR, err := c.Wallet(id)
	if err != nil {
		return nil
	}
	nwlt := walletResponseToWallet(*wltR)
	nwlt.nodeAddress = remoteNetwork.nodeAddress
	return nwlt
}

type WalletNode struct { //Implements WallentEnv interface
	wltService  *SkycoinRemoteNetwork
	nodeAddress string
}

func (wltEnv *WalletNode) GetStorage() core.WalletStorage {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteNetwork)
		wltEnv.wltService.nodeAddress = wltEnv.nodeAddress
	}
	return wltEnv.wltService
}

func (wltEnv *WalletNode) GetWalletSet() core.WalletSet {
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteNetwork)
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

type Network struct {
	Ip          string
	Port        string
	Source      string
	Block       string
	LastSeenIn  string
	LastSeenOut string
}

func (network Network) newClient() *api.Client {
	return api.NewClient(network.nodeAddress)
}
func (network Network) GetLabel() string {
	return network.Label
}

func (network Network) SetLabel(name string) {
	c := network.newClient()
	_ = c.UpdateWallet(network.Id, name)
}

func (network Network) GetId() string {
	return network.Id
}

func (network Network) Transfer(to core.Address, amount uint64) { //------TODO
	return
}

func (network Network) SendFromAddress(from, to core.Address, amount uint64) { //------TODO
	return
}

func (network Network) Spend(unspent, new []core.TransactionOutput) { //------TODO
	return
}

func (network Network) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {
	c := network.newClient()
	password, _ := pwd("Insert password")
	wltR, err := c.Wallet(network.Id)
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
		newAddrs, err := c.NewWalletAddress(network.Id, int(difference), password)
		if err != nil {
			return nil
		}
		for _, addr := range newAddrs {
			addresses = append(addresses, SkycoinAddress{addr})
		}
	}

	return NewSkycoinAddressIterator(addresses)

}

func (network Network) GetCryptoAccount() core.CryptoAccount {
	return network
}

func (network Network) GetLoadedAddresses() (core.AddressIterator, error) {
	c := network.newClient()
	wltR, err := c.Wallet(network.Id)
	if err != nil {
		return nil, err
	}
	addresses := make([]SkycoinAddress, 0)
	for _, entry := range wltR.Entries {
		addresses = append(addresses, walletEntryToAddress(entry))
	}

	return NewSkycoinAddressIterator(addresses), nil
}

func walletResponseToWallet(wltR api.WalletResponse) Network {
	wlt := Network{}
	wlt.CoinType = string(wltR.Meta.Coin)
	wlt.Encrypted = wltR.Meta.Encrypted
	wlt.Label = wltR.Meta.Label
	wlt.Id = wltR.Meta.Filename
	return wlt
}

func walletEntryToAddress(wltE readable.WalletEntry) SkycoinAddress {
	return SkycoinAddress{wltE.Address}
}
