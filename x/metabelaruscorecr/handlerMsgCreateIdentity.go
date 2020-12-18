package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
)

func handleMsgCreateIdentity(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateIdentity) (*sdk.Result, error) {
	var Identity = types.Identity{
		Creator: msg.Creator,
		ID:      msg.ID,
    	AccountID: msg.AccountID,
    	Details: msg.Details,
    	CreationDt: msg.CreationDt,
    	IdenitityType: msg.IdenitityType,
    	AuthPubKey: msg.AuthPubKey,
	}
	k.CreateIdentity(ctx, Identity)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
