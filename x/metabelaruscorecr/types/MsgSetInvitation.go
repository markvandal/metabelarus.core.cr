package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetInvitation{}

type MsgSetInvitation struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  InviterId string `json:"inviterId" yaml:"inviterId"`
  IdentityId string `json:"identityId" yaml:"identityId"`
  CreationDate string `json:"creationDate" yaml:"creationDate"`
  ActivationPubKey string `json:"activationPubKey" yaml:"activationPubKey"`
}

func NewMsgSetInvitation(creator sdk.AccAddress, id string, inviterId string, identityId string, creationDate string, activationPubKey string) MsgSetInvitation {
  return MsgSetInvitation{
    ID: id,
		Creator: creator,
    InviterId: inviterId,
    IdentityId: identityId,
    CreationDate: creationDate,
    ActivationPubKey: activationPubKey,
	}
}

func (msg MsgSetInvitation) Route() string {
  return RouterKey
}

func (msg MsgSetInvitation) Type() string {
  return "SetInvitation"
}

func (msg MsgSetInvitation) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetInvitation) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetInvitation) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}