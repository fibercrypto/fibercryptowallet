package util

import (
	"fmt"
	"testing"

	"github.com/fibercrypto/fibercryptowallet/src/coin/mocks"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisteredPlugin(t *testing.T) {
	fakeTicker := "MOCKSCOIN"
	fakeDesc := "Fake coin"
	fakeExp := 3
	fakeQuotient := 1000 // 10 ^ fakeExp
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

	require.Equal(t, fakeDesc, AltcoinCaption(fakeTicker))
	q, err := AltcoinQuotient(fakeTicker)
	require.NoError(t, err)
	require.Equal(t, uint64(fakeQuotient), q)
}

func TestUnknownPlugin(t *testing.T) {
	fakeTicker := "MOCKSCOIN_UNK"
	require.Equal(t, "MOCKSCOIN_UNK <Unregistered>", AltcoinCaption(fakeTicker))
	_, err := AltcoinQuotient(fakeTicker)
	require.Error(t, err)
}

func TestLookupSignerByUID(t *testing.T) {
	type WalletSigner struct {
		mocks.Wallet
		mocks.TxnSigner
	}

	emptyUID := core.UID("")
	uid := core.UID("walletid")
	other := core.UID("otherid")

	ws := new(WalletSigner)
	ws.TxnSigner.On("GetSignerUID").Return(uid)
	var signer core.TxnSigner
	signer = ws
	err := AttachSignService(signer)
	require.NoError(t, err)

	tests := []struct {
		wallet core.Wallet
		id     core.UID
		want   core.TxnSigner
	}{
		{wallet: new(mocks.Wallet), id: emptyUID, want: nil},
		{wallet: ws, id: emptyUID, want: signer},
		{wallet: ws, id: uid, want: signer},
		{wallet: new(mocks.Wallet), id: uid, want: signer},
		{wallet: new(mocks.Wallet), id: other, want: nil},
		{wallet: ws, id: other, want: nil},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("lookup%d", i), func(t *testing.T) {
			require.Equal(t, tt.want, LookupSignerByUID(tt.wallet, tt.id))
		})
	}
}
