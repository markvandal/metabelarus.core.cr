package crsign

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func handleMsgCreateSignature(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateSignature) (*sdk.Result, error) {
	k.CreateSignature(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateSignature(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateSignature) (*sdk.Result, error) {
	var signature = types.Signature{
		Id:       msg.Id,
		Identity: msg.Identity,
		Service:  msg.Service,
		// Key: msg.Key,
		// Secret: msg.Secret,
		// CreationDt: msg.CreationDt,
		// AvailabilityDt: msg.AvailabilityDt,
	}

	if msg.Creator != k.GetSignatureOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.UpdateSignature(ctx, signature)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteSignature(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteSignature) (*sdk.Result, error) {
	if !k.HasSignature(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}
	if msg.Creator != k.GetSignatureOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.DeleteSignature(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
