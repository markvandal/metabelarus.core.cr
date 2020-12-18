package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgSetIdentity{}

// MsgSetIdentity — Set Identity Message structure
type MsgSetIdentity struct {
	ID            string         `json:"id" yaml:"id"`
	AccountID     sdk.AccAddress `json:"accountID" yaml:"accountID"`
	IdenitityType IdentityType   `json:"idenitityType" yaml:"idenitityType"`
	Details       string         `json:"details" yaml:"details"`
	AuthPubKey    string         `json:"authPubKey" yaml:"authPubKey"`
	CreationDt    time.Time      `json:"creationDt" yaml:"creationDt"`
}

// NewMsgSetIdentity — create new identity message
func NewMsgSetIdentity(accountID sdk.AccAddress, idenitityType IdentityType, details string, creationDt time.Time) (MsgSetIdentity, error) {
	return MsgSetIdentity{
		ID:            uuid.New().String(),
		AccountID:     accountID,
		Details:       details,
		IdenitityType: idenitityType,
		AuthPubKey:    "",
		CreationDt:    time.Date(creationDt.Year(), creationDt.Month(), creationDt.Day(), 0, 0, 0, 0, BelarusLocation),
	}, nil
}

func (msg MsgSetIdentity) Route() string {
	return RouterKey
}

func (msg MsgSetIdentity) Type() string {
	return "SetIdentity"
}

func (msg MsgSetIdentity) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.AccountID)}
}

func (msg MsgSetIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic - validate if message has all data
func (msg MsgSetIdentity) ValidateBasic() error {
	if msg.AccountID.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Identity should be connected to an account")
	}

	nowTime := time.Now()
	nowDate := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, BelarusLocation)

	if msg.CreationDt.After(nowDate) {
		return sdkerrors.Wrap(ErrDateIssue, "Try to create indentity after the current time")
	}

	if nowDate.Before(msg.CreationDt) && nowTime.After(nowDate.Add(time.Minute*5)) {
		return sdkerrors.Wrap(ErrDateIssue, "Try to create idenitty that was created long ago")
	}

	return nil
}
