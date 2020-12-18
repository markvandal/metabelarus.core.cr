package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Identity struct {
	ID            string         `json:"id" yaml:"id"`
	AccountID     sdk.AccAddress `json:"accountID" yaml:"accountID"`
	Details       string         `json:"details" yaml:"details"`
	CreationDt    time.Time      `json:"creationDt" yaml:"creationDt"`
	IdenitityType IdentityType   `json:"idenitityType" yaml:"idenitityType"`
	AuthPubKey    string         `json:"authPubKey" yaml:"authPubKey"`
}

type IdentityType int

const (
	IdentityTypeCitizen IdentityType = iota
	IdentityTypeForeigner
	IdentityTypeDiasporaMember
)
