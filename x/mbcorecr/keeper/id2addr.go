package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

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
