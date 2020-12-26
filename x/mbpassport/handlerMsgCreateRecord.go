package mbpassport

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpassport/keeper"
)

func handleMsgCreateRecord(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateRecord) (*sdk.Result, error) {
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
	k.CreateRecord(ctx, record)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
