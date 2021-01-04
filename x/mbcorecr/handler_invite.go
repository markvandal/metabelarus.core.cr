package mbcorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	"github.com/metabelarus/mbcorecr/x/mbcorecr/helper"
)

func handleMsgCreateInvite(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateInvite) (*sdk.Result, error) {
	// Has invite coin of appropriate level
	inviteHelper, err := helper.NewInviteHelper(msg.Inviter, &ctx, &k.BankKeeper, &k.AuthKeeper)
	if err != nil {
		return nil, err
	}

	if err := inviteHelper.WithLevel(msg.Level); err != nil {
		return nil, err
	}

	// Create a temporary account
	tempAcc, err := helper.NewInviteAccount(msg.Uid, ctx, k.AuthKeeper)
	if err != nil {
		return nil, err
	}

	// Consume invite coin
	if err := inviteHelper.Pay(); err != nil {
		return nil, err
	}

	// Create invite entery
	// k.CreateInvite(ctx, *msg)

	// If level 4 set no balance (empty pack)
	if err := tempAcc.SetBalances(
		ctx, k.BankKeeper,
		types.SuperIdentityCoinsPack, // Cast pack from level
	); err != nil {
		return nil, err
	}

	invite := &types.Invite{
		Inviter:      msg.Inviter,
		Invitee:      "",
		Level:        msg.Level,
		IdentityType: msg.IdentityType,
		Key:          tempAcc.GetAddress(),
		CreationDt:   msg.CreationDt,
	}

	// [TODO] Make an owner from temporary account
	k.CreateInvite(ctx, invite)

	inviteAcc := types.InviteAccount{
		Uid:      tempAcc.Uid,
		Address:  tempAcc.Address,
		Mnemonic: tempAcc.Mnemonic,
		PubKey:   tempAcc.PubKey,
		PrivKey:  tempAcc.PrivKey,
		InviteID: invite.Id,
	}

	ecnryptedPayload, err := inviteAcc.EncryptData(inviteAcc)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventCreateInvite,
			sdk.NewAttribute(
				types.EventAttrIdentityType,
				types.IdentityType_name[int32(msg.IdentityType)],
			),
			sdk.NewAttribute(
				types.EventAttrIdentityPayload,
				ecnryptedPayload,
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
