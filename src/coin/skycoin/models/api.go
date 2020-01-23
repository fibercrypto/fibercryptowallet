package skycoin

import "github.com/fibercrypto/fibercryptowallet/src/coin/mocks"

// "testing"

type SkycoinApiMock struct {
	mocks.SkycoinAPI
}

func (m *SkycoinApiMock) Create() (interface{}, error) {
	return m, nil
}
