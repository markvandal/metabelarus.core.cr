package record_update

import (
	"github.com/metabelarus/mbcorecr/x/crsign/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type UpdateStatus interface {
	Dispatch(msg *types.MsgUpdateRecord) error
}

type StatusAbstract struct {
	record                 *types.Record
	action                 *types.MsgUpdateRecord
	actor                  string
	_providerUpdateChecked bool
	providerUpdate         bool
}

func (status *StatusAbstract) UpdateData() {
	status.record.Data = status.action.Data
	status.record.Signature = status.action.Signature
	status.record.UpdateDt = status.action.UpdateDt
}

func (status *StatusAbstract) CheckUpdate() error {
	if status._providerUpdateChecked {
		return nil
	}

	status._providerUpdateChecked = true
	status.providerUpdate = false

	if (status.action.Data != status.record.Data) != (status.action.Signature == "") {
		return sdkerrors.Wrap(types.ErrUpdateData, "Data is changed without signature")
	}

	// if types.RecordType_IDENTITY_PERMANENT_RECORD == status.record.RecordType {
	// 	return sdkerrors.Wrap(types.ErrUpdateImmutable, "Can't update data in permanent record")
	// }

	if types.IsIdentityRecord(status.record.RecordType) {
		if status.record.Identity != status.actor {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect updater of self record")
		}
	} else if status.record.Identity == status.actor {
		status.providerUpdate = false
	} else if status.record.Provider == status.actor {
		status.providerUpdate = true
	} else {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect updater")
	}

	if !status.providerUpdate {
		if status.record.RecordType == types.RecordType_PROVIDER_RECORD ||
			status.record.IsChildRecord() {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Impossible to change provided record by owner")
		}
	}

	return nil
}

func CreateUpdateStatus(record *types.Record, actor string) (UpdateStatus, error) {
	var ret UpdateStatus

	abstractStatus := StatusAbstract{
		record: record,
		actor:  actor,
	}
	switch record.Status {
	case types.RecordStatus_RECORD_OPEN:
		ret = &StatusOpen{StatusAbstract: abstractStatus}
		break
	case types.RecordStatus_RECORD_SIGNED:
		break
	case types.RecordStatus_RECORD_WITHDRAWN:
		break
	case types.RecordStatus_RECORD_REJECTED:
		break
	case types.RecordStatus_RECORD_SEALED:
		break
	default:
		return nil, sdkerrors.Wrap(types.ErrUpdateAction, "Unkwnown status")
	}

	return ret, nil
}
