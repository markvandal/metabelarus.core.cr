package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateAllowance{}, "mbpasstrust/CreateAllowance", nil)
		cdc.RegisterConcrete(MsgSetAllowance{}, "mbpasstrust/SetAllowance", nil)
		cdc.RegisterConcrete(MsgDeleteAllowance{}, "mbpasstrust/DeleteAllowance", nil)
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
