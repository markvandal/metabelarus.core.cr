package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/mbgovperm/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateConsent creates a consent
func (k Keeper) CreateConsent(ctx sdk.Context, consent types.Consent) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ConsentPrefix + consent.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(consent)
	store.Set(key, value)
}

// GetConsent returns the consent information
func (k Keeper) GetConsent(ctx sdk.Context, key string) (types.Consent, error) {
	store := ctx.KVStore(k.storeKey)
	var consent types.Consent
	byteKey := []byte(types.ConsentPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &consent)
	if err != nil {
		return consent, err
	}
	return consent, nil
}

// SetConsent sets a consent
func (k Keeper) SetConsent(ctx sdk.Context, consent types.Consent) {
	consentKey := consent.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(consent)
	key := []byte(types.ConsentPrefix + consentKey)
	store.Set(key, bz)
}

// DeleteConsent deletes a consent
func (k Keeper) DeleteConsent(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ConsentPrefix + key))
}

//
// Functions used by querier
//

func listConsent(ctx sdk.Context, k Keeper) ([]byte, error) {
	var consentList []types.Consent
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ConsentPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var consent types.Consent
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &consent)
		consentList = append(consentList, consent)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, consentList)
	return res, nil
}

func getConsent(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	consent, err := k.GetConsent(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, consent)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetConsentOwner(ctx sdk.Context, key string) sdk.AccAddress {
	consent, err := k.GetConsent(ctx, key)
	if err != nil {
		return nil
	}
	return consent.Creator
}


// Check if the key exists in the store
func (k Keeper) ConsentExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ConsentPrefix + key))
}
