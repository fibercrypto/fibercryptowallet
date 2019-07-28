package util

import (
	"github.com/skycoin/skycoin/src/api"
)

//Return address of daemon node
func nodeAddress() string {
	return "http://node.skycoin.net" // example only
}

//Return username of daemon node
func nodeUsername() string {
	return "user"   // example only
}

//Return password of daemon node
func nodePassword() string {
	return "password"   // example only
}

func NewClient() *api.Client {
	c := api.NewClient(nodeAddress())
	c.SetAuth(nodeUsername(), nodePassword())
	return c
}


