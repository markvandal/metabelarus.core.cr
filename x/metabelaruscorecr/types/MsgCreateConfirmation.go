package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateConfirmation{}

type MsgCreateConfirmation struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  IdenitityID string `json:"idenitityID" yaml:"idenitityID"`
  CreationDate string `json:"creationDate" yaml:"creationDate"`
  ExpirationDate string `json:"expirationDate" yaml:"expirationDate"`
  ConfirmatorID string `json:"confirmatorID" yaml:"confirmatorID"`
  CenterGeo string `json:"centerGeo" yaml:"centerGeo"`
  Status string `json:"status" yaml:"status"`
  NextTryDate string `json:"nextTryDate" yaml:"nextTryDate"`
}

func NewMsgCreateConfirmation(creator sdk.AccAddress, idenitityID string, creationDate string, expirationDate string, confirmatorID string, centerGeo string, status string, nextTryDate string) MsgCreateConfirmation {
  return MsgCreateConfirmation{
    ID: uuid.New().String(),
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

func (msg MsgCreateConfirmation) Route() string {
  return RouterKey
}

func (msg MsgCreateConfirmation) Type() string {
  return "CreateConfirmation"
}

func (msg MsgCreateConfirmation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateConfirmation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateConfirmation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}