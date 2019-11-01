package skycoin

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/errors"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/fibercrypto/FiberCryptoWallet/src/util/logging"
	"github.com/skycoin/skycoin/src/api"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/bip39"
	"github.com/skycoin/skycoin/src/coin"
	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/wallet"
)

var logWallet = logging.MustGetLogger("Skycoin Wallet")

const (
	Sky                     = params.SkycoinTicker
	CoinHour                = params.CoinHoursTicker
	CalculatedHour          = params.CalculatedHoursTicker
	WalletTypeDeterministic = "deterministic"
	WalletTypeCollection    = "collection"
	WalletTypeBip44         = "bip44"

	WalletTypeXPub        = "xpub"
	walletExt             = ".wlt"
	WalletTimestampFormat = "2006_01_02"

	SignerIDLocalWallet  = "sky.local"
	SignerIDRemoteWallet = "sky.remote"
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
	return !((it.current + 1) >= len(it.wallets))
}

func NewSkycoinWalletIterator(wallets []core.Wallet) *SkycoinWalletIterator {
	return &SkycoinWalletIterator{wallets: wallets, current: -1}
}

type SkycoinRemoteWallet struct {
	//Implements WalletStorage and WalletSet interfaces
	poolSection string
}

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

func (wltSrv *SkycoinRemoteWallet) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	logWallet.Info("Creating wallet")
	wlt := &RemoteWallet{} //nolint megacheck False negative
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return nil, err
	}
	defer ReturnSkycoinClient(c)
	if IsEncrypted {
		password, err := pwd("Enter your password")
		if err != nil {
			logWallet.WithError(err).Fatal("Something was wrong entering the password")
			return nil, err
		}
		wltOpt := api.CreateWalletOptions{}
		wltOpt.Type = WalletTypeDeterministic
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
		wltOpt.Type = WalletTypeDeterministic
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

func (wltSrv *SkycoinRemoteWallet) Encrypt(walletName string, pwd core.PasswordReader) {
	logWallet.Info("Encrypting remote wallet")
	c, err := NewSkycoinApiClient(wltSrv.poolSection)
	if err != nil {
		logWallet.WithError(err).Error("Couldn't get API client")
		return
	}
	defer ReturnSkycoinClient(c)
	password, err := pwd("Insert password")
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
	password, err := pwd("Insert password")
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
	//Implements WallentEnv interface
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

//Implements SeedGenerator interface
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

//Implements Wallet, TxnSigner and CryptoAccount interfaces
type RemoteWallet struct {
	Id          string
	Label       string
	CoinType    string
	Encrypted   bool
	poolSection string
	signers     map[core.UID]core.TxnSigner
}

func (wlt *RemoteWallet) Sign(txn core.Transaction, signerID core.UID, pwd core.PasswordReader, index []string) (signedTxn core.Transaction, err error) {
	logWallet.Info("Signing using remote wallet")
	var signer core.TxnSigner
	if signerID == wlt.GetSignerUID() {
		signer = wlt
	} else {
		var isBound bool
		if signer, isBound = wlt.signers[signerID]; !isBound {
			logWallet.Error(fmt.Sprintf("RemoteWallet '%s': Unsupported signer '%s'", wlt.Id, signerID))
			return nil, errors.ErrUnsupportedSigner
		}
	}
	signedTxn, err = signer.SignTransaction(txn, pwd, index)
	return
}

func (wlt *RemoteWallet) signSkycoinTxn(txn core.Transaction, pwd core.PasswordReader, index []int) (core.Transaction, error) {
	client, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		logWallet.WithError(err).Warn(err)
		return nil, err
	}
	defer ReturnSkycoinClient(client)
	skyTxn, isSkyTxn := txn.(skycoinTxn)
	if !isSkyTxn {
		logWallet.WithError(err).Warn(err)
		return nil, errors.ErrInvalidTxn
	}
	password, err := pwd(fmt.Sprintf("Enter password to decrypt wallet '%s'", wlt.Id))
	if err != nil {
		logWallet.WithError(err).Warn("Error getting password")
		return nil, err
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

func (wlt *RemoteWallet) Transfer(to core.Address, amount uint64, options core.KeyValueStorage) (core.Transaction, error) {
	logWallet.Info("Transfer from remote wallet")

	var txnOutput SkycoinTransactionOutput
	txnOutput.skyOut.Address = to.String()
	txnOutput.skyOut.Coins = strconv.FormatUint(amount/1e6, 10)

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

func createTransaction(from []core.Address, to, uxOut []core.TransactionOutput, change core.Address, options core.KeyValueStorage, createTxnFunc createTxn) (core.Transaction, error) {
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
			uxOuts = append(uxOuts, out.GetAddress().String())
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

func (wlt *RemoteWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
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

func (wlt *RemoteWallet) Spend(unspent, new []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
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
	password, err := pwd("Insert password")
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
	addresses := make([]core.Address, 0)
	for _, entry := range wltR.Entries[startIndex:int(util.Min(len(wltR.Entries), int(startIndex+count)))] {
		addresses = append(addresses, walletEntryToAddress(entry, wlt.poolSection))
	}
	//Checking if all the necessary addresses exists
	if uint32(len(wltR.Entries)) < (startIndex + count) {
		difference := (startIndex + count) - uint32(len(wltR.Entries))
		logWallet.Info("POST /api/v1/wallet/newAddress")
		newAddrs, err := c.NewWalletAddress(wlt.Id, int(difference), password)
		if err != nil {
			logWallet.WithError(err).Warn("Couldn't POST /api/v1/wallet/newAddress")
			return nil
		}
		for _, addr := range newAddrs {
			addresses = append(addresses, &SkycoinAddress{address: addr})
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
		addresses = append(addresses, walletEntryToAddress(entry, wlt.poolSection))
	}

	return NewSkycoinAddressIterator(addresses), nil
}

func (wlt *RemoteWallet) AttachSignService(signSrv core.TxnSigner) error {
	wlt.signers[signSrv.GetSignerUID()] = signSrv
	return nil
}

func (wlt *RemoteWallet) EnumerateSignServices() core.TxnSignerIterator {
	// TODO: Implement
	return nil
}

func (wlt *RemoteWallet) RemoveSignService(signSrv core.TxnSigner) error {
	uid := signSrv.GetSignerUID()
	if _, isBound := wlt.signers[uid]; isBound {
		delete(wlt.signers, uid)
		return nil
	}
	return errors.ErrInvalidValue
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
		Id:        wltR.Meta.Filename,
		signers:   make(map[core.UID]core.TxnSigner),
	}
}

func walletEntryToAddress(wltE readable.WalletEntry, poolSection string) *SkycoinAddress {
	return &SkycoinAddress{address: wltE.Address, poolSection: poolSection}
}

func NewWalletDirectory(dirPath string) *WalletDirectory {
	return &WalletDirectory{
		WalletDir: dirPath,
	}
}

type WalletDirectory struct {
	//Implements WallentEnv interface
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

//Implements WalletStorage and WalletSet interfaces
type SkycoinLocalWallet struct {
	walletDir string
}

func (wltSrv *SkycoinLocalWallet) ListWallets() core.WalletIterator {
	logWallet.Info("Listing Skycoin local wallets")
	wallets := make([]core.Wallet, 0)
	entries, err := ioutil.ReadDir(wltSrv.walletDir)
	if err != nil {
		logWallet.WithError(err).WithField("dirname", wltSrv.walletDir).Error("Call to ioutil.ReadDir(dirname) inside ListWallets failed.")
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
				signers:   make(map[core.UID]core.TxnSigner),
			})
		}
	}

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
		signers:   make(map[core.UID]core.TxnSigner),
	}
}

func (wltSrv *SkycoinLocalWallet) CreateWallet(label string, seed string, IsEncrypted bool, pwd core.PasswordReader, scanAddressesN int) (core.Wallet, error) {
	logWallet.Info("Creating Skycoin local wallet")
	password, err := pwd("Insert Password")

	if err != nil {
		logWallet.WithError(err).Fatal("Something was wrong entering the password")
		return nil, err
	}

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
		signers:   make(map[core.UID]core.TxnSigner),
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
	logWallet.Info("Encrypt Skycoin local wallet")
	wltName := filepath.Join(wltSrv.walletDir, walletName)
	wlt, err := wallet.Load(wltName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", wltName).Error("Call to wallet.Load(filename) inside Encrypt failed.")
		return
	}

	if wlt.IsEncrypted() {
		return
	}

	pwd, err := password("Insert Password")
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
	pwd, err := password("Insert Password")
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
	signers   map[core.UID]core.TxnSigner
}

func (wlt *LocalWallet) Sign(txn core.Transaction, signerID core.UID, pwd core.PasswordReader, index []string) (signedTxn core.Transaction, err error) {
	var signer core.TxnSigner
	logWallet.Info("Signing local wallet")
	if signerID == wlt.GetSignerUID() {
		signer = wlt
	} else {
		var isBound bool
		if signer, isBound = wlt.signers[signerID]; !isBound {
			logWallet.Error(fmt.Sprintf("RemoteWallet '%s': Unsupported signer '%s'", wlt.Id, signerID))
			return nil, errors.ErrUnsupportedSigner
		}
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

	walletDir := filepath.Join(wlt.WalletDir, wlt.Id)
	skyWlt, err := wallet.Load(walletDir)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't load api client")
		return nil, err
	}
	if rTxn, isReadableTxn := txn.(readableTxn); isReadableTxn {
		// Readable tranasctions should not need extra API calls
		cTxn, err := rTxn.ToCreatedTransaction()
		if err != nil {
			return nil, err
		}
		skyTxn, err = cTxn.ToTransaction()
		if err != nil {
			return nil, err
		}
		uxouts = make([]coin.UxOut, len(cTxn.In))
		txnHash, err := cipher.SHA256FromHex(cTxn.TxID)
		if err != nil {
			logWallet.Errorf("Error parsing transaction hash %s", cTxn.TxID)
			return nil, err
		}
		tmpInt64, err := strconv.ParseInt(cTxn.Fee, 10, 64)
		if err != nil {
			logWallet.Errorf("Error parsing fee of TxID %s : %s", cTxn.TxID, cTxn.Fee)
			return nil, err
		}
		txnFee = uint64(tmpInt64)
		for i, cIn := range cTxn.In {
			tmpInt64, err = strconv.ParseInt(cIn.Coins, 10, 64)
			if err != nil {
				logWallet.Errorf("Error parsing coins of uxto %s : %s", cIn.UxID, cIn.Coins)
				return nil, err
			}
			cInCoins := uint64(tmpInt64)
			tmpInt64, err = strconv.ParseInt(cIn.Hours, 10, 64)
			if err != nil {
				logWallet.Errorf("Error parsing hours of uxto %s : %s", cIn.UxID, cIn.Hours)
				return nil, err
			}
			cInHours := uint64(tmpInt64)
			cInAddr, err := cipher.DecodeBase58Address(cIn.Address)
			if err != nil {
				logWallet.Errorf("Error decoding base58 address for uxto %s : %s", cIn.UxID, cIn.Address)
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
			pass, err := pwd("Type your password")
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
	// FIXME: Return readable SkycoinCreatedTransaction since UX data is available
	resultTxn, err := NewUninjectedTransaction(signedTxn, txnFee)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't create an un injected transaction")
		return nil, err
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

func (wlt *LocalWallet) Transfer(to core.Address, amount uint64, options core.KeyValueStorage) (core.Transaction, error) {
	logWallet.Info("Sending form local wallet")
	quotient, err := util.AltcoinQuotient(Sky)
	if err != nil {
		logWallet.WithError(err).Warn("Couldn't get skycoin quotient")
		return nil, err
	}
	strAmount := util.FormatCoins(amount, quotient)

	var txnOutput SkycoinTransactionOutput
	txnOutput.skyOut.Address = to.String()
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

func (wlt LocalWallet) SendFromAddress(from []core.Address, to []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
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
func (wlt LocalWallet) Spend(unspent, new []core.TransactionOutput, change core.Address, options core.KeyValueStorage) (core.Transaction, error) {
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
	logWallet.Info("Generating addresses in local wallet")
	walletName := filepath.Join(wlt.WalletDir, wlt.Id)
	walletLoaded, err := wallet.Load(walletName)
	if err != nil {
		logWallet.WithError(err).WithField("filename", walletName).Error("Call to wallet.Load(filename) inside GenAddresses failed.")
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
					logWallet.WithError(err).Error("Something was wrong entering the password")
					return nil, nil
				}
				passwordBytes := []byte(password)
				var addrs []cipher.Addresser
				if err := wallet.GuardUpdate(w, passwordBytes, func(wlt wallet.Wallet) error {
					var err error
					addrs, err = wlt.GenerateAddresses(n)
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

	addrs := walletLoaded.GetAddresses()[startIndex : startIndex+count]
	skyAddrs := make([]core.Address, 0)
	for _, addr := range addrs {
		skyAddrs = append(skyAddrs, &SkycoinAddress{address: addr.String()})
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
		addrs = append(addrs, &SkycoinAddress{address: addr.String()})
	}

	return NewSkycoinAddressIterator(addrs), nil

}

func (wlt *LocalWallet) AttachSignService(signSrv core.TxnSigner) error {
	wlt.signers[signSrv.GetSignerUID()] = signSrv
	return nil
}

func (wlt *LocalWallet) EnumerateSignServices() core.TxnSignerIterator {
	// TODO: Implement
	return nil
}

func (wlt *LocalWallet) RemoveSignService(signSrv core.TxnSigner) error {
	uid := signSrv.GetSignerUID()
	if _, isBound := wlt.signers[uid]; isBound {
		delete(wlt.signers, uid)
		return nil
	}
	// FIXME: Global error object
	return errors.ErrInvalidValue
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
