package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteExtservice{}

type MsgDeleteExtservice struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteExtservice(id string, creator sdk.AccAddress) MsgDeleteExtservice {
  return MsgDeleteExtservice{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteExtservice) Route() string {
  return RouterKey
}

func (msg MsgDeleteExtservice) Type() string {
  return "DeleteExtservice"
}

func (msg MsgDeleteExtservice) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteExtservice) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteExtservice) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}