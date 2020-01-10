package util

import (
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		min  int
	}{
		{name: "comparison1", a: 10, b: 5, min: 5},
		{name: "comparison2", a: 3, b: 8, min: 3},
		{name: "comparison3", a: 1000000, b: 5, min: 5},
		{name: "comparison4", a: 0, b: 1, min: 0},
		{name: "comparison5", a: 10, b: 10, min: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.min, Min(tt.a, tt.b))
		})
	}
}

func TestFormatUint64(t *testing.T) {
	tests := []struct {
		name   string
		value  uint64
		result string
	}{
		{name: "format1", value: uint64(12345678), result: "12,345,678"},
		{name: "format2", value: uint64(15333), result: "15,333"},
		{name: "format3", value: uint64(1234), result: "1,234"},
		{name: "format4", value: uint64(10), result: "10"},
		{name: "format5", value: uint64(4242), result: "4,242"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.result, FormatUint64(tt.value))
		})
	}
}

func TestFormatCoins(t *testing.T) {
	tests := []struct {
		name     string
		value    uint64
		quotient uint64
		result   string
	}{
		{name: "format1", value: uint64(12345678), quotient: uint64(10), result: "1,234,567.8"},
		{name: "format2", value: uint64(1533300), result: "15,333", quotient: uint64(100)},
		{name: "format3", value: uint64(1234), result: "0.1234", quotient: uint64(10000)},
		{name: "format4", value: uint64(12340), result: "1.234", quotient: uint64(10000)},
		{name: "format5", value: uint64(421142), result: "4,211.42", quotient: uint64(100)},
		{name: "format6", value: uint64(0), result: "0", quotient: uint64(100)},
		{name: "format7", value: uint64(42), result: "0", quotient: uint64(10000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.result, FormatCoins(tt.value, tt.quotient))
		})
	}
}

func TestRemoveZeros(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		result string
	}{
		{name: "remove1", value: "10", result: "1"},
		{name: "remove2", value: "15.20", result: "15.2"},
		{name: "remove3", value: "42.0042", result: "42.0042"},
		{name: "remove4", value: "42.004200", result: "42.0042"},
		{name: "remove5", value: "42", result: "42"},
		{name: "remove6", value: "00000", result: ""},
		{name: "remove6", value: "00000.01", result: "00000.01"},
		{name: "remove7", value: "00000.00", result: "00000."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.result, RemoveZeros(tt.value))
		})
	}
}

func TestStringInList(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		list   []string
		result bool
	}{
		{name: "search1", value: "10", list: []string{"a", "", "10"}, result: true},
		{name: "search2", value: "15.20", list: []string{"15.20", "15", "15.2", "15.202"}, result: true},
		{name: "search3", value: "ab", list: []string{"a15a", "bca", "ab", "mn"}, result: true},
		{name: "search4", value: "42.004200", list: []string{}, result: false},
		{name: "search5", value: "42", list: []string{"40"}, result: false},
		{name: "search6", value: "00000", list: []string{"00000", "00000"}, result: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.result, StringInList(tt.value, tt.list))
		})
	}
}

func TestGetCoinValue(t *testing.T) {
	fakeTicker := "MOCKSCOIN"
	fakeDesc := "Fake coin"
	fakeExp := 3
	fakeMeta := core.AltcoinMetadata{
		Name:          fakeDesc,
		Ticker:        fakeTicker,
		Family:        fakeTicker,
		HasBip44:      false,
		Bip44CoinType: 0,
		Accuracy:      int32(fakeExp),
	}
	mockPlugin := new(mocks.AltcoinPlugin)
	mockPlugin.On("RegisterTo", mock.Anything).Return().Run(func(args mock.Arguments) {
		manager := args.Get(0).(core.AltcoinManager)
		manager.RegisterAltcoin(fakeMeta, mockPlugin)
	})
	mockPlugin.On("GetName").Return(fakeDesc)
	RegisterAltcoin(mockPlugin)

	tests := []struct {
		name   string
		value  string
		ticker string
		valid  bool
		want   uint64
	}{
		{name: "invalidGetCoinValue1", value: "10", ticker: "MYCOIN", valid: false, want: uint64(0)},
		{name: "invalidGetCoinValue2", value: "coin", ticker: fakeTicker, valid: false, want: uint64(0)},
		{name: "validGetCoinValue1", value: "10", ticker: fakeTicker, valid: true, want: uint64(10000)},
		{name: "validGetCoinValue2", value: "12.123456", ticker: fakeTicker, valid: true, want: uint64(12123)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				val, err := GetCoinValue(tt.value, tt.ticker)
				require.Nil(t, err)
				require.Equal(t, tt.want, val)
			} else {
				_, err := GetCoinValue(tt.value, tt.ticker)
				require.NotNil(t, err)
			}
		})
	}
}

func TestAddressFromString(t *testing.T) {
	fakeTicker := "MOCKSCOIN"
	fakeDesc := "Fake coin"
	fakeExp := 3
	fakeMeta := core.AltcoinMetadata{
		Name:          fakeDesc,
		Ticker:        fakeTicker,
		Family:        fakeTicker,
		HasBip44:      false,
		Bip44CoinType: 0,
		Accuracy:      int32(fakeExp),
	}
	mockPlugin := new(mocks.AltcoinPlugin)
	mockPlugin.On("RegisterTo", mock.Anything).Return().Run(func(args mock.Arguments) {
		manager := args.Get(0).(core.AltcoinManager)
		manager.RegisterAltcoin(fakeMeta, mockPlugin)
	})
	mockPlugin.On("GetName").Return(fakeDesc)
	strAddr, addr := "addr1", new(mocks.Address)
	mockPlugin.On("AddressFromString", strAddr, fakeTicker).Return(addr, nil)
	RegisterAltcoin(mockPlugin)

	_, err := AddressFromString("someAddress", "someCoin")
	require.NotNil(t, err)

	newAddr, err2 := AddressFromString("someAddress", "someCoin")
	require.Nil(t, err)
	require.Equal(t, addr, newAddr)
}
