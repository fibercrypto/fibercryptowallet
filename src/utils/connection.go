package utils

import "github.com/skycoin/skycoin/src/api"

// NewClient returns a new client
func NewClient() *api.Client {
	addr := "http://node.skycoin.net" // example only
	return api.NewClient(addr)
}
