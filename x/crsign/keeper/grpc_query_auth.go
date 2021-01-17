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

func (k Keeper) AuthAll(c context.Context, req *types.QueryAllAuthRequest) (*types.QueryAllAuthResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auths []*types.Auth
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	authStore := prefix.NewStore(store, types.KeyPrefix(types.AuthKey))

	pageRes, err := query.Paginate(authStore, req.Pagination, func(key []byte, value []byte) error {
		var auth types.Auth
		if err := k.cdc.UnmarshalBinaryBare(value, &auth); err != nil {
			return err
		}

		auths = append(auths, &auth)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAuthResponse{Auth: auths, Pagination: pageRes}, nil
}

func (k Keeper) Auth(c context.Context, req *types.QueryGetAuthRequest) (*types.QueryGetAuthResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auth types.Auth
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AuthKey + req.Id)), &auth)

	return &types.QueryGetAuthResponse{Auth: &auth}, nil
}
