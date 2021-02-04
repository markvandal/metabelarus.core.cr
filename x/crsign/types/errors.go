package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/crsign module sentinel errors
var (
	ErrSample              = sdkerrors.Register(ModuleName, 2100, "sample error")
	ErrNoIdentity          = sdkerrors.Register(ModuleName, 2200, "no identity")
	ErrIdNotService        = sdkerrors.Register(ModuleName, 2210, "requesting auth identity isn't the service")
	ErrAuthDuration        = sdkerrors.Register(ModuleName, 2300, "bad auth duration")
	ErrDateIssue           = sdkerrors.Register(ModuleName, 2400, "date issues")
	ErrRecIdentityProvider = sdkerrors.Register(ModuleName, 2500, "For self identity records provider and identity should match and otherwise")
	ErrNoRecord            = sdkerrors.Register(ModuleName, 2600, "No such record")
	ErrUpdateAction        = sdkerrors.Register(ModuleName, 2700, "Bad action")
	ErrUpdateImmutable     = sdkerrors.Register(ModuleName, 2800, "Can't update immutable record")
	ErrUpdateSeal          = sdkerrors.Register(ModuleName, 2710, "Can't seal record")
	ErrUpdateData          = sdkerrors.Register(ModuleName, 2900, "Can't update data without setting signature")
	ErrUpdateSign          = sdkerrors.Register(ModuleName, 2720, "Can't sign record")
	ErrUpdateStore         = sdkerrors.Register(ModuleName, 2730, "Can't store in record")
	ErrUpdateCancel        = sdkerrors.Register(ModuleName, 2740, "Can't perform action with canceled record")
	ErrDelete              = sdkerrors.Register(ModuleName, 2810, "Can't delete record")
)
