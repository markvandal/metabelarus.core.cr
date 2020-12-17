package mbgovperm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/keeper"
)

// Handle a message to delete name
func handleMsgDeleteConsent(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteConsent) (*sdk.Result, error) {
	if !k.ConsentExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetConsentOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteConsent(ctx, msg.ID)
	return &sdk.Result{}, nil
}
