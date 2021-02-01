package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

var _ sdk.Msg = &MsgCreateRecord{}

func NewMsgCreateRecord(
	creator string,
	provider string,
	key string,
	data string,
	signature string,
	recordType RecordType,
	publicity PublicityType,
	liveTime int32,
) *MsgCreateRecord {
	return &MsgCreateRecord{
		Creator:    creator,
		Provider:   provider,
		Key:        key,
		Data:       data,
		Signature:  signature,
		RecordType: recordType,
		Publicity:  publicity,
		LiveTime:   liveTime,
		CreationDt: mbutils.CreateCurrentDate(),
	}
}

func (msg *MsgCreateRecord) Route() string {
	return RouterKey
}

func (msg *MsgCreateRecord) Type() string {
	return "CreateRecord"
}

func (msg *MsgCreateRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRecord) ValidateBasic() error {
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

var _ sdk.Msg = &MsgUpdateRecord{}

func NewMsgUpdateRecord(
	updater string,
	id string,
	data string,
	signature string,
	liveTime int32,
	action RecordUpdate,
) *MsgUpdateRecord {
	return &MsgUpdateRecord{
		Id:        id,
		Updater:   updater,
		Data:      data,
		Signature: signature,
		LiveTime:  liveTime,
		Action:    action,
		UpdateDt:  mbutils.CreateCurrentDate(),
	}
}

func (msg *MsgUpdateRecord) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRecord) Type() string {
	return "UpdateRecord"
}

func (msg *MsgUpdateRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	created := mbutils.Created{CreationDt: msg.UpdateDt}

	if err := created.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrDateIssue, "invalid message date (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteRecord{}

func NewMsgDeleteRecord(deleter string, id string) *MsgDeleteRecord {
	return &MsgDeleteRecord{
		Id:      id,
		Deleter: deleter,
	}
}
func (msg *MsgDeleteRecord) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRecord) Type() string {
	return "DeleteRecord"
}

func (msg *MsgDeleteRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Deleter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Deleter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
