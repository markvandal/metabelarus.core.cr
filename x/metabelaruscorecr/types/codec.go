package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateConfirmation{}, "metabelaruscorecr/CreateConfirmation", nil)
		cdc.RegisterConcrete(MsgSetConfirmation{}, "metabelaruscorecr/SetConfirmation", nil)
		cdc.RegisterConcrete(MsgDeleteConfirmation{}, "metabelaruscorecr/DeleteConfirmation", nil)
		cdc.RegisterConcrete(MsgCreateInvitation{}, "metabelaruscorecr/CreateInvitation", nil)
		cdc.RegisterConcrete(MsgSetInvitation{}, "metabelaruscorecr/SetInvitation", nil)
		cdc.RegisterConcrete(MsgDeleteInvitation{}, "metabelaruscorecr/DeleteInvitation", nil)
		cdc.RegisterConcrete(MsgCreateIdentity{}, "metabelaruscorecr/CreateIdentity", nil)
		cdc.RegisterConcrete(MsgSetIdentity{}, "metabelaruscorecr/SetIdentity", nil)
		cdc.RegisterConcrete(MsgDeleteIdentity{}, "metabelaruscorecr/DeleteIdentity", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
