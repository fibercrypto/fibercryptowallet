package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/errors"
	"time"
)

// BalanceSnapshot act as account balances cache
type BalanceSnapshot struct {
	cache          map[string]uint64
	lastUpdate     int64
	updateInterval int64
}

const defaultUpdateInterval = 10000000

// NewBalanceSnapshot new balance snapshot with validity timeout
func NewBalanceSnapshot(updInterval int64) *BalanceSnapshot {
	if updInterval <= 0 {
	}
	return &BalanceSnapshot{
		cache:          make(map[string]uint64),
		lastUpdate:     0,
		updateInterval: updInterval,
	}
}

// SetCoins update balance for a given coin token
func (bs *BalanceSnapshot) SetCoins(ticker string, coins uint64) {
	bs.cache[ticker] = coins
}

// GetCoins retrieve balance for a given coin ticker
func (bs *BalanceSnapshot) GetCoins(ticker string) (uint64, error) {
	if coins, hasCoins := bs.cache[ticker]; hasCoins {
		return coins, nil
	}
	return 0, errors.ErrKeyNotFound
}

func (bs *BalanceSnapshot) IsUpdated() bool {
	return time.Now().UnixNano()-bs.lastUpdate < bs.updateInterval
}
