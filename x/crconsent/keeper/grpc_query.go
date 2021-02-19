package keeper

import (
	"github.com/metabelarus/mbcorecr/x/crconsent/types"
)

var _ types.QueryServer = Keeper{}
