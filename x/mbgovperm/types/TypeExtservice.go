package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Extservice struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Alias string `json:"alias" yaml:"alias"`
    DetailsUrl string `json:"detailsUrl" yaml:"detailsUrl"`
}