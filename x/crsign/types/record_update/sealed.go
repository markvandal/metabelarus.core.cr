package record_update

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

type StatusSealed struct {
	StatusAbstract
}

func (status *StatusSealed) Dispatch(msg *types.MsgUpdateRecord) error {
	status.action = msg
	switch msg.Action {
	case types.RecordUpdate_RECORD_UPDATE_STORE:
		return sdkerrors.Wrap(types.ErrUpdateStore, "Can't update a sealed record")
	case types.RecordUpdate_RECORD_UPDATE_SIGN:
		return sdkerrors.Wrap(types.ErrUpdateSign, "Can't sign a sealed record")
	case types.RecordUpdate_REOCRD_UPDATE_SEAL:
		return sdkerrors.Wrap(types.ErrUpdateSign, "Can't seal a sealed record")
	case types.RecordUpdate_REOCRD_UPDATE_REJECT:
		return sdkerrors.Wrap(types.ErrUpdateSign, "Can't reject a sealed record")
	case types.RecordUpdate_REOCRD_UPDATE_WITHDRAW:
		status.Withdraw()
	case types.RecordUpdate_REOCRD_UPDATE_REOPEN:
		err := status.CheckUpdate()
		if err != nil {
			return err
		}
		switch status.record.RecordType {
		default:
			return sdkerrors.Wrap(
				types.ErrUpdateSeal,
				fmt.Sprintf(
					"Can't reopen sealed record of type %s",
					types.RecordType_name[int32(status.record.RecordType)],
				),
			)
		case types.RecordType_PROVIDER_RECORD:
		case types.RecordType_PROVIDER_SIGNABLE_RECORD:
		}
		status.record.Status = types.RecordStatus_RECORD_OPEN
		status.record.UpdateDt = status.action.UpdateDt
	default:
		return sdkerrors.Wrap(types.ErrUpdateAction, "Unkwnown action")
	}

	return nil
}
