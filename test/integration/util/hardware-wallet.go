package util

import (
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	"github.com/fibercrypto/skywallet-go/src/skywallet/wire"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"testing"
	"github.com/stretchr/testify/require"
)

func forceWipe(t *testing.T, dev *skywallet.Device) {
	err := dev.SetAutoPressButton(true, skywallet.ButtonRight)
	require.NoError(t, err)
	msg, err := dev.Wipe()
	require.NoError(t, err)
	PressAcceptButton(t, dev, msg, messages.MessageType_MessageType_Success)
}

func PressAcceptButton(t *testing.T, dev skywallet.Devicer, prvMsg wire.Message, nextMsg messages.MessageType) wire.Message {
	require.Equal(t, uint16(messages.MessageType_MessageType_ButtonRequest), prvMsg.Kind)
	msg, err := dev.ButtonAck()
	require.NoError(t, err)
	require.Equal(t, uint16(nextMsg), msg.Kind)
	if msg.Kind == uint16(messages.MessageType_MessageType_Success) {
		successMsg, err := skywallet.DecodeSuccessMsg(msg)
		require.NoError(t, err)
		require.NotNil(t, successMsg)
	}
	return msg
}

func ForceSetMnemonic(t *testing.T, dev *skywallet.Device, mnemonic string) {
	forceWipe(t, dev)
	msg, err := dev.SetMnemonic(mnemonic)
	require.NoError(t, err)
	PressAcceptButton(t, dev, msg, messages.MessageType_MessageType_Success)
}
