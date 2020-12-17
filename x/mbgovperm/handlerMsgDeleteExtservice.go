package mbgovperm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/keeper"
)

// Handle a message to delete name
func handleMsgDeleteExtservice(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteExtservice) (*sdk.Result, error) {
	if !k.ExtserviceExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetExtserviceOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteExtservice(ctx, msg.ID)
	return &sdk.Result{}, nil
}
