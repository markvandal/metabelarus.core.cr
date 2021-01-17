package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSignature{}

func NewMsgCreateSignature(creator string, identity string, service string, key string, secret string, creationDt string, availabilityDt string) *MsgCreateSignature {
  return &MsgCreateSignature{
		Creator: creator,
    Identity: identity,
    Service: service,
    Key: key,
    Secret: secret,
    CreationDt: creationDt,
    AvailabilityDt: availabilityDt,
	}
}

func (msg *MsgCreateSignature) Route() string {
  return RouterKey
}

func (msg *MsgCreateSignature) Type() string {
  return "CreateSignature"
}

func (msg *MsgCreateSignature) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSignature) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSignature) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateSignature{}

func NewMsgUpdateSignature(creator string, id string, identity string, service string, key string, secret string, creationDt string, availabilityDt string) *MsgUpdateSignature {
  return &MsgUpdateSignature{
        Id: id,
		Creator: creator,
    Identity: identity,
    Service: service,
    Key: key,
    Secret: secret,
    CreationDt: creationDt,
    AvailabilityDt: availabilityDt,
	}
}

func (msg *MsgUpdateSignature) Route() string {
  return RouterKey
}

func (msg *MsgUpdateSignature) Type() string {
  return "UpdateSignature"
}

func (msg *MsgUpdateSignature) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSignature) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSignature) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgCreateSignature{}

func NewMsgDeleteSignature(creator string, id string) *MsgDeleteSignature {
  return &MsgDeleteSignature{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteSignature) Route() string {
  return RouterKey
}

func (msg *MsgDeleteSignature) Type() string {
  return "DeleteSignature"
}

func (msg *MsgDeleteSignature) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSignature) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSignature) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
