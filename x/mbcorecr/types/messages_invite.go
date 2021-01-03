package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
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
