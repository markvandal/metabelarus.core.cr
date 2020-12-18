package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
)

func handleMsgSetIdentity(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetIdentity) (*sdk.Result, error) {
	var Identity = types.Identity{
		Creator: msg.Creator,
		ID:      msg.ID,
    	AccountID: msg.AccountID,
    	Details: msg.Details,
    	CreationDt: msg.CreationDt,
    	IdenitityType: msg.IdenitityType,
    	AuthPubKey: msg.AuthPubKey,
	}
	if !msg.Creator.Equals(k.GetIdentityOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetIdentity(ctx, Identity)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
