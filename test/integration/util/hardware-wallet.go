package util

import (
	"github.com/fibercrypto/skywallet-go/src/skywallet"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/stretchr/testify/require"
	"testing"
)

func forceWipe(t *testing.T, dev skywallet.Devicer) {
	err := dev.SetAutoPressButton(true, skywallet.ButtonRight)
	require.NoError(t, err)
	msg, err := dev.Wipe()
	require.NoError(t, err)
	require.Equal(t, uint16(messages.MessageType_MessageType_Success), msg.Kind)
}

func ForceSetMnemonic(t *testing.T, dev skywallet.Devicer, mnemonic string) {
	forceWipe(t, dev)
	msg, err := dev.SetMnemonic(mnemonic)
	require.NoError(t, err)
	require.Equal(t, uint16(messages.MessageType_MessageType_Success), msg.Kind)
}
