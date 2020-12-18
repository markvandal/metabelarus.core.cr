package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteInvitation{}

type MsgDeleteInvitation struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteInvitation(id string, creator sdk.AccAddress) MsgDeleteInvitation {
  return MsgDeleteInvitation{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteInvitation) Route() string {
  return RouterKey
}

func (msg MsgDeleteInvitation) Type() string {
  return "DeleteInvitation"
}

func (msg MsgDeleteInvitation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteInvitation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteInvitation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}