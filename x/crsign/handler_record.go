package crsign

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	corecrtypes "github.com/metabelarus/mbcorecr/x/mbcorecr/types"

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

	if types.IsIdentityRecord(msg.RecordType) && (creatorIdentityId != identityId) {
		return nil, sdkerrors.Wrap(
			types.ErrRecIdentityProvider,
			fmt.Sprintf("Provided %s: %s %s", msg.RecordType.String(), msg.Provider, identityId),
		)
	}

	if types.IsProviderRecord(msg.RecordType) && (creatorIdentityId == identityId) {
		service := k.IdKeeper.ExportIdentity(ctx, creatorIdentityId)
		if !service.VerifyIdentityType(corecrtypes.IdentityType_SERVICE) {
			return nil, sdkerrors.Wrap(
				types.ErrRecIdentityProvider,
				fmt.Sprintf("Provided %s: %s %s", msg.RecordType.String(), msg.Provider, identityId),
			)
		}
	}

	record, err := k.CreateRecord(ctx, msg.ToRecord(identityId))
	if err != nil {
		return nil, err
	}
	k.IdKeeper.TouchId(ctx, creatorIdentityId, msg.CreationDt)

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

	updater, err := update.CreateUpdateStatus(record, updaterIdentityId)
	if err != nil {
		return nil, err
	}
	if err = updater.Dispatch(msg); err != nil {
		return nil, err
	}

	var mutualRecord *types.Record
	if updater.IsMutualUpdateRequired() {
		mutualMsg := &types.MsgUpdateRecord{
			LiveTime: msg.LiveTime,
			Action:   msg.Action,
			UpdateDt: msg.UpdateDt,
		}
		var mutualIdentity string
		if updater.IsChildUpdate() {
			mutualMsg.Id = record.GetParentId()
			if record.RecordType == types.RecordType_IDENTITY_RECORD {
				mutualIdentity = record.Provider
			} else {
				mutualIdentity = record.Identity
			}
		} else {
			mutualMsg.Id = record.GetChildId()
			if record.RecordType == types.RecordType_IDENTITY_MUTUAL_RECORD {
				mutualIdentity = record.Provider
			} else {
				mutualIdentity = record.Identity
			}
		}
		mutualMsg.Updater = k.IdKeeper.GetAddressFromId(ctx, mutualIdentity)
		mutualRecord = k.GetRecord(ctx, mutualMsg.Id)
		mutualUpdater, err := update.CreateUpdateStatus(mutualRecord, mutualIdentity)
		if err != nil {
			return nil, err
		}
		if err = mutualUpdater.Dispatch(mutualMsg); err != nil {
			return nil, err
		}
	}

	k.UpdateRecord(ctx, record)
	if mutualRecord != nil {
		k.UpdateRecord(ctx, mutualRecord)
	}

	k.IdKeeper.TouchId(ctx, record.Identity, msg.UpdateDt)
	if record.Identity != record.Provider {
		k.IdKeeper.TouchId(ctx, record.Provider, msg.UpdateDt)
	}

	// Produce response
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventUpdateRecord,
			sdk.NewAttribute(
				types.EventAttrRecordId,
				record.GetId(),
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteRecord(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteRecord) (*sdk.Result, error) {
	deleterIdentityId := k.IdKeeper.GetIdFromAddress(ctx, msg.Deleter)
	if deleterIdentityId == "" {
		return nil, sdkerrors.Wrap(types.ErrNoIdentity, "No user identity")
	}
	if !k.HasRecord(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}

	record := k.GetRecord(ctx, msg.Id)
	switch record.Status {
	default:
		return nil, sdkerrors.Wrap(
			types.ErrDelete,
			fmt.Sprintf(
				"Can't delete record from %s status",
				record.Status.String(),
			),
		)
	case types.RecordStatus_RECORD_OPEN:
	case types.RecordStatus_RECORD_WITHDRAWN:
	case types.RecordStatus_RECORD_REJECTED:
	}

	updater, err := update.CreateUpdateStatus(record, deleterIdentityId)
	if err != nil {
		return nil, err
	}

	err = updater.CheckUpdate()
	if err != nil {
		return nil, err
	}
	updater.RequireMutualUpdate()

	if updater.IsMutualUpdateRequired() {
		if updater.IsChildUpdate() {
			k.DeleteRecord(ctx, record.GetParentId())
		} else {
			k.DeleteRecord(ctx, record.GetChildId())
		}
	}

	k.DeleteRecord(ctx, msg.Id)

	// Produce response
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventDeleteRecord,
			sdk.NewAttribute(
				types.EventAttrRecordId,
				record.GetId(),
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
