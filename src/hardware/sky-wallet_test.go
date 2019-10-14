package hardware

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/hardware/mocks"
	"github.com/micro/protobuf/proto"
	"github.com/skycoin/hardware-wallet-go/src/skywallet/wire"
	"github.com/skycoin/hardware-wallet-protob/go"
	"github.com/stretchrcom/testify/require"
	"testing"
)

func TestGetSignerUIDShouldBeOk(t *testing.T) {
	// Giving
	dev := mocks.Devicer{}
	expectedDevId := "c24a046d-7d7b-484d-baf6-abedd883023f"
	f := messages.Features{
		DeviceId: proto.String(expectedDevId),
	}
	fb, err := proto.Marshal(&f)
	require.Nil(t, err)
	msg := wire.Message{
		Kind: uint16(messages.MessageType_MessageType_Features),
		Data: fb,
	}
	dev.On("GetFeatures").Return(msg, nil)
	sw := NewSkyWallet(&dev)

	// When
	devId := sw.GetSignerUID()

	// Then
	require.Equal(t, core.UID(expectedDevId), devId)
}