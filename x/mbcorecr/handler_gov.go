package mbcorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	"github.com/metabelarus/mbcorecr/x/mbcorecr/helper"
)

func handleMsgCreateSuperIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateSuperIdentity) (*sdk.Result, error) {
	inviteHelper, err := helper.NewInviteHelper(msg.Creator, &ctx, &k.BankKeeper, &k.AuthKeeper)
	if err != nil {
		return nil, err
	}

	if err := inviteHelper.WithDenom(types.SuperInviteDenom); err != nil {
		return nil, err
	}

	inviteAcc, err := helper.NewInviteAccount(msg.Uid, ctx, k.AuthKeeper)
	if err != nil {
		return nil, err
	}

	if err := inviteHelper.Pay(); err != nil {
		return nil, err
	}

	if err := inviteAcc.SetBalances(
		ctx, k.BankKeeper,
		types.SuperIdentityCoinsPack,
	); err != nil {
		return nil, err
	}

	datails, err := inviteAcc.EncryptStr("{}")
	if err != nil {
		return nil, err
	}

	// Create an identity
	identityKey := k.CreateIdentity(ctx, types.Identity{
		AccountID:    inviteAcc.PubKey,
		IdentityType: types.IdentityType_CITIZEN,
		Details:      datails,
		InvitationId: "",
		CreationDt:   msg.CreationDt,
	})

	// [TODO] Probably we need to add 1.000.000 simple tokens
	// [TODO] Probably we need to add 1.000 stake

	resp := types.IdentityAccount{
		Uid:        inviteAcc.Uid,
		Address:    inviteAcc.Address,
		Mnemonic:   inviteAcc.Mnemonic,
		PubKey:     inviteAcc.PubKey,
		PrivKey:    inviteAcc.PrivKey,
		IdentityID: identityKey,
	}

	ecnryptedPayload, err := inviteAcc.EncryptData(resp)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventGovCreateIdentity,
			sdk.NewAttribute(
				types.EventAttrIdentityType,
				string(types.AttrIdentityTypeSuper),
			),
			sdk.NewAttribute(
				types.EventAttrIdentityPayload,
				ecnryptedPayload,
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
