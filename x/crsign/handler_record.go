package crsign

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"

	update "github.com/metabelarus/mbcorecr/x/crsign/types/record_update"
)

func handleMsgCreateRecord(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateRecord) (*sdk.Result, error) {
	creatorIdentityId, err := k.IdKeeper.EnsureIdFromAddress(ctx, msg.Creator, msg.CreationDt)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "Can't initialize a user identity")
	}
	if creatorIdentityId == "" {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No user identity")
	}

	if msg.Provider == "" {
		msg.Provider = creatorIdentityId
	}

	var identityId string
	if types.IsProviderRecord(msg.RecordType) {
		identityId = msg.Provider
		msg.Provider = creatorIdentityId
	} else {
		identityId = creatorIdentityId
	}

	if types.IsIdentityRecord(msg.RecordType) != (msg.Provider == identityId) {
		return nil, sdkerrors.Wrap(
			types.ErrRecIdentityProvider,
			fmt.Sprintf("Provided %s: %s %s", msg.RecordType.String(), msg.Provider, identityId),
		)
	}

	record := k.CreateRecord(ctx, msg.ToRecord(identityId))

	// Produce response
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventCreateRecord,
			sdk.NewAttribute(
				types.EventAttrRecordId,
				record.GetId(),
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateRecord(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateRecord) (*sdk.Result, error) {
	updaterIdentityId, err := k.IdKeeper.EnsureIdFromAddress(ctx, msg.Updater, msg.UpdateDt)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "Can't initialize a user identity")
	}
	if updaterIdentityId == "" {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No user identity")
	}

	record := k.GetRecord(ctx, msg.Id)
	if record.Id == "" {
		return nil, sdkerrors.Wrap(types.ErrNoRecord, "No record")
	}

	update, err := update.CreateUpdateStatus(record, updaterIdentityId)
	if err != nil {
		return nil, err
	}
	if err = update.Dispatch(msg); err != nil {
		return nil, err
	}

	k.UpdateRecord(ctx, record)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteRecord(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteRecord) (*sdk.Result, error) {
	if !k.HasRecord(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}
	if msg.Creator != k.GetRecordOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.DeleteRecord(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
