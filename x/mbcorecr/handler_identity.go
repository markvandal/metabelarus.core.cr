package mbcorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

func handleMsgCreateIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateIdentity) (*sdk.Result, error) {
	k.CreateIdentity(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateIdentity) (*sdk.Result, error) {
	var identity = types.Identity{
		Creator:      msg.Creator,
		Id:           msg.Id,
		AccountID:    msg.AccountID,
		IdentityType: msg.IdentityType,
		Details:      msg.Details,
	}

	if msg.Creator != k.GetIdentityOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.UpdateIdentity(ctx, identity)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteIdentity) (*sdk.Result, error) {
	if !k.HasIdentity(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}
	if msg.Creator != k.GetIdentityOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.DeleteIdentity(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
