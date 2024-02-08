package txbuilder_test

import (
	"context"

	"github.com/ark-network/ark/internal/core/ports"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/stretchr/testify/mock"
)

type mockedWallet struct {
	mock.Mock
}

// BroadcastTransaction implements ports.WalletService.
func (m *mockedWallet) BroadcastTransaction(ctx context.Context, txHex string) (string, error) {
	args := m.Called(ctx, txHex)

	var res string
	if a := args.Get(0); a != nil {
		res = a.(string)
	}
	return res, args.Error(1)
}

// Close implements ports.WalletService.
func (m *mockedWallet) Close() {
	m.Called()
}

// DeriveAddresses implements ports.WalletService.
func (m *mockedWallet) DeriveAddresses(ctx context.Context, num int) ([]string, error) {
	args := m.Called(ctx, num)

	var res []string
	if a := args.Get(0); a != nil {
		res = a.([]string)
	}
	return res, args.Error(1)
}

// GetPubkey implements ports.WalletService.
func (m *mockedWallet) GetPubkey(ctx context.Context) (*secp256k1.PublicKey, error) {
	args := m.Called(ctx)

	var res *secp256k1.PublicKey
	if a := args.Get(0); a != nil {
		res = a.(*secp256k1.PublicKey)
	}
	return res, args.Error(1)
}

// SignPset implements ports.WalletService.
func (m *mockedWallet) SignPset(ctx context.Context, pset string, extractRawTx bool) (string, error) {
	args := m.Called(ctx, pset, extractRawTx)

	var res string
	if a := args.Get(0); a != nil {
		res = a.(string)
	}
	return res, args.Error(1)
}

// Status implements ports.WalletService.
func (m *mockedWallet) Status(ctx context.Context) (ports.WalletStatus, error) {
	args := m.Called(ctx)

	var res ports.WalletStatus
	if a := args.Get(0); a != nil {
		res = a.(ports.WalletStatus)
	}
	return res, args.Error(1)
}

func (m *mockedWallet) SelectUtxos(ctx context.Context, asset string, amount uint64) ([]ports.TxInput, uint64, error) {
	args := m.Called(ctx, asset, amount)

	var res0 func() []ports.TxInput
	if a := args.Get(0); a != nil {
		res0 = a.(func() []ports.TxInput)
	}
	var res1 uint64
	if a := args.Get(1); a != nil {
		res1 = a.(uint64)
	}
	return res0(), res1, args.Error(2)
}

func (m *mockedWallet) EstimateFees(ctx context.Context, pset string) (uint64, error) {
	args := m.Called(ctx, pset)

	var res uint64
	if a := args.Get(0); a != nil {
		res = a.(uint64)
	}
	return res, args.Error(1)
}

type mockedInput struct {
	mock.Mock
}

func (m *mockedInput) GetTxid() string {
	args := m.Called()

	var res string
	if a := args.Get(0); a != nil {
		res = a.(string)
	}

	return res
}

func (m *mockedInput) GetIndex() uint32 {
	args := m.Called()

	var res uint32
	if a := args.Get(0); a != nil {
		res = a.(uint32)
	}
	return res
}

func (m *mockedInput) GetScript() string {
	args := m.Called()

	var res string
	if a := args.Get(0); a != nil {
		res = a.(string)
	}
	return res
}

func (m *mockedInput) GetAsset() string {
	args := m.Called()

	var res string
	if a := args.Get(0); a != nil {
		res = a.(string)
	}
	return res
}

func (m *mockedInput) GetValue() uint64 {
	args := m.Called()

	var res uint64
	if a := args.Get(0); a != nil {
		res = a.(uint64)
	}
	return res
}