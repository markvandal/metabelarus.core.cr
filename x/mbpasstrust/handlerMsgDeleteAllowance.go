package mbpasstrust

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/keeper"
)

// Handle a message to delete name
func handleMsgDeleteAllowance(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteAllowance) (*sdk.Result, error) {
	if !k.AllowanceExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetAllowanceOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteAllowance(ctx, msg.ID)
	return &sdk.Result{}, nil
}
