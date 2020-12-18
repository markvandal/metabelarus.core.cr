package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
)

func handleMsgSetInvitation(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetInvitation) (*sdk.Result, error) {
	var invitation = types.Invitation{
		Creator: msg.Creator,
		ID:      msg.ID,
    	InviterId: msg.InviterId,
    	IdentityId: msg.IdentityId,
    	CreationDate: msg.CreationDate,
    	ActivationPubKey: msg.ActivationPubKey,
	}
	if !msg.Creator.Equals(k.GetInvitationOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetInvitation(ctx, invitation)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
