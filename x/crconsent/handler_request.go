package crconsent

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crconsent/keeper"
	"github.com/metabelarus/mbcorecr/x/crconsent/types"
)

func handleMsgCreateRequest(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateRequest) (*sdk.Result, error) {
	k.CreateRequest(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateRequest(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateRequest) (*sdk.Result, error) {
	var request = types.Request{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Initiator:   msg.Initiator,
		Recipient:   msg.Recipient,
		RequestType: msg.RequestType,
		Status:      msg.Status,
		Value:       msg.Value,
		Memo:        msg.Memo,
		PromoUrl:    msg.PromoUrl,
		CreationDt:  msg.CreationDt,
		FinalDt:     msg.FinalDt,
	}

	if msg.Creator != k.GetRequestOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.UpdateRequest(ctx, request)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteRequest(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteRequest) (*sdk.Result, error) {
	if !k.HasRequest(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}
	if msg.Creator != k.GetRequestOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.DeleteRequest(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
