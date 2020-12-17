package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Allowance struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    PassportId string `json:"passportId" yaml:"passportId"`
    AllowanceType string `json:"allowanceType" yaml:"allowanceType"`
    Resolution string `json:"resolution" yaml:"resolution"`
}