package record_update

import (
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

type StatusWithdrawn struct {
	StatusAbstract
}

func (status *StatusWithdrawn) Dispatch(msg *types.MsgUpdateRecord) error {
	status.action = msg
	switch msg.Action {
	case types.RecordUpdate_RECORD_UPDATE_STORE:

		break
	case types.RecordUpdate_RECORD_UPDATE_SIGN:

		break
	case types.RecordUpdate_REOCRD_UPDATE_SEAL:

		break
	case types.RecordUpdate_REOCRD_UPDATE_REJECT:

		break
	case types.RecordUpdate_REOCRD_UPDATE_WITHDRAW:

		break
	case types.RecordUpdate_REOCRD_UPDATE_REOPEN:

		break
	default:

		break
	}

	return nil
}
