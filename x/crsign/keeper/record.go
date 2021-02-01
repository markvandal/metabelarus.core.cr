package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) CreateRecord(ctx sdk.Context, msg types.Record) {
	// Create the record
	var status types.RecordStatus
	switch msg.RecordType {
	case types.RecordType_IDENTITY_PERMANENT_RECORD:
		status = types.RecordStatus_RECORD_SEALED
	default:
		status = types.RecordStatus_RECORD_OPEN
	}

	incrementRecord := false
	var id string
	if msg.Id != "" && msg.RecordType == types.RecordType_IDENTITY_MUTUAL_RECORD {
		id = msg.Id + "." + msg.Provider
	} else {
		id = strconv.FormatInt(k.GetRecordCount(ctx), 10)
		incrementRecord = true
	}

	// TODO Check signature for ownership

	count := k.GetRecordCount(ctx)
	var record = types.Record{
		Id:          id,
		Identity:    msg.Identity,
		Provider:    msg.Provider,
		Key:         msg.Key,
		Data:        msg.Data,
		Signature:   msg.Signature,
		RecordType:  msg.RecordType,
		Publicity:   msg.Publicity,
		Status:      status,
		LiveTime:    msg.LiveTime,
		CreationDt:  msg.CreationDt,
		SignatureDt: msg.CreationDt,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	key := types.KeyPrefix(types.RecordKey + record.Id)
	value := k.cdc.MustMarshalBinaryBare(&record)
	store.Set(key, value)

	// Update record count
	if incrementRecord {
		k.SetRecordCount(ctx, count+1)
	}
}

func (k Keeper) UpdateRecord(ctx sdk.Context, record types.Record) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	b := k.cdc.MustMarshalBinaryBare(&record)
	store.Set(types.KeyPrefix(types.RecordKey+record.Id), b)
}

func (k Keeper) GetRecord(ctx sdk.Context, id string) types.Record {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	var record types.Record
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.RecordKey+id)), &record)
	return record
}

func (k Keeper) HasRecord(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	return store.Has(types.KeyPrefix(types.RecordKey + id))
}

func (k Keeper) GetRecordOwner(ctx sdk.Context, key string) string {
	record := k.GetRecord(ctx, key)
	switch record.RecordType {
	case types.RecordType_PROVIDER_PERMISSION:
	case types.RecordType_PROVIDER_RECORD:
	case types.RecordType_PROVIDER_SIGNABLE_RECORD:
		return record.Provider
	}
	return record.Identity
}

// DeleteRecord deletes a signature
func (k Keeper) DeleteRecord(ctx sdk.Context, key string) {
	record := k.GetRecord(ctx, key)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	store.Delete(types.KeyPrefix(types.RecordKey + key))
	if parent := record.GetParentId(); "" != parent {
		k.DeleteRecord(ctx, parent)
	} else if child := record.GetChildId(); "" != child {
		k.DeleteRecord(ctx, child)
	}
}
