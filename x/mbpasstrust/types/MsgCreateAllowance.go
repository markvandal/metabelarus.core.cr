package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateAllowance{}

type MsgCreateAllowance struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  PassportId string `json:"passportId" yaml:"passportId"`
  AllowanceType string `json:"allowanceType" yaml:"allowanceType"`
  Resolution string `json:"resolution" yaml:"resolution"`
}

func NewMsgCreateAllowance(creator sdk.AccAddress, passportId string, allowanceType string, resolution string) MsgCreateAllowance {
  return MsgCreateAllowance{
    ID: uuid.New().String(),
		Creator: creator,
    PassportId: passportId,
    AllowanceType: allowanceType,
    Resolution: resolution,
	}
}

func (msg MsgCreateAllowance) Route() string {
  return RouterKey
}

func (msg MsgCreateAllowance) Type() string {
  return "CreateAllowance"
}

func (msg MsgCreateAllowance) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateAllowance) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateAllowance) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}