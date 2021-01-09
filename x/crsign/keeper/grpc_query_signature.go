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

func (k Keeper) SignatureAll(c context.Context, req *types.QueryAllSignatureRequest) (*types.QueryAllSignatureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signatures []*types.Signature
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signatureStore := prefix.NewStore(store, types.KeyPrefix(types.SignatureKey))

	pageRes, err := query.Paginate(signatureStore, req.Pagination, func(key []byte, value []byte) error {
		var signature types.Signature
		if err := k.cdc.UnmarshalBinaryBare(value, &signature); err != nil {
			return err
		}

		signatures = append(signatures, &signature)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignatureResponse{Signature: signatures, Pagination: pageRes}, nil
}

func (k Keeper) Signature(c context.Context, req *types.QueryGetSignatureRequest) (*types.QueryGetSignatureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signature types.Signature
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SignatureKey + req.Id)), &signature)

	return &types.QueryGetSignatureResponse{Signature: &signature}, nil
}
