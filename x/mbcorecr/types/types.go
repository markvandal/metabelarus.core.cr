package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	EventCreateInvite   = "mbcorecr:create.invite"
	EventCreateIdentity = "mbcorecr:create.identity"
)

const (
	EventAttrIdentityType    = "type"
	EventAttrTmpAddress      = "tmp_address"
	EventAttrIdentityAddress = "identity_address"
	EventAttrInviteId        = "invite_id"
	EventAttrIentityId       = "identity_id"

	DefaultWalletPath             = "44'/118'/0'/0/0" // @TODO should be read from config
	UnsecureNewAcctountPKPassword = "11112222"

	SuperInviteDenom = "invitesuper"
	Invite0Denom     = "invite0"
	Invite1Denom     = "invite1"
	Invite2Denom     = "invite2"
	Invite3Denom     = "invite3"
	Invite4Denom     = "invite4"
)

var (
	SuperIdentityCoinsPack = sdk.Coins{
		sdk.Coin{Denom: Invite0Denom, Amount: sdk.NewInt(150)},
		sdk.Coin{Denom: Invite1Denom, Amount: sdk.NewInt(100)},
		sdk.Coin{Denom: Invite2Denom, Amount: sdk.NewInt(50)},
		sdk.Coin{Denom: Invite3Denom, Amount: sdk.NewInt(15)},
		sdk.Coin{Denom: Invite4Denom, Amount: sdk.NewInt(5)},
	}

	IndentityCoinPacks = map[IdentityLevel]sdk.Coins{
		IdentityLevel_LevelSuper: sdk.Coins{
			sdk.Coin{Denom: Invite0Denom, Amount: sdk.NewInt(150)},
			sdk.Coin{Denom: Invite1Denom, Amount: sdk.NewInt(100)},
			sdk.Coin{Denom: Invite2Denom, Amount: sdk.NewInt(50)},
			sdk.Coin{Denom: Invite3Denom, Amount: sdk.NewInt(15)},
			sdk.Coin{Denom: Invite4Denom, Amount: sdk.NewInt(5)},
		},
		IdentityLevel_Level0: sdk.Coins{
			sdk.Coin{Denom: Invite1Denom, Amount: sdk.NewInt(100)},
			sdk.Coin{Denom: Invite2Denom, Amount: sdk.NewInt(50)},
			sdk.Coin{Denom: Invite3Denom, Amount: sdk.NewInt(15)},
			sdk.Coin{Denom: Invite4Denom, Amount: sdk.NewInt(5)},
		},
		IdentityLevel_Level1: sdk.Coins{
			sdk.Coin{Denom: Invite2Denom, Amount: sdk.NewInt(50)},
			sdk.Coin{Denom: Invite3Denom, Amount: sdk.NewInt(15)},
			sdk.Coin{Denom: Invite4Denom, Amount: sdk.NewInt(5)},
		},
		IdentityLevel_Level2: sdk.Coins{
			sdk.Coin{Denom: Invite3Denom, Amount: sdk.NewInt(15)},
			sdk.Coin{Denom: Invite4Denom, Amount: sdk.NewInt(5)},
		},
		IdentityLevel_Level3: sdk.Coins{
			sdk.Coin{Denom: Invite4Denom, Amount: sdk.NewInt(5)},
		},
		IdentityLevel_Level4: sdk.Coins{},
	}

	IdentityLevelToDenom = map[IdentityLevel]string{
		IdentityLevel_LevelSuper: SuperInviteDenom,
		IdentityLevel_Level0:     Invite0Denom,
		IdentityLevel_Level1:     Invite1Denom,
		IdentityLevel_Level2:     Invite2Denom,
		IdentityLevel_Level3:     Invite3Denom,
		IdentityLevel_Level4:     Invite4Denom,
	}
)
