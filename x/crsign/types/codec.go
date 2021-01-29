package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgRequestAuth{}, "crsign/RequestAuth", nil)
	cdc.RegisterConcrete(&MsgConfirmAuth{}, "crsign/ConfirmAuth", nil)

	cdc.RegisterConcrete(&MsgCreateSignature{}, "crsign/CreateSignature", nil)
	cdc.RegisterConcrete(&MsgUpdateSignature{}, "crsign/UpdateSignature", nil)
	cdc.RegisterConcrete(&MsgDeleteSignature{}, "crsign/DeleteSignature", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestAuth{},
		&MsgConfirmAuth{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSignature{},
		&MsgUpdateSignature{},
		&MsgDeleteSignature{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
