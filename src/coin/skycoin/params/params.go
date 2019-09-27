package params

import (
	skyparams "github.com/skycoin/skycoin/src/params"
)

type SkyFiberParams struct {
	Distribution skyparams.Distribution
}

var (
	SkycoinMainNetParams = SkyFiberParams{
		Distribution: skyparams.MainNetDistribution,
	}
)

const (
	SkycoinTicker              = "SKY"
	SkycoinName                = "Skycoin"
	SkycoinFamily              = "SkyFiber"
	SkycoinDescription         = "Skycoin is an entire cryptocurrency ecosystem aimed at eliminating mining rewards, developing energy-efficient custom hardware, speeding up transaction confirmation times, and the advancement of a more secure and private Internet"
	CoinHoursTicker            = "SKYCH"
	CoinHoursName              = "Coin Hours"
	CoinHoursFamily            = "SkyFiber"
	CoinHoursDescription       = "Coin Hours is the parallel asset used for transaction fee, for creating scarcity, and to increase transaction privacy"
	CalculatedHoursTicker      = "SKYCHC"
	CalculatedHoursName	       = "Calculated Hours"
	CalculatedHoursFamily      = "SkyFiber"
	CalculatedHoursDescription = "Calculated Hours are Coin Hours calculated considering the time since an output was created"
)
