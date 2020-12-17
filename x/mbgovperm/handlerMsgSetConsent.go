package mbgovperm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/keeper"
)

func handleMsgSetConsent(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetConsent) (*sdk.Result, error) {
	var consent = types.Consent{
		Creator: msg.Creator,
		ID:      msg.ID,
    	ExtserviceId: msg.ExtserviceId,
    	PassportId: msg.PassportId,
    	Resolution: msg.Resolution,
	}
	if !msg.Creator.Equals(k.GetConsentOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetConsent(ctx, consent)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
