package mbgovperm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/keeper"
)

func handleMsgCreateConsent(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateConsent) (*sdk.Result, error) {
	var consent = types.Consent{
		Creator: msg.Creator,
		ID:      msg.ID,
    	ExtserviceId: msg.ExtserviceId,
    	PassportId: msg.PassportId,
    	Resolution: msg.Resolution,
	}
	k.CreateConsent(ctx, consent)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
