package record_update

import (
	"github.com/metabelarus/mbcorecr/x/crsign/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type UpdateStatus interface {
	Dispatch(msg *types.MsgUpdateRecord) error
	IsMutualUpdateRequired() bool
	IsChildUpdate() bool
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
	case types.RecordType_IDENTITY_PERMANENT_RECORD:
	}
	status.record.Status = types.RecordStatus_RECORD_SEALED
	status.record.UpdateDt = status.action.UpdateDt

	status.RequireMutualUpdate()

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

	status.RequireMutualUpdate()
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

	err := status.CheckMutualRecordChangability()
	if err != nil {
		return err
	}

	return nil
}

func (status *StatusAbstract) CheckMutualRecordChangability() error {
	if status.providerUpdate {
		if status.record.IsChildRecord() &&
			status.record.RecordType == types.RecordType_IDENTITY_RECORD {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Impossible to change identity record by provider")
		}
	} else {
		if status.record.IsChildRecord() &&
			status.record.RecordType == types.RecordType_PROVIDER_RECORD {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Impossible to change provider record by identity")
		}
	}

	return nil
}

func (status *StatusAbstract) CheckReject() error {
	if types.IsIdentityRecord(status.record.RecordType) {
		if status.record.Identity != status.actor {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only personal record can be rejected")
		}
	}

	err := status.CheckMutualRecordChangability()
	if err != nil {
		return err
	}

	return nil
}

func (status *StatusAbstract) CheckWithdraw() error {
	if types.IsProviderRecord(status.record.RecordType) {
		if status.record.Provider != status.actor {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only provider record can be withdrawal")
		}
	}

	err := status.CheckMutualRecordChangability()
	if err != nil {
		return err
	}

	return nil
}

func (status *StatusAbstract) RequireMutualUpdate() {
	if status.record.IsChildRecord() {
		status.IsParentUpdateRequired = true
	}
	if status.record.IsParentRecord() {
		status.IsChildUpdateRequired = true
	}
}

func (status *StatusAbstract) IsMutualUpdateRequired() bool {
	return status.IsParentUpdateRequired || status.IsChildUpdateRequired
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
