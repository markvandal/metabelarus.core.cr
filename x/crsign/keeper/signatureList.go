package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// GetSignatureListCount get the total number of signatureList
func (k Keeper) GetSignatureListCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListCountKey))
	byteKey := types.KeyPrefix(types.SignatureListCountKey)
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

// SetSignatureListCount set the total number of signatureList
func (k Keeper) SetSignatureListCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListCountKey))
	byteKey := types.KeyPrefix(types.SignatureListCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateSignatureList(ctx sdk.Context, signatureList *types.SignatureList) {
	// Create the signatureList
	count := k.GetSignatureListCount(ctx)
	signatureList.Id = strconv.FormatInt(count, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListKey))
	key := types.KeyPrefix(types.SignatureListKey + signatureList.Id)
	value := k.cdc.MustMarshalBinaryBare(signatureList)
	store.Set(key, value)

	// Update signatureList count
	k.SetSignatureListCount(ctx, count+1)
}

func (k Keeper) UpdateSignatureList(ctx sdk.Context, signatureList types.SignatureList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListKey))
	b := k.cdc.MustMarshalBinaryBare(&signatureList)
	store.Set(types.KeyPrefix(types.SignatureListKey+signatureList.Id), b)
}

func (k Keeper) GetSignatureList(ctx sdk.Context, key string) types.SignatureList {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListKey))
	var signatureList types.SignatureList
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SignatureListKey+key)), &signatureList)
	return signatureList
}

func (k Keeper) HasSignatureList(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListKey))
	return store.Has(types.KeyPrefix(types.SignatureListKey + id))
}

// DeleteSignatureList deletes a signatureList
func (k Keeper) DeleteSignatureList(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListKey))
	store.Delete(types.KeyPrefix(types.SignatureListKey + key))
}

func (k Keeper) GetAllSignatureList(ctx sdk.Context) (msgs []types.SignatureList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.SignatureListKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.SignatureList
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
