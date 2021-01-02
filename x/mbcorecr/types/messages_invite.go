package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateInvite{}

func NewMsgCreateInvite(creator string, inviter string, invitee string, level string, key string, creationDt string) *MsgCreateInvite {
  return &MsgCreateInvite{
		Creator: creator,
    Inviter: inviter,
    Invitee: invitee,
    Level: level,
    Key: key,
    CreationDt: creationDt,
	}
}

func (msg *MsgCreateInvite) Route() string {
  return RouterKey
}

func (msg *MsgCreateInvite) Type() string {
  return "CreateInvite"
}

func (msg *MsgCreateInvite) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInvite) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInvite) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateInvite{}

func NewMsgUpdateInvite(creator string, id string, inviter string, invitee string, level string, key string, creationDt string) *MsgUpdateInvite {
  return &MsgUpdateInvite{
        Id: id,
		Creator: creator,
    Inviter: inviter,
    Invitee: invitee,
    Level: level,
    Key: key,
    CreationDt: creationDt,
	}
}

func (msg *MsgUpdateInvite) Route() string {
  return RouterKey
}

func (msg *MsgUpdateInvite) Type() string {
  return "UpdateInvite"
}

func (msg *MsgUpdateInvite) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateInvite) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateInvite) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgCreateInvite{}

func NewMsgDeleteInvite(creator string, id string) *MsgDeleteInvite {
  return &MsgDeleteInvite{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteInvite) Route() string {
  return RouterKey
}

func (msg *MsgDeleteInvite) Type() string {
  return "DeleteInvite"
}

func (msg *MsgDeleteInvite) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteInvite) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteInvite) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
