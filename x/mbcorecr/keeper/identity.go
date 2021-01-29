package keeper

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

// GetIdentityCount get the total number of identity
func (k Keeper) GetIdentityCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityCountKey))
	byteKey := types.KeyPrefix(types.IdentityCountKey)
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

// SetIdentityCount set the total number of identity
func (k Keeper) SetIdentityCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityCountKey))
	byteKey := types.KeyPrefix(types.IdentityCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) GetAddressesFromId(ctx sdk.Context, id string) types.Id2Addr {
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.IdToAddrKey),
	)
	var id2addr types.Id2Addr
	k.cdc.MustUnmarshalBinaryBare(
		store.Get(
			types.KeyPrefix(types.IdToAddrKey+id),
		),
		&id2addr,
	)
	return id2addr
}

func (k Keeper) EnsureIdFromAddress(ctx sdk.Context, address string, creationDt *time.Time) (string, error) {
	id := k.GetIdFromAddress(ctx, address)

	if id != "" {
		return id, nil
	}

	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrEnsureId, "account address is wrong")
	}

	if k.BankKeeper.HasBalance(
		ctx, accAddress,
		sdk.NewCoin(types.SuperInviteDenom, sdk.NewInt(1)),
	) {
		id = k.CreateIdentity(ctx, types.Identity{
			IdentityType: types.IdentityType_CITIZEN,
			InvitationId: "",
			CreationDt:   creationDt,
		}, address)
	} else {
		return "", sdkerrors.Wrap(types.ErrEnsureId, "this account can't have identity")
	}

	return id, nil
}

func (k Keeper) GetIdFromAddress(ctx sdk.Context, address string) string {
	addr2IdBytes := prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddrToIdKey),
	).Get(types.KeyPrefix(address))

	var addr2id types.Addr2Id
	k.cdc.MustUnmarshalBinaryBare(addr2IdBytes, &addr2id)

	return addr2id.Id
}

func (k Keeper) GetAddressFromId(ctx sdk.Context, id string) string {
	id2addr := k.GetAddressesFromId(ctx, id)
	for _, address := range id2addr.Addresses {
		if address.Main {
			return address.Address
		}
	}

	return ""
}

func (k Keeper) SetId2Addr(ctx sdk.Context, id string, addr types.Addr) {
	id2addr := k.GetAddressesFromId(ctx, id)
	var currentKey int32
	var currentMain int32
	for key, address := range id2addr.Addresses {
		if address.Address == addr.Address {
			currentKey = key
		}
		if address.Main {
			currentMain = key
		}
	}

	if _, ok := id2addr.Addresses[currentKey]; !ok {
		currentKey = int32(len(id2addr.Addresses))
	}

	if len(id2addr.Addresses) == 0 {
		id2addr.Addresses = make(map[int32]*types.Addr)
	}

	id2addr.Addresses[currentKey] = &addr
	if existingMain, ok := id2addr.Addresses[currentMain]; ok && addr.Main && currentMain != currentKey {
		existingMain.Main = false
	}

	prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdToAddrKey),
	).Set(
		types.KeyPrefix(types.IdToAddrKey+id),
		k.cdc.MustMarshalBinaryBare(&id2addr),
	)

	prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddrToIdKey),
	).Set(
		types.KeyPrefix(addr.Address),
		k.cdc.MustMarshalBinaryBare(&types.Addr2Id{Id: id}),
	)
}

func (k Keeper) CreateIdentity(ctx sdk.Context, identity types.Identity, address ...string) string {
	// Create the identity
	count := k.GetIdentityCount(ctx)
	identity.Id = strconv.FormatInt(count, 10)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	key := types.KeyPrefix(types.IdentityKey + identity.Id)
	value := k.cdc.MustMarshalBinaryBare(&identity)
	store.Set(key, value)

	// Update identity count
	k.SetIdentityCount(ctx, count+1)

	if len(address) > 0 {
		k.SetId2Addr(ctx, identity.Id, types.Addr{
			Address: address[0],
			Main:    true},
		)
	}

	return identity.Id
}

func (k Keeper) UpdateIdentity(ctx sdk.Context, identity types.Identity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	b := k.cdc.MustMarshalBinaryBare(&identity)
	store.Set(types.KeyPrefix(types.IdentityKey+identity.Id), b)
}

func (k Keeper) GetIdentity(ctx sdk.Context, id string) types.Identity {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	var identity types.Identity
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.IdentityKey+id)), &identity)
	return identity
}

func (k Keeper) ExportIdentity(ctx sdk.Context, id string) types.IdentityI {
	return k.GetIdentity(ctx, id)
}

func (k Keeper) HasIdentity(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	return store.Has(types.KeyPrefix(types.IdentityKey + id))
}

func (k Keeper) GetIdentityOwner(ctx sdk.Context, id string) string {
	return k.GetAddressFromId(ctx, id)
}

// DeleteIdentity deletes a identity
func (k Keeper) DeleteIdentity(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	store.Delete(types.KeyPrefix(types.IdentityKey + key))
}

func (k Keeper) GetAllIdentity(ctx sdk.Context) (msgs []types.Identity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.IdentityKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Identity
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
