package metabelaruscorecr

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/keeper"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateConfirmation:
			return handleMsgCreateConfirmation(ctx, k, msg)
		case types.MsgSetConfirmation:
			return handleMsgSetConfirmation(ctx, k, msg)
		case types.MsgDeleteConfirmation:
			return handleMsgDeleteConfirmation(ctx, k, msg)
		case types.MsgCreateInvitation:
			return handleMsgCreateInvitation(ctx, k, msg)
		case types.MsgSetInvitation:
			return handleMsgSetInvitation(ctx, k, msg)
		case types.MsgDeleteInvitation:
			return handleMsgDeleteInvitation(ctx, k, msg)
		case types.MsgCreateIdentity:
			return handleMsgCreateIdentity(ctx, k, msg)
		case types.MsgSetIdentity:
			return handleMsgSetIdentity(ctx, k, msg)
		case types.MsgDeleteIdentity:
			return handleMsgDeleteIdentity(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
