package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetRecord{}

type MsgSetRecord struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  IdentityId string `json:"IdentityId" yaml:"IdentityId"`
  ServiceId string `json:"ServiceId" yaml:"ServiceId"`
  ServiceType string `json:"ServiceType" yaml:"ServiceType"`
  Key string `json:"Key" yaml:"Key"`
  UserValue string `json:"UserValue" yaml:"UserValue"`
  ServiceValue string `json:"ServiceValue" yaml:"ServiceValue"`
  CreationDt string `json:"CreationDt" yaml:"CreationDt"`
  UpdateDt string `json:"UpdateDt" yaml:"UpdateDt"`
}

func NewMsgSetRecord(creator sdk.AccAddress, id string, IdentityId string, ServiceId string, ServiceType string, Key string, UserValue string, ServiceValue string, CreationDt string, UpdateDt string) MsgSetRecord {
  return MsgSetRecord{
    ID: id,
		Creator: creator,
    IdentityId: IdentityId,
    ServiceId: ServiceId,
    ServiceType: ServiceType,
    Key: Key,
    UserValue: UserValue,
    ServiceValue: ServiceValue,
    CreationDt: CreationDt,
    UpdateDt: UpdateDt,
	}
}

func (msg MsgSetRecord) Route() string {
  return RouterKey
}

func (msg MsgSetRecord) Type() string {
  return "SetRecord"
}

func (msg MsgSetRecord) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetRecord) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetRecord) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}