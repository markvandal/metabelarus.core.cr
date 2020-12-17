package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteAllowance{}

type MsgDeleteAllowance struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteAllowance(id string, creator sdk.AccAddress) MsgDeleteAllowance {
  return MsgDeleteAllowance{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteAllowance) Route() string {
  return RouterKey
}

func (msg MsgDeleteAllowance) Type() string {
  return "DeleteAllowance"
}

func (msg MsgDeleteAllowance) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteAllowance) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteAllowance) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}