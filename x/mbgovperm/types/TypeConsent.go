package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Consent struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    ExtserviceId string `json:"extserviceId" yaml:"extserviceId"`
    PassportId string `json:"passportId" yaml:"passportId"`
    Resolution string `json:"resolution" yaml:"resolution"`
}