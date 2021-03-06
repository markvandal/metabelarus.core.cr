package mbcorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	"github.com/metabelarus/mbcorecr/x/mbcorecr/helper"
)

func handleMsgAcceptInvite(ctx sdk.Context, k keeper.Keeper, msg *types.MsgAcceptInvite) (*sdk.Result, error) {
	invite := k.GetInvite(ctx, msg.InviteId)

	// Check if invite is not activated yet
	// Check if invitee is an owner
	if invite.Invitee != "" {
		return nil, sdkerrors.Wrap(types.ErrInvite, "Invite is already accepted")
	}
	if invite.Key != msg.TmpAddress {
		return nil, sdkerrors.Wrap(types.ErrInvite, "Wrong invite key")
	}
	// Create a new account
	newAcc, err := types.NewInviteAccount(
		msg.Address, msg.PubKey, &ctx, k.AuthKeeper,
	)
	// Set new account as invitee and remove key
	invite.AcceptanceDt = msg.AcceptanceDt

	// Add coins to the account
	coinsPack := types.IndentityCoinPacks[invite.Level]
	if len(coinsPack) > 0 {
		if err := newAcc.SetBalances(
			k.BankKeeper,
			coinsPack,
		); err != nil {
			return nil, err
		}
	}

	invite.Invitee = k.CreateIdentity(ctx, types.Identity{
		IdentityType: invite.IdentityType,
		InvitationId: invite.Id,
		CreationDt:   invite.AcceptanceDt,
	}, newAcc.Address)

	// Update invite
	k.UpdateInvite(ctx, invite)

	tmpAddr, err := sdk.AccAddressFromBech32(invite.Key)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvite, "Incorrect invite key")
	}

	// Delete temporary account
	k.AuthKeeper.RemoveAccount(
		ctx, k.AuthKeeper.GetAccount(ctx, tmpAddr),
	)

	k.TouchId(ctx, invite.Invitee, msg.AcceptanceDt)
	k.TouchId(ctx, invite.Inviter, msg.AcceptanceDt)

	// Build response
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventCreateIdentity,
			sdk.NewAttribute(
				types.EventAttrIdentityType,
				types.IdentityType_name[int32(invite.IdentityType)],
			),
			sdk.NewAttribute(
				types.EventAttrIdentityAddress,
				newAcc.Address,
			),
			sdk.NewAttribute(
				types.EventAttrIentityId,
				invite.Invitee,
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgCreateInvite(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateInvite) (*sdk.Result, error) {
	// Has invite coin of appropriate level
	inviteHelper, err := helper.NewInviteHelper(msg.Inviter, &ctx, &k)
	if err != nil {
		return nil, err
	}

	if err := inviteHelper.WithLevel(msg.Level); err != nil {
		return nil, err
	}

	// Create a temporary account
	tempAcc, err := types.NewInviteAccount(
		msg.Address, msg.PubKey, &ctx, k.AuthKeeper,
	)
	if err != nil {
		return nil, err
	}

	// Consume invite coin
	if err := inviteHelper.Pay(); err != nil {
		return nil, err
	}

	inviterId, err := k.EnsureIdFromAddress(ctx, msg.Inviter, msg.CreationDt)
	if err != nil {
		return nil, err
	}

	// Create invite entery
	invite := &types.Invite{
		Inviter:      inviterId,
		Invitee:      "",
		Level:        msg.Level,
		IdentityType: msg.IdentityType,
		Key:          tempAcc.Address,
		CreationDt:   msg.CreationDt,
	}

	k.CreateInvite(ctx, invite)
	k.TouchId(ctx, inviterId, msg.CreationDt)

	// Provide response
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventCreateInvite,
			sdk.NewAttribute(
				types.EventAttrIdentityType,
				types.IdentityType_name[int32(msg.IdentityType)],
			),
			sdk.NewAttribute(
				types.EventAttrTmpAddress,
				tempAcc.Address,
			),
			sdk.NewAttribute(
				types.EventAttrInviteId,
				invite.Id,
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
