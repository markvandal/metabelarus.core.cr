package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteConfirmation{}

type MsgDeleteConfirmation struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteConfirmation(id string, creator sdk.AccAddress) MsgDeleteConfirmation {
  return MsgDeleteConfirmation{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteConfirmation) Route() string {
  return RouterKey
}

func (msg MsgDeleteConfirmation) Type() string {
  return "DeleteConfirmation"
}

func (msg MsgDeleteConfirmation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteConfirmation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteConfirmation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}