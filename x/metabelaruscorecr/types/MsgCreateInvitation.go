package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateInvitation{}

type MsgCreateInvitation struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  InviterId string `json:"inviterId" yaml:"inviterId"`
  IdentityId string `json:"identityId" yaml:"identityId"`
  CreationDate string `json:"creationDate" yaml:"creationDate"`
  ActivationPubKey string `json:"activationPubKey" yaml:"activationPubKey"`
}

func NewMsgCreateInvitation(creator sdk.AccAddress, inviterId string, identityId string, creationDate string, activationPubKey string) MsgCreateInvitation {
  return MsgCreateInvitation{
    ID: uuid.New().String(),
		Creator: creator,
    InviterId: inviterId,
    IdentityId: identityId,
    CreationDate: creationDate,
    ActivationPubKey: activationPubKey,
	}
}

func (msg MsgCreateInvitation) Route() string {
  return RouterKey
}

func (msg MsgCreateInvitation) Type() string {
  return "CreateInvitation"
}

func (msg MsgCreateInvitation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateInvitation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateInvitation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}