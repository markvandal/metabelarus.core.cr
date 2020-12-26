package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateIdentity{}

// MsgCreateIdentity - message structure to create new Idenity
type MsgCreateIdentity struct {
	ID            string         `json:"id" yaml:"id"`
	AccountID     sdk.AccAddress `json:"accountID" yaml:"accountID"`
	IdenitityType IdentityType   `json:"idenitityType" yaml:"idenitityType"`
	Details       string         `json:"details" yaml:"details"`
	AuthPubKey    string         `json:"authPubKey" yaml:"authPubKey"`
	CreationDt    time.Time      `json:"creationDt" yaml:"creationDt"`
}

// NewMsgCreateIdentity - build message structure to create new Identity
func NewMsgCreateIdentity(accountID sdk.AccAddress, idenitityType IdentityType, details string, authPubKey string) MsgCreateIdentity {
	now := time.Now()

	return MsgCreateIdentity{
		ID:            uuid.New().String(),
		AccountID:     accountID,
		IdenitityType: idenitityType,
		Details:       details,
		AuthPubKey:    authPubKey,
		CreationDt:    time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, BelarusLocation),
	}
}

// Route - returns route key
func (msg MsgCreateIdentity) Route() string {
	return RouterKey
}

// Type - returns type
func (msg MsgCreateIdentity) Type() string {
	return "CreateIdentity"
}

// GetSigners - returns signers of Identity creation message
func (msg MsgCreateIdentity) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.AccountID)}
}

// GetSignBytes - returns bytes of the Identity creation message
func (msg MsgCreateIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic - Validate if creation message has account and correct date
func (msg MsgCreateIdentity) ValidateBasic() error {
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
