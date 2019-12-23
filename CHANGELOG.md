# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## Unreleased

### Added

- Spend in single transaction coins owned by multiple wallets (same altcoin plugin)
- Added logger for the app and proper error handling
- Built-in support for SkyWallet hardware wallet as signer for Skycoin transactions
- Built-in support iand GUI for applying operations upon SkyWallet hardware devices
- A UI to show blockchains
- GUI for status of peer-to-peer exchange network of cryptocurrencies
- GUI to render unspent outputs
- GUI listing transactions marked as pending in memory pool
- GUI to preview transactions and confirm submission
- GUI to enter password when sending a transaction
- GUI for transaction details
- GUI to list transactions history
- GUI for transactions management
- Documentation for the QML API
- GUI to confirm wallet removal
- GUI to manage wallets
- GUI for wallet details
- GUI to enter the PIN code as requested by hardware wallet
- GUI to inform the user to resolve security issues before using the wallet
- GUI to secure the wallet
- GUI to inform that the hardware wallet has been detected and configured. It requests a name for it
- GUI that informs to the user that the desktop wallet is busy checking the hardware wallet
- GUI to recover a hardware wallet from a backup
- GUI to configure an unconfigured hardware wallet
- GUI to report errors related to the hardware wallet
- GUI to create a new wallet or load an existing one
- GUI to edit wallets
- Backend models autodetect available wallets
- Continuous integration powered by Travis CI
- FiberCrypto wallet architecture and core interfaces
- [Skycoin] Backend to manage pending transactions
- [Skycoin] Backend clases for peer-to-peer exchange networking status
- [Skycoin] Backend to manage (un)spent outputs
- [Skycoin] Backend to manage the blockchain
- [Skycoin] Backend to manage transaction history
- [Skycoin] Backend classes for transactions
- [Skycoin] Backend to manage the list of wallets
- Project licensed under the terms of GPLv3
- About Qt dialog
- MenuBar in the header of the main application window
- Generic multi-coin connection pool
- [Skycoin] Support for spend in single transaction outputs managed by multiple wallets
- Internationalization for UI messages
- Add a fictional coin ticker `CalculatedHour` to disambiguate between coin hours and calculated hours
- GUI to select transaction change address
- Wallet GUI DARK (night) theme
- Generic dialogs to set and get passwords
- GUI (simple and advanced) to put together transaction details
- Dynamic multi-platform configuration for FiberCryptoWallet (powered by QT standard QSettings)
- Different account balances rendered in user interface controls
- Custom dropdown control in GUI for combo boxes that show addresses
- QR code images are buttons
- GUI to add wallets
- GUI to generate new addresses for a wallet
- Added Address Book
