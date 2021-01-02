package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetInviteCount get the total number of invite
func (k Keeper) GetInviteCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteCountKey))
	byteKey := types.KeyPrefix(types.InviteCountKey)
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

// SetInviteCount set the total number of invite
func (k Keeper) SetInviteCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteCountKey))
	byteKey := types.KeyPrefix(types.InviteCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateInvite(ctx sdk.Context, msg types.MsgCreateInvite) {
	// Create the invite
    count := k.GetInviteCount(ctx)
    var invite = types.Invite{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Inviter: msg.Inviter,
        Invitee: msg.Invitee,
        Level: msg.Level,
        Key: msg.Key,
        CreationDt: msg.CreationDt,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
    key := types.KeyPrefix(types.InviteKey + invite.Id)
    value := k.cdc.MustMarshalBinaryBare(&invite)
    store.Set(key, value)

    // Update invite count
    k.SetInviteCount(ctx, count+1)
}

func (k Keeper) UpdateInvite(ctx sdk.Context, invite types.Invite) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
	b := k.cdc.MustMarshalBinaryBare(&invite)
	store.Set(types.KeyPrefix(types.InviteKey + invite.Id), b)
}

func (k Keeper) GetInvite(ctx sdk.Context, key string) types.Invite {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
	var invite types.Invite
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.InviteKey + key)), &invite)
	return invite
}

func (k Keeper) HasInvite(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
	return store.Has(types.KeyPrefix(types.InviteKey + id))
}

func (k Keeper) GetInviteOwner(ctx sdk.Context, key string) string {
    return k.GetInvite(ctx, key).Creator
}

// DeleteInvite deletes a invite
func (k Keeper) DeleteInvite(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
	store.Delete(types.KeyPrefix(types.InviteKey + key))
}

func (k Keeper) GetAllInvite(ctx sdk.Context) (msgs []types.Invite) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.InviteKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Invite
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
