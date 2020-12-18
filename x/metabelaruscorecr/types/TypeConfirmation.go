package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Confirmation struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    IdenitityID string `json:"idenitityID" yaml:"idenitityID"`
    CreationDate string `json:"creationDate" yaml:"creationDate"`
    ExpirationDate string `json:"expirationDate" yaml:"expirationDate"`
    ConfirmatorID string `json:"confirmatorID" yaml:"confirmatorID"`
    CenterGeo string `json:"centerGeo" yaml:"centerGeo"`
    Status string `json:"status" yaml:"status"`
    NextTryDate string `json:"nextTryDate" yaml:"nextTryDate"`
}