package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetExtservice{}

type MsgSetExtservice struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Alias string `json:"alias" yaml:"alias"`
  DetailsUrl string `json:"detailsUrl" yaml:"detailsUrl"`
}

func NewMsgSetExtservice(creator sdk.AccAddress, id string, alias string, detailsUrl string) MsgSetExtservice {
  return MsgSetExtservice{
    ID: id,
		Creator: creator,
    Alias: alias,
    DetailsUrl: detailsUrl,
	}
}

func (msg MsgSetExtservice) Route() string {
  return RouterKey
}

func (msg MsgSetExtservice) Type() string {
  return "SetExtservice"
}

func (msg MsgSetExtservice) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetExtservice) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetExtservice) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}