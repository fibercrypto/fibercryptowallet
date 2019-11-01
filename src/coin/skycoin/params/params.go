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

// Constparams
const (
	// SkycoinTicker Skycoin coin identifier
	SkycoinTicker = "SKY"
	// SkycoinName human readable name associated to Skycoin
	SkycoinName = "Skycoin"
	// SkycoinFamily identifies Skyfiber coins
	SkycoinFamily = "SkyFiber"
	// SkycoinDescription verbose explanaitiion of Skycoin
	SkycoinDescription = "Skycoin is an entire cryptocurrency ecosystem aimed at eliminating mining rewards, developing energy-efficient custom hardware, speeding up transaction confirmation times, and the advancement of a more secure and private Internet"
	// CoinHoursTicker internal identifier to refer to Skycoin coin hours
	CoinHoursTicker = "SKY#CH"
	// CoinHoursName is the readable name for coin hours
	CoinHoursName = "Coin Hours"
	// CoinHoursDescription verbose explanaitiion of coin hours
	CoinHoursDescription = "Coin Hours is the parallel asset used for transaction fee, for creating scarcity, and to increase transaction privacy"
	// CoinHoursTicker internal identifier to refer to accumulated coin hours
	CalculatedHoursTicker = "SKY#CHC"
	// CoinHoursName is the readable name for accumulated coin hours
	CalculatedHoursName = "Calculated Hours"
	// CalculatedHoursDescription verbose explanaitiion of accumulated coin hours
	CalculatedHoursDescription = "Calculated Hours are Coin Hours calculated considering the time since an output was created"
)
