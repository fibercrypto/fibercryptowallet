package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseDate(t *testing.T) {
	some_date := int64(1577665665)
	datetime := time.Unix(some_date, 0)

	y, m, d := datetime.Date()
	h, min, s := datetime.Hour(), datetime.Minute(), datetime.Second()

	ny, nm, nd, nh, nmin, ns := ParseDate(some_date)
	require.Equal(t, y, ny)
	require.Equal(t, int(m), nm)
	require.Equal(t, d, nd)
	require.Equal(t, h, nh)
	require.Equal(t, min, nmin)
	require.Equal(t, s, ns)
}
