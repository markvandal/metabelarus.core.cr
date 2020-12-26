package metabelaruscorecr

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
)

func handleMsgCreateInvitation(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateInvitation) (*sdk.Result, error) {
	var invitation = types.Invitation{
		Creator: msg.Creator,
		ID:      msg.ID,
    	InviterId: msg.InviterId,
    	IdentityId: msg.IdentityId,
    	CreationDate: msg.CreationDate,
    	ActivationPubKey: msg.ActivationPubKey,
	}
	k.CreateInvitation(ctx, invitation)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
