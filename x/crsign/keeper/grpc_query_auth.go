package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Auth(c context.Context, req *types.QueryGetAuthRequest) (*types.QueryGetAuthResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auth types.Auth
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AuthKey+req.Id)), &auth)

	return &types.QueryGetAuthResponse{Auth: &auth}, nil
}

func (k Keeper) Id2Service(c context.Context, req *types.QueryGetId2ServiceRequest) (*types.QueryGetId2ServiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var id2service types.Id2Service
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2ServicesKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.Id2ServicesKey+req.Id)), &id2service)

	return &types.QueryGetId2ServiceResponse{Id2Service: &id2service}, nil
}
