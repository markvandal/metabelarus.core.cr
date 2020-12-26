package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	mbutils "github.com/metabelarus/mbcorecr/mb/utils"

	"time"
)

var _ sdk.Msg = &MsgCreateIdentity{}

// NewMsgCreateIdentity - build message structure to create new Identity
func NewMsgCreateIdentity(creator string, AccountID string, IdentityType IdentityType, Details string) *MsgCreateIdentity {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, mbutils.BelarusLocation)

	return &MsgCreateIdentity{
		Creator:      creator,
		AccountID:    AccountID,
		IdentityType: IdentityType,
		Details:      Details,
		CreationDt:   &tm,
	}
}

func (msg *MsgCreateIdentity) Route() string {
	return RouterKey
}

func (msg *MsgCreateIdentity) Type() string {
	return "CreateIdentity"
}

func (msg *MsgCreateIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	created := mbutils.Created{CreationDt: msg.CreationDt}
	err = created.ValidateBasic()
	if err != nil {
		return sdkerrors.Wrap(ErrDateIssue, err.Error())
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateIdentity{}

func NewMsgUpdateIdentity(creator string, id string, AccountID string, IdentityType IdentityType, Details string) *MsgUpdateIdentity {
	return &MsgUpdateIdentity{
		Id:           id,
		Creator:      creator,
		AccountID:    AccountID,
		IdentityType: IdentityType,
		Details:      Details,
	}
}

func (msg *MsgUpdateIdentity) Route() string {
	return RouterKey
}

func (msg *MsgUpdateIdentity) Type() string {
	return "UpdateIdentity"
}

func (msg *MsgUpdateIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgCreateIdentity{}

func NewMsgDeleteIdentity(creator string, id string) *MsgDeleteIdentity {
	return &MsgDeleteIdentity{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteIdentity) Route() string {
	return RouterKey
}

func (msg *MsgDeleteIdentity) Type() string {
	return "DeleteIdentity"
}

func (msg *MsgDeleteIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
