package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Signature(c context.Context, req *types.QueryGetSignatureRequest) (*types.QueryGetSignatureResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signature types.Signature
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignatureKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.SignatureKey+req.Id)), &signature)

	return &types.QueryGetSignatureResponse{Signature: &signature}, nil
}
