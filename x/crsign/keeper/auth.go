package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func (k Keeper) CreateAuth(ctx sdk.Context, auth *types.Auth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	key := types.KeyPrefix(types.AuthKey + auth.GetId())
	value := k.cdc.MustMarshalBinaryBare(auth)
	store.Set(key, value)
	k.AddService2Id(ctx, auth.Identity, auth.Service)
}

func (k Keeper) UpdateAuth(ctx sdk.Context, auth *types.Auth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	b := k.cdc.MustMarshalBinaryBare(auth)
	store.Set(types.KeyPrefix(types.AuthKey+auth.GetId()), b)
}

func (k Keeper) GetAuthByKey(ctx sdk.Context, key string) types.Auth {
	auth := &types.Auth{}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AuthKey+key)), auth)

	return *auth
}

func (k Keeper) GetAuth(ctx sdk.Context, service string, identity string) types.Auth {
	auth := &types.Auth{
		Identity: identity,
		Service:  service,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AuthKey+auth.GetId())), auth)

	return *auth
}

func (k Keeper) HasAuth(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	return store.Has(types.KeyPrefix(types.AuthKey + id))
}

// DeleteAuth deletes a auth
func (k Keeper) DeleteAuth(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	store.Delete(types.KeyPrefix(types.AuthKey + key))
}

func (k Keeper) GetAllAuth(ctx sdk.Context) (msgs []types.Auth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AuthKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Auth
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
