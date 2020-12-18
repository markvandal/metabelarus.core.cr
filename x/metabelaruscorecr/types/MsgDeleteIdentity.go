package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteIdentity{}

type MsgDeleteIdentity struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteIdentity(id string, creator sdk.AccAddress) MsgDeleteIdentity {
  return MsgDeleteIdentity{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteIdentity) Route() string {
  return RouterKey
}

func (msg MsgDeleteIdentity) Type() string {
  return "DeleteIdentity"
}

func (msg MsgDeleteIdentity) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteIdentity) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteIdentity) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}