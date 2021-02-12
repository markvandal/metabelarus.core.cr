package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Invite(c context.Context, req *types.QueryGetInviteRequest) (*types.QueryGetInviteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var invite types.Invite
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.InviteKey+req.Id)), &invite)

	return &types.QueryGetInviteResponse{Invite: &invite}, nil
}
