package mbpassport

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpassport/keeper"
)

func handleMsgSetRecord(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetRecord) (*sdk.Result, error) {
	var record = types.Record{
		Creator: msg.Creator,
		ID:      msg.ID,
    	IdentityId: msg.IdentityId,
    	ServiceId: msg.ServiceId,
    	ServiceType: msg.ServiceType,
    	Key: msg.Key,
    	UserValue: msg.UserValue,
    	ServiceValue: msg.ServiceValue,
    	CreationDt: msg.CreationDt,
    	UpdateDt: msg.UpdateDt,
	}
	if !msg.Creator.Equals(k.GetRecordOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetRecord(ctx, record)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
