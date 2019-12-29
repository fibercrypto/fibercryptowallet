package skycoin

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"

	"github.com/SkycoinProject/skycoin/src/readable"
)

func TestSkycoinPexNode(t *testing.T) {
	addr := "addr"
	port := uint16(8000)
	sent := int64(20)
	recived := int64(42)
	height := uint64(200)
	trusted := true

	conn := readable.Connection{
		Addr:          addr,
		ListenPort:    port,
		LastSent:      sent,
		LastReceived:  recived,
		Height:        height,
		IsTrustedPeer: trusted,
	}
	pex := connectionsToNetwork(conn)

	require.Equal(t, addr, pex.GetIp())
	require.Equal(t, port, pex.GetPort())
	require.Equal(t, sent, pex.GetLastSeenIn())
	require.Equal(t, recived, pex.GetLastSeenOut())
	require.Equal(t, height, pex.GetBlockHeight())
	require.Equal(t, trusted, pex.IsTrusted())
}

func TestSkycoinPexNodeIterator(t *testing.T) {
	pexs := make([]core.PexNode, 0)
	addrs := []string{"addr1", "addr2", "addr3"}
	for _, addr := range addrs {
		pex := &mocks.PexNode{}
		pex.On("GetIp").Return(addr)
		pexs = append(pexs, pex)
	}

	it := NewSkycoinPexNodeIterator(pexs)
	for _, addr := range addrs {
		require.Equal(t, true, it.Next())
		require.Equal(t, addr, it.Value().GetIp())
	}
	require.Equal(t, false, it.Next())
}
