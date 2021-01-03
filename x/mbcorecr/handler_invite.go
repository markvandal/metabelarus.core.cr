package mbcorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

func handleMsgCreateInvite(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateInvite) (*sdk.Result, error) {
	// k.CreateInvite(ctx, *msg)

	// Has invite coin of appropriate level

	// Create a temporary account 

	// Consume invite coin

	// Create invite entery

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
