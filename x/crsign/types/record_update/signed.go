package record_update

import (
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
		switch status.record.RecordType {
		default:
		}
		break
	default:
		return sdkerrors.Wrap(types.ErrUpdateAction, "Unkwnown action")
	}

	return nil
}
