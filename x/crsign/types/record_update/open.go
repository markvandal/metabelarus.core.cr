package record_update

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

type StatusOpen struct {
	StatusAbstract
}

func (status *StatusOpen) Dispatch(msg *types.MsgUpdateRecord) error {
	status.action = msg
	switch msg.Action {
	case types.RecordUpdate_RECORD_UPDATE_STORE:
		err := status.CheckUpdate()
		if err != nil {
			return err
		}
		status.UpdateData()
	case types.RecordUpdate_RECORD_UPDATE_SIGN:
		err := status.CheckUpdate()
		if err != nil {
			return err
		}
		status.record.Status = types.RecordStatus_RECORD_SIGNED
		status.record.SignatureDt = status.action.UpdateDt
		status.record.UpdateDt = status.action.UpdateDt
	case types.RecordUpdate_REOCRD_UPDATE_SEAL:
		status.Seal()
	case types.RecordUpdate_REOCRD_UPDATE_REJECT:
		status.Reject()
	case types.RecordUpdate_REOCRD_UPDATE_WITHDRAW:
		status.Withdraw()
	case types.RecordUpdate_REOCRD_UPDATE_REOPEN:
		return sdkerrors.Wrap(types.ErrUpdateAction, "Can't reopen opened record")
	default:
		return sdkerrors.Wrap(types.ErrUpdateAction, "Unkwnown action")
	}

	return nil
}
