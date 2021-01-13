package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"

	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateInvite{}

func NewMsgCreateInvite(inviter string, level IdentityLevel, identityType IdentityType) *MsgCreateInvite {
	return &MsgCreateInvite{
		Inviter:      inviter,
		Level:        level,
		IdentityType: identityType,
		CreationDt:   mbutils.CreateCurrentDate(),
	}
}

func (msg *MsgCreateInvite) Route() string {
	return RouterKey
}

func (msg *MsgCreateInvite) Type() string {
	return "CreateInvite"
}

func (msg *MsgCreateInvite) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Inviter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInvite) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInvite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Inviter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	created := mbutils.Created{CreationDt: msg.CreationDt}

	if err := created.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgAcceptInvite{}

func NewMsgAcceptInvite(inviteId string, invitee string) *MsgAcceptInvite {
	return &MsgAcceptInvite{
		InviteId:     inviteId,
		Invitee:      invitee,
		Uid:          uuid.New().String(),
		AcceptanceDt: mbutils.CreateCurrentDate(),
	}
}

func (msg *MsgAcceptInvite) Route() string {
	return RouterKey
}

func (msg *MsgAcceptInvite) Type() string {
	return "AcceptInvite"
}

func (msg *MsgAcceptInvite) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Invitee)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptInvite) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptInvite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Invitee)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	created := mbutils.Created{CreationDt: msg.AcceptanceDt}

	if err := created.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	return nil
}
