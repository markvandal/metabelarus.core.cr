package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

// Such trunsaction shouldn't exist separetly - it should be part of request confirmation
func handleMsgSetIdentity(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetIdentity) (*sdk.Result, error) {
	oldIdentity, err := k.GetIdentity(ctx, msg.ID)
	if err != nil {
		return nil, err
	}

	k.SetIdentity(ctx, types.Identity{
		ID:            msg.ID,
		Details:       msg.Details,
		IdenitityType: msg.IdenitityType,
		AuthPubKey:    msg.AuthPubKey,
		CreationDt:    oldIdentity.CreationDt,
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
