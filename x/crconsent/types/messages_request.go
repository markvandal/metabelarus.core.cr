package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRequest{}

func NewMsgCreateRequest(creator string, initiator string, recipient string, requestType string, status string, value int32, memo string, promoUrl string, creationDt string, finalDt string) *MsgCreateRequest {
	return &MsgCreateRequest{
		Creator:     creator,
		Initiator:   initiator,
		Recipient:   recipient,
		RequestType: requestType,
		Status:      status,
		Value:       value,
		Memo:        memo,
		PromoUrl:    promoUrl,
		CreationDt:  creationDt,
		FinalDt:     finalDt,
	}
}

func (msg *MsgCreateRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateRequest) Type() string {
	return "CreateRequest"
}

func (msg *MsgCreateRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRequest{}

func NewMsgUpdateRequest(creator string, id string, initiator string, recipient string, requestType string, status string, value int32, memo string, promoUrl string, creationDt string, finalDt string) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		Id:          id,
		Creator:     creator,
		Initiator:   initiator,
		Recipient:   recipient,
		RequestType: requestType,
		Status:      status,
		Value:       value,
		Memo:        memo,
		PromoUrl:    promoUrl,
		CreationDt:  creationDt,
		FinalDt:     finalDt,
	}
}

func (msg *MsgUpdateRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRequest) Type() string {
	return "UpdateRequest"
}

func (msg *MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateRequest{}

func NewMsgDeleteRequest(creator string, id string) *MsgDeleteRequest {
	return &MsgDeleteRequest{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRequest) Type() string {
	return "DeleteRequest"
}

func (msg *MsgDeleteRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
