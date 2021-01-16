package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateInvite{}, "mbcorecr/CreateInvite", nil)
	cdc.RegisterConcrete(&MsgAcceptInvite{}, "mbcorecr/AcceptInvite", nil)

	cdc.RegisterConcrete(&MsgUpdateIdentity{}, "mbcorecr/UpdateIdentity", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInvite{},
		&MsgAcceptInvite{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateIdentity{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
