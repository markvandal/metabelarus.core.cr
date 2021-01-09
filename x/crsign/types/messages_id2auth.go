package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateId2Auth{}

func NewMsgCreateId2Auth(creator string, identity string, auth string) *MsgCreateId2Auth {
	return &MsgCreateId2Auth{
		Creator:  creator,
		Identity: identity,
		Auth:     auth,
	}
}

func (msg *MsgCreateId2Auth) Route() string {
	return RouterKey
}

func (msg *MsgCreateId2Auth) Type() string {
	return "CreateId2Auth"
}

func (msg *MsgCreateId2Auth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateId2Auth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateId2Auth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateId2Auth{}

func NewMsgUpdateId2Auth(creator string, id string, identity string, auth string) *MsgUpdateId2Auth {
	return &MsgUpdateId2Auth{
		Id:       id,
		Creator:  creator,
		Identity: identity,
		Auth:     auth,
	}
}

func (msg *MsgUpdateId2Auth) Route() string {
	return RouterKey
}

func (msg *MsgUpdateId2Auth) Type() string {
	return "UpdateId2Auth"
}

func (msg *MsgUpdateId2Auth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateId2Auth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateId2Auth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateId2Auth{}

func NewMsgDeleteId2Auth(creator string, id string) *MsgDeleteId2Auth {
	return &MsgDeleteId2Auth{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteId2Auth) Route() string {
	return RouterKey
}

func (msg *MsgDeleteId2Auth) Type() string {
	return "DeleteId2Auth"
}

func (msg *MsgDeleteId2Auth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteId2Auth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteId2Auth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
