package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAuth{}

func NewMsgCreateAuth(creator string, identity string, service string, key string, status string, creationDt string, availabilityDt string) *MsgCreateAuth {
  return &MsgCreateAuth{
		Creator: creator,
    Identity: identity,
    Service: service,
    Key: key,
    Status: status,
    CreationDt: creationDt,
    AvailabilityDt: availabilityDt,
	}
}

func (msg *MsgCreateAuth) Route() string {
  return RouterKey
}

func (msg *MsgCreateAuth) Type() string {
  return "CreateAuth"
}

func (msg *MsgCreateAuth) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAuth) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAuth) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateAuth{}

func NewMsgUpdateAuth(creator string, id string, identity string, service string, key string, status string, creationDt string, availabilityDt string) *MsgUpdateAuth {
  return &MsgUpdateAuth{
        Id: id,
		Creator: creator,
    Identity: identity,
    Service: service,
    Key: key,
    Status: status,
    CreationDt: creationDt,
    AvailabilityDt: availabilityDt,
	}
}

func (msg *MsgUpdateAuth) Route() string {
  return RouterKey
}

func (msg *MsgUpdateAuth) Type() string {
  return "UpdateAuth"
}

func (msg *MsgUpdateAuth) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAuth) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAuth) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgCreateAuth{}

func NewMsgDeleteAuth(creator string, id string) *MsgDeleteAuth {
  return &MsgDeleteAuth{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteAuth) Route() string {
  return RouterKey
}

func (msg *MsgDeleteAuth) Type() string {
  return "DeleteAuth"
}

func (msg *MsgDeleteAuth) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAuth) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAuth) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
