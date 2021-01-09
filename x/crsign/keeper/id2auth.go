package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetId2AuthCount get the total number of id2auth
func (k Keeper) GetId2AuthCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthCountKey))
	byteKey := types.KeyPrefix(types.Id2AuthCountKey)
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

// SetId2AuthCount set the total number of id2auth
func (k Keeper) SetId2AuthCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthCountKey))
	byteKey := types.KeyPrefix(types.Id2AuthCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateId2Auth(ctx sdk.Context, msg types.  MsgCreateId2Auth) {
	// Create the id2auth
    count := k.GetId2AuthCount(ctx)
    var id2auth = types.Id2Auth{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Identity: msg.Identity,
        Auth: msg.Auth,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthKey))
    key := types.KeyPrefix(types.Id2AuthKey + id2auth.Id)
    value := k.cdc.MustMarshalBinaryBare(&id2auth)
    store.Set(key, value)

    // Update id2auth count
    k.SetId2AuthCount(ctx, count+1)
}

func (k Keeper) UpdateId2Auth(ctx sdk.Context, id2auth types.Id2Auth) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthKey))
	b := k.cdc.MustMarshalBinaryBare(&id2auth)
	store.Set(types.KeyPrefix(types.Id2AuthKey + id2auth.Id), b)
}

func (k Keeper) GetId2Auth(ctx sdk.Context, key string) types.Id2Auth {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthKey))
	var id2auth types.Id2Auth
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.Id2AuthKey + key)), &id2auth)
	return id2auth
}

func (k Keeper) HasId2Auth(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthKey))
	return store.Has(types.KeyPrefix(types.Id2AuthKey + id))
}

func (k Keeper) GetId2AuthOwner(ctx sdk.Context, key string) string {
    return k.GetId2Auth(ctx, key).Creator
}

// DeleteId2Auth deletes a id2auth
func (k Keeper) DeleteId2Auth(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthKey))
	store.Delete(types.KeyPrefix(types.Id2AuthKey + key))
}

func (k Keeper) GetAllId2Auth(ctx sdk.Context) (msgs []types.Id2Auth) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.Id2AuthKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Id2Auth
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
