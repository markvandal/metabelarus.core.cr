package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateConsent{}, "mbgovperm/CreateConsent", nil)
		cdc.RegisterConcrete(MsgSetConsent{}, "mbgovperm/SetConsent", nil)
		cdc.RegisterConcrete(MsgDeleteConsent{}, "mbgovperm/DeleteConsent", nil)
		cdc.RegisterConcrete(MsgCreateExtservice{}, "mbgovperm/CreateExtservice", nil)
		cdc.RegisterConcrete(MsgSetExtservice{}, "mbgovperm/SetExtservice", nil)
		cdc.RegisterConcrete(MsgDeleteExtservice{}, "mbgovperm/DeleteExtservice", nil)
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
