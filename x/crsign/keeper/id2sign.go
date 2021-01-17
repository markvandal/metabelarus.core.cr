package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetId2SignCount get the total number of id2sign
func (k Keeper) GetId2SignCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignCountKey))
	byteKey := types.KeyPrefix(types.Id2SignCountKey)
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

// SetId2SignCount set the total number of id2sign
func (k Keeper) SetId2SignCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignCountKey))
	byteKey := types.KeyPrefix(types.Id2SignCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateId2Sign(ctx sdk.Context, msg types.MsgCreateId2Sign) {
	// Create the id2sign
    count := k.GetId2SignCount(ctx)
    var id2sign = types.Id2Sign{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Identity: msg.Identity,
        Signature: msg.Signature,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignKey))
    key := types.KeyPrefix(types.Id2SignKey + id2sign.Id)
    value := k.cdc.MustMarshalBinaryBare(&id2sign)
    store.Set(key, value)

    // Update id2sign count
    k.SetId2SignCount(ctx, count+1)
}

func (k Keeper) UpdateId2Sign(ctx sdk.Context, id2sign types.Id2Sign) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignKey))
	b := k.cdc.MustMarshalBinaryBare(&id2sign)
	store.Set(types.KeyPrefix(types.Id2SignKey + id2sign.Id), b)
}

func (k Keeper) GetId2Sign(ctx sdk.Context, key string) types.Id2Sign {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignKey))
	var id2sign types.Id2Sign
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.Id2SignKey + key)), &id2sign)
	return id2sign
}

func (k Keeper) HasId2Sign(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignKey))
	return store.Has(types.KeyPrefix(types.Id2SignKey + id))
}

func (k Keeper) GetId2SignOwner(ctx sdk.Context, key string) string {
    return k.GetId2Sign(ctx, key).Creator
}

// DeleteId2Sign deletes a id2sign
func (k Keeper) DeleteId2Sign(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignKey))
	store.Delete(types.KeyPrefix(types.Id2SignKey + key))
}

func (k Keeper) GetAllId2Sign(ctx sdk.Context) (msgs []types.Id2Sign) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.Id2SignKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Id2Sign
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
