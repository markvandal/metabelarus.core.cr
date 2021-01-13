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

func (k Keeper) InviteAll(c context.Context, req *types.QueryAllInviteRequest) (*types.QueryAllInviteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var invites []*types.Invite
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inviteStore := prefix.NewStore(store, types.KeyPrefix(types.InviteKey))

	pageRes, err := query.Paginate(inviteStore, req.Pagination, func(key []byte, value []byte) error {
		var invite types.Invite
		if err := k.cdc.UnmarshalBinaryBare(value, &invite); err != nil {
			return err
		}

		invites = append(invites, &invite)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInviteResponse{Invite: invites, Pagination: pageRes}, nil
}

func (k Keeper) Invite(c context.Context, req *types.QueryGetInviteRequest) (*types.QueryGetInviteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var invite types.Invite
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InviteKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.InviteKey + req.Id)), &invite)

	return &types.QueryGetInviteResponse{Invite: &invite}, nil
}
