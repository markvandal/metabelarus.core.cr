package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateInvite{}, "mbcorecr/CreateInvite", nil)

	cdc.RegisterConcrete(&MsgUpdateIdentity{}, "mbcorecr/UpdateIdentity", nil)

	cdc.RegisterConcrete(&MsgCreateSuperIdentity{}, "mbcorecr/CreateSuperIdentity", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInvite{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateIdentity{},
		&MsgCreateSuperIdentity{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
