package crsign

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
)

func handleMsgCreateId2Sign(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateId2Sign) (*sdk.Result, error) {
	k.CreateId2Sign(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateId2Sign(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateId2Sign) (*sdk.Result, error) {
	var id2sign = types.Id2Sign{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Identity: msg.Identity,
    	Signature: msg.Signature,
	}

    if msg.Creator != k.GetId2SignOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdateId2Sign(ctx, id2sign)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteId2Sign(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteId2Sign) (*sdk.Result, error) {
    if !k.HasId2Sign(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if msg.Creator != k.GetId2SignOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeleteId2Sign(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
