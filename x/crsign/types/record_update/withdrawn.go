package record_update

import (
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

type StatusWithdrawn struct {
	StatusAbstract
}

func (status *StatusWithdrawn) Dispatch(msg *types.MsgUpdateRecord) error {
	return status.DispatchCanceled(msg)
}
