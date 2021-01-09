package crsign

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func handleMsgCreateAuth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateAuth) (*sdk.Result, error) {
	k.CreateAuth(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateAuth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateAuth) (*sdk.Result, error) {
	var auth = types.Auth{
		Id:       msg.Id,
		Identity: msg.Identity,
		Service:  msg.Service,
		Key:      msg.Key,
		// Status: msg.Status,
		// CreationDt: msg.CreationDt,
		// AvailabilityDt: msg.AvailabilityDt,
	}

	if msg.Creator != k.GetAuthOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.UpdateAuth(ctx, auth)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteAuth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteAuth) (*sdk.Result, error) {
	if !k.HasAuth(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}
	if msg.Creator != k.GetAuthOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.DeleteAuth(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
