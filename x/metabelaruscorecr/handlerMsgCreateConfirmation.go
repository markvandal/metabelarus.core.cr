package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
)

func handleMsgCreateConfirmation(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateConfirmation) (*sdk.Result, error) {
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
	k.CreateConfirmation(ctx, confirmation)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
