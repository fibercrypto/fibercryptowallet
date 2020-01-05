package skycoin

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/SkycoinProject/skycoin/src/api"
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/cipher/bip39"
	"github.com/SkycoinProject/skycoin/src/coin"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/SkycoinProject/skycoin/src/visor"
	"github.com/SkycoinProject/skycoin/src/wallet"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/params"
	"github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/skytypes"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
)

var logWallet = logging.MustGetLogger("Skycoin Wallet")

const (
	Sky            = params.SkycoinTicker
	CoinHour       = params.CoinHoursTicker
	CalculatedHour = params.CalculatedHoursTicker

	walletExt             = ".wlt"
	WalletTimestampFormat = "2006_01_02"

	SignerIDLocalWallet  = "sky.local"
	SignerIDRemoteWallet = "sky.remote"
)

// SkycoinWalletIterator implements WalletIterator interface
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
	return !((it.current + 1) >= len(it.wallets))
}

func NewSkycoinWalletIterator(wallets []core.Wallet) *SkycoinWalletIterator {
	return &SkycoinWalletIterator{wallets: wallets, current: -1}
}

type SkycoinRemoteWallet struct {
	// Implements WalletStorage and WalletSet interfaces
	poolSection string
}

// ListWallets returns an iterator over wallets in the set
func (wltSrv *SkycoinRemoteWallet) ListWallets() core.WalletIterator {
	logWallet.Info("Listing wallets")
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return nil
	}
	defer ReturnSkycoinClient(c)

	logWallet.Info("GET /api/v1/wallets")
	wlts, err := c.Wallets()
	if err != nil {
		logWallet.WithError(err).Error("Couldn't GET /api/v1/wallets")
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

// CreateWallet instantiates a new wallet given account seed
func (wltSrv *SkycoinRemoteWallet) CreateWallet(label string, seed string, wltType string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	logWallet.Info("Creating wallet")
	wlt := &RemoteWallet{} // nolint megacheck False negative
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return nil, err
	}
	defer ReturnSkycoinClient(c)
	if IsEncrypted {
		pwdCtx := util.NewKeyValueMap()
		pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletSet)
		pwdCtx.SetValue(core.StrMethodName, "CreateWallet")
		pwdCtx.SetValue(core.StrWalletLabel, label)
		password, err := pwd("Enter password to encrypt wallet", pwdCtx)
		if err != nil {
			logWallet.WithError(err).Fatal("Something was wrong entering the password")
			return nil, err
		}
		wltOpt := api.CreateWalletOptions{}
		wltOpt.Type = wltType
		wltOpt.Seed = seed
		wltOpt.Password = password
		wltOpt.Encrypt = true
		wltOpt.Label = label
		wltOpt.ScanN = scanAddressesN

		logWallet.Info("POST /api/v1/wallet/create")
		wltR, err := c.CreateWallet(wltOpt)
		if err != nil {
			logWallet.WithError(err).Error("Couldn't POST /api/v1/wallet/create")
			return nil, err
		}
		wlt = walletResponseToWallet(*wltR)
	} else {
		wltOpt := api.CreateWalletOptions{}
		wltOpt.Type = wltType
		wltOpt.Seed = seed
		wltOpt.Encrypt = false
		wltOpt.Label = label
		wltOpt.ScanN = scanAddressesN

		logWallet.Info("POST /api/v1/wallet/create")
		wltR, err := c.CreateWallet(wltOpt)
		if err != nil {
			logWallet.WithError(err).Error("Couldn't POST /api/v1/wallet/create")
			return nil, err
		}
		wlt = walletResponseToWallet(*wltR)
	}
	wlt.poolSection = wltSrv.poolSection
	return wlt, nil
}

// DefaultWalletType default wallet type
func (wltSrv *SkycoinRemoteWallet) DefaultWalletType() string {
	return wallet.WalletTypeBip44
}

// SupportedWalletTypes list supported wallet type names
func (wltSrv *SkycoinRemoteWallet) SupportedWalletTypes() []string {
	return []string{
		wallet.WalletTypeDeterministic,
		wallet.WalletTypeBip44,
	}
}

func (wltSrv *SkycoinRemoteWallet) Encrypt(walletName string, pwd core.PasswordReader) {
	logWallet.Info("Encrypting remote wallet")
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return
	}
	defer ReturnSkycoinClient(c)
	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletStorage)
	pwdCtx.SetValue(core.StrMethodName, "Encrypt")
	pwdCtx.SetValue(core.StrWalletName, walletName)
	password, err := pwd("Enter password to encrypt wallet", pwdCtx)
	if err != nil {
		logWallet.WithError(err).Fatal("Something was wrong entering the password")
		return
	}
	logWallet.Info("POST /api/v1/wallet/encrypt")
	_, err = c.EncryptWallet(walletName, password)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't POST /api/v1/wallet/encrypt")
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) Decrypt(walletName string, pwd core.PasswordReader) {
	logWallet.Info("Decrypting remote wallet")
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return
	}
	defer ReturnSkycoinClient(c)
	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletStorage)
	pwdCtx.SetValue(core.StrMethodName, "Decrypt")
	pwdCtx.SetValue(core.StrWalletName, walletName)
	password, err := pwd("Enter password to decrypt wallet", pwdCtx)
	if err != nil {
		logWallet.WithError(err).Fatal("Something was wrong entering the password")
		return
	}
	logWallet.Info("POST /api/v1/wallet/decrypt")
	_, err = c.DecryptWallet(walletName, password)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't POST /api/v1/wallet/decrypt")
		return
	}
}

func (wltSrv *SkycoinRemoteWallet) IsEncrypted(walletName string) (bool, error) {
	log.Info("Checking if remote wallet is Encrypted")
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return false, err
	}
	defer ReturnSkycoinClient(c)
	logWallet.Info("GET /api/v1/wallet")
	wlt, err := c.Wallet(walletName)
	if err != nil {
		logWallet.WithError(err).WithField("id", walletName).Error("Couldn't GET /api/v1/wallet")
		return false, err
	}
	return wlt.Meta.Encrypted, nil
}

// GetWallet to lookup wallet by ID
func (wltSrv *SkycoinRemoteWallet) GetWallet(id string) core.Wallet {
	logWallet.Info("Getting remote wallet")
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return nil
	}
	defer ReturnSkycoinClient(c)
	logWallet.Info("GET /api/v1/wallet")
	wltR, err := c.Wallet(id)
	if err != nil {
		logWallet.WithError(err).WithField("id", id).Error("Couldn't GET /api/v1/wallet")
		return nil
	}
	nwlt := walletResponseToWallet(*wltR)
	nwlt.poolSection = wltSrv.poolSection
	return nwlt
}

func NewWalletNode(nodeAddress string) *WalletNode {

	pool := core.GetMultiPool()
	sections, err := pool.ListSections()
	if err != nil {
		return nil
	}
	cont := 1
	var sect string
	for {
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

	err = pool.CreateSection(sect, NewSkycoinConnectionFactory(nodeAddress))
	if err != nil {
		return nil
	}
	return &WalletNode{
		NodeAddress: nodeAddress,
		poolSection: sect,
	}

}

type WalletNode struct {
	// Implements WallentEnv interface
	wltService  *SkycoinRemoteWallet
	NodeAddress string
	poolSection string
}

func (wltEnv *WalletNode) GetStorage() core.WalletStorage {
	logWallet.Info("Getting wallet node storage")
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteWallet)
		wltEnv.wltService.poolSection = wltEnv.poolSection
	}
	return wltEnv.wltService
}

func (wltEnv *WalletNode) GetWalletSet() core.WalletSet {
	logWallet.Info("Getting wallet node set")
	if wltEnv.wltService == nil {
		wltEnv.wltService = new(SkycoinRemoteWallet)
		wltEnv.wltService.poolSection = wltEnv.poolSection
	}
	return wltEnv.wltService
}

// Implements SeedGenerator interface
type SeedService struct{}

func (seedService *SeedService) GenerateMnemonic(entropyBits int) (string, error) {
	logWallet.Info("Generating mnemonic for Seed service")
	if entropyBits != 128 && entropyBits != 256 {
		return "", errors.ErrInvalidWalletEntropy
	}

	entropy, err := bip39.NewEntropy(entropyBits)
	if err != nil {
		logWallet.WithError(err).WithField("entropyBits", entropyBits).Error("Call to bip39.NewEntropy(entropyBits) inside GenerateMnemonic failed")
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		logWallet.WithError(err).WithField("entropy", entropy).Error("Call to bip39.NewMnemonic(entropy) inside GenerateMnemonic failed")
		return "", err
	}
	return mnemonic, nil
}

func (seedService *SeedService) VerifyMnemonic(seed string) (bool, error) {
	logWallet.Info("Checking mnemonic")
	err := bip39.ValidateMnemonic(seed)
	if err != nil {
		logWallet.WithError(err).WithField("seed", seed).Error("Call to bip39.ValidateMnemonic(seed) inside VerifyMnemonic failed")
		return false, err
	}
	return true, nil
}

type errorTickerInvalid struct {
	tickerUsed string
}

func (err errorTickerInvalid) Error() string {
	return err.tickerUsed + " is an invalid ticker. Use " + Sky + " or " + CoinHour
}

// Implements Wallet, TxnSigner and CryptoAccount interfaces
type RemoteWallet struct {
	Id          string
	Label       string
	CoinType    string
	Type        string
	Encrypted   bool
	poolSection string
	signers     map[core.UID]core.TxnSigner
}

func (wlt *RemoteWallet) GetSkycoinWalletType() string {
	return wlt.Type
}

func (wlt *RemoteWallet) Sign(txn core.Transaction, signer core.TxnSigner, pwd core.PasswordReader, index []string) (signedTxn core.Transaction, err error) {
	logWallet.Info("Signing using remote wallet")
	if signer == nil {
		signer = wlt
	}
	signedTxn, err = signer.SignTransaction(txn, pwd, index)
	return
}

func (wlt *RemoteWallet) signSkycoinTxn(txn core.Transaction, pwd core.PasswordReader, index []int) (core.Transaction, error) {
	client, err := NewSkycoinApiClient(PoolSection)
	var password string = ""
	if err != nil {
		logWallet.WithError(err).Warn(err)
		return nil, err
	}
	defer ReturnSkycoinClient(client)
	skyTxn, isSkyTxn := txn.(skytypes.SkycoinTxn)
	if !isSkyTxn {
		logWallet.WithError(err).Warn(err)
		return nil, errors.ErrInvalidTxn
	}
	if wlt.Encrypted {
		pwdCtx := util.NewKeyValueMap()
		pwdCtx.SetValue(core.StrTypeName, core.TypeNameWallet)
		pwdCtx.SetValue(core.StrMethodName, "Sign")
		pwdCtx.SetValue(core.StrWalletName, wlt.Id)
		pwdCtx.SetValue(core.StrWalletLabel, wlt.Label)
		password, err = pwd("Enter password", pwdCtx)
		if err != nil {
			logWallet.WithError(err).Warn("Error getting password")
			return nil, err
		}
	}

	txnBytes, err := skyTxn.EncodeSkycoinTransaction()
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't get Transaction Encoded")
		return nil, err
	}
	walletSignTxn := api.WalletSignTransactionRequest{
		EncodedTransaction: hex.EncodeToString(txnBytes),
		WalletID:           wlt.Id,
		Password:           password,
		SignIndexes:        index,
	}
	txnResponse, err := client.WalletSignTransaction(walletSignTxn)
	if err != nil {
		logWallet.WithError(err).Warn("Error signing transaction")
		return nil, err
	}
	cTxn := NewSkycoinCreatedTransaction(txnResponse.Transaction)
	return cTxn, nil
}

func (wlt *RemoteWallet) GetLabel() string {
	logWallet.Info("Getting label of remote wallet")
	return wlt.Label
}

func (wlt *RemoteWallet) SetLabel(name string) {
	logWallet.Info("Setting label to remote wallet")
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return
	}
	defer ReturnSkycoinClient(c)
	logWallet.Info("POST /api/v1/wallet/update")
	err = c.UpdateWallet(wlt.Id, name)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't POST /api/v1/wallet/update")
		return
	}
}

func (wlt *RemoteWallet) GetId() string {
	logWallet.Info("Getting Id of remote wallet")
	return wlt.Id
}

func (wlt *RemoteWallet) Transfer(destination core.TransactionOutput, options core.KeyValueStore) (core.Transaction, error) {
	logWallet.Info("Transfer from remote wallet")
	amount, err := destination.GetCoins(SkycoinTicker)
	if err != nil {
		logWallet.WithError(err).Warnf("Couldn't retrieve %s to transfer", params.SkycoinTicker)
		return nil, err
	}
	to := destination.GetAddress()

	var txnOutput SkycoinTransactionOutput
	txnOutput.skyOut.Address = to.String()
	quot, err := util.AltcoinQuotient(params.SkycoinTicker)
	if err != nil {
		logWallet.WithError(err).Warnf("Couldn't get quotient for %s", params.SkycoinTicker)
		return nil, err
	}
	txnOutput.skyOut.Coins = util.FormatCoins(amount, quot)
	createTxnFunc := func(txnR *api.CreateTransactionRequest) (core.Transaction, error) {
		logWallet.Info("Creating transaction for remote wallet")
		var req api.WalletCreateTransactionRequest
		req.Unsigned = true
		req.WalletID = wlt.Id
		req.CreateTransactionRequest = *txnR
		client, err := NewSkycoinApiClient(wlt.poolSection)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't load api client")
			return nil, err
		}
		defer ReturnSkycoinClient(client)

		txnResponse, err := client.WalletCreateTransaction(req)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't create wallet remote transaction")
			return nil, err
		}

		return fromTxnResponse(txnResponse), nil
	}

	return createTransaction(nil, []core.TransactionOutput{&txnOutput}, nil, nil, options, createTxnFunc)
}

type createTxn func(*api.CreateTransactionRequest) (core.Transaction, error)

func createTransaction(from []core.Address, to, uxOut []core.TransactionOutput, change core.Address, options core.KeyValueStore, createTxnFunc createTxn) (core.Transaction, error) {
	logWallet.Info("Creating transaction...")
	var req api.CreateTransactionRequest
	req.IgnoreUnconfirmed = false

	if from != nil {
		addrs := make([]string, 0)
		for _, addr := range from {
			addrs = append(addrs, addr.String())
		}
		req.Addresses = addrs
	}

	if uxOut != nil {
		uxOuts := make([]string, 0)
		for _, out := range uxOut {
			uxOuts = append(uxOuts, out.GetId())
		}
		req.UxOuts = uxOuts
	}

	obj := options.GetValue("CoinHoursSelectionType")
	coinHoursType, ok := obj.(string)
	if !ok {
		logWallet.WithError(nil).Warn("Couldn't get CoinHoursSelectionType")
		return nil, errors.ErrInvalidOptions
	}
	obj = options.GetValue("BurnFactor")

	burnFactor, ok := obj.(string)
	if !ok {
		logWallet.WithError(nil).Warn("Couldn't get BurnFactor")
		return nil, errors.ErrInvalidOptions
	}
	coinHoursSelection := api.HoursSelection{
		Type: "manual",
	}
	if coinHoursType == "auto" {
		coinHoursSelection.Type = "auto"
		coinHoursSelection.Mode = "share"
		coinHoursSelection.ShareFactor = burnFactor
	}
	req.HoursSelection = coinHoursSelection

	destination := make([]api.Receiver, 0)
	for _, out := range to {
		skyV, err := out.GetCoins(Sky)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't get Skycoin's")
			return nil, err
		}
		quotient, err := util.AltcoinQuotient(Sky)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't get Skycoin's quotient")
			return nil, err
		}
		strAmount := util.FormatCoins(skyV, quotient)
		recv := api.Receiver{}
		recv.Address = out.GetAddress().String()
		recv.Coins = strAmount
		if coinHoursSelection.Type == "manual" {
			chV, err := out.GetCoins(CoinHour)
			if err != nil {
				logWallet.WithError(err).Warn("Couldn't get CoinHours")
				return nil, err
			}
			quotient, err = util.AltcoinQuotient(CoinHour)
			if err != nil {
				logWallet.WithError(err).Warn("Couldn't get CoinHours quotient")
				return nil, err
			}

			recv.Hours = util.FormatCoins(chV, quotient)
		}
		destination = append(destination, recv)
	}
	req.To = destination

	if change != nil {
		ch := change.String()
		if ch != "" {
			req.ChangeAddress = &ch
		}
	}

	return createTxnFunc(&req)

}

func (wlt *RemoteWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStore) (core.Transaction, error) {
	logWallet.Info("Sending from address of remote wallets")
	createTxnFunc := func(txnR *api.CreateTransactionRequest) (core.Transaction, error) {
		logWallet.Info("Creating transaction for remote wallet")
		var req api.WalletCreateTransactionRequest
		req.Unsigned = true
		req.WalletID = wlt.Id
		req.CreateTransactionRequest = *txnR
		client, err := NewSkycoinApiClient(wlt.poolSection)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't get api client")
			return nil, err
		}
		defer ReturnSkycoinClient(client)

		txnResponse, err := client.WalletCreateTransaction(req)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't create wallet remote transaction")
			return nil, err
		}

		return fromTxnResponse(txnResponse), nil
	}

	return createTransaction(from, to, nil, change, options, createTxnFunc)
}

func (wlt *RemoteWallet) Spend(unspent, new []core.TransactionOutput, change core.Address, options core.KeyValueStore) (core.Transaction, error) {
	createTxnFunc := func(txnR *api.CreateTransactionRequest) (core.Transaction, error) {
		logWallet.Info("Spend using remote wallet")
		var req api.WalletCreateTransactionRequest
		req.Unsigned = true
		req.WalletID = wlt.Id
		req.CreateTransactionRequest = *txnR
		client, err := NewSkycoinApiClient(wlt.poolSection)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't load api client")
			return nil, err
		}
		defer ReturnSkycoinClient(client)

		txnResponse, err := client.WalletCreateTransaction(req)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't create wallet remote transaction")
			return nil, err
		}

		return fromTxnResponse(txnResponse), nil
	}

	return createTransaction(nil, new, unspent, change, options, createTxnFunc)
}

func (wlt *RemoteWallet) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {
	logWallet.Info("Generate new addresses in remote wallet")
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return nil
	}
	defer ReturnSkycoinClient(c)
	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWallet)
	pwdCtx.SetValue(core.StrMethodName, "GenAddresses")
	pwdCtx.SetValue(core.StrWalletName, wlt.Id)
	pwdCtx.SetValue(core.StrWalletLabel, wlt.Label)
	password, err := pwd("Enter password", pwdCtx)
	if err != nil {
		logWallet.WithError(err).Fatal("Something was wrong entering the password")
		return nil
	}
	logWallet.Info("GET /api/v1/wallet")
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		logWallet.WithError(err).WithField("id", wlt.Id).Error("Couldn't GET /api/v1/wallet")
		return nil
	}
	// FIXME: Lazy iterator wrapping wallet entries instead of copying to addresses slice
	addresses := make([]core.Address, 0)
	for _, entry := range wltR.Entries[startIndex:int(util.Min(len(wltR.Entries), int(startIndex+count)))] {
		addresses = append(addresses, walletEntryToAddress(entry))
	}
	// Checking if all the necessary addresses exists
	if uint32(len(wltR.Entries)) < (startIndex + count) {
		difference := (startIndex + count) - uint32(len(wltR.Entries))
		logWallet.Info("POST /api/v1/wallet/newAddress")
		newAddrs, err := c.NewWalletAddress(wlt.Id, int(difference), password)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't POST /api/v1/wallet/newAddress")
			return nil
		}
		for _, addr := range newAddrs {
			skyAddrs, err := NewSkycoinAddress(addr)
			if err != nil {
				logWallet.WithError(err).Warningf("GenAddresses: Unable to parse address %s", skyAddrs.String())
			} else if wlt.GetSkycoinWalletType() == wallet.WalletTypeBip44 {
				skyAddrs.isBip32 = true
			}
			addresses = append(addresses, &skyAddrs)
		}
	}

	return NewSkycoinAddressIterator(addresses)

}

func (wlt *RemoteWallet) GetCryptoAccount() core.CryptoAccount {
	logWallet.Info("Getting CryptoAccount of remote wallet")
	return wlt
}

func (wlt *RemoteWallet) GetLoadedAddresses() (core.AddressIterator, error) {
	logWallet.Info("Loading addresses from remote wallet")
	c, err := NewSkycoinApiClient(wlt.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return nil, err
	}
	defer ReturnSkycoinClient(c)
	logWallet.Info("GET /api/v1/wallet")
	wltR, err := c.Wallet(wlt.Id)
	if err != nil {
		logWallet.WithError(err).WithField("id", wlt.Id).Error("Couldn't GET /api/v1/wallet")
		return nil, err
	}
	addresses := make([]core.Address, 0)
	for _, entry := range wltR.Entries {
		addresses = append(addresses, walletEntryToAddress(entry))
	}

	return NewSkycoinAddressIterator(addresses), nil
}

// ReadyForTxn determines whether this signer instance can be used by wallet to sign given transaction
func (wlt *RemoteWallet) ReadyForTxn(w core.Wallet, txn core.Transaction) (bool, error) {
	return checkTxnSupported(wlt, w, txn)
}

// SignTransaction according to Skycoin SkyFiber rules
//
// @param txn Transacion object
// @param pwdReader password prompt to decode target wallet should it be needed
// @param strIdxs may be `nil` for full signing; if set should contain indices of outputs that need to be signed
func (wlt *RemoteWallet) SignTransaction(txn core.Transaction, pwdReader core.PasswordReader, strIdxs []string) (signedTxn core.Transaction, err error) {
	var indices []int
	if strIdxs == nil {
		indices = nil
	} else {
		indices, err = getHashIndices(txn.GetInputs(), strIdxs)
		if err != nil {
			logWallet.Error("Error parsing Skycoin transaction input indices array for signing")
			return nil, err
		}
	}
	signedTxn, err = wlt.signSkycoinTxn(txn, pwdReader, indices)
	return
}

func getHashIndices(ins []core.TransactionInput, strIdxs []string) (indices []int, err error) {
	cache := make(map[string]int, len(ins))
	indices = make([]int, len(strIdxs))
	scanIdx := 0
	for i, strIdx := range strIdxs {
		if strIdx[0] == '#' {
			// Parse index
			index, err := strconv.Atoi(strIdx[1:])
			if err != nil {
				return nil, errors.ErrIntegerInputsRequired
			}
			indices[i] = index
		} else if index, isCached := cache[strIdx]; isCached {
			// Found in previous scan
			indices[i] = index
		} else {
			logWallet.Infof("Scanning inputs array looking for %s", strIdx)
			// Continue scanning for UXID position in slice
			notfound := true
			for ; scanIdx < len(ins) && notfound; scanIdx++ {
				uxID := ins[scanIdx].GetId()
				logWallet.Infof("Scanning inputs array found %s", uxID)
				cache[uxID] = scanIdx
				if uxID == strIdx {
					indices[i] = scanIdx
					notfound = false
				}
			}
			if notfound {
				return nil, errors.ErrNotFound
			}
		}
	}
	return indices, nil
}

func (wlt *RemoteWallet) GetSignerUID() core.UID {
	return SignerIDRemoteWallet
}

func (wlt *RemoteWallet) GetSignerDescription() string {
	return "Remote Skycoin wallet " + wlt.Id
}

func walletResponseToWallet(wltR api.WalletResponse) *RemoteWallet {
	return &RemoteWallet{
		CoinType:  string(wltR.Meta.Coin),
		Encrypted: wltR.Meta.Encrypted,
		Label:     wltR.Meta.Label,
		Type:      wltR.Meta.Type,
		Id:        wltR.Meta.Filename,
		signers:   make(map[core.UID]core.TxnSigner),
	}
}

func walletEntryToAddress(wltE readable.WalletEntry) *SkycoinAddress {

	skyAddrs, err := NewSkycoinAddress(wltE.Address)
	if err != nil {
		logWallet.WithError(err).Error("Invalid address in wallet entry")
		return nil
	}

	return &skyAddrs
}

func NewWalletDirectory(dirPath string) *WalletDirectory {
	return &WalletDirectory{
		WalletDir: dirPath,
	}
}

type WalletDirectory struct {
	// Implements WallentEnv interface
	WalletDir  string
	wltService *SkycoinLocalWallet
}

func (wltDir *WalletDirectory) GetStorage() core.WalletStorage {
	logWallet.Info("Getting storage from wallet directory")
	if wltDir.wltService == nil {
		wltDir.wltService = &SkycoinLocalWallet{wltDir.WalletDir}
	}
	return wltDir.wltService
}

func (wltDir *WalletDirectory) GetWalletSet() core.WalletSet {
	logWallet.Info("Getting wallet set from wallet directory")
	if wltDir.wltService == nil {
		wltDir.wltService = &SkycoinLocalWallet{wltDir.WalletDir}
	}
	return wltDir.wltService
}

// Implements WalletStorage and WalletSet interfaces
type SkycoinLocalWallet struct {
	walletDir string
}

func (wltSrv *SkycoinLocalWallet) ListWallets() core.WalletIterator {
	logWallet.Info("Listing Skycoin local wallets")
	wallets := make([]core.Wallet, 0)

	logWallet.Debug("Reading wallet")
	entries, err := ioutil.ReadDir(wltSrv.walletDir)
	if err != nil {
		logWallet.WithError(err).WithField("dirname", wltSrv.walletDir).Error("Call to ioutil.ReadDir(dirname) inside ListWallets failed.")
		return nil
	}
	logWallet.Debug("Readed wallet")

	for i, e := range entries {
		logWallet.Debug("Entry " + strconv.Itoa(i) + " started")
		if e.Mode().IsRegular() {
			name := e.Name()
			if !strings.HasSuffix(name, walletExt) {
				continue
			}

			path := filepath.Join(wltSrv.walletDir, name)
			w, err := wallet.Load(path)
			if err != nil {
				logWallet.WithError(err).WithField("filename", path).Error("Call to wallet.Load(filename) inside ListWallets failed.")
				return nil
			}
			wallets = append(wallets, &LocalWallet{
				Id:        name,
				Label:     w.Label(),
				Encrypted: w.IsEncrypted(),
				Type:      w.Type(),
				CoinType:  string(w.Coin()),
				WalletDir: wltSrv.walletDir,
			})
		}
		logWallet.Debug("Entry " + strconv.Itoa(i) + " finished")
	}

	logWallet.Debug("number of wallets :=> " + strconv.Itoa(len(entries)))

	return NewSkycoinWalletIterator(wallets)
}

func (wltSrv *SkycoinLocalWallet) GetWallet(id string) core.Wallet {
	logWallet.Info("Getting Skycoin local wallet")
	path := filepath.Join(wltSrv.walletDir, id)
	w, err := wallet.Load(path)
	if err != nil {
		logWallet.WithError(err).WithField("filename", path).Error("Call to wallet.Load(filename) inside GetWallet failed.")
		return nil
	}
	return &LocalWallet{
		Id:        id,
		Label:     w.Label(),
		Encrypted: w.IsEncrypted(),
		Type:      w.Type(),
		CoinType:  string(w.Coin()),
		WalletDir: wltSrv.walletDir,
	}
}

func (wltSrv *SkycoinLocalWallet) CreateWallet(label string, seed string, wltType string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	logWallet.Info("Creating Skycoin local wallet")
	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletSet)
	pwdCtx.SetValue(core.StrMethodName, "CreateWallet")
	pwdCtx.SetValue(core.StrWalletLabel, label)
	password, err := pwd("Insert Password", pwdCtx)
	if err != nil {
		logWallet.WithError(err).Fatal("Something was wrong entering the password")
		return nil, err
	}

	passwordByte := []byte(password)

	opts := wallet.Options{
		Label:    label,
		Seed:     seed,
		Encrypt:  IsEncrypted,
		Type:     wltType,
		Password: passwordByte,
	}
	wltName := wltSrv.newUnicWalletFilename()
	var wlt wallet.Wallet

	if scanAddressesN > 0 {
		wlt, err = wallet.NewWalletScanAhead(wltName, opts, &TransactionFinder{})
		if err != nil {
			logWallet.WithError(err).WithField("wltName", wltName).Error("Call to wallet.NewWalletScanAhead(wltName, opts, &TransactionFinder{}) inside CreateWallet failed")
			return nil, err
		}
	} else {
		wlt, err = wallet.NewWallet(wltName, opts)
		if err != nil {
			logWallet.WithError(err).WithField("wltName", wltName).Error("Call to wallet.NewWallet(wltName, opts) inside CreateWallet failed")
			return nil, err
		}
	}

	if err := wallet.Save(wlt, wltSrv.walletDir); err != nil {
		logWallet.WithError(err).WithField("dir", wltSrv.walletDir).Error("Call to wallet.Save(wlt, dir) inside CreateWallet failed")
		return nil, err
	}

	return &LocalWallet{
		Id:        wltName,
		Label:     wlt.Label(),
		Encrypted: wlt.IsEncrypted(),
		Type:      wlt.Type(),
		CoinType:  string(wlt.Coin()),
		WalletDir: wltSrv.walletDir,
	}, nil
}

// DefaultWalletType default wallet type
func (wltSrv *SkycoinLocalWallet) DefaultWalletType() string {
	return wallet.WalletTypeBip44
}

// SupportedWalletTypes list supported wallet type names
func (wltSrv *SkycoinLocalWallet) SupportedWalletTypes() []string {
	return []string{
		wallet.WalletTypeDeterministic,
		wallet.WalletTypeBip44,
	}
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
	logWallet.Info("Encrypt Skycoin local wallet")
	wltName := filepath.Join(wltSrv.walletDir, walletName)
	wlt, err := wallet.Load(wltName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", wltName).Error("Call to wallet.Load(filename) inside Encrypt failed.")
		return
	}

	wltLabel := wlt.Label()
	if wlt.IsEncrypted() {
		return
	}

	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletStorage)
	pwdCtx.SetValue(core.StrMethodName, "Encrypt")
	pwdCtx.SetValue(core.StrWalletName, wltName)
	pwdCtx.SetValue(core.StrWalletLabel, wltLabel)
	pwd, err := password("Enter Password", pwdCtx)
	if err != nil {
		logWallet.WithError(err).Fatal("Something was wrong entering the password")
		return
	}
	pwdBytes := []byte(pwd)

	if err := wallet.Lock(wlt, pwdBytes, "scrypt-chacha20poly1305"); err != nil {
		logWallet.WithError(err).Error("Call to wallet.Lock() inside Encrypt failed")
		return
	}

	if err := wallet.Save(wlt, wltSrv.walletDir); err != nil {
		logWallet.WithError(err).WithField("dir", wltSrv.walletDir).Error("Call to wallet.Save(wlt, dir) inside Encrypt failed")
		return
	}

}

func (wltSrv *SkycoinLocalWallet) Decrypt(walletName string, password core.PasswordReader) {
	logWallet.Info("Decrypt Skycoin local wallet")
	wltName := filepath.Join(wltSrv.walletDir, walletName)
	wlt, err := wallet.Load(wltName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", wltName).Error("Call to wallet.Load(filename) inside Decrypt failed.")
		return
	}
	if !wlt.IsEncrypted() {
		return
	}
	wltLabel := wlt.Label()
	pwdCtx := util.NewKeyValueMap()
	pwdCtx.SetValue(core.StrTypeName, core.TypeNameWalletStorage)
	pwdCtx.SetValue(core.StrMethodName, "Decrypt")
	pwdCtx.SetValue(core.StrWalletName, wltName)
	pwdCtx.SetValue(core.StrWalletLabel, wltLabel)
	pwd, err := password("Enter Password", pwdCtx)
	if err != nil {
		logWallet.WithError(err).Fatal("Something was wrong entering the password")
		return
	}
	pwdBytes := []byte(pwd)

	unlockedWallet, err := wallet.Unlock(wlt, pwdBytes)
	if err != nil {
		logWallet.WithError(err).Error("Call to wallet.Unlock() inside Decrypt failed")
		return
	}
	if err := wallet.Save(unlockedWallet, wltSrv.walletDir); err != nil {
		logWallet.WithError(err).WithField("dir", wltSrv.walletDir).Error("Call to wallet.Save(wlt, dir) inside Decrypt failed")
		return
	}
}

func (wltSrv *SkycoinLocalWallet) IsEncrypted(walletName string) (bool, error) {
	logWallet.Info("Checking if Skycoin local wallet is encrypted")
	wltName := filepath.Join(wltSrv.walletDir, walletName)

	wlt, err := wallet.Load(wltName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", wltName).Error("Call to wallet.Load(filename) inside IsEncrypted failed.")
		return false, err
	}
	return wlt.IsEncrypted(), nil
}

type TransactionFinder struct {
}

func (tf *TransactionFinder) AddressesActivity(addresses []cipher.Address) ([]bool, error) {
	logWallet.Info("Getting Addresses activity")
	addrs := make([]string, 0)
	for _, addr := range addresses {
		addrs = append(addrs, addr.String())
	}
	answer := make([]bool, len(addrs))
	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return nil, err
	}
	defer ReturnSkycoinClient(c)

	for i := 0; i < len(addrs); i++ {
		logWallet.Info("POST /api/v1/transactions")
		txn, err := c.Transactions([]string{addrs[i]})
		if err != nil {
			logWallet.WithError(err).WithField("addr", addrs[i]).Error("Couldn't POST /api/v1/transactions")
			return nil, err
		}
		if len(txn) != 0 {
			answer[i] = true
		}
	}
	return answer, nil
}

type LocalWallet struct {
	Id        string
	Label     string
	CoinType  string
	Encrypted bool
	Type      string
	WalletDir string
}

func (wlt *LocalWallet) Sign(txn core.Transaction, signer core.TxnSigner, pwd core.PasswordReader, index []string) (signedTxn core.Transaction, err error) {
	logWallet.Info("Signing local wallet")
	if signer == nil {
		signer = wlt
	}
	signedTxn, err = signer.SignTransaction(txn, pwd, index)
	return
}

func copyTransaction(txn *coin.Transaction) *coin.Transaction {
	txnHash := txn.Hash()
	txnInnerHash := txn.HashInner()

	txn2 := *txn
	txn2.Sigs = make([]cipher.Sig, len(txn.Sigs))
	copy(txn2.Sigs, txn.Sigs)
	txn2.In = make([]cipher.SHA256, len(txn.In))
	copy(txn2.In, txn.In)
	txn2.Out = make([]coin.TransactionOutput, len(txn.Out))
	copy(txn2.Out, txn.Out)

	if txnInnerHash != txn2.HashInner() {
		logWallet.Panic("copyTransaction copy broke InnerHash")
	}
	if txnHash != txn2.Hash() {
		logWallet.Panic("copyTransaction copy broke Hash")
	}

	return &txn2
}

func (wlt *LocalWallet) signSkycoinTxn(txn core.Transaction, pwd core.PasswordReader, index []int) (core.Transaction, error) {
	var skyTxn *coin.Transaction
	var err error
	var uxouts []coin.UxOut
	var txnFee uint64
	var resultTxn core.Transaction
	walletDir := filepath.Join(wlt.WalletDir, wlt.Id)
	skyWlt, err := wallet.Load(walletDir)
	var originalInputs []api.CreatedTransactionInput
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't load api client")
		return nil, err
	}
	rTxn, isReadableTxn := txn.(skytypes.ReadableTxn)
	if isReadableTxn {
		// Readable tranasctions should not need extra API calls
		cTxn, err := rTxn.ToCreatedTransaction()
		if err != nil {
			logWallet.WithError(err).Warn("Failed to convert to readable transaction")
			return nil, err
		}
		originalInputs = cTxn.In

		if skyWlt.IsEncrypted() {
			pwdCtx := util.NewKeyValueMap()
			pwdCtx.SetValue(core.StrTypeName, core.TypeNameWallet)
			pwdCtx.SetValue(core.StrMethodName, "Sign")
			pwdCtx.SetValue(core.StrWalletName, wlt.Id)
			pwdCtx.SetValue(core.StrWalletLabel, wlt.Label)
			pass, err := pwd("Enter password", pwdCtx)
			if err != nil {
				logWallet.WithError(err).Warn("Couldn't get password")
				return nil, err
			}

			skyWlt, err = wallet.Unlock(skyWlt, []byte(pass))
			if err != nil {
				logWallet.WithError(err).Warn("Couldn't unlock local wallet")
				return nil, err
			}
		}

		skyTxn, err = cTxn.ToTransaction()
		if err != nil {
			return nil, err
		}
		uxouts = make([]coin.UxOut, len(cTxn.In))
		txnHash, err := cipher.SHA256FromHex(cTxn.TxID)
		if err != nil {
			logWallet.WithError(err).Errorf("Error parsing transaction hash %s", cTxn.TxID)
			return nil, err
		}
		tmpInt64, err := util.GetCoinValue(cTxn.Fee, params.CoinHoursTicker)
		if err != nil {
			logWallet.WithError(err).Errorf("Error parsing fee of TxID %s : %s", cTxn.TxID, cTxn.Fee)
			return nil, err
		}
		txnFee = uint64(tmpInt64)
		for i, cIn := range cTxn.In {
			tmpInt64, err = util.GetCoinValue(cIn.Coins, params.SkycoinTicker)
			if err != nil {
				logWallet.WithError(err).Errorf("Error parsing coins of uxto %s : %s", cIn.UxID, cIn.Coins)
				return nil, err
			}
			cInCoins := uint64(tmpInt64)
			tmpInt64, err = util.GetCoinValue(cIn.Hours, params.CoinHoursTicker)
			if err != nil {
				logWallet.WithError(err).Errorf("Error parsing hours of uxto %s : %s", cIn.UxID, cIn.Hours)
				return nil, err
			}
			cInHours := uint64(tmpInt64)
			cInAddr, err := cipher.DecodeBase58Address(cIn.Address)
			if err != nil {
				logWallet.WithError(err).Errorf("Error decoding base58 address for uxto %s : %s", cIn.UxID, cIn.Address)
				return nil, err
			}

			uxouts[i] = coin.UxOut{
				Head: coin.UxHead{
					Time:  cIn.Time,
					BkSeq: cIn.Block,
				},
				Body: coin.UxBody{
					SrcTransaction: txnHash,
					Address:        cInAddr,
					Coins:          cInCoins,
					Hours:          cInHours,
				},
			}
		}
	} else {
		// Raw transaction
		unTxn, ok := txn.(*SkycoinUninjectedTransaction)
		if !ok {
			logWallet.WithError(err).Warn("Couldn't load transaction un injected")
			return nil, errors.ErrInvalidTxn
		}

		// Uninjected transactions
		txnFee = unTxn.fee
		skyTxn = copyTransaction(unTxn.txn)
		clt, err := NewSkycoinApiClient(PoolSection)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't load skycoin wallet from local path")
			return nil, err
		}
		defer ReturnSkycoinClient(clt)

		if skyWlt.IsEncrypted() {

			pwdCtx := util.NewKeyValueMap()
			pwdCtx.SetValue(core.StrTypeName, core.TypeNameWallet)
			pwdCtx.SetValue(core.StrMethodName, "Sign")
			pwdCtx.SetValue(core.StrWalletName, wlt.Id)
			pass, err := pwd("Enter password", pwdCtx)
			if err != nil {
				logWallet.WithError(err).Warn("Couldn't get password")
				return nil, err
			}

			skyWlt, err = wallet.Unlock(skyWlt, []byte(pass))
			if err != nil {
				logWallet.WithError(err).Warn("Couldn't unlock local wallet")
				return nil, err
			}
		}

		uxouts = make([]coin.UxOut, 0)
		for _, in := range unTxn.txn.In {
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
	}
	logWallet.Info("Signing transaction using local transaction")
	// Transaction sigs array may not be empty
	if len(skyTxn.Sigs) == 0 {
		skyTxn.Sigs = make([]cipher.Sig, len(skyTxn.In))
	}
	signedTxn, err := wallet.SignTransaction(skyWlt, skyTxn, index, uxouts)

	if err != nil {
		logWallet.WithError(err).Warn("Couldn't sign transaction using local wallet")
		return nil, err
	}
	if isReadableTxn {
		vins := make([]visor.TransactionInput, 0)
		for i, ux := range uxouts {
			calCh, err := util.GetCoinValue(originalInputs[i].CalculatedHours, CoinHour)
			if err != nil {
				return nil, err
			}
			vin := visor.TransactionInput{
				UxOut:           ux,
				CalculatedHours: calCh,
			}
			if err != nil {
				logWallet.WithError(err).Warn("Couldn't create a transaction input")
				return nil, err
			}
			vins = append(vins, vin)
		}
		crtTxn, err := api.NewCreatedTransaction(signedTxn, vins)
		if err != nil {
			return nil, err
		}
		crtTxn.In = originalInputs
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't create an un SkycoinCreatedTransaction")
			return nil, err
		}
		resultTxn = NewSkycoinCreatedTransaction(*crtTxn)
	} else {
		resultTxn, err = NewUninjectedTransaction(signedTxn, txnFee)
		if err != nil {
			return nil, err
		}
	}
	return resultTxn, nil

}

func (wlt *LocalWallet) GetId() string {
	logWallet.Info("Getting Id if local wallet")
	return wlt.Id
}

func (wlt *LocalWallet) GetLabel() string {
	logWallet.Info("Getting label from local wallet")
	return wlt.Label
}

func (wlt *LocalWallet) SetLabel(wltName string) {
	logWallet.Info("Setting label to local wallet")
	wltFile, err := wallet.Load(filepath.Join(wlt.WalletDir, wlt.GetId()))
	if err != nil {
		logWallet.WithError(err).WithField("filename", filepath.Join(wlt.WalletDir, wlt.GetId())).Error("Call to wallet.Load(filename) inside SetLabel failed.")
		return
	}
	wltFile.SetLabel(wltName)

	err = wallet.Save(wltFile, wlt.WalletDir)
	if err != nil {
		logWallet.WithError(err).WithField("dir", wlt.WalletDir).Error("Call to wallet.Save(wlt, dir) inside SetLabel failed")
		return
	}
	wlt.Label = wltName

}

func fromTxnResponse(txnResponse *api.CreateTransactionResponse) *SkycoinCreatedTransaction {
	return NewSkycoinCreatedTransaction(txnResponse.Transaction)
}

func skyAPICreateTxn(txnReq *api.CreateTransactionRequest) (core.Transaction, error) {
	client, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't load api client")
		return nil, err
	}
	defer ReturnSkycoinClient(client)
	txnR, err := client.CreateTransaction(*txnReq)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't create transaction")
		return nil, err
	}
	return fromTxnResponse(txnR), nil
}

func (wlt *LocalWallet) Transfer(to core.TransactionOutput, options core.KeyValueStore) (core.Transaction, error) {
	logWallet.Info("Sending form local wallet")
	quotient, err := util.AltcoinQuotient(Sky)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't get skycoin quotient")
		return nil, err
	}
	amount, err := to.GetCoins(params.SkycoinTicker)
	if err != nil {
		logWallet.WithError(err).Warnf("Couldn't get ticker %s from TransactionOutput", params.SkycoinTicker)
		return nil, err
	}
	strAmount := util.FormatCoins(amount, quotient)

	var txnOutput SkycoinTransactionOutput
	txnOutput.skyOut.Address = to.GetAddress().String()
	txnOutput.skyOut.Coins = strAmount
	addresses := make([]core.Address, 0)
	iterAddr, err := wlt.GetLoadedAddresses()
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't get loaded addresses")
		return nil, err
	}
	for iterAddr.Next() {
		addresses = append(addresses, iterAddr.Value())
	}

	createTxnFunc := skyAPICreateTxn
	return createTransaction(addresses, []core.TransactionOutput{&txnOutput}, nil, nil, options, createTxnFunc)
}

func (wlt LocalWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStore) (core.Transaction, error) {
	logWallet.Info("Sending from addresses in local wallet")
	createTxnFunc := func(txnReq *api.CreateTransactionRequest) (core.Transaction, error) {
		client, err := NewSkycoinApiClient(PoolSection)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't load api client")
			return nil, err
		}
		defer ReturnSkycoinClient(client)
		txnR, err := client.CreateTransaction(*txnReq)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't create transaction")
			return nil, err
		}
		return fromTxnResponse(txnR), nil

	}

	return createTransaction(from, to, nil, change, options, createTxnFunc)

}
func (wlt LocalWallet) Spend(unspent, new []core.TransactionOutput, change core.Address, options core.KeyValueStore) (core.Transaction, error) {
	logWallet.Info("Spending from local wallet")
	createTxnFunc := func(txnReq *api.CreateTransactionRequest) (core.Transaction, error) {
		client, err := NewSkycoinApiClient(PoolSection)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't load api client")
			return nil, err
		}
		defer ReturnSkycoinClient(client)
		txnR, err := client.CreateTransaction(*txnReq)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't load api client")
			return nil, err
		}
		return fromTxnResponse(txnR), nil

	}

	return createTransaction(nil, new, unspent, change, options, createTxnFunc)
}

func (wlt *LocalWallet) GenAddresses(addrType core.AddressType, startIndex, count uint32, pwd core.PasswordReader) core.AddressIterator {

	if addrType != core.AccountAddress && addrType != core.ChangeAddress {
		logWallet.Errorf("Incorret address type %d", addrType)
		return nil
	}
	logWallet.Info("Generating addresses in local wallet")
	walletName := filepath.Join(wlt.WalletDir, wlt.Id)
	walletLoaded, err := wallet.Load(walletName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", walletName).Error("Call to wallet.Load(filename) inside GenAddresses failed.")
		return nil
	}

	if addrType == core.ChangeAddress && walletLoaded.Type() != wallet.WalletTypeBip44 {
		logWallet.Error("Change addresses may be used with Skycoin BIP44 HD wallets only")
		return nil
	}

	genAddr := func(w wallet.Wallet, n uint64) ([]cipher.Addresser, error) {
		return w.GenerateAddresses(n)
	}

	addressCount := walletLoaded.EntriesLen()

	getAddrs := func(w wallet.Wallet) []cipher.Addresser {
		return w.GetAddresses()[startIndex : startIndex+count]
	}

	if walletLoaded.Type() == wallet.WalletTypeBip44 {
		addressCount = len((*(walletLoaded.(*wallet.Bip44Wallet))).ExternalEntries)

		getAddrs = func(w wallet.Wallet) []cipher.Addresser {
			addresser := make([]cipher.Addresser, 0)
			for _, entry := range (*(walletLoaded.(*wallet.Bip44Wallet))).ExternalEntries[startIndex : startIndex+count] {
				addresser = append(addresser, entry.Address)
			}
			return addresser
		}

	}
	if addrType == core.ChangeAddress {
		addressCount = len((*(walletLoaded.(*wallet.Bip44Wallet))).ChangeEntries)

		genAddr = func(w wallet.Wallet, n uint64) ([]cipher.Addresser, error) {
			addresser := make([]cipher.Addresser, 0)
			for i := uint64(0); i < n; i++ {
				addrs, err := w.(*wallet.Bip44Wallet).GenerateChangeEntry()
				if err != nil {
					return nil, err
				}
				addresser = append(addresser, addrs.Address)
			}
			return addresser, nil
		}

		getAddrs = func(w wallet.Wallet) []cipher.Addresser {
			addresser := make([]cipher.Addresser, 0)
			for _, entry := range (*(walletLoaded.(*wallet.Bip44Wallet))).ChangeEntries[startIndex : startIndex+count] {
				addresser = append(addresser, entry.Address)
			}
			return addresser
		}
	}

	if uint32(addressCount) < (startIndex + count) {
		diff := (startIndex + count) - uint32(addressCount)
		genAddressesInFile := func(w wallet.Wallet, n uint64) ([]cipher.Addresser, error) {
			return genAddr(w, n)
		}

		if walletLoaded.IsEncrypted() {
			genAddressesInFile = func(w wallet.Wallet, n uint64) ([]cipher.Addresser, error) {
				pwdCtx := util.NewKeyValueMap()
				pwdCtx.SetValue(core.StrTypeName, core.TypeNameWallet)
				pwdCtx.SetValue(core.StrMethodName, "GenAddresses")
				pwdCtx.SetValue(core.StrWalletName, wlt.Id)
				pwdCtx.SetValue(core.StrWalletLabel, wlt.Label)
				password, err := pwd("Enter password", pwdCtx)
				if err != nil {
					logWallet.WithError(err).Error("Something was wrong entering the password")
					return nil, nil
				}
				passwordBytes := []byte(password)
				var addrs []cipher.Addresser
				if err := wallet.GuardUpdate(w, passwordBytes, func(wlt wallet.Wallet) error {
					var err error

					addrs, err = genAddr(wlt, n)

					logWallet.WithError(err).WithField("num", n).Error("Call to wlt.GenerateAddresses(num) inside wallet.GuardUpdate failed")

					return err
				}); err != nil {
					logWallet.WithError(err).Error("Call to wallet.GuardUpdate inside genAddressesInFile failed")
					return nil, err
				}

				return addrs, nil
			}
		}
		_, err := genAddressesInFile(walletLoaded, uint64(diff))

		if err != nil {
			logWallet.WithError(err).Error("Call to genAddressesInFile inside GenAddresses failed")
			return nil
		}

		if err := wallet.Save(walletLoaded, wlt.WalletDir); err != nil {
			logWallet.WithError(err).WithField("dir", wlt.WalletDir).Error("Call to wallet.Save(wlt, dir) inside GenAddresses failed")
			return nil
		}
	}

	walletLoaded, err = wallet.Load(walletName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", walletName).Error("Call to wallet.Load(filename) inside GenAddresses failed.")
		return nil
	}

	addrs := getAddrs(walletLoaded)
	skyAddrs := make([]core.Address, 0)
	for _, addr := range addrs {
		newSkyAddrs, err := NewSkycoinAddress(addr.String())
		if err != nil {
			logWallet.WithError(err).Warningf("GenAddresses: Unable to parse Skycoin address %s", addr.String())
		} else if wlt.GetSkycoinWalletType() == wallet.WalletTypeBip44 {
			newSkyAddrs.isBip32 = true
		}

		skyAddrs = append(skyAddrs, &newSkyAddrs)
	}
	return NewSkycoinAddressIterator(skyAddrs)

}

func (wlt *LocalWallet) GetCryptoAccount() core.CryptoAccount {
	logWallet.Info("Getting CryptoAccount from local wallet")
	return wlt
}

func (wlt *LocalWallet) GetLoadedAddresses() (core.AddressIterator, error) {
	logWallet.Info("Getting loaded addresses from local wallet")
	walletName := filepath.Join(wlt.WalletDir, wlt.Id)
	walletLoaded, err := wallet.Load(walletName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", walletName).Error("Call to wallet.Load(filename) inside GetLoadedAddresses failed.")
		return nil, err
	}
	addrs := make([]core.Address, 0)
	addresses := walletLoaded.GetAddresses()
	for _, addr := range addresses {
		newSkyAddrs, err := NewSkycoinAddress(addr.String())
		if err != nil {
			logWallet.WithError(err).Warningf("GetLoadedAddresses: Unable to parse Skycoin address %s", addr.String())
		} else if wlt.GetSkycoinWalletType() == wallet.WalletTypeBip44 {
			newSkyAddrs.isBip32 = true
		}
		addrs = append(addrs, &newSkyAddrs)
	}

	return NewSkycoinAddressIterator(addrs), nil

}

func (wlt *LocalWallet) GetSkycoinWalletType() string {
	return wlt.Type
}

func checkEquivalentSkycoinWallets(wlt1, wlt2 core.Wallet) (bool, error) {
	if wlt1 == wlt2 {
		return true, nil
	}
	// Must be Skycoin wallet
	skyWlt1, isSkycoinWallet := wlt1.(skytypes.SkycoinWallet)
	if !isSkycoinWallet {
		return false, nil
	}
	skyWlt2, isSkycoinWallet := wlt2.(skytypes.SkycoinWallet)
	if !isSkycoinWallet {
		return false, nil
	}
	// Must be of same Skycoin wallet type
	// FIXME: Is this enough for distribution-only wallets (i.e. xpub , collection) ?
	if skyWlt1.GetSkycoinWalletType() != skyWlt2.GetSkycoinWalletType() {
		return false, nil
	}
	// Must have a match for first address in deterministic sequence
	addrs1, err := wlt1.GetLoadedAddresses()
	if err != nil {
		return false, err
	}
	addrs2, err := wlt2.GetLoadedAddresses()
	if err != nil {
		return false, err
	}
	return addrs1.HasNext() && addrs2.HasNext() && addrs1.Value().String() == addrs2.Value().String(), nil
}

func checkTxnSupported(wlt1, wlt2 core.Wallet, txn core.Transaction) (bool, error) {
	// Wallets must match
	if isMatch, err := checkEquivalentSkycoinWallets(wlt1, wlt2); err != nil || !isMatch {
		return false, err
	}
	// Necessary and sufficient to be Skycoin transaction
	_, isSkycoinTxn := txn.(skytypes.SkycoinTxn)
	return isSkycoinTxn, nil
}

// ReadyForTxn determines whether transaction can be signed with this signer instance
func (wlt *LocalWallet) ReadyForTxn(w core.Wallet, txn core.Transaction) (bool, error) {
	return checkTxnSupported(wlt, w, txn)
}

// SignTransaction according to Skycoin SkyFiber rules
//
// @param txn Transacion object
// @param pwdReader password prompt to decode target wallet should it be needed
// @param strIdxs may be `nil` for full signing; if set should contain indices of outputs that need to be signed
func (wlt *LocalWallet) SignTransaction(txn core.Transaction, pwdReader core.PasswordReader, strIdxs []string) (signedTxn core.Transaction, err error) {
	var indices []int
	if strIdxs == nil {
		indices = nil
	} else {
		indices, err = getHashIndices(txn.GetInputs(), strIdxs)
		if err != nil {
			logWallet.Error("Error parsing Skycoin transaction input indices array for signing")
			return nil, err
		}
	}
	signedTxn, err = wlt.signSkycoinTxn(txn, pwdReader, indices)
	return
}

func (wlt *LocalWallet) GetSignerUID() core.UID {
	return SignerIDLocalWallet
}

func (wlt *LocalWallet) GetSignerDescription() string {
	return "Remote Skycoin wallet " + wlt.Id
}

// Typoe assertions
var (
	_ core.Wallet            = &LocalWallet{}
	_ core.Wallet            = &RemoteWallet{}
	_ skytypes.SkycoinWallet = &LocalWallet{}
	_ skytypes.SkycoinWallet = &RemoteWallet{}
	_ core.WalletEnv         = &WalletNode{}
	_ core.WalletEnv         = &WalletDirectory{}
	_ core.WalletSet         = &SkycoinRemoteWallet{}
	_ core.WalletStorage     = &SkycoinRemoteWallet{}
	_ core.WalletSet         = &SkycoinLocalWallet{}
	_ core.WalletStorage     = &SkycoinLocalWallet{}
)
