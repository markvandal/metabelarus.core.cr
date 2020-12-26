package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
)

func handleMsgSetConfirmation(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetConfirmation) (*sdk.Result, error) {
	var confirmation = types.Confirmation{
		Creator: msg.Creator,
		ID:      msg.ID,
    	IdenitityID: msg.IdenitityID,
    	CreationDate: msg.CreationDate,
    	ExpirationDate: msg.ExpirationDate,
    	ConfirmatorID: msg.ConfirmatorID,
    	CenterGeo: msg.CenterGeo,
    	Status: msg.Status,
    	NextTryDate: msg.NextTryDate,
	}
	if !msg.Creator.Equals(k.GetConfirmationOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetConfirmation(ctx, confirmation)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
