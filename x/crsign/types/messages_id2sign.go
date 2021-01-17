package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateId2Sign{}

func NewMsgCreateId2Sign(creator string, identity string, signature string) *MsgCreateId2Sign {
  return &MsgCreateId2Sign{
		Creator: creator,
    Identity: identity,
    Signature: signature,
	}
}

func (msg *MsgCreateId2Sign) Route() string {
  return RouterKey
}

func (msg *MsgCreateId2Sign) Type() string {
  return "CreateId2Sign"
}

func (msg *MsgCreateId2Sign) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateId2Sign) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateId2Sign) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateId2Sign{}

func NewMsgUpdateId2Sign(creator string, id string, identity string, signature string) *MsgUpdateId2Sign {
  return &MsgUpdateId2Sign{
        Id: id,
		Creator: creator,
    Identity: identity,
    Signature: signature,
	}
}

func (msg *MsgUpdateId2Sign) Route() string {
  return RouterKey
}

func (msg *MsgUpdateId2Sign) Type() string {
  return "UpdateId2Sign"
}

func (msg *MsgUpdateId2Sign) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateId2Sign) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateId2Sign) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgCreateId2Sign{}

func NewMsgDeleteId2Sign(creator string, id string) *MsgDeleteId2Sign {
  return &MsgDeleteId2Sign{
        Id: id,
		Creator: creator,
	}
} 
func (msg *MsgDeleteId2Sign) Route() string {
  return RouterKey
}

func (msg *MsgDeleteId2Sign) Type() string {
  return "DeleteId2Sign"
}

func (msg *MsgDeleteId2Sign) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteId2Sign) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteId2Sign) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}
