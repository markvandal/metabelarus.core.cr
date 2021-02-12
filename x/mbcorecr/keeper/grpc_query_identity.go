package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdentityAll(c context.Context, req *types.QueryAllIdentityRequest) (*types.QueryAllIdentityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var identitys []*types.Identity
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	identityStore := prefix.NewStore(store, types.KeyPrefix(types.IdentityKey))

	pageRes, err := query.Paginate(identityStore, req.Pagination, func(key []byte, value []byte) error {
		var identity types.Identity
		if err := k.cdc.UnmarshalBinaryBare(value, &identity); err != nil {
			return err
		}

		identitys = append(identitys, &identity)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIdentityResponse{Identity: identitys, Pagination: pageRes}, nil
}

func (k Keeper) Identity(c context.Context, req *types.QueryGetIdentityRequest) (*types.QueryGetIdentityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var identity types.Identity
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentityKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.IdentityKey+req.Id)), &identity)

	return &types.QueryGetIdentityResponse{Identity: &identity}, nil
}

func (k Keeper) AddrToId(c context.Context, req *types.QueryAddrToIdRequest) (*types.QueryAddrToIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addrToId := &types.Addr2Id{
		Id: k.GetIdFromAddress(sdk.UnwrapSDKContext(c), req.Address),
	}

	return &types.QueryAddrToIdResponse{Addr2Id: addrToId}, nil
}

func (k Keeper) IdToAddr(c context.Context, req *types.QueryIdToAddrRequest) (*types.QueryIdToAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr := &types.Addr{
		Address: k.GetAddressFromId(sdk.UnwrapSDKContext(c), req.Id),
	}

	return &types.QueryIdToAddrResponse{Addr: addr}, nil
}
