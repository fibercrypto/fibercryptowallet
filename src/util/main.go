package util

import (
	"errors"

	"github.com/fibercrypto/fibercryptowallet/src/core"
	local "github.com/fibercrypto/fibercryptowallet/src/main"
)

func PubKeyFromBytes(ticker string, bytes []byte) (core.PubKey, error) {
	if plugin, isRegistered := local.LoadAltcoinManager().LookupAltcoinPlugin(ticker); isRegistered {
		return plugin.PubKeyFromBytes(bytes)
	}
	return nil, errors.New(ticker + " <Unregistered>")
}

func SecKeyFromBytes(ticker string, bytes []byte) (core.PubKey, error) {
	if plugin, isRegistered := local.LoadAltcoinManager().LookupAltcoinPlugin(ticker); isRegistered {
		return plugin.SecKeyFromBytes(bytes)
	}
	return nil, errors.New(ticker + " <Unregistered>")
}
