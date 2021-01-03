package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

var _ sdk.Msg = &MsgCreateSuperIdentity{}

// NewMsgCreateSuperIdentity - create message to generate supper identity
func NewMsgCreateSuperIdentity(creator string) *MsgCreateSuperIdentity {
	return &MsgCreateSuperIdentity{
		Creator:    creator,
		Uid:        uuid.New().String(),
		CreationDt: mbutils.CreateCurrentDate(),
	}
}

func (msg *MsgCreateSuperIdentity) Route() string {
	return RouterKey
}

func (msg *MsgCreateSuperIdentity) Type() string {
	return "CreateSuperIdentity"
}

func (msg *MsgCreateSuperIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSuperIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSuperIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	created := mbutils.Created{CreationDt: msg.CreationDt}

	if err := created.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	return nil
}
