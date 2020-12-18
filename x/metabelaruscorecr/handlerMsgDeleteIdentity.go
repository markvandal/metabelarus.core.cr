package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
)

// Handle a message to delete name
func handleMsgDeleteIdentity(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteIdentity) (*sdk.Result, error) {
	if !k.IdentityExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetIdentityOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteIdentity(ctx, msg.ID)
	return &sdk.Result{}, nil
}
