package util

import "strconv"

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func GetCoinValue(value string) (uint64, error) {
	coin, err := strconv.ParseFloat(value, 64)
	return uint64(coin * 1000000), err
}
