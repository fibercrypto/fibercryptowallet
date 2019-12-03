package core

// UID is a type that holds unique ID values, including UUIDs.
// Because we don't ONLY use UUIDs, this is an alias to string.
// Being a type captures intent and helps make sure that UIDs and names do not get conflated.
type UID string

// AltcoinMetadata describes a cryptocurrency asset token
type AltcoinMetadata struct {
	// Name of the altcoin
	Name string
	// Ticker is he short name altcoin idenifier
	Ticker string
	// Family used to group altcoins having similar implementations
	Family string
	// HasBip44 highlights whether coin deesign and plugin support BIP44 address
	HasBip44 bool
	// Bip44CoinType numeric ID identiying coin segment in BIP32 derivation paths
	Bip44CoinType int32
	// Accuracy decimal places seen in coin fractions
	Accuracy int32
}

// AltcoinPlugin is the entry point to every cryptocurrency APIs
// These shall comprise at least the following:
//
// - Wallet environments list wallets managing private keys used in transactions
// - Peer exchange (a.k.a PEX) API to interact with cryptocurrency P2P network
type AltcoinPlugin interface {
	// ListSupportedAltcoins to enumerate supported assets and related metadata
	ListSupportedAltcoins() []AltcoinMetadata
	// ListSupportedFamilies classifies similar cryptocurrencies into a family
	ListSupportedFamilies() []string
	// RegisterTo boilerplate to register this plugin against an altcoin manager and enable it
	RegisterTo(manager AltcoinManager)
	// GetName provides concise human-readable caption o identify this plugin
	GetName() string
	// GetDescription describes plugin and its features
	GetDescription() string
	// LoadWalletEnvs loads wallet environments to lookup and create wallets
	// used to create and sign transactions
	LoadWalletEnvs() []WalletEnv
	// LoadPEX instantiates proxy object to interact with nodes nodes of the P2P network
	LoadPEX(netType string) (PEX, error)
	// LoadTransactionAPI blockchain transaction API entry poiny
	LoadTransactionAPI(netType string) (BlockchainTransactionAPI, error)
	// LoadSignService sign service entry point
	LoadSignService() (BlockchainSignService, error)
}

// AltcoinManager defines the contract for altcoin repositories
type AltcoinManager interface {
	// RegisterPlugin extends manager with support for another altcoin
	RegisterPlugin(p AltcoinPlugin)
	// RegisterAltcoin should be invoked in plugin's RegisterTo so as to announce support for an altcoin
	RegisterAltcoin(info AltcoinMetadata, plugin AltcoinPlugin)
	// ListRegisteredPlugins enumerates instances of AltcoinPlugin , previously registered with RegisterAltcoin
	ListRegisteredPlugins() []AltcoinPlugin
	// LookupAltcoinPlugin to instantiate plugin implementing support for coin represented by ticker
	LookupAltcoinPlugin(ticker string) (AltcoinPlugin, bool)
	// DescribeAltcoin returns metadata for coin identified by ticker
	DescribeAltcoin(ticker string) (AltcoinMetadata, bool)
	// AttachSignService registers a signing strategy for use by wallets
	AttachSignService(TxnSigner) error
	// LookupSignService returns a reference to signer identified by ID
	LookupSignService(UID) TxnSigner
	// RemoveSignService detaches a signing strategy
	RemoveSignService(TxnSigner) error
	// EnumerateSignServices returns an object to iterate over global signing srategies
	EnumerateSignServices() TxnSignerIterator
	// SignServicesForTxn returns an object to iterate over strategies supported to sign a given transaction on behalf of a wallet
	SignServicesForTxn(Wallet, Transaction) TxnSignerIterator
}
