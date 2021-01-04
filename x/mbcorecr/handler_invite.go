package mbcorecr

import (
	"encoding/base64"

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
	if invite.Key != msg.Invitee {
		return nil, sdkerrors.Wrap(types.ErrInvite, "Invite is already accepted")
	}
	// Create a new account
	newAcc, err := helper.NewInviteAccount(msg.Uid, ctx, k.AuthKeeper)
	// Set new account as invitee and remove key
	invite.Invitee = newAcc.Address
	invite.Key = ""
	invite.AcceptanceDt = msg.AcceptanceDt

	// Delete temporary account
	err = helper.DeleteAccount(ctx, k.AuthKeeper, msg.Invitee)
	if err != nil {
		return nil, err
	}

	// Add coins to the account
	coinsPack := types.IndentityCoinPacks[invite.Level]
	if len(coinsPack) > 0 {
		if err := newAcc.SetBalances(
			ctx, k.BankKeeper,
			coinsPack,
		); err != nil {
			return nil, err
		}
	}
	// Create identity and backreference identity to invite
	datails, err := newAcc.EncryptStr("{}")
	if err != nil {
		return nil, err
	}

	invite.IdentityId = k.CreateIdentity(ctx, types.Identity{
		AccountID:    newAcc.Address,
		IdentityType: invite.IdentityType,
		Details:      datails,
		InvitationId: invite.Id,
		CreationDt:   invite.AcceptanceDt,
	})

	// Update invite
	k.UpdateInvite(ctx, invite)

	// Build response
	newAcc.InviteID = invite.IdentityId

	ecnryptedPayload, err := newAcc.EncryptData(nil)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventCreateIdentity,
			sdk.NewAttribute(
				types.EventAttrIdentityType,
				types.IdentityType_name[int32(invite.IdentityType)],
			),
			sdk.NewAttribute(
				types.EventAttrIdentityPayload,
				ecnryptedPayload,
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

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
	invite := &types.Invite{
		Inviter:      msg.Inviter,
		Invitee:      "",
		Level:        msg.Level,
		IdentityType: msg.IdentityType,
		Key:          tempAcc.GetAddress(),
		CreationDt:   msg.CreationDt,
	}

	k.CreateInvite(ctx, invite)

	// Provide response
	inviteAcc := types.InviteAccount{
		Uid:      tempAcc.Uid,
		Address:  tempAcc.Address,
		Mnemonic: base64.URLEncoding.EncodeToString([]byte(tempAcc.Mnemonic)),
		PubKey:   tempAcc.PubKey,
		PrivKey:  tempAcc.PrivKey,
		InviteID: invite.Id,
	}

	ecnryptedPayload, err := inviteHelper.EncryptData(inviteAcc)
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
