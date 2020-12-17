package keeper

// TODO: Define if your module needs Parameters, if not this can be deleted

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpasstrust/types"
)

// GetParams returns the total set of mbpasstrust parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the mbpasstrust parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}

