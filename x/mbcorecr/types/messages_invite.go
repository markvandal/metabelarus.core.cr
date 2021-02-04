package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

var _ sdk.Msg = &MsgCreateInvite{}

func NewMsgCreateInvite(
	inviter string,
	level IdentityLevel,
	identityType IdentityType,
	address string,
	pubKey string,
) *MsgCreateInvite {
	return &MsgCreateInvite{
		Inviter:      inviter,
		Level:        level,
		IdentityType: identityType,
		Address:      address,
		PubKey:       pubKey,
		CreationDt:   mbutils.CreateCurrentTime(),
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

	_, err = sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, msg.PubKey)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "Invalid public key (%s)", err)
	}

	created := mbutils.TimePoint{msg.CreationDt}

	if err := created.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgAcceptInvite{}

func NewMsgAcceptInvite(inviteId string, tmpAddress string, address string, pubKey string) *MsgAcceptInvite {
	return &MsgAcceptInvite{
		InviteId:     inviteId,
		TmpAddress:   tmpAddress,
		Address:      address,
		PubKey:       pubKey,
		AcceptanceDt: mbutils.CreateCurrentTime(),
	}
}

func (msg *MsgAcceptInvite) Route() string {
	return RouterKey
}

func (msg *MsgAcceptInvite) Type() string {
	return "AcceptInvite"
}

func (msg *MsgAcceptInvite) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TmpAddress)
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
	_, err := sdk.AccAddressFromBech32(msg.TmpAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid tmp account address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new account address (%s)", err)
	}

	_, err = sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, msg.PubKey)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "Invalid public key (%s)", err)
	}

	created := mbutils.TimePoint{msg.AcceptanceDt}

	if err := created.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	return nil
}
