package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateAllowance creates a allowance
func (k Keeper) CreateAllowance(ctx sdk.Context, allowance types.Allowance) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AllowancePrefix + allowance.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(allowance)
	store.Set(key, value)
}

// GetAllowance returns the allowance information
func (k Keeper) GetAllowance(ctx sdk.Context, key string) (types.Allowance, error) {
	store := ctx.KVStore(k.storeKey)
	var allowance types.Allowance
	byteKey := []byte(types.AllowancePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &allowance)
	if err != nil {
		return allowance, err
	}
	return allowance, nil
}

// SetAllowance sets a allowance
func (k Keeper) SetAllowance(ctx sdk.Context, allowance types.Allowance) {
	allowanceKey := allowance.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(allowance)
	key := []byte(types.AllowancePrefix + allowanceKey)
	store.Set(key, bz)
}

// DeleteAllowance deletes a allowance
func (k Keeper) DeleteAllowance(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.AllowancePrefix + key))
}

//
// Functions used by querier
//

func listAllowance(ctx sdk.Context, k Keeper) ([]byte, error) {
	var allowanceList []types.Allowance
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.AllowancePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var allowance types.Allowance
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &allowance)
		allowanceList = append(allowanceList, allowance)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, allowanceList)
	return res, nil
}

func getAllowance(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	allowance, err := k.GetAllowance(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, allowance)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetAllowanceOwner(ctx sdk.Context, key string) sdk.AccAddress {
	allowance, err := k.GetAllowance(ctx, key)
	if err != nil {
		return nil
	}
	return allowance.Creator
}


// Check if the key exists in the store
func (k Keeper) AllowanceExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.AllowancePrefix + key))
}
