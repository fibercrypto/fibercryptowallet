package skycoin

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/coin/skycoin/params"
	skyparams "github.com/skycoin/skycoin/src/params"
)

var (
	SkycoinMainNetParams = params.SkyFiberParams{
		Distribution: skyparams.MainNetDistribution,
	}
)

const (
	SkycoinTicker              = params.SkycoinTicker
	SkycoinName                = params.SkycoinName
	SkycoinFamily              = params.SkycoinFamily
	SkycoinDescription         = params.SkycoinDescription
	CoinHoursTicker            = params.CoinHoursTicker
	CoinHoursName              = params.CoinHoursName
	CoinHoursFamily            = params.CoinHoursFamily
	CoinHoursDescription       = params.CoinHoursDescription
	CalculatedHoursTicker      = params.CalculatedHoursTicker
	CalculatedHoursName        = params.CalculatedHoursName
	CalculatedHoursFamily      = params.CalculatedHoursFamily
	CalculatedHoursDescription = params.CalculatedHoursDescription
)
