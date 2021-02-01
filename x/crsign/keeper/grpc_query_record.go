package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Record(c context.Context, req *types.QueryGetRecordRequest) (*types.QueryGetRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	record := k.GetRecord(sdk.UnwrapSDKContext(c), req.Id)

	return &types.QueryGetRecordResponse{Record: &record}, nil
}

func (k Keeper) Id2Record(c context.Context, req *types.QueryGetId2RecordRequest) (*types.QueryGetId2RecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	id2record := k.GetRecordsFromId(sdk.UnwrapSDKContext(c), req.Id)

	return &types.QueryGetId2RecordResponse{Id2Record: &id2record}, nil
}
