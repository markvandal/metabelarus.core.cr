package mbgovperm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/keeper"
)

func handleMsgCreateExtservice(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateExtservice) (*sdk.Result, error) {
	var extservice = types.Extservice{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Alias: msg.Alias,
    	DetailsUrl: msg.DetailsUrl,
	}
	k.CreateExtservice(ctx, extservice)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
