package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteRecord{}

type MsgDeleteRecord struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteRecord(id string, creator sdk.AccAddress) MsgDeleteRecord {
  return MsgDeleteRecord{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteRecord) Route() string {
  return RouterKey
}

func (msg MsgDeleteRecord) Type() string {
  return "DeleteRecord"
}

func (msg MsgDeleteRecord) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteRecord) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteRecord) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}