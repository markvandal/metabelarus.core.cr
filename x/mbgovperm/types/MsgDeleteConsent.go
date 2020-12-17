package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteConsent{}

type MsgDeleteConsent struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteConsent(id string, creator sdk.AccAddress) MsgDeleteConsent {
  return MsgDeleteConsent{
    ID: id,
		Creator: creator,
	}
}

func (msg MsgDeleteConsent) Route() string {
  return RouterKey
}

func (msg MsgDeleteConsent) Type() string {
  return "DeleteConsent"
}

func (msg MsgDeleteConsent) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteConsent) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteConsent) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}