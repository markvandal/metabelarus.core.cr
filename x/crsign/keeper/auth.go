package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func (k Keeper) CreateAuth(ctx sdk.Context, auth types.Auth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	key := types.KeyPrefix(types.AuthKey + auth.GetId())
	value := k.cdc.MustMarshalBinaryBare(&auth)
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
	auth := types.Auth{}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	k.cdc.MustUnmarshalBinaryBare(
		store.Get(types.KeyPrefix(types.AuthKey+identity+"."+service)),
		&auth,
	)

	return auth
}

func (k Keeper) CheckAuthorization(ctx sdk.Context, service string, identity string) error {
	auth := k.GetAuth(ctx, service, identity)
	if auth.Status != types.AuthStatus_AUTH_SIGNED {
		return sdkerrors.Wrap(
			sdkerrors.ErrUnauthorized,
			"Provider isn't authorized to create a record for identity",
		)
	}

	now := mbutils.CreateCurrentTime()
	if now.After(*auth.AvailabilityDt) {
		return sdkerrors.Wrap(
			types.ErrAuthDuration,
			"Provider authentication is expired",
		)
	}

	return nil
}

func (k Keeper) HasAuth(ctx sdk.Context, key string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	return store.Has(types.KeyPrefix(types.AuthKey + key))
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
