package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetAllowance{}

type MsgSetAllowance struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  PassportId string `json:"passportId" yaml:"passportId"`
  AllowanceType string `json:"allowanceType" yaml:"allowanceType"`
  Resolution string `json:"resolution" yaml:"resolution"`
}

func NewMsgSetAllowance(creator sdk.AccAddress, id string, passportId string, allowanceType string, resolution string) MsgSetAllowance {
  return MsgSetAllowance{
    ID: id,
		Creator: creator,
    PassportId: passportId,
    AllowanceType: allowanceType,
    Resolution: resolution,
	}
}

func (msg MsgSetAllowance) Route() string {
  return RouterKey
}

func (msg MsgSetAllowance) Type() string {
  return "SetAllowance"
}

func (msg MsgSetAllowance) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetAllowance) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetAllowance) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}