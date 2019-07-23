package util

import (
	"github.com/skycoin/skycoin/src/api"
)

//Return address of daemon node----TODO
func nodeAddress() string {
	return "http://127.0.0.1:38391"
}

//Return username of daemon node----TODO
func nodeUsername() string {
	return "Kid"
}

//Return password of daemon node-----TODO
func nodePassword() string {
	return "P@ssw0rd!"
}

func newClient() *api.Client {
	c := api.NewClient(nodeAddress())
	//c.SetAuth(nodeUsername(), nodePassword())
	return c
}