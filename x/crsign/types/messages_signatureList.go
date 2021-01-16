package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSignatureList{}

func NewMsgCreateSignatureList(creator string, rootSignatureId string) *MsgCreateSignatureList {
	return &MsgCreateSignatureList{
		Creator:         creator,
		RootSignatureId: rootSignatureId,
	}
}

func (msg *MsgCreateSignatureList) Route() string {
	return RouterKey
}

func (msg *MsgCreateSignatureList) Type() string {
	return "CreateSignatureList"
}

func (msg *MsgCreateSignatureList) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignatureList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignatureList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSignatureList{}

func NewMsgUpdateSignatureList(creator string, id string, rootSignatureId string, lastSignatureId string, nextSignatureId string, metadata string) *MsgUpdateSignatureList {
	return &MsgUpdateSignatureList{
		Id:              id,
		Creator:         creator,
		RootSignatureId: rootSignatureId,
		LastSignatureId: lastSignatureId,
		NextSignatureId: nextSignatureId,
		Metadata:        metadata,
	}
}

func (msg *MsgUpdateSignatureList) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSignatureList) Type() string {
	return "UpdateSignatureList"
}

func (msg *MsgUpdateSignatureList) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSignatureList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSignatureList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
