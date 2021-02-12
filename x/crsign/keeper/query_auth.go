package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func getAuth(ctx sdk.Context, key string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	msg := keeper.GetAuthByKey(ctx, key)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func getId2Service(ctx sdk.Context, id string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	msg := keeper.GetServicesFromId(ctx, id)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
