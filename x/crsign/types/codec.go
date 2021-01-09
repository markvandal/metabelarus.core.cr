package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&  MsgCreateId2Auth{}, "crsign/CreateId2Auth", nil)
cdc.RegisterConcrete(&MsgUpdateId2Auth{}, "crsign/UpdateId2Auth", nil)
cdc.RegisterConcrete(&MsgDeleteId2Auth{}, "crsign/DeleteId2Auth", nil)

cdc.RegisterConcrete(&MsgCreateAuth{}, "crsign/CreateAuth", nil)
cdc.RegisterConcrete(&MsgUpdateAuth{}, "crsign/UpdateAuth", nil)
cdc.RegisterConcrete(&MsgDeleteAuth{}, "crsign/DeleteAuth", nil)

cdc.RegisterConcrete(&MsgCreateId2Sign{}, "crsign/CreateId2Sign", nil)
cdc.RegisterConcrete(&MsgUpdateId2Sign{}, "crsign/UpdateId2Sign", nil)
cdc.RegisterConcrete(&MsgDeleteId2Sign{}, "crsign/DeleteId2Sign", nil)

cdc.RegisterConcrete(&MsgCreateSignature{}, "crsign/CreateSignature", nil)
cdc.RegisterConcrete(&MsgUpdateSignature{}, "crsign/UpdateSignature", nil)
cdc.RegisterConcrete(&MsgDeleteSignature{}, "crsign/DeleteSignature", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&  MsgCreateId2Auth{},
	&MsgUpdateId2Auth{},
	&MsgDeleteId2Auth{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateAuth{},
	&MsgUpdateAuth{},
	&MsgDeleteAuth{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateId2Sign{},
	&MsgUpdateId2Sign{},
	&MsgDeleteId2Sign{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateSignature{},
	&MsgUpdateSignature{},
	&MsgDeleteSignature{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
