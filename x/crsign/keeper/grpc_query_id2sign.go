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

func (k Keeper) Id2SignAll(c context.Context, req *types.QueryAllId2SignRequest) (*types.QueryAllId2SignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var id2signs []*types.Id2Sign
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	id2signStore := prefix.NewStore(store, types.KeyPrefix(types.Id2SignKey))

	pageRes, err := query.Paginate(id2signStore, req.Pagination, func(key []byte, value []byte) error {
		var id2sign types.Id2Sign
		if err := k.cdc.UnmarshalBinaryBare(value, &id2sign); err != nil {
			return err
		}

		id2signs = append(id2signs, &id2sign)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllId2SignResponse{Id2Sign: id2signs, Pagination: pageRes}, nil
}

func (k Keeper) Id2Sign(c context.Context, req *types.QueryGetId2SignRequest) (*types.QueryGetId2SignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var id2sign types.Id2Sign
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2SignKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.Id2SignKey + req.Id)), &id2sign)

	return &types.QueryGetId2SignResponse{Id2Sign: &id2sign}, nil
}
