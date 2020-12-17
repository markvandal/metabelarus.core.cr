package mbgovperm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/keeper"
)

func handleMsgSetExtservice(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetExtservice) (*sdk.Result, error) {
	var extservice = types.Extservice{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Alias: msg.Alias,
    	DetailsUrl: msg.DetailsUrl,
	}
	if !msg.Creator.Equals(k.GetExtserviceOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetExtservice(ctx, extservice)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
