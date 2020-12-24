package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSetIdentity{}

// MsgSetIdentity — Set Identity Message structure
type MsgSetIdentity struct {
	ID            string         `json:"id" yaml:"id"`
	AccountID     sdk.AccAddress `json:"accountID" yaml:"accountID"`
	IdenitityType IdentityType   `json:"idenitityType" yaml:"idenitityType"`
	Details       string         `json:"details" yaml:"details"`
	AuthPubKey    string         `json:"authPubKey" yaml:"authPubKey"`
}

// NewMsgSetIdentity — create new identity message
func NewMsgSetIdentity(ID string, accountID sdk.AccAddress, idenitityType IdentityType, details string, authPubKey string) MsgSetIdentity {
	return MsgSetIdentity{
		ID:            ID,
		AccountID:     accountID,
		Details:       details,
		IdenitityType: idenitityType,
		AuthPubKey:    authPubKey,
		// CreationDt:    time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, BelarusLocation),
	}
}

func (msg MsgSetIdentity) Route() string {
	return RouterKey
}

func (msg MsgSetIdentity) Type() string {
	return "SetIdentity"
}

// GetSigners - it should return participated signers
/**
 * @TODO it looks like it doesn't work without some signer as create-Identity does
 * We need to find another signer (e.g. the request owner)
 */
func (msg MsgSetIdentity) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.AccountID)}
}

func (msg MsgSetIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic - validate if message has all data
func (msg MsgSetIdentity) ValidateBasic() error {
	return nil
}
