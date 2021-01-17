package keeper

import (
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

var _ types.QueryServer = Keeper{}
