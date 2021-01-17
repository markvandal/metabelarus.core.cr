package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

var _ sdk.Msg = &MsgRequestAuth{}

func NewMsgRequestAuth(creator string, service string, identity string, key string) *MsgRequestAuth {
	return &MsgRequestAuth{
		Creator:    creator,
		Identity:   identity,
		Service:    service,
		Key:        key,
		CreationDt: mbutils.CreateCurrentDate(),
	}
}

func (msg *MsgRequestAuth) Route() string {
	return RouterKey
}

func (msg *MsgRequestAuth) Type() string {
	return "CreateAuth"
}

func (msg *MsgRequestAuth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestAuth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAuth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// @TODO CreationDt should be checked for an hour +\-
	return nil
}

var _ sdk.Msg = &MsgConfirmAuth{}

func NewMsgConfirmAuth(creator string, service string, identity string) *MsgConfirmAuth {
	return &MsgConfirmAuth{
		Creator:  creator,
		Identity: identity,
		Service:  service,
	}
}

func (msg *MsgConfirmAuth) Route() string {
	return RouterKey
}

func (msg *MsgConfirmAuth) Type() string {
	return "ConfirmAuth"
}

func (msg *MsgConfirmAuth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConfirmAuth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfirmAuth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
