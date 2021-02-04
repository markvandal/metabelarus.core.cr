package record_update

import (
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

type StatusRejected struct {
	StatusAbstract
}

func (status *StatusRejected) Dispatch(msg *types.MsgUpdateRecord) error {
	return status.DispatchCanceled(msg)
}
