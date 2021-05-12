package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crconsent/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetRequestCount get the total number of request
func (k Keeper) GetRequestCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKey))
	byteKey := types.KeyPrefix(types.RequestCountKey)
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

// SetRequestCount set the total number of request
func (k Keeper) SetRequestCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKey))
	byteKey := types.KeyPrefix(types.RequestCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateRequest(ctx sdk.Context, msg types.MsgCreateRequest) {
	// Create the request
    count := k.GetRequestCount(ctx)
    var request = types.Request{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Initiator: msg.Initiator,
        Recipient: msg.Recipient,
        RequestType: msg.RequestType,
        Status: msg.Status,
        Value: msg.Value,
        Memo: msg.Memo,
        PromoUrl: msg.PromoUrl,
        CreationDt: msg.CreationDt,
        FinalDt: msg.FinalDt,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
    key := types.KeyPrefix(types.RequestKey + request.Id)
    value := k.cdc.MustMarshalBinaryBare(&request)
    store.Set(key, value)

    // Update request count
    k.SetRequestCount(ctx, count+1)
}

func (k Keeper) UpdateRequest(ctx sdk.Context, request types.Request) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	b := k.cdc.MustMarshalBinaryBare(&request)
	store.Set(types.KeyPrefix(types.RequestKey + request.Id), b)
}

func (k Keeper) GetRequest(ctx sdk.Context, key string) types.Request {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	var request types.Request
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.RequestKey + key)), &request)
	return request
}

func (k Keeper) HasRequest(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	return store.Has(types.KeyPrefix(types.RequestKey + id))
}

func (k Keeper) GetRequestOwner(ctx sdk.Context, key string) string {
    return k.GetRequest(ctx, key).Creator
}

func (k Keeper) GetAllRequest(ctx sdk.Context) (msgs []types.Request) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.RequestKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Request
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
