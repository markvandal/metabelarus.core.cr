package keeper

import (
	// this line is used by starport scaffolding # 1

	"github.com/metabelarus/mbcorecr/x/crsign/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		// this line is used by starport scaffolding # 1
		case types.QueryGetAuth:
			return getAuth(ctx, path[1], k, legacyQuerierCdc)
		case types.QueryGetSignature:
			return getSignature(ctx, path[1], k, legacyQuerierCdc)
		case types.QueryGetId2Service:
			return getId2Service(ctx, path[1], k, legacyQuerierCdc)

		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}
