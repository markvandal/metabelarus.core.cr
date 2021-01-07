package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/mbcorecr module sentinel errors
var (
	ErrSample    = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrDateIssue = sdkerrors.Register(ModuleName, 1200, "Issue with a date")
)
