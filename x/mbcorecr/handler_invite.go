package mbcorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
)

func handleMsgCreateInvite(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateInvite) (*sdk.Result, error) {
	k.CreateInvite(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateInvite(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateInvite) (*sdk.Result, error) {
	var invite = types.Invite{
		Creator: msg.Creator,
		Id:      msg.Id,
    	Inviter: msg.Inviter,
    	Invitee: msg.Invitee,
    	Level: msg.Level,
    	Key: msg.Key,
    	CreationDt: msg.CreationDt,
	}

    if msg.Creator != k.GetInviteOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error                                                                                             
    }          

	k.UpdateInvite(ctx, invite)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteInvite(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteInvite) (*sdk.Result, error) {
    if !k.HasInvite(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)                                                                                                                                
    }                                                                                                                                                                                                  
    if msg.Creator != k.GetInviteOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")                                                                                                                       
    } 

	k.DeleteInvite(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
