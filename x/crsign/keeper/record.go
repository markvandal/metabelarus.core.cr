package keeper

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// GetRecordCount get the total number of signature
func (k Keeper) GetRecordCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordCountKey))
	byteKey := types.KeyPrefix(types.RecordCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetRecordCount set the total number of signature
func (k Keeper) SetRecordCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordCountKey))
	byteKey := types.KeyPrefix(types.RecordCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateRecord(ctx sdk.Context, msg *types.Record) (*types.Record, error) {
	count := k.GetRecordCount(ctx)
	id := strconv.FormatInt(count, 10)

	var recordKey types.Id2KeyRecord
	keyStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2KeyRecordKey))
	id2KeyRecordKey := types.KeyPrefix(types.Id2KeyRecordKey + msg.Identity + "-" + msg.Key)
	k.cdc.MustUnmarshalBinaryBare(keyStore.Get(id2KeyRecordKey), &recordKey)
	if recordKey.Id != "" {
		return nil, sdkerrors.Wrap(
			types.ErrRecordExists,
			fmt.Sprintf("Record with key %s already exist in identity %s", msg.Key, msg.Identity),
		)
	}
	recordKey = types.Id2KeyRecord{Id: id}

	// TODO Check signature for ownership
	record := &types.Record{
		Id:          id,
		Identity:    msg.Identity,
		Provider:    msg.Provider,
		Key:         msg.Key,
		Data:        msg.Data,
		Signature:   msg.Signature,
		RecordType:  msg.RecordType,
		Publicity:   msg.Publicity,
		Status:      types.RecordStatus_RECORD_OPEN,
		LiveTime:    msg.LiveTime,
		CreationDt:  msg.CreationDt,
		SignatureDt: msg.CreationDt,
		UpdateDt:    msg.CreationDt,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	key := types.KeyPrefix(types.RecordKey + record.Id)
	value := k.cdc.MustMarshalBinaryBare(record)
	store.Set(key, value)

	keyStore.Set(id2KeyRecordKey, k.cdc.MustMarshalBinaryBare(&recordKey))

	// Update record count
	k.SetRecordCount(ctx, count+1)

	err := k.AddRecord2Id(ctx, msg.Identity, id)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (k Keeper) UpdateRecord(ctx sdk.Context, record *types.Record) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	b := k.cdc.MustMarshalBinaryBare(record)
	store.Set(types.KeyPrefix(types.RecordKey+record.Id), b)
}

func (k Keeper) GetRecord(ctx sdk.Context, id string) *types.Record {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	var record types.Record
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.RecordKey+id)), &record)
	return &record
}

func (k Keeper) HasRecord(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	return store.Has(types.KeyPrefix(types.RecordKey + id))
}

func (k Keeper) GetRecordOwner(ctx sdk.Context, key string) string {
	record := k.GetRecord(ctx, key)
	switch record.RecordType {
	case types.RecordType_PROVIDER_PERMISSION:
		return record.Provider
	case types.RecordType_PROVIDER_RECORD:
		return record.Provider
	case types.RecordType_PROVIDER_SIGNABLE_RECORD:
		return record.Provider
	}
	return record.Identity
}

// DeleteRecord deletes a record
func (k Keeper) DeleteRecord(ctx sdk.Context, key string) {
	record := k.GetRecord(ctx, key)
	k.DeleteRecordFromId(ctx, record.Identity, record.Id)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	store.Delete(types.KeyPrefix(types.RecordKey + key))
}
