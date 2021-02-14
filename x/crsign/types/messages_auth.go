package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

var _ sdk.Msg = &MsgRequestAuth{}

func NewMsgRequestAuth(service string, identity string, key string) *MsgRequestAuth {
	return &MsgRequestAuth{
		Identity:   identity,
		Service:    service,
		Key:        key,
		CreationDt: mbutils.CreateCurrentTime(),
	}
}

func (msg *MsgRequestAuth) Route() string {
	return RouterKey
}

func (msg *MsgRequestAuth) Type() string {
	return "CreateAuth"
}

func (msg *MsgRequestAuth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Service)
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
	_, err := sdk.AccAddressFromBech32(msg.Service)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := mbutils.ValidateKey(msg.Key, "Auth Key"); err != nil {
		return err
	}
	if err := mbutils.ValidateId(msg.Identity, "Identity ID"); err != nil {
		return err
	}

	created := mbutils.TimePoint{msg.CreationDt}
	if err := created.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgConfirmAuth{}

func NewMsgConfirmAuth(identity string, service string) *MsgConfirmAuth {
	return &MsgConfirmAuth{
		Identity:       identity,
		Service:        service,
		ConfirmationDt: mbutils.CreateCurrentTime(),
	}
}

func (msg *MsgConfirmAuth) Route() string {
	return RouterKey
}

func (msg *MsgConfirmAuth) Type() string {
	return "ConfirmAuth"
}

func (msg *MsgConfirmAuth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Identity)
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
	_, err := sdk.AccAddressFromBech32(msg.Identity)

	created := mbutils.TimePoint{msg.ConfirmationDt}

	if err := created.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
