package skywallet

import (
	"math/rand"
	"testing"

	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/stretchr/testify/suite"
)

type bitEncodedFlagsSuit struct {
	suite.Suite
}

func TestBitEncodedFlagsSuit(t *testing.T) {
	suite.Run(t, new(bitEncodedFlagsSuit))
}

func (suite *bitEncodedFlagsSuit) TestOperationsAreReversible() {
	for i := 0; i < 100; i++ {
		// NOTE: Giving
		flags := rand.Uint64() % 32
		ff := NewFirmwareFeatures(flags)
		// NOTE: When
		suite.NoError(ff.Unmarshal())
		f, e := ff.Marshal()
		// NOTE: Assert
		suite.NoError(e)
		suite.Equal(flags, f)
	}
}

func (suite *bitEncodedFlagsSuit) TestShouldHaveRdpEnabled() {
	// NOTE: Giving
	ff := NewFirmwareFeatures(uint64(messages.FirmwareFeatures_FirmwareFeatures_RdpMemProtect))
	// NOTE: When
	suite.NoError(ff.Unmarshal())
	// NOTE: Assert
	suite.True(ff.HasRdpMemProtectEnabled())
}

func (suite *bitEncodedFlagsSuit) TestShouldHaveNtRdpEnabled() {
	// NOTE: Giving
	ff := NewFirmwareFeatures(uint64(messages.FirmwareFeatures_FirmwareFeatures_RdpDebugDisabled))
	// NOTE: When
	suite.NoError(ff.Unmarshal())
	// NOTE: Assert
	suite.False(ff.HasRdpMemProtectEnabled())
}
