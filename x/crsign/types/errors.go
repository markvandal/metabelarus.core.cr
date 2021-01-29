package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/crsign module sentinel errors
var (
	ErrSample       = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNoIdentity   = sdkerrors.Register(ModuleName, 1200, "no identity")
	ErrIdNotService = sdkerrors.Register(ModuleName, 1210, "requesting auth identity isn't the service")
	ErrAuthDuration = sdkerrors.Register(ModuleName, 1300, "bad auth duration")
)
