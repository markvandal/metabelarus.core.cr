package crconsent

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crconsent/keeper"
	"github.com/metabelarus/mbcorecr/x/crconsent/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
case *types.MsgCreateRequest:
	return handleMsgCreateRequest(ctx, k, msg)

case *types.MsgUpdateRequest:
	return handleMsgUpdateRequest(ctx, k, msg)

case *types.MsgDeleteRequest:
	return handleMsgDeleteRequest(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
