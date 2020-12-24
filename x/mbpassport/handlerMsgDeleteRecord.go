package mbpassport

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpassport/keeper"
)

// Handle a message to delete name
func handleMsgDeleteRecord(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteRecord) (*sdk.Result, error) {
	if !k.RecordExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetRecordOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteRecord(ctx, msg.ID)
	return &sdk.Result{}, nil
}
