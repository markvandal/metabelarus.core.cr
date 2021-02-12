package crsign

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	corecrtypes "github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

func handleMsgRequestAuth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgRequestAuth) (*sdk.Result, error) {
	serviceId := k.IdKeeper.GetIdFromAddress(ctx, msg.Service)
	if serviceId == "" {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No service identity")
	}

	service := k.IdKeeper.ExportIdentity(ctx, serviceId)
	if !service.VerifyIdentityType(corecrtypes.IdentityType_SERVICE) {
		return nil, sdkerrors.Wrap(types.ErrIdNotService, "Is not a service")
	}

	if !k.IdKeeper.HasIdentity(ctx, msg.Identity) {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No user identity")
	}

	// Reset the auth object for update
	auth := &types.Auth{
		Service:        serviceId,
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

func handleMsgConfirmAuth(ctx sdk.Context, k keeper.Keeper, msg *types.MsgConfirmAuth) (*sdk.Result, error) {
	identityId := k.IdKeeper.GetIdFromAddress(ctx, msg.Identity)
	if identityId == "" {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No user identity")
	}

	if !k.IdKeeper.HasIdentity(ctx, msg.Service) {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No service identity")
	}

	auth := k.GetAuth(ctx, msg.Service, identityId)
	auth.Status = types.AuthStatus_AUTH_SIGNED

	duration, err := time.ParseDuration(types.DefaultAuthLifeTime) // @TODO should be variable and limited
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
