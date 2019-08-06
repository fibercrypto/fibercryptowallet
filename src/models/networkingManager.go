package models

//import (
//	qtcore "github.com/therecipe/qt/core"
//)
//
//type NetworkingManager struct {
//	qtcore.QObject
//
//	_ func()                                                                   `constructor:"init"`
//	_ func() []*QNetworking                                                    `slot:"getNetworks"`
//}
//
//func (net *NetworkingManager) init() {
//	net.ConnectGetNetworks(net.getNetworks)
//
//}
//
//
//func (net *NetworkingManager) decryptWallet(id, password string) {
//	pwd := func(message string) (string, error) {
//		return password, nil
//	}
//	net.WalletEnv.GetStorage().Decrypt(id, pwd)
//}
//
//func (net *NetworkingManager) getNetworks() []*QNetworking {
//	networks := make([]*QNetworking, 0)
//	it := net.WalletEnv.GetWalletSet().ListWallets()
//	for it.Next() {
//		encrypted, err := net.WalletEnv.GetStorage().IsEncrypted(it.Value().GetId())
//		if err != nil {
//			continue
//		}
//		if encrypted {
//			qwallets = append(qwallets, fromWalletToQWallet(it.Value(), true))
//		} else {
//			qwallets = append(qwallets, fromWalletToQWallet(it.Value(), false))
//		}
//
//	}
//	return qwallets
//}
