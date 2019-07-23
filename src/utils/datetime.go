package utils

import "time"

// ParseDate parse timestamp to a date represented as a tuple
func ParseDate(timeStamp int64) (int, int, int, int, int, int) {
	t := time.Unix(timeStamp, 0) //Fixme: use or not UTC() for local time or for server time?
	y, _m, d := t.Date()

	return y, int(_m), d, t.Hour(), t.Minute(), t.Second()
}
