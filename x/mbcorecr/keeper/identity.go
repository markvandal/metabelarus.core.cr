package keeper

import (
	"encoding/hex"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

// GetIdentityCount get the total number of identity
func (k Keeper) GetIdentityCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityCountKey))
	byteKey := types.KeyPrefix(types.IdentityCountKey)
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

// SetIdentityCount set the total number of identity
func (k Keeper) SetIdentityCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityCountKey))
	byteKey := types.KeyPrefix(types.IdentityCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateIdentity(ctx sdk.Context, identity types.Identity) string {
	// Create the identity
	count := k.GetIdentityCount(ctx)
	identity.Id = strconv.FormatInt(count, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	key := types.KeyPrefix(types.IdentityKey + identity.Id)
	value := k.cdc.MustMarshalBinaryBare(&identity)
	store.Set(key, value)

	// Update identity count
	k.SetIdentityCount(ctx, count+1)

	return hex.EncodeToString(key)
}

func (k Keeper) UpdateIdentity(ctx sdk.Context, identity types.Identity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	b := k.cdc.MustMarshalBinaryBare(&identity)
	store.Set(types.KeyPrefix(types.IdentityKey+identity.Id), b)
}

func (k Keeper) GetIdentity(ctx sdk.Context, key string) types.Identity {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	var identity types.Identity
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.IdentityKey+key)), &identity)
	return identity
}

func (k Keeper) HasIdentity(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	return store.Has(types.KeyPrefix(types.IdentityKey + id))
}

func (k Keeper) GetIdentityOwner(ctx sdk.Context, key string) string {
	return k.GetIdentity(ctx, key).AccountID
}

// DeleteIdentity deletes a identity
func (k Keeper) DeleteIdentity(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	store.Delete(types.KeyPrefix(types.IdentityKey + key))
}

func (k Keeper) GetAllIdentity(ctx sdk.Context) (msgs []types.Identity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.IdentityKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Identity
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
