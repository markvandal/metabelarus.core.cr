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

	cdc.RegisterConcrete(&MsgCreateRecord{}, "crsign/CreateRecord", nil)
	cdc.RegisterConcrete(&MsgUpdateRecord{}, "crsign/UpdateRecord", nil)
	cdc.RegisterConcrete(&MsgDeleteRecord{}, "crsign/DeleteRecord", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestAuth{},
		&MsgConfirmAuth{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRecord{},
		&MsgUpdateRecord{},
		&MsgDeleteRecord{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
