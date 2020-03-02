package skycoin

import "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin/skymocks"

// "testing"

type SkycoinApiMock struct {
	skymocks.SkycoinAPI
}

func (m *SkycoinApiMock) Create() (interface{}, error) {
	return m, nil
}
