package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateInvitation creates a invitation
func (k Keeper) CreateInvitation(ctx sdk.Context, invitation types.Invitation) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.InvitationPrefix + invitation.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(invitation)
	store.Set(key, value)
}

// GetInvitation returns the invitation information
func (k Keeper) GetInvitation(ctx sdk.Context, key string) (types.Invitation, error) {
	store := ctx.KVStore(k.storeKey)
	var invitation types.Invitation
	byteKey := []byte(types.InvitationPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &invitation)
	if err != nil {
		return invitation, err
	}
	return invitation, nil
}

// SetInvitation sets a invitation
func (k Keeper) SetInvitation(ctx sdk.Context, invitation types.Invitation) {
	invitationKey := invitation.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(invitation)
	key := []byte(types.InvitationPrefix + invitationKey)
	store.Set(key, bz)
}

// DeleteInvitation deletes a invitation
func (k Keeper) DeleteInvitation(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.InvitationPrefix + key))
}

//
// Functions used by querier
//

func listInvitation(ctx sdk.Context, k Keeper) ([]byte, error) {
	var invitationList []types.Invitation
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.InvitationPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var invitation types.Invitation
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &invitation)
		invitationList = append(invitationList, invitation)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, invitationList)
	return res, nil
}

func getInvitation(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	invitation, err := k.GetInvitation(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, invitation)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetInvitationOwner(ctx sdk.Context, key string) sdk.AccAddress {
	invitation, err := k.GetInvitation(ctx, key)
	if err != nil {
		return nil
	}
	return invitation.Creator
}


// Check if the key exists in the store
func (k Keeper) InvitationExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.InvitationPrefix + key))
}
