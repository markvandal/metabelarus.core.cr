package record_update

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

type StatusSigned struct {
	StatusAbstract
}

func (status *StatusSigned) Dispatch(msg *types.MsgUpdateRecord) error {
	status.action = msg
	switch msg.Action {
	case types.RecordUpdate_RECORD_UPDATE_STORE:
		return sdkerrors.Wrap(types.ErrUpdateStore, "Can't update signed record")
	case types.RecordUpdate_RECORD_UPDATE_SIGN:
		return sdkerrors.Wrap(types.ErrUpdateSign, "Can't sign signed record")
	case types.RecordUpdate_REOCRD_UPDATE_SEAL:
		status.Seal()
		break
	case types.RecordUpdate_REOCRD_UPDATE_REJECT:
		status.Reject()
		break
	case types.RecordUpdate_REOCRD_UPDATE_WITHDRAW:
		status.Withdraw()
		break
	case types.RecordUpdate_REOCRD_UPDATE_REOPEN:
		err := status.CheckUpdate()
		if err != nil {
			return err
		}
		switch status.record.RecordType {
		case types.RecordType_IDENTITY_PERMANENT_RECORD:
		case types.RecordType_PROVIDER_PERMISSION:
			return sdkerrors.Wrap(
				types.ErrUpdateSign,
				fmt.Sprintf(
					"Can't reopen record %s",
					types.RecordType_name[int32(status.record.RecordType)],
				),
			)
		}
		status.record.Status = types.RecordStatus_RECORD_OPEN
		status.record.UpdateDt = status.action.UpdateDt
		status.RequireMutualUpdate()
		break
	default:
		return sdkerrors.Wrap(types.ErrUpdateAction, "Unkwnown action")
	}

	return nil
}
