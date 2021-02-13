package record_update

import (
	"fmt"

	"github.com/metabelarus/mbcorecr/x/crsign/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type UpdateStatus interface {
	Dispatch(msg *types.MsgUpdateRecord) error
	IsChildUpdate() bool
	CheckUpdate() error
}

type StatusAbstract struct {
	record                 *types.Record
	action                 *types.MsgUpdateRecord
	actor                  string
	_providerUpdateChecked bool
	providerUpdate         bool
	IsParentUpdateRequired bool
	IsChildUpdateRequired  bool
}

func (status *StatusAbstract) UpdateData() {
	status.record.Data = status.action.Data
	status.record.Signature = status.action.Signature
	status.record.UpdateDt = status.action.UpdateDt
}

func (status *StatusAbstract) Seal() error {
	switch status.record.RecordType {
	default:
		return sdkerrors.Wrap(types.ErrUpdateSeal, "Can't seal open record")
	case types.RecordType_PROVIDER_RECORD:
	case types.RecordType_IDENTITY_RECORD:
	}
	status.record.Status = types.RecordStatus_RECORD_SEALED
	status.record.UpdateDt = status.action.UpdateDt

	return nil
}

func (status *StatusAbstract) Reject() error {
	err := status.CheckReject()
	if err != nil {
		return err
	}
	status.Cancel(types.RecordStatus_RECORD_REJECTED)

	return nil
}

func (status *StatusAbstract) Withdraw() error {
	err := status.CheckWithdraw()
	if err != nil {
		return err
	}
	status.Cancel(types.RecordStatus_RECORD_WITHDRAWN)

	return nil
}

func (status *StatusAbstract) Cancel(newStatus types.RecordStatus) {
	status.record.Data = ""
	status.record.Signature = ""
	status.record.Status = newStatus
	status.record.UpdateDt = status.action.UpdateDt
}

func (status *StatusAbstract) DispatchCanceled(msg *types.MsgUpdateRecord) error {
	status.action = msg
	switch msg.Action {
	case types.RecordUpdate_REOCRD_UPDATE_REOPEN:
		err := status.CheckUpdate()
		if err != nil {
			return err
		}
		status.record.Status = types.RecordStatus_RECORD_OPEN
		status.record.UpdateDt = status.action.UpdateDt
	default:
		return sdkerrors.Wrap(
			types.ErrUpdateCancel,
			fmt.Sprintf(
				"Action: %s",
				types.RecordUpdate_name[int32(status.action.Action)],
			),
		)
	}

	return nil
}

func (status *StatusAbstract) CheckUpdate() error {
	if status._providerUpdateChecked {
		return nil
	}

	status._providerUpdateChecked = true
	status.providerUpdate = false

	if status.record.Publicity == types.PublicityType_PRIVATE {
		if status.action != nil && (status.action.Data != status.record.Data) && (status.action.Signature == "") {
			return sdkerrors.Wrap(types.ErrUpdateData, "Data is changed without signature")
		}
	}

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

	return nil
}

func (status *StatusAbstract) CheckReject() error {
	if types.IsIdentityRecord(status.record.RecordType) {
		if status.record.Identity != status.actor {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only personal record can be rejected")
		}
	}

	return nil
}

func (status *StatusAbstract) CheckWithdraw() error {
	if types.IsProviderRecord(status.record.RecordType) {
		if status.record.Provider != status.actor {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only provider record can be withdrawal")
		}
	}

	return nil
}

func (status *StatusAbstract) IsChildUpdate() bool {
	return status.IsChildUpdateRequired
}

func CreateUpdateStatus(record *types.Record, actor string) (UpdateStatus, error) {
	var ret UpdateStatus

	abstractStatus := StatusAbstract{
		record:                 record,
		actor:                  actor,
		IsParentUpdateRequired: false,
		IsChildUpdateRequired:  false,
	}
	switch record.Status {
	case types.RecordStatus_RECORD_OPEN:
		ret = &StatusOpen{StatusAbstract: abstractStatus}
	case types.RecordStatus_RECORD_SIGNED:
		ret = &StatusSigned{StatusAbstract: abstractStatus}
	case types.RecordStatus_RECORD_WITHDRAWN:
		ret = &StatusWithdrawn{StatusAbstract: abstractStatus}
	case types.RecordStatus_RECORD_REJECTED:
		ret = &StatusRejected{StatusAbstract: abstractStatus}
	case types.RecordStatus_RECORD_SEALED:
		ret = &StatusSealed{StatusAbstract: abstractStatus}
	default:
		return nil, sdkerrors.Wrap(types.ErrUpdateAction, "Unkwnown status")
	}

	return ret, nil
}
