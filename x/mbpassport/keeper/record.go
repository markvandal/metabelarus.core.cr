package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateRecord creates a record
func (k Keeper) CreateRecord(ctx sdk.Context, record types.Record) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.RecordPrefix + record.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(record)
	store.Set(key, value)
}

// GetRecord returns the record information
func (k Keeper) GetRecord(ctx sdk.Context, key string) (types.Record, error) {
	store := ctx.KVStore(k.storeKey)
	var record types.Record
	byteKey := []byte(types.RecordPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &record)
	if err != nil {
		return record, err
	}
	return record, nil
}

// SetRecord sets a record
func (k Keeper) SetRecord(ctx sdk.Context, record types.Record) {
	recordKey := record.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(record)
	key := []byte(types.RecordPrefix + recordKey)
	store.Set(key, bz)
}

// DeleteRecord deletes a record
func (k Keeper) DeleteRecord(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.RecordPrefix + key))
}

//
// Functions used by querier
//

func listRecord(ctx sdk.Context, k Keeper) ([]byte, error) {
	var recordList []types.Record
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.RecordPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var record types.Record
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &record)
		recordList = append(recordList, record)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, recordList)
	return res, nil
}

func getRecord(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	record, err := k.GetRecord(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, record)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetRecordOwner(ctx sdk.Context, key string) sdk.AccAddress {
	record, err := k.GetRecord(ctx, key)
	if err != nil {
		return nil
	}
	return record.Creator
}


// Check if the key exists in the store
func (k Keeper) RecordExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.RecordPrefix + key))
}
