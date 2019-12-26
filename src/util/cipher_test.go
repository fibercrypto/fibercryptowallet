package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGenericAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want GenericAddress
	}{
		{
			name: "valid_Addrs",
			args: args{addr: "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt"},
			want: GenericAddress{Address: "2kvLEyXwAYvHfJuFCkjnYNRTUfHPyWgVwKt"},
		}, {
			name: "valid_Addrs2",
			args: args{addr: "R6aHqKWSQfvpdo2fGSrq4F1RYXkBWR9HHJ"},
			want: GenericAddress{Address: "R6aHqKWSQfvpdo2fGSrq4F1RYXkBWR9HHJ"},
		}, {
			name: "invalid_Addrs",
			args: args{addr: ""},
			want: GenericAddress{Address: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGenericAddress(tt.args.addr)
			require.Equal(t, got, tt.want)
			require.Equal(t, got.Bytes(), []byte(tt.args.addr))
			require.Equal(t, got.String(), tt.args.addr)
			require.False(t, got.IsBip32())
			if tt.name == "invalid_Addrs" {
				require.True(t, got.Null())
			} else {
				require.False(t, got.Null())
			}

			require.Nil(t, got.Verify(nil))
			require.Nil(t, got.GetCryptoAccount())
			require.Equal(t, got.Checksum(), core.Checksum{})

		})
	}
}
