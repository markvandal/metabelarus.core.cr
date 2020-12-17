package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateExtservice{}

type MsgCreateExtservice struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Alias string `json:"alias" yaml:"alias"`
  DetailsUrl string `json:"detailsUrl" yaml:"detailsUrl"`
}

func NewMsgCreateExtservice(creator sdk.AccAddress, alias string, detailsUrl string) MsgCreateExtservice {
  return MsgCreateExtservice{
    ID: uuid.New().String(),
		Creator: creator,
    Alias: alias,
    DetailsUrl: detailsUrl,
	}
}

func (msg MsgCreateExtservice) Route() string {
  return RouterKey
}

func (msg MsgCreateExtservice) Type() string {
  return "CreateExtservice"
}

func (msg MsgCreateExtservice) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateExtservice) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateExtservice) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}