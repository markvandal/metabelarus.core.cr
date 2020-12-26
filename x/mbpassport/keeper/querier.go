package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
		
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
)

// NewQuerier creates a new querier for mbpassport clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListRecord:
			return listRecord(ctx, k)
		case types.QueryGetRecord:
			return getRecord(ctx, path[1:], k)
		case types.QueryParams:
			return queryParams(ctx, k)
			// TODO: Put the modules query routes
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown mbpassport query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
