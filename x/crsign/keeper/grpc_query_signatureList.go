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

func (k Keeper) SignatureListAll(c context.Context, req *types.QueryAllSignatureListRequest) (*types.QueryAllSignatureListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signatureLists []*types.SignatureList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	signatureListStore := prefix.NewStore(store, types.KeyPrefix(types.SignatureListKey))

	pageRes, err := query.Paginate(signatureListStore, req.Pagination, func(key []byte, value []byte) error {
		var signatureList types.SignatureList
		if err := k.cdc.UnmarshalBinaryBare(value, &signatureList); err != nil {
			return err
		}

		signatureLists = append(signatureLists, &signatureList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignatureListResponse{SignatureList: signatureLists, Pagination: pageRes}, nil
}

func (k Keeper) SignatureList(c context.Context, req *types.QueryGetSignatureListRequest) (*types.QueryGetSignatureListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signatureList types.SignatureList
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureListKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SignatureListKey + req.Id)), &signatureList)

	return &types.QueryGetSignatureListResponse{SignatureList: &signatureList}, nil
}
