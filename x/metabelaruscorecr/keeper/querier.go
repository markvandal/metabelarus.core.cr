package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for metabelaruscorecr clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListConfirmation:
			return listConfirmation(ctx, k)
		case types.QueryGetConfirmation:
			return getConfirmation(ctx, path[1:], k)
		case types.QueryListInvitation:
			return listInvitation(ctx, k)
		case types.QueryGetInvitation:
			return getInvitation(ctx, path[1:], k)
		case types.QueryListIdentity:
			return listIdentity(ctx, k)
		case types.QueryGetIdentity:
			return getIdentity(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown metabelaruscorecr query endpoint")
		}
	}
}
