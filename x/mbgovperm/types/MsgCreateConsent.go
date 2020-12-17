package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateConsent{}

type MsgCreateConsent struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  ExtserviceId string `json:"extserviceId" yaml:"extserviceId"`
  PassportId string `json:"passportId" yaml:"passportId"`
  Resolution string `json:"resolution" yaml:"resolution"`
}

func NewMsgCreateConsent(creator sdk.AccAddress, extserviceId string, passportId string, resolution string) MsgCreateConsent {
  return MsgCreateConsent{
    ID: uuid.New().String(),
		Creator: creator,
    ExtserviceId: extserviceId,
    PassportId: passportId,
    Resolution: resolution,
	}
}

func (msg MsgCreateConsent) Route() string {
  return RouterKey
}

func (msg MsgCreateConsent) Type() string {
  return "CreateConsent"
}

func (msg MsgCreateConsent) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateConsent) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateConsent) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}