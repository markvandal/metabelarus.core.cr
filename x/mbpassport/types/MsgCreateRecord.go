package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateRecord{}

type MsgCreateRecord struct {
  ID      string
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

func NewMsgCreateRecord(creator sdk.AccAddress, IdentityId string, ServiceId string, ServiceType string, Key string, UserValue string, ServiceValue string, CreationDt string, UpdateDt string) MsgCreateRecord {
  return MsgCreateRecord{
    ID: uuid.New().String(),
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

func (msg MsgCreateRecord) Route() string {
  return RouterKey
}

func (msg MsgCreateRecord) Type() string {
  return "CreateRecord"
}

func (msg MsgCreateRecord) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateRecord) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateRecord) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}