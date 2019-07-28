package util

import (
	"github.com/skycoin/skycoin/src/api"
	//"github.com/therecipe/qt/core"
)

//Return address of daemon node----TODO
func nodeAddress() string {
	return "http://127.0.0.1:33411"
}

//Return username of daemon node----TODO
func nodeUsername() string {
	return "Kid"
}

//Return password of daemon node-----TODO
func nodePassword() string {
	return "P@ssw0rd!"
}

func NewClient() *api.Client {
	c := api.NewClient(nodeAddress())
	//c.SetAuth(nodeUsername(), nodePassword())
	return c
}

type Wallet struct {
	Sky uint64
}
