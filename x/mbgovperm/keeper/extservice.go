package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateExtservice creates a extservice
func (k Keeper) CreateExtservice(ctx sdk.Context, extservice types.Extservice) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ExtservicePrefix + extservice.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(extservice)
	store.Set(key, value)
}

// GetExtservice returns the extservice information
func (k Keeper) GetExtservice(ctx sdk.Context, key string) (types.Extservice, error) {
	store := ctx.KVStore(k.storeKey)
	var extservice types.Extservice
	byteKey := []byte(types.ExtservicePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &extservice)
	if err != nil {
		return extservice, err
	}
	return extservice, nil
}

// SetExtservice sets a extservice
func (k Keeper) SetExtservice(ctx sdk.Context, extservice types.Extservice) {
	extserviceKey := extservice.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(extservice)
	key := []byte(types.ExtservicePrefix + extserviceKey)
	store.Set(key, bz)
}

// DeleteExtservice deletes a extservice
func (k Keeper) DeleteExtservice(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ExtservicePrefix + key))
}

//
// Functions used by querier
//

func listExtservice(ctx sdk.Context, k Keeper) ([]byte, error) {
	var extserviceList []types.Extservice
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ExtservicePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var extservice types.Extservice
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &extservice)
		extserviceList = append(extserviceList, extservice)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, extserviceList)
	return res, nil
}

func getExtservice(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	extservice, err := k.GetExtservice(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, extservice)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetExtserviceOwner(ctx sdk.Context, key string) sdk.AccAddress {
	extservice, err := k.GetExtservice(ctx, key)
	if err != nil {
		return nil
	}
	return extservice.Creator
}


// Check if the key exists in the store
func (k Keeper) ExtserviceExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ExtservicePrefix + key))
}
