package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Record struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    IdentityId string `json:"IdentityId" yaml:"IdentityId"`
    ServiceId string `json:"ServiceId" yaml:"ServiceId"`
    ServiceType string `json:"ServiceType" yaml:"ServiceType"`
    Key string `json:"Key" yaml:"Key"`
    UserValue string `json:"UserValue" yaml:"UserValue"`
    ServiceValue string `json:"ServiceValue" yaml:"ServiceValue"`
    CreationDt string `json:"CreationDt" yaml:"CreationDt"`
    UpdateDt string `json:"UpdateDt" yaml:"UpdateDt"`
}