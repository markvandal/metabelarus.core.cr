package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateIdentity creates a Identity
func (k Keeper) CreateIdentity(ctx sdk.Context, Identity types.Identity) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.IdentityPrefix + Identity.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(Identity)
	store.Set(key, value)
}

// GetIdentity returns the Identity information
func (k Keeper) GetIdentity(ctx sdk.Context, key string) (types.Identity, error) {
	store := ctx.KVStore(k.storeKey)
	var Identity types.Identity
	byteKey := []byte(types.IdentityPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &Identity)
	if err != nil {
		return Identity, err
	}
	return Identity, nil
}

// SetIdentity sets a Identity
func (k Keeper) SetIdentity(ctx sdk.Context, Identity types.Identity) {
	IdentityKey := Identity.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(Identity)
	key := []byte(types.IdentityPrefix + IdentityKey)
	store.Set(key, bz)
}

// DeleteIdentity deletes a Identity
func (k Keeper) DeleteIdentity(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.IdentityPrefix + key))
}

//
// Functions used by querier
//

func listIdentity(ctx sdk.Context, k Keeper) ([]byte, error) {
	var IdentityList []types.Identity
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.IdentityPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var Identity types.Identity
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &Identity)
		IdentityList = append(IdentityList, Identity)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, IdentityList)
	return res, nil
}

func getIdentity(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	Identity, err := k.GetIdentity(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, Identity)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetIdentityOwner(ctx sdk.Context, key string) sdk.AccAddress {
	Identity, err := k.GetIdentity(ctx, key)
	if err != nil {
		return nil
	}
	return Identity.Creator
}


// Check if the key exists in the store
func (k Keeper) IdentityExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.IdentityPrefix + key))
}
