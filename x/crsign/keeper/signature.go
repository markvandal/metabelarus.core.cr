package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// GetSignatureCount get the total number of signature
func (k Keeper) GetSignatureCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureCountKey))
	byteKey := types.KeyPrefix(types.SignatureCountKey)
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

// SetSignatureCount set the total number of signature
func (k Keeper) SetSignatureCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureCountKey))
	byteKey := types.KeyPrefix(types.SignatureCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateSignature(ctx sdk.Context, msg types.MsgCreateSignature) {
	// Create the signature
	count := k.GetSignatureCount(ctx)
	var signature = types.Signature{
		Id:       strconv.FormatInt(count, 10),
		Identity: msg.Identity,
		Service:  msg.Service,
		// Key: msg.Key,
		// Secret: msg.Secret,
		// CreationDt: msg.CreationDt,
		// AvailabilityDt: msg.AvailabilityDt,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	key := types.KeyPrefix(types.SignatureKey + signature.Id)
	value := k.cdc.MustMarshalBinaryBare(&signature)
	store.Set(key, value)

	// Update signature count
	k.SetSignatureCount(ctx, count+1)
}

func (k Keeper) UpdateSignature(ctx sdk.Context, signature types.Signature) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	b := k.cdc.MustMarshalBinaryBare(&signature)
	store.Set(types.KeyPrefix(types.SignatureKey+signature.Id), b)
}

func (k Keeper) GetSignature(ctx sdk.Context, key string) types.Signature {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	var signature types.Signature
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SignatureKey+key)), &signature)
	return signature
}

func (k Keeper) HasSignature(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	return store.Has(types.KeyPrefix(types.SignatureKey + id))
}

func (k Keeper) GetSignatureOwner(ctx sdk.Context, key string) string {
	return k.GetSignature(ctx, key).Service
}

// DeleteSignature deletes a signature
func (k Keeper) DeleteSignature(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	store.Delete(types.KeyPrefix(types.SignatureKey + key))
}

func (k Keeper) GetAllSignature(ctx sdk.Context) (msgs []types.Signature) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.SignatureKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Signature
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
