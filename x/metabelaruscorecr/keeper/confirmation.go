package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateConfirmation creates a confirmation
func (k Keeper) CreateConfirmation(ctx sdk.Context, confirmation types.Confirmation) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ConfirmationPrefix + confirmation.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(confirmation)
	store.Set(key, value)
}

// GetConfirmation returns the confirmation information
func (k Keeper) GetConfirmation(ctx sdk.Context, key string) (types.Confirmation, error) {
	store := ctx.KVStore(k.storeKey)
	var confirmation types.Confirmation
	byteKey := []byte(types.ConfirmationPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &confirmation)
	if err != nil {
		return confirmation, err
	}
	return confirmation, nil
}

// SetConfirmation sets a confirmation
func (k Keeper) SetConfirmation(ctx sdk.Context, confirmation types.Confirmation) {
	confirmationKey := confirmation.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(confirmation)
	key := []byte(types.ConfirmationPrefix + confirmationKey)
	store.Set(key, bz)
}

// DeleteConfirmation deletes a confirmation
func (k Keeper) DeleteConfirmation(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ConfirmationPrefix + key))
}

//
// Functions used by querier
//

func listConfirmation(ctx sdk.Context, k Keeper) ([]byte, error) {
	var confirmationList []types.Confirmation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ConfirmationPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var confirmation types.Confirmation
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &confirmation)
		confirmationList = append(confirmationList, confirmation)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, confirmationList)
	return res, nil
}

func getConfirmation(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	confirmation, err := k.GetConfirmation(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, confirmation)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetConfirmationOwner(ctx sdk.Context, key string) sdk.AccAddress {
	confirmation, err := k.GetConfirmation(ctx, key)
	if err != nil {
		return nil
	}
	return confirmation.Creator
}


// Check if the key exists in the store
func (k Keeper) ConfirmationExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ConfirmationPrefix + key))
}
