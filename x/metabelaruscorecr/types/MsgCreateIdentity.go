package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateIdentity{}

type MsgCreateIdentity struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  AccountID string `json:"accountID" yaml:"accountID"`
  Details string `json:"details" yaml:"details"`
  CreationDt string `json:"creationDt" yaml:"creationDt"`
  IdenitityType string `json:"idenitityType" yaml:"idenitityType"`
  AuthPubKey string `json:"authPubKey" yaml:"authPubKey"`
}

func NewMsgCreateIdentity(creator sdk.AccAddress, accountID string, details string, creationDt string, idenitityType string, authPubKey string) MsgCreateIdentity {
  return MsgCreateIdentity{
    ID: uuid.New().String(),
		Creator: creator,
    AccountID: accountID,
    Details: details,
    CreationDt: creationDt,
    IdenitityType: idenitityType,
    AuthPubKey: authPubKey,
	}
}

func (msg MsgCreateIdentity) Route() string {
  return RouterKey
}

func (msg MsgCreateIdentity) Type() string {
  return "CreateIdentity"
}

func (msg MsgCreateIdentity) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateIdentity) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateIdentity) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}