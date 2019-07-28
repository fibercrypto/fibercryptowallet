package api

import "github.com/fibercrypto/FiberCryptoWallet/src/core"

type WalletNode struct {
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
