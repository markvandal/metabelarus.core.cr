package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreateIdentity{}, "mbcorecr/CreateIdentity", nil)
cdc.RegisterConcrete(&MsgUpdateIdentity{}, "mbcorecr/UpdateIdentity", nil)
cdc.RegisterConcrete(&MsgDeleteIdentity{}, "mbcorecr/DeleteIdentity", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateIdentity{},
	&MsgUpdateIdentity{},
	&MsgDeleteIdentity{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
