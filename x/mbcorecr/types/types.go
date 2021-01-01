package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	EventGovCreateIdentity = "mbcorecr.gov:create.identity"
)

const (
	EventAttrIdentityType    = "type"
	EventAttrIdentityPayload = "payload"

	DefaultWalletPath             = "0'"
	UnsecureNewAcctountPKPassword = "11112222"

	SuperInviteDenom = "invitesuper"
	Invite0Denom     = "invite0"
	Invite1Denom     = "invite1"
	Invite2Denom     = "invite2"
	Invite3Denom     = "invite3"
	Invite4Denom     = "invite4"
)

type AttrIdentityType string

const (
	AttrIdentityTypeSuper AttrIdentityType = "super"
)

var (
	SuperIdentityCoinsPack = sdk.Coins{
		sdk.Coin{Denom: Invite0Denom, Amount: sdk.NewInt(150)},
		sdk.Coin{Denom: Invite1Denom, Amount: sdk.NewInt(100)},
		sdk.Coin{Denom: Invite2Denom, Amount: sdk.NewInt(50)},
		sdk.Coin{Denom: Invite3Denom, Amount: sdk.NewInt(15)},
		sdk.Coin{Denom: Invite4Denom, Amount: sdk.NewInt(5)},
	}
)
