package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetConsent{}

type MsgSetConsent struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ExtserviceId string `json:"extserviceId" yaml:"extserviceId"`
  PassportId string `json:"passportId" yaml:"passportId"`
  Resolution string `json:"resolution" yaml:"resolution"`
}

func NewMsgSetConsent(creator sdk.AccAddress, id string, extserviceId string, passportId string, resolution string) MsgSetConsent {
  return MsgSetConsent{
    ID: id,
		Creator: creator,
    ExtserviceId: extserviceId,
    PassportId: passportId,
    Resolution: resolution,
	}
}

func (msg MsgSetConsent) Route() string {
  return RouterKey
}

func (msg MsgSetConsent) Type() string {
  return "SetConsent"
}

func (msg MsgSetConsent) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetConsent) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetConsent) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}