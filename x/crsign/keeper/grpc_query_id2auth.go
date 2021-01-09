package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Id2AuthAll(c context.Context, req *types.QueryAllId2AuthRequest) (*types.QueryAllId2AuthResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var id2auths []*types.Id2Auth
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	id2authStore := prefix.NewStore(store, types.KeyPrefix(types.Id2AuthKey))

	pageRes, err := query.Paginate(id2authStore, req.Pagination, func(key []byte, value []byte) error {
		var id2auth types.Id2Auth
		if err := k.cdc.UnmarshalBinaryBare(value, &id2auth); err != nil {
			return err
		}

		id2auths = append(id2auths, &id2auth)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllId2AuthResponse{Id2Auth: id2auths, Pagination: pageRes}, nil
}

func (k Keeper) Id2Auth(c context.Context, req *types.QueryGetId2AuthRequest) (*types.QueryGetId2AuthResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var id2auth types.Id2Auth
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2AuthKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.Id2AuthKey + req.Id)), &id2auth)

	return &types.QueryGetId2AuthResponse{Id2Auth: &id2auth}, nil
}
