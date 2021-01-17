package crsign

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgRequestAuth:
			return handleMsgRequestAuth(ctx, k, msg)

		case *types.MsgConfirmAuth:
			return handleMsgConfirmAuth(ctx, k, msg)

		case *types.MsgCreateId2Sign:
			return handleMsgCreateId2Sign(ctx, k, msg)

		case *types.MsgUpdateId2Sign:
			return handleMsgUpdateId2Sign(ctx, k, msg)

		case *types.MsgDeleteId2Sign:
			return handleMsgDeleteId2Sign(ctx, k, msg)

		case *types.MsgCreateSignature:
			return handleMsgCreateSignature(ctx, k, msg)

		case *types.MsgUpdateSignature:
			return handleMsgUpdateSignature(ctx, k, msg)

		case *types.MsgDeleteSignature:
			return handleMsgDeleteSignature(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
