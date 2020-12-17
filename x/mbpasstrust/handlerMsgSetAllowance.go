package mbpasstrust

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/keeper"
)

func handleMsgSetAllowance(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetAllowance) (*sdk.Result, error) {
	var allowance = types.Allowance{
		Creator: msg.Creator,
		ID:      msg.ID,
    	PassportId: msg.PassportId,
    	AllowanceType: msg.AllowanceType,
    	Resolution: msg.Resolution,
	}
	if !msg.Creator.Equals(k.GetAllowanceOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetAllowance(ctx, allowance)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
