package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateRecord{}, "mbpassport/CreateRecord", nil)
		cdc.RegisterConcrete(MsgSetRecord{}, "mbpassport/SetRecord", nil)
		cdc.RegisterConcrete(MsgDeleteRecord{}, "mbpassport/DeleteRecord", nil)
	// TODO: Register the modules msgs
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
