package crsign

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func handleMsgCreateId2Auth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateId2Auth) (*sdk.Result, error) {
	k.CreateId2Auth(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateId2Auth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateId2Auth) (*sdk.Result, error) {
	var id2auth = types.Id2Auth{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Identity: msg.Identity,
		Auth:     msg.Auth,
	}

	if msg.Creator != k.GetId2AuthOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.UpdateId2Auth(ctx, id2auth)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteId2Auth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteId2Auth) (*sdk.Result, error) {
	if !k.HasId2Auth(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}
	if msg.Creator != k.GetId2AuthOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.DeleteId2Auth(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
