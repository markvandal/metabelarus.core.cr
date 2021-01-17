package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/crsign module sentinel errors
var (
	ErrSample       = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNoIdentity   = sdkerrors.Register(ModuleName, 1200, "no identity")
	ErrAuthDuration = sdkerrors.Register(ModuleName, 1300, "bad auth duration")
)
