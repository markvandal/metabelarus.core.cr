package crsign

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func handleMsgRequestAuth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgRequestAuth) (*sdk.Result, error) {
	if !k.IdKeeper.HasIdentity(ctx, msg.Service) {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No service identity")
	}

	if !k.IdKeeper.HasIdentity(ctx, msg.Identity) {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No user identity")
	}

	identity := k.IdKeeper.ExportIdentity(ctx, msg.Service)
	if identity.ExportAddress() != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "Incorrect service identity")
	}

	// Reset the auth object for update
	auth := &types.Auth{
		Service:        msg.Service,
		Identity:       msg.Identity,
		Key:            msg.Key,
		Status:         types.AuthStatus_AUTH_OPEN,
		CreationDt:     msg.CreationDt,
		AvailabilityDt: msg.CreationDt,
	}

	k.CreateAuth(ctx, auth)

	// Produce response
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRequestAuth,
			sdk.NewAttribute(
				types.EventAttrAuthId,
				auth.GetId(),
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

// @TODO register this handler everywhere
func handleMsgConfirmAuth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgConfirmAuth) (*sdk.Result, error) {
	if !k.IdKeeper.HasIdentity(ctx, msg.Identity) {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No user identity")
	}

	if !k.IdKeeper.HasIdentity(ctx, msg.Service) {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No service identity")
	}

	identity := k.IdKeeper.ExportIdentity(ctx, msg.Identity)
	if identity.ExportAddress() != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "Incorrect user identity")
	}

	auth := k.GetAuth(ctx, msg.Service, msg.Identity)

	auth.Status = types.AuthStatus_AUTH_SIGNED

	duration, err := time.ParseDuration("12h") // @TODO should be variable and limited
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrAuthDuration, "Incorrect duration format for auth")
	}
	newDuration := auth.AvailabilityDt.Add(duration)
	auth.AvailabilityDt = &newDuration

	k.UpdateAuth(ctx, &auth)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventConfirmAuth,
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
