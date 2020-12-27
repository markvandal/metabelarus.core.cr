package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSuperIdentity{}

// NewNewMsgCreateIdentity - create message to generate supper identity
func NewNewMsgCreateIdentity(creator string) *MsgCreateSuperIdentity {
	return &MsgCreateSuperIdentity{
		Creator: creator,
	}
}

func (msg *MsgCreateSuperIdentity) Route() string {
	return RouterKey
}

func (msg *MsgCreateSuperIdentity) Type() string {
	return "CreateSuperIdentity"
}

func (msg *MsgCreateSuperIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSuperIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSuperIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
