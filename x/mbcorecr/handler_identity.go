package mbcorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

func handleMsgUpdateIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateIdentity) (*sdk.Result, error) {
	var identity = types.Identity{
		Id:        msg.Id,
		AccountID: msg.AccountID,
		Details:   msg.Details,
	}

	if msg.AccountID != k.GetIdentityOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.UpdateIdentity(ctx, identity)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
