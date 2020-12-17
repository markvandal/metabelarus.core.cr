package mbpasstrust

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/keeper"
)

func handleMsgCreateAllowance(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateAllowance) (*sdk.Result, error) {
	var allowance = types.Allowance{
		Creator: msg.Creator,
		ID:      msg.ID,
    	PassportId: msg.PassportId,
    	AllowanceType: msg.AllowanceType,
    	Resolution: msg.Resolution,
	}
	k.CreateAllowance(ctx, allowance)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
