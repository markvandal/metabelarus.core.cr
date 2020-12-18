package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetConfirmation{}

type MsgSetConfirmation struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  IdenitityID string `json:"idenitityID" yaml:"idenitityID"`
  CreationDate string `json:"creationDate" yaml:"creationDate"`
  ExpirationDate string `json:"expirationDate" yaml:"expirationDate"`
  ConfirmatorID string `json:"confirmatorID" yaml:"confirmatorID"`
  CenterGeo string `json:"centerGeo" yaml:"centerGeo"`
  Status string `json:"status" yaml:"status"`
  NextTryDate string `json:"nextTryDate" yaml:"nextTryDate"`
}

func NewMsgSetConfirmation(creator sdk.AccAddress, id string, idenitityID string, creationDate string, expirationDate string, confirmatorID string, centerGeo string, status string, nextTryDate string) MsgSetConfirmation {
  return MsgSetConfirmation{
    ID: id,
		Creator: creator,
    IdenitityID: idenitityID,
    CreationDate: creationDate,
    ExpirationDate: expirationDate,
    ConfirmatorID: confirmatorID,
    CenterGeo: centerGeo,
    Status: status,
    NextTryDate: nextTryDate,
	}
}

func (msg MsgSetConfirmation) Route() string {
  return RouterKey
}

func (msg MsgSetConfirmation) Type() string {
  return "SetConfirmation"
}

func (msg MsgSetConfirmation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetConfirmation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetConfirmation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}