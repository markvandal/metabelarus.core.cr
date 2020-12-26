package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Invitation struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    InviterId string `json:"inviterId" yaml:"inviterId"`
    IdentityId string `json:"identityId" yaml:"identityId"`
    CreationDate string `json:"creationDate" yaml:"creationDate"`
    ActivationPubKey string `json:"activationPubKey" yaml:"activationPubKey"`
}